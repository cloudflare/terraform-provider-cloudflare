resource "cloudflare_snippet" "%[2]s" {
  zone_id     = "%[1]s"
  name        = "%[2]s"
  main_module = "main.js"

  files {
    name    = "main.js"
    content = "import { helper } from './helper.js'; export default {async fetch(request) {return helper(request)}};"
  }

  files {
    name    = "helper.js"
    content = "export function helper(request) { return fetch(request); }"
  }

  files {
    name    = "utils.js"
    content = "export const VERSION = '1.0.0';"
  }
}