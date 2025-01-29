// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancer_monitor

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*LoadBalancerMonitorDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"monitor_id": schema.StringAttribute{
				Optional: true,
			},
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"allow_insecure": schema.BoolAttribute{
				Description: "Do not validate the certificate when monitor use HTTPS. This parameter is currently only valid for HTTP and HTTPS monitors.",
				Computed:    true,
			},
			"consecutive_down": schema.Int64Attribute{
				Description: "To be marked unhealthy the monitored origin must fail this healthcheck N consecutive times.",
				Computed:    true,
			},
			"consecutive_up": schema.Int64Attribute{
				Description: "To be marked healthy the monitored origin must pass this healthcheck N consecutive times.",
				Computed:    true,
			},
			"created_on": schema.StringAttribute{
				Computed: true,
			},
			"description": schema.StringAttribute{
				Description: "Object description.",
				Computed:    true,
			},
			"expected_body": schema.StringAttribute{
				Description: "A case-insensitive sub-string to look for in the response body. If this string is not found, the origin will be marked as unhealthy. This parameter is only valid for HTTP and HTTPS monitors.",
				Computed:    true,
			},
			"expected_codes": schema.StringAttribute{
				Description: "The expected HTTP response code or code range of the health check. This parameter is only valid for HTTP and HTTPS monitors.",
				Computed:    true,
			},
			"follow_redirects": schema.BoolAttribute{
				Description: "Follow redirects if returned by the origin. This parameter is only valid for HTTP and HTTPS monitors.",
				Computed:    true,
			},
			"interval": schema.Int64Attribute{
				Description: "The interval between each health check. Shorter intervals may improve failover time, but will increase load on the origins as we check from multiple locations.",
				Computed:    true,
			},
			"method": schema.StringAttribute{
				Description: "The method to use for the health check. This defaults to 'GET' for HTTP/HTTPS based checks and 'connection_established' for TCP based health checks.",
				Computed:    true,
			},
			"modified_on": schema.StringAttribute{
				Computed: true,
			},
			"path": schema.StringAttribute{
				Description: "The endpoint path you want to conduct a health check against. This parameter is only valid for HTTP and HTTPS monitors.",
				Computed:    true,
			},
			"port": schema.Int64Attribute{
				Description: "The port number to connect to for the health check. Required for TCP, UDP, and SMTP checks. HTTP and HTTPS checks should only define the port when using a non-standard port (HTTP: default 80, HTTPS: default 443).",
				Computed:    true,
			},
			"probe_zone": schema.StringAttribute{
				Description: "Assign this monitor to emulate the specified zone while probing. This parameter is only valid for HTTP and HTTPS monitors.",
				Computed:    true,
			},
			"retries": schema.Int64Attribute{
				Description: "The number of retries to attempt in case of a timeout before marking the origin as unhealthy. Retries are attempted immediately.",
				Computed:    true,
			},
			"timeout": schema.Int64Attribute{
				Description: "The timeout (in seconds) before marking the health check as failed.",
				Computed:    true,
			},
			"type": schema.StringAttribute{
				Description: "The protocol to use for the health check. Currently supported protocols are 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"http",
						"https",
						"tcp",
						"udp_icmp",
						"icmp_ping",
						"smtp",
					),
				},
			},
			"header": schema.MapAttribute{
				Description: "The HTTP request headers to send in the health check. It is recommended you set a Host header by default. The User-Agent header cannot be overridden. This parameter is only valid for HTTP and HTTPS monitors.",
				Computed:    true,
				CustomType:  customfield.NewMapType[customfield.List[types.String]](ctx),
				ElementType: types.ListType{
					ElemType: types.StringType,
				},
			},
		},
	}
}

func (d *LoadBalancerMonitorDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *LoadBalancerMonitorDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
