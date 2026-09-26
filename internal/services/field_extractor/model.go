// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package field_extractor

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type FieldExtractorResultEnvelope struct {
	Result FieldExtractorModel `json:"result"`
}

type FieldExtractorModel struct {
	ID        types.String                 `tfsdk:"id" json:"-,computed"`
	Extractor types.String                 `tfsdk:"extractor" path:"extractor,required"`
	AccountID types.String                 `tfsdk:"account_id" path:"account_id,required"`
	Rules     *[]*FieldExtractorRulesModel `tfsdk:"rules" json:"rules,required"`
}

func (m FieldExtractorModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m FieldExtractorModel) MarshalJSONForUpdate(state FieldExtractorModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type FieldExtractorRulesModel struct {
	Fields      *[]*FieldExtractorRulesFieldsModel `tfsdk:"fields" json:"fields,required"`
	Ref         types.String                       `tfsdk:"ref" json:"ref,required"`
	Description types.String                       `tfsdk:"description" json:"description,optional"`
}

type FieldExtractorRulesFieldsModel struct {
	Expression types.String `tfsdk:"expression" json:"expression,required"`
	Name       types.String `tfsdk:"name" json:"name,required"`
}
