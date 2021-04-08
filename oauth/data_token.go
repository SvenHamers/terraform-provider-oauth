package oauth

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceToken() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTokenRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"token": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"token_endpoint": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"client_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"client_secret": &schema.Schema{
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"scopes": &schema.Schema{
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Description: "scopes",
			},
		},
	}
}

func dataSourceTokenRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	tokenEndpoint := d.Get("token_endpoint").(string)
	clientID := d.Get("client_id").(string)
	clientSecret := d.Get("client_secret").(string)
	scopes := d.Get("scopes").([]interface{})
	
	val, present := os.LookupEnv("oauth_client_secret")
	
	if present == true {
        clientSecret = val
    	}

	opt := &apiClientOpt{
		oauth_client_id:     clientID,
		oauth_client_secret: clientSecret,
		oauth_token_url:     tokenEndpoint,
		oauth_scopes:        expandStringSet(scopes),
	}

	client, _ := NewAPIClient(opt)

	tokenSource := client.oauth_config.TokenSource(context.Background())
	token, err := tokenSource.Token()
	if err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(token.AccessToken)
	d.Set("token", token.AccessToken)

	return diags
}

func expandStringSet(configured []interface{}) []string {
	return expandStringList(configured)
}

func expandStringList(configured []interface{}) []string {
	vs := make([]string, 0, len(configured))
	for _, v := range configured {
		val, ok := v.(string)
		if ok && val != "" {
			vs = append(vs, v.(string))
		}
	}
	return vs
}
