resource "cloudflare_account" "%[1]s" {
  name = "%[2]s"
  type = "%[3]s"
  settings = {
    enforce_twofactor   = "%[4]t"
    abuse_contact_email = "%[5]s"
  }
}