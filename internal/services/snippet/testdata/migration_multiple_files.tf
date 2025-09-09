resource "cloudflare_snippet" "%[1]s" {
  zone_id     = "%[2]s"
  name        = "%[1]s"
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