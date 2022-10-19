package pingaccess

import (
	"context"
	"io/ioutil"
	// "encoding/json"
	// "encoding/base64"
	"fmt"
	"net/http"
	// "strconv"
	"time"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)
type Item struct {
    items []string
}
// PA schema to read id,name,port, secure and trustedCertificateGroupID as per the PA api
func enginelistener() *schema.Resource {
	return &schema.Resource{
		ReadContext: enginelistenerRead,
		Schema: map[string]*schema.Schema{
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"port": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"secure": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

// Function to read the PA api engine listeners
// func basicAuth(username, password string) string {
// 	auth := username + ":" + password
// 	return base64.StdEncoding.EncodeToString([]byte(auth))
//   }
  
// func redirectPolicyFunc(req *http.Request, via []*http.Request) error{
// 	req.Header.Add("Authorization","Basic " + basicAuth("administrator","2FederateM0re"))
// 	return nil

func enginelistenerRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	// username := "administrator"
	// password := "2FederateM0re"

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/pa-admin-api/v3/engineListeners", "https://localhost:9000"), nil)
	if err != nil {
		return diag.FromErr(err)
	}
	req.SetBasicAuth("Administrator", "2FederateM0re")
	req.Header.Set("X-Xsrf-Header", "PingAccess")
	resp, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte
	fmt.Println(string(body)) 
	return diags
}
