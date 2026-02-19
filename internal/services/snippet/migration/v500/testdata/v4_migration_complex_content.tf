resource "cloudflare_snippet" "%[2]s" {
  zone_id     = "%[1]s"
  name        = "%[2]s"
  main_module = "worker.js"

  files {
    name    = "worker.js"
    content = <<-EOT
      // Complex worker with multiple features
      export default {
        async fetch(request, env, ctx) {
          const url = new URL(request.url);

          // Handle different paths
          if (url.pathname === '/api/data') {
            return new Response(JSON.stringify({
              message: "Hello from Cloudflare",
              path: url.pathname,
              timestamp: Date.now()
            }), {
              headers: {
                'Content-Type': 'application/json',
                'X-Custom-Header': 'test-value'
              }
            });
          }

          // Default response
          return fetch(request);
        }
      };
    EOT
  }
}