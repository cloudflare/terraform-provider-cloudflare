data "cloudflare_streams" "example_streams" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  id = "ea95132c15732412d22c1476fa83f27a"
  after = "2019-12-27T18:11:19.117Z"
  before = "2019-12-27T18:11:19.117Z"
  creator = "creator-id_abcde12345"
  end = "2014-01-02T02:20:00Z"
  limit = 1
  live_input_id = "live_input_id"
  name = "name"
  search = "puppy.mp4"
  start = "2014-01-02T02:20:00Z"
  status = "inprogress"
  type = "live"
  video_name = "puppy.mp4"
}
