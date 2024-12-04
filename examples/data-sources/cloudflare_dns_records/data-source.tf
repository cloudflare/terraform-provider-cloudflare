data "cloudflare_dns_records" "example_dns_records" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
  comment = {
    absent = "absent"
    contains = "ello, worl"
    endswith = "o, world"
    exact = "Hello, world"
    present = "present"
    startswith = "Hello, w"
  }
  content = {
    contains = "7.0.0."
    endswith = ".0.1"
    exact = "127.0.0.1"
    startswith = "127.0."
  }
  name = {
    contains = "w.example."
    endswith = ".example.com"
    exact = "www.example.com"
    startswith = "www.example"
  }
  search = "www.cloudflare.com"
  tag = {
    absent = "important"
    contains = "greeting:ello, worl"
    endswith = "greeting:o, world"
    exact = "greeting:Hello, world"
    present = "important"
    startswith = "greeting:Hello, w"
  }
  type = "A"
}
