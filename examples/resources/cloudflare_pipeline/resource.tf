resource "cloudflare_pipeline" "example_pipeline" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  destination = {
    batch = {
      max_bytes = 1000
      max_duration_s = 0.25
      max_rows = 100
    }
    compression = {
      type = "gzip"
    }
    credentials = {
      access_key_id = "<access key id>"
      endpoint = "https://123f8a8258064ed892a347f173372359.r2.cloudflarestorage.com"
      secret_access_key = "<secret key>"
    }
    format = "json"
    path = {
      bucket = "bucket"
      filename = "${slug}${extension}"
      filepath = "${date}/${hour}"
      prefix = "base"
    }
    type = "r2"
  }
  name = "sample_pipeline"
  source = [{
    format = "json"
    type = "type"
    authentication = true
    cors = {
      origins = ["*"]
    }
  }]
}
