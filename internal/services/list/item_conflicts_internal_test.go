package list

import (
	"context"
	"fmt"
	"math/big"
	"sort"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// conflictsWithSchema returns the resource schema with the per-attribute
// ConflictsWith validators that listItemConflictsValidator replaces.
func conflictsWithSchema(t testing.TB, ctx context.Context) schema.Schema {
	t.Helper()
	s := ResourceSchema(ctx)
	items := mustAs[schema.SetNestedAttribute](t, s.Attributes["items"])
	attrs := items.NestedObject.Attributes
	sibling := func(name string) path.Expression { return path.MatchRelative().AtParent().AtName(name) }

	asn := mustAs[schema.Int64Attribute](t, attrs["asn"])
	asn.Validators = append(asn.Validators,
		int64validator.ConflictsWith(sibling("ip")),
		int64validator.ConflictsWith(sibling("hostname")),
		int64validator.ConflictsWith(sibling("redirect")),
	)
	attrs["asn"] = asn

	hostname := mustAs[schema.SingleNestedAttribute](t, attrs["hostname"])
	hostname.Validators = append(hostname.Validators,
		objectvalidator.ConflictsWith(sibling("ip")),
		objectvalidator.ConflictsWith(sibling("asn")),
		objectvalidator.ConflictsWith(sibling("redirect")),
	)
	attrs["hostname"] = hostname

	ip := mustAs[schema.StringAttribute](t, attrs["ip"])
	ip.Validators = append(ip.Validators,
		stringvalidator.ConflictsWith(sibling("asn")),
		stringvalidator.ConflictsWith(sibling("hostname")),
		stringvalidator.ConflictsWith(sibling("redirect")),
	)
	attrs["ip"] = ip

	redirect := mustAs[schema.SingleNestedAttribute](t, attrs["redirect"])
	redirect.Validators = append(redirect.Validators,
		objectvalidator.ConflictsWith(sibling("ip")),
		objectvalidator.ConflictsWith(sibling("asn")),
		objectvalidator.ConflictsWith(sibling("hostname")),
	)
	attrs["redirect"] = redirect

	return s
}

// conflictsWithResource is cloudflare_list as it was validated before
// listItemConflictsValidator: ConflictsWith in the schema, and only
// ListValidator as a config validator. The schema is built by the test, so
// that a failed type assertion fails the test.
type conflictsWithResource struct {
	*ListResource
	schema schema.Schema
}

func (r conflictsWithResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = r.schema
}

func (r conflictsWithResource) ConfigValidators(context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{ListValidator{}}
}

type conflictsWithProvider struct {
	protoProvider
	schema schema.Schema
}

func (p *conflictsWithProvider) Resources(context.Context) []func() resource.Resource {
	return []func() resource.Resource{func() resource.Resource {
		return conflictsWithResource{ListResource: &ListResource{}, schema: p.schema}
	}}
}

func protoDiagKeys(diags []*tfprotov6.Diagnostic) []string {
	keys := make([]string, 0, len(diags))
	for _, d := range diags {
		attrPath := ""
		if d.Attribute != nil {
			attrPath = d.Attribute.String()
		}
		keys = append(keys, fmt.Sprintf("%s | %s | %s | %s", d.Severity, attrPath, d.Summary, d.Detail))
	}
	sort.Strings(keys)
	return keys
}

// TestListItemConflictsMatchesConflictsWith checks that the config validator
// reports exactly the diagnostics the ConflictsWith validators reported.
func TestListItemConflictsMatchesConflictsWith(t *testing.T) {
	ctx := context.Background()
	h := newProtoHarness(t, ctx)
	reference := providerserver.NewProtocol6(&conflictsWithProvider{
		protoProvider: protoProvider{baseURL: h.api.server.URL},
		schema:        conflictsWithSchema(t, ctx),
	})()

	itemType := h.item
	hostnameType := mustAs[tftypes.Object](t, itemType.AttributeTypes["hostname"])
	redirectType := mustAs[tftypes.Object](t, itemType.AttributeTypes["redirect"])
	unknown := func(name string) tftypes.Value {
		return tftypes.NewValue(itemType.AttributeTypes[name], tftypes.UnknownValue)
	}
	set := map[string]tftypes.Value{
		"ip":      tftypes.NewValue(tftypes.String, "192.0.2.1"),
		"asn":     tftypes.NewValue(tftypes.Number, big.NewFloat(13335)),
		"comment": tftypes.NewValue(tftypes.String, "c"),
		"hostname": tftypes.NewValue(hostnameType, map[string]tftypes.Value{
			"url_hostname":           tftypes.NewValue(tftypes.String, "example.com"),
			"exclude_exact_hostname": tftypes.NewValue(tftypes.Bool, nil),
		}),
		"redirect": tftypes.NewValue(redirectType, map[string]tftypes.Value{
			"source_url":            tftypes.NewValue(tftypes.String, "example.com/a"),
			"target_url":            tftypes.NewValue(tftypes.String, "https://example.com/b"),
			"include_subdomains":    tftypes.NewValue(tftypes.Bool, nil),
			"preserve_path_suffix":  tftypes.NewValue(tftypes.Bool, nil),
			"preserve_query_string": tftypes.NewValue(tftypes.Bool, nil),
			"status_code":           tftypes.NewValue(tftypes.Number, nil),
			"subpath_matching":      tftypes.NewValue(tftypes.Bool, nil),
		}),
	}
	// item builds an element from attribute names; "name?" is unknown.
	item := func(names ...string) tftypes.Value {
		attrs := make(map[string]tftypes.Value, len(itemType.AttributeTypes))
		for name, typ := range itemType.AttributeTypes {
			attrs[name] = tftypes.NewValue(typ, nil)
		}
		for _, name := range names {
			if n := len(name); name[n-1] == '?' {
				attrs[name[:n-1]] = unknown(name[:n-1])
				continue
			}
			attrs[name] = set[name]
		}
		return tftypes.NewValue(itemType, attrs)
	}
	itemsType := h.obj.AttributeTypes["items"]
	items := func(elems ...tftypes.Value) tftypes.Value { return tftypes.NewValue(itemsType, elems) }

	for _, tc := range []struct {
		name  string
		items tftypes.Value
	}{
		{"null items", tftypes.NewValue(itemsType, nil)},
		{"unknown items", tftypes.NewValue(itemsType, tftypes.UnknownValue)},
		{"empty items", items()},
		{"no conflicts", items(item("ip", "comment"), item("ip"))},
		{"ip and asn", items(item("ip", "asn"))},
		{"ip and hostname", items(item("ip", "hostname"))},
		{"ip and redirect", items(item("ip", "redirect"))},
		{"asn and hostname", items(item("asn", "hostname"))},
		{"asn and redirect", items(item("asn", "redirect"))},
		{"hostname and redirect", items(item("hostname", "redirect"))},
		{"three set", items(item("ip", "asn", "redirect"))},
		{"all four set", items(item("ip", "asn", "hostname", "redirect", "comment"))},
		{"ip and unknown asn", items(item("ip", "asn?"))},
		{"unknown ip and asn", items(item("ip?", "asn"))},
		{"unknown before set sibling", items(item("ip", "asn?", "hostname"))},
		// An unknown hostname is not covered: ListValidator dereferences a nil
		// hostname model in that case, with or without this change.
		{"unknown after set sibling", items(item("ip", "asn", "redirect?"))},
		{"unknown redirect with set hostname", items(item("hostname", "redirect?"))},
		{"unknown comment with conflict", items(item("ip", "asn", "comment?"))},
		{"several items", items(item("ip"), item("ip", "asn"), item("hostname", "redirect"), item("ip", "comment"))},
		{"unknown element", items(item("ip"), tftypes.NewValue(itemType, tftypes.UnknownValue))},
	} {
		t.Run(tc.name, func(t *testing.T) {
			config := h.config("d", tc.items)
			req := &tfprotov6.ValidateResourceConfigRequest{TypeName: "cloudflare_list", Config: h.dynamic(config)}

			want, err := reference.ValidateResourceConfig(ctx, req)
			if err != nil {
				t.Fatal(err)
			}
			got, err := h.server.ValidateResourceConfig(ctx, req)
			if err != nil {
				t.Fatal(err)
			}

			wantKeys, gotKeys := protoDiagKeys(want.Diagnostics), protoDiagKeys(got.Diagnostics)
			if fmt.Sprint(wantKeys) != fmt.Sprint(gotKeys) {
				t.Errorf("diagnostics differ\n got: %q\nwant: %q", gotKeys, wantKeys)
			}
		})
	}
}

// TestListItemConflictsReportsConflicts guards against the equivalence test
// passing because neither side reports anything.
func TestListItemConflictsReportsConflicts(t *testing.T) {
	ctx := context.Background()
	h := newProtoHarness(t, ctx)

	elems := []tftypes.Value{}
	for _, it := range protoIPItems(2) {
		attrs := make(map[string]tftypes.Value, len(h.item.AttributeTypes))
		for name, typ := range h.item.AttributeTypes {
			attrs[name] = tftypes.NewValue(typ, nil)
		}
		attrs["ip"] = tftypes.NewValue(tftypes.String, it.ip)
		attrs["asn"] = tftypes.NewValue(tftypes.Number, big.NewFloat(13335))
		elems = append(elems, tftypes.NewValue(h.item, attrs))
	}
	config := h.config("d", tftypes.NewValue(h.obj.AttributeTypes["items"], elems))

	resp, err := h.server.ValidateResourceConfig(ctx, &tfprotov6.ValidateResourceConfigRequest{
		TypeName: "cloudflare_list", Config: h.dynamic(config),
	})
	if err != nil {
		t.Fatal(err)
	}
	// ip conflicts with asn and asn with ip, for each of the two items.
	if n := len(resp.Diagnostics); n != 4 {
		t.Fatalf("got %d diagnostics, want 4: %q", n, protoDiagKeys(resp.Diagnostics))
	}
	for _, d := range resp.Diagnostics {
		if d.Summary != "Invalid Attribute Combination" {
			t.Errorf("unexpected diagnostic: %s: %s", d.Summary, d.Detail)
		}
	}
}

// TestListItemConflictsNotInSchema fails if a regeneration of schema.go
// restores the per-attribute ConflictsWith validators on items.* or drops
// listItemConflictsValidator. The framework deduplicates equal diagnostics,
// so restored validators would not change any output; they would only make
// ValidateResourceConfig quadratic in the number of items again.
func TestListItemConflictsNotInSchema(t *testing.T) {
	ctx := context.Background()
	items, ok := ResourceSchema(ctx).Attributes["items"].(schema.SetNestedAttribute)
	if !ok {
		t.Fatal("items is not a SetNestedAttribute")
	}

	for _, c := range listItemConflictsWith {
		var validators []any
		switch a := items.NestedObject.Attributes[c.name].(type) {
		case schema.Int64Attribute:
			for _, v := range a.Validators {
				validators = append(validators, v)
			}
		case schema.StringAttribute:
			for _, v := range a.Validators {
				validators = append(validators, v)
			}
		case schema.SingleNestedAttribute:
			for _, v := range a.Validators {
				validators = append(validators, v)
			}
		default:
			t.Fatalf("items.%s: unexpected attribute type %T", c.name, a)
		}
		for _, v := range validators {
			if typ := fmt.Sprintf("%T", v); strings.Contains(typ, "ConflictsWith") {
				t.Errorf("items.%s has %s; conflicts are checked by listItemConflictsValidator", c.name, typ)
			}
		}
	}

	found := false
	for _, v := range (&ListResource{}).ConfigValidators(ctx) {
		if _, ok := v.(listItemConflictsValidator); ok {
			found = true
		}
	}
	if !found {
		t.Error("ConfigValidators does not include listItemConflictsValidator")
	}
}
