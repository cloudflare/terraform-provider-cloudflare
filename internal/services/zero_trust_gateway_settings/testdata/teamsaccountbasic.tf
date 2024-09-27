
resource "cloudflare_zero_trust_gateway_settings" "%[1]s" {
  account_id = "%[2]s"
  tls_decrypt_enabled = true
  protocol_detection_enabled = true
  activity_log_enabled = true
  url_browser_isolation_enabled = true
  non_identity_browser_isolation_enabled = false
  block_page = {
    name = "%[1]s"
    enabled = true
    footer_text = "hello"
    header_text = "hello"
    logo_path = "https://example.com"
    background_color = "#000000"
	mailto_subject = "hello"
	mailto_address = "test@cloudflare.com"
  }
  body_scanning = {
    inspection_mode = "deep"
  }
  fips = {
    tls = true
  }
  antivirus = {
    enabled_download_phase = true
    enabled_upload_phase = false
    fail_closed = true
	notification_settings = {
		enabled = true
		message = "msg"
		support_url = "https://hello.com/"
	}
  }
  proxy = {
    tcp = true
    udp = false
	root_ca = true
	virtual_ip = true
  }
  logging = {
    redact_pii = true
    settings_by_rule_type =[ {
      dns = {
    log_all = false
        log_blocks = true
  }
      http = {
    log_all = true
        log_blocks = true
  }
      l4 = {
    log_all = false
        log_blocks = true
  }
    }]
  }
  ssh_session_log = {
    public_key = "testvSXw8BfbrGCi0fhGiD/3yXk2SiV1Nzg2lru3oj0="
  }
  payload_log = {
    public_key = "EmpOvSXw8BfbrGCi0fhGiD/3yXk2SiV1Nzg2lru3oj0="
  }
  extended_email_matching =[ {
	enabled = true
  }]
  custom_certificate =[ {
	enabled = false
  }]
}
