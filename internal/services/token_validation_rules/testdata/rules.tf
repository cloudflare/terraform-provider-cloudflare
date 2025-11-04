resource "cloudflare_token_validation_config" "%[1]s" {
  zone_id = "%[2]s"
  token_type = "JWT"
  title = "Test config"
  description = "Terraform acceptance test config"
  token_sources = [
    "http.request.headers[\"authorization\"][0]"
  ]
  credentials = {
    keys = [
        {
            alg = "ES256"
            kid = "some-kid"
            kty = "EC"
            crv = "P-256"
            x = "yl_BZSxUG5II7kJCMxDfWImiU6zkcJcBYaTgzV3Jgnk"
            y = "0qAzLQe_YGEdotb54qWq00k74QdiTOiWnuw_YzuIqr0"
        }
    ]
  }
}

resource "cloudflare_api_shield_operation" "%[1]s" {
	zone_id  = "%[2]s"
	method   = "GET"
	host     = "example.com"
	endpoint = "/excluded"
}


resource "cloudflare_token_validation_rules" "%[1]s" {
 zone_id = "%[2]s"
 title = "%[3]s"
 description = "%[4]s"
 action = "%[5]s"
 enabled = %[6]s
 # reference the ID of the generated token config, this constructs: is_jwt_valid("<uuid>")
 expression = format("(is_jwt_valid(%%q))", cloudflare_token_validation_config.%[1]s.id)
 selector = {
    include = [{
        host = ["example.com"]
    }]
    exclude = [{
        # reference the ID of the generated operation to exclude it
        operation_ids = ["${cloudflare_api_shield_operation.%[1]s.id}"]
    }]
 }
}