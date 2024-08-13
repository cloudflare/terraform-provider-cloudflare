// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logpush_job

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = &LogpushJobDataSource{}

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.",
				Optional:    true,
			},
			"job_id": schema.Int64Attribute{
				Description: "Unique id of the job.",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"zone_id": schema.StringAttribute{
				Description: "The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.",
				Optional:    true,
			},
			"frequency": schema.StringAttribute{
				Description: "This field is deprecated. Please use `max_upload_*` parameters instead. The frequency at which Cloudflare sends batches of logs to your destination. Setting frequency to high sends your logs in larger quantities of smaller files. Setting frequency to low sends logs in smaller quantities of larger files.",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("high", "low"),
				},
			},
			"max_upload_interval_seconds": schema.Int64Attribute{
				Description: "The maximum interval in seconds for log batches. This setting must be between 30 and 300 seconds (5 minutes), or `0` to disable it. Note that you cannot specify a minimum interval for log batches; this means that log files may be sent in shorter intervals than this. This parameter is only used for jobs with `edge` as its kind.",
				Computed:    true,
				Validators: []validator.Int64{
					int64validator.Between(30, 300),
				},
			},
			"max_upload_records": schema.Int64Attribute{
				Description: "The maximum number of log lines per batch. This setting must be between 1000 and 1,000,000 lines, or `0` to disable it. Note that you cannot specify a minimum number of log lines per batch; this means that log files may contain many fewer lines than this. This parameter is not available for jobs with `edge` as its kind.",
				Computed:    true,
				Validators: []validator.Int64{
					int64validator.Between(1000, 1000000),
				},
			},
			"dataset": schema.StringAttribute{
				Description: "Name of the dataset. A list of supported datasets can be found on the [Developer Docs](https://developers.cloudflare.com/logs/reference/log-fields/).",
				Computed:    true,
				Optional:    true,
			},
			"destination_conf": schema.StringAttribute{
				Description: "Uniquely identifies a resource (such as an s3 bucket) where data will be pushed. Additional configuration parameters supported by the destination may be included.",
				Computed:    true,
				Optional:    true,
			},
			"enabled": schema.BoolAttribute{
				Description: "Flag that indicates if the job is enabled.",
				Computed:    true,
				Optional:    true,
			},
			"error_message": schema.StringAttribute{
				Description: "If not null, the job is currently failing. Failures are usually repetitive (example: no permissions to write to destination bucket). Only the last failure is recorded. On successful execution of a job the error_message and last_error are set to null.",
				Computed:    true,
				Optional:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"id": schema.Int64Attribute{
				Description: "Unique id of the job.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"kind": schema.StringAttribute{
				Description: "The kind parameter (optional) is used to differentiate between Logpush and Edge Log Delivery jobs. Currently, Edge Log Delivery is only supported for the `http_requests` dataset.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("edge"),
				},
			},
			"last_complete": schema.StringAttribute{
				Description: "Records the last time for which logs have been successfully pushed. If the last successful push was for logs range 2018-07-23T10:00:00Z to 2018-07-23T10:01:00Z then the value of this field will be 2018-07-23T10:01:00Z. If the job has never run or has just been enabled and hasn't run yet then the field will be empty.",
				Computed:    true,
				Optional:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"last_error": schema.StringAttribute{
				Description: "Records the last time the job failed. If not null, the job is currently failing. If null, the job has either never failed or has run successfully at least once since last failure. See also the error_message field.",
				Computed:    true,
				Optional:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"logpull_options": schema.StringAttribute{
				Description: "This field is deprecated. Use `output_options` instead. Configuration string. It specifies things like requested fields and timestamp formats. If migrating from the logpull api, copy the url (full url or just the query string) of your call here, and logpush will keep on making this call for you, setting start and end times appropriately.",
				Computed:    true,
				Optional:    true,
			},
			"max_upload_bytes": schema.Int64Attribute{
				Description: "The maximum uncompressed file size of a batch of logs. This setting value must be between `5 MB` and `1 GB`, or `0` to disable it. Note that you cannot set a minimum file size; this means that log files may be much smaller than this batch size. This parameter is not available for jobs with `edge` as its kind.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.Between(5000000, 1000000000),
				},
			},
			"name": schema.StringAttribute{
				Description: "Optional human readable job name. Not unique. Cloudflare suggests that you set this to a meaningful string, like the domain name, to make it easier to identify your job.",
				Computed:    true,
				Optional:    true,
			},
			"output_options": schema.SingleNestedAttribute{
				Description: "The structured replacement for `logpull_options`. When including this field, the `logpull_option` field will be ignored.",
				Computed:    true,
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"batch_prefix": schema.StringAttribute{
						Description: "String to be prepended before each batch.",
						Computed:    true,
					},
					"batch_suffix": schema.StringAttribute{
						Description: "String to be appended after each batch.",
						Computed:    true,
					},
					"cve_2021_4428": schema.BoolAttribute{
						Description: "If set to true, will cause all occurrences of `${` in the generated files to be replaced with `x{`.",
						Computed:    true,
					},
					"field_delimiter": schema.StringAttribute{
						Description: "String to join fields. This field be ignored when `record_template` is set.",
						Computed:    true,
					},
					"field_names": schema.ListAttribute{
						Description: "List of field names to be included in the Logpush output. For the moment, there is no option to add all fields at once, so you must specify all the fields names you are interested in.",
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"output_type": schema.StringAttribute{
						Description: "Specifies the output type, such as `ndjson` or `csv`. This sets default values for the rest of the settings, depending on the chosen output type. Some formatting rules, like string quoting, are different between output types.",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("ndjson", "csv"),
						},
					},
					"record_delimiter": schema.StringAttribute{
						Description: "String to be inserted in-between the records as separator.",
						Computed:    true,
					},
					"record_prefix": schema.StringAttribute{
						Description: "String to be prepended before each record.",
						Computed:    true,
					},
					"record_suffix": schema.StringAttribute{
						Description: "String to be appended after each record.",
						Computed:    true,
					},
					"record_template": schema.StringAttribute{
						Description: "String to use as template for each record instead of the default comma-separated list. All fields used in the template must be present in `field_names` as well, otherwise they will end up as null. Format as a Go `text/template` without any standard functions, like conditionals, loops, sub-templates, etc.",
						Computed:    true,
					},
					"sample_rate": schema.Float64Attribute{
						Description: "Floating number to specify sampling rate. Sampling is applied on top of filtering, and regardless of the current `sample_interval` of the data.",
						Computed:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 1),
						},
					},
					"timestamp_format": schema.StringAttribute{
						Description: "String to specify the format for timestamps, such as `unixnano`, `unix`, or `rfc3339`.",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"unixnano",
								"unix",
								"rfc3339",
							),
						},
					},
				},
			},
		},
	}
}

func (d *LogpushJobDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *LogpushJobDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
