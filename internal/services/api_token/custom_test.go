package api_token

import (
	"encoding/json"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestMarshalCustom(t *testing.T) {
	permGroups := []*APITokenPoliciesPermissionGroupsModel{
		{
			ID: types.StringValue("permgroup1"),
		},
	}
	policies := []*APITokenPoliciesModel{
		{
			Effect:           types.StringValue("effect"),
			PermissionGroups: &permGroups,
			Resources:        types.StringValue("{\"com.cloudflare.api.account.123\": \"*\"}"),
		},
	}
	tt := APITokenModel{
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
	permGroups := []*APITokenPoliciesPermissionGroupsModel{
		{
			ID: types.StringValue("permgroup1"),
		},
	}
	policies := []*APITokenPoliciesModel{
		{
			Effect:           types.StringValue("effect"),
			PermissionGroups: &permGroups,
			Resources:        types.StringValue("{\"com.cloudflare.api.account.123\": \"*\""),
		},
	}
	tt := APITokenModel{
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
	env := APITokenResultEnvelope{}
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
	permGroups1 := []*APITokenPoliciesPermissionGroupsModel{
		{
			ID: types.StringValue("permgroup1"),
		},
	}
	// Prior model has policies in [nested_resources, flat] order
	policies := []*APITokenPoliciesModel{
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
	env := APITokenResultEnvelope{
		Result: APITokenModel{
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
	pgs := []*APITokenPoliciesPermissionGroupsModel{
		{ID: types.StringValue("ccc")},
		{ID: types.StringValue("aaa")},
		{ID: types.StringValue("bbb")},
	}
	policies := []*APITokenPoliciesModel{
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
	pg := []*APITokenPoliciesPermissionGroupsModel{{ID: types.StringValue("pg1")}}
	policies := []*APITokenPoliciesModel{
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

// TestUnmarshalComputedCustom_GitHub7125 reproduces the bug from
// https://github.com/cloudflare/terraform-provider-cloudflare/issues/7125.
//
// On update, the planned model contains the user's two policies (each with a
// distinct set of permission_groups and distinct resources). The user changed
// one permission_group ID inside policy A. The API stored the new policies and
// returned them in swapped order: [B, A_new].
//
// Because Policies is tagged required (not computed), apijson.UnmarshalComputed
// leaves the planned policies in place. The custom unmarshal then walks the API
// response by index and overwrites each model policy's Resources using
// policyJsons[i]. When the API reorders policies relative to what was sent,
// this assigns the wrong resources to each policy: policy A ends up with
// policy B's resources and vice-versa. The subsequent reorder step then sees
// policies whose fingerprints no longer match the snapshot, so they fall
// through to the canonical sort and the swap is persisted to state.
//
// A correct implementation must match API policies to model policies by
// identity (effect + permission_group set), not by positional index.
func TestUnmarshalComputedCustom_GitHub7125(t *testing.T) {
	const resA = `{"com.cloudflare.api.account.acct123":"*"}`
	const resB = `{"com.cloudflare.api.account.acct123":{"com.cloudflare.api.account.zone.*":"*"}}`

	// Planned model: policy A (R2-ish PGs, flat resources) then policy B
	// (DNS/Zone PGs, nested resources). One PG ID in A was just changed by
	// the user (pgA3-new replaces the previous pgA3).
	pgsA := []*APITokenPoliciesPermissionGroupsModel{
		{ID: types.StringValue("pgA1")},
		{ID: types.StringValue("pgA2")},
		{ID: types.StringValue("pgA3-new")},
		{ID: types.StringValue("pgA4")},
		{ID: types.StringValue("pgA5")},
		{ID: types.StringValue("pgA6")},
	}
	pgsB := []*APITokenPoliciesPermissionGroupsModel{
		{ID: types.StringValue("pgB1")},
		{ID: types.StringValue("pgB2")},
		{ID: types.StringValue("pgB3")},
		{ID: types.StringValue("pgB4")},
	}
	policies := []*APITokenPoliciesModel{
		{
			Effect:           types.StringValue("allow"),
			PermissionGroups: &pgsA,
			Resources:        types.StringValue(resA),
		},
		{
			Effect:           types.StringValue("allow"),
			PermissionGroups: &pgsB,
			Resources:        types.StringValue(resB),
		},
	}
	env := APITokenResultEnvelope{
		Result: APITokenModel{
			Name:     types.StringValue("name"),
			Policies: &policies,
		},
	}

	// API response: same two policies as the plan, but in swapped order
	// ([B, A_new] instead of [A_new, B]).
	data := json.RawMessage(`{
		"result":{
			"name":"name",
			"policies":[
				{
					"effect":"allow",
					"permission_groups":[
						{"id":"pgB1"},
						{"id":"pgB2"},
						{"id":"pgB3"},
						{"id":"pgB4"}
					],
					"resources":` + resB + `
				},
				{
					"effect":"allow",
					"permission_groups":[
						{"id":"pgA1"},
						{"id":"pgA2"},
						{"id":"pgA3-new"},
						{"id":"pgA4"},
						{"id":"pgA5"},
						{"id":"pgA6"}
					],
					"resources":` + resA + `
				}
			]
		}
	}`)

	if err := UnmarshalComputedCustom(data, &env); err != nil {
		t.Fatal(err)
	}

	// Each policy in the final model must carry the resources that belong to
	// its own permission_group set, regardless of the order the API returned
	// the policies in. Verify by matching on a stable PG ID inside each set.
	got := *env.Result.Policies
	if len(got) != 2 {
		t.Fatalf("expected 2 policies, got %d", len(got))
	}

	pgsContain := func(p *APITokenPoliciesModel, id string) bool {
		if p.PermissionGroups == nil {
			return false
		}
		for _, pg := range *p.PermissionGroups {
			if pg.ID.ValueString() == id {
				return true
			}
		}
		return false
	}

	for i, p := range got {
		switch {
		case pgsContain(p, "pgA1"):
			// This is policy A — must have policy A's flat resources and
			// must carry the new (post-edit) PG ID.
			if p.Resources.ValueString() != resA {
				t.Fatalf("policies[%d] is policy A but has wrong resources:\n  want: %s\n  got:  %s",
					i, resA, p.Resources.ValueString())
			}
			if !pgsContain(p, "pgA3-new") {
				t.Fatalf("policies[%d] is policy A but does not contain the updated permission_group pgA3-new", i)
			}
		case pgsContain(p, "pgB1"):
			// This is policy B — must have policy B's nested resources.
			if p.Resources.ValueString() != resB {
				t.Fatalf("policies[%d] is policy B but has wrong resources:\n  want: %s\n  got:  %s",
					i, resB, p.Resources.ValueString())
			}
		default:
			t.Fatalf("policies[%d] has unrecognized permission_groups", i)
		}
	}

	// And the order in state should still match the planned order [A, B],
	// so that the next plan is a no-op.
	if !pgsContain(got[0], "pgA1") {
		t.Fatalf("expected policies[0] to be policy A (matching planned order), got policy with pgs %+v",
			*got[0].PermissionGroups)
	}
	if !pgsContain(got[1], "pgB1") {
		t.Fatalf("expected policies[1] to be policy B (matching planned order), got policy with pgs %+v",
			*got[1].PermissionGroups)
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
	env := APITokenResultEnvelope{}
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
