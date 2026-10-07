package workers_script_test

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/workers_script"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

const (
	readTestAccountID  = "acc-123"
	readTestScriptName = "my-worker"
	readTestIndexJS    = `import { helper } from "./util.js"; export default { fetch() { return new Response(helper()); } };`
	readTestUtilJS     = `export function helper() { return "hi"; }`

	readTestServiceWorkerJS = `addEventListener("fetch", (e) => e.respondWith(new Response("hi")));`
)

type readTestPart struct {
	name        string
	content     string
	contentType string
}

func newWorkersScriptResource(t *testing.T, baseURL string) *workers_script.WorkersScriptResource {
	t.Helper()
	r := workers_script.NewResource().(*workers_script.WorkersScriptResource)
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIToken("test-token"),
		option.WithMaxRetries(0),
	)
	resp := &resource.ConfigureResponse{}
	r.Configure(context.Background(), resource.ConfigureRequest{ProviderData: client}, resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("Configure failed: %v", resp.Diagnostics)
	}
	return r
}

// makeWorkersScriptState builds a state where every attribute is null except
// the identifiers and the given overrides.
func makeWorkersScriptState(t *testing.T, ctx context.Context, overrides map[string]tftypes.Value) tfsdk.State {
	t.Helper()
	schema := workers_script.ResourceSchema(ctx)
	objType := schema.Type().TerraformType(ctx).(tftypes.Object)
	attrs := make(map[string]tftypes.Value, len(objType.AttributeTypes))
	for name, typ := range objType.AttributeTypes {
		attrs[name] = tftypes.NewValue(typ, nil)
	}
	attrs["id"] = tftypes.NewValue(tftypes.String, readTestScriptName)
	attrs["script_name"] = tftypes.NewValue(tftypes.String, readTestScriptName)
	attrs["account_id"] = tftypes.NewValue(tftypes.String, readTestAccountID)
	for name, v := range overrides {
		attrs[name] = v
	}
	return tfsdk.State{Schema: schema, Raw: tftypes.NewValue(objType, attrs)}
}

// newWorkersScriptServer serves the endpoints hit by Read. The content
// endpoint returns the given parts as multipart/form-data, with CF-Entrypoint
// set when entrypoint is non-empty.
func newWorkersScriptServer(t *testing.T, entrypoint string, parts []readTestPart) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case strings.HasSuffix(r.URL.Path, fmt.Sprintf("/accounts/%s/workers/services/%s", readTestAccountID, readTestScriptName)):
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, `{"success":true,"errors":[],"messages":[],"result":{"id":%q,"default_environment":{"environment":"production","script":{"id":%q}}}}`, readTestScriptName, readTestScriptName)
		case strings.HasSuffix(r.URL.Path, fmt.Sprintf("/accounts/%s/workers/scripts/%s/settings", readTestAccountID, readTestScriptName)):
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `{"success":true,"errors":[],"messages":[],"result":{"compatibility_date":"2024-01-01","placement":{}}}`)
		case strings.HasSuffix(r.URL.Path, fmt.Sprintf("/accounts/%s/workers/scripts/%s/content/v2", readTestAccountID, readTestScriptName)):
			var body bytes.Buffer
			mw := multipart.NewWriter(&body)
			for _, p := range parts {
				h := make(map[string][]string)
				h["Content-Disposition"] = []string{fmt.Sprintf(`form-data; name=%q; filename=%q`, p.name, p.name)}
				h["Content-Type"] = []string{p.contentType}
				pw, err := mw.CreatePart(h)
				if err != nil {
					t.Errorf("failed to create part: %v", err)
					return
				}
				pw.Write([]byte(p.content))
			}
			mw.Close()
			if entrypoint != "" {
				w.Header().Set("CF-Entrypoint", entrypoint)
			}
			w.Header().Set("Content-Type", mw.FormDataContentType())
			w.Write(body.Bytes())
		default:
			t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
			w.WriteHeader(http.StatusNotFound)
		}
	}))
}

func readWorkersScript(t *testing.T, ctx context.Context, entrypoint string, parts []readTestPart, state tfsdk.State) workers_script.WorkersScriptModel {
	t.Helper()
	ts := newWorkersScriptServer(t, entrypoint, parts)
	defer ts.Close()

	r := newWorkersScriptResource(t, ts.URL)
	resp := &resource.ReadResponse{State: state}
	r.Read(ctx, resource.ReadRequest{State: state}, resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("unexpected error: %v", resp.Diagnostics)
	}

	var result workers_script.WorkersScriptModel
	if diags := resp.State.Get(ctx, &result); diags.HasError() {
		t.Fatalf("failed to get state: %v", diags)
	}
	return result
}

var readTestModuleParts = []readTestPart{
	{name: "index.js", content: readTestIndexJS, contentType: "application/javascript+module"},
	{name: "util.js", content: readTestUtilJS, contentType: "application/javascript+module"},
}

func assertModuleRead(t *testing.T, result workers_script.WorkersScriptModel) {
	t.Helper()
	if got := result.MainModule.ValueString(); got != "index.js" {
		t.Errorf("main_module = %q, want %q", got, "index.js")
	}
	if got := result.Content.ValueString(); got != readTestIndexJS {
		t.Errorf("content = %q, want index.js content", got)
	}
	if result.Files == nil {
		t.Fatal("files is nil, want util.js")
	}
	files := *result.Files
	if len(files) != 1 {
		t.Errorf("files has %d entries, want 1: %v", len(files), files)
	}
	util, ok := files["util.js"]
	if !ok {
		t.Fatal("util.js missing from files")
	}
	if got := util.ContentBase64.ValueString(); got != base64.StdEncoding.EncodeToString([]byte(readTestUtilJS)) {
		t.Errorf("util.js content_base64 = %q", got)
	}
}

func assertServiceWorkerRead(t *testing.T, result workers_script.WorkersScriptModel) {
	t.Helper()
	if !result.MainModule.IsNull() {
		t.Errorf("main_module = %q, want null", result.MainModule.ValueString())
	}
	if got := result.BodyPart.ValueString(); got != "worker.js" {
		t.Errorf("body_part = %q, want %q", got, "worker.js")
	}
	if got := result.Content.ValueString(); got != readTestServiceWorkerJS {
		t.Errorf("content = %q, want worker.js content", got)
	}
}

var readTestServiceWorkerParts = []readTestPart{
	{name: "worker.js", content: readTestServiceWorkerJS, contentType: "application/javascript"},
}

func TestRead_MainModule(t *testing.T) {
	tests := []struct {
		name       string
		state      map[string]tftypes.Value
		entrypoint string
		parts      []readTestPart
		assert     func(t *testing.T, result workers_script.WorkersScriptModel)
	}{
		{
			name: "stale main_module uses entrypoint",
			state: map[string]tftypes.Value{
				"main_module": tftypes.NewValue(tftypes.String, "worker.js"),
				"content":     tftypes.NewValue(tftypes.String, "stale"),
			},
			entrypoint: "index.js",
			parts:      readTestModuleParts,
			assert:     assertModuleRead,
		},
		{
			name: "main_module matches entrypoint",
			state: map[string]tftypes.Value{
				"main_module": tftypes.NewValue(tftypes.String, "index.js"),
				"content":     tftypes.NewValue(tftypes.String, readTestIndexJS),
			},
			entrypoint: "index.js",
			parts:      readTestModuleParts,
			assert:     assertModuleRead,
		},
		{
			name: "no entrypoint uses state main_module",
			state: map[string]tftypes.Value{
				"main_module": tftypes.NewValue(tftypes.String, "index.js"),
				"content":     tftypes.NewValue(tftypes.String, readTestIndexJS),
			},
			parts:  readTestModuleParts,
			assert: assertModuleRead,
		},
		{
			name: "service worker with entrypoint stays body_part",
			state: map[string]tftypes.Value{
				"body_part": tftypes.NewValue(tftypes.String, "worker.js"),
				"content":   tftypes.NewValue(tftypes.String, readTestServiceWorkerJS),
			},
			entrypoint: "worker.js",
			parts:      readTestServiceWorkerParts,
			assert:     assertServiceWorkerRead,
		},
		{
			name: "service worker with stale body_part uses entrypoint",
			state: map[string]tftypes.Value{
				"body_part": tftypes.NewValue(tftypes.String, "script"),
				"content":   tftypes.NewValue(tftypes.String, "stale"),
			},
			entrypoint: "worker.js",
			parts:      readTestServiceWorkerParts,
			assert:     assertServiceWorkerRead,
		},
		{
			name: "service worker without entrypoint uses state body_part",
			state: map[string]tftypes.Value{
				"body_part": tftypes.NewValue(tftypes.String, "worker.js"),
				"content":   tftypes.NewValue(tftypes.String, readTestServiceWorkerJS),
			},
			parts:  readTestServiceWorkerParts,
			assert: assertServiceWorkerRead,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			state := makeWorkersScriptState(t, ctx, tt.state)
			result := readWorkersScript(t, ctx, tt.entrypoint, tt.parts, state)
			tt.assert(t, result)
		})
	}
}
