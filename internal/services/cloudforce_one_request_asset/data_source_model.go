// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one_request_asset

import (
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type CloudforceOneRequestAssetResultDataSourceEnvelope struct {
Result CloudforceOneRequestAssetDataSourceModel `json:"result,computed"`
}

type CloudforceOneRequestAssetDataSourceModel struct {
AccountIdentifier types.String `tfsdk:"account_identifier" path:"account_identifier,required"`
AssetIdentifer types.String `tfsdk:"asset_identifer" path:"asset_identifer,required"`
RequestIdentifier types.String `tfsdk:"request_identifier" path:"request_identifier,required"`
Created timetypes.RFC3339 `tfsdk:"created" json:"created,computed" format:"date-time"`
Description types.String `tfsdk:"description" json:"description,computed"`
FileType types.String `tfsdk:"file_type" json:"file_type,computed"`
ID types.Int64 `tfsdk:"id" json:"id,computed"`
Name types.String `tfsdk:"name" json:"name,computed"`
}
