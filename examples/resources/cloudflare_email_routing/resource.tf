
# Managed destination address
resource "cloudflare_email_routing_address" "example" {
    account_id = "01a7362d577a6c3019a474fd6f485823"
    email = "user@example.com"
}

# Manage Email Routing on zone
resource "cloudflare_email_routing_settings" "my_zone" {
    zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
    enabled = "true"
}

resource "cloudflare_email_routing_rule" "main" {
    zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
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
