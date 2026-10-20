package snippet

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

const (
	readTestZoneID      = "023e105f4ecef8ad9ca31a8372d0c353"
	readTestSnippetName = "my_snippet"
	readTestCode        = "export default { async fetch(request) { return fetch(request); } };"
)

// newReadTestServer fakes the two endpoints Read calls: the snippet metadata (JSON, which
// carries no main module) and its content (multipart, with the entrypoint in Cf-Entrypoint).
func newReadTestServer(t *testing.T) *httptest.Server {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case fmt.Sprintf("/zones/%s/snippets/%s", readTestZoneID, readTestSnippetName):
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, `{"success":true,"errors":[],"messages":[],"result":{"snippet_name":%q,"created_on":"2026-01-01T00:00:00Z","modified_on":"2026-01-01T00:00:00Z"}}`, readTestSnippetName)
		case fmt.Sprintf("/zones/%s/snippets/%s/content", readTestZoneID, readTestSnippetName):
			const boundary = "readtestboundary"
			w.Header().Set("Content-Type", "multipart/form-data; boundary="+boundary)
			w.Header().Set("Cf-Entrypoint", "snippet.js")
			fmt.Fprint(w, strings.Join([]string{
				"--" + boundary,
				`Content-Disposition: form-data; name="files"; filename="snippet.js"`,
				"Content-Type: application/javascript+module",
				"",
				readTestCode,
				"--" + boundary + "--",
				"",
			}, "\r\n"))
		default:
			http.NotFound(w, r)
		}
	}))
	t.Cleanup(srv.Close)
	return srv
}

func readSnippet(t *testing.T, prior SnippetModel) SnippetModel {
	t.Helper()
	ctx := context.Background()
	srv := newReadTestServer(t)
	r := &SnippetResource{client: cloudflare.NewClient(option.WithBaseURL(srv.URL), option.WithAPIToken("test"))}
	sch := ResourceSchema(ctx)

	state := tfsdk.State{Schema: sch, Raw: tftypes.NewValue(sch.Type().TerraformType(ctx), nil)}
	if diags := state.Set(ctx, &prior); diags.HasError() {
		t.Fatalf("setting prior state: %v", diags)
	}
	resp := &resource.ReadResponse{State: tfsdk.State{Schema: sch, Raw: state.Raw.Copy()}}

	r.Read(ctx, resource.ReadRequest{State: state}, resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("Read returned errors: %v", resp.Diagnostics)
	}
	var got SnippetModel
	if diags := resp.State.Get(ctx, &got); diags.HasError() {
		t.Fatalf("reading new state: %v", diags)
	}
	return got
}

func TestSnippetRead_PriorStateWithMetadata(t *testing.T) {
	files := []SnippetFile{NewSnippetsFileValueMust("snippet.js", readTestCode)}
	got := readSnippet(t, SnippetModel{
		ID:          types.StringValue(readTestSnippetName),
		SnippetName: types.StringValue(readTestSnippetName),
		ZoneID:      types.StringValue(readTestZoneID),
		Files:       &files,
		Metadata:    &SnippetMetadataModel{MainModule: types.StringValue("snippet.js")},
	})
	if got.Metadata == nil || got.Metadata.MainModule.ValueString() != "snippet.js" {
		t.Fatalf("expected metadata.main_module %q, got %+v", "snippet.js", got.Metadata)
	}
}

// State produced by ImportState has no metadata: the snippet GET response doesn't include it
// and the attribute is no_refresh. Read must not dereference the nil pointer
// (https://github.com/cloudflare/terraform-provider-cloudflare/issues/7378).
func TestSnippetRead_PriorStateWithoutMetadata(t *testing.T) {
	got := readSnippet(t, SnippetModel{
		ID:          types.StringValue(readTestSnippetName),
		SnippetName: types.StringValue(readTestSnippetName),
		ZoneID:      types.StringValue(readTestZoneID),
	})
	if got.Metadata == nil || got.Metadata.MainModule.ValueString() != "snippet.js" {
		t.Fatalf("expected metadata.main_module %q from Cf-Entrypoint, got %+v", "snippet.js", got.Metadata)
	}
	if got.Files == nil || len(*got.Files) != 1 || (*got.Files)[0].Name() != "snippet.js" {
		t.Fatalf("expected files to be read from the content endpoint, got %+v", got.Files)
	}
}
