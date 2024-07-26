
	resource "cloudflare_api_shield_operation" "terraform_test_acc_operation" {
		zone_id = "%[2]s"
		host = "foo.com"
		method = "GET"
        endpoint = "/api"
	}
	resource "cloudflare_api_shield_operation_schema_validation_settings" "%[1]s" {
		zone_id = "%[2]s"
		operation_id = cloudflare_api_shield_operation.terraform_test_acc_operation.id
		mitigation_action = %[3]s
	}
