package v500

import (
	"context"
	"net"
	"net/url"
	"strings"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Transform converts a v4 cloudflare_list state to v5 state.
// This flattens the nested item/value block structure into a flat items set.
func Transform(ctx context.Context, source SourceListModel) (*TargetListModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	target := &TargetListModel{
		ID:          source.ID,
		AccountID:   source.AccountID,
		Kind:        source.Kind,
		Name:        source.Name,
		Description: source.Description,
	}

	// Transform items from v4 item blocks to v5 flat items
	var targetItems []TargetListItemModel
	for _, item := range source.Item {
		targetItem, itemDiags := transformItem(ctx, item, source.Kind.ValueString())
		diags.Append(itemDiags...)
		if itemDiags.HasError() {
			continue
		}
		if targetItem != nil {
			targetItems = append(targetItems, *targetItem)
		}
	}

	// Set items on target
	if len(targetItems) > 0 {
		target.Items = customfield.NewObjectSetMust[TargetListItemModel](ctx, targetItems)
		target.NumItems = types.Float64Value(float64(len(targetItems)))
	} else if len(source.Item) > 0 {
		// Had items but none transformed successfully
		target.Items = customfield.NewObjectSetMust[TargetListItemModel](ctx, []TargetListItemModel{})
		target.NumItems = types.Float64Value(0)
	} else {
		target.Items = customfield.NullObjectSet[TargetListItemModel](ctx)
		target.NumItems = source.NumItems
	}

	return target, diags
}

// transformItem converts a single v4 item block to a v5 item.
func transformItem(ctx context.Context, item SourceItemModel, kind string) (*TargetListItemModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	target := &TargetListItemModel{
		Comment: item.Comment,
	}

	// Extract value from the nested value block (MaxItems:1)
	if len(item.Value) == 0 {
		return nil, diags
	}
	value := item.Value[0]

	switch kind {
	case "ip":
		target.IP = normalizeIPAddress(value.IP)
		target.Hostname = customfield.NullObject[TargetHostnameModel](ctx)
		target.Redirect = customfield.NullObject[TargetRedirectModel](ctx)

	case "asn":
		// ASN 0 is not valid, treat as null
		if !value.ASN.IsNull() && !value.ASN.IsUnknown() && value.ASN.ValueInt64() != 0 {
			target.ASN = value.ASN
		}
		target.Hostname = customfield.NullObject[TargetHostnameModel](ctx)
		target.Redirect = customfield.NullObject[TargetRedirectModel](ctx)

	case "hostname":
		target.Hostname = transformHostname(ctx, value.Hostname)
		target.Redirect = customfield.NullObject[TargetRedirectModel](ctx)

	case "redirect":
		target.Hostname = customfield.NullObject[TargetHostnameModel](ctx)
		target.Redirect = transformRedirect(ctx, value.Redirect)
	default:
		target.Hostname = customfield.NullObject[TargetHostnameModel](ctx)
		target.Redirect = customfield.NullObject[TargetRedirectModel](ctx)
	}

	return target, diags
}

// transformHostname converts a v4 hostname list (MaxItems:1) to a v5 nested object.
func transformHostname(ctx context.Context, hostnames []SourceHostnameModel) customfield.NestedObject[TargetHostnameModel] {
	if len(hostnames) == 0 {
		return customfield.NullObject[TargetHostnameModel](ctx)
	}

	hostnameModel := &TargetHostnameModel{
		URLHostname: hostnames[0].URLHostname,
	}
	return customfield.NewObjectMust[TargetHostnameModel](ctx, hostnameModel)
}

// transformRedirect converts a v4 redirect list (MaxItems:1) to a v5 nested object.
// Converts "enabled"/"disabled" string booleans to actual booleans.
func transformRedirect(ctx context.Context, redirects []SourceRedirectModel) customfield.NestedObject[TargetRedirectModel] {
	if len(redirects) == 0 {
		return customfield.NullObject[TargetRedirectModel](ctx)
	}

	sourceRedirect := redirects[0]
	redirectModel := &TargetRedirectModel{
		SourceURL:           ensureSourceURLHasPath(sourceRedirect.SourceURL),
		TargetURL:           sourceRedirect.TargetURL,
		StatusCode:          sourceRedirect.StatusCode,
		IncludeSubdomains:   convertEnabledDisabledToBool(sourceRedirect.IncludeSubdomains),
		SubpathMatching:     convertEnabledDisabledToBool(sourceRedirect.SubpathMatching),
		PreserveQueryString: convertEnabledDisabledToBool(sourceRedirect.PreserveQueryString),
		PreservePathSuffix:  convertEnabledDisabledToBool(sourceRedirect.PreservePathSuffix),
	}
	return customfield.NewObjectMust[TargetRedirectModel](ctx, redirectModel)
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

// normalizeIPAddress normalizes list IP/CIDR values to match list validator behavior.
// Keep CIDRs intact except host CIDRs (/32 IPv4, /128 IPv6), which normalize to plain IP.
func normalizeIPAddress(v types.String) types.String {
	if v.IsNull() || v.IsUnknown() {
		return v
	}
	ip := v.ValueString()
	if ip == "" {
		return v
	}

	parsedIP, network, err := net.ParseCIDR(ip)
	if err == nil && network != nil {
		ones, bits := network.Mask.Size()
		if (bits == 32 && ones == 32) || (bits == 128 && ones == 128) {
			return types.StringValue(parsedIP.String())
		}
		return v
	}

	return v
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
