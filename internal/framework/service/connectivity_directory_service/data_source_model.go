package connectivity_directory_service

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ConnectivityDirectoryServiceDataSourceModel is the single-item data source model.
type ConnectivityDirectoryServiceDataSourceModel struct {
	ID        types.String                           `tfsdk:"id"`
	AccountID types.String                           `tfsdk:"account_id"`
	ServiceID types.String                           `tfsdk:"service_id"`
	Name      types.String                           `tfsdk:"name"`
	Type      types.String                           `tfsdk:"type"`
	Host      *ConnectivityDirectoryServiceHostModel `tfsdk:"host"`
	HTTPPort  types.Int64                            `tfsdk:"http_port"`
	HTTPSPort types.Int64                            `tfsdk:"https_port"`
	CreatedAt types.String                           `tfsdk:"created_at"`
	UpdatedAt types.String                           `tfsdk:"updated_at"`
}

// ConnectivityDirectoryServicesDataSourceModel is the list data source model.
type ConnectivityDirectoryServicesDataSourceModel struct {
	ID        types.String                                        `tfsdk:"id"`
	AccountID types.String                                        `tfsdk:"account_id"`
	Type      types.String                                        `tfsdk:"type"`
	Services  []ConnectivityDirectoryServiceDataSourceResultModel `tfsdk:"services"`
}

// ConnectivityDirectoryServiceDataSourceResultModel is a single item in the list result.
type ConnectivityDirectoryServiceDataSourceResultModel struct {
	ID        types.String                           `tfsdk:"id"`
	ServiceID types.String                           `tfsdk:"service_id"`
	Name      types.String                           `tfsdk:"name"`
	Type      types.String                           `tfsdk:"type"`
	Host      *ConnectivityDirectoryServiceHostModel `tfsdk:"host"`
	HTTPPort  types.Int64                            `tfsdk:"http_port"`
	HTTPSPort types.Int64                            `tfsdk:"https_port"`
	CreatedAt types.String                           `tfsdk:"created_at"`
	UpdatedAt types.String                           `tfsdk:"updated_at"`
}
