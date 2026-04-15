// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package leaked_credential_check_rule

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*LeakedCredentialCheckRuleDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Account WAF Read",
				"Account WAF Write",
				"Zone WAF Read",
				"Zone WAF Write",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Defines the unique ID for this custom detection.",
				Computed:    true,
			},
			"detection_id": schema.StringAttribute{
				Description: "Defines the unique ID for this custom detection.",
				Required:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "Defines an identifier.",
				Optional:    true,
			},
			"password": schema.StringAttribute{
				Description: "Defines ehe ruleset expression to use in matching the password in a request.",
				Computed:    true,
			},
			"username": schema.StringAttribute{
				Description: "Defines the ruleset expression to use in matching the username in a request.",
				Computed:    true,
			},
		},
	}
}

func (d *LeakedCredentialCheckRuleDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *LeakedCredentialCheckRuleDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
