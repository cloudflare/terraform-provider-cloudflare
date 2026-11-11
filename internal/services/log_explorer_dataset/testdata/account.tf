resource "cloudflare_log_explorer_dataset" "test" {
  account_id          = "%[1]s"
  dataset             = "magic_bgp_logs"
  deletion_protection = %[2]t
%[3]s}
