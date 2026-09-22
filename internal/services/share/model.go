// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package share

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ShareResultEnvelope struct {
	Result ShareModel `json:"result"`
}

type ShareModel struct {
	ID                           types.String             `tfsdk:"id" json:"id,computed"`
	AccountID                    types.String             `tfsdk:"account_id" path:"account_id,required"`
	Recipients                   *[]*ShareRecipientsModel `tfsdk:"recipients" json:"recipients,required,no_refresh"`
	Resources                    *[]*ShareResourcesModel  `tfsdk:"resources" json:"resources,required"`
	Name                         types.String             `tfsdk:"name" json:"name,required"`
	AccountName                  types.String             `tfsdk:"account_name" json:"account_name,computed"`
	AssociatedRecipientCount     types.Int64              `tfsdk:"associated_recipient_count" json:"associated_recipient_count,computed"`
	AssociatingRecipientCount    types.Int64              `tfsdk:"associating_recipient_count" json:"associating_recipient_count,computed"`
	Created                      timetypes.RFC3339        `tfsdk:"created" json:"created,computed" format:"date-time"`
	DisassociatedRecipientCount  types.Int64              `tfsdk:"disassociated_recipient_count" json:"disassociated_recipient_count,computed"`
	DisassociatingRecipientCount types.Int64              `tfsdk:"disassociating_recipient_count" json:"disassociating_recipient_count,computed"`
	Kind                         types.String             `tfsdk:"kind" json:"kind,computed"`
	Modified                     timetypes.RFC3339        `tfsdk:"modified" json:"modified,computed" format:"date-time"`
	OrganizationID               types.String             `tfsdk:"organization_id" json:"organization_id,computed"`
	Status                       types.String             `tfsdk:"status" json:"status,computed"`
	TargetType                   types.String             `tfsdk:"target_type" json:"target_type,computed"`
}

func (m ShareModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ShareModel) MarshalJSONForUpdate(state ShareModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ShareRecipientsModel struct {
	OrganizationID     types.String `tfsdk:"organization_id" json:"organization_id,optional"`
	RecipientAccountID types.String `tfsdk:"recipient_account_id" json:"recipient_account_id,optional"`
}

type ShareResourcesModel struct {
	Meta              jsontypes.Normalized `tfsdk:"meta" json:"meta,required"`
	ResourceAccountID types.String         `tfsdk:"resource_account_id" json:"resource_account_id,required"`
	ResourceID        types.String         `tfsdk:"resource_id" json:"resource_id,required"`
	ResourceType      types.String         `tfsdk:"resource_type" json:"resource_type,required"`
}
