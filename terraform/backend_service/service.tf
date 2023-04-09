resource "kubernetes_service" "backend_service_service" {
  metadata {
    name = local.backend_service_label
  }
  spec {
    selector = {
      App = local.backend_deployment_label
    }
    port {
      port        = local.backend_service_port
      target_port = local.backend_service_port
    }

    type = "NodePort"
  }
}
