// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield_schema

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = &APIShieldSchemasDataSource{}
var _ datasource.DataSourceWithValidateConfig = &APIShieldSchemasDataSource{}

func (r APIShieldSchemasDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"omit_source": schema.BoolAttribute{
				Description: "Omit the source-files of schemas and only retrieve their meta-data.",
				Computed:    true,
				Optional:    true,
			},
			"page": schema.Int64Attribute{
				Description: "Page number of paginated results.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"per_page": schema.Int64Attribute{
				Description: "Maximum number of results per page.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.Between(5, 50),
				},
			},
			"validation_enabled": schema.BoolAttribute{
				Description: "Flag whether schema is enabled for validation.",
				Optional:    true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
			},
			"result": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"created_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"kind": schema.StringAttribute{
							Description: "Kind of schema",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("openapi_v3"),
							},
						},
						"name": schema.StringAttribute{
							Description: "Name of the schema",
							Computed:    true,
						},
						"schema_id": schema.StringAttribute{
							Description: "UUID",
							Computed:    true,
						},
						"source": schema.StringAttribute{
							Description: "Source of the schema",
							Computed:    true,
							Optional:    true,
						},
						"validation_enabled": schema.BoolAttribute{
							Description: "Flag whether schema is enabled for validation.",
							Computed:    true,
							Optional:    true,
						},
					},
				},
			},
		},
	}
}

func (r *APIShieldSchemasDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *APIShieldSchemasDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
