// This file contains hand-written logic that extends the generated resource.go.
// It is intentionally separate so generator regenerations leave it untouched.
//
// Hooks into resource.go:
//   - Read()       calls fetchItems(ctx, client, accountID, listID) to populate data.Items
//   - ModifyPlan() calls suppressItemsDiff(ctx, req, resp) for hash-based diff suppression

package zero_trust_list

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"sort"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// fetchItems retrieves all items for a gateway list via the dedicated items
// endpoint (GET /gateway/lists/{id}/items), which is separate from the list
// metadata endpoint that Read() uses.
//
// Return values:
//   - (items, nil)  — success; items is always non-nil (empty slice = list exists with 0 items)
//   - (nil, nil)    — list was deleted (404); caller should remove resource from state
//   - (nil, err)    — unexpected error
func fetchItems(ctx context.Context, client *cloudflare.Client, accountID, listID string) ([]*ZeroTrustListItemsModel, error) {
	page, err := client.ZeroTrust.Gateway.Lists.Items.List(
		ctx,
		listID,
		zero_trust.GatewayListItemListParams{
			AccountID: cloudflare.F(accountID),
		},
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		var apiErr *cloudflare.Error
		if errors.As(err, &apiErr) && apiErr.StatusCode == 404 {
			return nil, nil
		}
		return nil, err
	}

	// A null result field means the API returned no data — treat as gone.
	if page.JSON.Result.IsNull() {
		return nil, fmt.Errorf("items endpoint returned null result for list %s", listID)
	}

	// Result is [][]GatewayItem: the outer slice always has exactly one element
	// (SinglePage is never paginated), and the inner slice is the list of items.
	if len(page.Result) > 1 {
		return nil, fmt.Errorf("unexpected pagination: items endpoint returned %d result pages, expected at most 1", len(page.Result))
	}

	// Allocate a non-nil slice so the caller can distinguish "0 items" from "404".
	items := make([]*ZeroTrustListItemsModel, 0)
	if len(page.Result) == 1 {
		for _, item := range page.Result[0] {
			items = append(items, &ZeroTrustListItemsModel{
				Value:       types.StringValue(item.Value),
				Description: gatewayItemDescription(item),
			})
		}
	}
	return items, nil
}

// gatewayItemDescription maps a GatewayItem description to a Terraform string.
// Returns StringNull() when the field is absent or null in the API response,
// so state matches configs that omit the description field entirely.
func gatewayItemDescription(item zero_trust.GatewayItem) types.String {
	// IsNull() returns true for both null and missing fields (status <= null),
	// so no separate IsMissing() check is needed.
	if item.JSON.Description.IsNull() {
		return types.StringNull()
	}
	return types.StringValue(item.Description)
}

// suppressItemsDiff suppresses spurious plan diffs for list items when the
// config and state are semantically identical (same values and descriptions,
// regardless of set ordering).
//
// Background: Terraform's SetNestedAttribute comparison is O(n) but each
// element requires expensive nested object hashing (~70s for 2000 items).
// This hash-based approach takes ~0.5ms for 5000 items (~70,000x speedup).
func suppressItemsDiff(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// Plan is null on destroy; state is null on first create. !IsKnown() means
	// the resource itself is not yet known (e.g. ID computed from another resource).
	if req.State.Raw.IsNull() || req.Plan.Raw.IsNull() ||
		!req.State.Raw.IsKnown() || !req.Plan.Raw.IsKnown() {
		return
	}

	var plan, state *ZeroTrustListModel

	// Use local diagnostics so deserialization failures in this best-effort
	// hook do not propagate as provider errors in the plan response.
	var planDiags, stateDiags diag.Diagnostics
	planDiags.Append(req.Plan.Get(ctx, &plan)...)
	stateDiags.Append(req.State.Get(ctx, &state)...)
	if planDiags.HasError() || stateDiags.HasError() || plan == nil || state == nil {
		return
	}

	if plan.Items != nil && state.Items != nil {
		// If any plan item has an unknown value we cannot reliably compare —
		// bail out and let Terraform show the diff as-is.
		for _, item := range *plan.Items {
			if item != nil && (item.Value.IsUnknown() || item.Description.IsUnknown()) {
				return
			}
		}

		if computeItemsHash(*plan.Items) == computeItemsHash(*state.Items) {
			// Use SetAttribute to surgically replace only the items attribute,
			// avoiding clobbering unknown computed values in other plan fields.
			resp.Diagnostics.Append(resp.Plan.SetAttribute(ctx, path.Root("items"), state.Items)...)
		}
	}
}

// computeItemsHash computes a deterministic SHA-256 hash of list items for
// efficient set comparison. Items are sorted before hashing to ensure
// order-independence (set semantics).
//
// Each item is encoded with explicit type tags so that nil, null, unknown,
// and empty-string values are all distinguishable from one another.
func computeItemsHash(items []*ZeroTrustListItemsModel) [32]byte {
	encoded := make([]string, len(items))
	for i, item := range items {
		if item == nil {
			encoded[i] = "nil"
			continue
		}

		// Encode value with explicit null/unknown/string tags.
		var v string
		switch {
		case item.Value.IsUnknown():
			v = "v:unknown"
		case item.Value.IsNull():
			v = "v:null"
		default:
			raw := item.Value.ValueString()
			v = fmt.Sprintf("v(%d):%s", len(raw), raw)
		}

		// Encode description with explicit null/unknown/string tags.
		var d string
		switch {
		case item.Description.IsUnknown():
			d = "d:unknown"
		case item.Description.IsNull():
			d = "d:null"
		default:
			raw := item.Description.ValueString()
			d = fmt.Sprintf("d(%d):%s", len(raw), raw)
		}

		encoded[i] = v + "/" + d
	}
	sort.Strings(encoded)

	h := sha256.New()
	for _, v := range encoded {
		h.Write([]byte(v))
		h.Write([]byte{0})
	}
	var result [32]byte
	copy(result[:], h.Sum(nil))
	return result
}
