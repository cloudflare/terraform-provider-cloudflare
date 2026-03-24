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
	"io"
	"net/http"
	"sort"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
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

	// Result is [][]GatewayItem: the outer slice always has exactly one element
	// (SinglePage is never paginated), and the inner slice is the list of items.
	// Allocate a non-nil slice so the caller can distinguish "0 items" from "404".
	items := make([]*ZeroTrustListItemsModel, 0)
	if len(page.Result) > 0 {
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
// Returns StringNull() when the field is absent so state matches configs that
// omit the description field entirely, avoiding perpetual plan diffs.
func gatewayItemDescription(item zero_trust.GatewayItem) types.String {
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
	// Plan is null on destroy; state is null on first create.
	if req.State.Raw.IsNull() || req.Plan.Raw.IsNull() {
		return
	}

	var plan, state *ZeroTrustListModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() || plan == nil || state == nil {
		return
	}

	if plan.Items != nil && state.Items != nil {
		if computeItemsHash(*plan.Items) == computeItemsHash(*state.Items) {
			plan.Items = state.Items
			resp.Diagnostics.Append(resp.Plan.Set(ctx, plan)...)
		}
	}
}

// computeItemsHash computes a deterministic SHA-256 hash of list items for
// efficient set comparison. Items are sorted before hashing to ensure
// order-independence (set semantics). Each item is encoded as
// "<len(value)>:<value>/<len(desc)>:<desc>" using length-prefixed fields
// to avoid separator-collision ambiguities.
func computeItemsHash(items []*ZeroTrustListItemsModel) [32]byte {
	encoded := make([]string, len(items))
	for i, item := range items {
		if item != nil {
			v := item.Value.ValueString()
			d := item.Description.ValueString()
			encoded[i] = fmt.Sprintf("%d:%s/%d:%s", len(v), v, len(d), d)
		}
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

// readBody reads all bytes from an http.Response body.
func readBody(res *http.Response) ([]byte, error) {
	defer res.Body.Close()
	return io.ReadAll(res.Body)
}
