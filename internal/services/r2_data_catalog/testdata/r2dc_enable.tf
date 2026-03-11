resource "cloudflare_r2_bucket" "%[1]s" {
    account_id = "%[2]s"
    name       = "%[1]s"
}

resource "cloudflare_r2_data_catalog" "%[1]s" {
    account_id  = "%[2]s"
    bucket_name = cloudflare_r2_bucket.%[1]s.name
}
