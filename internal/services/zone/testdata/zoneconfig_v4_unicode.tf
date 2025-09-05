resource "cloudflare_zone" "%[1]s" {
  zone       = "例え.テスト"
  account_id = "%[3]s"
  type       = "full"
}