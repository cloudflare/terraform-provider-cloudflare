// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package share_recipient

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ShareRecipientResultEnvelope struct {
	Result ShareRecipientModel `json:"result"`
}

type ShareRecipientModel struct {
	ID                 types.String                                               `tfsdk:"id" json:"id,computed"`
	AccountID          types.String                                               `tfsdk:"account_id" path:"account_id,required"`
	ShareID            types.String                                               `tfsdk:"share_id" path:"share_id,required"`
	OrganizationID     types.String                                               `tfsdk:"organization_id" json:"organization_id,optional,no_refresh"`
	RecipientAccountID types.String                                               `tfsdk:"recipient_account_id" json:"recipient_account_id,optional,no_refresh"`
	AssociationStatus  types.String                                               `tfsdk:"association_status" json:"association_status,computed"`
	Created            timetypes.RFC3339                                          `tfsdk:"created" json:"created,computed" format:"date-time"`
	Modified           timetypes.RFC3339                                          `tfsdk:"modified" json:"modified,computed" format:"date-time"`
	Resources          customfield.NestedObjectList[ShareRecipientResourcesModel] `tfsdk:"resources" json:"resources,computed"`
}

func (m ShareRecipientModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ShareRecipientModel) MarshalJSONForUpdate(state ShareRecipientModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ShareRecipientResourcesModel struct {
	Error           types.String `tfsdk:"error" json:"error,computed"`
	ResourceID      types.String `tfsdk:"resource_id" json:"resource_id,computed"`
	ResourceVersion types.Int64  `tfsdk:"resource_version" json:"resource_version,computed"`
	Terminal        types.Bool   `tfsdk:"terminal" json:"terminal,computed"`
}
