package netbox

import (
	"errors"
	"regexp"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	pkgerrors "github.com/pkg/errors"
	netboxclient "github.com/tomasherout/go-netbox/netbox/client"
	"github.com/tomasherout/go-netbox/netbox/client/ipam"
	"github.com/tomasherout/go-netbox/netbox/models"
)

func resourceNetboxIpamIPByPrefix() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxIpamIPByPrefixCreate,
		Read:   resourceNetboxIpamIPByPrefixRead,
		Update: resourceNetboxIpamIPByPrefixUpdate,
		Delete: resourceNetboxIpamIPByPrefixsDelete,

		Schema: map[string]*schema.Schema{
			"address": {
				Type:     schema.TypeString,
				Computed: true,
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
			"search_prefix_ids": {
				Type:     schema.TypeSet,
				Required: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
		},
	}
}

func resourceNetboxIpamIPByPrefixCreate(d *schema.ResourceData, m interface{}) error {

	client := m.(*netboxclient.NetBoxAPI)

	prefixIds := d.Get("search_prefix_ids").(*schema.Set)

	if prefixIds.Len() == 0 {
		return errors.New("search_prefix_ids nemůže být předáno prázdné")
	}

	// projít jednotlivé prefixy a zkusit v nich získat volnou IP adresu
	for _, prefixId := range prefixIds.List() {
		resource := ipam.NewIpamPrefixesAvailableIpsCreateParams().WithID(int64(prefixId.(int)))

		// chce: IpamPrefixesAvailableIpsCreateParams
		resourceCreated, err := client.Ipam.IpamPrefixesAvailableIpsCreate(resource, nil)

		if err != nil {
			// status 204 se vrací v případě, že je subnet plný
			if m, _ := regexp.MatchString("status 204", err.Error()); m {
				continue
			} else {
				return err
			}
		}

		// uložíme si ID
		// resourceCreated.Payload.ID
		d.SetId(strconv.FormatInt(resourceCreated.Payload.ID, 10))

		// a přečteme IP adresu pro získání zbylých údajů a máme hotovo
		return resourceNetboxIpamIPByPrefixRead(d, m)
	}

	return errors.New("žádný ze subnetů nemá volnou IP adresu")
}

func resourceNetboxIpamIPByPrefixRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*netboxclient.NetBoxAPI)

	ipID := d.Id()
	ipIDInt64, err := strconv.ParseInt(ipID, 10, 64)

	if err != nil {
		return err
	}

	params := ipam.NewIpamIPAddressesReadParams().WithID(ipIDInt64)
	resource, err := client.Ipam.IpamIPAddressesRead(params, nil)

	// pokud se vrátí chyba
	if err != nil {
		if m, _ := regexp.MatchString("status 404", err.Error()); m {
			// chyba je 404
			// IP adresa neexistuje
			d.SetId("")
			return nil
		} else {
			return err
		}
	}

	payload := resource.Payload

	if err = d.Set("address", payload.Address); err != nil {
		return err
	}

	if err = d.Set("description", payload.Description); err != nil {
		return err
	}

	if err = d.Set("dns_name", payload.DNSName); err != nil {
		return err
	}

	if payload.AssignedObjectID != nil {
		if *payload.AssignedObjectType == "dcim.interface" {
			if err = d.Set("interface_id", payload.AssignedObjectID); err != nil {
				return err
			}
		}
	}

	if payload.NatInside == nil {
		if err = d.Set("nat_inside_id", nil); err != nil {
			return err
		}
	} else {
		if err = d.Set("nat_inside_id", payload.NatInside.ID); err != nil {
			return err
		}
	}

	if payload.NatOutside == nil {
		if err = d.Set("nat_outside_id", nil); err != nil {
			return err
		}
	} else {
		if err = d.Set("nat_outside_id", payload.NatOutside.ID); err != nil {
			return err
		}
	}

	if payload.Role == nil {
		if err = d.Set("role", nil); err != nil {
			return err
		}
	} else {
		if err = d.Set("role", payload.Role.Value); err != nil {
			return err
		}
	}

	if payload.Status == nil {
		if err = d.Set("status", nil); err != nil {
			return err
		}
	} else {
		if err = d.Set("status", payload.Status.Value); err != nil {
			return err
		}
	}

	if err = d.Set("tags", payload.Tags); err != nil {
		return err
	}

	if payload.Tenant == nil {
		if err = d.Set("tenant_id", nil); err != nil {
			return err
		}
	} else {
		if err = d.Set("tenant_id", payload.Tenant.ID); err != nil {
			return err
		}
	}

	if payload.Vrf == nil {
		if err = d.Set("vrf_id", nil); err != nil {
			return err
		}
	} else {
		if err = d.Set("vrf_id", payload.Vrf.ID); err != nil {
			return err
		}
	}

	return nil
}

func resourceNetboxIpamIPByPrefixUpdate(d *schema.ResourceData, m interface{}) error {

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

func resourceNetboxIpamIPByPrefixsDelete(d *schema.ResourceData, m interface{}) error {

	client := m.(*netboxclient.NetBoxAPI)

	ipID := d.Id()
	ipIDInt64, err := strconv.ParseInt(ipID, 10, 64)

	if err != nil {
		return err
	}

	resource := ipam.NewIpamIPAddressesDeleteParams().WithID(ipIDInt64)

	if _, err := client.Ipam.IpamIPAddressesDelete(resource, nil); err != nil {
		return err
	}

	d.SetId("")

	return nil
}
