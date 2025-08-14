package ast_test

import (
	"strings"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/ast"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

func TestWarning(t *testing.T) {
	str := `
resource "dog" "dog" {
	dog = {
		poodle = "big"
	}
}
`
	parsed, _ := hclwrite.ParseConfig([]byte(str), "dog.hcl", hcl.InitialPos)
	attr := parsed.Body().Blocks()[0].Body().GetAttribute("dog")

	cases := map[string]string{
		string(ast.WarnManualMigration4AttrWrite("resources/account_member", nil).BuildTokens(nil).Bytes()): `
regex(<<-WARNING

  > We were unable to automatically migrate this resource to the new provider.
  > Please refer to "https://registry.terraform.io/providers/cloudflare/cloudflare/latest/docs/resources/account_member" for manual migration.

WARNING
, "")`,
		string(ast.WarnManualMigration4AttrWrite("resources/account_member", attr).BuildTokens(nil).Bytes()): `
regex(<<-WARNING

  > We were unable to automatically migrate this resource to the new provider.
  > Please refer to "https://registry.terraform.io/providers/cloudflare/cloudflare/latest/docs/resources/account_member" for manual migration.

  > Your previous configuration was:
   dog = {
    poodle = "big"
   }

WARNING
, "")`,
	}
	for str, expected := range cases {
		expected := strings.ReplaceAll(strings.Trim(expected, " \n"), "\t", " ")
		if str != expected {
			t.Errorf("Expected:\n%q\nActual:\n%q", expected, str)
		}
	}

}
