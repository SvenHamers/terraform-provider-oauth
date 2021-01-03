package main

import (
	"github.com/svenhamers/terraform-provider-oauth/oauth"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: oauth.Provider,
	})
}