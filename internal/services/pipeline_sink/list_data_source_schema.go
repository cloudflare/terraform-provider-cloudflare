// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pipeline_sink

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

var _ datasource.DataSourceWithConfigValidators = (*PipelineSinksDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Specifies the public ID of the account.",
				Required:    true,
			},
			"pipeline_id": schema.StringAttribute{
				Optional: true,
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
				CustomType:  customfield.NewNestedObjectListType[PipelineSinksResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Indicates a unique identifier for this sink.",
							Computed:    true,
						},
						"created_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"modified_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"name": schema.StringAttribute{
							Description: "Defines the name of the Sink.",
							Computed:    true,
						},
						"type": schema.StringAttribute{
							Description: "Specifies the type of sink.\nAvailable values: \"r2\", \"r2_data_catalog\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("r2", "r2_data_catalog"),
							},
						},
						"config": schema.SingleNestedAttribute{
							Description: "Defines the configuration of the R2 Sink.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[PipelineSinksConfigDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"account_id": schema.StringAttribute{
									Description: "Cloudflare Account ID for the bucket",
									Computed:    true,
								},
								"bucket": schema.StringAttribute{
									Description: "R2 Bucket to write to",
									Computed:    true,
								},
								"file_naming": schema.SingleNestedAttribute{
									Description: "Controls filename prefix/suffix and strategy.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[PipelineSinksConfigFileNamingDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"prefix": schema.StringAttribute{
											Description: "The prefix to use in file name. i.e prefix-<uuid>.parquet",
											Computed:    true,
										},
										"strategy": schema.StringAttribute{
											Description: "Filename generation strategy.\nAvailable values: \"serial\", \"uuid\", \"uuid_v7\", \"ulid\".",
											Computed:    true,
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
											Computed:    true,
										},
									},
								},
								"jurisdiction": schema.StringAttribute{
									Description: "Jurisdiction this bucket is hosted in",
									Computed:    true,
								},
								"partitioning": schema.SingleNestedAttribute{
									Description: "Data-layout partitioning for sinks.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[PipelineSinksConfigPartitioningDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"time_pattern": schema.StringAttribute{
											Description: "The pattern of the date string",
											Computed:    true,
										},
									},
								},
								"path": schema.StringAttribute{
									Description: "Subpath within the bucket to write to",
									Computed:    true,
								},
								"rolling_policy": schema.SingleNestedAttribute{
									Description: "Rolling policy for file sinks (when & why to close a file and open a new one).",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[PipelineSinksConfigRollingPolicyDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"file_size_bytes": schema.Int64Attribute{
											Description: "Files will be rolled after reaching this number of bytes",
											Computed:    true,
											Validators: []validator.Int64{
												int64validator.AtLeast(0),
											},
										},
										"inactivity_seconds": schema.Int64Attribute{
											Description: "Number of seconds of inactivity to wait before rolling over to a new file",
											Computed:    true,
											Validators: []validator.Int64{
												int64validator.AtLeast(1),
											},
										},
										"interval_seconds": schema.Int64Attribute{
											Description: "Number of seconds to wait before rolling over to a new file",
											Computed:    true,
											Validators: []validator.Int64{
												int64validator.AtLeast(1),
											},
										},
									},
								},
								"table_name": schema.StringAttribute{
									Description: "Table name",
									Computed:    true,
								},
								"namespace": schema.StringAttribute{
									Description: "Table namespace",
									Computed:    true,
								},
							},
						},
						"format": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[PipelineSinksFormatDataSourceModel](ctx),
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
						"schema": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[PipelineSinksSchemaDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"fields": schema.ListNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectListType[PipelineSinksSchemaFieldsDataSourceModel](ctx),
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
									CustomType: customfield.NewNestedObjectType[PipelineSinksSchemaFormatDataSourceModel](ctx),
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
					},
				},
			},
		},
	}
}

func (d *PipelineSinksDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *PipelineSinksDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
