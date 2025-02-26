// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package page_rule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*PageRuleDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"pagerule_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"created_on": schema.StringAttribute{
				Description: "The timestamp of when the Page Rule was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"id": schema.StringAttribute{
				Description: "Identifier",
				Computed:    true,
			},
			"modified_on": schema.StringAttribute{
				Description: "The timestamp of when the Page Rule was last modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"priority": schema.Int64Attribute{
				Description: "The priority of the rule, used to define which Page Rule is processed\nover another. A higher number indicates a higher priority. For example,\nif you have a catch-all Page Rule (rule A: `/images/*`) but want a more\nspecific Page Rule to take precedence (rule B: `/images/special/*`),\nspecify a higher priority for rule B so it overrides rule A.",
				Computed:    true,
			},
			"status": schema.StringAttribute{
				Description: "The status of the Page Rule.\navailable values: \"active\", \"disabled\"",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("active", "disabled"),
				},
			},
		},
	}
}

func (d *PageRuleDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *PageRuleDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
