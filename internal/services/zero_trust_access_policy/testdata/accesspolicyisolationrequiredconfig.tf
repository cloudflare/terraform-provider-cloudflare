
	resource "cloudflare_zero_trust_gateway_settings" "%[1]s" {
		account_id = "%[3]s"
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
  }
		proxy = {
    tcp = true
		  udp = false
		  root_ca = true
		  virtual_ip = false
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
	  }

    resource "cloudflare_zero_trust_access_application" "%[1]s" {
      name       = "%[1]s"
      account_id = "%[3]s"
      domain     = "%[1]s.%[2]s"
	  depends_on = ["cloudflare_zero_trust_gateway_settings.%[1]s"]
    }

    resource "cloudflare_zero_trust_access_policy" "%[1]s" {
      application_id = cloudflare_zero_trust_access_application.%[1]s.id
      name           = "%[1]s"
      account_id     = "%[3]s"
      decision       = "allow"
      precedence     = "1"

      include =[ {
        email = ["a@example.com", "b@example.com"]
      }]

      isolation_required = "true"
    }

  