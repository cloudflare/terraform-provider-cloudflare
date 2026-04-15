// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pipeline_stream

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*PipelineStreamDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Specifies the public ID of the stream.",
				Computed:    true,
			},
			"stream_id": schema.StringAttribute{
				Description: "Specifies the public ID of the stream.",
				Optional:    true,
			},
			"account_id": schema.StringAttribute{
				Description: "Specifies the public ID of the account.",
				Required:    true,
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
			"name": schema.StringAttribute{
				Description: "Indicates the name of the Stream.",
				Computed:    true,
			},
			"version": schema.Int64Attribute{
				Description: "Indicates the current version of this stream.",
				Computed:    true,
			},
			"format": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[PipelineStreamFormatDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"type": schema.StringAttribute{
						Description: `Available values: "json", "parquet".`,
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("json", "parquet"),
						},
					},
					"decimal_encoding": schema.StringAttribute{
						Description: `Available values: "number", "string", "bytes".`,
						Computed:    true,
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
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("rfc3339", "unix_millis"),
						},
					},
					"unstructured": schema.BoolAttribute{
						Computed: true,
					},
					"compression": schema.StringAttribute{
						Description: `Available values: "uncompressed", "snappy", "gzip", "zstd", "lz4".`,
						Computed:    true,
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
						Computed: true,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},
				},
			},
			"http": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[PipelineStreamHTTPDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"authentication": schema.BoolAttribute{
						Description: "Indicates that authentication is required for the HTTP endpoint.",
						Computed:    true,
					},
					"enabled": schema.BoolAttribute{
						Description: "Indicates that the HTTP endpoint is enabled.",
						Computed:    true,
					},
					"cors": schema.SingleNestedAttribute{
						Description: "Specifies the CORS options for the HTTP endpoint.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[PipelineStreamHTTPCORSDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"origins": schema.ListAttribute{
								Computed:    true,
								CustomType:  customfield.NewListType[types.String](ctx),
								ElementType: types.StringType,
							},
						},
					},
				},
			},
			"schema": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[PipelineStreamSchemaDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"fields": schema.ListNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectListType[PipelineStreamSchemaFieldsDataSourceModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"type": schema.StringAttribute{
									Description: `Available values: "int32", "int64", "float32", "float64", "bool", "string", "binary", "timestamp", "json".`,
									Computed:    true,
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
									Computed: true,
								},
								"name": schema.StringAttribute{
									Computed: true,
								},
								"required": schema.BoolAttribute{
									Computed: true,
								},
								"sql_name": schema.StringAttribute{
									Computed: true,
								},
								"unit": schema.StringAttribute{
									Description: `Available values: "second", "millisecond", "microsecond", "nanosecond".`,
									Computed:    true,
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
						Computed:   true,
						CustomType: customfield.NewNestedObjectType[PipelineStreamSchemaFormatDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Description: `Available values: "json", "parquet".`,
								Computed:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("json", "parquet"),
								},
							},
							"decimal_encoding": schema.StringAttribute{
								Description: `Available values: "number", "string", "bytes".`,
								Computed:    true,
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
								Computed:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("rfc3339", "unix_millis"),
								},
							},
							"unstructured": schema.BoolAttribute{
								Computed: true,
							},
							"compression": schema.StringAttribute{
								Description: `Available values: "uncompressed", "snappy", "gzip", "zstd", "lz4".`,
								Computed:    true,
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
								Computed: true,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
								},
							},
						},
					},
					"inferred": schema.BoolAttribute{
						Computed: true,
					},
				},
			},
			"worker_binding": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[PipelineStreamWorkerBindingDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"enabled": schema.BoolAttribute{
						Description: "Indicates that the worker binding is enabled.",
						Computed:    true,
					},
				},
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"pipeline_id": schema.StringAttribute{
						Description: "Specifies the public ID of the pipeline.",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (d *PipelineStreamDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *PipelineStreamDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("stream_id"), path.MatchRoot("filter")),
	}
}
