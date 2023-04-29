resource "kubernetes_secret" "backend_service_secret" {
  metadata {
    name = local.backend_service_secret
  }

  data = {
    JWT_SECRET       = random_password.jwt_secret.result
    OAUTH2_CLIENT_ID = var.oauth2_client_id
    OAUTH2_SECRET    = var.oauth2_secret
    MONGO_URL        = var.mongo_url
    SS_API_KEY       = var.ss_api_key
  }
}

resource "random_password" "jwt_secret" {
  length  = 64
  special = false
  lower   = false
}