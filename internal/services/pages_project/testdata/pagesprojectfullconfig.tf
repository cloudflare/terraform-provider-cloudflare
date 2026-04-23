resource "cloudflare_workers_kv_namespace" "%[1]s_kv_namespace" {
  account_id = "%[2]s"
  title = "tfacctest-pages-project-kv-namespace"
}

resource "cloudflare_pages_project" "%[1]s" {
	account_id = "%[2]s"
	name = "%[3]s"
	production_branch = "main"
	
	build_config = {
		build_caching = true
		build_command = "npm run build"
		destination_dir = "dist"
		root_dir = "/app"
		web_analytics_tag = "test-tag-123"
		web_analytics_token = "test-token-456"
	}

	deployment_configs = {
		preview = {
			compatibility_date = "2023-01-15"
			compatibility_flags = ["preview_flag_1", "preview_flag_2"]
			
			env_vars = {
				ENVIRONMENT = {
					type = "plain_text"
					value = "preview"
				}
				SECRET_KEY = {
					type = "secret_text"
					value = "preview-secret-123"
				}
			}
			
			kv_namespaces = {
				KV_PREVIEW = {
					namespace_id = cloudflare_workers_kv_namespace.%[1]s_kv_namespace.id
				}
			}
			
			d1_databases = {
				D1_PREVIEW = {
					id = "preview-d1-database-id"
				}
			}
			
			r2_buckets = {
				R2_PREVIEW = {
					name = "preview-bucket"
					jurisdiction = "eu"
				}
			}
			
			ai_bindings = {
				AI_PREVIEW = {
					project_id = "preview-ai-project-id"
				}
			}
			
			analytics_engine_datasets = {
				ANALYTICS_PREVIEW = {
					dataset = "preview-analytics-dataset"
				}
			}
				
			browsers = {
				BROWSER_PREVIEW = {}
			}
			
			hyperdrive_bindings = {
				HYPERDRIVE_PREVIEW = {
					id = "preview-hyperdrive-id"
				}
			}
			
			mtls_certificates = {
				MTLS_PREVIEW = {
					certificate_id = "preview-mtls-cert-id"
				}
			}
			
			queue_producers = {
				QUEUE_PREVIEW = {
					name = "preview-queue"
				}
			}
			
			services = {
				SERVICE_PREVIEW = {
					service = "preview-service"
					environment = "preview"
					entrypoint = "main"
				}
			}
			
			vectorize_bindings = {
				VECTORIZE_PREVIEW = {
					index_name = "preview-vector-index"
				}
			}
			
			placement = {
				mode = "smart"
			}
		}
		
		production = {
			compatibility_date = "2023-01-16"
			compatibility_flags = ["production_flag"]
			
			env_vars = {
				ENVIRONMENT = {
					type = "plain_text"
					value = "production"
				}
			}
			
			kv_namespaces = {
				KV_PROD = {
					namespace_id = cloudflare_workers_kv_namespace.%[1]s_kv_namespace.id
				}
			}
			
			placement = {
				mode = "smart"
			}
		}
	}

	source = {
		type = "github"
		config = {
			owner = "%[4]s"
			repo_name = "%[5]s"
			production_branch = "main"
			pr_comments_enabled = true
			deployments_enabled = true
			production_deployments_enabled = true
			preview_deployment_setting = "all"
			path_includes = ["src/**", "public/**"]
			path_excludes = ["*.test.js", "node_modules/**"]
			preview_branch_includes = ["dev", "staging"]
			preview_branch_excludes = ["main"]
		}
	}
}
