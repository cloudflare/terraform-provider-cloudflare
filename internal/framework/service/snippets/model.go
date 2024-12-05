package snippets

import "github.com/hashicorp/terraform-plugin-framework/types"

type Snippet struct {
	ZoneID      types.String  `tfsdk:"zone_id"`
	SnippetFile []SnippetFile `tfsdk:"files"`
	SnippetName types.String  `tfsdk:"name"`
	MainModule  types.String  `tfsdk:"main_module"`
}

type SnippetFile struct {
	FileName types.String `tfsdk:"name"`
	Content  types.String `tfsdk:"content"`
}
