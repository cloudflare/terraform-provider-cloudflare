resource "cloudflare_zero_trust_organization" "example_zero_trust_organization" {
  zone_id = "zone_id"
  allow_authenticate_via_warp = true
  auth_domain = "test.cloudflareaccess.com"
  auto_redirect_to_identity = true
  custom_pages = {
    forbidden = "699d98642c564d2e855e9661899b7252"
    identity_denied = "699d98642c564d2e855e9661899b7252"
  }
  deny_unmatched_requests = true
  deny_unmatched_requests_exempted_zone_names = ["example.com"]
  is_ui_read_only = true
  login_design = {
    background_color = "#c5ed1b"
    footer_text = "This is an example description."
    header_text = "This is an example description."
    logo_path = "https://example.com/logo.png"
    text_color = "#c5ed1b"
  }
  mfa_config = {
    allowed_authenticators = ["totp", "biometrics", "security_key"]
    amr_matching_session_duration = "12h"
    required_aaguids = "2fc0579f-8113-47ea-b116-bb5a8db9202a"
    session_duration = "24h"
  }
  mfa_required_for_all_apps = false
  mfa_ssh_piv_key_requirements = {
    pin_policy = "always"
    require_fips_device = true
    ssh_key_size = [256, 2048]
    ssh_key_type = ["ecdsa", "rsa"]
    touch_policy = "always"
  }
  name = "Widget Corps Internal Applications"
  session_duration = "24h"
  ui_read_only_toggle_reason = "Temporarily turn off the UI read only lock to make a change via the UI"
  user_seat_expiration_inactive_time = "730h"
  warp_auth_session_duration = "24h"
}
