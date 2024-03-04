package dlp_datasets

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DlpDatasetsModel struct {
	AccountID types.String       `tfsdk:"account_id"`
	Datasets  []*DlpDatasetModel `tfsdk:"datasets"`
}

type DlpDatasetModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	Status      types.String `tfsdk:"status"`
	Secret      types.Bool   `tfsdk:"secret"`
}
