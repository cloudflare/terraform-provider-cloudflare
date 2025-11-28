resource "cloudflare_pages_project" "example_pages_project" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  name = "my-pages-app"
  production_branch = "main"
  build_config = {
    build_caching = true
    build_command = "npm run build"
    destination_dir = "build"
    root_dir = "/"
    web_analytics_tag = "cee1c73f6e4743d0b5e6bb1a0bcaabcc"
    web_analytics_token = "021e1057c18547eca7b79f2516f06o7x"
  }
  deployment_configs = {
    preview = {
      ai_bindings = {
        AI_BINDING = {
          project_id = "some-project-id"
        }
      }
      always_use_latest_compatibility_date = false
      analytics_engine_datasets = {
        ANALYTICS_ENGINE_BINDING = {
          dataset = "api_analytics"
        }
      }
      browsers = {
        BROWSER = {

        }
      }
      build_image_major_version = 3
      compatibility_date = "2025-01-01"
      compatibility_flags = ["url_standard"]
      d1_databases = {
        D1_BINDING = {
          id = "445e2955-951a-43f8-a35b-a4d0c8138f63"
        }
      }
      durable_object_namespaces = {
        DO_BINDING = {
          namespace_id = "5eb63bbbe01eeed093cb22bb8f5acdc3"
        }
      }
      env_vars = {
        foo = {
          type = "plain_text"
          value = "hello world"
        }
      }
      fail_open = true
      hyperdrive_bindings = {
        HYPERDRIVE = {
          id = "a76a99bc342644deb02c38d66082262a"
        }
      }
      kv_namespaces = {
        KV_BINDING = {
          namespace_id = "5eb63bbbe01eeed093cb22bb8f5acdc3"
        }
      }
      limits = {
        cpu_ms = 100
      }
      mtls_certificates = {
        MTLS = {
          certificate_id = "d7cdd17c-916f-4cb7-aabe-585eb382ec4e"
        }
      }
      placement = {
        mode = "smart"
      }
      queue_producers = {
        QUEUE_PRODUCER_BINDING = {
          name = "some-queue"
        }
      }
      r2_buckets = {
        R2_BINDING = {
          name = "some-bucket"
          jurisdiction = "eu"
        }
      }
      services = {
        SERVICE_BINDING = {
          service = "example-worker"
          entrypoint = "MyHandler"
          environment = "production"
        }
      }
      usage_model = "standard"
      vectorize_bindings = {
        VECTORIZE = {
          index_name = "my_index"
        }
      }
      wrangler_config_hash = "abc123def456"
    }
    production = {
      ai_bindings = {
        AI_BINDING = {
          project_id = "some-project-id"
        }
      }
      always_use_latest_compatibility_date = false
      analytics_engine_datasets = {
        ANALYTICS_ENGINE_BINDING = {
          dataset = "api_analytics"
        }
      }
      browsers = {
        BROWSER = {

        }
      }
      build_image_major_version = 3
      compatibility_date = "2025-01-01"
      compatibility_flags = ["url_standard"]
      d1_databases = {
        D1_BINDING = {
          id = "445e2955-951a-43f8-a35b-a4d0c8138f63"
        }
      }
      durable_object_namespaces = {
        DO_BINDING = {
          namespace_id = "5eb63bbbe01eeed093cb22bb8f5acdc3"
        }
      }
      env_vars = {
        foo = {
          type = "plain_text"
          value = "hello world"
        }
      }
      fail_open = true
      hyperdrive_bindings = {
        HYPERDRIVE = {
          id = "a76a99bc342644deb02c38d66082262a"
        }
      }
      kv_namespaces = {
        KV_BINDING = {
          namespace_id = "5eb63bbbe01eeed093cb22bb8f5acdc3"
        }
      }
      limits = {
        cpu_ms = 100
      }
      mtls_certificates = {
        MTLS = {
          certificate_id = "d7cdd17c-916f-4cb7-aabe-585eb382ec4e"
        }
      }
      placement = {
        mode = "smart"
      }
      queue_producers = {
        QUEUE_PRODUCER_BINDING = {
          name = "some-queue"
        }
      }
      r2_buckets = {
        R2_BINDING = {
          name = "some-bucket"
          jurisdiction = "eu"
        }
      }
      services = {
        SERVICE_BINDING = {
          service = "example-worker"
          entrypoint = "MyHandler"
          environment = "production"
        }
      }
      usage_model = "standard"
      vectorize_bindings = {
        VECTORIZE = {
          index_name = "my_index"
        }
      }
      wrangler_config_hash = "abc123def456"
    }
  }
  source = {
    config = {
      deployments_enabled = true
      owner = "my-org"
      owner_id = "12345678"
      path_excludes = ["string"]
      path_includes = ["string"]
      pr_comments_enabled = true
      preview_branch_excludes = ["string"]
      preview_branch_includes = ["string"]
      preview_deployment_setting = "all"
      production_branch = "main"
      production_deployments_enabled = true
      repo_id = "12345678"
      repo_name = "my-repo"
    }
    type = "github"
  }
}
