// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package snippet

import (
	"bytes"
	"mime/multipart"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/apiform"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SnippetResultEnvelope struct {
	Result SnippetModel `json:"result"`
}

type SnippetModel struct {
	SnippetName types.String                                   `tfsdk:"snippet_name" path:"snippet_name,required"`
	ZoneID      types.String                                   `tfsdk:"zone_id" path:"zone_id,required"`
	Files       types.String                                   `tfsdk:"files" json:"files,optional"`
	Metadata    customfield.NestedObject[SnippetMetadataModel] `tfsdk:"metadata" json:"metadata,computed_optional"`
	CreatedOn   types.String                                   `tfsdk:"created_on" json:"created_on,computed"`
	ModifiedOn  types.String                                   `tfsdk:"modified_on" json:"modified_on,computed"`
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
	MainModule types.String `tfsdk:"main_module" json:"main_module,optional"`
}
