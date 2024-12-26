package snippet_rules

import "github.com/hashicorp/terraform-plugin-framework/types"

type SnippetRules struct {
	ZoneID types.String  `tfsdk:"zone_id"`
	Rules  []SnippetRule `tfsdk:"rules"`
}

type SnippetRule struct {
	Enabled     types.Bool   `tfsdk:"enabled"`
	Expression  types.String `tfsdk:"expression"`
	Description types.String `tfsdk:"description"`
	SnippetName types.String `tfsdk:"snippet_name"`
}
