resource "cloudflare_ai_search_namespace" "example_ai_search_namespace" {
  account_id = "c3dc5f0b34a14ff8e1b3ec04895e1b22"
  name = "name"
  description = "Production environment"
  public_endpoint_params = {
    authorized_hosts = ["string"]
    chat_completions_endpoint = {
      disabled = true
    }
    custom_domains = ["search.example.com"]
    default_domain_enabled = true
    enabled = true
    instances_allowed = ["docs", "blog"]
    mcp = {
      description = "description"
      disabled = true
    }
    rate_limit = {
      period_ms = 60000
      requests = 1
      technique = "fixed"
    }
    search_endpoint = {
      disabled = true
    }
  }
}
