resource "cloudflare_user_agent_blocking_rule" "example-1" {
	zone_id      = "24670de1ee826d9591121cbd94418f7b"
	mode         = "js_challenge"
	paused       = false
	description  = "My description 1"
	configuration {
		target = "ua"
		value  = "Chrome"
	}
}
resource "cloudflare_user_agent_blocking_rule" "example-2" {
	zone_id      = "24670de1ee826d9591121cbd94418f7b"
	mode         = "managed_challenge"
	paused       = true
	description  = "My description 22"
	configuration {
		target = "ua"
		value  = "Mozilla"
	}
}