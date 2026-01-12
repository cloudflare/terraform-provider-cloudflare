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
		_ = env.Result.ManagedBy.As(context.Background(), &managedBy, basetypes.ObjectAsOptions{})

		// Check if the user has 'unit' block in their HCL
		isUnitInConfig := configuredModel.Unit != nil
		// Identify an Import: the HCL is empty. Since 'name' is Required, if it's Null in the config but we have an ID in the result, it's an import.
		isImport := configuredModel.Name.IsNull() && !env.Result.ID.IsNull()
		// Identify an existing resource (Read/Update): We can check if the ID is already present in "Result"
		isExisting := !env.Result.ID.IsNull() && !env.Result.ID.IsUnknown()

		if (isImport || (isExisting && isUnitInConfig)) && !managedBy.ParentOrgID.IsNull() {
			env.Result.Unit = &AccountUnitModel{
				ID: managedBy.ParentOrgID,
			}
		} else {
			// During a fresh 'Create', isExisting is false or we fall through here.
			env.Result.Unit = nil
		}
	}

	return &env.Result, nil
}
