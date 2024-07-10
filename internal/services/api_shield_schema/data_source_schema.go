// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield_schema

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = &APIShieldSchemaDataSource{}
var _ datasource.DataSourceWithValidateConfig = &APIShieldSchemaDataSource{}

func (r APIShieldSchemaDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"schema_id": schema.StringAttribute{
				Description: "UUID identifier",
				Computed:    true,
				Optional:    true,
			},
			"omit_source": schema.BoolAttribute{
				Description: "Omit the source-files of schemas and only retrieve their meta-data.",
				Optional:    true,
			},
			"created_at": schema.StringAttribute{
				Computed: true,
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
			"find_one_by": schema.SingleNestedAttribute{
				Optional: true,
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
					"page": schema.StringAttribute{
						Description: "Page number of paginated results.",
						Computed:    true,
						Optional:    true,
					},
					"per_page": schema.StringAttribute{
						Description: "Maximum number of results per page.",
						Computed:    true,
						Optional:    true,
					},
					"validation_enabled": schema.BoolAttribute{
						Description: "Flag whether schema is enabled for validation.",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (r *APIShieldSchemaDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *APIShieldSchemaDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
