// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ip_ranges

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/ips"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type IPRangesResultDataSourceEnvelope struct {
	Result IPRangesDataSourceModel `json:"result,computed"`
}

type IPRangesDataSourceModel struct {
	Networks     types.String                   `tfsdk:"networks" query:"networks,optional"`
	Etag         types.String                   `tfsdk:"etag" json:"etag,computed"`
	IPV4CIDRs    customfield.List[types.String] `tfsdk:"ipv4_cidrs" json:"ipv4_cidrs,computed"`
	IPV6CIDRs    customfield.List[types.String] `tfsdk:"ipv6_cidrs" json:"ipv6_cidrs,computed"`
	JDCloudCIDRs customfield.List[types.String] `tfsdk:"jdcloud_cidrs" json:"jdcloud_cidrs,computed"`
}

func (m *IPRangesDataSourceModel) toReadParams(_ context.Context) (params ips.IPListParams, diags diag.Diagnostics) {
	params = ips.IPListParams{}

	if !m.Networks.IsNull() {
		params.Networks = cloudflare.F(m.Networks.ValueString())
	}

	return
}
