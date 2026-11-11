package list

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"sync"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	providerschema "github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-plugin-log/tfsdklogtest"
)

// These tests drive cloudflare_list through the framework's protocol server
// against a mock API, so they cover how the framework uses ListItemsSet during
// validation, plan, apply, refresh and state upgrade, not just the type itself.

const (
	protoAccountID = "0123456789abcdef0123456789abcdef"
	protoListID    = "2c0fc9fa937b11eaa1b71c4d701ab86e"
)

// protoAPI is a minimal in-memory implementation of the Lists API endpoints
// used by the resource.
type protoAPI struct {
	t      testing.TB
	server *httptest.Server

	mu          sync.Mutex
	kind, name  string
	description string
	items       []map[string]any
	reverse     bool     // return items in reverse order
	itemPuts    []string // request bodies of PUT .../items
}

func newProtoAPI(t testing.TB) *protoAPI {
	api := &protoAPI{t: t}
	api.server = httptest.NewServer(http.HandlerFunc(api.handle))
	t.Cleanup(api.server.Close)
	return api
}

func (api *protoAPI) seed(kind, name, description string, items []map[string]any) {
	api.mu.Lock()
	defer api.mu.Unlock()
	api.kind, api.name, api.description, api.items = kind, name, description, items
}

func (api *protoAPI) takeItemPuts() []string {
	api.mu.Lock()
	defer api.mu.Unlock()
	puts := api.itemPuts
	api.itemPuts = nil
	return puts
}

func (api *protoAPI) handle(w http.ResponseWriter, r *http.Request) {
	api.mu.Lock()
	defer api.mu.Unlock()

	body, _ := io.ReadAll(r.Body)
	listPath := "/accounts/" + protoAccountID + "/rules/lists"
	write := func(result any, extra map[string]any) {
		env := map[string]any{"result": result, "success": true, "errors": []any{}, "messages": []any{}}
		for k, v := range extra {
			env[k] = v
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(env)
	}
	list := func() map[string]any {
		l := map[string]any{
			"id": protoListID, "name": api.name, "kind": api.kind,
			"num_items": len(api.items), "num_referencing_filters": 0,
			"created_on": "2020-01-01T00:00:00Z", "modified_on": "2020-01-02T00:00:00Z",
		}
		if api.description != "" {
			l["description"] = api.description
		}
		return l
	}

	switch {
	case r.Method == http.MethodPost && r.URL.Path == listPath:
		var in map[string]any
		_ = json.Unmarshal(body, &in)
		api.name, _ = in["name"].(string)
		api.kind, _ = in["kind"].(string)
		api.description, _ = in["description"].(string)
		api.items = nil
		write(list(), nil)
	case r.Method == http.MethodGet && r.URL.Path == listPath+"/"+protoListID:
		write(list(), nil)
	case r.Method == http.MethodPut && r.URL.Path == listPath+"/"+protoListID:
		var in map[string]any
		_ = json.Unmarshal(body, &in)
		api.description, _ = in["description"].(string)
		write(list(), nil)
	case r.Method == http.MethodPut && r.URL.Path == listPath+"/"+protoListID+"/items":
		api.itemPuts = append(api.itemPuts, string(body))
		api.items = nil
		_ = json.Unmarshal(body, &api.items)
		write(map[string]any{"operation_id": "op1"}, nil)
	case r.Method == http.MethodGet && r.URL.Path == listPath+"/bulk_operations/op1":
		write(map[string]any{"id": "op1", "status": "completed"}, nil)
	case r.Method == http.MethodGet && r.URL.Path == listPath+"/"+protoListID+"/items":
		items := append([]map[string]any(nil), api.items...)
		if api.reverse {
			for i, j := 0, len(items)-1; i < j; i, j = i+1, j-1 {
				items[i], items[j] = items[j], items[i]
			}
		}
		// Small pages, so that pagination is always exercised.
		const pageSize = 5
		start, _ := strconv.Atoi(r.URL.Query().Get("cursor"))
		end := min(start+pageSize, len(items))
		page := []map[string]any{}
		if start < len(items) {
			page = items[start:end]
		}
		after := ""
		if end < len(items) {
			after = strconv.Itoa(end)
		}
		write(page, map[string]any{"result_info": map[string]any{"cursors": map[string]any{"after": after}}})
	default:
		api.t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
		w.WriteHeader(http.StatusNotFound)
	}
}

type protoProvider struct{ baseURL string }

func (p *protoProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "cloudflare"
}

func (p *protoProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = providerschema.Schema{}
}

func (p *protoProvider) Configure(_ context.Context, _ provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	resp.ResourceData = cloudflare.NewClient(
		option.WithBaseURL(p.baseURL),
		option.WithAPIToken("test-token"),
		option.WithMaxRetries(0),
	)
}

func (p *protoProvider) Resources(context.Context) []func() resource.Resource {
	return []func() resource.Resource{NewResource}
}

func (p *protoProvider) DataSources(context.Context) []func() datasource.DataSource { return nil }

// protoHarness wraps a configured protocol server for cloudflare_list.
type protoHarness struct {
	t      testing.TB
	api    *protoAPI
	server tfprotov6.ProviderServer
	obj    tftypes.Object // resource object type
	item   tftypes.Object // items element type
}

func newProtoHarness(t testing.TB, ctx context.Context) *protoHarness {
	t.Helper()
	api := newProtoAPI(t)
	server := providerserver.NewProtocol6(&protoProvider{baseURL: api.server.URL})()

	emptyType := tftypes.Object{AttributeTypes: map[string]tftypes.Type{}}
	cfg, err := tfprotov6.NewDynamicValue(emptyType, tftypes.NewValue(emptyType, map[string]tftypes.Value{}))
	if err != nil {
		t.Fatal(err)
	}
	resp, err := server.ConfigureProvider(ctx, &tfprotov6.ConfigureProviderRequest{Config: &cfg})
	if err != nil {
		t.Fatal(err)
	}
	protoNoDiags(t, "configure", resp.Diagnostics)

	obj := mustAs[tftypes.Object](t, ResourceSchema(ctx).Type().TerraformType(ctx))
	item := mustAs[tftypes.Object](t, mustAs[tftypes.Set](t, obj.AttributeTypes["items"]).ElementType)
	return &protoHarness{t: t, api: api, server: server, obj: obj, item: item}
}

// mustAs returns v as a T, failing the test if it is not one.
func mustAs[T any](t testing.TB, v any) T {
	t.Helper()
	got, ok := v.(T)
	if !ok {
		t.Fatalf("got %T, want %T", v, *new(T))
	}
	return got
}

func protoNoDiags(t testing.TB, step string, diags []*tfprotov6.Diagnostic) {
	t.Helper()
	for _, d := range diags {
		t.Errorf("%s: unexpected diagnostic: %s: %s: %s", step, d.Severity, d.Summary, d.Detail)
	}
	if len(diags) > 0 {
		t.FailNow()
	}
}

func (h *protoHarness) dynamic(v tftypes.Value) *tfprotov6.DynamicValue {
	h.t.Helper()
	dv, err := tfprotov6.NewDynamicValue(h.obj, v)
	if err != nil {
		h.t.Fatal(err)
	}
	return &dv
}

func (h *protoHarness) value(dv *tfprotov6.DynamicValue) tftypes.Value {
	h.t.Helper()
	v, err := dv.Unmarshal(h.obj)
	if err != nil {
		h.t.Fatal(err)
	}
	return v
}

func (h *protoHarness) validate(ctx context.Context, config tftypes.Value) {
	h.t.Helper()
	resp, err := h.server.ValidateResourceConfig(ctx, &tfprotov6.ValidateResourceConfigRequest{
		TypeName: "cloudflare_list", Config: h.dynamic(config),
	})
	if err != nil {
		h.t.Fatal(err)
	}
	protoNoDiags(h.t, "validate", resp.Diagnostics)
}

func (h *protoHarness) plan(ctx context.Context, prior, config tftypes.Value) *tfprotov6.PlanResourceChangeResponse {
	h.t.Helper()
	resp, err := h.server.PlanResourceChange(ctx, &tfprotov6.PlanResourceChangeRequest{
		TypeName:         "cloudflare_list",
		PriorState:       h.dynamic(prior),
		ProposedNewState: h.dynamic(h.proposedNew(prior, config)),
		Config:           h.dynamic(config),
	})
	if err != nil {
		h.t.Fatal(err)
	}
	protoNoDiags(h.t, "plan", resp.Diagnostics)
	return resp
}

func (h *protoHarness) apply(ctx context.Context, prior, planned, config tftypes.Value) tftypes.Value {
	h.t.Helper()
	resp, err := h.server.ApplyResourceChange(ctx, &tfprotov6.ApplyResourceChangeRequest{
		TypeName:     "cloudflare_list",
		PriorState:   h.dynamic(prior),
		PlannedState: h.dynamic(planned),
		Config:       h.dynamic(config),
	})
	if err != nil {
		h.t.Fatal(err)
	}
	protoNoDiags(h.t, "apply", resp.Diagnostics)
	return h.value(resp.NewState)
}

func (h *protoHarness) read(ctx context.Context, state tftypes.Value) tftypes.Value {
	h.t.Helper()
	resp, err := h.server.ReadResource(ctx, &tfprotov6.ReadResourceRequest{
		TypeName: "cloudflare_list", CurrentState: h.dynamic(state),
	})
	if err != nil {
		h.t.Fatal(err)
	}
	protoNoDiags(h.t, "read", resp.Diagnostics)
	return h.value(resp.NewState)
}

func (h *protoHarness) upgrade(ctx context.Context, version int64, rawJSON string) tftypes.Value {
	h.t.Helper()
	resp, err := h.server.UpgradeResourceState(ctx, &tfprotov6.UpgradeResourceStateRequest{
		TypeName: "cloudflare_list", Version: version, RawState: &tfprotov6.RawState{JSON: []byte(rawJSON)},
	})
	if err != nil {
		h.t.Fatal(err)
	}
	protoNoDiags(h.t, "upgrade", resp.Diagnostics)
	return h.value(resp.UpgradedState)
}

// proposedNew mimics Terraform Core for this schema: configured values, with
// computed attributes taken from the prior state where the config is null.
func (h *protoHarness) proposedNew(prior, config tftypes.Value) tftypes.Value {
	if config.IsNull() || prior.IsNull() {
		return config
	}
	cfg, pr := protoAttrs(h.t, config), protoAttrs(h.t, prior)
	for _, name := range []string{"id", "created_on", "modified_on", "num_items", "num_referencing_filters"} {
		if cfg[name].IsNull() {
			cfg[name] = pr[name]
		}
	}
	return tftypes.NewValue(h.obj, cfg)
}

func (h *protoHarness) object(attrs map[string]tftypes.Value) tftypes.Value {
	full := make(map[string]tftypes.Value, len(h.obj.AttributeTypes))
	for name, typ := range h.obj.AttributeTypes {
		full[name] = tftypes.NewValue(typ, nil)
	}
	for name, v := range attrs {
		full[name] = v
	}
	return tftypes.NewValue(h.obj, full)
}

func (h *protoHarness) config(description string, items tftypes.Value) tftypes.Value {
	return h.object(map[string]tftypes.Value{
		"account_id":  tftypes.NewValue(tftypes.String, protoAccountID),
		"kind":        tftypes.NewValue(tftypes.String, "ip"),
		"name":        tftypes.NewValue(tftypes.String, "list"),
		"description": tftypes.NewValue(tftypes.String, description),
		"items":       items,
	})
}

// configFromState returns the user-settable attributes of state as a config.
func (h *protoHarness) configFromState(state tftypes.Value) tftypes.Value {
	attrs := protoAttrs(h.t, state)
	return h.object(map[string]tftypes.Value{
		"account_id":  attrs["account_id"],
		"kind":        attrs["kind"],
		"name":        attrs["name"],
		"description": attrs["description"],
		"items":       attrs["items"],
	})
}

// state returns a prior state for the list currently served by the mock API.
func (h *protoHarness) state(description string, items tftypes.Value) tftypes.Value {
	return h.object(map[string]tftypes.Value{
		"id":                      tftypes.NewValue(tftypes.String, protoListID),
		"account_id":              tftypes.NewValue(tftypes.String, protoAccountID),
		"kind":                    tftypes.NewValue(tftypes.String, "ip"),
		"name":                    tftypes.NewValue(tftypes.String, "list"),
		"description":             tftypes.NewValue(tftypes.String, description),
		"created_on":              tftypes.NewValue(tftypes.String, "2020-01-01T00:00:00Z"),
		"modified_on":             tftypes.NewValue(tftypes.String, "2020-01-02T00:00:00Z"),
		"num_items":               tftypes.NewValue(tftypes.Number, big.NewFloat(0)),
		"num_referencing_filters": tftypes.NewValue(tftypes.Number, big.NewFloat(0)),
		"items":                   items,
	})
}

// protoItem describes an ip list item. A nil comment is null.
type protoItem struct {
	ip             string
	comment        *string
	unknownComment bool
}

func protoIPItems(n int) []protoItem {
	items := make([]protoItem, n)
	for i := range items {
		items[i].ip = fmt.Sprintf("10.%d.%d.%d", i>>16&255, i>>8&255, i&255)
		if i%4 != 3 {
			c := fmt.Sprintf("comment %d", i)
			items[i].comment = &c
		}
	}
	return items
}

func (h *protoHarness) items(items []protoItem) tftypes.Value {
	elems := make([]tftypes.Value, 0, len(items))
	for _, it := range items {
		attrs := make(map[string]tftypes.Value, len(h.item.AttributeTypes))
		for name, typ := range h.item.AttributeTypes {
			attrs[name] = tftypes.NewValue(typ, nil)
		}
		attrs["ip"] = tftypes.NewValue(tftypes.String, it.ip)
		switch {
		case it.unknownComment:
			attrs["comment"] = tftypes.NewValue(tftypes.String, tftypes.UnknownValue)
		case it.comment != nil:
			attrs["comment"] = tftypes.NewValue(tftypes.String, *it.comment)
		}
		elems = append(elems, tftypes.NewValue(h.item, attrs))
	}
	return tftypes.NewValue(h.obj.AttributeTypes["items"], elems)
}

func protoAPIItems(items []protoItem) []map[string]any {
	out := make([]map[string]any, 0, len(items))
	for i, it := range items {
		m := map[string]any{"id": fmt.Sprintf("item%d", i), "ip": it.ip, "created_on": "2020-01-01T00:00:00Z"}
		if it.comment != nil {
			m["comment"] = *it.comment
		}
		// The API returns empty nested objects for ip items.
		m["hostname"] = map[string]any{}
		m["redirect"] = map[string]any{}
		out = append(out, m)
	}
	return out
}

// protoAPIItemsFromState converts the items of a state value into API items.
func protoAPIItemsFromState(t testing.TB, state tftypes.Value) []map[string]any {
	t.Helper()
	var elems []tftypes.Value
	if err := protoAttrs(t, state)["items"].As(&elems); err != nil {
		t.Fatal(err)
	}
	out := make([]map[string]any, 0, len(elems))
	for i, e := range elems {
		m := mustAs[map[string]any](t, protoJSON(t, e))
		m["id"] = fmt.Sprintf("item%d", i)
		out = append(out, m)
	}
	return out
}

func protoJSON(t testing.TB, v tftypes.Value) any {
	t.Helper()
	switch {
	case v.Type().Is(tftypes.Object{}):
		attrs := protoAttrs(t, v)
		out := make(map[string]any, len(attrs))
		for name, attr := range attrs {
			if !attr.IsNull() {
				out[name] = protoJSON(t, attr)
			}
		}
		return out
	case v.Type().Is(tftypes.String):
		var s string
		_ = v.As(&s)
		return s
	case v.Type().Is(tftypes.Bool):
		var b bool
		_ = v.As(&b)
		return b
	case v.Type().Is(tftypes.Number):
		var f big.Float
		_ = v.As(&f)
		i, _ := f.Int64()
		return i
	}
	t.Fatalf("unsupported value %s", v)
	return nil
}

func protoAttrs(t testing.TB, v tftypes.Value) map[string]tftypes.Value {
	t.Helper()
	var attrs map[string]tftypes.Value
	if err := v.As(&attrs); err != nil {
		t.Fatal(err)
	}
	return attrs
}

// protoAssertEqual compares with tftypes.Value.Equal, which ignores set
// element order, like Terraform Core.
func protoAssertEqual(t testing.TB, what string, got, want tftypes.Value) {
	t.Helper()
	if !got.Equal(want) {
		t.Errorf("%s differs\n got: %s\nwant: %s", what, got, want)
	}
}

func TestListResourceProtocolRefreshReorderedAPIResults(t *testing.T) {
	// Enable framework trace logs to check how semantic equality ran.
	t.Setenv("TF_LOG_SDK_FRAMEWORK", "TRACE")
	var logs bytes.Buffer
	ctx := tfsdklogtest.RootLogger(context.Background(), &logs)

	h := newProtoHarness(t, ctx)
	items := protoIPItems(12)
	h.api.seed("ip", "list", "d", protoAPIItems(items))
	h.api.reverse = true

	prior := h.state("d", h.items(items))
	logs.Reset()
	got := h.read(ctx, prior)

	attrs := protoAttrs(t, got)
	protoAssertEqual(t, "items", attrs["items"], h.items(items))
	if !attrs["num_items"].Equal(tftypes.NewValue(tftypes.Number, big.NewFloat(12))) {
		t.Errorf("num_items = %s, want 12", attrs["num_items"])
	}

	// The prior and refreshed sets are equal, so ListItemsSet.SetSemanticEquals
	// must have been called and must have accepted the prior value without the
	// framework walking the elements (which logs once per kept element).
	calledSemanticEquals, elementsKept := false, 0
	dec := json.NewDecoder(&logs)
	for {
		var entry map[string]any
		if err := dec.Decode(&entry); err == io.EOF {
			break
		} else if err != nil {
			t.Fatal(err)
		}
		attrPath, _ := entry["tf_attribute_path"].(string)
		switch entry["@message"] {
		case "Called provider defined type-based SemanticEquals":
			calledSemanticEquals = calledSemanticEquals || attrPath == "items"
		case "Value switched to prior value due to semantic equality logic":
			if strings.HasPrefix(attrPath, "items[") {
				elementsKept++
			}
		}
	}
	if !calledSemanticEquals {
		t.Errorf("framework did not call ListItemsSet.SetSemanticEquals for items")
	}
	if elementsKept != 0 {
		t.Errorf("framework compared %d set elements individually, want 0", elementsKept)
	}

	// A second refresh is stable.
	protoAssertEqual(t, "second refresh", h.read(ctx, got), got)
}

func TestListResourceProtocolRefreshDrift(t *testing.T) {
	ctx := context.Background()
	h := newProtoHarness(t, ctx)
	prior := protoIPItems(12)

	drifted := append([]protoItem(nil), prior[1:]...) // first item removed
	changed := "changed outside terraform"
	drifted[4].comment = &changed                         // one comment changed
	drifted = append(drifted, protoItem{ip: "192.0.2.1"}) // one item added
	h.api.seed("ip", "list", "d", protoAPIItems(drifted))
	h.api.reverse = true

	got := h.read(ctx, h.state("d", h.items(prior)))
	protoAssertEqual(t, "items", protoAttrs(t, got)["items"], h.items(drifted))
}

func TestListResourceProtocolApplyOneCommentChanged(t *testing.T) {
	ctx := context.Background()
	h := newProtoHarness(t, ctx)
	items := protoIPItems(12)
	h.api.seed("ip", "list", "d", protoAPIItems(items))
	h.api.reverse = true
	state := h.read(ctx, h.state("d", h.items(items)))

	desired := append([]protoItem(nil), items...)
	for i, j := 0, len(desired)-1; i < j; i, j = i+1, j-1 {
		desired[i], desired[j] = desired[j], desired[i]
	}
	changed := "changed"
	desired[5].comment = &changed
	config := h.config("d", h.items(desired))

	h.validate(ctx, config)
	planResp := h.plan(ctx, state, config)
	if len(planResp.RequiresReplace) != 0 {
		t.Errorf("RequiresReplace = %v, want none", planResp.RequiresReplace)
	}
	planned := h.value(planResp.PlannedState)
	protoAssertEqual(t, "planned items", protoAttrs(t, planned)["items"], h.items(desired))

	applied := h.apply(ctx, state, planned, config)
	protoAssertEqual(t, "applied items", protoAttrs(t, applied)["items"], h.items(desired))
	puts := h.api.takeItemPuts()
	if len(puts) != 1 || !strings.Contains(puts[0], `"comment":"changed"`) {
		t.Errorf("PUT items requests = %q, want one containing the changed comment", puts)
	}

	// After a refresh, the same configuration plans no changes.
	refreshed := h.read(ctx, applied)
	protoAssertEqual(t, "refreshed state", refreshed, applied)
	protoAssertEqual(t, "follow-up plan", h.value(h.plan(ctx, refreshed, config).PlannedState), refreshed)
}

func TestListResourceProtocolPartiallyUnknownItems(t *testing.T) {
	ctx := context.Background()
	h := newProtoHarness(t, ctx)
	items := protoIPItems(12)
	h.api.seed("ip", "list", "d", protoAPIItems(items))
	state := h.read(ctx, h.state("d", h.items(items)))

	desired := append([]protoItem(nil), items...)
	desired[3].unknownComment = true
	config := h.config("d", h.items(desired))

	h.validate(ctx, config)
	planned := h.value(h.plan(ctx, state, config).PlannedState)
	protoAssertEqual(t, "planned items", protoAttrs(t, planned)["items"], h.items(desired))

	var elems []tftypes.Value
	if err := protoAttrs(t, planned)["items"].As(&elems); err != nil {
		t.Fatal(err)
	}
	unknown := 0
	for _, e := range elems {
		if !e.IsFullyKnown() {
			unknown++
		}
	}
	if len(elems) != 12 || unknown != 1 {
		t.Errorf("planned %d items with %d partially unknown, want 12 with 1", len(elems), unknown)
	}
}

func TestListResourceProtocolUpgradeState(t *testing.T) {
	for _, tc := range []struct {
		name    string
		version int64
		raw     string
	}{
		{
			name:    "v0 with v5 items",
			version: 0,
			raw: `{"id":"` + protoListID + `","account_id":"` + protoAccountID + `","kind":"hostname","name":"list","description":"d","num_items":2,` +
				`"created_on":"2020-01-01T00:00:00Z","modified_on":"2020-01-02T00:00:00Z","num_referencing_filters":0,` +
				`"items":[{"hostname":{"url_hostname":"example.com","exclude_exact_hostname":null},"comment":"x"},` +
				`{"hostname":{"url_hostname":"*.example.org","exclude_exact_hostname":true}}]}`,
		},
		{
			name:    "v0 with v4 item blocks",
			version: 0,
			raw: `{"id":"` + protoListID + `","account_id":"` + protoAccountID + `","kind":"ip","name":"list","description":"d","num_items":3,` +
				`"item":[{"comment":"c1","value":[{"ip":"192.0.2.1","asn":null,"hostname":[],"redirect":[]}]},` +
				`{"comment":null,"value":[{"ip":"192.0.2.0/24","asn":null,"hostname":[],"redirect":[]}]},` +
				`{"comment":"c3","value":[{"ip":"2001:db8::1","asn":null,"hostname":[],"redirect":[]}]}]}`,
		},
		{
			name:    "v1",
			version: 1,
			raw: `{"id":"` + protoListID + `","account_id":"` + protoAccountID + `","kind":"ip","name":"list","description":"d","num_items":2,` +
				`"created_on":"2020-01-01T00:00:00Z","modified_on":"2020-01-02T00:00:00Z","num_referencing_filters":0,` +
				`"items":[{"ip":"192.0.2.1","comment":"a","asn":null,"hostname":null,"redirect":null},` +
				`{"ip":"192.0.2.2","comment":null,"asn":null,"hostname":null,"redirect":null}]}`,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			h := newProtoHarness(t, ctx)
			upgraded := h.upgrade(ctx, tc.version, tc.raw)

			attrs := protoAttrs(t, upgraded)
			var kind, name, description string
			_ = attrs["kind"].As(&kind)
			_ = attrs["name"].As(&name)
			_ = attrs["description"].As(&description)
			if attrs["items"].IsNull() {
				t.Fatal("upgraded state has no items")
			}

			// The API still holds the same items.
			h.api.seed(kind, name, description, protoAPIItemsFromState(t, upgraded))
			h.api.reverse = true
			refreshed := h.read(ctx, upgraded)
			protoAssertEqual(t, "refreshed items", protoAttrs(t, refreshed)["items"], attrs["items"])

			config := h.configFromState(refreshed)
			h.validate(ctx, config)
			protoAssertEqual(t, "plan", h.value(h.plan(ctx, refreshed, config).PlannedState), refreshed)
		})
	}
}

// Benchmarks for refresh and plan through the protocol server. Run with:
//
//	go test ./internal/services/list/ -run '^$' -bench 'ListResource' -benchmem
func BenchmarkListResourceRead(b *testing.B) {
	for _, n := range []int{100, 400, 829} {
		for _, tc := range []string{"unchanged", "reordered", "one_drifted"} {
			b.Run(fmt.Sprintf("%s/%d", tc, n), func(b *testing.B) {
				ctx := context.Background()
				h := newProtoHarness(b, ctx)
				items := protoIPItems(n)
				apiItems := protoAPIItems(items)
				if tc == "one_drifted" {
					apiItems[n/2]["comment"] = "drifted"
				}
				h.api.seed("ip", "list", "d", apiItems)
				h.api.reverse = tc == "reordered"
				req := &tfprotov6.ReadResourceRequest{TypeName: "cloudflare_list", CurrentState: h.dynamic(h.state("d", h.items(items)))}

				b.ReportAllocs()
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					resp, err := h.server.ReadResource(ctx, req)
					if err != nil || len(resp.Diagnostics) > 0 {
						b.Fatal(err, resp.Diagnostics)
					}
				}
			})
		}
	}
}

func BenchmarkListResourcePlan(b *testing.B) {
	for _, n := range []int{100, 400, 829} {
		for _, tc := range []string{"unchanged", "reordered", "one_changed"} {
			b.Run(fmt.Sprintf("%s/%d", tc, n), func(b *testing.B) {
				ctx := context.Background()
				h := newProtoHarness(b, ctx)
				items := protoIPItems(n)
				h.api.seed("ip", "list", "d", protoAPIItems(items))
				state := h.read(ctx, h.state("d", h.items(items)))

				desired := append([]protoItem(nil), items...)
				if tc != "unchanged" {
					for i, j := 0, len(desired)-1; i < j; i, j = i+1, j-1 {
						desired[i], desired[j] = desired[j], desired[i]
					}
				}
				if tc == "one_changed" {
					changed := "changed"
					desired[n/2].comment = &changed
				}
				config := h.config("d", h.items(desired))
				req := &tfprotov6.PlanResourceChangeRequest{
					TypeName:         "cloudflare_list",
					PriorState:       h.dynamic(state),
					ProposedNewState: h.dynamic(h.proposedNew(state, config)),
					Config:           h.dynamic(config),
				}

				b.ReportAllocs()
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					resp, err := h.server.PlanResourceChange(ctx, req)
					if err != nil || len(resp.Diagnostics) > 0 {
						b.Fatal(err, resp.Diagnostics)
					}
				}
			})
		}
	}
}

func BenchmarkListResourceValidate(b *testing.B) {
	for _, n := range []int{100, 400, 829} {
		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			ctx := context.Background()
			h := newProtoHarness(b, ctx)
			req := &tfprotov6.ValidateResourceConfigRequest{
				TypeName: "cloudflare_list",
				Config:   h.dynamic(h.config("d", h.items(protoIPItems(n)))),
			}

			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				resp, err := h.server.ValidateResourceConfig(ctx, req)
				if err != nil || len(resp.Diagnostics) > 0 {
					b.Fatal(err, resp.Diagnostics)
				}
			}
		})
	}
}
