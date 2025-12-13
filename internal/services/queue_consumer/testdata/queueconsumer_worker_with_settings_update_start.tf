resource "cloudflare_queue" "test_queue" {
  account_id = "%[1]s"
  queue_name = "%[2]s"
}

resource "cloudflare_workers_script" "worker_script_with_settings" {
  account_id  = "%[5]s"
  script_name = "test-worker-consumer-worker-with-settings-update"
  bindings = [
    {
      type       = "queue"
      name       = "incoming"
      queue_name = cloudflare_queue.test_queue.queue_name
    }
  ]
  content             = <<-EOT
export default {
  async queue(batch, env, ctx) {
    for (const message of batch.messages) {
      console.log('Received', message);
    }
  }
};
EOT
  main_module         = "index.js"


  depends_on = [cloudflare_queue.test_queue]
}

resource "cloudflare_queue_consumer" "%[3]s" {
  account_id  = "%[4]s"
  queue_id    = cloudflare_queue.test_queue.id
  type        = "worker"
  script_name = cloudflare_workers_script.worker_script_with_settings.script_name

  settings = {
    batch_size        = 10
    max_retries       = 3
    max_wait_time_ms  = 5000
    retry_delay = 0
  }
  depends_on = [cloudflare_workers_script.worker_script_with_settings, cloudflare_queue.test_queue]
}