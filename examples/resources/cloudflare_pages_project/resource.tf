
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
      environment_variables = {
        ENVIRONMENT = "preview"
      }
    }
    production {
      environment_variables = {
        ENVIRONMENT = "production"
      }
    }
  }
}

# Manage compatiablity date and flags
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
      compatibility_date = "2022-08-15"
      compatibility_flags = ["url_standard"]
    }
    production {
      compatibility_date = "2022-08-15"
      compatibility_flags = ["url_standard"]
    }
  }
}

# Add custom domain to pages project
resource "cloudflare_pages_domain" "my-domain" {
  account_id = var.account_id
  project_name = cloudflare_pages_project.build_config.name
  domain = "example.com"
}