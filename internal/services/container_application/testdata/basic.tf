resource "cloudflare_container_application" "example" {
  account_id  = "%[1]s"
  script_name = "my-worker"
  class_name  = "MyDurableObject"
  image       = "registry.example.com/my-image:latest"
}
