// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package d1_database

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type D1DatabaseResultEnvelope struct {
	Result D1DatabaseModel `json:"result,computed"`
}

type D1DatabaseModel struct {
	ID                  types.String      `tfsdk:"id" json:"-,computed"`
	Name                types.String      `tfsdk:"name" json:"name"`
	AccountID           types.String      `tfsdk:"account_id" path:"account_id"`
	PrimaryLocationHint types.String      `tfsdk:"primary_location_hint" json:"primary_location_hint"`
	CreatedAt           timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed"`
	FileSize            types.Float64     `tfsdk:"file_size" json:"file_size,computed"`
	NumTables           types.Float64     `tfsdk:"num_tables" json:"num_tables,computed"`
	UUID                types.String      `tfsdk:"uuid" json:"uuid,computed"`
	Version             types.String      `tfsdk:"version" json:"version,computed"`
}
