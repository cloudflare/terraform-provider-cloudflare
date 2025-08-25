package ast_test

import (
	"strings"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/ast"
)

func TestDynamicBody(t *testing.T) {
	tf := `
resource "cloudflare_fallback_domain" "example" {
  account_id = "abc123"
  dynamic "domains" {
    for_each = toset(["intranet", "internal", "private", "localdomain", "domain", "lan", "home", "host", "corp", "local", "localhost", "home.arpa", "invalid", "test"])
    content {
      suffix = domains.value
    }
  }

  domains {
    suffix      = "example.com"
    description = "Example domain"
    dns_server  = ["192.0.2.0", "192.0.2.1"]
  }
}
`
	expected := `
resource "cloudflare_fallback_domain" "example" {
  account_id = "abc123"
  domains = concat([
    for domains in
    [
      for value in
      toset(["intranet", "internal", "private", "localdomain", "domain", "lan", "home", "host", "corp", "local", "localhost", "home.arpa", "invalid", "test"])
      : {
        key   = value
        value = value
      }
    ]
    : { content = { suffix = domains.value } }
    ], [{
      description = "Example domain"
      dns_server  = ["192.0.2.0", "192.0.2.1"]
      suffix      = "example.com"
  }])
}
`

	diags := ast.Diagnostics{}
	body := ast.ParseIntoSyntaxBody([]byte(tf), "test.tf", diags)
	b2 := ast.Block2AttrsRoot(*body, diags)
	actual := ast.FormatString(ast.Body2S(b2, diags))

	if actual != strings.Trim(expected, "\n") {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, actual)
	}
}
