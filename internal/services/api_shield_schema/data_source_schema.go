// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield_schema

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = &APIShieldSchemaDataSource{}

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"schema_id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"omit_source": schema.BoolAttribute{
				Description: "Omit the source-files of schemas and only retrieve their meta-data.",
				Optional:    true,
			},
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
			"filter": schema.SingleNestedAttribute{
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
					"validation_enabled": schema.BoolAttribute{
						Description: "Flag whether schema is enabled for validation.",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (d *APIShieldSchemaDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *APIShieldSchemaDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.RequiredTogether(path.MatchRoot("schema_id"), path.MatchRoot("zone_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("schema_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("zone_id")),
	}
}
