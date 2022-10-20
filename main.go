package main

import (
	"fmt"
	"net/http"
	"crypto/tls"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
  
	"terraform-provider-pa/pingaccess"
)

func main() {
//  Skip the localhost cert prompt when using local docker instance, might need to remove this in prod

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	_, err := http.Get("https://localhost:9000/")
	if err != nil {
		fmt.Println(err)
	}
// the main file should be executed as plugin.
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return pingaccess.Provider()
		},
	})
}
