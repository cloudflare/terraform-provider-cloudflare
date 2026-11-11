resource "cloudflare_log_explorer_dataset" "test" {
  zone_id             = "%[1]s"
  dataset             = "dns_logs"
  deletion_protection = %[2]t
%[3]s
}
