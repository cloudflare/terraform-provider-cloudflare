resource "cloudflare_workers_kv_namespace" "%[1]s_kv_namespace" {
  account_id = "%[2]s"
  title = "tfacctest-pages-project-kv-namespace"
}

resource "cloudflare_pages_project" "%[1]s" {
  account_id        = "%[2]s"
  name              = "%[3]s"
  production_branch = "main"
  deployment_configs = {
    preview = {
      env_vars = {
        ENVIRONMENT = {
          type = "plain_text"
          value = "preview"
        }
        TURNSTILE_SECRET = {
          type = "secret_text"
          value = "1x0000000000000000000000000000000AA"
        }
      }
      kv_namespaces = {
        KV_BINDING = {
          namespace_id = cloudflare_workers_kv_namespace.%[1]s_kv_namespace.id
        }
      }
      durable_object_namespaces = {
        DO_BINDING = {
          namespace_id = "5eb63bbbe01eeed093cb22bb8f5acdc3"
        }
      }
      r2_buckets = {
        R2_BINDING = {
          name = "some-bucket"
        }
      }
      d1_databases = {
        D1_BINDING = {
          id = "445e2955-951a-4358-a35b-a4d0c813f63"
        }
      }
      service_binding = [{
        name        = "MY_SERVICE_BINDING"
        service     = "my-service"
        environment = "preview"
      }]
      compatibility_date                   = "2022-08-15"
      compatibility_flags                  = ["preview_flag"]
      fail_open                            = true
      always_use_latest_compatibility_date = true
      usage_model                          = "standard"
    }
    production = {
      env_vars = {
        ENVIRONMENT = {
          type = "plain_text"
          value = "production"
        }
        OTHER_VALUE = {
          type = "plain_text"
          value = "other value"
        }
        TURNSTILE_SECRET = {
          type = "secret_text"
          value = "1x0000000000000000000000000000000AA"
        }
        TURNSTILE_INVIS_SECRET = {
          type = "secret_text"
          value = "2x0000000000000000000000000000000AA"
        }
      }
      kv_namespaces = {
        KV_BINDING_1 = {
          namespace_id = cloudflare_workers_kv_namespace.%[1]s_kv_namespace.id
        }
        KV_BINDING_2 = {
          namespace_id = cloudflare_workers_kv_namespace.%[1]s_kv_namespace.id
        }
      }
      durable_object_namespaces = {
        DO_BINDING_1 = {
          namespace_id = "5eb63bbbe01eeed093cb22bb8f5acdc3"
        }
        DO_BINDING_2 = {
          namespace_id = "3cdca5f8bb22bc390deee10ebbb36be5"
        }
      }
      r2_buckets = {
        R2_BINDING_1 = {
          name = "some-bucket"
        }
        R2_BINDING_2 = {
          name = "other-bucket"
        }
      }
      d1_databases = {
        D1_BINDING_1 = {
          id = "445e2955-951a-4358-a35b-a4d0c813f63"
        }
        D1_BINDING_2 = {
          id = "a399414b-c697-409a-a688-377db6433cd9"
        }
      }
      service_binding = [{
        name        = "MY_SERVICE_BINDING"
        service     = "my-service"
        environment = "production"
      }]
      compatibility_date                   = "2022-08-16"
      compatibility_flags                  = ["production_flag", "second flag"]
      fail_open                            = true
      always_use_latest_compatibility_date = false
      usage_model                          = "standard"
      placement = {
        mode = "smart"
      }
    }
  }
}
