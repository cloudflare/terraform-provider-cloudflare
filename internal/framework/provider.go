package framework

import (
	"context"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure CloudflareProvider satisfies various provider interfaces.
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
	APIKey            types.String `tfsdk:"api_key"`
	APIUserServiceKey types.String `tfsdk:"api_user_service_key"`
	Email             types.String `tfsdk:"email"`
	MinBackOff        types.Int64  `tfsdk:"min_backoff"`
	RPS               types.Int64  `tfsdk:"rps"`
	AccountID         types.String `tfsdk:"account_id"`
	APIBasePath       types.String `tfsdk:"api_base_path"`
	APIToken          types.String `tfsdk:"api_token"`
	Retries           types.Int64  `tfsdk:"retries"`
	MaxBackoff        types.Int64  `tfsdk:"max_backoff"`
	APIClientLogging  types.Bool   `tfsdk:"api_client_logging"`
	APIHostname       types.String `tfsdk:"api_hostname"`
}

func (p *CloudflareProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "cloudflare"
	resp.Version = p.version
}

func (p *CloudflareProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"email": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "A registered Cloudflare email address. Alternatively, can be configured using the `CLOUDFLARE_EMAIL` environment variable.",
			},

			"api_key": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "The API key for operations. Alternatively, can be configured using the `CLOUDFLARE_API_KEY` environment variable. API keys are [now considered legacy by Cloudflare](https://developers.cloudflare.com/api/keys/#limitations), API tokens should be used instead.",
			},

			"api_token": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "The API Token for operations. Alternatively, can be configured using the `CLOUDFLARE_API_TOKEN` environment variable.",
			},

			"api_user_service_key": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "A special Cloudflare API key good for a restricted set of endpoints. Alternatively, can be configured using the `CLOUDFLARE_API_USER_SERVICE_KEY` environment variable.",
			},

			"rps": schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: "RPS limit to apply when making calls to the API. Alternatively, can be configured using the `CLOUDFLARE_RPS` environment variable.",
			},

			"retries": schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: "Maximum number of retries to perform when an API request fails. Alternatively, can be configured using the `CLOUDFLARE_RETRIES` environment variable.",
			},

			"min_backoff": schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: "Minimum backoff period in seconds after failed API calls. Alternatively, can be configured using the `CLOUDFLARE_MIN_BACKOFF` environment variable.",
			},

			"max_backoff": schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: "Maximum backoff period in seconds after failed API calls. Alternatively, can be configured using the `CLOUDFLARE_MAX_BACKOFF` environment variable.",
			},

			"api_client_logging": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Whether to print logs from the API client (using the default log library logger). Alternatively, can be configured using the `CLOUDFLARE_API_CLIENT_LOGGING` environment variable.",
			},

			"account_id": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Configure API client to always use a specific account. Alternatively, can be configured using the `CLOUDFLARE_ACCOUNT_ID` environment variable.",
				DeprecationMessage:  "Use resource specific `account_id` attributes instead.",
			},

			"api_hostname": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Configure the hostname used by the API client. Alternatively, can be configured using the `CLOUDFLARE_API_HOSTNAME` environment variable.",
			},

			"api_base_path": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Configure the base path used by the API client. Alternatively, can be configured using the `CLOUDFLARE_API_BASE_PATH` environment variable.",
			},
		},
	}
}

func (p *CloudflareProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data CloudflareProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Configuration values are now available.
	// if data.Endpoint.IsNull() { /* ... */ }

	// Example client configuration for data sources and resources
	options := []cloudflare.Option{}

	options = append(options, cloudflare.BaseURL("https://api.cloudflare.com/client/v4"))
	config := Config{Options: options}

	client, _ := config.Client()
	// if err != nil {
	// 	resp.Diagnostics.Append(err...)
	// 	return
	// }
	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *CloudflareProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewExampleResource,
	}
}

func (p *CloudflareProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewExampleDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &CloudflareProvider{
			version: version,
		}
	}
}
