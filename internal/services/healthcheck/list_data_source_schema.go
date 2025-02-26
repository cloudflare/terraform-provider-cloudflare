// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package healthcheck

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*HealthchecksDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
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
				CustomType:  customfield.NewNestedObjectListType[HealthchecksResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Identifier",
							Computed:    true,
						},
						"address": schema.StringAttribute{
							Description: "The hostname or IP address of the origin server to run health checks on.",
							Computed:    true,
						},
						"check_regions": schema.ListAttribute{
							Description: "A list of regions from which to run health checks. Null means Cloudflare will pick a default region.",
							Computed:    true,
							Validators: []validator.List{
								listvalidator.ValueStringsAre(
									stringvalidator.OneOfCaseInsensitive(
										"WNAM",
										"ENAM",
										"WEU",
										"EEU",
										"NSAM",
										"SSAM",
										"OC",
										"ME",
										"NAF",
										"SAF",
										"IN",
										"SEAS",
										"NEAS",
										"ALL_REGIONS",
									),
								),
							},
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
						"consecutive_fails": schema.Int64Attribute{
							Description: "The number of consecutive fails required from a health check before changing the health to unhealthy.",
							Computed:    true,
						},
						"consecutive_successes": schema.Int64Attribute{
							Description: "The number of consecutive successes required from a health check before changing the health to healthy.",
							Computed:    true,
						},
						"created_on": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"description": schema.StringAttribute{
							Description: "A human-readable description of the health check.",
							Computed:    true,
						},
						"failure_reason": schema.StringAttribute{
							Description: "The current failure reason if status is unhealthy.",
							Computed:    true,
						},
						"http_config": schema.SingleNestedAttribute{
							Description: "Parameters specific to an HTTP or HTTPS health check.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[HealthchecksHTTPConfigDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"allow_insecure": schema.BoolAttribute{
									Description: "Do not validate the certificate when the health check uses HTTPS.",
									Computed:    true,
								},
								"expected_body": schema.StringAttribute{
									Description: "A case-insensitive sub-string to look for in the response body. If this string is not found, the origin will be marked as unhealthy.",
									Computed:    true,
								},
								"expected_codes": schema.ListAttribute{
									Description: "The expected HTTP response codes (e.g. \"200\") or code ranges (e.g. \"2xx\" for all codes starting with 2) of the health check.",
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"follow_redirects": schema.BoolAttribute{
									Description: "Follow redirects if the origin returns a 3xx status code.",
									Computed:    true,
								},
								"header": schema.MapAttribute{
									Description: "The HTTP request headers to send in the health check. It is recommended you set a Host header by default. The User-Agent header cannot be overridden.",
									Computed:    true,
									CustomType:  customfield.NewMapType[customfield.List[types.String]](ctx),
									ElementType: types.ListType{
										ElemType: types.StringType,
									},
								},
								"method": schema.StringAttribute{
									Description: "The HTTP method to use for the health check.\nAvailable values: \"GET\", \"HEAD\".",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("GET", "HEAD"),
									},
								},
								"path": schema.StringAttribute{
									Description: "The endpoint path to health check against.",
									Computed:    true,
								},
								"port": schema.Int64Attribute{
									Description: "Port number to connect to for the health check. Defaults to 80 if type is HTTP or 443 if type is HTTPS.",
									Computed:    true,
								},
							},
						},
						"interval": schema.Int64Attribute{
							Description: "The interval between each health check. Shorter intervals may give quicker notifications if the origin status changes, but will increase load on the origin as we check from multiple locations.",
							Computed:    true,
						},
						"modified_on": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"name": schema.StringAttribute{
							Description: "A short name to identify the health check. Only alphanumeric characters, hyphens and underscores are allowed.",
							Computed:    true,
						},
						"retries": schema.Int64Attribute{
							Description: "The number of retries to attempt in case of a timeout before marking the origin as unhealthy. Retries are attempted immediately.",
							Computed:    true,
						},
						"status": schema.StringAttribute{
							Description: "The current status of the origin server according to the health check.\nAvailable values: \"unknown\", \"healthy\", \"unhealthy\", \"suspended\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"unknown",
									"healthy",
									"unhealthy",
									"suspended",
								),
							},
						},
						"suspended": schema.BoolAttribute{
							Description: "If suspended, no health checks are sent to the origin.",
							Computed:    true,
						},
						"tcp_config": schema.SingleNestedAttribute{
							Description: "Parameters specific to TCP health check.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[HealthchecksTCPConfigDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"method": schema.StringAttribute{
									Description: "The TCP connection method to use for the health check.\nAvailable values: \"connection_established\".",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("connection_established"),
									},
								},
								"port": schema.Int64Attribute{
									Description: "Port number to connect to for the health check. Defaults to 80.",
									Computed:    true,
								},
							},
						},
						"timeout": schema.Int64Attribute{
							Description: "The timeout (in seconds) before marking the health check as failed.",
							Computed:    true,
						},
						"type": schema.StringAttribute{
							Description: "The protocol to use for the health check. Currently supported protocols are 'HTTP', 'HTTPS' and 'TCP'.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *HealthchecksDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *HealthchecksDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
