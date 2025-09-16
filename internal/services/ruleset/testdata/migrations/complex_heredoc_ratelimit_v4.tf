resource "cloudflare_ruleset" "%[2]s" {
  account_id  = "%[1]s"
  kind        = "zone"
  name        = "Complex Heredoc Test %[2]s"
  phase       = "http_ratelimit"
  description = "Test complex heredoc with ratelimit"
  
  rules {
    action      = "log"
    description = "Log rate limit with complex expression"
    enabled     = true
    expression  = <<-EOF
      not (
        any(http.request.headers["authorization"][*] contains "sess-")
        or any(http.request.headers["authorization"][*] contains "sk-")
        or http.cookie contains "__Secure-next-auth."
      )
      and not http.user_agent eq "Stripe/1.0 (+https://stripe.com/docs/webhooks)"
    EOF
    ratelimit {
      characteristics     = ["cf.unique_visitor_id", "cf.colo.id"]
      mitigation_timeout  = 3600
      period              = 10
      requests_per_period = 1000
    }
  }
}