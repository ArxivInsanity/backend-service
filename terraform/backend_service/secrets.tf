resource "kubernetes_secret" "backend_service_secret" {
  metadata {
    name = local.backend_service_secret
  }

  data = {
    JWT_SECRET       = var.jwt_secret
    OAUTH2_CLIENT_ID = var.oauth2_client_id
    OAUTH2_SECRET    = var.oauth2_secret
    MONGO_URL        = var.mongo_url
    SS_API_KEY       = var.ss_api_key
  }
}