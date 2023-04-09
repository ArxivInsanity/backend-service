variable "jwt_secret" {
  description = "The secret for the JWT"
}
variable "oauth2_client_id" {
  description = "The client id for google oauth"
}
variable "oauth2_secret" {
  description = "The secret for google oauth"
}
variable "mongo_url" {
  description = "The url for handling mongo db connection"
}
variable "backend_service_image" {
  description = "The docker image for backend service application that should be deployed in kubernetes pod"
}