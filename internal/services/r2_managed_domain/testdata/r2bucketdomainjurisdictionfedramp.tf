resource "cloudflare_r2_bucket" "%[1]s" {
    account_id    = "%[2]s"
    name          = "%[1]s"
    jurisdiction  = "fedramp"
}

resource "cloudflare_r2_managed_domain" "%[1]s" {
    account_id = "%[2]s"
    bucket_name = "%[1]s"
    enabled = "true"
    jurisdiction = "fedramp"
    depends_on = [
        cloudflare_r2_bucket.%[1]s
    ]
}
