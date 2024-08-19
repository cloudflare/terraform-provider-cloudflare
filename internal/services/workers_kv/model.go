// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_kv

import (
	"bytes"
	"mime/multipart"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/apiform"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersKVResultEnvelope struct {
	Result WorkersKVModel `json:"result,computed"`
}

type WorkersKVModel struct {
	ID          types.String `tfsdk:"id" json:"-,computed"`
	KeyName     types.String `tfsdk:"key_name" path:"key_name"`
	AccountID   types.String `tfsdk:"account_id" path:"account_id"`
	NamespaceID types.String `tfsdk:"namespace_id" path:"namespace_id"`
	Metadata    types.String `tfsdk:"metadata" json:"metadata"`
	Value       types.String `tfsdk:"value" json:"value"`
}

func (r WorkersKVModel) MarshalMultipart() (data []byte, contentType string, err error) {
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
