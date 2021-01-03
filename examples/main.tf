terraform {
  required_providers {
    oauth = {
      source = "terraform.ebmracecloud.nl/embracecloud/oauth"
      version = "1.14.1"
    }
  }
}

provider "oauth" {

}


data "oauth_token" "example" {
      client_id = "root"
    client_secret = "Ez14ROb6CiFWC27e"
    token_endpoint = "https://squidex-gewoon.staging.embracecloud.io/identity-server/connect/token"
    scopes = [""]
}


output "name" {
  value = data.oauth_token.example
}
