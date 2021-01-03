package oauth

import (
	"net/http"

	"golang.org/x/oauth2/clientcredentials"
)

type api_client struct {
	http_client  *http.Client
	oauth_config *clientcredentials.Config
}

type apiClientOpt struct {
	oauth_client_id     string
	oauth_client_secret string
	oauth_scopes        []string
	oauth_token_url     string
}

func NewAPIClient(opt *apiClientOpt) (*api_client, error) {

	client := api_client{}

	client.oauth_config = &clientcredentials.Config{
		ClientID:     opt.oauth_client_id,
		ClientSecret: opt.oauth_client_secret,
		TokenURL:     opt.oauth_token_url,
		Scopes:       opt.oauth_scopes,
	}

	return &client, nil
}
