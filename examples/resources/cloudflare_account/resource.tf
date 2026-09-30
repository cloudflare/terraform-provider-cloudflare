resource "cloudflare_account" "example_account" {
  name = "name"
  standalone = true
  type = "standard"
  unit = {
    id = "f267e341f3dd4697bd3b9f71dd96247f"
  }
}
