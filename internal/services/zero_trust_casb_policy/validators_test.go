package zero_trust_casb_policy

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
)

func TestConfigValidators(t *testing.T) {
	t.Parallel()

	validators := (&ZeroTrustCasbPolicyResource{}).ConfigValidators(context.Background())
	assert.Len(t, validators, 1)
	assert.IsType(t, integrationScopeValidator{}, validators[0])
}

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
				types.StringValue("019e702a-f9ab-7a3a-8790-2c1e434acf47"),
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
