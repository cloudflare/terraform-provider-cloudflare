// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package record

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RecordResultEnvelope struct {
	Result RecordModel `json:"result"`
}

type RecordModel struct {
	ID             types.String `tfsdk:"id" json:"id,computed"`
	ZoneID         types.String `tfsdk:"zone_id" path:"zone_id,required"`
	Name           types.String `tfsdk:"name" json:"name,computed"`
	Hostname       types.String `tfsdk:"hostname" json:"hostname,computed"`
	Type           types.String `tfsdk:"type" json:"type,computed"`
	Value          types.String `tfsdk:"value" json:"value,computed"`
	Content        types.String `tfsdk:"content" json:"content,computed"`
	Data           types.List   `tfsdk:"data" json:"data,computed"`
	TTL            types.Int64  `tfsdk:"ttl" json:"ttl,computed"`
	Priority       types.Int64  `tfsdk:"priority" json:"priority,computed"`
	Proxied        types.Bool   `tfsdk:"proxied" json:"proxied,computed"`
	CreatedOn      types.String `tfsdk:"created_on" json:"created_on,computed"`
	Metadata       types.Map    `tfsdk:"metadata" json:"metadata,computed"`
	ModifiedOn     types.String `tfsdk:"modified_on" json:"modified_on,computed"`
	Proxiable      types.Bool   `tfsdk:"proxiable" json:"proxiable,computed"`
	AllowOverwrite types.Bool   `tfsdk:"allow_overwrite" json:"allow_overwrite,computed"`
	Comment        types.String `tfsdk:"comment" json:"comment,computed"`
	Tags           types.Set    `tfsdk:"tags" json:"tags,computed"`
}

func (m RecordModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m RecordModel) MarshalJSONForUpdate(state RecordModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
