package netbox

import (
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	netboxclient "github.com/netbox-community/go-netbox/netbox/client"
	"github.com/netbox-community/go-netbox/netbox/client/ipam"
	"github.com/netbox-community/go-netbox/netbox/models"
	pkgerrors "github.com/pkg/errors"
)

func resourceNetboxIpamPrefix() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxIpamPrefixCreate,
		Read:   resourceNetboxIpamPrefixRead,
		Update: resourceNetboxIpamPrefixUpdate,
		Delete: resourceNetboxIpamPrefixDelete,
		Exists: resourceNetboxIpamPrefixExists,

		Schema: map[string]*schema.Schema{
			"description": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(1, 100),
			},
			"is_pool": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  nil,
			},
			"prefix": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsCIDRNetwork(0, 256),
			},
			"role_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"site_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "active",
				ValidateFunc: validation.StringInSlice([]string{"container", "active",
					"reserved", "deprecated"}, false),
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
			"vlan_id": {
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

func resourceNetboxIpamPrefixCreate(d *schema.ResourceData,
	m interface{}) error {
	client := m.(*netboxclient.NetBox)

	log.Printf("[DEBUG] Start resourceNetboxIpamPrefixCreate func")

	description := d.Get("description").(string)
	isPool := d.Get("is_pool").(bool)
	prefix := d.Get("prefix").(string)
	roleID := int64(d.Get("role_id").(int))
	siteID := int64(d.Get("site_id").(int))
	status := d.Get("status").(string)
	tags := d.Get("tags").(*schema.Set).List()
	tenantID := int64(d.Get("tenant_id").(int))
	vlanID := int64(d.Get("vlan_id").(int))
	vrfID := int64(d.Get("vrf_id").(int))

	log.Printf("[DEBUG] Get description (%s) from terraform code", description)
	log.Printf("[DEBUG] Get isPool (%v) from terraform code", isPool)
	log.Printf("[DEBUG] Get prefix (%s) from terraform code", prefix)
	log.Printf("[DEBUG] Get roleID (%v) from terraform code", roleID)
	log.Printf("[DEBUG] Get siteID (%v) from terraform code", siteID)
	log.Printf("[DEBUG] Get status (%s) from terraform code", status)
	log.Printf("[DEBUG] Get tags (%s) from terraform code", tags)
	log.Printf("[DEBUG] Get tenantID (%v) from terraform code", tenantID)
	log.Printf("[DEBUG] Get vlanID (%v) from terraform code", vlanID)
	log.Printf("[DEBUG] Get vrfID (%v) from terraform code", vrfID)

	newResource := &models.WritablePrefix{
		Description: description,
		IsPool:      isPool,
		Prefix:      &prefix,
		Status:      status,
		Tags:        expandToStringSlice(tags),
	}

	if roleID != 0 {
		newResource.Role = &roleID
	}

	if siteID != 0 {
		newResource.Site = &siteID
	}

	if tenantID != 0 {
		newResource.Tenant = &tenantID
	}

	if vlanID != 0 {
		newResource.Vlan = &vlanID
	}

	if vrfID != 0 {
		newResource.Vrf = &vrfID
	}

	resource := ipam.NewIpamPrefixesCreateParams().WithData(newResource)

	resourceCreated, err := client.Ipam.IpamPrefixesCreate(resource, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(resourceCreated.Payload.ID, 10))

	log.Printf("[DEBUG] End resourceNetboxIpamPrefixCreate func")

	return resourceNetboxIpamPrefixRead(d, m)
}

func resourceNetboxIpamPrefixRead(d *schema.ResourceData,
	m interface{}) error {
	client := m.(*netboxclient.NetBox)

	log.Printf("[DEBUG] Start resourceNetboxIpamPrefixCreate func")

	resourceID := d.Id()
	log.Printf("[DEBUG] Get resource id (%v) from terraform state", resourceID)

	params := ipam.NewIpamPrefixesListParams().WithID(&resourceID)
	resources, err := client.Ipam.IpamPrefixesList(params, nil)
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Fetch each resource to see if id is found")
	for _, resource := range resources.Payload.Results {
		log.Printf("[DEBUG] Check resource with ID: %v",
			strconv.FormatInt(resource.ID, 10))
		if strconv.FormatInt(resource.ID, 10) == d.Id() {
			if err = d.Set("description", resource.Description); err != nil {
				return err
			}

			if err = d.Set("is_pool", resource.IsPool); err != nil {
				return err
			}

			if err = d.Set("prefix", resource.Prefix); err != nil {
				return err
			}

			if resource.Role == nil {
				if err = d.Set("role_id", nil); err != nil {
					return err
				}
			} else {
				if err = d.Set("role_id", resource.Role.ID); err != nil {
					return err
				}
			}

			if resource.Site == nil {
				if err = d.Set("site_id", nil); err != nil {
					return err
				}
			} else {
				if err = d.Set("site_id", resource.Site.ID); err != nil {
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

			if resource.Vlan == nil {
				if err = d.Set("vlan_id", nil); err != nil {
					return err
				}
			} else {
				if err = d.Set("vlan_id", resource.Vlan.ID); err != nil {
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

	//d.SetId("")

	log.Printf("[DEBUG] End resourceNetboxIpamPrefixCreate func")

	return nil
}

func resourceNetboxIpamPrefixUpdate(d *schema.ResourceData,
	m interface{}) error {
	client := m.(*netboxclient.NetBox)
	params := &models.WritablePrefix{}

	if d.HasChange("description") {
		description := d.Get("description").(string)
		params.Description = description
	}

	params.IsPool = d.Get("is_pool").(bool)

	prefix := d.Get("prefix").(string)
	params.Prefix = &prefix

	if d.HasChange("role_id") {
		roleID := int64(d.Get("role_id").(int))
		if roleID != 0 {
			params.Role = &roleID
		}
	}

	if d.HasChange("site_id") {
		siteID := int64(d.Get("site_id").(int))
		if siteID != 0 {
			params.Site = &siteID
		}
	}

	if d.HasChange("status") {
		params.Status = d.Get("status").(string)
	}

	tags := d.Get("tags").(*schema.Set).List()
	params.Tags = expandToStringSlice(tags)

	if d.HasChange("tenant_id") {
		tenantID := int64(d.Get("tenant_id").(int))
		if tenantID != 0 {
			params.Tenant = &tenantID
		}
	}

	if d.HasChange("vlan_id") {
		vlanID := int64(d.Get("vlan_id").(int))
		if vlanID != 0 {
			params.Vlan = &vlanID
		}
	}

	if d.HasChange("vrf_id") {
		vrfID := int64(d.Get("vrf_id").(int))
		if vrfID != 0 {
			params.Vrf = &vrfID
		}
	}

	resource := ipam.NewIpamPrefixesPartialUpdateParams().WithData(params)

	resourceID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return pkgerrors.New("Unable to convert ID into int64")
	}

	resource.SetID(resourceID)

	_, err = client.Ipam.IpamPrefixesPartialUpdate(resource, nil)
	if err != nil {
		return err
	}

	return resourceNetboxIpamPrefixRead(d, m)
}

func resourceNetboxIpamPrefixDelete(d *schema.ResourceData,
	m interface{}) error {
	client := m.(*netboxclient.NetBox)

	resourceExists, err := resourceNetboxIpamPrefixExists(d, m)
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

	resource := ipam.NewIpamPrefixesDeleteParams().WithID(id)
	if _, err := client.Ipam.IpamPrefixesDelete(resource, nil); err != nil {
		return err
	}

	return nil
}

func resourceNetboxIpamPrefixExists(d *schema.ResourceData,
	m interface{}) (b bool, e error) {
	client := m.(*netboxclient.NetBox)
	resourceExist := false

	resourceID := d.Id()
	params := ipam.NewIpamPrefixesListParams().WithID(&resourceID)
	resources, err := client.Ipam.IpamPrefixesList(params, nil)
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
