// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_subscription

import (
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type ZoneSubscriptionResultDataSourceEnvelope struct {
Result ZoneSubscriptionDataSourceModel `json:"result,computed"`
}

type ZoneSubscriptionDataSourceModel struct {
Identifier types.String `tfsdk:"identifier" path:"identifier,required"`
}
