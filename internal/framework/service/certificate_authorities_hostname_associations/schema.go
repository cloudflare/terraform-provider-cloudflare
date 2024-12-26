package certificate_authorities_hostname_associations

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *CertificateAuthoritiesHostnameAssociationsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Provides a Cloudflare Certificate Authorities Hostname Associations resource.",
		Attributes: map[string]schema.Attribute{
			consts.ZoneIDSchemaKey: schema.StringAttribute{
				Description: consts.ZoneIDSchemaDescription,
				Required:    true,
			},
			"mtls_certificate_id": schema.StringAttribute{
				Description: "TODO",
				Optional:    true,
			},
			"hostnames": schema.ListAttribute{
				Description: "TODO",
				Required:    true,
				ElementType: types.StringType,
			},
		},
	}
}
