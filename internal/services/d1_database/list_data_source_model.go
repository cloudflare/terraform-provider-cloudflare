// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package d1_database

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/d1"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type D1DatabasesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[D1DatabasesResultDataSourceModel] `json:"result,computed"`
}

type D1DatabasesDataSourceModel struct {
	AccountID types.String                                                   `tfsdk:"account_id" path:"account_id,required"`
	Name      types.String                                                   `tfsdk:"name" query:"name,optional"`
	MaxItems  types.Int64                                                    `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[D1DatabasesResultDataSourceModel] `tfsdk:"result"`
}

func (m *D1DatabasesDataSourceModel) toListParams(_ context.Context) (params d1.DatabaseListParams, diags diag.Diagnostics) {
	params = d1.DatabaseListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Name.IsNull() {
		params.Name = cloudflare.F(m.Name.ValueString())
	}

	return
}

type D1DatabasesResultDataSourceModel struct {
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Name      types.String      `tfsdk:"name" json:"name,computed"`
	UUID      types.String      `tfsdk:"uuid" json:"uuid,computed"`
	Version   types.String      `tfsdk:"version" json:"version,computed"`
}
