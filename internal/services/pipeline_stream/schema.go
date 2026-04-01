// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pipeline_stream

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*PipelineStreamResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Version: 500,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Indicates a unique identifier for this stream.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description:   "Specifies the public ID of the account.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description:   "Specifies the name of the Stream.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
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
			"http": schema.SingleNestedAttribute{
				Computed:   true,
				Optional:   true,
				CustomType: customfield.NewNestedObjectType[PipelineStreamHTTPModel](ctx),
				Attributes: map[string]schema.Attribute{
					"authentication": schema.BoolAttribute{
						Description: "Indicates that authentication is required for the HTTP endpoint.",
						Required:    true,
					},
					"enabled": schema.BoolAttribute{
						Description: "Indicates that the HTTP endpoint is enabled.",
						Required:    true,
					},
					"cors": schema.SingleNestedAttribute{
						Description: "Specifies the CORS options for the HTTP endpoint.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"origins": schema.ListAttribute{
								Optional:    true,
								ElementType: types.StringType,
							},
						},
					},
				},
			},
			"worker_binding": schema.SingleNestedAttribute{
				Computed:   true,
				Optional:   true,
				CustomType: customfield.NewNestedObjectType[PipelineStreamWorkerBindingModel](ctx),
				Attributes: map[string]schema.Attribute{
					"enabled": schema.BoolAttribute{
						Description: "Indicates that the worker binding is enabled.",
						Required:    true,
					},
				},
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"endpoint": schema.StringAttribute{
				Description: "Indicates the endpoint URL of this stream.",
				Computed:    true,
			},
			"modified_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"version": schema.Int64Attribute{
				Description: "Indicates the current version of this stream.",
				Computed:    true,
			},
		},
	}
}

func (r *PipelineStreamResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *PipelineStreamResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
