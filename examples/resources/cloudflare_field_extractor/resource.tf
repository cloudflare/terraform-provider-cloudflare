resource "cloudflare_field_extractor" "example_field_extractor" {
  account_id = "123456"
  extractor = "llm_prompts"
  rules = [{
    fields = [{
      expression = "x"
      name = "x"
    }]
    ref = "x"
    description = "description"
  }]
}
