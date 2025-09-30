// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_tags

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZoneTagsModel struct {
	ZoneID types.String `tfsdk:"zone_id" json:"-"`
	Tags   types.Map    `tfsdk:"tags" json:"tags"`
}