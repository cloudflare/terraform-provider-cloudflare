package hyperdrive_config

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestModifyPlanModifiedOn(t *testing.T) {
	ctx := context.Background()

	createdOn, diags := timetypes.NewRFC3339Value("2026-07-15T06:00:00Z")
	if diags.HasError() {
		t.Fatalf("failed to create created_on value: %v", diags)
	}
	modifiedOn, diags := timetypes.NewRFC3339Value("2026-07-15T07:00:00Z")
	if diags.HasError() {
		t.Fatalf("failed to create modified_on value: %v", diags)
	}

	newModel := func(mtls *HyperdriveConfigMTLSModel, created, modified timetypes.RFC3339) *HyperdriveConfigModel {
		return &HyperdriveConfigModel{
			ID:        types.StringValue("hyperdrive-id"),
			AccountID: types.StringValue("account-id"),
			Name:      types.StringValue("example"),
			Origin: &HyperdriveConfigOriginModel{
				Database:           types.StringValue("example"),
				Host:               types.StringNull(),
				Password:           types.StringValue("password"),
				Port:               types.Int64Null(),
				Scheme:             types.StringValue("postgresql"),
				User:               types.StringValue("example"),
				AccessClientID:     types.StringNull(),
				AccessClientSecret: types.StringValue("access-secret"),
				ServiceID:          types.StringValue("vpc-service-id"),
			},
			OriginConnectionLimit: types.Int64Value(20),
			Caching: &HyperdriveConfigCachingModel{
				Disabled:             types.BoolValue(false),
				MaxAge:               types.Int64Null(),
				StaleWhileRevalidate: types.Int64Null(),
			},
			MTLS:       mtls,
			CreatedOn:  created,
			ModifiedOn: modified,
		}
	}

	tests := []struct {
		name                    string
		mutate                  func(state, plan *HyperdriveConfigModel)
		expectModifiedOnUnknown bool
	}{
		{
			name: "no configuration changes",
		},
		{
			name: "password added after import",
			mutate: func(state, _ *HyperdriveConfigModel) {
				state.Origin.Password = types.StringNull()
			},
			expectModifiedOnUnknown: true,
		},
		{
			name: "access client secret added",
			mutate: func(state, _ *HyperdriveConfigModel) {
				state.Origin.AccessClientSecret = types.StringNull()
			},
			expectModifiedOnUnknown: true,
		},
		{
			name: "service ID changed",
			mutate: func(_, plan *HyperdriveConfigModel) {
				plan.Origin.ServiceID = types.StringValue("new-vpc-service-id")
			},
			expectModifiedOnUnknown: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			stateData := newModel(
				&HyperdriveConfigMTLSModel{
					CACertificateID:   types.StringNull(),
					MTLSCertificateID: types.StringNull(),
					Sslmode:           types.StringNull(),
				},
				createdOn,
				modifiedOn,
			)
			planData := newModel(
				nil,
				timetypes.NewRFC3339Unknown(),
				timetypes.NewRFC3339Unknown(),
			)
			if test.mutate != nil {
				test.mutate(stateData, planData)
			}

			schema := ResourceSchema(ctx)
			state := tfsdk.State{Schema: schema}
			if diags := state.Set(ctx, stateData); diags.HasError() {
				t.Fatalf("failed to build state: %v", diags)
			}
			plan := tfsdk.Plan{Schema: schema}
			if diags := plan.Set(ctx, planData); diags.HasError() {
				t.Fatalf("failed to build plan: %v", diags)
			}

			resp := &resource.ModifyPlanResponse{Plan: plan}
			modifyPlan(ctx, resource.ModifyPlanRequest{Plan: plan, State: state}, resp)
			if resp.Diagnostics.HasError() {
				t.Fatalf("modifyPlan returned diagnostics: %v", resp.Diagnostics)
			}

			var updatedPlan *HyperdriveConfigModel
			if diags := resp.Plan.Get(ctx, &updatedPlan); diags.HasError() {
				t.Fatalf("failed to read updated plan: %v", diags)
			}
			if updatedPlan == nil {
				t.Fatal("expected a populated updated plan")
			}
			if updatedPlan.CreatedOn.IsUnknown() || updatedPlan.CreatedOn.ValueString() != createdOn.ValueString() {
				t.Fatalf("expected created_on to be preserved as %q, got %q", createdOn.ValueString(), updatedPlan.CreatedOn.ValueString())
			}

			if test.expectModifiedOnUnknown {
				if !updatedPlan.ModifiedOn.IsUnknown() {
					t.Fatalf("expected modified_on to remain unknown during update, got %q", updatedPlan.ModifiedOn.ValueString())
				}
				return
			}
			if updatedPlan.ModifiedOn.IsUnknown() || updatedPlan.ModifiedOn.ValueString() != modifiedOn.ValueString() {
				t.Fatalf("expected modified_on to be preserved as %q, got %q", modifiedOn.ValueString(), updatedPlan.ModifiedOn.ValueString())
			}
		})
	}
}

func TestPreserveWriteOnlyFields(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		source   *HyperdriveConfigModel
		dest     *HyperdriveConfigModel
		expected *HyperdriveConfigModel
	}{
		{
			name:   "nil source does nothing",
			source: nil,
			dest: &HyperdriveConfigModel{
				Origin: &HyperdriveConfigOriginModel{
					Password: types.StringNull(),
				},
			},
			expected: &HyperdriveConfigModel{
				Origin: &HyperdriveConfigOriginModel{
					Password: types.StringNull(),
				},
			},
		},
		{
			name: "nil dest does nothing",
			source: &HyperdriveConfigModel{
				Origin: &HyperdriveConfigOriginModel{
					Password: types.StringValue("secret"),
				},
			},
			dest:     nil,
			expected: nil,
		},
		{
			name: "preserves password from source to dest",
			source: &HyperdriveConfigModel{
				Origin: &HyperdriveConfigOriginModel{
					Password:           types.StringValue("my-secret-password"),
					AccessClientSecret: types.StringNull(),
				},
			},
			dest: &HyperdriveConfigModel{
				Origin: &HyperdriveConfigOriginModel{
					Password:           types.StringNull(),
					AccessClientSecret: types.StringNull(),
				},
			},
			expected: &HyperdriveConfigModel{
				Origin: &HyperdriveConfigOriginModel{
					Password:           types.StringValue("my-secret-password"),
					AccessClientSecret: types.StringNull(),
				},
			},
		},
		{
			name: "preserves access_client_secret from source to dest",
			source: &HyperdriveConfigModel{
				Origin: &HyperdriveConfigOriginModel{
					Password:           types.StringValue("password"),
					AccessClientSecret: types.StringValue("my-access-secret"),
				},
			},
			dest: &HyperdriveConfigModel{
				Origin: &HyperdriveConfigOriginModel{
					Password:           types.StringNull(),
					AccessClientSecret: types.StringNull(),
				},
			},
			expected: &HyperdriveConfigModel{
				Origin: &HyperdriveConfigOriginModel{
					Password:           types.StringValue("password"),
					AccessClientSecret: types.StringValue("my-access-secret"),
				},
			},
		},
		{
			name: "preserves both password and access_client_secret",
			source: &HyperdriveConfigModel{
				Origin: &HyperdriveConfigOriginModel{
					Database:           types.StringValue("mydb"),
					Host:               types.StringValue("localhost"),
					Password:           types.StringValue("password123"),
					AccessClientID:     types.StringValue("client-id"),
					AccessClientSecret: types.StringValue("client-secret"),
				},
			},
			dest: &HyperdriveConfigModel{
				Origin: &HyperdriveConfigOriginModel{
					Database:           types.StringValue("mydb"),
					Host:               types.StringValue("localhost"),
					Password:           types.StringNull(), // API doesn't return this
					AccessClientID:     types.StringValue("client-id"),
					AccessClientSecret: types.StringNull(), // API doesn't return this
				},
			},
			expected: &HyperdriveConfigModel{
				Origin: &HyperdriveConfigOriginModel{
					Database:           types.StringValue("mydb"),
					Host:               types.StringValue("localhost"),
					Password:           types.StringValue("password123"),
					AccessClientID:     types.StringValue("client-id"),
					AccessClientSecret: types.StringValue("client-secret"),
				},
			},
		},
		{
			name: "does not overwrite when source has unknown value",
			source: &HyperdriveConfigModel{
				Origin: &HyperdriveConfigOriginModel{
					Password: types.StringUnknown(),
				},
			},
			dest: &HyperdriveConfigModel{
				Origin: &HyperdriveConfigOriginModel{
					Password: types.StringNull(),
				},
			},
			expected: &HyperdriveConfigModel{
				Origin: &HyperdriveConfigOriginModel{
					Password: types.StringNull(),
				},
			},
		},
		{
			name: "handles nil origin in source",
			source: &HyperdriveConfigModel{
				Origin: nil,
			},
			dest: &HyperdriveConfigModel{
				Origin: &HyperdriveConfigOriginModel{
					Password: types.StringNull(),
				},
			},
			expected: &HyperdriveConfigModel{
				Origin: &HyperdriveConfigOriginModel{
					Password: types.StringNull(),
				},
			},
		},
		{
			name: "handles nil origin in dest",
			source: &HyperdriveConfigModel{
				Origin: &HyperdriveConfigOriginModel{
					Password: types.StringValue("password"),
				},
			},
			dest: &HyperdriveConfigModel{
				Origin: nil,
			},
			expected: &HyperdriveConfigModel{
				Origin: nil,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			preserveWriteOnlyFields(tt.source, tt.dest)

			// For nil dest, just verify it stayed nil
			if tt.expected == nil {
				if tt.dest != nil {
					t.Errorf("expected dest to be nil, got %v", tt.dest)
				}
				return
			}

			// For nil expected origin, verify dest origin is nil
			if tt.expected.Origin == nil {
				if tt.dest.Origin != nil {
					t.Errorf("expected dest.Origin to be nil, got %v", tt.dest.Origin)
				}
				return
			}

			// Verify password
			if !tt.dest.Origin.Password.Equal(tt.expected.Origin.Password) {
				t.Errorf("password mismatch: got %v, want %v",
					tt.dest.Origin.Password, tt.expected.Origin.Password)
			}

			// Verify access_client_secret
			if !tt.dest.Origin.AccessClientSecret.Equal(tt.expected.Origin.AccessClientSecret) {
				t.Errorf("access_client_secret mismatch: got %v, want %v",
					tt.dest.Origin.AccessClientSecret, tt.expected.Origin.AccessClientSecret)
			}
		})
	}
}
