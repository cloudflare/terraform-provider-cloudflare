resource "cloudflare_notification_policy_webhooks" "%[1]s" {
    account_id  = "%[2]s"
    name        = "my webhooks destination for receiving Cloudflare notifications"
    url         = "https://httpbin.org/post"
    secret      =  "my-secret"
}
