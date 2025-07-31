resource "cloudflare_snippets" "%[1]s" {
  zone_id      = "%[2]s"
  snippet_name = "%[1]s"
  files = [
    {
      name    = "main.js"
      content = <<-EOT
      export default {
        async fetch(request) {
          // Get the current timestamp
          const timestamp = Date.now();

          // Convert the timestamp to hexadecimal format
          const hexTimestamp = timestamp.toString(16);

          // Clone the request and add the custom header
          const modifiedRequest = new Request(request, {
              headers: new Headers(request.headers)
          });
          modifiedRequest.headers.set("X-Hex-Timestamp", hexTimestamp);

          // Pass the modified request to the origin
          const response = await fetch(modifiedRequest);

          return response;
        },
      }
      EOT
    }
  ]
  metadata = {
    main_module = "main.js"
  }
}
