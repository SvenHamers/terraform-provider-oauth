---
page_title: "oauth_token Data Source - terraform-provider-oauth"
subcategory: ""
description: |-
  
---

# Data Source `oauth_token`


```
data "oauth_token" "example" {
    client_id = "xxxxxxx" (required)
    client_secret = "xxxxx" (required)
    token_endpoint = "xxxxxxx" (required)
    scopes = [""]
}
```