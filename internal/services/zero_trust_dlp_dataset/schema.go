// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_dataset

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*ZeroTrustDLPDatasetResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"dataset_id": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"encoding_version": schema.Int64Attribute{
				Description: "Dataset encoding version\n\nNon-secret custom word lists with no header are always version 1.\nSecret EDM lists with no header are version 1.\nMulticolumn CSV with headers are version 2.\nOmitting this field provides the default value 0, which is interpreted\nthe same as 1.",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"secret": schema.BoolAttribute{
				Description:   "Generate a secret dataset.\n\nIf true, the response will include a secret to use with the EDM encoder.\nIf false, the response has no secret and the dataset is uploaded in plaintext.",
				Optional:      true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"case_sensitive": schema.BoolAttribute{
				Description: "Only applies to custom word lists.\nDetermines if the words should be matched in a case-sensitive manner\nCannot be set to false if `secret` is true or undefined",
				Optional:    true,
			},
			"description": schema.StringAttribute{
				Description: "The description of the dataset.",
				Optional:    true,
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"max_cells": schema.Int64Attribute{
				Computed: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"num_cells": schema.Int64Attribute{
				Computed: true,
			},
			"status": schema.StringAttribute{
				Description: `Available values: "empty", "uploading", "processing", "failed", "complete".`,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"empty",
						"uploading",
						"processing",
						"failed",
						"complete",
					),
				},
			},
			"updated_at": schema.StringAttribute{
				Description: "When the dataset was last updated.\n\nThis includes name or description changes as well as uploads.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"version": schema.Int64Attribute{
				Description: "The version to use when uploading the dataset.",
				Computed:    true,
			},
			"columns": schema.ListNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[ZeroTrustDLPDatasetColumnsModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"entry_id": schema.StringAttribute{
							Computed: true,
						},
						"header_name": schema.StringAttribute{
							Computed: true,
						},
						"num_cells": schema.Int64Attribute{
							Computed: true,
						},
						"upload_status": schema.StringAttribute{
							Description: `Available values: "empty", "uploading", "processing", "failed", "complete".`,
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"empty",
									"uploading",
									"processing",
									"failed",
									"complete",
								),
							},
						},
					},
				},
			},
			"dataset": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[ZeroTrustDLPDatasetDatasetModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
					"columns": schema.ListNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectListType[ZeroTrustDLPDatasetDatasetColumnsModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"entry_id": schema.StringAttribute{
									Computed: true,
								},
								"header_name": schema.StringAttribute{
									Computed: true,
								},
								"num_cells": schema.Int64Attribute{
									Computed: true,
								},
								"upload_status": schema.StringAttribute{
									Description: `Available values: "empty", "uploading", "processing", "failed", "complete".`,
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"empty",
											"uploading",
											"processing",
											"failed",
											"complete",
										),
									},
								},
							},
						},
					},
					"created_at": schema.StringAttribute{
						Computed:   true,
						CustomType: timetypes.RFC3339Type{},
					},
					"encoding_version": schema.Int64Attribute{
						Computed: true,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},
					"name": schema.StringAttribute{
						Computed: true,
					},
					"num_cells": schema.Int64Attribute{
						Computed: true,
					},
					"secret": schema.BoolAttribute{
						Computed: true,
					},
					"status": schema.StringAttribute{
						Description: `Available values: "empty", "uploading", "processing", "failed", "complete".`,
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"empty",
								"uploading",
								"processing",
								"failed",
								"complete",
							),
						},
					},
					"updated_at": schema.StringAttribute{
						Description: "When the dataset was last updated.\n\nThis includes name or description changes as well as uploads.",
						Computed:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
					"uploads": schema.ListNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectListType[ZeroTrustDLPDatasetDatasetUploadsModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"num_cells": schema.Int64Attribute{
									Computed: true,
								},
								"status": schema.StringAttribute{
									Description: `Available values: "empty", "uploading", "processing", "failed", "complete".`,
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"empty",
											"uploading",
											"processing",
											"failed",
											"complete",
										),
									},
								},
								"version": schema.Int64Attribute{
									Computed: true,
								},
							},
						},
					},
					"case_sensitive": schema.BoolAttribute{
						Computed: true,
					},
					"description": schema.StringAttribute{
						Description: "The description of the dataset.",
						Computed:    true,
					},
				},
			},
			"uploads": schema.ListNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[ZeroTrustDLPDatasetUploadsModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"num_cells": schema.Int64Attribute{
							Computed: true,
						},
						"status": schema.StringAttribute{
							Description: `Available values: "empty", "uploading", "processing", "failed", "complete".`,
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"empty",
									"uploading",
									"processing",
									"failed",
									"complete",
								),
							},
						},
						"version": schema.Int64Attribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func (r *ZeroTrustDLPDatasetResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ZeroTrustDLPDatasetResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
