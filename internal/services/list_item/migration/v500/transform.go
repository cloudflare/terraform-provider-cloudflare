package v500

import (
	"context"
	"net/url"
	"strings"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Transform converts a v4 cloudflare_list_item state to v5 state.
func Transform(ctx context.Context, source SourceListItemModel) (*TargetListItemModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	target := &TargetListItemModel{
		AccountID: source.AccountID,
		ListID:    source.ListID,
		ID:        source.ID,
		Comment:   source.Comment,
		CreatedOn: source.CreatedOn,
	}

	// IP: copy directly (no type change)
	target.IP = source.IP

	// ASN: copy directly (no type change)
	target.ASN = source.ASN

	// Hostname: List (MaxItems:1) → SingleNestedAttribute
	if len(source.Hostname) > 0 {
		hostnameModel := &TargetHostnameModel{
			URLHostname: source.Hostname[0].URLHostname,
		}
		target.Hostname = customfield.NewObjectMust[TargetHostnameModel](ctx, hostnameModel)
	} else {
		target.Hostname = customfield.NullObject[TargetHostnameModel](ctx)
	}

	// Redirect: List (MaxItems:1) → SingleNestedAttribute with boolean conversion
	if len(source.Redirect) > 0 {
		sourceRedirect := source.Redirect[0]
		redirectModel := &TargetRedirectModel{
			SourceURL:           ensureSourceURLHasPath(sourceRedirect.SourceURL),
			TargetURL:           sourceRedirect.TargetURL,
			StatusCode:          sourceRedirect.StatusCode,
			IncludeSubdomains:   convertEnabledDisabledToBool(sourceRedirect.IncludeSubdomains),
			SubpathMatching:     convertEnabledDisabledToBool(sourceRedirect.SubpathMatching),
			PreserveQueryString: convertEnabledDisabledToBool(sourceRedirect.PreserveQueryString),
			PreservePathSuffix:  convertEnabledDisabledToBool(sourceRedirect.PreservePathSuffix),
		}
		target.Redirect = customfield.NewObjectMust[TargetRedirectModel](ctx, redirectModel)
	} else {
		target.Redirect = customfield.NullObject[TargetRedirectModel](ctx)
	}

	return target, diags
}

// convertEnabledDisabledToBool converts v4 "enabled"/"disabled" strings to v5 boolean values.
func convertEnabledDisabledToBool(v types.String) types.Bool {
	if v.IsNull() || v.IsUnknown() {
		return types.BoolNull()
	}
	switch v.ValueString() {
	case "enabled":
		return types.BoolValue(true)
	case "disabled":
		return types.BoolValue(false)
	default:
		return types.BoolNull()
	}
}

// ensureSourceURLHasPath ensures the source_url has a path component.
// The v5 provider requires source_url to have a non-empty path.
func ensureSourceURLHasPath(v types.String) types.String {
	if v.IsNull() || v.IsUnknown() {
		return v
	}
	rawURL := v.ValueString()
	if rawURL == "" {
		return v
	}

	// Add scheme if missing for URL parsing
	parseURL := rawURL
	if !strings.HasPrefix(parseURL, "http://") && !strings.HasPrefix(parseURL, "https://") {
		parseURL = "https://" + parseURL
	}
	u, err := url.Parse(parseURL)
	if err != nil {
		return v
	}
	if u.Path == "" {
		return types.StringValue(rawURL + "/")
	}
	return v
}
