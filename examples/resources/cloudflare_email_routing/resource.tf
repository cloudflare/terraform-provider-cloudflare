
# Managed destination address
resource "cloudflare_email_routing_address" "example" {
    account_id = "f037e56e89293a057740de681ac9abbe"
    email = "user@example.com"
}

# Managed Email Routing Settings on zone
resource "cloudflare_email_routing_settings" "my_zone" {
    zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
    enabled = "true"
}

# Managed Email Routing Rule
resource "cloudflare_email_routing_rule" "main" {
    zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
    name = "terraform rule"
    enabled =  true
    matchers = [
        {
            type = "literal",
            field = "to",
            value = "test@example.com"
        }
    ] 

    actions = [
        {
            type = "forward"
            value = ["destinationaddress@example.net"]
        }   
    ]
}

# Managed Email Routing Rule
resource "cloudflare_email_routing_rule_catch_all" "main" {
    zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
    name = "terraform rule"
    enabled =  true
    matchers = [
        {
            type = "all",
        }
    ] 

    actions = [
        {
            type = "forward"
            value = ["destinationaddress@example.net"]
        }   
    ]
}