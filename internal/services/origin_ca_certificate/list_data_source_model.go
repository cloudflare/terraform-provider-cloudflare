// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package origin_ca_certificate

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type OriginCACertificatesResultListDataSourceEnvelope struct {
	Result *[]*OriginCACertificatesItemsDataSourceModel `json:"result,computed"`
}

type OriginCACertificatesDataSourceModel struct {
	ZoneID   types.String                                 `tfsdk:"zone_id" query:"zone_id"`
	MaxItems types.Int64                                  `tfsdk:"max_items"`
	Items    *[]*OriginCACertificatesItemsDataSourceModel `tfsdk:"items"`
}

type OriginCACertificatesItemsDataSourceModel struct {
	Csr               types.String            `tfsdk:"csr" json:"csr,computed"`
	Hostnames         *[]jsontypes.Normalized `tfsdk:"hostnames" json:"hostnames,computed"`
	RequestType       types.String            `tfsdk:"request_type" json:"request_type,computed"`
	RequestedValidity types.Float64           `tfsdk:"requested_validity" json:"requested_validity,computed"`
	ID                types.String            `tfsdk:"id" json:"id,computed"`
	Certificate       types.String            `tfsdk:"certificate" json:"certificate,computed"`
	ExpiresOn         timetypes.RFC3339       `tfsdk:"expires_on" json:"expires_on,computed"`
}
