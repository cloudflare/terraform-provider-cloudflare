package list

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/helpers/validatordiag"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// listItemConflictsWith lists, for each item attribute, the sibling
// attributes it conflicts with, in the order the generated schema declared
// them as ConflictsWith validators.
//
// Those per-attribute ConflictsWith validators resolved
// path.MatchRelative().AtParent().AtName(...) against the whole config for
// every item. Matching a path inside a set walks every element, so validation
// was quadratic in the number of items (about 13s for 829 items). The same
// check is done here in a single pass over the items, with the same
// diagnostics.
var listItemConflictsWith = []struct {
	name      string
	conflicts []string
}{
	{name: "asn", conflicts: []string{"ip", "hostname", "redirect"}},
	{name: "hostname", conflicts: []string{"ip", "asn", "redirect"}},
	{name: "ip", conflicts: []string{"asn", "hostname", "redirect"}},
	{name: "redirect", conflicts: []string{"ip", "asn", "hostname"}},
}

var _ resource.ConfigValidator = listItemConflictsValidator{}

// listItemConflictsValidator rejects list items that set more than one of
// asn, hostname, ip and redirect.
type listItemConflictsValidator struct{}

func (v listItemConflictsValidator) Description(context.Context) string {
	return "Ensure that each list item sets at most one of asn, hostname, ip and redirect"
}

func (v listItemConflictsValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v listItemConflictsValidator) ValidateResource(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var items ListItemsSet
	diags := req.Config.GetAttribute(ctx, path.Root("items"), &items)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() || items.IsNull() || items.IsUnknown() {
		return
	}

	resp.Diagnostics.Append(listItemConflicts(ctx, path.Root("items"), items.Elements())...)
}

// listItemConflicts returns the diagnostics the ConflictsWith validators
// produced for the given item elements. For each attribute that is set (known
// and not null), each conflicting sibling that is also set is reported on the
// attribute's path. Unknown values are skipped until they are known.
func listItemConflicts(ctx context.Context, itemsPath path.Path, elems []attr.Value) diag.Diagnostics {
	var diags diag.Diagnostics

	for _, elem := range elems {
		obj, ok := elem.(basetypes.ObjectValuable)
		if !ok || elem.IsNull() || elem.IsUnknown() {
			continue
		}
		objValue, d := obj.ToObjectValue(ctx)
		if d.HasError() {
			continue
		}
		attrs := objValue.Attributes()
		elemPath := itemsPath.AtSetValue(elem)

		for _, c := range listItemConflictsWith {
			value, ok := attrs[c.name]
			if !ok || value.IsNull() || value.IsUnknown() {
				continue
			}
			attrPath := elemPath.AtName(c.name)

			for _, other := range c.conflicts {
				otherValue, ok := attrs[other]
				if !ok {
					continue
				}
				// Each sibling had its own ConflictsWith validator, so an
				// unknown sibling only skips that pair.
				if otherValue.IsUnknown() || otherValue.IsNull() {
					continue
				}
				diags.Append(validatordiag.InvalidAttributeCombinationDiagnostic(
					attrPath,
					fmt.Sprintf("Attribute %q cannot be specified when %q is specified", elemPath.AtName(other), attrPath),
				))
			}
		}
	}

	return diags
}
