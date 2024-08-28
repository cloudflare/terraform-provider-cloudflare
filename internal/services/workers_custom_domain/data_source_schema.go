// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_custom_domain

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
)

var _ datasource.DataSourceWithConfigValidators = (*WorkersCustomDomainDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifer of the account.",
				Optional:    true,
			},
			"domain_id": schema.StringAttribute{
				Description: "Identifer of the Worker Domain.",
				Optional:    true,
			},
			"environment": schema.StringAttribute{
				Description: "Worker environment associated with the zone and hostname.",
				Computed:    true,
			},
			"hostname": schema.StringAttribute{
				Description: "Hostname of the Worker Domain.",
				Computed:    true,
			},
			"id": schema.StringAttribute{
				Description: "Identifer of the Worker Domain.",
				Computed:    true,
			},
			"service": schema.StringAttribute{
				Description: "Worker service associated with the zone and hostname.",
				Computed:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "Identifier of the zone.",
				Computed:    true,
			},
			"zone_name": schema.StringAttribute{
				Description: "Name of the zone.",
				Computed:    true,
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Description: "Identifer of the account.",
						Required:    true,
					},
					"environment": schema.StringAttribute{
						Description: "Worker environment associated with the zone and hostname.",
						Optional:    true,
					},
					"hostname": schema.StringAttribute{
						Description: "Hostname of the Worker Domain.",
						Optional:    true,
					},
					"service": schema.StringAttribute{
						Description: "Worker service associated with the zone and hostname.",
						Optional:    true,
					},
					"zone_id": schema.StringAttribute{
						Description: "Identifier of the zone.",
						Optional:    true,
					},
					"zone_name": schema.StringAttribute{
						Description: "Name of the zone.",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (d *WorkersCustomDomainDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *WorkersCustomDomainDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.RequiredTogether(path.MatchRoot("account_id"), path.MatchRoot("domain_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("account_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("domain_id")),
	}
}
