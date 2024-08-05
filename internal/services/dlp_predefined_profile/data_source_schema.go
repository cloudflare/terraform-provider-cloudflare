// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dlp_predefined_profile

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = &DLPPredefinedProfileDataSource{}

func (d *DLPPredefinedProfileDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
					stringvalidator.OneOfCaseInsensitive("predefined"),
				},
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
						CustomType:  customfield.NewNestedObjectType[DLPPredefinedProfileContextAwarenessSkipDataSourceModel](ctx),
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
						"profile_id": schema.StringAttribute{
							Description: "Unique identifier for a DLP profile",
							Computed:    true,
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

func (d *DLPPredefinedProfileDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
