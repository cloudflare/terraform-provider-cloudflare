---
page_title: "cloudflare_zero_trust_access_identity_provider Resource - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_zero_trust_access_identity_provider (Resource)



## Example Usage

```terraform
# one time pin
resource "cloudflare_zero_trust_access_identity_provider" "pin_login" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "PIN login"
  type       = "onetimepin"
}

# oauth
resource "cloudflare_zero_trust_access_identity_provider" "github_oauth" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "GitHub OAuth"
  type       = "github"
  config = {
    client_id     = "example"
    client_secret = "secret_key"
  }
}

# saml
resource "cloudflare_zero_trust_access_identity_provider" "jumpcloud_saml" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "JumpCloud SAML"
  type       = "saml"
  config = {
    issuer_url      = "jumpcloud"
    sso_target_url  = "https://sso.myexample.jumpcloud.com/saml2/cloudflareaccess"
    attributes      = ["email", "username"]
    sign_request    = false
    idp_public_cert = "MIIDpDCCAoygAwIBAgIGAV2ka+55MA0GCSqGSIb3DQEBCwUAMIGSMQswCQ...GF/Q2/MHadws97cZg\nuTnQyuOqPuHbnN83d/2l1NSYKCbHt24o"
  }
}

# okta
resource "cloudflare_zero_trust_access_identity_provider" "okta" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "Okta"
  type       = "okta"
  config = {
    client_id     = "example"
    client_secret = "secret_key"
    api_token     = "okta_api_token"
    okta_account  = "https://example.com"
  }
}
```
<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `config` (Attributes) The configuration parameters for the identity provider. To view the required parameters for a specific provider, refer to our [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/). (see [below for nested schema](#nestedatt--config))
- `name` (String) The name of the identity provider, shown to users on the login page.
- `type` (String) The type of identity provider. To determine the value for a specific provider, refer to our [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).

### Optional

- `account_id` (String) The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
- `scim_config` (Attributes) The configuration settings for enabling a System for Cross-Domain Identity Management (SCIM) with the identity provider. (see [below for nested schema](#nestedatt--scim_config))
- `zone_id` (String) The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.

### Read-Only

- `id` (String) UUID

<a id="nestedatt--config"></a>
### Nested Schema for `config`

Optional:

- `apps_domain` (String) Your companies TLD
- `attributes` (List of String) A list of SAML attribute names that will be added to your signed JWT token and can be used in SAML policy rules.
- `auth_url` (String) The authorization_endpoint URL of your IdP
- `authorization_server_id` (String) Your okta authorization server id
- `centrify_account` (String) Your centrify account url
- `centrify_app_id` (String) Your centrify app id
- `certs_url` (String) The jwks_uri endpoint of your IdP to allow the IdP keys to sign the tokens
- `claims` (List of String) Custom claims
- `client_id` (String) Your OAuth Client ID
- `client_secret` (String) Your OAuth Client Secret
- `conditional_access_enabled` (Boolean) Should Cloudflare try to load authentication contexts from your account
- `directory_id` (String) Your Azure directory uuid
- `email_attribute_name` (String) The attribute name for email in the SAML response.
- `email_claim_name` (String) The claim name for email in the id_token response.
- `header_attributes` (Attributes List) Add a list of attribute names that will be returned in the response header from the Access callback. (see [below for nested schema](#nestedatt--config--header_attributes))
- `idp_public_certs` (List of String) X509 certificate to verify the signature in the SAML authentication response
- `issuer_url` (String) IdP Entity ID or Issuer URL
- `okta_account` (String) Your okta account url
- `onelogin_account` (String) Your OneLogin account url
- `ping_env_id` (String) Your PingOne environment identifier
- `prompt` (String) Indicates the type of user interaction that is required. prompt=login forces the user to enter their credentials on that request, negating single-sign on. prompt=none is the opposite. It ensures that the user isn't presented with any interactive prompt. If the request can't be completed silently by using single-sign on, the Microsoft identity platform returns an interaction_required error. prompt=select_account interrupts single sign-on providing account selection experience listing all the accounts either in session or any remembered account or an option to choose to use a different account altogether.
- `scopes` (List of String) OAuth scopes
- `sign_request` (Boolean) Sign the SAML authentication request with Access credentials. To verify the signature, use the public key from the Access certs endpoints.
- `sso_target_url` (String) URL to send the SAML authentication requests to
- `support_groups` (Boolean) Should Cloudflare try to load groups from your account
- `token_url` (String) The token_endpoint URL of your IdP

<a id="nestedatt--config--header_attributes"></a>
### Nested Schema for `config.header_attributes`

Optional:

- `attribute_name` (String) attribute name from the IDP
- `header_name` (String) header that will be added on the request to the origin



<a id="nestedatt--scim_config"></a>
### Nested Schema for `scim_config`

Optional:

- `enabled` (Boolean) A flag to enable or disable SCIM for the identity provider.
- `group_member_deprovision` (Boolean) A flag to revoke a user's session in Access and force a reauthentication on the user's Gateway session when they have been added or removed from a group in the Identity Provider.
- `seat_deprovision` (Boolean) A flag to remove a user's seat in Zero Trust when they have been deprovisioned in the Identity Provider.  This cannot be enabled unless user_deprovision is also enabled.
- `secret` (String) A read-only token generated when the SCIM integration is enabled for the first time.  It is redacted on subsequent requests.  If you lose this you will need to refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
- `user_deprovision` (Boolean) A flag to enable revoking a user's session in Access and Gateway when they have been deprovisioned in the Identity Provider.

## Import

Import is supported using the following syntax:

```shell
$ terraform import cloudflare_zero_trust_access_identity_provider.example <account_id>/<identity_provider_id>
```