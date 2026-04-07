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
	// API returns policies in [flat, nested_resources] order
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
	// Prior model has policies in [nested_resources, flat] order
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
	// With reorder-to-match-prior, result should be [nested_resources, flat] matching prior order
	foundPolicies := env.Result.Policies
	if (*foundPolicies)[0].Resources != types.StringValue("{\"com.cloudflare.api.account.123\":{\"com.cloudflare.api.account.zone.*\":\"*\"}}") {
		t.Fatalf("expected nested resources first, got %s", (*foundPolicies)[0].Resources)
	}
	if (*foundPolicies)[1].Resources != types.StringValue("{\"com.cloudflare.api.account.123\":\"*\"}") {
		t.Fatalf("expected flat resources second, got %s", (*foundPolicies)[1].Resources)
	}
}

func TestSortPolicies_SortsPermissionGroupsByID(t *testing.T) {
	pgs := []*AccountTokenPoliciesPermissionGroupsModel{
		{ID: types.StringValue("ccc")},
		{ID: types.StringValue("aaa")},
		{ID: types.StringValue("bbb")},
	}
	policies := []*AccountTokenPoliciesModel{
		{
			Effect:           types.StringValue("allow"),
			PermissionGroups: &pgs,
			Resources:        types.StringValue("{}"),
		},
	}
	sortPolicies(&policies)

	sorted := *policies[0].PermissionGroups
	expected := []string{"aaa", "bbb", "ccc"}
	for i, exp := range expected {
		got := sorted[i].ID.ValueString()
		if got != exp {
			t.Fatalf("permission_groups[%d]: expected %q, got %q", i, exp, got)
		}
	}
}

func TestSortPolicies_SortsPoliciesByEffectThenResources(t *testing.T) {
	pg := []*AccountTokenPoliciesPermissionGroupsModel{{ID: types.StringValue("pg1")}}
	policies := []*AccountTokenPoliciesModel{
		{Effect: types.StringValue("deny"), PermissionGroups: &pg, Resources: types.StringValue("b")},
		{Effect: types.StringValue("allow"), PermissionGroups: &pg, Resources: types.StringValue("z")},
		{Effect: types.StringValue("allow"), PermissionGroups: &pg, Resources: types.StringValue("a")},
	}
	sortPolicies(&policies)

	type expected struct {
		effect    string
		resources string
	}
	expectations := []expected{
		{"allow", "a"},
		{"allow", "z"},
		{"deny", "b"},
	}
	for i, exp := range expectations {
		got := policies[i]
		if got.Effect.ValueString() != exp.effect || got.Resources.ValueString() != exp.resources {
			t.Fatalf("policies[%d]: expected {%s, %s}, got {%s, %s}",
				i, exp.effect, exp.resources, got.Effect.ValueString(), got.Resources.ValueString())
		}
	}
}

func TestUnmarshalCustom_SortsPermissionGroups(t *testing.T) {
	// API returns permission groups in [zzz, aaa, mmm] order, no prior state
	data := json.RawMessage(`{
		"result":{
			"name":"name",
			"policies":[
				{
					"effect":"allow",
					"permission_groups":[{"id":"zzz"},{"id":"aaa"},{"id":"mmm"}],
					"resources":{"com.cloudflare.api.account.123":"*"}
				}
			]
		}
	}`)
	env := AccountTokenResultEnvelope{}
	err := UnmarshalCustom(data, &env)
	if err != nil {
		t.Fatal(err)
	}
	// With no prior state, canonical sort should order by ID: [aaa, mmm, zzz]
	pgs := *(*env.Result.Policies)[0].PermissionGroups
	expected := []string{"aaa", "mmm", "zzz"}
	for i, exp := range expected {
		got := pgs[i].ID.ValueString()
		if got != exp {
			t.Fatalf("permission_groups[%d]: expected %q, got %q", i, exp, got)
		}
	}
}
