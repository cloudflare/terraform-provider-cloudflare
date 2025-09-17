resource "cloudflare_snippet" "%[1]s" {
  zone_id     = "%[2]s"
  name        = "%[1]s"
  main_module = "worker.js"
  files {
    name    = "worker.js"
    content = <<-EOT
    // Complex worker with multiple features
    export default {
      async fetch(request, env, ctx) {
        const url = new URL(request.url);
        
        // Special characters: $, @, #, &, *, {}
        const specialPath = "/api/v1/$${resource}";
        
        // Multi-line string
        const message = `Request received:
          Method: $${request.method}
          Path: $${url.pathname}
          Time: $${new Date().toISOString()}`;
        
        // Handle different routes
        if (url.pathname === '/') {
          return new Response('Hello World!', {
            headers: { 'content-type': 'text/plain' },
          });
        } else if (url.pathname.startsWith('/api')) {
          return new Response(JSON.stringify({ message, specialPath }), {
            headers: { 'content-type': 'application/json' },
          });
        }
        
        // Pass through to origin
        return fetch(request);
      }
    }
    EOT
  }
}