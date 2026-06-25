package load_balancer_pool

import (
	"sort"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
)

// The Load Balancing API canonically returns a pool's origins sorted by name
// (ORDER BY name), regardless of the order they were submitted in. The Terraform
// resource models `origins` as an ordered ListNestedAttribute (Stainless cannot
// emit a SetNestedAttribute here because the nested object contains purely
// computed children such as `disabled_at` and `port`, whose values are unknown
// at plan time and therefore cannot participate in set element identity).
//
// As a result, a config that lists origins in any order other than ascending by
// name produces a perpetual diff on every plan/apply (the planned order never
// matches the server's alphabetical order). See LB-5712 / GitHub #7179 (and the
// earlier, incompletely-fixed #6140 / LB-5179).
//
// UnmarshalCustom / UnmarshalComputedCustom wrap the standard apijson unmarshal
// and reorder the origins returned by the API back into the order the user
// planned (captured from the envelope before unmarshal), matching by origin
// identity (name, tie-broken by address). When there is no prior order to
// align to (e.g. import), origins fall back to a canonical sort by name, which
// matches the server's own ordering and is therefore stable.

// UnmarshalCustom unmarshals an API response into env, then reorders origins to
// match the order already present in env.Result.Origins (the prior state/plan).
func UnmarshalCustom(data []byte, env *LoadBalancerPoolResultEnvelope) error {
	snap := snapshotOriginOrder(env.Result.Origins)
	if err := apijson.Unmarshal(data, env); err != nil {
		return err
	}
	reorderOriginsFromSnapshot(env.Result.Origins, snap)
	return nil
}

// UnmarshalComputedCustom is like UnmarshalCustom but uses apijson.UnmarshalComputed,
// for the create/update paths where computed fields are populated from the response.
func UnmarshalComputedCustom(data []byte, env *LoadBalancerPoolResultEnvelope) error {
	snap := snapshotOriginOrder(env.Result.Origins)
	if err := apijson.UnmarshalComputed(data, env); err != nil {
		return err
	}
	reorderOriginsFromSnapshot(env.Result.Origins, snap)
	return nil
}

// originOrderSnapshot captures the prior ordering of origins as plain-string
// identity fingerprints.
type originOrderSnapshot struct {
	fingerprints []string
}

// originFingerprint builds a stable identity key for an origin: name + address.
// Both are user-supplied and, in combination, unique within a pool.
func originFingerprint(name, address string) string {
	return name + "|" + address
}

func originFingerprintFromModel(o *LoadBalancerPoolOriginsModel) string {
	if o == nil {
		return "|"
	}
	return originFingerprint(o.Name.ValueString(), o.Address.ValueString())
}

// snapshotOriginOrder records the current origin order. Returns nil when there
// are no prior origins to align to.
func snapshotOriginOrder(origins *[]*LoadBalancerPoolOriginsModel) *originOrderSnapshot {
	if origins == nil || len(*origins) == 0 {
		return nil
	}
	snap := &originOrderSnapshot{
		fingerprints: make([]string, len(*origins)),
	}
	for i, o := range *origins {
		snap.fingerprints[i] = originFingerprintFromModel(o)
	}
	return snap
}

// reorderOriginsFromSnapshot reorders origins to match the prior snapshot. When
// snap is nil (no prior order), it applies a canonical sort by name (matching
// the API's own ORDER BY name) so the result is deterministic.
func reorderOriginsFromSnapshot(origins *[]*LoadBalancerPoolOriginsModel, snap *originOrderSnapshot) {
	if origins == nil || len(*origins) == 0 {
		return
	}

	if snap == nil {
		sortOriginsByName(origins)
		return
	}

	// fingerprint -> desired position from prior order
	priorIndex := make(map[string]int, len(snap.fingerprints))
	for i, fp := range snap.fingerprints {
		// First occurrence wins; duplicate identities are unexpected within a pool.
		if _, ok := priorIndex[fp]; !ok {
			priorIndex[fp] = i
		}
	}

	os := *origins
	sort.SliceStable(os, func(i, j int) bool {
		fpI := originFingerprintFromModel(os[i])
		fpJ := originFingerprintFromModel(os[j])
		posI, okI := priorIndex[fpI]
		posJ, okJ := priorIndex[fpJ]

		switch {
		case okI && okJ:
			return posI < posJ
		case okI:
			return true // matched (known) origins before unmatched
		case okJ:
			return false
		default:
			// Both unmatched: canonical order by name.
			return originLessByName(os[i], os[j])
		}
	})
}

// sortOriginsByName applies a canonical, deterministic sort by name (tie-broken
// by address) matching the server's ORDER BY name.
func sortOriginsByName(origins *[]*LoadBalancerPoolOriginsModel) {
	os := *origins
	sort.SliceStable(os, func(i, j int) bool {
		return originLessByName(os[i], os[j])
	})
}

func originLessByName(a, b *LoadBalancerPoolOriginsModel) bool {
	na, nb := "", ""
	if a != nil {
		na = a.Name.ValueString()
	}
	if b != nil {
		nb = b.Name.ValueString()
	}
	if na != nb {
		return na < nb
	}
	aa, ab := "", ""
	if a != nil {
		aa = a.Address.ValueString()
	}
	if b != nil {
		ab = b.Address.ValueString()
	}
	return aa < ab
}
