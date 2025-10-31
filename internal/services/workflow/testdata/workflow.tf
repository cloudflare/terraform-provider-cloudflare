resource "cloudflare_worker" "workflow_worker" {
    account_id  = "%[2]s"
    name        = "%[4]s"
}

resource "cloudflare_worker_version" "workflow_version" {
    account_id  = "%[2]s"
    worker_id = cloudflare_worker.workflow_worker.id
    compatibility_date = "2024-11-13"
    main_module = "worker.js"
    modules = [{
        content_file = "%[6]s"
        content_type = "application/javascript+module"
        name = "worker.js"
    }]
    depends_on = [
        cloudflare_worker.workflow_worker
    ]
}

resource "cloudflare_workers_deployment" "workflow_deployment" {
  account_id = "%[2]s"
  script_name = cloudflare_worker.workflow_worker.name
  strategy = "percentage"
  versions = [{
    percentage = 100
    version_id = cloudflare_worker_version.workflow_version.id
  }]
  depends_on = [
    cloudflare_worker_version.workflow_version
  ]
}

resource "cloudflare_workflow" "%[1]s" {
   account_id    = "%[2]s"
   workflow_name = "%[3]s"
   script_name   = cloudflare_worker.workflow_worker.name
   class_name    = "%[5]s"
   depends_on = [
      cloudflare_workers_deployment.workflow_deployment,
   ]
 }
