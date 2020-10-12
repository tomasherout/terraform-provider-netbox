package netbox

import (
	"regexp"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	pkgerrors "github.com/pkg/errors"
	netboxclient "github.com/tomasherout/go-netbox/netbox/client"
	"github.com/tomasherout/go-netbox/netbox/client/ipam"
)

func dataNetboxIpamRole() *schema.Resource {
	return &schema.Resource{
		Read: dataNetboxIpamRoleRead,

		Schema: map[string]*schema.Schema{
			"slug": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringMatch(
					regexp.MustCompile("^[-a-zA-Z0-9_]{1,50}$"),
					"Must be like ^[-a-zA-Z0-9_]{1,50}$"),
			},
		},
	}
}

func dataNetboxIpamRoleRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*netboxclient.NetBoxAPI)

	slug := d.Get("slug").(string)

	p := ipam.NewIpamRolesListParams().WithSlug(&slug)

	list, err := client.Ipam.IpamRolesList(p, nil)
	if err != nil {
		return err
	}

	if *list.Payload.Count == 1 {
		d.SetId(strconv.FormatInt(list.Payload.Results[0].ID, 10))
	} else {
		return pkgerrors.New("Data results for netbox_ipam_role returns 0 or " +
			"more than one result.")
	}

	return nil
}
