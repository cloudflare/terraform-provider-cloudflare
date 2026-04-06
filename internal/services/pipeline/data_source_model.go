// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pipeline

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/pipelines"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PipelineResultDataSourceEnvelope struct {
	Result PipelineDataSourceModel `json:"result,computed"`
}

type PipelineDataSourceModel struct {
	ID            types.String                                                `tfsdk:"id" path:"pipeline_id,computed"`
	PipelineID    types.String                                                `tfsdk:"pipeline_id" path:"pipeline_id,required"`
	AccountID     types.String                                                `tfsdk:"account_id" path:"account_id,required"`
	CreatedAt     types.String                                                `tfsdk:"created_at" json:"created_at,computed"`
	FailureReason types.String                                                `tfsdk:"failure_reason" json:"failure_reason,computed"`
	ModifiedAt    types.String                                                `tfsdk:"modified_at" json:"modified_at,computed"`
	Name          types.String                                                `tfsdk:"name" json:"name,computed"`
	Sql           types.String                                                `tfsdk:"sql" json:"sql,computed"`
	Status        types.String                                                `tfsdk:"status" json:"status,computed"`
	Tables        customfield.NestedObjectList[PipelineTablesDataSourceModel] `tfsdk:"tables" json:"tables,computed"`
}

func (m *PipelineDataSourceModel) toReadParams(_ context.Context) (params pipelines.PipelineGetV1Params, diags diag.Diagnostics) {
	params = pipelines.PipelineGetV1Params{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type PipelineTablesDataSourceModel struct {
	ID      types.String `tfsdk:"id" json:"id,computed"`
	Latest  types.Int64  `tfsdk:"latest" json:"latest,computed"`
	Name    types.String `tfsdk:"name" json:"name,computed"`
	Type    types.String `tfsdk:"type" json:"type,computed"`
	Version types.Int64  `tfsdk:"version" json:"version,computed"`
}
