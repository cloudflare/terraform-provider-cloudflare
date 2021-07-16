---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_access_identity_provider"
sidebar_current: "docs-cloudflare-resource-access-identity-provider"
description: |-
  Provides a Cloudflare Access Identity Provider resource.
---

# cloudflare_access_identity_provider

Provides a Cloudflare Access Identity Provider resource. Identity Providers are
used as an authentication or authorisation source within Access.

## Example Usage

```hcl
# one time pin
resource "cloudflare_access_identity_provider" "pin_login" {
  account_id = "1d5fdc9e88c8a8c4518b068cd94331fe"
  name       = "PIN login"
  type       = "onetimepin"
}

# oauth
resource "cloudflare_access_identity_provider" "github_oauth" {
  account_id = "1d5fdc9e88c8a8c4518b068cd94331fe"
  name       = "GitHub OAuth"
  type       = "github"
  config {
    client_id     = "example"
    client_secret = "secret_key"
  }
}

# saml
resource "cloudflare_access_identity_provider" "jumpcloud_saml" {
  account_id = "1d5fdc9e88c8a8c4518b068cd94331fe"
  name       = "JumpCloud SAML"
  type       = "saml"
  config {
    issuer_url = "jumpcloud"
    sso_target_url = "https://sso.myexample.jumpcloud.com/saml2/cloudflareaccess"
    attributes = [ "email", "username" ]
    sign_request = false
    idp_public_cert = "MIIDpDCCAoygAwIBAgIGAV2ka+55MA0GCSqGSIb3DQEBCwUAMIGSMQswCQ...GF/Q2/MHadws97cZg\nuTnQyuOqPuHbnN83d/2l1NSYKCbHt24o"
  }
}

# okta
resource "cloudflare_access_identity_provider" "okta" {
  account_id = "1d5fdc9e88c8a8c4518b068cd94331fe"
  name       = "Okta"
  type       = "okta"
  config {
    client_id     = "example"
    client_secret = "secret_key"
    api_token     = "okta_api_token"
  }
}
```

Please refer to the [developers.cloudflare.com Access documentation][access_identity_provider_guide]
for full reference on what is available and how to configure your provider.

## Argument Reference

The following arguments are supported:

-> **Note:** It's required that an `account_id` or `zone_id` is provided and in most cases using either is fine. However, if you're using a scoped access token, you must provide the argument that matches the token's scope. For example, an access token that is scoped to the "example.com" zone needs to use the `zone_id` argument.

* `account_id` - (Optional) The account ID the provider should be associated with. Conflicts with `zone_id`.
* `zone_id` - (Optional) The zone ID the provider should be associated with. Conflicts with `account_id`.
* `name` - (Required) Friendly name of the Access Identity Provider configuration.
* `type` - (Required) The provider type to use. Must be one of: `"centrify"`,
  `"facebook"`, `"google-apps"`, `"oidc"`, `"github"`, `"google"`, `"saml"`,
  `"linkedin"`, `"azureAD"`, `"okta"`, `"onetimepin"`, `"onelogin"`, `"yandex"`.
* `config` - (Optional) Provider configuration from the [developer documentation][access_identity_provider_guide].

## Attributes Reference

The following additional attributes are exported:

* `id` - ID of the Access Identity Provider
* `name` - Friendly name of the Access Identity Provider configuration.
* `type` - The provider type to use.
* `config` - Access Identity Provider configuration.

## Import

Access Identity Providers can be imported using a composite ID formed of account
ID and Access Identity Provider ID.

```
$ terraform import cloudflare_access_identity_provider.my_idp cb029e245cfdd66dc8d2e570d5dd3322/e00e1c13-e350-44fe-96c5-fb75c954871c
```

[access_identity_provider_guide]: https://developers.cloudflare.com/access/configuring-identity-providers/
