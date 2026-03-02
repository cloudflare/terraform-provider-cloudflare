resource "cloudflare_pipeline" "example_pipeline" {
  account_id = "0123105f4ecef8ad9ca31a8372d0c353"
  name = "my_pipeline"
  sql = "insert into sink select * from source;"
}
