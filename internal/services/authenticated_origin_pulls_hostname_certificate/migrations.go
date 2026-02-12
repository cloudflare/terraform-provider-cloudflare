// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package authenticated_origin_pulls_hostname_certificate

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ resource.ResourceWithMoveState = (*AuthenticatedOriginPullsHostnameCertificateResource)(nil)

// authenticatedOriginPullsCertificateSourceSchema defines the source schema for moves
// from cloudflare_authenticated_origin_pulls_certificate with type="per-hostname"
// This represents the v4 cloudflare_authenticated_origin_pulls_certificate schema
var authenticatedOriginPullsCertificateSourceSchema = schema.Schema{
	Version: 0,
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed: true,
		},
		"zone_id": schema.StringAttribute{
			Required: true,
		},
		"certificate": schema.StringAttribute{
			Required: true,
		},
		"private_key": schema.StringAttribute{
			Required: true,
		},
		"type": schema.StringAttribute{
			Required: true,
		},
		"issuer": schema.StringAttribute{
			Computed: true,
		},
		"signature": schema.StringAttribute{
			Computed: true,
		},
		"serial_number": schema.StringAttribute{
			Computed: true,
		},
		"expires_on": schema.StringAttribute{
			Computed:   true,
			CustomType: timetypes.RFC3339Type{},
		},
		"status": schema.StringAttribute{
			Computed: true,
		},
		"uploaded_on": schema.StringAttribute{
			Computed:   true,
			CustomType: timetypes.RFC3339Type{},
		},
	},
}

// MoveState implements ResourceWithMoveState interface
// This enables moving state from cloudflare_authenticated_origin_pulls_certificate
// (with type="per-hostname") to cloudflare_authenticated_origin_pulls_hostname_certificate
func (r *AuthenticatedOriginPullsHostnameCertificateResource) MoveState(ctx context.Context) []resource.StateMover {
	return []resource.StateMover{
		{
			SourceSchema: &authenticatedOriginPullsCertificateSourceSchema,
			StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
				tflog.Info(ctx, "Starting state move from cloudflare_authenticated_origin_pulls_certificate to cloudflare_authenticated_origin_pulls_hostname_certificate")

				// Define source state structure (from cloudflare_authenticated_origin_pulls_certificate)
				var sourceState struct {
					ID           types.String      `tfsdk:"id"`
					ZoneID       types.String      `tfsdk:"zone_id"`
					Certificate  types.String      `tfsdk:"certificate"`
					PrivateKey   types.String      `tfsdk:"private_key"`
					Type         types.String      `tfsdk:"type"`
					Issuer       types.String      `tfsdk:"issuer"`
					Signature    types.String      `tfsdk:"signature"`
					SerialNumber types.String      `tfsdk:"serial_number"`
					ExpiresOn    timetypes.RFC3339 `tfsdk:"expires_on"`
					Status       types.String      `tfsdk:"status"`
					UploadedOn   timetypes.RFC3339 `tfsdk:"uploaded_on"`
				}

				// Get the source state
				resp.Diagnostics.Append(req.SourceState.Get(ctx, &sourceState)...)
				if resp.Diagnostics.HasError() {
					tflog.Error(ctx, "Failed to get source state during move")
					return
				}

				tflog.Debug(ctx, "Source state retrieved", map[string]interface{}{
					"zone_id": sourceState.ZoneID.ValueString(),
					"type":    sourceState.Type.ValueString(),
					"id":      sourceState.ID.ValueString(),
				})

				// Validate that this is a per-hostname certificate (the only type we should move)
				if !sourceState.Type.IsNull() && sourceState.Type.ValueString() != "per-hostname" {
					resp.Diagnostics.AddError(
						"Invalid State Move",
						fmt.Sprintf("Cannot move cloudflare_authenticated_origin_pulls_certificate with type='%s' to cloudflare_authenticated_origin_pulls_hostname_certificate. Only type='per-hostname' should be moved to this resource type.", sourceState.Type.ValueString()),
					)
					return
				}

				// Create the target state (for cloudflare_authenticated_origin_pulls_hostname_certificate)
				// Map all fields directly - the schemas are compatible except for the type field
				targetState := AuthenticatedOriginPullsHostnameCertificateModel{
					ID:           sourceState.ID,
					ZoneID:       sourceState.ZoneID,
					Certificate:  sourceState.Certificate,
					PrivateKey:   sourceState.PrivateKey,
					Issuer:       sourceState.Issuer,
					Signature:    sourceState.Signature,
					SerialNumber: sourceState.SerialNumber,
					ExpiresOn:    sourceState.ExpiresOn,
					Status:       sourceState.Status,
					UploadedOn:   sourceState.UploadedOn,
				}

				tflog.Info(ctx, "State move transformation completed", map[string]interface{}{
					"source_type": sourceState.Type.ValueString(),
					"zone_id":     targetState.ZoneID.ValueString(),
				})

				// Set the target state
				resp.Diagnostics.Append(resp.TargetState.Set(ctx, &targetState)...)
				if resp.Diagnostics.HasError() {
					tflog.Error(ctx, "Failed to set target state during move")
					return
				}

				tflog.Info(ctx, "State move from cloudflare_authenticated_origin_pulls_certificate to cloudflare_authenticated_origin_pulls_hostname_certificate completed successfully")
			},
		},
	}
}
