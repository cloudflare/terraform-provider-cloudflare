package workers_script

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"io"
	"mime"
	"mime/multipart"
	"os"
	"path/filepath"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestMarshalMultipartIncludesWASMWithDurableObjectMigration(t *testing.T) {
	ctx := context.Background()
	wasm := []byte("\x00asm\x01\x00\x00\x00")
	files := map[string]WorkersScriptFileModel{
		"module.wasm": {
			ContentType:   types.StringValue("application/wasm"),
			ContentBase64: types.StringValue(base64.StdEncoding.EncodeToString(wasm)),
		},
	}
	newClasses := []types.String{types.StringValue("Test")}
	migrations, diags := customfield.NewObject(ctx, &WorkersScriptMetadataMigrationsModel{
		NewSqliteClasses: &newClasses,
		NewTag:           types.StringValue("v1"),
	})
	if diags.HasError() {
		t.Fatalf("creating migrations: %v", diags)
	}
	bindings, diags := customfield.NewObjectList(ctx, []WorkersScriptMetadataBindingsModel{
		{
			Name:      types.StringValue("TEST"),
			Type:      types.StringValue("durable_object_namespace"),
			ClassName: types.StringValue("Test"),
		},
		{
			Name: types.StringValue("ADD_WASM"),
			Type: types.StringValue("wasm_module"),
			Part: types.StringValue("module.wasm"),
		},
	})
	if diags.HasError() {
		t.Fatalf("creating bindings: %v", diags)
	}
	model := WorkersScriptModel{
		Content:     types.StringValue("addEventListener('fetch', () => {})"),
		ContentType: types.StringValue("application/javascript"),
		Files:       &files,
		WorkersScriptMetadataModel: WorkersScriptMetadataModel{
			Bindings:   bindings,
			Migrations: migrations,
		},
	}

	body, contentType, err := model.MarshalMultipart()
	if err != nil {
		t.Fatal(err)
	}
	_, params, err := mime.ParseMediaType(contentType)
	if err != nil {
		t.Fatal(err)
	}

	parts := make(map[string][]byte)
	partTypes := make(map[string]string)
	reader := multipart.NewReader(bytes.NewReader(body), params["boundary"])
	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatal(err)
		}
		content, err := io.ReadAll(part)
		if err != nil {
			t.Fatal(err)
		}
		parts[part.FormName()] = content
		partTypes[part.FormName()] = part.Header.Get("Content-Type")
	}

	for name, expected := range map[string][]byte{
		"script":      []byte("addEventListener('fetch', () => {})"),
		"module.wasm": wasm,
	} {
		if !bytes.Equal(parts[name], expected) {
			t.Errorf("part %q = %q, want %q", name, parts[name], expected)
		}
	}
	if partTypes["module.wasm"] != "application/wasm" {
		t.Errorf("WASM content type = %q", partTypes["module.wasm"])
	}
	if _, ok := parts["metadata"]; !ok {
		t.Error("metadata part is missing")
	}
	var metadata struct {
		Bindings []struct {
			Name string `json:"name"`
			Part string `json:"part"`
			Type string `json:"type"`
		} `json:"bindings"`
		Migrations struct {
			NewSqliteClasses []string `json:"new_sqlite_classes"`
			NewTag           string   `json:"new_tag"`
		} `json:"migrations"`
	}
	if err := json.Unmarshal(parts["metadata"], &metadata); err != nil {
		t.Fatal(err)
	}
	if metadata.Migrations.NewTag != "v1" || len(metadata.Migrations.NewSqliteClasses) != 1 || metadata.Migrations.NewSqliteClasses[0] != "Test" {
		t.Fatalf("Durable Object migration missing from metadata: %#v", metadata.Migrations)
	}
	if len(metadata.Bindings) != 2 || metadata.Bindings[1].Type != "wasm_module" || metadata.Bindings[1].Part != "module.wasm" {
		t.Fatalf("WASM binding missing from metadata: %#v", metadata.Bindings)
	}
}

func TestMarshalMultipartRejectsInvalidFileBase64(t *testing.T) {
	files := map[string]WorkersScriptFileModel{"module.wasm": {
		ContentType:   types.StringValue("application/wasm"),
		ContentBase64: types.StringValue("not base64"),
	}}
	_, _, err := (WorkersScriptModel{Files: &files}).MarshalMultipart()
	if err == nil {
		t.Fatal("expected invalid base64 error")
	}
}

func TestPrepareFileUploadsReadsBinaryFileWithoutChangingPlan(t *testing.T) {
	content := []byte("\x00asm\x01\x00\x00\x00")
	path := filepath.Join(t.TempDir(), "module.wasm")
	if err := os.WriteFile(path, content, 0o600); err != nil {
		t.Fatal(err)
	}
	files := map[string]WorkersScriptFileModel{"module.wasm": {
		ContentFile: types.StringValue(path),
		ContentType: types.StringValue("application/wasm"),
	}}
	model := WorkersScriptModel{Files: &files}

	planFiles, err := prepareFileUploads(&model)
	if err != nil {
		t.Fatal(err)
	}
	if !(*planFiles)["module.wasm"].ContentBase64.IsNull() {
		t.Fatal("prepared upload changed planned content_base64")
	}
	decoded, err := base64.StdEncoding.DecodeString((*model.Files)["module.wasm"].ContentBase64.ValueString())
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(decoded, content) {
		t.Fatalf("uploaded content = %q, want %q", decoded, content)
	}
}

func TestPrepareFileUploadsRejectsInvalidNames(t *testing.T) {
	for _, name := range []string{"", "metadata", "bad\r\nname"} {
		t.Run(name, func(t *testing.T) {
			files := map[string]WorkersScriptFileModel{name: {
				ContentBase64: types.StringValue(""),
				ContentType:   types.StringValue("text/plain"),
			}}
			_, err := prepareFileUploads(&WorkersScriptModel{Files: &files})
			if err == nil {
				t.Fatal("expected invalid name error")
			}
		})
	}
}
