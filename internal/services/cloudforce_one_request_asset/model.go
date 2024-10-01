// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one_request_asset

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CloudforceOneRequestAssetResultEnvelope struct {
	Result CloudforceOneRequestAssetModel `json:"result"`
}

type CloudforceOneRequestAssetModel struct {
	AccountIdentifier types.String                                                         `tfsdk:"account_identifier" path:"account_identifier,required"`
	RequestIdentifier types.String                                                         `tfsdk:"request_identifier" path:"request_identifier,required"`
	AssetIdentifer    types.String                                                         `tfsdk:"asset_identifer" path:"asset_identifer,optional"`
	Page              types.Int64                                                          `tfsdk:"page" json:"page,required"`
	PerPage           types.Int64                                                          `tfsdk:"per_page" json:"per_page,required"`
	Source            types.String                                                         `tfsdk:"source" json:"source,optional"`
	Created           timetypes.RFC3339                                                    `tfsdk:"created" json:"created,computed" format:"date-time"`
	Description       types.String                                                         `tfsdk:"description" json:"description,computed"`
	FileType          types.String                                                         `tfsdk:"file_type" json:"file_type,computed"`
	ID                types.Int64                                                          `tfsdk:"id" json:"id,computed"`
	Name              types.String                                                         `tfsdk:"name" json:"name,computed"`
	Success           types.Bool                                                           `tfsdk:"success" json:"success,computed"`
	Errors            customfield.NestedObjectList[CloudforceOneRequestAssetErrorsModel]   `tfsdk:"errors" json:"errors,computed"`
	Messages          customfield.NestedObjectList[CloudforceOneRequestAssetMessagesModel] `tfsdk:"messages" json:"messages,computed"`
}

func (m CloudforceOneRequestAssetModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m CloudforceOneRequestAssetModel) MarshalJSONForUpdate(state CloudforceOneRequestAssetModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type CloudforceOneRequestAssetErrorsModel struct {
	Code    types.Int64  `tfsdk:"code" json:"code,computed"`
	Message types.String `tfsdk:"message" json:"message,computed"`
}

type CloudforceOneRequestAssetMessagesModel struct {
	Code    types.Int64  `tfsdk:"code" json:"code,computed"`
	Message types.String `tfsdk:"message" json:"message,computed"`
}
