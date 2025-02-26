// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_dataset

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustDLPDatasetsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"result": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustDLPDatasetsResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed: true,
						},
						"columns": schema.ListNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectListType[ZeroTrustDLPDatasetsColumnsDataSourceModel](ctx),
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
										Description: "available values: \"empty\", \"uploading\", \"processing\", \"failed\", \"complete\"",
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
							Description: "available values: \"empty\", \"uploading\", \"processing\", \"failed\", \"complete\"",
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
							CustomType: customfield.NewNestedObjectListType[ZeroTrustDLPDatasetsUploadsDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"num_cells": schema.Int64Attribute{
										Computed: true,
									},
									"status": schema.StringAttribute{
										Description: "available values: \"empty\", \"uploading\", \"processing\", \"failed\", \"complete\"",
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
						"description": schema.StringAttribute{
							Description: "The description of the dataset",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *ZeroTrustDLPDatasetsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *ZeroTrustDLPDatasetsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
