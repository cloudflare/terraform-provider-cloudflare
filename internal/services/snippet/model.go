// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package snippet

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"strings"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/apiform"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

type SnippetResultEnvelope struct {
	Result SnippetModel `json:"result"`
}

var SnippetFileType = snippetFileType{
	ObjectType: types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":    types.StringType,
			"content": types.StringType,
		},
	},
}

type SnippetModel struct {
	SnippetName types.String          `tfsdk:"snippet_name" path:"snippet_name,required"`
	ZoneID      types.String          `tfsdk:"zone_id" path:"zone_id,required"`
	Files       *[]SnippetFile        `tfsdk:"files" json:"files,metadata,required"`
	Metadata    *SnippetMetadataModel `tfsdk:"metadata" json:"metadata,required,no_refresh"`
	CreatedOn   timetypes.RFC3339     `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	ModifiedOn  timetypes.RFC3339     `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
}

func (r SnippetModel) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

type SnippetMetadataModel struct {
	MainModule types.String `tfsdk:"main_module" json:"main_module,required"`
}

func (r *SnippetModel) UnmarshalMultipart(data []byte, contentType string) error {
	mediaType, params, err := mime.ParseMediaType(contentType)
	if err != nil {
		return fmt.Errorf("failed to parse media type: %w", err)
	}
	if mediaType != "multipart/form-data" {
		return fmt.Errorf("expected media type %q, got %q", "multipart/form-data", mediaType)
	}
	reader := multipart.NewReader(bytes.NewReader(data), params["boundary"])
	var files []SnippetFile
	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to get multipart part: %w", err)
		}
		if part.FormName() == "files" {
			bytes, err := io.ReadAll(part)
			if err != nil {
				return fmt.Errorf("failed to read multipart part: %w", err)
			}
			files = append(files, NewSnippetsFileValueMust(
				part.FileName(),
				string(bytes),
			))
		}
	}
	r.Files = &files
	return nil
}

type snippetFileType struct {
	types.ObjectType
}

func (t snippetFileType) Equal(other attr.Type) bool {
	_, ok := other.(snippetFileType)

	return ok
}

func (t snippetFileType) String() string {
	return "SnippetsFileContentType"
}

func (t snippetFileType) ValueFromTerraform(
	ctx context.Context,
	in tftypes.Value,
) (attr.Value, error) {
	val, err := t.ObjectType.ValueFromTerraform(ctx, in)
	if err != nil {
		return nil, err
	}

	obj, ok := val.(types.Object)
	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", val)
	}

	return SnippetFile{obj, new(int64)}, nil
}

func (t snippetFileType) ValueType(_ context.Context) attr.Value {
	return SnippetFile{}
}

func (t snippetFileType) ValueFromObject(
	_ context.Context,
	obj basetypes.ObjectValue,
) (basetypes.ObjectValuable, diag.Diagnostics) {
	return SnippetFile{obj, new(int64)}, nil
}

type SnippetFile struct {
	types.Object
	offset *int64
}

func NewSnippetsFileValueMust(name string, content string) SnippetFile {
	return SnippetFile{types.ObjectValueMust(
		SnippetFileType.AttrTypes,
		map[string]attr.Value{
			"name":    types.StringValue(name),
			"content": types.StringValue(content),
		},
	), new(int64)}
}

func (f SnippetFile) Type(_ context.Context) attr.Type {
	return SnippetFileType
}

func (f SnippetFile) Equal(other attr.Value) bool {
	o, ok := other.(SnippetFile)
	if !ok {
		return false
	}

	return f.Object.Equal(o.Object)
}

func (f SnippetFile) Name() string {
	return f.Object.Attributes()["name"].(types.String).ValueString()
}

func (f SnippetFile) ContentType() string {
	return "application/javascript+module"
}

func (f SnippetFile) Read(p []byte) (n int, err error) {
	content := f.Object.Attributes()["content"].(types.String).ValueString()

	reader := strings.NewReader(content)

	if _, err := reader.Seek(*f.offset, io.SeekStart); err != nil {
		return 0, err
	}

	n, err = reader.Read(p)

	*f.offset += int64(n)

	return n, err
}
