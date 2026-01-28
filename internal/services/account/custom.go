package account

import (
	"context"
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (m *AccountModel) marshalCustom() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m *AccountModel) marshalCustomForUpdate(state AccountModel) (data []byte, err error) {
	// Removing type data that is no longer accepted by the API
	savedType := m.Type
	m.Type = types.StringNull()
	bytes, err := m.MarshalJSONForUpdate(state)
	// Adding it back
	m.Type = savedType
	return bytes, err
}

func unmarshalCustom(ctx context.Context, bytes []byte, configuredModel *AccountModel) (*AccountModel, error) {
	var env AccountResultEnvelope
	if err := apijson.Unmarshal(bytes, &env); err != nil {
		return nil, err
	}
	result := &env.Result

	return parseUnitAndManagedBy(ctx, bytes, result)
}

func parseUnitAndManagedBy(ctx context.Context, bytes []byte, result *AccountModel) (*AccountModel, error) {
	// Manually extract parent_org_id from the response JSON
	var fullResponse struct {
		Result struct {
			ManagedBy struct {
				ParentOrgID types.String `json:"parent_org_id"`
			} `json:"managed_by"`
		} `json:"result"`
	}

	if err := apijson.Unmarshal(bytes, &fullResponse); err != nil {
		return nil, err
	}

	// Direct mapping: Whenever parent_org_id is not null, map it to unit.id
	if !fullResponse.Result.ManagedBy.ParentOrgID.IsNull() {
		unit := AccountUnitModel{
			ID: fullResponse.Result.ManagedBy.ParentOrgID,
		}

		unitAttrTypes := map[string]attr.Type{"id": types.StringType}
		unitObj, diags := types.ObjectValueFrom(ctx, unitAttrTypes, unit)
		if diags.HasError() {
			return nil, fmt.Errorf("failed to create unit object: %s", diags.Errors())
		}

		result.Unit.ObjectValue = unitObj
	}

	return result, nil
}
