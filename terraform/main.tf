terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "3.52.0"
    }
    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = ">= 2.0.1"
    }
  }
  cloud {
    hostname     = "app.terraform.io"
    organization = "Arxiv-Insanity"
    workspaces {
      name = "app-deployment"
    }
  }
}

provider "google" {
  project = var.project
  region  = var.region
  zone    = var.zone
}

data "terraform_remote_state" "gke" {
  backend = "remote"
  config = {
    organization = "Arxiv-Insanity"
    workspaces = {
      name = "app-infra"
    }
  }
}

data "google_client_config" "default" {}

data "google_container_cluster" "my_cluster" {
  name     = data.terraform_remote_state.gke.outputs.gke_outputs.cluster_name
  location = data.terraform_remote_state.gke.outputs.gke_outputs.location
}

provider "kubernetes" {
  host = "https://${data.terraform_remote_state.gke.outputs.gke_outputs.cluster_host}"

  token                  = data.google_client_config.default.access_token
  cluster_ca_certificate = base64decode(data.google_container_cluster.my_cluster.master_auth[0].cluster_ca_certificate)
}

module "backend_service" {
  source = "./backend_service"

  jwt_secret            = var.jwt_secret
  oauth2_client_id      = var.oauth2_client_id
  oauth2_secret         = var.oauth2_secret
  mongo_url             = var.mongo_url
  backend_service_image = var.backend_service_image
}