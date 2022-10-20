package pingaccess

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// PingAccess Provider function
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{
			"pingaccess_enginelistener": enginelistener(),
		},
	}
}
