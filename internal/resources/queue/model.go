// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package queue

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type QueueResultEnvelope struct {
	Result QueueModel `json:"result,computed"`
}

type QueueModel struct {
	QueueID   types.String `tfsdk:"queue_id" json:"queue_id,computed"`
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
}
