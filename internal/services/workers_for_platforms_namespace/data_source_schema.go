// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms_namespace

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = &WorkersForPlatformsNamespaceDataSource{}
var _ datasource.DataSourceWithValidateConfig = &WorkersForPlatformsNamespaceDataSource{}

func (r WorkersForPlatformsNamespaceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"dispatch_namespace": schema.StringAttribute{
				Description: "Name of the Workers for Platforms dispatch namespace.",
				Optional:    true,
			},
			"created_by": schema.StringAttribute{
				Description: "Identifier",
				Computed:    true,
				Optional:    true,
			},
			"created_on": schema.StringAttribute{
				Description: "When the script was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"modified_by": schema.StringAttribute{
				Description: "Identifier",
				Computed:    true,
				Optional:    true,
			},
			"modified_on": schema.StringAttribute{
				Description: "When the script was last modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"namespace_id": schema.StringAttribute{
				Description: "API Resource UUID tag.",
				Computed:    true,
				Optional:    true,
			},
			"namespace_name": schema.StringAttribute{
				Description: "Name of the Workers for Platforms dispatch namespace.",
				Computed:    true,
				Optional:    true,
			},
			"script_count": schema.Int64Attribute{
				Description: "The current number of scripts in this Dispatch Namespace",
				Computed:    true,
				Optional:    true,
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Description: "Identifier",
						Required:    true,
					},
				},
			},
		},
	}
}

func (r *WorkersForPlatformsNamespaceDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *WorkersForPlatformsNamespaceDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
