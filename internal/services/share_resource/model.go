// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package share_resource

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ShareResourceResultEnvelope struct {
	Result ShareResourceModel `json:"result"`
}

type ShareResourceModel struct {
	ID                types.String         `tfsdk:"id" json:"id,computed"`
	AccountID         types.String         `tfsdk:"account_id" path:"account_id,required"`
	ShareID           types.String         `tfsdk:"share_id" path:"share_id,required"`
	ResourceAccountID types.String         `tfsdk:"resource_account_id" json:"resource_account_id,required"`
	ResourceID        types.String         `tfsdk:"resource_id" json:"resource_id,required"`
	ResourceType      types.String         `tfsdk:"resource_type" json:"resource_type,required"`
	Meta              jsontypes.Normalized `tfsdk:"meta" json:"meta,required"`
	Created           timetypes.RFC3339    `tfsdk:"created" json:"created,computed" format:"date-time"`
	Modified          timetypes.RFC3339    `tfsdk:"modified" json:"modified,computed" format:"date-time"`
	ResourceVersion   types.Int64          `tfsdk:"resource_version" json:"resource_version,computed"`
	Status            types.String         `tfsdk:"status" json:"status,computed"`
}

func (m ShareResourceModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ShareResourceModel) MarshalJSONForUpdate(state ShareResourceModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
