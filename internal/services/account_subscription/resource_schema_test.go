// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_subscription_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/account_subscription"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestAccountSubscriptionModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*account_subscription.AccountSubscriptionModel)(nil)
	schema := account_subscription.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
