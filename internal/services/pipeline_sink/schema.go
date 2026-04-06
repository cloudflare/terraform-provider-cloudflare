// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pipeline_sink

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*PipelineSinkResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Version: 500,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Indicates a unique identifier for this sink.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"account_id": schema.StringAttribute{
				Description:   "Specifies the public ID of the account.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description:   "Defines the name of the Sink.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"type": schema.StringAttribute{
				Description: "Specifies the type of sink.\nAvailable values: \"r2\", \"r2_data_catalog\".",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("r2", "r2_data_catalog"),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"config": schema.SingleNestedAttribute{
				Description: "Defines the configuration of the R2 Sink.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Description: "Cloudflare Account ID for the bucket",
						Required:    true,
					},
					"bucket": schema.StringAttribute{
						Description: "R2 Bucket to write to",
						Required:    true,
					},
					"credentials": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"access_key_id": schema.StringAttribute{
								Description: "Cloudflare Account ID for the bucket",
								Required:    true,
							},
							"secret_access_key": schema.StringAttribute{
								Description: "Cloudflare Account ID for the bucket",
								Required:    true,
								Sensitive:   true,
							},
						},
					},
					"file_naming": schema.SingleNestedAttribute{
						Description: "Controls filename prefix/suffix and strategy.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"prefix": schema.StringAttribute{
								Description: "The prefix to use in file name. i.e prefix-<uuid>.parquet",
								Optional:    true,
							},
							"strategy": schema.StringAttribute{
								Description: "Filename generation strategy.\nAvailable values: \"serial\", \"uuid\", \"uuid_v7\", \"ulid\".",
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive(
										"serial",
										"uuid",
										"uuid_v7",
										"ulid",
									),
								},
							},
							"suffix": schema.StringAttribute{
								Description: "This will overwrite the default file suffix. i.e .parquet, use with caution",
								Optional:    true,
							},
						},
					},
					"jurisdiction": schema.StringAttribute{
						Description: "Jurisdiction this bucket is hosted in",
						Optional:    true,
					},
					"partitioning": schema.SingleNestedAttribute{
						Description: "Data-layout partitioning for sinks.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"time_pattern": schema.StringAttribute{
								Description: "The pattern of the date string",
								Optional:    true,
							},
						},
					},
					"path": schema.StringAttribute{
						Description: "Subpath within the bucket to write to",
						Optional:    true,
					},
					"rolling_policy": schema.SingleNestedAttribute{
						Description: "Rolling policy for file sinks (when & why to close a file and open a new one).",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"file_size_bytes": schema.Int64Attribute{
								Description: "Files will be rolled after reaching this number of bytes",
								Optional:    true,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
								},
							},
							"inactivity_seconds": schema.Int64Attribute{
								Description: "Number of seconds of inactivity to wait before rolling over to a new file",
								Optional:    true,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
								},
							},
							"interval_seconds": schema.Int64Attribute{
								Description: "Number of seconds to wait before rolling over to a new file",
								Optional:    true,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
								},
							},
						},
					},
					"token": schema.StringAttribute{
						Description: "Authentication token",
						Optional:    true,
						Sensitive:   true,
					},
					"table_name": schema.StringAttribute{
						Description: "Table name",
						Optional:    true,
					},
					"namespace": schema.StringAttribute{
						Description: "Table namespace",
						Optional:    true,
					},
				},
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
			},
			"format": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"type": schema.StringAttribute{
						Description: `Available values: "json", "parquet".`,
						Required:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("json", "parquet"),
						},
					},
					"decimal_encoding": schema.StringAttribute{
						Description: `Available values: "number", "string", "bytes".`,
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"number",
								"string",
								"bytes",
							),
						},
					},
					"timestamp_format": schema.StringAttribute{
						Description: `Available values: "rfc3339", "unix_millis".`,
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("rfc3339", "unix_millis"),
						},
					},
					"unstructured": schema.BoolAttribute{
						Optional: true,
					},
					"compression": schema.StringAttribute{
						Description: `Available values: "uncompressed", "snappy", "gzip", "zstd", "lz4".`,
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"uncompressed",
								"snappy",
								"gzip",
								"zstd",
								"lz4",
							),
						},
					},
					"row_group_bytes": schema.Int64Attribute{
						Optional: true,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},
				},
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
			},
			"schema": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"fields": schema.ListNestedAttribute{
						Optional: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"type": schema.StringAttribute{
									Description: `Available values: "int32", "int64", "float32", "float64", "bool", "string", "binary", "timestamp", "json".`,
									Required:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"int32",
											"int64",
											"float32",
											"float64",
											"bool",
											"string",
											"binary",
											"timestamp",
											"json",
										),
									},
								},
								"metadata_key": schema.StringAttribute{
									Optional: true,
								},
								"name": schema.StringAttribute{
									Optional: true,
								},
								"required": schema.BoolAttribute{
									Optional: true,
								},
								"sql_name": schema.StringAttribute{
									Optional: true,
								},
								"unit": schema.StringAttribute{
									Description: `Available values: "second", "millisecond", "microsecond", "nanosecond".`,
									Optional:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"second",
											"millisecond",
											"microsecond",
											"nanosecond",
										),
									},
								},
							},
						},
					},
					"format": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Description: `Available values: "json", "parquet".`,
								Required:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("json", "parquet"),
								},
							},
							"decimal_encoding": schema.StringAttribute{
								Description: `Available values: "number", "string", "bytes".`,
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive(
										"number",
										"string",
										"bytes",
									),
								},
							},
							"timestamp_format": schema.StringAttribute{
								Description: `Available values: "rfc3339", "unix_millis".`,
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("rfc3339", "unix_millis"),
								},
							},
							"unstructured": schema.BoolAttribute{
								Optional: true,
							},
							"compression": schema.StringAttribute{
								Description: `Available values: "uncompressed", "snappy", "gzip", "zstd", "lz4".`,
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive(
										"uncompressed",
										"snappy",
										"gzip",
										"zstd",
										"lz4",
									),
								},
							},
							"row_group_bytes": schema.Int64Attribute{
								Optional: true,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
								},
							},
						},
					},
					"inferred": schema.BoolAttribute{
						Optional: true,
					},
				},
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"modified_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
		},
	}
}

func (r *PipelineSinkResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *PipelineSinkResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
