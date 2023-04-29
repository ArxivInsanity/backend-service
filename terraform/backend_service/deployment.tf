resource "kubernetes_deployment" "backend_service_deployment" {
  depends_on = [kubernetes_secret.backend_service_secret, kubernetes_config_map.backend_service_config_map]
  metadata {
    name = local.backend_deployment_label
    labels = {
      App = local.backend_deployment_label
    }
  }

  spec {
    replicas = 1
    selector {
      match_labels = {
        App = local.backend_deployment_label
      }
    }
    template {
      metadata {
        labels = {
          App = local.backend_deployment_label
        }
      }
      spec {
        container {
          image             = var.backend_service_image
          name              = local.backend_service_label
          image_pull_policy = "Always"

          port {
            container_port = local.backend_service_port
          }

          resources {
            limits = {
              cpu    = "0.5"
              memory = "512Mi"
            }
            requests = {
              cpu    = "250m"
              memory = "50Mi"
            }
          }
          env_from {
            secret_ref {
              name = local.backend_service_secret
            }
            config_map_ref {
              name = local.backend_service_config_map
            }
          }
          readiness_probe {
            http_get {
              path = "/"
              port = local.backend_service_port
            }
          }
          liveness_probe {
            http_get {
              path = "/"
              port = local.backend_service_port
            }
          }
        }
      }
    }
  }
}
