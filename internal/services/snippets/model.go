// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package snippets

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

var SnippetsFileType = snippetsFileType{
	ObjectType: types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":    types.StringType,
			"content": types.StringType,
		},
	},
}

type SnippetsResultEnvelope struct {
	Result SnippetsModel `json:"result"`
}

type SnippetsModel struct {
	SnippetName types.String           `tfsdk:"snippet_name" path:"snippet_name,required"`
	ZoneID      types.String           `tfsdk:"zone_id" path:"zone_id,required"`
	Files       *[]SnippetsFile        `tfsdk:"files" json:"files,metadata,required"`
	Metadata    *SnippetsMetadataModel `tfsdk:"metadata" json:"metadata,metadata,required"`
	CreatedOn   timetypes.RFC3339      `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	ModifiedOn  timetypes.RFC3339      `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
}

func (r SnippetsModel) MarshalMultipart() (data []byte, contentType string, err error) {
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

func (r *SnippetsModel) UnmarshalMultipart(data []byte, contentType string) error {
	mediaType, params, err := mime.ParseMediaType(contentType)
	if err != nil {
		return fmt.Errorf("failed to parse media type: %w", err)
	}
	if mediaType != "multipart/form-data" {
		return fmt.Errorf("expected media type %q, got %q", "multipart/form-data", mediaType)
	}
	reader := multipart.NewReader(bytes.NewReader(data), params["boundary"])
	var files []SnippetsFile
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

type snippetsFileType struct {
	types.ObjectType
}

func (t snippetsFileType) Equal(other attr.Type) bool {
	_, ok := other.(snippetsFileType)

	return ok
}

func (t snippetsFileType) String() string {
	return "SnippetsFileContentType"
}

func (t snippetsFileType) ValueFromTerraform(
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

	return SnippetsFile{obj, new(int64)}, nil
}

func (t snippetsFileType) ValueType(_ context.Context) attr.Value {
	return SnippetsFile{}
}

func (t snippetsFileType) ValueFromObject(
	_ context.Context,
	obj basetypes.ObjectValue,
) (basetypes.ObjectValuable, diag.Diagnostics) {
	return SnippetsFile{obj, new(int64)}, nil
}

type SnippetsFile struct {
	types.Object
	offset *int64
}

func NewSnippetsFileValueMust(name string, content string) SnippetsFile {
	return SnippetsFile{types.ObjectValueMust(
		SnippetsFileType.AttrTypes,
		map[string]attr.Value{
			"name":    types.StringValue(name),
			"content": types.StringValue(content),
		},
	), new(int64)}
}

func (f SnippetsFile) Type(_ context.Context) attr.Type {
	return SnippetsFileType
}

func (f SnippetsFile) Equal(other attr.Value) bool {
	o, ok := other.(SnippetsFile)
	if !ok {
		return false
	}

	return f.Object.Equal(o.Object)
}

func (f SnippetsFile) Name() string {
	return f.Object.Attributes()["name"].(types.String).ValueString()
}

func (f SnippetsFile) ContentType() string {
	return "application/javascript+module"
}

func (f SnippetsFile) Read(p []byte) (n int, err error) {
	content := f.Object.Attributes()["content"].(types.String).ValueString()

	reader := strings.NewReader(content)

	if _, err := reader.Seek(*f.offset, io.SeekStart); err != nil {
		return 0, err
	}

	n, err = reader.Read(p)

	*f.offset += int64(n)

	return n, err
}

type SnippetsMetadataModel struct {
	MainModule types.String `tfsdk:"main_module" json:"main_module,required"`
}
