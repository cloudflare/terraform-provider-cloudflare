package main

import (
	"testing"
)

func TestZeroTrustAccessGroupTransformation(t *testing.T) {
	tests := []TestCase{
		{
			Name: "basic boolean attributes to empty objects",
			Config: `resource "cloudflare_zero_trust_access_group" "example" {
  name = "example-group"
  include = [{
    everyone = true
    certificate = true
    any_valid_service_token = true
  }]
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_group" "example" {
  name = "example-group"
  include = [{
    everyone = {}
  }, {
    certificate = {}
  }, {
    any_valid_service_token = {}
  }]
}`},
		},
		{
			Name: "array expansion for simple attributes",
			Config: `resource "cloudflare_zero_trust_access_group" "example" {
  name = "example-group"
  include = [{
    email = ["user1@example.com", "user2@example.com"]
    ip = ["192.0.2.1/32", "192.0.2.2/32"]
    geo = ["US", "CA"]
  }]
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_group" "example" {
  name = "example-group"
  include = [{
    email = {
      email = "user1@example.com"
    }
  }, {
    email = {
      email = "user2@example.com"
    }
  }, {
    ip = {
      ip = "192.0.2.1/32"
    }
  }, {
    ip = {
      ip = "192.0.2.2/32"
    }
  }, {
    geo = {
      country_code = "US"
    }
  }, {
    geo = {
      country_code = "CA"
    }
  }]
}`},
		},
		{
			Name: "email domain and other list attributes",
			Config: `resource "cloudflare_zero_trust_access_group" "example" {
  name = "example-group"
  include = [{
    email_domain = ["company.com", "trusted.com"]
    group = ["group1", "group2"]
    service_token = ["token1", "token2"]
  }]
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_group" "example" {
  name = "example-group"
  include = [{
    email_domain = {
      domain = "company.com"
    }
  }, {
    email_domain = {
      domain = "trusted.com"
    }
  }, {
    group = {
      id = "group1"
    }
  }, {
    group = {
      id = "group2"
    }
  }, {
    service_token = {
      token_id = "token1"
    }
  }, {
    service_token = {
      token_id = "token2"
    }
  }]
}`},
		},
		{
			Name: "common_names array expansion",
			Config: `resource "cloudflare_zero_trust_access_group" "example" {
  name = "example-group"
  include = [{
    common_names = ["cert1.example.com", "cert2.example.com"]
  }]
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_group" "example" {
  name = "example-group"
  include = [{
    common_name = {
      common_name = "cert1.example.com"
    }
  }, {
    common_name = {
      common_name = "cert2.example.com"
    }
  }]
}`},
		},
		{
			Name: "azure blocks renamed to azure_ad with ID expansion",
			Config: `resource "cloudflare_zero_trust_access_group" "example" {
  name = "example-group"
  include = [{
    azure = [{
      id = ["group1", "group2"]
      identity_provider_id = "azure-provider"
    }]
  }]
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_group" "example" {
  name = "example-group"
  include = [{
    azure_ad = {
      id                   = "group1"
      identity_provider_id = "azure-provider"
    }
  }, {
    azure_ad = {
      id                   = "group2"
      identity_provider_id = "azure-provider"
    }
  }]
}`},
		},
		{
			Name: "github blocks renamed to github_organization with teams expansion",
			Config: `resource "cloudflare_zero_trust_access_group" "example" {
  name = "example-group"
  include = [{
    github = [{
      name = "example-org"
      teams = ["team1", "team2", "team3"]
      identity_provider_id = "github-provider"
    }]
  }]
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_group" "example" {
  name = "example-group"
  include = [{
    github_organization = {
      name                 = "example-org"
      team                 = "team1"
      identity_provider_id = "github-provider"
    }
  }, {
    github_organization = {
      name                 = "example-org"
      team                 = "team2"
      identity_provider_id = "github-provider"
    }
  }, {
    github_organization = {
      name                 = "example-org"
      team                 = "team3"
      identity_provider_id = "github-provider"
    }
  }]
}`},
		},
		{
			Name: "gsuite blocks with email expansion",
			Config: `resource "cloudflare_zero_trust_access_group" "example" {
  name = "example-group"
  include = [{
    gsuite = [{
      email = ["user1@gsuite.com", "user2@gsuite.com"]
      identity_provider_id = "gsuite-provider"
    }]
  }]
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_group" "example" {
  name = "example-group"
  include = [{
    gsuite = {
      email                = "user1@gsuite.com"
      identity_provider_id = "gsuite-provider"
    }
  }, {
    gsuite = {
      email                = "user2@gsuite.com"
      identity_provider_id = "gsuite-provider"
    }
  }]
}`},
		},
		{
			Name: "okta blocks with name expansion",
			Config: `resource "cloudflare_zero_trust_access_group" "example" {
  name = "example-group"
  include = [{
    okta = [{
      name = ["group1", "group2"]
      identity_provider_id = "okta-provider"
    }]
  }]
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_group" "example" {
  name = "example-group"
  include = [{
    okta = {
      name                 = "group1"
      identity_provider_id = "okta-provider"
    }
  }, {
    okta = {
      name                 = "group2"
      identity_provider_id = "okta-provider"
    }
  }]
}`},
		},
		{
			Name: "saml blocks preserved as-is",
			Config: `resource "cloudflare_zero_trust_access_group" "example" {
  name = "example-group"
  include = [{
    saml = [{
      attribute_name = "department"
      attribute_value = "engineering"
      identity_provider_id = "saml-provider"
    }]
  }]
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_group" "example" {
  name = "example-group"
  include = [{
    saml = {
      attribute_name       = "department"
      attribute_value      = "engineering"
      identity_provider_id = "saml-provider"
    }
  }]
}`},
		},
		{
			Name: "external_evaluation blocks preserved as-is",
			Config: `resource "cloudflare_zero_trust_access_group" "example" {
  name = "example-group"
  include = [{
    external_evaluation = [{
      evaluate_url = "https://example.com/evaluate"
      keys_url = "https://example.com/keys"
    }]
  }]
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_group" "example" {
  name = "example-group"
  include = [{
    external_evaluation = {
      evaluate_url = "https://example.com/evaluate"
      keys_url     = "https://example.com/keys"
    }
  }]
}`},
		},
		{
			Name: "exclude and require rules transformation",
			Config: `resource "cloudflare_zero_trust_access_group" "example" {
  name = "example-group"
  include = [{
    email = ["user@company.com"]
  }]
  exclude = [{
    ip = ["192.168.1.100/32"]
    geo = ["CN"]
  }]
  require = [{
    email_domain = ["company.com"]
    certificate = true
  }]
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_group" "example" {
  name = "example-group"
  include = [{
    email = {
      email = "user@company.com"
    }
  }]
  exclude = [{
    ip = {
      ip = "192.168.1.100/32"
    }
  }, {
    geo = {
      country_code = "CN"
    }
  }]
  require = [{
    email_domain = {
      domain = "company.com"
    }
  }, {
    certificate = {}
  }]
}`},
		},
		{
			Name: "complex mixed scenario with all rule types",
			Config: `resource "cloudflare_zero_trust_access_group" "example" {
  name = "example-group"
  include = [{
    email = ["admin@company.com", "manager@company.com"]
    ip = ["10.0.1.0/24"]
    azure = [{
      id = ["admin-group", "dev-group"]
      identity_provider_id = "azure-ad-provider"
    }]
    github = [{
      name = "company-org"
      teams = ["backend-team", "frontend-team"]
      identity_provider_id = "github-provider"
    }]
    common_names = ["client1.company.com", "client2.company.com"]
    everyone = true
    certificate = true
  }]
  exclude = [{
    email = ["blocked@company.com"]
    geo = ["RU"]
  }]
  require = [{
    email_domain = ["company.com"]
    saml = [{
      attribute_name = "role"
      attribute_value = "employee"
      identity_provider_id = "saml-provider"
    }]
  }]
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_group" "example" {
  name = "example-group"
  include = [{
    email = {
      email = "admin@company.com"
    }
  }, {
    email = {
      email = "manager@company.com"
    }
  }, {
    ip = {
      ip = "10.0.1.0/24"
    }
  }, {
    azure_ad = {
      id                   = "admin-group"
      identity_provider_id = "azure-ad-provider"
    }
  }, {
    azure_ad = {
      id                   = "dev-group"
      identity_provider_id = "azure-ad-provider"
    }
  }, {
    github_organization = {
      name                 = "company-org"
      team                 = "backend-team"
      identity_provider_id = "github-provider"
    }
  }, {
    github_organization = {
      name                 = "company-org"
      team                 = "frontend-team"
      identity_provider_id = "github-provider"
    }
  }, {
    common_name = {
      common_name = "client1.company.com"
    }
  }, {
    common_name = {
      common_name = "client2.company.com"
    }
  }, {
    everyone = {}
  }, {
    certificate = {}
  }]
  exclude = [{
    email = {
      email = "blocked@company.com"
    }
  }, {
    geo = {
      country_code = "RU"
    }
  }]
  require = [{
    email_domain = {
      domain = "company.com"
    }
  }, {
    saml = {
      attribute_name       = "role"
      attribute_value      = "employee"
      identity_provider_id = "saml-provider"
    }
  }]
}`},
		},
		{
			Name: "additional list attributes",
			Config: `resource "cloudflare_zero_trust_access_group" "example" {
  name = "example-group"
  include = [{
    email_list = ["list1", "list2"]
    ip_list = ["iplist1", "iplist2"]
    login_method = ["method1", "method2"]
    device_posture = ["posture1", "posture2"]
  }]
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_group" "example" {
  name = "example-group"
  include = [{
    email_list = {
      id = "list1"
    }
  }, {
    email_list = {
      id = "list2"
    }
  }, {
    ip_list = {
      id = "iplist1"
    }
  }, {
    ip_list = {
      id = "iplist2"
    }
  }, {
    login_method = {
      id = "method1"
    }
  }, {
    login_method = {
      id = "method2"
    }
  }, {
    device_posture = {
      integration_uid = "posture1"
    }
  }, {
    device_posture = {
      integration_uid = "posture2"
    }
  }]
}`},
		},
	}

	RunTransformationTests(t, tests, transformFile)
}