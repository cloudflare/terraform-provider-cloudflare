// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_dnssec

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func (r ZoneDNSSECResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Identifier",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"zone_id": schema.StringAttribute{
				Description:   "Identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"dnssec_multi_signer": schema.BoolAttribute{
				Description: "If true, multi-signer DNSSEC is enabled on the zone, allowing multiple\nproviders to serve a DNSSEC-signed zone at the same time.\nThis is required for DNSKEY records (except those automatically\ngenerated by Cloudflare) to be added to the zone.\n\nSee [Multi-signer DNSSEC](https://developers.cloudflare.com/dns/dnssec/multi-signer-dnssec/) for details.",
				Optional:    true,
			},
			"dnssec_presigned": schema.BoolAttribute{
				Description: "If true, allows Cloudflare to transfer in a DNSSEC-signed zone\nincluding signatures from an external provider, without requiring\nCloudflare to sign any records on the fly.\n\nNote that this feature has some limitations.\nSee [Cloudflare as Secondary](https://developers.cloudflare.com/dns/zone-setups/zone-transfers/cloudflare-as-secondary/setup/#dnssec) for details.",
				Optional:    true,
			},
			"status": schema.StringAttribute{
				Description: "Status of DNSSEC, based on user-desired state and presence of necessary records.",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("active", "disabled"),
				},
			},
			"algorithm": schema.StringAttribute{
				Description: "Algorithm key code.",
				Computed:    true,
			},
			"digest": schema.StringAttribute{
				Description: "Digest hash.",
				Computed:    true,
			},
			"digest_algorithm": schema.StringAttribute{
				Description: "Type of digest algorithm.",
				Computed:    true,
			},
			"digest_type": schema.StringAttribute{
				Description: "Coded type for digest algorithm.",
				Computed:    true,
			},
			"ds": schema.StringAttribute{
				Description: "Full DS record.",
				Computed:    true,
			},
			"flags": schema.Float64Attribute{
				Description: "Flag for DNSSEC record.",
				Computed:    true,
			},
			"key_tag": schema.Float64Attribute{
				Description: "Code for key tag.",
				Computed:    true,
			},
			"key_type": schema.StringAttribute{
				Description: "Algorithm key type.",
				Computed:    true,
			},
			"modified_on": schema.StringAttribute{
				Description: "When DNSSEC was last modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"public_key": schema.StringAttribute{
				Description: "Public key for DS record.",
				Computed:    true,
			},
		},
	}
}
