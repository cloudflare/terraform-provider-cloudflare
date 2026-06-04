package account_token

import (
	"encoding/json"
	"sort"
	"strings"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func MarshalCustom(m AccountTokenModel) (data []byte, err error) {
	if data, err = apijson.MarshalRoot(m); err != nil {
		return
	}
	var base map[string]json.RawMessage
	if err := json.Unmarshal(data, &base); err != nil {
		return nil, err
	}

	// for each policy, marshal the resources string as raw json
	policyJsons := make([]json.RawMessage, len(*m.Policies))
	for i, policy := range *m.Policies {
		policyData, err := apijson.MarshalRoot(policy)
		if err != nil {
			return nil, err
		}
		var policyBase map[string]json.RawMessage
		if err := json.Unmarshal(policyData, &policyBase); err != nil {
			return nil, err
		}
		resources := json.RawMessage(policy.Resources.ValueString())
		policyBase["resources"] = resources
		policyJsons[i], err = json.Marshal(policyBase)
		if err != nil {
			return nil, err
		}
	}
	base["policies"], err = json.Marshal(policyJsons)
	if err != nil {
		return nil, err
	}

	return json.Marshal(base)
}

func UnmarshalCustom(data []byte, model *AccountTokenResultEnvelope) (err error) {
	// Snapshot prior policy order before apijson.Unmarshal overwrites the model
	snap := snapshotPolicyOrder(model.Result.Policies)

	if err = apijson.Unmarshal(data, model); err != nil {
		return
	}

	return unmarshalCustomWithSnapshot(data, model, snap)
}

func UnmarshalComputedCustom(data []byte, model *AccountTokenResultEnvelope) (err error) {
	// Snapshot prior policy order BEFORE UnmarshalComputed overwrites the model
	snap := snapshotPolicyOrder(model.Result.Policies)

	if err = apijson.UnmarshalComputed(data, model); err != nil {
		return
	}
	return unmarshalCustomWithSnapshot(data, model, snap)
}

// unmarshalCustomWithSnapshot is like UnmarshalCustom but uses a provided
// snapshot instead of capturing its own. apijson.Unmarshal is NOT called
// because the caller has already done the initial unmarshal.
//
// Note: when this is reached via UnmarshalComputedCustom (update path),
// apijson.UnmarshalComputed leaves the planned Policies slice in place
// because the field is tagged "required", not "computed". That means the
// model still carries the user's planned effect + permission_groups for
// each policy, and only the Resources string still needs to be populated
// from the API response. The API may reorder policies relative to what
// was sent, so we must match by policy identity (effect + permission group
// set), NOT by positional index.
func unmarshalCustomWithSnapshot(data []byte, model *AccountTokenResultEnvelope, snap *policyOrderSnapshot) error {
	// pull out the raw JSON values for each policy resource and map to the model
	var base map[string]json.RawMessage
	if err := json.Unmarshal(data, &base); err != nil {
		return err
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(base["result"], &result); err != nil {
		return err
	}
	var policyJsons []json.RawMessage
	if err := json.Unmarshal(result["policies"], &policyJsons); err != nil {
		return err
	}

	// Parse each API policy into an identity (effect + permission group IDs)
	// and a resources string so we can match by identity instead of by index.
	type apiPolicy struct {
		effect    string
		pgIDs     []string
		resources string
	}
	apiPolicies := make([]apiPolicy, 0, len(policyJsons))
	for _, pj := range policyJsons {
		var p struct {
			Effect           string `json:"effect"`
			PermissionGroups []struct {
				ID string `json:"id"`
			} `json:"permission_groups"`
			Resources json.RawMessage `json:"resources"`
		}
		if err := json.Unmarshal(pj, &p); err != nil {
			return err
		}
		ids := make([]string, len(p.PermissionGroups))
		for i, pg := range p.PermissionGroups {
			ids[i] = pg.ID
		}
		apiPolicies = append(apiPolicies, apiPolicy{
			effect:    p.Effect,
			pgIDs:     ids,
			resources: string(p.Resources),
		})
	}

	// identityKey is intentionally narrower than policyFingerprint: it omits
	// resources because that's the field we're trying to assign. Sorting the
	// permission group IDs makes the key invariant to PG order changes within
	// a policy (those are handled later by reorderPoliciesFromSnapshot).
	identityKey := func(effect string, pgIDs []string) string {
		sorted := make([]string, len(pgIDs))
		copy(sorted, pgIDs)
		sort.Strings(sorted)
		return effect + "|" + strings.Join(sorted, ",")
	}

	// Index API policies by identity. Use a slice of indices per key so that
	// duplicate identities (multiple policies with the same effect + PG set
	// but different resources) are consumed in encounter order.
	available := make(map[string][]int, len(apiPolicies))
	for i, ap := range apiPolicies {
		k := identityKey(ap.effect, ap.pgIDs)
		available[k] = append(available[k], i)
	}

	// For each model policy, find the matching API policy by identity and
	// copy its resources string into the model. If no match is found (which
	// shouldn't happen for a healthy round-trip), fall back to positional
	// assignment to preserve previous behavior.
	policies := *model.Result.Policies
	for i, p := range policies {
		var pgIDs []string
		if p.PermissionGroups != nil {
			pgIDs = make([]string, len(*p.PermissionGroups))
			for j, pg := range *p.PermissionGroups {
				pgIDs[j] = pg.ID.ValueString()
			}
		}
		k := identityKey(p.Effect.ValueString(), pgIDs)
		if idxs, ok := available[k]; ok && len(idxs) > 0 {
			apIdx := idxs[0]
			available[k] = idxs[1:]
			policies[i].Resources = types.StringValue(apiPolicies[apIdx].resources)
			continue
		}
		// Fallback: positional. Preserves prior (buggy) behavior only when
		// identity lookup fails, which is better than dropping the value.
		if i < len(apiPolicies) {
			policies[i].Resources = types.StringValue(apiPolicies[i].resources)
		}
	}

	// Reorder to match prior state (or canonical sort if no prior)
	reorderPoliciesFromSnapshot(model.Result.Policies, snap)
	return nil
}

// ============================================================================
// Policy order snapshot and reordering
// ============================================================================

// policySnapshot captures a policy's identity as plain strings.
type policySnapshot struct {
	effect           string
	resources        string
	permissionGroups []string
}

// policyOrderSnapshot captures the full prior ordering of policies.
type policyOrderSnapshot struct {
	policies []policySnapshot
}

// snapshotPolicyOrder extracts the current policy order as plain strings.
// Returns nil if there are no prior policies.
func snapshotPolicyOrder(policies *[]*AccountTokenPoliciesModel) *policyOrderSnapshot {
	if policies == nil || len(*policies) == 0 {
		return nil
	}
	snap := &policyOrderSnapshot{
		policies: make([]policySnapshot, len(*policies)),
	}
	for i, p := range *policies {
		var pgIDs []string
		if p.PermissionGroups != nil {
			pgIDs = make([]string, len(*p.PermissionGroups))
			for j, pg := range *p.PermissionGroups {
				pgIDs[j] = pg.ID.ValueString()
			}
		}
		snap.policies[i] = policySnapshot{
			effect:           p.Effect.ValueString(),
			resources:        p.Resources.ValueString(),
			permissionGroups: pgIDs,
		}
	}
	return snap
}

// policyFingerprint builds a unique key for a policy using effect + resources +
// sorted permission group IDs, joined with "|".
func policyFingerprint(effect, resources string, pgIDs []string) string {
	sorted := make([]string, len(pgIDs))
	copy(sorted, pgIDs)
	sort.Strings(sorted)
	return effect + "|" + resources + "|" + strings.Join(sorted, ",")
}

// policyFingerprintFromModel builds a fingerprint from a policy model.
func policyFingerprintFromModel(p *AccountTokenPoliciesModel) string {
	var pgIDs []string
	if p.PermissionGroups != nil {
		pgIDs = make([]string, len(*p.PermissionGroups))
		for i, pg := range *p.PermissionGroups {
			pgIDs[i] = pg.ID.ValueString()
		}
	}
	return policyFingerprint(p.Effect.ValueString(), p.Resources.ValueString(), pgIDs)
}

// reorderPoliciesFromSnapshot reorders policies to match the prior snapshot.
// If snap is nil (no prior state), falls back to canonical sort.
func reorderPoliciesFromSnapshot(newPolicies *[]*AccountTokenPoliciesModel, snap *policyOrderSnapshot) {
	if newPolicies == nil || len(*newPolicies) == 0 {
		return
	}

	if snap == nil {
		sortPolicies(newPolicies)
		return
	}

	// Build index: fingerprint → desired position from prior snapshot
	priorIndex := make(map[string]int, len(snap.policies))
	for i, sp := range snap.policies {
		fp := policyFingerprint(sp.effect, sp.resources, sp.permissionGroups)
		priorIndex[fp] = i
	}

	policies := *newPolicies

	// Stable sort: policies that match prior go to their prior position;
	// unmatched policies sort after matched ones in canonical order.
	sort.SliceStable(policies, func(i, j int) bool {
		fpI := policyFingerprintFromModel(policies[i])
		fpJ := policyFingerprintFromModel(policies[j])
		posI, okI := priorIndex[fpI]
		posJ, okJ := priorIndex[fpJ]

		switch {
		case okI && okJ:
			return posI < posJ
		case okI:
			return true // matched before unmatched
		case okJ:
			return false
		default:
			// Both unmatched: canonical order
			return canonicalPolicyLess(policies[i], policies[j])
		}
	})

	// Reorder permission_groups within each policy to match prior
	priorPGIndex := make(map[string]map[string]int, len(snap.policies))
	for _, sp := range snap.policies {
		fp := policyFingerprint(sp.effect, sp.resources, sp.permissionGroups)
		pgIdx := make(map[string]int, len(sp.permissionGroups))
		for j, pgID := range sp.permissionGroups {
			pgIdx[pgID] = j
		}
		priorPGIndex[fp] = pgIdx
	}

	for _, p := range policies {
		if p.PermissionGroups == nil || len(*p.PermissionGroups) == 0 {
			continue
		}
		fp := policyFingerprintFromModel(p)
		pgIdx, ok := priorPGIndex[fp]
		if !ok {
			// No prior — canonical sort by ID
			pgs := *p.PermissionGroups
			sort.SliceStable(pgs, func(i, j int) bool {
				return pgs[i].ID.ValueString() < pgs[j].ID.ValueString()
			})
			continue
		}

		pgs := *p.PermissionGroups
		sort.SliceStable(pgs, func(i, j int) bool {
			idI := pgs[i].ID.ValueString()
			idJ := pgs[j].ID.ValueString()
			posI, okI := pgIdx[idI]
			posJ, okJ := pgIdx[idJ]
			switch {
			case okI && okJ:
				return posI < posJ
			case okI:
				return true
			case okJ:
				return false
			default:
				return idI < idJ
			}
		})
	}
}

// sortPolicies applies canonical sort: permission_groups by ID within each
// policy, then policies by effect then resources.
func sortPolicies(policies *[]*AccountTokenPoliciesModel) {
	if policies == nil || len(*policies) == 0 {
		return
	}
	ps := *policies

	// Sort permission groups within each policy
	for _, p := range ps {
		if p.PermissionGroups == nil || len(*p.PermissionGroups) == 0 {
			continue
		}
		pgs := *p.PermissionGroups
		sort.SliceStable(pgs, func(i, j int) bool {
			return pgs[i].ID.ValueString() < pgs[j].ID.ValueString()
		})
	}

	// Sort policies by effect, then resources
	sort.SliceStable(ps, func(i, j int) bool {
		return canonicalPolicyLess(ps[i], ps[j])
	})
}

// canonicalPolicyLess defines the canonical ordering: by effect, then resources.
func canonicalPolicyLess(a, b *AccountTokenPoliciesModel) bool {
	ea := a.Effect.ValueString()
	eb := b.Effect.ValueString()
	if ea != eb {
		return ea < eb
	}
	return a.Resources.ValueString() < b.Resources.ValueString()
}
