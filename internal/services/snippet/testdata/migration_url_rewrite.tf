resource "cloudflare_snippet" "%[1]s" {
  zone_id     = "%[2]s"
  name        = "%[1]s"
  main_module = "rewrite.js"

  files {
    name    = "rewrite.js"
    content = <<-JS
// Reroutes requests from example.com/old/* to example.com/new/*
export default {
  async fetch(request) {
    const url = new URL(request.url);
    const match = url.pathname.match(/^\/old\/(.+)$/);
    
    if (!match) {
      return fetch(request);
    }
    
    url.pathname = '/new/' + match[1];
    return fetch(url, request);
  }
};
JS
  }
}