// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package page_rule

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*PageRuleDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Access: Apps and Policies Read",
				"Access: Apps and Policies Revoke",
				"Access: Apps and Policies Write",
				"Access: Mutual TLS Certificates Write",
				"Access: Organizations, Identity Providers, and Groups Write",
				"Analytics Read",
				"Apps Write",
				"Cache Purge",
				"DNS Read",
				"DNS Write",
				"Firewall Services Read",
				"Firewall Services Write",
				"Load Balancers Read",
				"Load Balancers Write",
				"Logs Read",
				"Logs Write",
				"Page Rules Read",
				"Page Rules Write",
				"SSL and Certificates Read",
				"SSL and Certificates Write",
				"Stream Read",
				"Stream Write",
				"Trust and Safety Read",
				"Trust and Safety Write",
				"Workers Routes Read",
				"Workers Routes Write",
				"Workers Scripts Read",
				"Workers Scripts Write",
				"Zaraz Admin",
				"Zaraz Edit",
				"Zaraz Read",
				"Zero Trust: PII Read",
				"Zone Read",
				"Zone Settings Read",
				"Zone Settings Write",
				"Zone Write",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Identifier.",
				Computed:    true,
			},
			"pagerule_id": schema.StringAttribute{
				Description: "Identifier.",
				Required:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "Identifier.",
				Required:    true,
			},
			"created_on": schema.StringAttribute{
				Description: "The timestamp of when the Page Rule was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
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
				Description: "The status of the Page Rule.\nAvailable values: \"active\", \"disabled\".",
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
