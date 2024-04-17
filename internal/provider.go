// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package internal

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-terraform/internal/resources/addressing_address_maps"
	"github.com/cloudflare/cloudflare-terraform/internal/resources/custom_hostnames"
	"github.com/cloudflare/cloudflare-terraform/internal/resources/dns_records"
	"github.com/cloudflare/cloudflare-terraform/internal/resources/email_routing_rules"
	"github.com/cloudflare/cloudflare-terraform/internal/resources/keyless_certificates"
	"github.com/cloudflare/cloudflare-terraform/internal/resources/logpush_jobs"
	"github.com/cloudflare/cloudflare-terraform/internal/resources/waiting_rooms"
	"github.com/cloudflare/cloudflare-terraform/internal/resources/waiting_rooms_events"
	"github.com/cloudflare/cloudflare-terraform/internal/resources/web3_hostnames"
	"github.com/cloudflare/cloudflare-terraform/internal/resources/zero_trust_access_applications"
	"github.com/cloudflare/cloudflare-terraform/internal/resources/zero_trust_access_certificates"
	"github.com/cloudflare/cloudflare-terraform/internal/resources/zero_trust_devices_posture_integrations"
	"github.com/cloudflare/cloudflare-terraform/internal/resources/zero_trust_identity_providers"
	"github.com/cloudflare/cloudflare-terraform/internal/resources/zones"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ provider.Provider = &CloudflareProvider{}

// CloudflareProvider defines the provider implementation.
type CloudflareProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// CloudflareProviderModel describes the provider data model.
type CloudflareProviderModel struct {
	BaseURL        types.String `tfsdk:"base_url" json:"base_url"`
	APIToken       types.String `tfsdk:"api_token" json:"api_token"`
	APIKey         types.String `tfsdk:"api_key" json:"api_key"`
	APIEmail       types.String `tfsdk:"api_email" json:"api_email"`
	UserServiceKey types.String `tfsdk:"user_service_key" json:"user_service_key"`
}

func (p *CloudflareProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "cloudflare"
	resp.Version = p.version
}

func (p CloudflareProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"base_url": schema.StringAttribute{
				Description: "Set the base url that the provider connects to. This can be used for testing in other environments.",
				Optional:    true,
			},
			"api_token": schema.StringAttribute{
				Optional: true,
			},
			"api_key": schema.StringAttribute{
				Optional: true,
			},
			"api_email": schema.StringAttribute{
				Optional: true,
			},
			"user_service_key": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}

func (p *CloudflareProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

	// TODO(terraform): apiKey := os.Getenv("API_KEY")

	var data CloudflareProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	opts := []option.RequestOption{}

	if !data.BaseURL.IsNull() {
		opts = append(opts, option.WithBaseURL(data.BaseURL.ValueString()))
	}
	if !data.APIToken.IsNull() {
		opts = append(opts, option.WithAPIToken(data.APIToken.ValueString()))
	}
	if !data.APIKey.IsNull() {
		opts = append(opts, option.WithAPIKey(data.APIKey.ValueString()))
	}
	if !data.APIEmail.IsNull() {
		opts = append(opts, option.WithAPIEmail(data.APIEmail.ValueString()))
	}
	if !data.UserServiceKey.IsNull() {
		opts = append(opts, option.WithUserServiceKey(data.UserServiceKey.ValueString()))
	}

	client := cloudflare.NewClient(
		opts...,
	)

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *CloudflareProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		zones.NewResource,
		custom_hostnames.NewResource,
		dns_records.NewResource,
		email_routing_rules.NewResource,
		keyless_certificates.NewResource,
		logpush_jobs.NewResource,
		waiting_rooms.NewResource,
		waiting_rooms_events.NewResource,
		web3_hostnames.NewResource,
		addressing_address_maps.NewResource,
		zero_trust_devices_posture_integrations.NewResource,
		zero_trust_identity_providers.NewResource,
		zero_trust_access_applications.NewResource,
		zero_trust_access_certificates.NewResource,
	}
}

func (p *CloudflareProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func NewProvider(version string) func() provider.Provider {
	return func() provider.Provider {
		return &CloudflareProvider{
			version: version,
		}
	}
}
