resource "cloudflare_container_application" "example" {
  account_id  = "%[1]s"
  script_name = "my-worker"
  class_name  = "MyDurableObject"
  name        = "custom-app-name"
  image       = "registry.example.com/my-image:v3"

  max_instances = 100
  instance_type = "standard-1"
}
