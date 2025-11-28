// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_settings

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*ZeroTrustDeviceSettingsResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"disable_for_time": schema.Float64Attribute{
				Description: "Sets the time limit, in seconds, that a user can use an override code to bypass WARP.",
				Optional:    true,
			},
			"external_emergency_signal_enabled": schema.BoolAttribute{
				Description: "Controls whether the external emergency disconnect feature is enabled.",
				Optional:    true,
			},
			"external_emergency_signal_fingerprint": schema.StringAttribute{
				Description: "The SHA256 fingerprint (64 hexadecimal characters) of the HTTPS server certificate for the external_emergency_signal_url. If provided, the WARP client will use this value to verify the server's identity. The device will ignore any response if the server's certificate fingerprint does not exactly match this value.",
				Optional:    true,
			},
			"external_emergency_signal_interval": schema.StringAttribute{
				Description: `The interval at which the WARP client fetches the emergency disconnect signal, formatted as a duration string (e.g., "5m", "2m30s", "1h"). Minimum 30 seconds.`,
				Optional:    true,
			},
			"external_emergency_signal_url": schema.StringAttribute{
				Description: "The HTTPS URL from which to fetch the emergency disconnect signal. Must use HTTPS and have an IPv4 or IPv6 address as the host.",
				Optional:    true,
			},
			"gateway_proxy_enabled": schema.BoolAttribute{
				Description: "Enable gateway proxy filtering on TCP.",
				Optional:    true,
			},
			"gateway_udp_proxy_enabled": schema.BoolAttribute{
				Description: "Enable gateway proxy filtering on UDP.",
				Optional:    true,
			},
			"root_certificate_installation_enabled": schema.BoolAttribute{
				Description: "Enable installation of cloudflare managed root certificate.",
				Optional:    true,
			},
			"use_zt_virtual_ip": schema.BoolAttribute{
				Description: "Enable using CGNAT virtual IPv4.",
				Optional:    true,
			},
		},
	}
}

func (r *ZeroTrustDeviceSettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ZeroTrustDeviceSettingsResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
