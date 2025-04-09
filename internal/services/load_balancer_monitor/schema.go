// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancer_monitor

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
  "github.com/hashicorp/terraform-plugin-framework/resource"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
  "github.com/hashicorp/terraform-plugin-framework/schema/validator"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*LoadBalancerMonitorResource)(nil)

func ResourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "id": schema.StringAttribute{
        Computed: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
      },
      "account_id": schema.StringAttribute{
        Description: "Identifier",
        Required: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
      },
      "description": schema.StringAttribute{
        Description: "Object description.",
        Optional: true,
      },
      "expected_body": schema.StringAttribute{
        Description: "A case-insensitive sub-string to look for in the response body. If this string is not found, the origin will be marked as unhealthy. This parameter is only valid for HTTP and HTTPS monitors.",
        Optional: true,
      },
      "expected_codes": schema.StringAttribute{
        Description: "The expected HTTP response code or code range of the health check. This parameter is only valid for HTTP and HTTPS monitors.",
        Optional: true,
      },
      "probe_zone": schema.StringAttribute{
        Description: "Assign this monitor to emulate the specified zone while probing. This parameter is only valid for HTTP and HTTPS monitors.",
        Optional: true,
      },
      "header": schema.MapAttribute{
        Description: "The HTTP request headers to send in the health check. It is recommended you set a Host header by default. The User-Agent header cannot be overridden. This parameter is only valid for HTTP and HTTPS monitors.",
        Optional: true,
        ElementType: types.ListType{
          ElemType: types.StringType,
        },
      },
      "allow_insecure": schema.BoolAttribute{
        Description: "Do not validate the certificate when monitor use HTTPS. This parameter is currently only valid for HTTP and HTTPS monitors.",
        Computed: true,
        Optional: true,
        Default: booldefault.  StaticBool(false),
      },
      "consecutive_down": schema.Int64Attribute{
        Description: "To be marked unhealthy the monitored origin must fail this healthcheck N consecutive times.",
        Computed: true,
        Optional: true,
        Default: int64default.  StaticInt64(0),
      },
      "consecutive_up": schema.Int64Attribute{
        Description: "To be marked healthy the monitored origin must pass this healthcheck N consecutive times.",
        Computed: true,
        Optional: true,
        Default: int64default.  StaticInt64(0),
      },
      "follow_redirects": schema.BoolAttribute{
        Description: "Follow redirects if returned by the origin. This parameter is only valid for HTTP and HTTPS monitors.",
        Computed: true,
        Optional: true,
        Default: booldefault.  StaticBool(false),
      },
      "interval": schema.Int64Attribute{
        Description: "The interval between each health check. Shorter intervals may improve failover time, but will increase load on the origins as we check from multiple locations.",
        Computed: true,
        Optional: true,
        Default: int64default.  StaticInt64(60),
      },
      "method": schema.StringAttribute{
        Description: "The method to use for the health check. This defaults to 'GET' for HTTP/HTTPS based checks and 'connection_established' for TCP based health checks.",
        Computed: true,
        Optional: true,
        Default: stringdefault.  StaticString("GET"),
      },
      "path": schema.StringAttribute{
        Description: "The endpoint path you want to conduct a health check against. This parameter is only valid for HTTP and HTTPS monitors.",
        Computed: true,
        Optional: true,
        Default: stringdefault.  StaticString("/"),
      },
      "port": schema.Int64Attribute{
        Description: "The port number to connect to for the health check. Required for TCP, UDP, and SMTP checks. HTTP and HTTPS checks should only define the port when using a non-standard port (HTTP: default 80, HTTPS: default 443).",
        Computed: true,
        Optional: true,
        Default: int64default.  StaticInt64(0),
      },
      "retries": schema.Int64Attribute{
        Description: "The number of retries to attempt in case of a timeout before marking the origin as unhealthy. Retries are attempted immediately.",
        Computed: true,
        Optional: true,
        Default: int64default.  StaticInt64(2),
      },
      "timeout": schema.Int64Attribute{
        Description: "The timeout (in seconds) before marking the health check as failed.",
        Computed: true,
        Optional: true,
        Default: int64default.  StaticInt64(5),
      },
      "type": schema.StringAttribute{
        Description: "The protocol to use for the health check. Currently supported protocols are 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.\nAvailable values: \"http\", \"https\", \"tcp\", \"udp_icmp\", \"icmp_ping\", \"smtp\".",
        Computed: true,
        Optional: true,
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
        Default: stringdefault.  StaticString("http"),
      },
      "created_on": schema.StringAttribute{
        Computed: true,
      },
      "modified_on": schema.StringAttribute{
        Computed: true,
      },
    },
  }
}

func (r *LoadBalancerMonitorResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
  resp.Schema = ResourceSchema(ctx)
}

func (r *LoadBalancerMonitorResource) ConfigValidators(_ context.Context) ([]resource.ConfigValidator) {
  return []resource.ConfigValidator{
  }
}
