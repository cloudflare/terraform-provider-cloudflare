data "cloudflare_cloudforce_one_requests" "example_cloudforce_one_requests" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  page = 0
  per_page = 10
  completed_after = "2022-01-01T00:00:00Z"
  completed_before = "2024-01-01T00:00:00Z"
  created_after = "2022-01-01T00:00:00Z"
  created_before = "2024-01-01T00:00:00Z"
  request_type = "Victomology"
  sort_by = "created"
  sort_order = "asc"
  status = "open"
}
