// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package schema_validation_schemas

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SchemaValidationSchemasResultEnvelope struct {
	Result SchemaValidationSchemasModel `json:"result"`
}

type SchemaValidationSchemasModel struct {
	ID                types.String      `tfsdk:"id" json:"-,computed"`
	SchemaID          types.String      `tfsdk:"schema_id" json:"schema_id,computed"`
	ZoneID            types.String      `tfsdk:"zone_id" path:"zone_id,required"`
	Kind              types.String      `tfsdk:"kind" json:"kind,required"`
	Name              types.String      `tfsdk:"name" json:"name,required"`
	Source            types.String      `tfsdk:"source" json:"source,required"`
	ValidationEnabled types.Bool        `tfsdk:"validation_enabled" json:"validation_enabled,optional"`
	CreatedAt         timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
}

func (m SchemaValidationSchemasModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m SchemaValidationSchemasModel) MarshalJSONForUpdate(state SchemaValidationSchemasModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}
