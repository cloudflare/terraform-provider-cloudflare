resource "cloudflare_magic_transit_connector" "%[1]s" {
  account_id = "%[2]s"
  device = {
    provision_license = true
  }
  activated = %[3]s
  notes     = "%[4]s"
  interrupt_window_duration_hours = %[5]s
  interrupt_window_hour_of_day    = %[6]s
}