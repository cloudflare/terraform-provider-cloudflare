
resource "cloudflare_access_group" "%[1]s" {
  %[2]s_id = "%[3]s"
  name     = "%[1]s"

  include {
    common_name = "common"
  }
}