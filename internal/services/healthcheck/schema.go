// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package healthcheck

import (
  "context"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
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

var _ resource.ResourceWithConfigValidators = (*HealthcheckResource)(nil)

func ResourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "id": schema.StringAttribute{
        Description: "Identifier",
        Computed: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
      },
      "zone_id": schema.StringAttribute{
        Description: "Identifier",
        Required: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
      },
      "address": schema.StringAttribute{
        Description: "The hostname or IP address of the origin server to run health checks on.",
        Required: true,
      },
      "name": schema.StringAttribute{
        Description: "A short name to identify the health check. Only alphanumeric characters, hyphens and underscores are allowed.",
        Required: true,
      },
      "description": schema.StringAttribute{
        Description: "A human-readable description of the health check.",
        Optional: true,
      },
      "check_regions": schema.ListAttribute{
        Description: "A list of regions from which to run health checks. Null means Cloudflare will pick a default region.",
        Optional: true,
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
        ElementType: types.StringType,
      },
      "consecutive_fails": schema.Int64Attribute{
        Description: "The number of consecutive fails required from a health check before changing the health to unhealthy.",
        Computed: true,
        Optional: true,
        Default: int64default.  StaticInt64(1),
      },
      "consecutive_successes": schema.Int64Attribute{
        Description: "The number of consecutive successes required from a health check before changing the health to healthy.",
        Computed: true,
        Optional: true,
        Default: int64default.  StaticInt64(1),
      },
      "interval": schema.Int64Attribute{
        Description: "The interval between each health check. Shorter intervals may give quicker notifications if the origin status changes, but will increase load on the origin as we check from multiple locations.",
        Computed: true,
        Optional: true,
        Default: int64default.  StaticInt64(60),
      },
      "retries": schema.Int64Attribute{
        Description: "The number of retries to attempt in case of a timeout before marking the origin as unhealthy. Retries are attempted immediately.",
        Computed: true,
        Optional: true,
        Default: int64default.  StaticInt64(2),
      },
      "suspended": schema.BoolAttribute{
        Description: "If suspended, no health checks are sent to the origin.",
        Computed: true,
        Optional: true,
        Default: booldefault.  StaticBool(false),
      },
      "timeout": schema.Int64Attribute{
        Description: "The timeout (in seconds) before marking the health check as failed.",
        Computed: true,
        Optional: true,
        Default: int64default.  StaticInt64(5),
      },
      "type": schema.StringAttribute{
        Description: "The protocol to use for the health check. Currently supported protocols are 'HTTP', 'HTTPS' and 'TCP'.",
        Computed: true,
        Optional: true,
        Default: stringdefault.  StaticString("HTTP"),
      },
      "http_config": schema.SingleNestedAttribute{
        Description: "Parameters specific to an HTTP or HTTPS health check.",
        Computed: true,
        Optional: true,
        CustomType: customfield.NewNestedObjectType[HealthcheckHTTPConfigModel](ctx),
        Attributes: map[string]schema.Attribute{
          "allow_insecure": schema.BoolAttribute{
            Description: "Do not validate the certificate when the health check uses HTTPS.",
            Computed: true,
            Optional: true,
            Default: booldefault.  StaticBool(false),
          },
          "expected_body": schema.StringAttribute{
            Description: "A case-insensitive sub-string to look for in the response body. If this string is not found, the origin will be marked as unhealthy.",
            Optional: true,
          },
          "expected_codes": schema.ListAttribute{
            Description: `The expected HTTP response codes (e.g. "200") or code ranges (e.g. "2xx" for all codes starting with 2) of the health check.`,
            Optional: true,
            ElementType: types.StringType,
          },
          "follow_redirects": schema.BoolAttribute{
            Description: "Follow redirects if the origin returns a 3xx status code.",
            Computed: true,
            Optional: true,
            Default: booldefault.  StaticBool(false),
          },
          "header": schema.MapAttribute{
            Description: "The HTTP request headers to send in the health check. It is recommended you set a Host header by default. The User-Agent header cannot be overridden.",
            Optional: true,
            ElementType: types.ListType{
              ElemType: types.StringType,
            },
          },
          "method": schema.StringAttribute{
            Description: "The HTTP method to use for the health check.\nAvailable values: \"GET\", \"HEAD\".",
            Computed: true,
            Optional: true,
            Validators: []validator.String{
            stringvalidator.OneOfCaseInsensitive("GET", "HEAD"),
            },
            Default: stringdefault.  StaticString("GET"),
          },
          "path": schema.StringAttribute{
            Description: "The endpoint path to health check against.",
            Computed: true,
            Optional: true,
            Default: stringdefault.  StaticString("/"),
          },
          "port": schema.Int64Attribute{
            Description: "Port number to connect to for the health check. Defaults to 80 if type is HTTP or 443 if type is HTTPS.",
            Computed: true,
            Optional: true,
            Default: int64default.  StaticInt64(80),
          },
        },
      },
      "tcp_config": schema.SingleNestedAttribute{
        Description: "Parameters specific to TCP health check.",
        Computed: true,
        Optional: true,
        CustomType: customfield.NewNestedObjectType[HealthcheckTCPConfigModel](ctx),
        Attributes: map[string]schema.Attribute{
          "method": schema.StringAttribute{
            Description: "The TCP connection method to use for the health check.\nAvailable values: \"connection_established\".",
            Computed: true,
            Optional: true,
            Validators: []validator.String{
            stringvalidator.OneOfCaseInsensitive("connection_established"),
            },
            Default: stringdefault.  StaticString("connection_established"),
          },
          "port": schema.Int64Attribute{
            Description: "Port number to connect to for the health check. Defaults to 80.",
            Computed: true,
            Optional: true,
            Default: int64default.  StaticInt64(80),
          },
        },
      },
      "created_on": schema.StringAttribute{
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "failure_reason": schema.StringAttribute{
        Description: "The current failure reason if status is unhealthy.",
        Computed: true,
      },
      "modified_on": schema.StringAttribute{
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "status": schema.StringAttribute{
        Description: "The current status of the origin server according to the health check.\nAvailable values: \"unknown\", \"healthy\", \"unhealthy\", \"suspended\".",
        Computed: true,
        Validators: []validator.String{
        stringvalidator.OneOfCaseInsensitive(
          "unknown",
          "healthy",
          "unhealthy",
          "suspended",
        ),
        },
      },
    },
  }
}

func (r *HealthcheckResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
  resp.Schema = ResourceSchema(ctx)
}

func (r *HealthcheckResource) ConfigValidators(_ context.Context) ([]resource.ConfigValidator) {
  return []resource.ConfigValidator{
  }
}
