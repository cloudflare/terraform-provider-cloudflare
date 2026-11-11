resource "cloudflare_ai_security_custom_topics" "example_ai_security_custom_topics" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
  topics = [{
    label = "credit-cards"
    topic = "credit card numbers"
  }]
}
