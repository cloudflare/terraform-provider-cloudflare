resource "cloudflare_queue" "test_queue" {
  account_id = "%[1]s"
  queue_name = "%[2]s"
}

resource "cloudflare_queue" "dlq1" {
  account_id = "%[3]s"
  queue_name = "%[4]s"
}

resource "cloudflare_workers_script" "worker_script_with_dead_letter" {
  account_id  = "%[7]s"
  script_name = "test-worker-consumer-worker-with-dead-letter"
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

resource "cloudflare_queue_consumer" "%[5]s" {
  account_id        = "%[6]s"
  queue_id          = cloudflare_queue.test_queue.id
  type              = "worker"
  script_name       = cloudflare_workers_script.worker_script_with_dead_letter.script_name
  dead_letter_queue = cloudflare_queue.dlq1.queue_name

  lifecycle {
    ignore_changes = [
      settings
    ]
  }
  depends_on = [cloudflare_workers_script.worker_script_with_dead_letter, cloudflare_queue.test_queue]
}