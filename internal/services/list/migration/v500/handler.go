package v500

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeFromV0 handles state upgrades from schema_version=0 to v500.
// Both v4 (SDKv2) and v5 (Framework) state arrive with schema_version=0 because
// the published v5 provider has no Version field set (defaults to 0).
//
// Detection:
//   - v4 state: has Item blocks populated (len(Item) > 0)
//   - v5 state or empty: no Item blocks (len(Item) == 0)
//
// For v4 state, the full Transform() pipeline runs (item/value flattening,
// boolean conversion, IP normalization, etc.).
// For v5 state, items are converted from plain types to customfield types.
func UpgradeFromV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	var sourceState SourceListModel
	resp.Diagnostics.Append(req.State.Get(ctx, &sourceState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var targetState *TargetListModel
	var diags diag.Diagnostics

	if len(sourceState.Item) > 0 {
		// v4 state with item blocks - run full transform
		tflog.Info(ctx, "Upgrading cloudflare_list state from v4 (schema_version=0)")
		targetState, diags = Transform(ctx, sourceState)
	} else {
		// v5 state or state without items - pass through with type conversion
		tflog.Info(ctx, "Upgrading cloudflare_list state from v5 (schema_version=0)")
		targetState, diags = transformV5State(ctx, sourceState)
	}

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, targetState)...)
}

// UpgradeFromV1 handles state upgrades from v5 provider (schema_version=1) to v500.
// This is a no-op upgrade that just bumps the version.
// Some published v5 versions set Version: 1 before schema version was introduced.
func UpgradeFromV1(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading cloudflare_list state from version=1 to version=500 (no-op)")
	resp.State.Raw = req.State.Raw
	tflog.Info(ctx, "State version bump from 1 to 500 completed")
}

// transformV5State converts v5 state (or empty state) to the target model.
// The main work is converting items from plain types.Set to customfield.NestedObjectSet.
func transformV5State(ctx context.Context, source SourceListModel) (*TargetListModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	target := &TargetListModel{
		ID:                    source.ID,
		AccountID:             source.AccountID,
		Kind:                  source.Kind,
		Name:                  source.Name,
		Description:           source.Description,
		CreatedOn:             source.CreatedOn,
		ModifiedOn:            source.ModifiedOn,
		NumItems:              source.NumItems,
		NumReferencingFilters: source.NumReferencingFilters,
	}

	if source.Items.IsNull() {
		target.Items = customfield.NullObjectSet[TargetListItemModel](ctx)
		return target, diags
	}

	if source.Items.IsUnknown() {
		target.Items = customfield.NullObjectSet[TargetListItemModel](ctx)
		return target, diags
	}

	// Convert each item from types.Object to TargetListItemModel
	var targetItems []TargetListItemModel
	for _, elem := range source.Items.Elements() {
		obj, ok := elem.(basetypes.ObjectValue)
		if !ok {
			continue
		}
		attrs := obj.Attributes()

		targetItem := TargetListItemModel{
			ASN:     attrs["asn"].(types.Int64),
			Comment: attrs["comment"].(types.String),
			IP:      attrs["ip"].(types.String),
		}

		// Convert hostname: types.Object → customfield.NestedObject[TargetHostnameModel]
		hostnameObj := attrs["hostname"].(basetypes.ObjectValue)
		if hostnameObj.IsNull() {
			targetItem.Hostname = customfield.NullObject[TargetHostnameModel](ctx)
		} else {
			hostnameAttrs := hostnameObj.Attributes()
			targetItem.Hostname = customfield.NewObjectMust[TargetHostnameModel](ctx, &TargetHostnameModel{
				URLHostname:          hostnameAttrs["url_hostname"].(types.String),
				ExcludeExactHostname: hostnameAttrs["exclude_exact_hostname"].(types.Bool),
			})
		}

		// Convert redirect: types.Object → customfield.NestedObject[TargetRedirectModel]
		redirectObj := attrs["redirect"].(basetypes.ObjectValue)
		if redirectObj.IsNull() {
			targetItem.Redirect = customfield.NullObject[TargetRedirectModel](ctx)
		} else {
			redirectAttrs := redirectObj.Attributes()
			targetItem.Redirect = customfield.NewObjectMust[TargetRedirectModel](ctx, &TargetRedirectModel{
				SourceURL:           redirectAttrs["source_url"].(types.String),
				TargetURL:           redirectAttrs["target_url"].(types.String),
				StatusCode:          redirectAttrs["status_code"].(types.Int64),
				IncludeSubdomains:   redirectAttrs["include_subdomains"].(types.Bool),
				SubpathMatching:     redirectAttrs["subpath_matching"].(types.Bool),
				PreserveQueryString: redirectAttrs["preserve_query_string"].(types.Bool),
				PreservePathSuffix:  redirectAttrs["preserve_path_suffix"].(types.Bool),
			})
		}

		targetItems = append(targetItems, targetItem)
	}

	if len(targetItems) > 0 {
		itemsSet, itemDiags := customfield.NewObjectSet[TargetListItemModel](ctx, targetItems)
		diags.Append(itemDiags...)
		target.Items = itemsSet
	} else {
		target.Items = customfield.NewObjectSetMust[TargetListItemModel](ctx, []TargetListItemModel{})
	}

	return target, diags
}
