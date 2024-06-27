// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package authenticated_origin_pulls

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func (r AuthenticatedOriginPullsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"hostname": schema.StringAttribute{
				Description: "The hostname on the origin for which the client certificate uploaded will be used.",
				Optional:    true,
			},
			"config": schema.ListNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"cert_id": schema.StringAttribute{
							Description: "Certificate identifier tag.",
							Optional:    true,
						},
						"enabled": schema.BoolAttribute{
							Description: "Indicates whether hostname-level authenticated origin pulls is enabled. A null value voids the association.",
							Optional:    true,
						},
						"hostname": schema.StringAttribute{
							Description: "The hostname on the origin for which the client certificate uploaded will be used.",
							Optional:    true,
						},
					},
				},
			},
			"cert_id": schema.StringAttribute{
				Description: "Identifier",
				Computed:    true,
			},
			"cert_status": schema.StringAttribute{
				Description: "Status of the certificate or the association.",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("initializing", "pending_deployment", "pending_deletion", "active", "deleted", "deployment_timed_out", "deletion_timed_out"),
				},
			},
			"cert_updated_at": schema.StringAttribute{
				Description: "The time when the certificate was updated.",
				Computed:    true,
			},
			"cert_uploaded_on": schema.StringAttribute{
				Description: "The time when the certificate was uploaded.",
				Computed:    true,
			},
			"certificate": schema.StringAttribute{
				Description: "The hostname certificate.",
				Computed:    true,
			},
			"created_at": schema.StringAttribute{
				Description: "The time when the certificate was created.",
				Computed:    true,
			},
			"enabled": schema.BoolAttribute{
				Description: "Indicates whether hostname-level authenticated origin pulls is enabled. A null value voids the association.",
				Computed:    true,
			},
			"expires_on": schema.StringAttribute{
				Description: "The date when the certificate expires.",
				Computed:    true,
			},
			"issuer": schema.StringAttribute{
				Description: "The certificate authority that issued the certificate.",
				Computed:    true,
			},
			"serial_number": schema.StringAttribute{
				Description: "The serial number on the uploaded certificate.",
				Computed:    true,
			},
			"signature": schema.StringAttribute{
				Description: "The type of hash used for the certificate.",
				Computed:    true,
			},
			"status": schema.StringAttribute{
				Description: "Status of the certificate or the association.",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("initializing", "pending_deployment", "pending_deletion", "active", "deleted", "deployment_timed_out", "deletion_timed_out"),
				},
			},
			"updated_at": schema.StringAttribute{
				Description: "The time when the certificate was updated.",
				Computed:    true,
			},
		},
	}
}
