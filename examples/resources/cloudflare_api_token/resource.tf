# User permissions
data "cloudflare_api_token_permission_groups" "all" {}

# Token allowed to create new tokens.
# Can only be used from specific ip range.
resource "cloudflare_api_token" "api_token_create" {
  name = "api_token_create"

  policy {
    permission_groups = [
      data.cloudflare_api_token_permission_groups.all.user["API Tokens Write"],
    ]
    resources = {
      "com.cloudflare.api.user.${var.user_id}" = "*"
    }
  }

  not_before = "2018-07-01T05:20:00Z"
  expires_on = "2020-01-01T00:00:00Z"

  condition {
    request_ip {
      in     = ["192.0.2.1/32"]
      not_in = ["198.51.100.1/32"]
    }
  }
}

# Account permissions
data "cloudflare_api_token_permission_groups" "all" {}

# Token allowed to read audit logs from all accounts.
resource "cloudflare_api_token" "logs_account_all" {
  name = "logs_account_all"

  policy {
    permission_groups = [
      data.cloudflare_api_token_permission_groups.all.account["Access: Audit Logs Read"],
    ]
    resources = {
      "com.cloudflare.api.account.*" = "*"
    }
  }
}

# Token allowed to read audit logs from specific account.
resource "cloudflare_api_token" "logs_account" {
  name = "logs_account"

  policy {
    permission_groups = [
      data.cloudflare_api_token_permission_groups.all.account["Access: Audit Logs Read"],
    ]
    resources = {
      "com.cloudflare.api.account.${var.account_id}" = "*"
    }
  }
}

# Zone permissions
data "cloudflare_api_token_permission_groups" "all" {}

# Token allowed to edit DNS entries and TLS certs for specific zone.
resource "cloudflare_api_token" "dns_tls_edit" {
  name = "dns_tls_edit"

  policy {
    permission_groups = [
      data.cloudflare_api_token_permission_groups.all.zone["DNS Write"],
      data.cloudflare_api_token_permission_groups.all.zone["SSL and Certificates Write"],
    ]
    resources = {
      "com.cloudflare.api.account.zone.${var.zone_id}" = "*"
    }
  }
}

# Token allowed to edit DNS entries for all zones except one.
resource "cloudflare_api_token" "dns_tls_edit_all_except_one" {
  name = "dns_tls_edit_all_except_one"

  # include all zones
  policy {
    permission_groups = [
      data.cloudflare_api_token_permission_groups.all.zone["DNS Write"],
    ]
    resources = {
      "com.cloudflare.api.account.zone.*" = "*"
    }
  }

  # exclude (deny) specific zone
  policy {
    permission_groups = [
      data.cloudflare_api_token_permission_groups.all.zone["DNS Write"],
    ]
    resources = {
      "com.cloudflare.api.account.zone.${var.zone_id}" = "*"
    }
    effect = "deny"
  }
}


# Token allowed to edit DNS entries for all zones from specific account.
resource "cloudflare_api_token" "dns_edit_all_account" {
  name = "dns_edit_all_account"

  # include all zones from specific account
  policy {
    permission_groups = [
      data.cloudflare_api_token_permission_groups.all.zone["DNS Write"],
    ]
    resources = {
      "com.cloudflare.api.account.${var.account_id}" = jsonencode({
        "com.cloudflare.api.account.zone.*" = "*"
      })
    }
  }
}
