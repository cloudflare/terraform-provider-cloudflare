// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_entry

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustDLPEntryDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"entry_id": schema.StringAttribute{
				Optional: true,
			},
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"enabled": schema.BoolAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"profile_id": schema.StringAttribute{
				Computed: true,
			},
			"secret": schema.BoolAttribute{
				Computed: true,
			},
			"type": schema.StringAttribute{
				Description: "available values: \"custom\"",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"custom",
						"predefined",
						"integration",
						"exact_data",
						"word_list",
					),
				},
			},
			"updated_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"confidence": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[ZeroTrustDLPEntryConfidenceDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"ai_context_available": schema.BoolAttribute{
						Description: "Indicates whether this entry has AI remote service validation",
						Computed:    true,
					},
					"available": schema.BoolAttribute{
						Description: "Indicates whether this entry has any form of validation that is not an AI remote service",
						Computed:    true,
					},
				},
			},
			"pattern": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[ZeroTrustDLPEntryPatternDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"regex": schema.StringAttribute{
						Computed: true,
					},
					"validation": schema.StringAttribute{
						Description: "available values: \"luhn\"",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("luhn"),
						},
					},
				},
			},
			"word_list": schema.StringAttribute{
				Computed:   true,
				CustomType: jsontypes.NormalizedType{},
			},
		},
	}
}

func (d *ZeroTrustDLPEntryDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ZeroTrustDLPEntryDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
