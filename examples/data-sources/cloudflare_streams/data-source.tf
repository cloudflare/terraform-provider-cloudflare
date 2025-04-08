data "cloudflare_streams" "example_streams" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  creator = "creator-id_abcde12345"
  end = "2014-01-02T02:20:00Z"
  search = "puppy.mp4"
  start = "2014-01-02T02:20:00Z"
  status = "inprogress"
  type = "live"
}
