resource "cloudflare_teams_list" "example" {
  account_id  = "1d5fdc9e88c8a8c4518b068cd94331fe"
  name        = "Corporate devices"
  type        = "SERIAL"
  description = "Serial numbers for all corporate devices."
  items       = ["8GE8721REF", "5RE8543EGG", "1YE2880LNP"]
}
