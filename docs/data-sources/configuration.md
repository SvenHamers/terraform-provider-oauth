---
page_title: "oauth_token Data Source - terraform-provider-oauth"
subcategory: ""
description: |-
  
---

# Data Source `oauth_token`


```
data "oauth_configuration" "example" {
    well_known_endpoint = "wel known endpoint" (required)
}

```

## Arguments Reference

The following arguments are supported:

* `access_token_issuer`
* `authorization_endpoint` 
* `device_authorization_endpoint`
* `end_session_endpoint`
* `issuer` 
* `jwks_uri`
* `token_endpoint` 
* `userinfo_endpoint` 