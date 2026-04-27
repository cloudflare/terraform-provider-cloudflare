resource "cloudflare_email_security_impersonation_registry" "example_email_security_impersonation_registry" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  email = "john.doe@example.com"
  is_email_regex = false
  name = "John Doe"
  comments = "comments"
  directory_id = 0
  directory_node_id = 0
  external_directory_node_id = "external_directory_node_id"
  provenance = "A1S_INTERNAL"
}
