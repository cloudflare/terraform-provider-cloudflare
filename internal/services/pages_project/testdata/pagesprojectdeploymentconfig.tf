
resource "cloudflare_pages_project" "%[1]s" {
  account_id        = "%[2]s"
  name              = "%[3]s"
  production_branch = "main"
  deployment_configs = {
    preview = {
      environment_variables = {
        ENVIRONMENT = "preview"
      }
      secrets = {
        TURNSTILE_SECRET = "1x0000000000000000000000000000000AA"
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
      service_binding = [{
        name        = "MY_SERVICE_BINDING"
        service     = "my-service"
        environment = "preview"
      }]
      compatibility_date                   = "2022-08-15"
      compatibility_flags                  = ["preview_flag"]
      fail_open                            = true
      always_use_latest_compatibility_date = true
      usage_model                          = "unbound"
    }
    production = {
      environment_variables = {
        ENVIRONMENT = "production"
        OTHER_VALUE = "other value"
      }
      secrets = {
        TURNSTILE_SECRET       = "1x0000000000000000000000000000000AA"
        TURNSTILE_INVIS_SECRET = "2x0000000000000000000000000000000AA"
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
      service_binding = [{
        name        = "MY_SERVICE_BINDING"
        service     = "my-service"
        environment = "production"
      }]
      compatibility_date                   = "2022-08-16"
      compatibility_flags                  = ["production_flag", "second flag"]
      fail_open                            = true
      always_use_latest_compatibility_date = false
      usage_model                          = "bundled"
      placement = {
        mode = "smart"
      }
    }
  }
}
