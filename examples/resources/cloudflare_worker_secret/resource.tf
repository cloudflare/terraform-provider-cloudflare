resource "cloudflare_worker_secret" "my_secret" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  name        = "MY_EXAMPLE_SECRET_TEXT"
  script_name = "script_1"
  secret_text = "my_secret_value"
}
