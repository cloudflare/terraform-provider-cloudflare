// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package worker_domain

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = &WorkerDomainDataSource{}
var _ datasource.DataSourceWithValidateConfig = &WorkerDomainDataSource{}

func (r WorkerDomainDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Optional: true,
			},
			"domain_id": schema.StringAttribute{
				Description: "Identifer of the Worker Domain.",
				Optional:    true,
			},
			"id": schema.StringAttribute{
				Description: "Identifer of the Worker Domain.",
				Optional:    true,
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
			"find_one_by": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Required: true,
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

func (r *WorkerDomainDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *WorkerDomainDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
