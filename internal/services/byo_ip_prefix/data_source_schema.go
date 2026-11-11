// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package byo_ip_prefix

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*ByoIPPrefixDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"IP Prefixes: BGP On Demand Read",
				"IP Prefixes: BGP On Demand Write",
				"IP Prefixes: Read",
				"IP Prefixes: Write",
				"Magic Transit Read",
				"Magic Transit Write",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Identifier of an IP Prefix.",
				Computed:    true,
			},
			"prefix_id": schema.StringAttribute{
				Description: "Identifier of an IP Prefix.",
				Required:    true,
			},
			"account_id": schema.StringAttribute{
				Description: "Identifier of a Cloudflare account.",
				Optional:    true,
			},
			"advertised": schema.BoolAttribute{
				Description:        "Prefix advertisement status to the Internet. This field is only not 'null' if on demand is enabled.",
				Computed:           true,
				DeprecationMessage: "Prefer the [BGP Prefixes API](https://developers.cloudflare.com/api/resources/addressing/subresources/prefixes/subresources/bgp_prefixes/) instead, which allows for advertising multiple BGP routes within a single IP Prefix.",
			},
			"advertised_modified_at": schema.StringAttribute{
				Description:        "Last time the advertisement status was changed. This field is only not 'null' if on demand is enabled.",
				Computed:           true,
				DeprecationMessage: "Prefer the [BGP Prefixes API](https://developers.cloudflare.com/api/resources/addressing/subresources/prefixes/subresources/bgp_prefixes/) instead, which allows for advertising multiple BGP routes within a single IP Prefix.",
				CustomType:         timetypes.RFC3339Type{},
			},
			"approved": schema.StringAttribute{
				Description: "Approval state of the prefix (P = pending, V = approved). Approval does not advertise the prefix or activate traffic.",
				Computed:    true,
			},
			"asn": schema.Int64Attribute{
				Description: "Autonomous System Number (ASN) the prefix will be advertised under. Use 13335 for Cloudflare's ASN, or your own ASN as an integer without the AS prefix. The IRR route and RPKI ROA must authorize this ASN.",
				Computed:    true,
			},
			"asn_ownership_validation_state": schema.StringAttribute{
				Description: "State of ASN ownership validation for a customer-owned Autonomous System Number (ASN). Returned when an ASN ownership validation result exists. Omitted when validation does not apply, such as for Cloudflare-owned ASNs. Verification is reused for prefixes with the same account and ASN. Refer to [ASN ownership validation states](https://developers.cloudflare.com/byoip/troubleshooting/prefix-validation/#asn-ownership-validation) for state meanings and resolutions.\nAvailable values: \"pending\", \"valid\", \"missing\", \"invalid\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"pending",
						"valid",
						"missing",
						"invalid",
					),
				},
			},
			"aspa_validation_state": schema.StringAttribute{
				Description: "State of Autonomous System Provider Authorization (ASPA) validation for a customer-owned ASN. Returned when an ASPA validation result exists. Publishing an ASPA is optional. A published ASPA must authorize AS13335 as a provider. An invalid result blocks approval. If omitted, no ASPA validation result is available. Omission does not indicate successful validation or confirm that no ASPA exists. Refer to [ASPA validation states](https://developers.cloudflare.com/byoip/troubleshooting/prefix-validation/#aspa-validation) for state meanings and resolutions, and [ASPA authorization](https://developers.cloudflare.com/byoip/concepts/route-filtering-rpki/#aspa-authorization) for publication guidance.\nAvailable values: \"pending\", \"valid\", \"invalid\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"pending",
						"valid",
						"invalid",
					),
				},
			},
			"cidr": schema.StringAttribute{
				Description: "IP Prefix in Classless Inter-Domain Routing format.",
				Computed:    true,
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"delegate_loa_creation": schema.BoolAttribute{
				Description: "Whether Cloudflare is allowed to generate the LOA document on behalf of the prefix owner.",
				Computed:    true,
			},
			"description": schema.StringAttribute{
				Description: "Description of the prefix.",
				Computed:    true,
			},
			"irr_validation_state": schema.StringAttribute{
				Description: "State of one kind of validation for an IP prefix.",
				Computed:    true,
			},
			"loa_document_id": schema.StringAttribute{
				Description: "Identifier for the uploaded LOA document.",
				Computed:    true,
			},
			"modified_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"on_demand_enabled": schema.BoolAttribute{
				Description:        "Whether advertisement of the prefix to the Internet may be dynamically enabled or disabled.",
				Computed:           true,
				DeprecationMessage: "Prefer the [BGP Prefixes API](https://developers.cloudflare.com/api/resources/addressing/subresources/prefixes/subresources/bgp_prefixes/) instead, which allows for advertising multiple BGP routes within a single IP Prefix.",
			},
			"on_demand_locked": schema.BoolAttribute{
				Description:        "Whether advertisement status of the prefix is locked, meaning it cannot be changed.",
				Computed:           true,
				DeprecationMessage: "Prefer the [BGP Prefixes API](https://developers.cloudflare.com/api/resources/addressing/subresources/prefixes/subresources/bgp_prefixes/) instead, which allows for advertising multiple BGP routes within a single IP Prefix.",
			},
			"ownership_validation_state": schema.StringAttribute{
				Description: "State of one kind of validation for an IP prefix.",
				Computed:    true,
			},
			"ownership_validation_token": schema.StringAttribute{
				Description: "Token used to verify ownership of the prefix and, when applicable, its ASN. Use the same token for both checks. Refer to the [BYOIP onboarding guide](https://developers.cloudflare.com/byoip/get-started/) for verification instructions.",
				Computed:    true,
			},
			"rpki_validation_state": schema.StringAttribute{
				Description: "State of one kind of validation for an IP prefix.",
				Computed:    true,
			},
		},
	}
}

func (d *ByoIPPrefixDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ByoIPPrefixDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
