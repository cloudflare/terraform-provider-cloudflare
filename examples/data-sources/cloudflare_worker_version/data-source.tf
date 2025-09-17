data "cloudflare_worker_version" "example_worker_version" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  worker_id = "worker_id"
  version_id = "version_id"
  include = "modules"
}
