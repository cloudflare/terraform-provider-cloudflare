package account_member

import (
	"testing"
)

// The API returns one roles entry per policy that grants a role, so the same
// role id can appear more than once when it is scoped to several resources.
func TestUnmarshalComputedCustomDuplicateRoles(t *testing.T) {
	data := []byte(`{"result":{"id":"member-1","roles":[{"id":"role-1"},{"id":"role-1"}]}}`)

	result, err := unmarshalComputedCustom(data, &AccountMemberModel{})
	if err != nil {
		t.Fatalf("unexpected error parsing duplicate roles: %s", err)
	}

	if got := len(result.Roles.Elements()); got != 1 {
		t.Errorf("expected 1 unique role, got %d", got)
	}
}
