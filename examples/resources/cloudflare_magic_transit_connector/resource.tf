resource "cloudflare_magic_transit_connector" "example_magic_transit_connector" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  device = {
    id = "id"
    serial_number = "serial_number"
  }
  activated = true
  interrupt_window_duration_hours = 0
  interrupt_window_hour_of_day = 0
  notes = "notes"
  timezone = "timezone"
}
