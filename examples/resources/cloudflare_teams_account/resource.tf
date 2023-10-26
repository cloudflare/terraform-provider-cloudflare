resource "cloudflare_teams_account" "example" {
  account_id                 = "f037e56e89293a057740de681ac9abbe"
  tls_decrypt_enabled        = true
  protocol_detection_enabled = true

  block_page {
    footer_text      = "hello"
    header_text      = "hello"
    logo_path        = "https://example.com/logo.jpg"
    background_color = "#000000"
  }

  body_scanning {
    inspection_mode = "deep"
  }

  antivirus {
    enabled_download_phase = true
    enabled_upload_phase   = false
    fail_closed            = true
  }

  fips {
    tls = true
  }

  proxy {
    tcp     = true
    udp     = true
    root_ca = true
  }

  url_browser_isolation_enabled = true

  logging {
    redact_pii = true
    settings_by_rule_type {
      dns {
        log_all    = false
        log_blocks = true
      }
      http {
        log_all    = true
        log_blocks = true
      }
      l4 {
        log_all    = false
        log_blocks = true
      }
    }
  }
}
