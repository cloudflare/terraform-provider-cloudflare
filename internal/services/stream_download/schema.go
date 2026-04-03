// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream_download

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*StreamDownloadResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description:   "Identifier.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"identifier": schema.StringAttribute{
				Description:   "A Cloudflare-generated unique identifier for a media item.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"percent_complete": schema.Float64Attribute{
				Description: "Indicates the progress as a percentage between 0 and 100.",
				Computed:    true,
				Validators: []validator.Float64{
					float64validator.Between(0, 100),
				},
			},
			"status": schema.StringAttribute{
				Description: "The status of a generated download.\nAvailable values: \"ready\", \"inprogress\", \"error\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"ready",
						"inprogress",
						"error",
					),
				},
			},
			"url": schema.StringAttribute{
				Description: "The URL to access the generated download.",
				Computed:    true,
			},
			"audio": schema.SingleNestedAttribute{
				Description: "The audio-only download. Only present if this download type has been created.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[StreamDownloadAudioModel](ctx),
				Attributes: map[string]schema.Attribute{
					"percent_complete": schema.Float64Attribute{
						Description: "Indicates the progress as a percentage between 0 and 100.",
						Computed:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 100),
						},
					},
					"status": schema.StringAttribute{
						Description: "The status of a generated download.\nAvailable values: \"ready\", \"inprogress\", \"error\".",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"ready",
								"inprogress",
								"error",
							),
						},
					},
					"url": schema.StringAttribute{
						Description: "The URL to access the generated download.",
						Computed:    true,
					},
				},
			},
			"default": schema.SingleNestedAttribute{
				Description: "The default video download. Only present if this download type has been created.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[StreamDownloadDefaultModel](ctx),
				Attributes: map[string]schema.Attribute{
					"percent_complete": schema.Float64Attribute{
						Description: "Indicates the progress as a percentage between 0 and 100.",
						Computed:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 100),
						},
					},
					"status": schema.StringAttribute{
						Description: "The status of a generated download.\nAvailable values: \"ready\", \"inprogress\", \"error\".",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"ready",
								"inprogress",
								"error",
							),
						},
					},
					"url": schema.StringAttribute{
						Description: "The URL to access the generated download.",
						Computed:    true,
					},
				},
			},
		},
	}
}

func (r *StreamDownloadResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *StreamDownloadResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
