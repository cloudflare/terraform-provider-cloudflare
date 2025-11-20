resource "cloudflare_worker" "example_worker" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  name = "my-worker"
  logpush = true
  observability = {
    enabled = true
    head_sampling_rate = 1
    logs = {
      enabled = true
      head_sampling_rate = 1
      invocation_logs = true
    }
  }
  subdomain = {
    enabled = true
    previews_enabled = true
  }
  tags = ["my-team", "my-public-api"]
  tail_consumers = [{
    name = "my-tail-consumer"
  }]
}
