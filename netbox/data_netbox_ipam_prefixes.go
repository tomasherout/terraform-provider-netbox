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
				Optional: true,
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

	params := ipam.NewIpamPrefixesListParams()

	// pokud jsou zadané tagy
	if tags := d.Get("tags").(*schema.Set); len(tags.List()) > 0 {
		// převést na pole stringů
		tagsStr := make([]string, tags.Len())
		for i, elem := range tags.List() {
			tagsStr[i] = elem.(string)
		}
		// přidat do parametrů (filtr)
		params.SetTag(tagsStr)
	}

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
