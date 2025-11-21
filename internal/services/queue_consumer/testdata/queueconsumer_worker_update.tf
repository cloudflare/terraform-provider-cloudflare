resource "cloudflare_queue" "test_queue" {
  account_id = "%[1]s"
  queue_name = "%[2]s"
}

resource "cloudflare_workers_script" "worker_script" {
  account_id  = "%[5]s"
  script_name = "test-worker-2"
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
      console.log('Updated, received', message);
    }
  }
};
EOT
  main_module         = "index.js"
}

resource "cloudflare_queue_consumer" "%[3]s" {
  account_id  = "%[4]s"
  queue_id    = cloudflare_queue.test_queue.id
  type        = "worker"
  script_name = cloudflare_workers_script.worker_script.script_name

  depends_on = [cloudflare_workers_script.worker_script]
}