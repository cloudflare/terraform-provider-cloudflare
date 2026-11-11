// This file is intentionally NOT code-generated.

package v500_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_device_default_profile"
	v500 "github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_device_default_profile/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

const v4DefaultProfileState = `{
	"id": "f037e56e89293a057740de681ac9abbe",
	"account_id": "f037e56e89293a057740de681ac9abbe",
	"default": true,
	"name": "Default",
	"description": "default device settings",
	"enabled": true,
	"allow_mode_switch": true,
	"auto_connect": 0,
	"captive_portal": 180,
	"service_mode_v2_mode": "proxy",
	"service_mode_v2_port": 8080,
	"tunnel_protocol": "wireguard"
}`

// The target model must convert into the live schema, or both the moved block
// and UpgradeState fail with a Value Conversion Error.
func TestDeviceDefaultProfileMigrationFromV4SetsState(t *testing.T) {
	ctx := context.Background()
	sourceSchema := v500.SourceDeviceProfileSchema()
	targetSchema := zero_trust_device_default_profile.ResourceSchema(ctx)
	raw, err := (&tfprotov6.RawState{JSON: []byte(v4DefaultProfileState)}).Unmarshal(sourceSchema.Type().TerraformType(ctx))
	if err != nil {
		t.Fatalf("could not decode v4 state: %v", err)
	}
	sourceState := tfsdk.State{Raw: raw, Schema: sourceSchema}

	for _, tc := range []struct {
		name    string
		migrate func(t *testing.T) tfsdk.State
	}{
		{
			name: "MoveState",
			migrate: func(t *testing.T) tfsdk.State {
				resp := resource.MoveStateResponse{TargetState: tfsdk.State{Schema: targetSchema}}
				v500.MoveDeviceProfilesToDefaultProfile(ctx, resource.MoveStateRequest{
					SourceTypeName: "cloudflare_zero_trust_device_profiles",
					SourceState:    &sourceState,
				}, &resp)
				if resp.Diagnostics.HasError() {
					t.Fatalf("MoveState failed: %v", resp.Diagnostics)
				}
				return resp.TargetState
			},
		},
		{
			name: "UpgradeFromV4",
			migrate: func(t *testing.T) tfsdk.State {
				resp := resource.UpgradeStateResponse{State: tfsdk.State{Schema: targetSchema}}
				v500.UpgradeFromV4(ctx, resource.UpgradeStateRequest{State: &sourceState}, &resp)
				if resp.Diagnostics.HasError() {
					t.Fatalf("UpgradeFromV4 failed: %v", resp.Diagnostics)
				}
				return resp.State
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			state := tc.migrate(t)
			var got zero_trust_device_default_profile.ZeroTrustDeviceDefaultProfileModel
			if diags := state.Get(ctx, &got); diags.HasError() {
				t.Fatalf("could not decode migrated state: %v", diags)
			}
			smv2, diags := got.ServiceModeV2.Value(ctx)
			if diags.HasError() || smv2 == nil {
				t.Fatalf("service_mode_v2 missing after migration: %v", diags)
			}
			if !smv2.Mode.Equal(types.StringValue("proxy")) || !smv2.Port.Equal(types.Float64Value(8080)) {
				t.Errorf("service_mode_v2 = %+v, want proxy on 8080", smv2)
			}
			if !got.CaptivePortal.Equal(types.Float64Value(180)) {
				t.Errorf("captive_portal = %s, want 180", got.CaptivePortal)
			}
			if !got.EnableDeviceCredentialAuthentication.IsNull() {
				t.Errorf("enable_device_credential_authentication = %s, want null", got.EnableDeviceCredentialAuthentication)
			}
			if got.GlobalAcceleration != nil {
				t.Errorf("global_acceleration = %+v, want null", got.GlobalAcceleration)
			}
		})
	}
}

// A configured global_acceleration must also convert into the live schema,
// which needs allow_manual_override and autoswitch on the target model.
func TestDeviceDefaultProfileMigrationTargetGlobalAccelerationSetsState(t *testing.T) {
	ctx := context.Background()
	target, diags := v500.TransformToDefaultProfile(ctx, v500.SourceDeviceProfileModel{
		ID:        types.StringValue("f037e56e89293a057740de681ac9abbe"),
		AccountID: types.StringValue("f037e56e89293a057740de681ac9abbe"),
	})
	if diags.HasError() {
		t.Fatalf("transform failed: %v", diags)
	}
	endpoints := customfield.NewListMust[types.String](ctx, []attr.Value{types.StringValue("192.0.2.1:443")})
	target.GlobalAcceleration, diags = customfield.NewObject(ctx, &v500.TargetDefaultProfileGlobalAccelerationModel{
		APIEndpoints:        endpoints,
		Enabled:             types.BoolValue(true),
		MasqueEndpoints:     endpoints,
		WireguardEndpoints:  endpoints,
		AllowManualOverride: types.BoolValue(true),
		Autoswitch:          types.BoolValue(false),
	})
	if diags.HasError() {
		t.Fatalf("could not build global_acceleration: %v", diags)
	}

	state := tfsdk.State{Schema: zero_trust_device_default_profile.ResourceSchema(ctx)}
	if diags := state.Set(ctx, target); diags.HasError() {
		t.Fatalf("could not set target into live schema: %v", diags)
	}
	var got zero_trust_device_default_profile.ZeroTrustDeviceDefaultProfileModel
	if diags := state.Get(ctx, &got); diags.HasError() {
		t.Fatalf("could not decode state: %v", diags)
	}
	if got.GlobalAcceleration == nil || !got.GlobalAcceleration.Enabled.ValueBool() {
		t.Errorf("global_acceleration = %+v, want enabled", got.GlobalAcceleration)
	}
}
