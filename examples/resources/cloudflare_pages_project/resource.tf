
account_id = ""
# Basic project
resource "cloudflare_pages_project" "basic_project" {
  account_id = var.account_id
  name = "this-is-my-project-01"
}

# Manage build config
resource "cloudflare_pages_project" "build_config" {
  account_id = var.account_id
  name = "this-is-my-project-01"
  build_config {
    build_command = "npm run build"
    destination_dir = "build"
    root_dir = "/"
    web_analytics_tag = "cee1c73f6e4743d0b5e6bb1a0bcaabcc"
    web_analytics_token = "021e1057c18547eca7b79f2516f06o7x"
  }
}

# Manage deployment_configs
resource "cloudflare_pages_project" "build_config" {
  account_id = var.account_id
  name = "this-is-my-project-01"
  build_config {
    build_command = "npm run build"
    destination_dir = "build"
    root_dir = "/"
    web_analytics_tag = "cee1c73f6e4743d0b5e6bb1a0bcaabcc"
    web_analytics_token = "021e1057c18547eca7b79f2516f06o7x"
  }
  deployment_configs {
    preview {
      env_vars {
        BUILD_VERSION {
          value = "3.3"
        }
      }
    }
    production {
      env_vars {
        BUILD_VERSION {
          value = "3.3"
        }
      }
    }
  }
}