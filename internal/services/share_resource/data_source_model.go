// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package share_resource

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/resource_sharing"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ShareResourceResultDataSourceEnvelope struct {
	Result ShareResourceDataSourceModel `json:"result,computed"`
}

type ShareResourceDataSourceModel struct {
	ID                types.String                           `tfsdk:"id" path:"share_resource_id,computed"`
	ShareResourceID   types.String                           `tfsdk:"share_resource_id" path:"share_resource_id,optional"`
	AccountID         types.String                           `tfsdk:"account_id" path:"account_id,required"`
	ShareID           types.String                           `tfsdk:"share_id" path:"share_id,required"`
	Created           timetypes.RFC3339                      `tfsdk:"created" json:"created,computed" format:"date-time"`
	Modified          timetypes.RFC3339                      `tfsdk:"modified" json:"modified,computed" format:"date-time"`
	ResourceAccountID types.String                           `tfsdk:"resource_account_id" json:"resource_account_id,computed"`
	ResourceID        types.String                           `tfsdk:"resource_id" json:"resource_id,computed"`
	ResourceType      types.String                           `tfsdk:"resource_type" json:"resource_type,computed"`
	ResourceVersion   types.Int64                            `tfsdk:"resource_version" json:"resource_version,computed"`
	Status            types.String                           `tfsdk:"status" json:"status,computed"`
	Meta              jsontypes.Normalized                   `tfsdk:"meta" json:"meta,computed"`
	Filter            *ShareResourceFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *ShareResourceDataSourceModel) toReadParams(_ context.Context) (params resource_sharing.ResourceGetParams, diags diag.Diagnostics) {
	params = resource_sharing.ResourceGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *ShareResourceDataSourceModel) toListParams(_ context.Context) (params resource_sharing.ResourceListParams, diags diag.Diagnostics) {
	params = resource_sharing.ResourceListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Filter.ResourceType.IsNull() {
		params.ResourceType = cloudflare.F(resource_sharing.ResourceListParamsResourceType(m.Filter.ResourceType.ValueString()))
	}
	if !m.Filter.Status.IsNull() {
		params.Status = cloudflare.F(resource_sharing.ResourceListParamsStatus(m.Filter.Status.ValueString()))
	}

	return
}

type ShareResourceFindOneByDataSourceModel struct {
	ResourceType types.String `tfsdk:"resource_type" query:"resource_type,optional"`
	Status       types.String `tfsdk:"status" query:"status,optional"`
}
