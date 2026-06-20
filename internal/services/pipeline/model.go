// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pipeline

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PipelineResultEnvelope struct {
	Result PipelineModel `json:"result"`
}

type PipelineModel struct {
	ID            types.String                                      `tfsdk:"id" json:"id,computed"`
	AccountID     types.String                                      `tfsdk:"account_id" path:"account_id,required"`
	Name          types.String                                      `tfsdk:"name" json:"name,required"`
	Sql           types.String                                      `tfsdk:"sql" json:"sql,required"`
	CreatedAt     types.String                                      `tfsdk:"created_at" json:"created_at,computed"`
	FailureReason types.String                                      `tfsdk:"failure_reason" json:"failure_reason,computed"`
	ModifiedAt    types.String                                      `tfsdk:"modified_at" json:"modified_at,computed"`
	Status        types.String                                      `tfsdk:"status" json:"status,computed"`
	Tables        customfield.NestedObjectList[PipelineTablesModel] `tfsdk:"tables" json:"tables,computed"`
}

func (m PipelineModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m PipelineModel) MarshalJSONForUpdate(state PipelineModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type PipelineTablesModel struct {
	ID      types.String `tfsdk:"id" json:"id,computed"`
	Latest  types.Int64  `tfsdk:"latest" json:"latest,computed"`
	Name    types.String `tfsdk:"name" json:"name,computed"`
	Type    types.String `tfsdk:"type" json:"type,computed"`
	Version types.Int64  `tfsdk:"version" json:"version,computed"`
}
