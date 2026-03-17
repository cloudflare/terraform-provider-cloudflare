resource "cloudflare_queue" "test_queue" {
  account_id = "%[1]s"
  queue_name = "%[2]s"
}

resource "cloudflare_workers_script" "worker_script" {
  account_id  = "%[1]s"
  script_name = "%[2]s"
  content     = <<-EOT
export default {
  async queue(batch, env, ctx) {
    for (const message of batch.messages) {
      console.log('Received', message);
    }
  }
};
EOT
  main_module = "index.js"
  depends_on  = [cloudflare_queue.test_queue]
}

resource "cloudflare_queue_consumer" "%[2]s" {
  account_id  = "%[1]s"
  queue_id    = cloudflare_queue.test_queue.id
  type        = "worker"
  script_name = cloudflare_workers_script.worker_script.script_name
  depends_on  = [cloudflare_workers_script.worker_script]

  lifecycle {
    # API returns "script" in Read but we send "script_name" on Create
    # settings is also returned differently for worker vs http_pull consumers
    ignore_changes = [script_name, settings]
  }
}
