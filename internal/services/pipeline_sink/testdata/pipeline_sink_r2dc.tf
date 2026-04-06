resource "cloudflare_r2_bucket" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
}

resource "cloudflare_r2_data_catalog" "%[1]s" {
  account_id  = "%[2]s"
  bucket_name = cloudflare_r2_bucket.%[1]s.name
}

data "cloudflare_account_api_token_permission_groups_list" "r2_bucket_item_write" {
  account_id = "%[2]s"
  name       = "Workers R2 Storage Bucket Item Write"
}

data "cloudflare_account_api_token_permission_groups_list" "r2_data_catalog_write" {
  account_id = "%[2]s"
  name       = "Workers R2 Data Catalog Write"
}

resource "cloudflare_account_token" "%[1]s" {
  name       = "%[1]s"
  account_id = "%[2]s"

  policies = [{
    effect = "allow"
    permission_groups = [
      { id = data.cloudflare_account_api_token_permission_groups_list.r2_bucket_item_write.result[0].id },
      { id = data.cloudflare_account_api_token_permission_groups_list.r2_data_catalog_write.result[0].id },
    ]
    resources = jsonencode({
      "com.cloudflare.api.account.%[2]s" = "*"
    })
  }]
}

resource "cloudflare_pipeline_sink" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  type       = "r2_data_catalog"
  format = {
    type = "parquet"
  }
  schema = {
    fields = []
  }
  config = {
    account_id = "%[2]s"
    bucket     = cloudflare_r2_bucket.%[1]s.name
    table_name = cloudflare_r2_data_catalog.%[1]s.name
    token      = cloudflare_account_token.%[1]s.value
  }
}
