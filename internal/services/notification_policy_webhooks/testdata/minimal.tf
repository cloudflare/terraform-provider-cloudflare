resource "cloudflare_notification_policy_webhooks" "%[1]s" {
    account_id  = "%[2]s"
    name        = "%[3]s"
    url         = "%[4]s"
}