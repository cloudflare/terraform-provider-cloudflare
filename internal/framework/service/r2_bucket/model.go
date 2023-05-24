package r2_bucket

import "github.com/hashicorp/terraform-plugin-framework/types"

type R2BucketModel struct {
	AccountID types.String `tfsdk:"account_id"`
	Name      types.String `tfsdk:"name"`
	ID        types.String `tfsdk:"id"`
	Location  types.String `tfsdk:"location"`
}
