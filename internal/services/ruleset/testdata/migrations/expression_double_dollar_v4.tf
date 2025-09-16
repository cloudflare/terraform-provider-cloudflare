locals {
  disallowed_regions_expression = "ip.geoip.country in {\"CN\" \"RU\" \"IR\"}"
}

variable "android_chat_openai_com_domain" {
  type    = string
  default = "android-chat.openai.com"
}

variable "ios_chat_openai_com_domain" {
  type    = string
  default = "ios-chat.openai.com"
}

variable "api_chatgpt_com_domain" {
  type    = string
  default = "api.chatgpt.com"
}

variable "api_openai_com_domain" {
  type    = string
  default = "api.openai.com"
}

variable "auth_openai_com_domain" {
  type    = string
  default = "auth.openai.com"
}

resource "cloudflare_ruleset" "%[2]s" {
  zone_id     = "%[1]s"
  name        = "Disallowed Countries %[2]s"
  description = "Test ruleset for expression double dollar issue"
  kind        = "zone"
  phase       = "http_request_firewall_custom"

  # Rule with complex heredoc expression containing multiple variables
  rules {
    action = "block"
    action_parameters {
      response {
        content = jsonencode({
          error = {
            message = "Country, region, or territory not supported"
            type    = "request_forbidden"
            param   = null
            code    = "unsupported_country_region_territory"
          }
        })
        content_type = "application/json"
        status_code  = 403
      }
    }
    description = "Block API traffic from disallowed countries"
    enabled     = true
    expression  = <<EOF
    ${local.disallowed_regions_expression}
    and (
      (cf.zone.name in {"${var.android_chat_openai_com_domain}" "${var.ios_chat_openai_com_domain}" "${var.api_chatgpt_com_domain}" "${var.api_openai_com_domain}"})
      or
      (
        (cf.zone.name eq "${var.auth_openai_com_domain}") and
        (http.request.uri.path matches "^/(api|oauth)/.*")
      )
    )
    EOF
  }

  # Rule with simple heredoc expression
  rules {
    action = "log"
    expression = <<EOF
    ${local.disallowed_regions_expression}
    EOF
    description = "Log simple condition"
    enabled = true
  }
}