# From modules/blog/main.tf
data "cloudflare_ruleset_rules" "origin" {
    phase = "http_request_origin"
    rules = [
        {
            action = "set_config"
            action_parameters = {
                automatic_https_rewrites = true
            }
            expression = "(http.host == \"blog.example.com\" and http.request.uri.path contains \"/admin\")"
        }
    ]
}

output "origin_rules" {
  value = data.cloudflare_ruleset_rules.origin.rules
}

# From modules/hr/main.tf
data "cloudflare_ruleset_rules" "origin" {
    phase = "http_request_origin"
    rules = [
        {
            action = "route"
            action_parameters = {
                "host_header": "hr-timetracking.example.com",
                "origin": {
                    "host": "hr-timetracking.example.com"
                }
            }
            expression = "(http.host == \"hr.example.com\" and starts_with(http.request.uri.path, \"/timetracking\"))"
        }
    ]
}

output "origin_rules" {
  value = data.cloudflare_ruleset_rules.origin.rules
}

# From main.tf

module "blog" {
    source = "./modules/blog"
}

module "hr" {
    source = "./modules/hr"
}

resource "cloudflare_ruleset" "test" {
    // zone id for example.com zone
    zone_id = "12345678901234567890123456789012"
    name = "default"
    kind = "zone"
    phase = "http_request_origin"
    rules = flatten([
        module.blog.origin_rules,
        module.hr.origin_rules
    ])
}
