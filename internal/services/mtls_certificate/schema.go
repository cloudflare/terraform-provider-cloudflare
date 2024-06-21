// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mtls_certificate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func (r MTLSCertificateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"ca": schema.BoolAttribute{
				Description: "Indicates whether the certificate is a CA or leaf certificate.",
				Required:    true,
			},
			"certificates": schema.StringAttribute{
				Description: "The uploaded root CA certificate.",
				Required:    true,
			},
			"name": schema.StringAttribute{
				Description: "Optional unique name for the certificate. Only used for human readability.",
				Optional:    true,
			},
			"private_key": schema.StringAttribute{
				Description: "The private key for the certificate",
				Optional:    true,
			},
		},
	}
}
