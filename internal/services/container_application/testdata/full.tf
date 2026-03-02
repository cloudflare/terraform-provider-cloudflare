resource "cloudflare_container_application" "example" {
  account_id  = "%[1]s"
  script_name = "my-worker"
  class_name  = "MyDurableObject"
  name        = "custom-app-name"
  image       = "registry.example.com/my-image:v2"

  max_instances               = 50
  instance_type               = "standard-1"
  scheduling_policy           = "regional"
  rollout_step_percentage     = 10
  rollout_kind                = "full_manual"
  rollout_active_grace_period = 60

  constraints {
    tiers   = [1, 2]
    regions = ["us-east", "eu-west"]
    cities  = []
  }
}
