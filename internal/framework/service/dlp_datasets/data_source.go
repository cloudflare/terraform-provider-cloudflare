package dlp_datasets

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = &CloudflareDlpDatasetsDataSource{}

func NewDataSource() datasource.DataSource {
	return &CloudflareDlpDatasetsDataSource{}
}

type CloudflareDlpDatasetsDataSource struct {
	client *cloudflare.API
}

func (d *CloudflareDlpDatasetsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dlp_datasets"
}

func (d *CloudflareDlpDatasetsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*cloudflare.API)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected resource configure type",
			fmt.Sprintf("Expected *cloudflare.API, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	d.client = client
}

func (d *CloudflareDlpDatasetsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data DlpDatasetsModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	accountDatasets, err := d.client.ListDLPDatasets(ctx, cloudflare.AccountIdentifier(data.AccountID.ValueString()), cloudflare.ListDLPDatasetsParams{})
	if err != nil {
		resp.Diagnostics.AddError("failed to fetch DLP Datasets: %w", err.Error())
		return
	}

	if len(accountDatasets) < 1 {
		return
	}

	var datasets []*DlpDatasetModel
	for _, dataset := range accountDatasets {
		datasets = append(datasets, &DlpDatasetModel{
			ID:          types.StringValue(dataset.ID),
			Name:        types.StringValue(dataset.Name),
			Description: types.StringValue(dataset.Description),
			Status:      types.StringValue(dataset.Status),
			Secret:      types.BoolValue(*dataset.Secret),
		})
	}

	data.Datasets = datasets
	resp.Diagnostics.Append(resp.State.Set(ctx, datasets)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
