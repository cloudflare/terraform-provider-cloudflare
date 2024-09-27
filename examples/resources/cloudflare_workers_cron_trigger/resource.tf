resource "cloudflare_workers_script" "example_script" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "example-script"
  content    = file("path/to/my.js")
}

resource "cloudflare_workers_cron_trigger" "example_trigger" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  script_name = cloudflare_workers_script.example_script.name
  schedules = [
    "*/5 * * * *",      # every 5 minutes
    "10 7 * * mon-fri", # 7:10am every weekday
  ]
}
