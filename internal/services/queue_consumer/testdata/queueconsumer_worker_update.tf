resource "cloudflare_queue" "test_queue" {
  account_id = "%s"
  queue_name = "%s"
}

resource "cloudflare_queue_consumer" "%s" {
  account_id  = "%s"
  queue_id    = cloudflare_queue.test_queue.id
  type        = "worker"
  script_name = "test-worker-2"
}

resource "cloudflare_workers_script" "worker_script" {
  account_id  = "%s"
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