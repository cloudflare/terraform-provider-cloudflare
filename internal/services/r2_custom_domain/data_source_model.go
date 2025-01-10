// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_custom_domain

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/r2"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type R2CustomDomainResultDataSourceEnvelope struct {
	Result R2CustomDomainDataSourceModel `json:"result,computed"`
}

type R2CustomDomainDataSourceModel struct {
	AccountID  types.String                                                  `tfsdk:"account_id" path:"account_id,required"`
	BucketName types.String                                                  `tfsdk:"bucket_name" path:"bucket_name,required"`
	DomainName types.String                                                  `tfsdk:"domain_name" path:"domain_name,required"`
	Domain     types.String                                                  `tfsdk:"domain" json:"domain,computed"`
	Enabled    types.Bool                                                    `tfsdk:"enabled" json:"enabled,computed"`
	MinTLS     types.String                                                  `tfsdk:"min_tls" json:"minTLS,computed"`
	ZoneID     types.String                                                  `tfsdk:"zone_id" json:"zoneId,computed"`
	ZoneName   types.String                                                  `tfsdk:"zone_name" json:"zoneName,computed"`
	Status     customfield.NestedObject[R2CustomDomainStatusDataSourceModel] `tfsdk:"status" json:"status,computed"`
}

func (m *R2CustomDomainDataSourceModel) toReadParams(_ context.Context) (params r2.BucketDomainCustomGetParams, diags diag.Diagnostics) {
	params = r2.BucketDomainCustomGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type R2CustomDomainStatusDataSourceModel struct {
	Ownership types.String `tfsdk:"ownership" json:"ownership,computed"`
	SSL       types.String `tfsdk:"ssl" json:"ssl,computed"`
}
