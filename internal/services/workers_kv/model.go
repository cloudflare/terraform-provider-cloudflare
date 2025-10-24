// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_kv

import (
	"bytes"
	"mime/multipart"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersKVResultEnvelope struct {
	Result WorkersKVModel `json:"result"`
}

type WorkersKVModel struct {
	ID          types.String         `tfsdk:"id" json:"-,computed"`
	KeyName     types.String         `tfsdk:"key_name" path:"key_name,required"`
	AccountID   types.String         `tfsdk:"account_id" path:"account_id,required"`
	NamespaceID types.String         `tfsdk:"namespace_id" path:"namespace_id,required"`
	Value       types.String         `tfsdk:"value" json:"value,required,no_refresh"`
	Metadata    jsontypes.Normalized `tfsdk:"metadata" json:"metadata,optional"`
}

func (r WorkersKVModel) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)

	if err := writer.WriteField("value", r.Value.ValueString()); err != nil {
		writer.Close()
		return nil, "", err
	}

	if !r.Metadata.IsNull() && !r.Metadata.IsUnknown() {
		if err := writer.WriteField("metadata", r.Metadata.ValueString()); err != nil {
			writer.Close()
			return nil, "", err
		}
	}

	if err := writer.Close(); err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}
