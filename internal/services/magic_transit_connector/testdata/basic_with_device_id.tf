resource "cloudflare_magic_transit_connector" "%[1]s" {
  account_id = "%[2]s"
  device = {
    id = "%[3]s"
  }
  activated = %[4]s
  notes     = "%[5]s"
  interrupt_window_duration_hours = %[6]s
  interrupt_window_hour_of_day    = %[7]s
}