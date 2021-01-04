

provider "oauth" {

}


data "oauth_token" "example" {
      client_id = "xxxxxxx"
    client_secret = "xxxxx"
    token_endpoint = "https://test.test.io/oauth/connect/token"
    scopes = [""]
}


output "name" {
  value = data.oauth_token.example
}
