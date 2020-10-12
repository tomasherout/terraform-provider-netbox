package netbox

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	netboxclient "github.com/tomasherout/go-netbox/netbox/client"
	"github.com/tomasherout/go-netbox/netbox/client/ipam"
	"strconv"
	"time"
)

func dataNetboxIpamIPPrefixes() *schema.Resource {
	return &schema.Resource{
		Read: dataNetboxIpamIPPrefixesRead,

		Schema: map[string]*schema.Schema{
			"tags": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ids": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
		},
	}
}

func dataNetboxIpamIPPrefixesRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*netboxclient.NetBoxAPI)

	tags := d.Get("tags").(*schema.Set)
	tagsStr := make([]string, tags.Len())

	for i, elem := range tags.List() {
		tagsStr[i] = elem.(string)
	}

	params := ipam.NewIpamPrefixesListParams().WithTag(tagsStr)

	res, err := client.Ipam.IpamPrefixesList(params, nil)

	if err != nil {
		return err
	}

	ids := make([]int64, len(res.Payload.Results))

	for i, elem := range res.Payload.Results {
		ids[i] = elem.ID
	}

	d.Set("ids", ids)

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return nil
}
