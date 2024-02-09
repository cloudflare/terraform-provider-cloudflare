package zaraz

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the desired interfaces.
var _ datasource.DataSource = &ZarazConfigDataSource{}

func NewDataSource() datasource.DataSource {
	return &ZarazConfigDataSource{}
}

type ZarazConfigDataSource struct {
	client *cloudflare.API
}

type ZarazConfigDataSourceModel struct {
	ZoneID   types.String `tfsdk:"zone_id"`
	DebugKey types.String `tfsdk:"debugKey"`
}

func (d *ZarazConfigDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "_zaraz_config"
}

func (d *ZarazConfigDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

	d.client = client
}

func (d *ZarazConfigDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			consts.ZoneIDSchemaKey: schema.StringAttribute{
				MarkdownDescription: consts.ZoneIDSchemaDescription,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.ConflictsWith(
						path.Expression(path.MatchRoot((consts.AccountIDSchemaKey))),
					),
				},
			},
		},
	}
}

func (d *ZarazConfigDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data ZarazConfigDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	// Typically data sources will make external calls, however this example
	// hardcodes setting the id attribute to a specific value for brevity.
	rc := cloudflare.ZoneIdentifier(data.ZoneID.ValueString())
	response, _ := d.client.GetZarazConfig(ctx, rc)

	config := response.Result
	data = ZarazConfigDataSourceModel{
		ZoneID:   types.StringValue(data.ZoneID.ValueString()),
		DebugKey: types.StringValue(config.DebugKey),
	}
	data.ZoneID = types.StringValue(data.ZoneID.ValueString())

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
