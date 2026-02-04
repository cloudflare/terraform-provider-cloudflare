resource "cloudflare_ai_search_instance" "example_ai_search_instance" {
  account_id = "c3dc5f0b34a14ff8e1b3ec04895e1b22"
  id = "my-ai-search"
  source = "source"
  type = "r2"
  ai_gateway_id = "ai_gateway_id"
  aisearch_model = "@cf/meta/llama-3.3-70b-instruct-fp8-fast"
  chunk = true
  chunk_overlap = 0
  chunk_size = 64
  custom_metadata = [{
    data_type = "text"
    field_name = "x"
  }]
  embedding_model = "@cf/qwen/qwen3-embedding-0.6b"
  hybrid_search_enabled = true
  max_num_results = 1
  metadata = {
    created_from_aisearch_wizard = true
    worker_domain = "worker_domain"
  }
  public_endpoint_params = {
    authorized_hosts = ["string"]
    chat_completions_endpoint = {
      disabled = true
    }
    enabled = true
    mcp = {
      disabled = true
    }
    rate_limit = {
      period_ms = 60000
      requests = 1
      technique = "fixed"
    }
    search_endpoint = {
      disabled = true
    }
  }
  reranking = true
  reranking_model = "@cf/baai/bge-reranker-base"
  rewrite_model = "@cf/meta/llama-3.3-70b-instruct-fp8-fast"
  rewrite_query = true
  score_threshold = 0
  source_params = {
    exclude_items = ["/admin/**", "/private/**", "**\\temp\\**"]
    include_items = ["/blog/**", "/docs/**/*.html", "**\\blog\\**.html"]
    prefix = "prefix"
    r2_jurisdiction = "r2_jurisdiction"
    web_crawler = {
      parse_options = {
        include_headers = {
          foo = "string"
        }
        include_images = true
        specific_sitemaps = ["https://example.com/sitemap.xml", "https://example.com/blog-sitemap.xml"]
        use_browser_rendering = true
      }
      parse_type = "sitemap"
      store_options = {
        storage_id = "storage_id"
        r2_jurisdiction = "r2_jurisdiction"
        storage_type = "r2"
      }
    }
  }
  token_id = "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"
}
