// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dlp_custom_profile

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = &DLPCustomProfileDataSource{}
var _ datasource.DataSourceWithValidateConfig = &DLPCustomProfileDataSource{}

func (r DLPCustomProfileDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"profile_id": schema.StringAttribute{
				Description: "Unique identifier for a DLP profile",
				Required:    true,
			},
			"created_at": schema.StringAttribute{
				Optional:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"description": schema.StringAttribute{
				Description: "The description of the profile.",
				Optional:    true,
			},
			"id": schema.StringAttribute{
				Description: "Unique identifier for a DLP profile",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the profile.",
				Optional:    true,
			},
			"ocr_enabled": schema.BoolAttribute{
				Description: "If true, scan images via OCR to determine if any text present matches filters.",
				Optional:    true,
			},
			"type": schema.StringAttribute{
				Description: "The type of the profile.",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("custom"),
				},
			},
			"updated_at": schema.StringAttribute{
				Optional:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"context_awareness": schema.SingleNestedAttribute{
				Description: "Scan the context of predefined entries to only return matches surrounded by keywords.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"enabled": schema.BoolAttribute{
						Description: "If true, scan the context of predefined entries to only return matches surrounded by keywords.",
						Computed:    true,
					},
					"skip": schema.SingleNestedAttribute{
						Description: "Content types to exclude from context analysis and return all matches.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[DLPCustomProfileContextAwarenessSkipDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"files": schema.BoolAttribute{
								Description: "If the content type is a file, skip context analysis and return all matches.",
								Computed:    true,
							},
						},
					},
				},
			},
			"entries": schema.ListNestedAttribute{
				Description: "The entries for this profile.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Unique identifier for a DLP entry",
							Computed:    true,
						},
						"created_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"enabled": schema.BoolAttribute{
							Description: "Whether the entry is enabled or not.",
							Computed:    true,
							Optional:    true,
						},
						"name": schema.StringAttribute{
							Description: "The name of the entry.",
							Computed:    true,
							Optional:    true,
						},
						"pattern": schema.SingleNestedAttribute{
							Description: "A pattern that matches an entry",
							Computed:    true,
							Optional:    true,
							Attributes: map[string]schema.Attribute{
								"regex": schema.StringAttribute{
									Description: "The regex pattern.",
									Computed:    true,
								},
								"validation": schema.StringAttribute{
									Description: "Validation algorithm for the pattern. This algorithm will get run on potential matches, and if it returns false, the entry will not be matched.",
									Computed:    true,
									Optional:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("luhn"),
									},
								},
							},
						},
						"profile_id": schema.StringAttribute{
							Description: "Unique identifier for a DLP profile",
							Computed:    true,
						},
						"updated_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
					},
				},
			},
			"allowed_match_count": schema.Float64Attribute{
				Description: "Related DLP policies will trigger when the match count exceeds the number set.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.Between(0, 1000),
				},
			},
		},
	}
}

func (r *DLPCustomProfileDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *DLPCustomProfileDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
