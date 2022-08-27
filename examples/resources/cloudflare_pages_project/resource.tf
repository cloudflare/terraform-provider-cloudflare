# Cloudflare Account ID
account_id = ""

# Basic project
resource "cloudflare_pages_project" "basic_project" {
  account_id = var.account_id
  name = "this-is-my-project-01"
  production_branch = "main"
}

# Manage build config
resource "cloudflare_pages_project" "build_config" {
  account_id = var.account_id
  name = "this-is-my-project-01"
  production_branch = "main"
  build_config {
    build_command = "npm run build"
    destination_dir = "build"
    root_dir = "/"
    web_analytics_tag = "cee1c73f6e4743d0b5e6bb1a0bcaabcc"
    web_analytics_token = "021e1057c18547eca7b79f2516f06o7x"
  }
}


# Manage project source
resource "cloudflare_pages_project" "%[1]s" {
  account_id = var.account_id
  name = "this-is-my-project-01"
  production_branch = "main"
  source {
    type = "github"
    config {
      owner = "cloudflare"
      repo_name = "ninjakittens"
      production_branch = "main"
      pr_comments_enabled = true
      deployments_enabled = true
      production_deployment_enabled = true
      preview_deployment_setting = "custom"
      preview_branch_includes = ["dev","preview"]
      preview_branch_excludes = ["main", "prod"]

    }
  }
}
# Manage deployment configs
resource "cloudflare_pages_project" "deployment_configs" {
  account_id = var.account_id
  name = "this-is-my-project-01"
  production_branch = "main"
  deployment_configs {
    preview {
      environment_variables = {
        ENVIRONMENT = "preview"
      }
      kv_namespaces = {
        KV_BINDING = "5eb63bbbe01eeed093cb22bb8f5acdc3"
      }
      durable_object_namespaces = {
        DO_BINDING = "5eb63bbbe01eeed093cb22bb8f5acdc3"
      }
      r2_buckets = {
        R2_BINDING = "some-bucket"
      }
      d1_databases = {
        D1_BINDING = "445e2955-951a-4358-a35b-a4d0c813f63"
      }
      compatibility_date = "2022-08-15"
      compatibility_flags = ["preview_flag"]
    }
    production {
      environment_variables = {
        ENVIRONMENT = "production"
        OTHER_VALUE = "other value"
      }
      kv_namespaces = {
        KV_BINDING_1 = "5eb63bbbe01eeed093cb22bb8f5acdc3"
        KV_BINDING_2 = "3cdca5f8bb22bc390deee10ebbb36be5"
      }
      durable_object_namespaces = {
        DO_BINDING_1 = "5eb63bbbe01eeed093cb22bb8f5acdc3"
        DO_BINDING_2 = "3cdca5f8bb22bc390deee10ebbb36be5"
      }
      r2_buckets = {
        R2_BINDING_1 = "some-bucket"
        R2_BINDING_2 = "other-bucket"
      }
      d1_databases = {
        D1_BINDING_1 = "445e2955-951a-4358-a35b-a4d0c813f63"
        D1_BINDING_2 = "a399414b-c697-409a-a688-377db6433cd9"
      }
      compatibility_date = "2022-08-16"
      compatibility_flags = ["production_flag", "second flag"]
    }
}
}



# Add custom domain to pages project
resource "cloudflare_pages_domain" "my-domain" {
  account_id = var.account_id
  project_name = cloudflare_pages_project.build_config.name
  domain = "example.com"
}