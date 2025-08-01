resource "cloudflare_queue_consumer" "%[1]s" {
  account_id  = "%[2]s"
  queue_id    = "%[1]s"
  script_name = "test_script"
  settings = {
    batch_size       = 50
    max_concurrency  = 10
    max_retries      = 5
    max_wait_time_ms = 5000
    retry_delay      = 10
  }
  type = "worker"
}