resource "cloudflare_teams_list" "example" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  name        = "Corporate devices"
  type        = "SERIAL"
  description = "Serial numbers for all corporate devices."
  items       = ["8GE8721REF", "5RE8543EGG", "1YE2880LNP"]
}
