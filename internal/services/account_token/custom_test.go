package account_token

import (
	"encoding/json"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestMarshalCustom(t *testing.T) {
	permGroups := []*AccountTokenPoliciesPermissionGroupsModel{
		{
			ID: types.StringValue("permgroup1"),
		},
	}
	policies := []*AccountTokenPoliciesModel{
		{
			Effect:           types.StringValue("effect"),
			PermissionGroups: &permGroups,
			Resources:        types.StringValue("{\"com.cloudflare.api.account.123\": \"*\"}"),
		},
	}
	tt := AccountTokenModel{
		ID:        types.StringValue("id"),
		Name:      types.StringValue("name"),
		Policies:  &policies,
		Condition: nil,
	}
	data, err := MarshalCustom(tt)
	if err != nil {
		t.Fatal(err)
	}
	expectedData := json.RawMessage("{\"name\":\"name\",\"policies\":[{\"effect\":\"effect\",\"permission_groups\":[{\"id\":\"permgroup1\"}],\"resources\":{\"com.cloudflare.api.account.123\":\"*\"}}]}")
	if string(data) != string(expectedData) {
		t.Fatalf("expected %s, got %s", string(expectedData), string(data))
	}
}

func TestMarshalCustom_BadJson(t *testing.T) {
	permGroups := []*AccountTokenPoliciesPermissionGroupsModel{
		{
			ID: types.StringValue("permgroup1"),
		},
	}
	policies := []*AccountTokenPoliciesModel{
		{
			Effect:           types.StringValue("effect"),
			PermissionGroups: &permGroups,
			Resources:        types.StringValue("{\"com.cloudflare.api.account.123\": \"*\""),
		},
	}
	tt := AccountTokenModel{
		ID:        types.StringValue("id"),
		Name:      types.StringValue("name"),
		Policies:  &policies,
		Condition: nil,
	}
	_, err := MarshalCustom(tt)
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestUnmarshalCustom(t *testing.T) {
	data := json.RawMessage(`{
		"result":{
			"name":"name",
			"policies":[
				{
					"effect":"effect",
					"permission_groups":[{"id":"permgroup1"}],
					"resources":{"com.cloudflare.api.account.123":"*"}
				},
				{
					"effect":"effect",
					"permission_groups":[{"id":"permgroup1"}],
					"resources":{"com.cloudflare.api.account.123":{"com.cloudflare.api.account.zone.*":"*"}}
				}
			]
		}
	}`)
	env := AccountTokenResultEnvelope{}
	err := UnmarshalCustom(data, &env)
	if err != nil {
		t.Fatal(err)
	}
	if env.Result.Name != types.StringValue("name") {
		t.Fatalf("expected %s, got %s", "name", env.Result.Name)
	}
	policies := env.Result.Policies
	if (*policies)[0].Resources != types.StringValue("{\"com.cloudflare.api.account.123\":\"*\"}") {
		t.Fatalf("expected %s, got %s", "{\"com.cloudflare.api.account.123\":\"*\"}", (*policies)[0].Resources)
	}
	if (*policies)[1].Resources != types.StringValue("{\"com.cloudflare.api.account.123\":{\"com.cloudflare.api.account.zone.*\":\"*\"}}") {
		t.Fatalf("expected %s, got %s", "{\"com.cloudflare.api.account.123\":{\"com.cloudflare.api.account.zone.*\":\"*\"}}", (*policies)[0].Resources)
	}
}

func TestUnmarshalComputedCustom(t *testing.T) {
	data := json.RawMessage(`{
		"result":{
			"name":"name",
			"policies":[
				{
					"effect":"effect",
					"permission_groups":[{"id":"permgroup1"}],
					"resources":{"com.cloudflare.api.account.123":"*"}
				},
				{
					"effect":"effect",
					"permission_groups":[{"id":"permgroup1"}],
					"resources":{"com.cloudflare.api.account.123":{"com.cloudflare.api.account.zone.*":"*"}}
				}
			]
		}
	}`)
	permGroups1 := []*AccountTokenPoliciesPermissionGroupsModel{
		{
			ID: types.StringValue("permgroup1"),
		},
	}
	// Policy order in client-side model is not the same as the server-side data
	// model.
	policies := []*AccountTokenPoliciesModel{
		{
			Effect:           types.StringValue("effect"),
			PermissionGroups: &permGroups1,
			Resources:        types.StringValue("{\"com.cloudflare.api.account.123\":{\"com.cloudflare.api.account.zone.*\":\"*\"}}"),
		},
		{
			Effect:           types.StringValue("effect"),
			PermissionGroups: &permGroups1,
			Resources:        types.StringValue("{\"com.cloudflare.api.account.123\": \"*\"}"),
		},
	}
	env := AccountTokenResultEnvelope{
		Result: AccountTokenModel{
			Name:     types.StringValue("name"),
			Policies: &policies,
		},
	}
	err := UnmarshalComputedCustom(data, &env)
	if err != nil {
		t.Fatal(err)
	}
	// The order of policies should be the same as the server-side data model
	foundPolicies := env.Result.Policies
	if (*foundPolicies)[0].Resources != types.StringValue("{\"com.cloudflare.api.account.123\":\"*\"}") {
		t.Fatalf("expected %s, got %s", "{\"com.cloudflare.api.account.123\":\"*\"}", (*foundPolicies)[0].Resources)
	}
	if (*foundPolicies)[1].Resources != types.StringValue("{\"com.cloudflare.api.account.123\":{\"com.cloudflare.api.account.zone.*\":\"*\"}}") {
		t.Fatalf("expected %s, got %s", "{\"com.cloudflare.api.account.123\":{\"com.cloudflare.api.account.zone.*\":\"*\"}}", (*foundPolicies)[0].Resources)
	}
}
