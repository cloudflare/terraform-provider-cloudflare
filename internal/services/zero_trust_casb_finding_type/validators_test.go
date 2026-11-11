package zero_trust_casb_finding_type

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testIntegrationID = "019e702a-f9ab-7a3a-8790-2c1e434acf47"

func TestIntegrationScopeInvalid(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	tests := []struct {
		name            string
		appliesToAll    types.Bool
		integrationIDs  customfield.List[types.String]
		expectedInvalid bool
	}{
		{
			name:            "false with null integration IDs",
			appliesToAll:    types.BoolValue(false),
			integrationIDs:  customfield.NullList[types.String](ctx),
			expectedInvalid: true,
		},
		{
			name:            "false with empty integration IDs",
			appliesToAll:    types.BoolValue(false),
			integrationIDs:  customfield.NewListMust[types.String](ctx, []attr.Value{}),
			expectedInvalid: true,
		},
		{
			name:         "false with integration IDs",
			appliesToAll: types.BoolValue(false),
			integrationIDs: customfield.NewListMust[types.String](ctx, []attr.Value{
				types.StringValue(testIntegrationID),
			}),
		},
		{
			name:           "true with empty integration IDs",
			appliesToAll:   types.BoolValue(true),
			integrationIDs: customfield.NewListMust[types.String](ctx, []attr.Value{}),
		},
		{
			name:           "null applies to all",
			appliesToAll:   types.BoolNull(),
			integrationIDs: customfield.NullList[types.String](ctx),
		},
		{
			name:           "unknown applies to all",
			appliesToAll:   types.BoolUnknown(),
			integrationIDs: customfield.NullList[types.String](ctx),
		},
		{
			name:           "unknown integration IDs",
			appliesToAll:   types.BoolValue(false),
			integrationIDs: customfield.UnknownList[types.String](ctx),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expectedInvalid, integrationScopeInvalid(tt.appliesToAll, tt.integrationIDs))
		})
	}
}

func TestAppliesToAllWithIDs(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	someIDs := customfield.NewListMust[types.String](ctx, []attr.Value{types.StringValue(testIntegrationID)})

	tests := []struct {
		name         string
		appliesToAll types.Bool
		configIDs    customfield.List[types.String]
		expected     bool
	}{
		{name: "true with integration IDs", appliesToAll: types.BoolValue(true), configIDs: someIDs, expected: true},
		{
			name:         "true with empty integration IDs",
			appliesToAll: types.BoolValue(true),
			configIDs:    customfield.NewListMust[types.String](ctx, []attr.Value{}),
		},
		{name: "true with null integration IDs", appliesToAll: types.BoolValue(true), configIDs: customfield.NullList[types.String](ctx)},
		{name: "false with integration IDs", appliesToAll: types.BoolValue(false), configIDs: someIDs},
		{name: "unknown applies to all", appliesToAll: types.BoolUnknown(), configIDs: someIDs},
		{name: "unknown integration IDs", appliesToAll: types.BoolValue(true), configIDs: customfield.UnknownList[types.String](ctx)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expected, appliesToAllWithIDs(tt.appliesToAll, tt.configIDs))
		})
	}
}

func TestModifyIntegrationScopePlan(t *testing.T) {
	t.Parallel()

	ids := tftypes.NewValue(tftypes.List{ElementType: tftypes.String},
		[]tftypes.Value{tftypes.NewValue(tftypes.String, testIntegrationID)})
	nullIDs := tftypes.NewValue(tftypes.List{ElementType: tftypes.String}, nil)
	unknownIDs := tftypes.NewValue(tftypes.List{ElementType: tftypes.String}, tftypes.UnknownValue)

	tests := []struct {
		name      string
		state     map[string]tftypes.Value // nil means create
		config    map[string]tftypes.Value
		plan      map[string]tftypes.Value
		wantError string
		wantIDs   tftypes.Value
	}{
		{
			name:    "switching to all integrations drops preserved IDs",
			state:   scope(false, ids),
			config:  scope(true, nullIDs),
			plan:    scope(true, ids),
			wantIDs: unknownIDs,
		},
		{
			name:    "unchanged all integrations plan keeps null IDs",
			state:   scope(true, nullIDs),
			config:  scope(true, nullIDs),
			plan:    scope(true, nullIDs),
			wantIDs: nullIDs,
		},
		{
			name:    "scoped update keeps preserved IDs",
			state:   scope(false, ids),
			config:  scope(false, nullIDs),
			plan:    scope(false, ids),
			wantIDs: ids,
		},
		{
			name:    "create for all integrations",
			config:  scope(true, nullIDs),
			plan:    scope(true, unknownIDs),
			wantIDs: unknownIDs,
		},
		{
			name:      "create scoped without IDs",
			config:    scope(false, nullIDs),
			plan:      scope(false, unknownIDs),
			wantError: "Missing integration IDs",
		},
		{
			name:      "switching from all integrations to scoped without IDs",
			state:     scope(true, nullIDs),
			config:    scope(false, nullIDs),
			plan:      scope(false, unknownIDs),
			wantError: "Missing integration IDs",
		},
		{
			name:    "switching from all integrations to scoped with IDs",
			state:   scope(true, nullIDs),
			config:  scope(false, ids),
			plan:    scope(false, ids),
			wantIDs: ids,
		},
		{
			name:      "all integrations with configured IDs",
			state:     scope(false, ids),
			config:    scope(true, ids),
			plan:      scope(true, ids),
			wantError: "Unexpected integration IDs",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()
			s := ResourceSchema(ctx)
			typ := s.Type().TerraformType(ctx)

			state := tfsdk.State{Schema: s, Raw: tftypes.NewValue(typ, nil)}
			if tt.state != nil {
				state.Raw = objectValue(t, typ, tt.state)
			}
			req := resource.ModifyPlanRequest{
				State:  state,
				Config: tfsdk.Config{Schema: s, Raw: objectValue(t, typ, tt.config)},
				Plan:   tfsdk.Plan{Schema: s, Raw: objectValue(t, typ, tt.plan)},
			}
			resp := &resource.ModifyPlanResponse{Plan: req.Plan}

			modifyIntegrationScopePlan(ctx, req, resp)

			if tt.wantError != "" {
				require.True(t, resp.Diagnostics.HasError())
				assert.Equal(t, tt.wantError, resp.Diagnostics.Errors()[0].Summary())
				return
			}
			require.False(t, resp.Diagnostics.HasError(), "%v", resp.Diagnostics)

			raw, _, err := tftypes.WalkAttributePath(resp.Plan.Raw, tftypes.NewAttributePath().WithAttributeName("integration_ids"))
			require.NoError(t, err)
			got, ok := raw.(tftypes.Value)
			require.True(t, ok)
			assert.True(t, tt.wantIDs.Equal(got), "integration_ids = %s, want %s", got, tt.wantIDs)
		})
	}
}

func scope(appliesToAll bool, integrationIDs tftypes.Value) map[string]tftypes.Value {
	return map[string]tftypes.Value{
		"applies_to_all_integrations": tftypes.NewValue(tftypes.Bool, appliesToAll),
		"integration_ids":             integrationIDs,
	}
}

// objectValue builds a resource object with the given attributes and every other attribute null.
func objectValue(t *testing.T, typ tftypes.Type, attrs map[string]tftypes.Value) tftypes.Value {
	t.Helper()
	objectType, ok := typ.(tftypes.Object)
	require.True(t, ok, "resource schema type is %T, want tftypes.Object", typ)
	values := make(map[string]tftypes.Value, len(objectType.AttributeTypes))
	for name, attrType := range objectType.AttributeTypes {
		if v, ok := attrs[name]; ok {
			values[name] = v
			continue
		}
		values[name] = tftypes.NewValue(attrType, nil)
	}
	return tftypes.NewValue(objectType, values)
}
