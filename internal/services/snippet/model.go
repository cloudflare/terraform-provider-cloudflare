// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package snippet

import (
	"bytes"
	"errors"
	"mime/multipart"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/apiform"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SnippetResultEnvelope struct {
	Result SnippetModel `json:"result"`
}

type SnippetModel struct {
	SnippetName types.String          `tfsdk:"snippet_name" path:"snippet_name,required"`
	ZoneID      types.String          `tfsdk:"zone_id" path:"zone_id,required"`
	Metadata    *SnippetMetadataModel `tfsdk:"metadata" json:"metadata,required,no_refresh"`
	CreatedOn   timetypes.RFC3339     `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	ModifiedOn  timetypes.RFC3339     `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
}

func (r SnippetModel) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		if e := writer.Close(); e != nil {
			err = errors.Join(err, e)
		}
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
