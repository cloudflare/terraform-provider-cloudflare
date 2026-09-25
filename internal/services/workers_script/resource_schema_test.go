// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_script_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/workers_script"
	resource_schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func TestWorkersScriptSchemaSupportsMultipartFiles(t *testing.T) {
	files, ok := workers_script.ResourceSchema(context.Background()).Attributes["files"]
	if !ok {
		t.Fatal("cloudflare_workers_script has no files attribute for WASM multipart uploads")
	}
	if _, ok := files.(resource_schema.MapNestedAttribute); !ok {
		t.Fatalf("files attribute has type %T, want schema.MapNestedAttribute", files)
	}
}

func TestWorkersScriptModelSchemaParity(t *testing.T) {
	diagnostics := workers_script.ResourceSchema(context.Background()).ValidateImplementation(context.Background())
	if diagnostics.HasError() {
		t.Fatalf("invalid resource schema: %v", diagnostics)
	}
}
