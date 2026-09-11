resource "cloudflare_workers_script" "%[1]s" {
  account_id  = "%[2]s"
  script_name = "%[1]s"
  content     = "export default { fetch() { return new Response('Hello world'); } };"
  main_module = "worker.mjs"
  bindings = [
    {
      name = "LOADER"
      type = "worker_loader"
    }
  ]
}
