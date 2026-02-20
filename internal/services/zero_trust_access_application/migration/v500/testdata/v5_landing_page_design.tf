resource "cloudflare_zero_trust_access_application" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  type       = "app_launcher"

  landing_page_design = {
    title             = "Welcome to App"
    message           = "Please select an application"
    button_color      = "#0051c3"
    button_text_color = "#ffffff"
  }
}
