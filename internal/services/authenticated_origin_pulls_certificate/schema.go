// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package authenticated_origin_pulls_certificate

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

var _ resource.ResourceWithConfigValidators = (*AuthenticatedOriginPullsCertificateResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description:   "Identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"certificate_id": schema.StringAttribute{
				Description:   "Identifier",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"certificate": schema.StringAttribute{
				Description:   "The zone's leaf certificate.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"private_key": schema.StringAttribute{
				Description:   "The zone's private key.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"enabled": schema.BoolAttribute{
				Description: "Indicates whether zone-level authenticated origin pulls is enabled.",
				Computed:    true,
			},
			"expires_on": schema.StringAttribute{
				Description: "When the certificate from the authority expires.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"id": schema.StringAttribute{
				Description: "Identifier",
				Computed:    true,
			},
			"issuer": schema.StringAttribute{
				Description: "The certificate authority that issued the certificate.",
				Computed:    true,
			},
			"signature": schema.StringAttribute{
				Description: "The type of hash used for the certificate.",
				Computed:    true,
			},
			"status": schema.StringAttribute{
				Description: "Status of the certificate activation.\navailable values: \"initializing\", \"pending_deployment\", \"pending_deletion\", \"active\", \"deleted\", \"deployment_timed_out\", \"deletion_timed_out\"",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"initializing",
						"pending_deployment",
						"pending_deletion",
						"active",
						"deleted",
						"deployment_timed_out",
						"deletion_timed_out",
					),
				},
			},
			"uploaded_on": schema.StringAttribute{
				Description: "This is the time the certificate was uploaded.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
		},
	}
}

func (r *AuthenticatedOriginPullsCertificateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *AuthenticatedOriginPullsCertificateResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
