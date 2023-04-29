resource "kubernetes_config_map" "backend_service_config_map" {
  metadata {
    name = local.backend_service_config_map
  }

  data = {
    GRAPH_SERVICE_ENDPOINT = var.graph_service_endpoint
  }
}