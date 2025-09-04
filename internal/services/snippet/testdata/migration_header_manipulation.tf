resource "cloudflare_snippet" "%[1]s" {
  zone_id     = "%[2]s"
  name        = "%[1]s"
  main_module = "headers.js"

  files {
    name    = "headers.js"
    content = <<-JS
// Adds custom headers and modifies requests
export default {
  async fetch(request) {
    const newRequest = new Request(request);
    
    // Add custom headers
    newRequest.headers.set("X-Custom-Header", "OpenAI");
    newRequest.headers.set("X-Request-ID", crypto.randomUUID());
    
    // Remove sensitive headers
    newRequest.headers.delete("Referer");
    newRequest.headers.delete("X-Forwarded-For");
    
    return fetch(newRequest);
  }
};
JS
  }
}