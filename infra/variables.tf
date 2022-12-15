variable "project" {
  type    = string
  default = "xoq-infrastructure"
}

variable "location" {
  type    = string
  default = "europe-west1"
}

variable "bucket_storage_class" {
  type    = string
  default = "STANDARD"
}

variable "bucket_force_destroy" {
  type    = bool
  default = false
}

variable "state_location" {
  type    = string
  default = "us-central1"
}

variable "redis_name" {
  type    = string
  default = "redis"
}

variable "redis_memory_size_gb" {
  type    = number
  default = 1
}

variable "ingress" {
  type    = string
  default = "INGRESS_TRAFFIC_ALL"
}

variable "iam_binding_role" {
  type    = string
  default = "roles/run.invoker"
}

variable "iam_binding_members" {
  type    = list(string)
  default = [
    "allUsers"
  ]
}

variable "backend_name" {
  type    = string
  default = "backend"
}

variable "backend_image" {
  type    = string
  default = "europe-docker.pkg.dev/xoq-infrastructure/eu.gcr.io/xoq-backend:latest"
}

variable "backend_resource_limits" {
  type    = map(string)
  default = {
    "cpu" : "1"
    "memory" : "128Mi"
  }
}

variable "backend_cpu_idle" {
  type    = bool
  default = true
}

variable "backend_min_instances" {
  type    = number
  default = 1
}

variable "backend_max_instances" {
  type    = number
  default = 2
}

variable "backend_domain" {
  type    = string
  default = "api.xoq-idpa.com"
}

variable "backend_cors_allow_origins_env_name" {
  type    = string
  default = "CORS_ALLOW_ORIGINS"
}

variable "backend_cors_allow_origins_env_value" {
  type    = string
  default = "https://xoq-idpa.com"
}

variable "secret_id_redis_host" {
  type    = string
  default = "projects/890972837267/secrets/xoq-redis-host"
}

variable "backend_redis_host_env_name" {
  type    = string
  default = "REDIS_HOST"
}

variable "secret_id_redis_user" {
  type    = string
  default = "projects/890972837267/secrets/xoq-redis-user"
}

variable "backend_redis_user_env_name" {
  type    = string
  default = "REDIS_USER"
}

variable "secret_id_redis_password" {
  type    = string
  default = "projects/890972837267/secrets/xoq-redis-password"
}

variable "backend_redis_password_env_name" {
  type    = string
  default = "REDIS_PASSWORD"
}

variable "frontend_name" {
  type    = string
  default = "frontend"
}

variable "frontend_image" {
  type    = string
  default = "europe-docker.pkg.dev/xoq-infrastructure/eu.gcr.io/xoq-frontend:latest"
}

variable "frontend_resource_limits" {
  type    = map(string)
  default = {
    "cpu" : "1"
    "memory" : "128Mi"
  }
}

variable "frontend_cpu_idle" {
  type    = bool
  default = true
}

variable "frontend_min_instances" {
  type    = number
  default = 1
}

variable "frontend_max_instances" {
  type    = number
  default = 2
}

variable "frontend_host_env_name" {
  type    = string
  default = "HOST"
}

variable "frontend_host_env_value" {
  type    = string
  default = "xoq-idpa.com"
}

variable "frontend_domain" {
  type    = string
  default = "xoq-idpa.com"
}
