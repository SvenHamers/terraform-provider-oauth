package oauth

import (
	"context"

	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type OauthConfig struct {
	AccessTokenIssuer           string `access_token_issuer`
	AuthorizationEndpoint       string `authorization_endpoint`
	DeviceQuthorizationEndpoint string `device_authorization_endpoint`
	EndSessionEndpoint          string `end_session_endpoint`
	Issuer                      string `issuer`
	JwksUri                     string `jwks_uri`
	TokenUndpoint               string `token_endpoint`
	UserinfoEndpoint            string `userinfo_endpoint`
}

func dataSourceConfiguration() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceoConfigurationRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"well_known_endpoint": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"access_token_issuer": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"authorization_endpoint": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"device_authorization_endpoint": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"end_session_endpoint": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"issuer": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"jwks_uri": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"token_endpoint": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"userinfo_endpoint": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceoConfigurationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	wellKnown := d.Get("well_known_endpoint").(string)

	resp, err := http.Get(wellKnown)
	if err != nil {
		return diag.FromErr(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte

	var result OauthConfig

	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		return diag.FromErr(err)
	}

	// always run
	d.SetId(wellKnown)
	d.Set("access_token_issuer", result.AccessTokenIssuer)
	d.Set("authorization_endpoint", result.AuthorizationEndpoint)
	d.Set("device_authorization_endpoint", result.DeviceQuthorizationEndpoint)
	d.Set("end_session_endpoint", result.EndSessionEndpoint)
	d.Set("issuer", result.Issuer)
	d.Set("jwks_uri", result.JwksUri)
	d.Set("token_endpoint", result.TokenUndpoint)
	d.Set("userinfo_endpoint", result.UserinfoEndpoint)

	return diags
}
