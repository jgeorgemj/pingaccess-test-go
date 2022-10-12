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
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	_, err := http.Get("https://localhost:9000/")
	if err != nil {
		fmt.Println(err)
	}

	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return pingaccess.Provider()
		},
	})
}
