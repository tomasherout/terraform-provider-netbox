package netbox

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	pkgerrors "github.com/pkg/errors"
	netboxclient "github.com/tomasherout/go-netbox/netbox/client"
	"github.com/tomasherout/go-netbox/netbox/client/ipam"
)

func dataNetboxIpamVlan() *schema.Resource {
	return &schema.Resource{
		Read: dataNetboxIpamVlanRead,

		Schema: map[string]*schema.Schema{
			"vlan_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"vlan_group_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func dataNetboxIpamVlanRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*netboxclient.NetBoxAPI)

	id := int64(d.Get("vlan_id").(int))
	idStr := strconv.FormatInt(id, 10)
	groupID := int64(d.Get("vlan_group_id").(int))
	groupIDStr := strconv.FormatInt(groupID, 10)

	p := ipam.NewIpamVlansListParams().WithVid(&idStr)
	if groupID != 0 {
		p.SetGroupID(&groupIDStr)
	}

	list, err := client.Ipam.IpamVlansList(p, nil)
	if err != nil {
		return err
	}

	if *list.Payload.Count == 1 {
		d.SetId(strconv.FormatInt(list.Payload.Results[0].ID, 10))
	} else {
		return pkgerrors.New("Data results for netbox_ipam_vlan returns 0 or " +
			"more than one result.")
	}

	return nil
}
