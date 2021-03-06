package netbox

import (
	"regexp"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	pkgerrors "github.com/pkg/errors"
	netboxclient "github.com/tomasherout/go-netbox/netbox/client"
	"github.com/tomasherout/go-netbox/netbox/client/ipam"
	"github.com/tomasherout/go-netbox/netbox/models"
)

func resourceNetboxIpamIPAddresses() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxIpamIPAddressesCreate,
		Read:   resourceNetboxIpamIPAddressesRead,
		Update: resourceNetboxIpamIPAddressesUpdate,
		Delete: resourceNetboxIpamIPAddressesDelete,
		Exists: resourceNetboxIpamIPAddressesExists,

		Schema: map[string]*schema.Schema{
			"address": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringMatch(
					regexp.MustCompile("^[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}/"+
						"[0-9]{1,2}$"), "Must be like 192.168.56.1/24"),
			},
			"description": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(1, 100),
			},
			"dns_name": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
				ValidateFunc: validation.StringMatch(
					regexp.MustCompile("^[-a-zA-Z0-9_.]{1,255}$"),
					"Must be like ^[-a-zA-Z0-9_.]{1,255}$"),
			},
			"interface_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"nat_inside_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"nat_outside_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"role": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
				ValidateFunc: validation.StringInSlice([]string{"loopback",
					"secondary", "anycast", "vip", "vrrp", "hsrp", "glbp", "carp"},
					false),
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "active",
				ValidateFunc: validation.StringInSlice([]string{"container", "active",
					"reserved", "deprecated", "dhcp"}, false),
			},
			"tags": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"tenant_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vrf_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceNetboxIpamIPAddressesCreate(d *schema.ResourceData,
	m interface{}) error {
	client := m.(*netboxclient.NetBoxAPI)

	address := d.Get("address").(string)
	description := d.Get("description").(string)
	dnsName := d.Get("dns_name").(string)
	interfaceID := int64(d.Get("interface_id").(int))
	natInsideID := int64(d.Get("nat_inside_id").(int))
	natOutsideID := int64(d.Get("nat_outside_id").(int))
	role := d.Get("role").(string)
	status := d.Get("status").(string)
	tags := d.Get("tags").(*schema.Set).List()
	tenantID := int64(d.Get("tenant_id").(int))
	vrfID := int64(d.Get("vrf_id").(int))

	newResource := &models.WritableIPAddress{
		Address:     &address,
		Description: description,
		DNSName:     dnsName,
		Role:        role,
		Status:      status,
		Tags:        expandToStringSlice(tags),
	}

	if interfaceID != 0 {
		dcimInterface := "dcim.interface"
		newResource.AssignedObjectType = &dcimInterface
		newResource.AssignedObjectID = &interfaceID
	}

	if natInsideID != 0 {
		newResource.NatInside = &natInsideID
	}

	if natOutsideID != 0 {
		newResource.NatOutside = &natOutsideID
	}

	if tenantID != 0 {
		newResource.Tenant = &tenantID
	}

	if vrfID != 0 {
		newResource.Vrf = &vrfID
	}

	resource := ipam.NewIpamIPAddressesCreateParams().WithData(newResource)

	resourceCreated, err := client.Ipam.IpamIPAddressesCreate(resource, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(resourceCreated.Payload.ID, 10))

	return resourceNetboxIpamIPAddressesRead(d, m)
}

func resourceNetboxIpamIPAddressesRead(d *schema.ResourceData,
	m interface{}) error {
	client := m.(*netboxclient.NetBoxAPI)

	resourceID := d.Id()
	params := ipam.NewIpamIPAddressesListParams().WithID(&resourceID)
	resources, err := client.Ipam.IpamIPAddressesList(params, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Payload.Results {
		if strconv.FormatInt(resource.ID, 10) == d.Id() {
			if err = d.Set("address", resource.Address); err != nil {
				return err
			}

			if err = d.Set("description", resource.Description); err != nil {
				return err
			}

			if err = d.Set("dns_name", resource.DNSName); err != nil {
				return err
			}

			if resource.AssignedObjectID == nil {
				if *resource.AssignedObjectType == "dcim.interface" {
					if err = d.Set("interface_id", resource.AssignedObjectID); err != nil {
						return err
					}
				}
			}

			if resource.NatInside == nil {
				if err = d.Set("nat_inside_id", nil); err != nil {
					return err
				}
			} else {
				if err = d.Set("nat_inside_id", resource.NatInside.ID); err != nil {
					return err
				}
			}

			if resource.NatOutside == nil {
				if err = d.Set("nat_outside_id", nil); err != nil {
					return err
				}
			} else {
				if err = d.Set("nat_outside_id", resource.NatOutside.ID); err != nil {
					return err
				}
			}

			if resource.Role == nil {
				if err = d.Set("role", nil); err != nil {
					return err
				}
			} else {
				if err = d.Set("role", resource.Role.Value); err != nil {
					return err
				}
			}

			if resource.Status == nil {
				if err = d.Set("status", nil); err != nil {
					return err
				}
			} else {
				if err = d.Set("status", resource.Status.Value); err != nil {
					return err
				}
			}

			if err = d.Set("tags", resource.Tags); err != nil {
				return err
			}

			if resource.Tenant == nil {
				if err = d.Set("tenant_id", nil); err != nil {
					return err
				}
			} else {
				if err = d.Set("tenant_id", resource.Tenant.ID); err != nil {
					return err
				}
			}

			if resource.Vrf == nil {
				if err = d.Set("vrf_id", nil); err != nil {
					return err
				}
			} else {
				if err = d.Set("vrf_id", resource.Vrf.ID); err != nil {
					return err
				}
			}

			return nil
		}
	}

	d.SetId("")

	return nil
}

func resourceNetboxIpamIPAddressesUpdate(d *schema.ResourceData,
	m interface{}) error {
	client := m.(*netboxclient.NetBoxAPI)
	params := &models.WritableIPAddress{}

	address := d.Get("address").(string)
	params.Address = &address

	if d.HasChange("description") {
		if description, exist := d.GetOk("description"); exist {
			params.Description = description.(string)
		} else {
			params.Description = " "
		}
	}

	if d.HasChange("dns_name") {
		params.DNSName = d.Get("dns_name").(string)
	}

	if d.HasChange("interface_id") {
		interfaceID := int64(d.Get("interface_id").(int))
		if interfaceID != 0 {
			dcimInterface := "dcim.interface"
			params.AssignedObjectType = &dcimInterface
			params.AssignedObjectID = &interfaceID
		}
	}

	if d.HasChange("nat_inside_id") {
		natInsideID := int64(d.Get("nat_inside_id").(int))
		if natInsideID != 0 {
			params.NatInside = &natInsideID
		}
	}

	if d.HasChange("nat_outside_id") {
		natInsideID := int64(d.Get("nat_outside_id").(int))
		if natInsideID != 0 {
			params.NatInside = &natInsideID
		}
	}

	if d.HasChange("role") {
		role := d.Get("role").(string)
		params.Role = role
	}

	if d.HasChange("status") {
		status := d.Get("status").(string)
		params.Status = status
	}

	tags := d.Get("tags").(*schema.Set).List()
	params.Tags = expandToStringSlice(tags)

	if d.HasChange("tenant_id") {
		tenantID := int64(d.Get("tenant_id").(int))
		if tenantID != 0 {
			params.Tenant = &tenantID
		}
	}

	if d.HasChange("vrf_id") {
		vrfID := int64(d.Get("vrf_id").(int))
		if vrfID != 0 {
			params.Vrf = &vrfID
		}
	}

	resource := ipam.NewIpamIPAddressesPartialUpdateParams().WithData(
		params)

	resourceID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return pkgerrors.New("Unable to convert ID into int64")
	}

	resource.SetID(resourceID)

	_, err = client.Ipam.IpamIPAddressesPartialUpdate(resource, nil)
	if err != nil {
		return err
	}

	return resourceNetboxIpamIPAddressesRead(d, m)
}

func resourceNetboxIpamIPAddressesDelete(d *schema.ResourceData,
	m interface{}) error {
	client := m.(*netboxclient.NetBoxAPI)

	resourceExists, err := resourceNetboxIpamIPAddressesExists(d, m)
	if err != nil {
		return err
	}

	if !resourceExists {
		return nil
	}

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return pkgerrors.New("Unable to convert ID into int64")
	}

	resource := ipam.NewIpamIPAddressesDeleteParams().WithID(id)
	if _, err := client.Ipam.IpamIPAddressesDelete(resource, nil); err != nil {
		return err
	}

	return nil
}

func resourceNetboxIpamIPAddressesExists(d *schema.ResourceData,
	m interface{}) (b bool, e error) {
	client := m.(*netboxclient.NetBoxAPI)
	resourceExist := false

	resourceID := d.Id()
	params := ipam.NewIpamIPAddressesListParams().WithID(&resourceID)
	resources, err := client.Ipam.IpamIPAddressesList(params, nil)
	if err != nil {
		return resourceExist, err
	}

	for _, resource := range resources.Payload.Results {
		if strconv.FormatInt(resource.ID, 10) == d.Id() {
			resourceExist = true
		}
	}

	return resourceExist, nil
}
