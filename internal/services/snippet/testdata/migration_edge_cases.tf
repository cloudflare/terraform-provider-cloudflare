resource "cloudflare_snippet" "%[1]s" {
  zone_id     = "%[2]s"
  name        = "%[1]s"
  main_module = "edge.js"

  files {
    name    = "edge.js"
    content = <<-JS
// Test edge cases with special characters and escaping
export default {
  async fetch(request) {
    const url = new URL(request.url);
    
    // Test various special characters
    const patterns = {
      "test": "value with \"quotes\"",
      'single': 'value with \'apostrophes\'',
      "backslash": "path\\with\\backslashes",
      "newline": "line1\nline2",
      "tab": "col1\tcol2",
      "unicode": "emoji 🚀 and symbols ™ © ®",
      "regex": "/^test.*\\.js$/",
      "html": "<script>alert('xss')</script>",
      "template": "$${variable} and $$$${escaped}",
    };
    
    // Check if path matches any test pattern
    for (const [key, value] of Object.entries(patterns)) {
      if (url.pathname.includes(key)) {
        return new Response(value, {
          headers: { "Content-Type": "text/plain" }
        });
      }
    }
    
    return fetch(request);
  }
};
JS
  }
}