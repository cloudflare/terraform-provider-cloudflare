data "cloudflare_filters" "example_filters" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
  id = "372e67954025e0ba6aaa6d586b9e0b61"
  description = "browsers"
  expression = "php"
  paused = false
  ref = "FIL-100"
}
