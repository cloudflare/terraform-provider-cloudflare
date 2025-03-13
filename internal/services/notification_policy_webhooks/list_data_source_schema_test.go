// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package notification_policy_webhooks_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/notification_policy_webhooks"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestNotificationPolicyWebhooksListDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*notification_policy_webhooks.NotificationPolicyWebhooksListDataSourceModel)(nil)
  schema := notification_policy_webhooks.ListDataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
