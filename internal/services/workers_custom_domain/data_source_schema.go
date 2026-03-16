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
			"id": schema.StringAttribute{
				Description: "ID of the domain.",
				Computed:    true,
			},
			"domain_id": schema.StringAttribute{
				Description: "ID of the domain.",
				Optional:    true,
			},
			"account_id": schema.StringAttribute{
				Description: "Identifier.",
				Required:    true,
			},
			"cert_id": schema.StringAttribute{
				Description: "ID of the TLS certificate issued for the domain.",
				Computed:    true,
			},
			"environment": schema.StringAttribute{
				Description:        "Worker environment associated with the domain.",
				Computed:           true,
				DeprecationMessage: "This attribute is deprecated.",
			},
			"hostname": schema.StringAttribute{
				Description: "Hostname of the domain. Can be either the zone apex or a subdomain of the zone. Requests to this hostname will be routed to the configured Worker.",
				Computed:    true,
			},
			"service": schema.StringAttribute{
				Description: "Name of the Worker associated with the domain. Requests to the configured hostname will be routed to this Worker.",
				Computed:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "ID of the zone containing the domain hostname.",
				Computed:    true,
			},
			"zone_name": schema.StringAttribute{
				Description: "Name of the zone containing the domain hostname.",
				Computed:    true,
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"environment": schema.StringAttribute{
						Description: "Worker environment associated with the domain.",
						Optional:    true,
					},
					"hostname": schema.StringAttribute{
						Description: "Hostname of the domain.",
						Optional:    true,
					},
					"service": schema.StringAttribute{
						Description: "Name of the Worker associated with the domain.",
						Optional:    true,
					},
					"zone_id": schema.StringAttribute{
						Description: "ID of the zone containing the domain hostname.",
						Optional:    true,
					},
					"zone_name": schema.StringAttribute{
						Description: "Name of the zone containing the domain hostname.",
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
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("domain_id"), path.MatchRoot("filter")),
	}
}
