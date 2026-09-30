# Exactly one vendor and exactly one authentication method must be configured.
#
# Confidential arguments (API keys, service account keys) are write-only:
# Terraform sends them to the provider but does not store them directly in state.
# Only a secure digest is stored so Terraform can detect secret rotation. Rotating
# a secret updates the integration in place; it does not recreate it.
#
# Non-confidential arguments are stored in state and diffed normally.
#
# Changing the vendor or the authentication method does force the integration to
# be recreated, because the update API cannot change either one.
#
# Write-only arguments require Terraform 1.11 or later, and may be fed from
# ephemeral resources so the secret never touches disk.

resource "cloudflare_zero_trust_casb_integration" "google_cloud_platform" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  name       = "My Google Cloud Platform integration"
  paused     = false

  google_cloud_platform = {
    google_cloud_platform_service_account = {
      service_account_key_json = file("service-account-key.json")
    }
  }
}

resource "cloudflare_zero_trust_casb_integration" "google_workspace" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  name       = "My Google Workspace integration"
  paused     = false

  google_workspace = {
    google_domain_wide_delegation_service_account = {
      service_account_key_json = file("service-account-key.json")
      administrator_email      = "admin@example.com"
    }
  }
}

resource "cloudflare_zero_trust_casb_integration" "anthropic" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  name       = "My Anthropic integration"
  paused     = false

  anthropic = {
    # One of anthropic_admin_api_key, anthropic_workspace_api_key or
    # anthropic_compliance_api_key.
    anthropic_admin_api_key = {
      api_key = var.anthropic_admin_api_key

      # Optional: auto-extracted from the key when omitted.
      tenant_id = "org_01234567-89ab-cdef-0123-456789abcdef"
    }
  }
}

resource "cloudflare_zero_trust_casb_integration" "openai" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  name       = "My OpenAI integration"
  paused     = false

  # Also enrol the integration in continuous evaluation and let it remediate
  # findings automatically. Defaults to ["casb"] when omitted.
  use_cases = ["casb", "ces", "auto_remediation"]

  openai = {
    # chatgpt_compliance_api_key requires an OpenAI Enterprise plan.
    chatgpt_standard_api_key = {
      admin_api_key   = var.openai_admin_api_key
      organization_id = "org-abc123"

      # Optional: only needed for DLP scanning.
      project_api_key = var.openai_project_api_key
      project_id      = "proj_abc123"
    }
  }
}

resource "cloudflare_zero_trust_casb_integration" "aws" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  name       = "My AWS integration"
  paused     = false

  aws = {
    # Nothing here is confidential: the external ID is useless without
    # Cloudflare's AWS principal, so no secrets_digest is needed to track it.
    aws_iam_role = {
      role_arn    = "arn:aws:iam::123456789012:role/Cloudflare_DSPM_Auditor"
      external_id = var.aws_external_id
    }
  }
}

# Before creating a Box integration, add the Cloudflare CASB application to
# your Box account:
#
# 1. Sign in to the Box Admin Console with Admin permission.
# 2. Open Integrations > Platform Apps Manager and select Server Authentication Apps.
# 3. Click Add Platform App and enter the Cloudflare CASB client ID:
#    puaghckpy0578r8p6f3g0rf860unup4r
# 4. In Accounts & Billing, copy the Enterprise ID and use it below.
#
# For details, refer to:
# https://developer.box.com/guides/authorization/platform-app-approval/#as-an-admin
resource "cloudflare_zero_trust_casb_integration" "box" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  name       = "My Box integration"
  paused     = false

  # Restrict DLP scanning to specific profiles. Omit to let Cloudflare assign
  # the account defaults.
  dlp_profiles = ["b0d1e2f3-4567-89ab-cdef-0123456789ab"]

  box = {
    box_server_authentication = {
      enterprise_id = "1234567"
    }
  }
}
