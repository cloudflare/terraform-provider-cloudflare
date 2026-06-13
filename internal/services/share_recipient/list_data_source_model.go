// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package share_recipient

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/resource_sharing"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ShareRecipientsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ShareRecipientsResultDataSourceModel] `json:"result,computed"`
}

type ShareRecipientsDataSourceModel struct {
	AccountID        types.String                                                       `tfsdk:"account_id" path:"account_id,required"`
	ShareID          types.String                                                       `tfsdk:"share_id" path:"share_id,required"`
	IncludeResources types.Bool                                                         `tfsdk:"include_resources" query:"include_resources,optional"`
	MaxItems         types.Int64                                                        `tfsdk:"max_items"`
	Result           customfield.NestedObjectList[ShareRecipientsResultDataSourceModel] `tfsdk:"result"`
}

func (m *ShareRecipientsDataSourceModel) toListParams(_ context.Context) (params resource_sharing.RecipientListParams, diags diag.Diagnostics) {
	params = resource_sharing.RecipientListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.IncludeResources.IsNull() {
		params.IncludeResources = cloudflare.F(m.IncludeResources.ValueBool())
	}

	return
}

type ShareRecipientsResultDataSourceModel struct {
	ID                types.String                                                          `tfsdk:"id" json:"id,computed"`
	AccountID         types.String                                                          `tfsdk:"account_id" json:"account_id,computed"`
	AssociationStatus types.String                                                          `tfsdk:"association_status" json:"association_status,computed"`
	Created           timetypes.RFC3339                                                     `tfsdk:"created" json:"created,computed" format:"date-time"`
	Modified          timetypes.RFC3339                                                     `tfsdk:"modified" json:"modified,computed" format:"date-time"`
	Resources         customfield.NestedObjectList[ShareRecipientsResourcesDataSourceModel] `tfsdk:"resources" json:"resources,computed"`
}

type ShareRecipientsResourcesDataSourceModel struct {
	Error           types.String `tfsdk:"error" json:"error,computed"`
	ResourceID      types.String `tfsdk:"resource_id" json:"resource_id,computed"`
	ResourceVersion types.Int64  `tfsdk:"resource_version" json:"resource_version,computed"`
	Terminal        types.Bool   `tfsdk:"terminal" json:"terminal,computed"`
}
