// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_certificates

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r ZeroTrustAccessCertificatesResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.",
				Optional:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.",
				Optional:    true,
			},
			"uuid": schema.StringAttribute{
				Description: "UUID",
				Optional:    true,
			},
			"certificate": schema.StringAttribute{
				Description: "The certificate content.",
				Required:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the certificate.",
				Required:    true,
			},
			"associated_hostnames": schema.ListAttribute{
				Description: "The hostnames of the applications that will use this certificate.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"id": schema.StringAttribute{
				Description: "The ID of the application that will use this certificate.",
				Optional:    true,
			},
			"created_at": schema.StringAttribute{
				Computed: true,
			},
			"expires_on": schema.StringAttribute{
				Computed: true,
			},
			"fingerprint": schema.StringAttribute{
				Description: "The MD5 fingerprint of the certificate.",
				Optional:    true,
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}
