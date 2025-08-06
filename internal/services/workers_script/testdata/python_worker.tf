resource "cloudflare_workers_script" "%[1]s" {
  account_id = "%[2]s"
  script_name = "%[1]s"
  content = <<EOT
from workers import Response
def on_fetch(request):
  return Response("Hello World")
EOT
  content_type = "text/x-python"
  main_module = "index.py"
  compatibility_date = "2025-07-22"
  compatibility_flags = ["python_workers"]
}
