resource "kubernetes_ingress_v1" "backend_ingress" {
  wait_for_load_balancer = true
  metadata {
    name = local.backend_service_ingress
    annotations = {
      "kubernetes.io/ingress.class" = "gce"
    }
  }
  spec {
    rule {
      http {
        path {
          path = "/*"
          backend {
            service {
              name = local.backend_service_label
              port {
                number = local.backend_service_port
              }
            }
          }
        }
      }
    }
  }
}