// This file is not generated. It is maintained by hand.

package turnstile_widget

import (
	"sort"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// The Turnstile API returns a widget's domains sorted alphabetically,
// regardless of the order they were submitted in. The Terraform resource
// models `domains` as an ordered ListAttribute (not a SetAttribute), so
// order matters for diff detection.
//
// As a result, a config that lists domains in any order other than ascending
// alphabetical produces a perpetual diff on every plan/apply (the planned
// order never matches the server's alphabetical order). See GitHub #7028.
//
// UnmarshalCustom / UnmarshalComputedCustom wrap the standard apijson
// unmarshal and reorder the domains returned by the API back into the order
// the user planned (captured from the envelope before unmarshal). When there
// is no prior order to align to (e.g. import), domains fall back to a
// canonical alphabetical sort, which matches the server's own ordering and
// is therefore stable.

// UnmarshalCustom unmarshals an API response into env, then reorders domains
// to match the order already present in env.Result.Domains (the prior
// state/plan).
func UnmarshalCustom(data []byte, env *TurnstileWidgetResultEnvelope) error {
	snap := snapshotDomainOrder(env.Result.Domains)
	if err := apijson.Unmarshal(data, env); err != nil {
		return err
	}
	reorderDomainsFromSnapshot(env.Result.Domains, snap)
	return nil
}

// UnmarshalComputedCustom is like UnmarshalCustom but uses
// apijson.UnmarshalComputed, for the create/update paths where computed
// fields are populated from the response.
//
// Although domains is a required (non-computed) field, Create/Update still
// need reordering because the API response overwrites the plan value during
// unmarshal. Without reordering, the state written after Create would contain
// alphabetically-sorted domains, causing a diff on the next plan when
// compared against the user's non-alphabetical config.
func UnmarshalComputedCustom(data []byte, env *TurnstileWidgetResultEnvelope) error {
	snap := snapshotDomainOrder(env.Result.Domains)
	if err := apijson.UnmarshalComputed(data, env); err != nil {
		return err
	}
	reorderDomainsFromSnapshot(env.Result.Domains, snap)
	return nil
}

// domainOrderSnapshot captures the prior ordering of domains as plain strings.
type domainOrderSnapshot struct {
	values []string
}

// snapshotDomainOrder records the current domain order. Returns nil when
// there are no prior domains to align to.
func snapshotDomainOrder(domains *[]types.String) *domainOrderSnapshot {
	if domains == nil || len(*domains) == 0 {
		return nil
	}
	snap := &domainOrderSnapshot{
		values: make([]string, len(*domains)),
	}
	for i, d := range *domains {
		snap.values[i] = d.ValueString()
	}
	return snap
}

// reorderDomainsFromSnapshot reorders domains to match the prior snapshot.
// When snap is nil (no prior order), it applies a canonical alphabetical
// sort matching the server's own ordering so the result is deterministic.
func reorderDomainsFromSnapshot(domains *[]types.String, snap *domainOrderSnapshot) {
	if domains == nil || len(*domains) == 0 {
		return
	}

	if snap == nil {
		sortDomainsAlphabetically(domains)
		return
	}

	// domain string -> desired position from prior order
	priorIndex := make(map[string]int, len(snap.values))
	for i, v := range snap.values {
		// First occurrence wins; duplicate domains are unexpected.
		if _, ok := priorIndex[v]; !ok {
			priorIndex[v] = i
		}
	}

	ds := *domains
	sort.SliceStable(ds, func(i, j int) bool {
		vi := ds[i].ValueString()
		vj := ds[j].ValueString()
		posI, okI := priorIndex[vi]
		posJ, okJ := priorIndex[vj]

		switch {
		case okI && okJ:
			return posI < posJ
		case okI:
			return true // matched (known) domains before unmatched
		case okJ:
			return false
		default:
			// Both unmatched: canonical alphabetical order.
			return vi < vj
		}
	})
}

// sortDomainsAlphabetically applies a canonical, deterministic alphabetical
// sort matching the server's own ordering.
func sortDomainsAlphabetically(domains *[]types.String) {
	ds := *domains
	sort.SliceStable(ds, func(i, j int) bool {
		return ds[i].ValueString() < ds[j].ValueString()
	})
}
