package account

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func (m *AccountModel) marshalCustom() (data []byte, err error) {
	// Removing type data that is no longer accepted by the API
	savedType := m.Type
	m.Type = types.StringNull()
	bytes, err := apijson.MarshalRoot(m)
	// Adding it back
	m.Type = savedType
	return bytes, err
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

func unmarshalCustom(bytes []byte, configuredModel *AccountModel) (*AccountModel, error) {
	env := AccountResultEnvelope{}
	err := apijson.Unmarshal(bytes, &env)
	if err != nil {
		return nil, err
	}
	// Setting type to whatever the configured type is to avoid state/drift issues
	env.Result.Type = configuredModel.Type

	if !env.Result.ManagedBy.IsNull() && !env.Result.ManagedBy.IsUnknown() {
		var managedBy AccountManagedByModel
		diag := env.Result.ManagedBy.As(context.Background(), &managedBy, basetypes.ObjectAsOptions{})

		if !diag.HasError() {
			// Backfill unit from managed_by if user expects it OR if we are in discovery/import mode
			// This allows drift detection while preventing 'Inconsistent Result' panics during Apply
			if env.Result.Unit == nil || env.Result.Unit.ID.IsNull() || env.Result.Unit.ID.IsUnknown() {
				unitInConfig := configuredModel.Unit != nil && !configuredModel.Unit.ID.IsNull()
				if (unitInConfig) || configuredModel.ID.IsNull() {
					env.Result.Unit = &AccountUnitModel{
						ID: managedBy.ParentOrgID,
					}
				}
			}
		}
	}

	return &env.Result, nil
}
