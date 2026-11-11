// This file is intentionally NOT code-generated.

package snippet

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Import stores state without metadata, and the Read that follows used to
// dereference the nil Metadata and crash the provider (APIX-1183).
func TestReadAfterImportSetsMainModule(t *testing.T) {
	t.Parallel()

	const boundary = "snippet-boundary"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		switch {
		case req.Method == http.MethodGet && req.URL.Path == "/zones/test-zone/snippets/test_snippet":
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `{"success":true,"errors":[],"messages":[],"result":{"snippet_name":"test_snippet","created_on":"2026-01-01T00:00:00Z","modified_on":"2026-01-02T00:00:00Z"}}`)
		case req.Method == http.MethodGet && req.URL.Path == "/zones/test-zone/snippets/test_snippet/content":
			w.Header().Set("Content-Type", "multipart/form-data; boundary="+boundary)
			w.Header().Set("Cf-Entrypoint", "main.js")
			fmt.Fprintf(w, "--%[1]s\r\nContent-Disposition: form-data; name=\"files\"; filename=\"main.js\"\r\n\r\nexport default {}\r\n--%[1]s--\r\n", boundary)
		default:
			t.Errorf("unexpected request: %s %s", req.Method, req.URL.Path)
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	r := &SnippetResource{client: cloudflare.NewClient(
		option.WithBaseURL(server.URL),
		option.WithAPIToken("test-token"),
		option.WithMaxRetries(0),
	)}
	ctx := context.Background()

	importResp := resource.ImportStateResponse{State: tfsdk.State{Schema: ResourceSchema(ctx)}}
	r.ImportState(ctx, resource.ImportStateRequest{ID: "test-zone/test_snippet"}, &importResp)
	if importResp.Diagnostics.HasError() {
		t.Fatalf("import failed: %v", importResp.Diagnostics)
	}

	readResp := resource.ReadResponse{State: importResp.State}
	r.Read(ctx, resource.ReadRequest{State: importResp.State}, &readResp)
	if readResp.Diagnostics.HasError() {
		t.Fatalf("read failed: %v", readResp.Diagnostics)
	}

	var got SnippetModel
	if diags := readResp.State.Get(ctx, &got); diags.HasError() {
		t.Fatalf("decode state: %v", diags)
	}
	if got.Metadata == nil || !got.Metadata.MainModule.Equal(types.StringValue("main.js")) {
		t.Fatalf("metadata = %+v, want main_module main.js", got.Metadata)
	}
	if got.Files == nil || len(*got.Files) != 1 {
		t.Fatalf("files = %v, want one file", got.Files)
	}
}
