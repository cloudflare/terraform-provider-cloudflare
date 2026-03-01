package v500

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeFromV0 handles state upgrades from v4 (schema_version=0) to v5.
// This is the main v4→v5 migration path for cloudflare_list_item.
// Key transformations:
//   - hostname: List (MaxItems:1) → SingleNestedAttribute
//   - redirect: List (MaxItems:1) → SingleNestedAttribute
//   - redirect boolean fields: "enabled"/"disabled" strings → true/false booleans
func UpgradeFromV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading cloudflare_list_item state from v4 (schema_version=0)")

	var sourceState SourceListItemModel
	resp.Diagnostics.Append(req.State.Get(ctx, &sourceState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	targetState, diags := Transform(ctx, sourceState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, targetState)...)

	tflog.Info(ctx, "State upgrade for cloudflare_list_item from v4 completed successfully")
}

// UpgradeFromV1 handles state upgrades from v4.52.5 framework (schema_version=1) to v500.
// In v4.52.5, hostname and redirect are ListNestedBlocks (lists with max 1 element).
// This upgrade converts them to SingleNestedAttributes (objects).
// Redirect boolean fields are already bools in v1 (no string conversion needed).
func UpgradeFromV1(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading cloudflare_list_item state from v4.52.5 framework (schema_version=1)")

	var sourceState SourceListItemV1Model
	resp.Diagnostics.Append(req.State.Get(ctx, &sourceState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	target := &TargetListItemModel{
		AccountID: sourceState.AccountID,
		ListID:    sourceState.ListID,
		ID:        sourceState.ID,
		IP:        sourceState.IP,
		ASN:       sourceState.ASN,
		Comment:   sourceState.Comment,
	}

	// Hostname: List (max 1) → SingleNestedAttribute
	if len(sourceState.Hostname) > 0 {
		target.Hostname = customfield.NewObjectMust[TargetHostnameModel](ctx, &TargetHostnameModel{
			URLHostname: sourceState.Hostname[0].URLHostname,
		})
	} else {
		target.Hostname = customfield.NullObject[TargetHostnameModel](ctx)
	}

	// Redirect: List (max 1) → SingleNestedAttribute (booleans are already bools)
	if len(sourceState.Redirect) > 0 {
		src := sourceState.Redirect[0]
		target.Redirect = customfield.NewObjectMust[TargetRedirectModel](ctx, &TargetRedirectModel{
			SourceURL:           ensureSourceURLHasPath(src.SourceURL),
			TargetURL:           src.TargetURL,
			StatusCode:          src.StatusCode,
			IncludeSubdomains:   src.IncludeSubdomains,
			SubpathMatching:     src.SubpathMatching,
			PreserveQueryString: src.PreserveQueryString,
			PreservePathSuffix:  src.PreservePathSuffix,
		})
	} else {
		target.Redirect = customfield.NullObject[TargetRedirectModel](ctx)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, target)...)
	tflog.Info(ctx, "State upgrade for cloudflare_list_item from v1 completed successfully")
}
