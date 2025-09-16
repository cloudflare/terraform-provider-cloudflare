resource "cloudflare_ruleset" "%[2]s" {
  account_id  = "%[1]s"
  kind        = "zone"
  name        = "Global rate limit %[2]s"
  phase       = "http_ratelimit"
  description = "Test multiple rules with ratelimit"

  rules {
    action      = "log"
    description = "Log Global rate limit 1000/10s non authenticated"
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

  rules {
    action      = "log"
    description = "Log Global rate limit 2500/1m non authenticated"
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
      period              = 60
      requests_per_period = 2500
    }
  }

  rules {
    action      = "log"
    description = "Log Unauthed rate limit DDoS User Agents"
    enabled     = true
    expression  = <<-EOF
      (
        (http.user_agent eq "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36")
        or (http.user_agent eq "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36")
      )
      and not (
        any(http.request.headers["authorization"][*] contains "sess-")
        or any(http.request.headers["authorization"][*] contains "sk-")
        or http.cookie contains "__Secure-next-auth."
      )
    EOF
    ratelimit {
      characteristics     = ["ip.src", "cf.colo.id"]
      mitigation_timeout  = 300
      period              = 10
      requests_per_period = 300
    }
  }
}