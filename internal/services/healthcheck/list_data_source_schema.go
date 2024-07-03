// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package healthcheck

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = &HealthchecksDataSource{}
var _ datasource.DataSourceWithValidateConfig = &HealthchecksDataSource{}

func (r HealthchecksDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"page": schema.StringAttribute{
				Description: "Page number of paginated results.",
				Computed:    true,
				Optional:    true,
			},
			"per_page": schema.StringAttribute{
				Description: "Maximum number of results per page. Must be a multiple of 5.",
				Computed:    true,
				Optional:    true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
			},
			"items": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Identifier",
							Computed:    true,
						},
						"address": schema.StringAttribute{
							Description: "The hostname or IP address of the origin server to run health checks on.",
							Computed:    true,
						},
						"check_regions": schema.ListAttribute{
							Description: "A list of regions from which to run health checks. Null means Cloudflare will pick a default region.",
							Computed:    true,
							ElementType: types.StringType,
						},
						"consecutive_fails": schema.Int64Attribute{
							Description: "The number of consecutive fails required from a health check before changing the health to unhealthy.",
							Computed:    true,
						},
						"consecutive_successes": schema.Int64Attribute{
							Description: "The number of consecutive successes required from a health check before changing the health to healthy.",
							Computed:    true,
						},
						"created_on": schema.StringAttribute{
							Computed: true,
						},
						"description": schema.StringAttribute{
							Description: "A human-readable description of the health check.",
							Computed:    true,
						},
						"failure_reason": schema.StringAttribute{
							Description: "The current failure reason if status is unhealthy.",
							Computed:    true,
						},
						"interval": schema.Int64Attribute{
							Description: "The interval between each health check. Shorter intervals may give quicker notifications if the origin status changes, but will increase load on the origin as we check from multiple locations.",
							Computed:    true,
						},
						"modified_on": schema.StringAttribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Description: "A short name to identify the health check. Only alphanumeric characters, hyphens and underscores are allowed.",
							Computed:    true,
						},
						"retries": schema.Int64Attribute{
							Description: "The number of retries to attempt in case of a timeout before marking the origin as unhealthy. Retries are attempted immediately.",
							Computed:    true,
						},
						"status": schema.StringAttribute{
							Description: "The current status of the origin server according to the health check.",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("unknown", "healthy", "unhealthy", "suspended"),
							},
						},
						"suspended": schema.BoolAttribute{
							Description: "If suspended, no health checks are sent to the origin.",
							Computed:    true,
						},
						"timeout": schema.Int64Attribute{
							Description: "The timeout (in seconds) before marking the health check as failed.",
							Computed:    true,
						},
						"type": schema.StringAttribute{
							Description: "The protocol to use for the health check. Currently supported protocols are 'HTTP', 'HTTPS' and 'TCP'.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (r *HealthchecksDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *HealthchecksDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
