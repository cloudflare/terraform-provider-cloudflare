package origin_ca_certificate

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = &CloudflareOriginCACertificateDataSource{}

func NewDataSource() datasource.DataSource {
	return &CloudflareOriginCACertificateDataSource{}
}

type CloudflareOriginCACertificateDataSource struct {
	client *cloudflare.API
}

func (r *CloudflareOriginCACertificateDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_origin_ca_certificate"
}

func (r *CloudflareOriginCACertificateDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*cloudflare.API)

	if !ok {
		resp.Diagnostics.AddError(
			"unexpected resource configure type",
			fmt.Sprintf("expected *cloudflare.API, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *CloudflareOriginCACertificateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data CloudflareOriginCACertificateModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	cert, err := r.client.GetOriginCACertificate(ctx, data.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("failed to fetch Origin CA: %w", err.Error())
		return
	}

	data = CloudflareOriginCACertificateModel{
		ID:          types.StringValue(cert.ID),
		Certificate: types.StringValue(cert.Certificate),
		ExpiresOn:   types.StringValue(cert.ExpiresOn.String()),
		RequestType: types.StringValue(cert.RequestType),
		RevokedAt:   types.StringValue(cert.RevokedAt.String()),
	}

	elements := make([]attr.Value, 0, len(cert.Hostnames))
	for _, hostname := range cert.Hostnames {
		elements = append(elements, types.StringValue(hostname))
	}
	listValue, diags := types.ListValue(types.StringType, elements)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	data.Hostnames = listValue

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
