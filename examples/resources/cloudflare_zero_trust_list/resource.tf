resource "cloudflare_zero_trust_list" "example_zero_trust_list" {
  account_id = "699d98642c564d2e855e9661899b7252"
  name = "Admin Serial Numbers"
  type = "SERIAL"
  description = "The serial numbers for administrators"
  items = [{
    created_at = "2014-01-01T05:20:00.12345Z"
    description = "Austin office IP"
    value = "8GE8721REF"
  }]
}
