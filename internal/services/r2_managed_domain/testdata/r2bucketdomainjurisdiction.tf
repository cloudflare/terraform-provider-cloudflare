resource "cloudflare_r2_bucket" "%[1]s" {
    account_id    = "%[2]s"
    name          = "%[1]s"
    jurisdiction  = "eu"
}

resource "cloudflare_r2_managed_domain" "%[1]s" {
    account_id = "%[2]s"
    bucket_name = "%[1]s"
    enabled = "false"
    jurisdiction = "eu"
    depends_on = [
        cloudflare_r2_bucket.%[1]s
    ]
}
