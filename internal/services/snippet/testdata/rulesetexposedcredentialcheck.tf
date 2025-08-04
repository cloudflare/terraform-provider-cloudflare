
resource "cloudflare_ruleset" "%[1]s" {
  account_id  = "%[3]s"
  name        = "%[2]s"
  description = "This ruleset includes a rule checking for exposed credentials."
  kind        = "custom"
  phase       = "http_request_firewall_custom"

  rules = [{
    action      = "log"
    expression  = "http.request.method == \"POST\" && http.request.uri == \"/login.php\""
    enabled     = true
    description = "example exposed credential check"
    exposed_credential_check = {
      username_expression = "url_decode(http.request.body.form[\"username\"][0])"
      password_expression = "url_decode(http.request.body.form[\"password\"][0])"
    }
  }]
}
