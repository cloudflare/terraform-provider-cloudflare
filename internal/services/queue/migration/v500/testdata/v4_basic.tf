resource "cloudflare_queue" "%[1]s" {
  account_id = "%[2]s"
  name       = "tf-acc-test-queue-%[1]s"
}
