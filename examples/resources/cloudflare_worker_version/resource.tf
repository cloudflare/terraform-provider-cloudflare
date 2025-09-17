resource "cloudflare_worker_version" "example_worker_version" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  worker_id = "023e105f4ecef8ad9ca31a8372d0c353"
  annotations = {
    workers_message = "Fixed bug."
    workers_tag = "v1.0.1"
  }
  assets = {
    config = {
      html_handling = "auto-trailing-slash"
      not_found_handling = "404-page"
      run_worker_first = ["string"]
    }
    jwt = "jwt"
  }
  bindings = [{
    name = "MY_ENV_VAR"
    text = "my_data"
    type = "plain_text"
  }]
  compatibility_date = "2021-01-01"
  compatibility_flags = ["nodejs_compat"]
  limits = {
    cpu_ms = 50
  }
  main_module = "index.js"
  migrations = {
    deleted_classes = ["string"]
    new_classes = ["string"]
    new_sqlite_classes = ["string"]
    new_tag = "v2"
    old_tag = "v1"
    renamed_classes = [{
      from = "from"
      to = "to"
    }]
    transferred_classes = [{
      from = "from"
      from_script = "from_script"
      to = "to"
    }]
  }
  modules = [{
    content_file = "dist/index.js"
    content_type = "application/javascript+module"
    name = "index.js"
  }]
  placement = {
    mode = "smart"
  }
}
