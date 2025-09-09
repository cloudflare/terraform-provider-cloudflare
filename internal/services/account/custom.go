package account

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
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
	return &env.Result, nil
}
