package origin_ca_certificate

import (
	"context"
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *CloudflareOriginCACertificateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Use this data source to retrieve an existing origin ca certificate.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The Origin CA Certificate unique identifier.",
				Required:    true,
			},
			"certificate": schema.StringAttribute{
				Computed:    true,
				Description: "The Origin CA certificate.",
			},
			"expires_on": schema.StringAttribute{
				Computed:    true,
				Description: "The timestamp when the certificate will expire.",
			},
			"hostnames": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "A list of hostnames or wildcard names bound to the certificate.",
			},
			"request_type": schema.StringAttribute{
				Computed:    true,
				Description: fmt.Sprintf("The signature type desired on the certificate. %s", utils.RenderAvailableDocumentationValuesStringSlice([]string{"origin-rsa", "origin-ecc", "keyless-certificate"})),
			},
			"revoked_at": schema.StringAttribute{
				Computed:    true,
				Description: "The timestamp when the certificate was revoked.",
			},
		},
	}
}
