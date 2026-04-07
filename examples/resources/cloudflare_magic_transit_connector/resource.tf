resource "cloudflare_magic_transit_connector" "example_magic_transit_connector" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  device = {
    id = "id"
    provision_license = true
    serial_number = "serial_number"
  }
  activated = true
  interrupt_window_days_of_week = ["Sunday"]
  interrupt_window_duration_hours = 1
  interrupt_window_embargo_dates = ["string"]
  interrupt_window_hour_of_day = 0
  notes = "notes"
  timezone = "timezone"
}
