resource "cloudflare_workers_script" "%[1]s" {
  account_id  = "%[2]s"
  script_name = "%[1]s"
  content = <<-EOT
    import {DurableObject} from "cloudflare:workers"
    import wasm from "./module.wasm"
    export class MyDurableObject extends DurableObject {}
    export default { fetch() {return new Response(wasm instanceof WebAssembly.Module ? "ok" : "error")} }
  EOT
  main_module = "worker.js"
  files = {
    "module.wasm" = {
      content_base64 = "%[3]s"
      content_type   = "application/wasm"
    }
  }
  migrations = {
    new_tag            = "v1"
    new_sqlite_classes = ["MyDurableObject"]
  }
  bindings = [
    {
      name       = "MY_DO"
      type       = "durable_object_namespace"
      class_name = "MyDurableObject"
    }
  ]
}
