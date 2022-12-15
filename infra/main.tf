resource "google_cloud_run_v2_service" "backend" {
  name     = "${var.backend_name}-service"
  location = var.location
  ingress  = var.ingress
  template {
    containers {
      image = var.backend_image
      env {
        name  = var.backend_cors_allow_origins_env_name
        value = var.backend_cors_allow_origins_env_value
      }
      resources {
        limits   = var.backend_resource_limits
        cpu_idle = var.backend_cpu_idle
      }
    }
    scaling {
      min_instance_count = var.backend_min_instances
      max_instance_count = var.backend_max_instances
    }
  }
}

resource "google_cloud_run_service_iam_binding" "backend_iam_binding" {
  location = google_cloud_run_v2_service.backend.location
  service  = google_cloud_run_v2_service.backend.name
  role     = var.iam_binding_role
  members  = var.iam_binding_members
}

resource "google_cloud_run_domain_mapping" "backend_domain_mapping" {
  location = var.location
  name     = var.backend_domain
  metadata {
    namespace = var.project
  }
  spec {
    route_name = google_cloud_run_v2_service.backend.name
  }
}

resource "google_cloud_run_v2_service" "frontend" {
  name     = "${var.frontend_name}-service"
  location = var.location
  ingress  = var.ingress
  template {
    containers {
      image = var.frontend_image
      env {
        name  = var.frontend_host_env_name
        value = var.frontend_host_env_value
      }
      resources {
        limits   = var.frontend_resource_limits
        cpu_idle = var.frontend_cpu_idle
      }
    }
    scaling {
      min_instance_count = var.frontend_min_instances
      max_instance_count = var.frontend_max_instances
    }
  }
}

resource "google_cloud_run_service_iam_binding" "frontend_iam_binding" {
  location = google_cloud_run_v2_service.frontend.location
  service  = google_cloud_run_v2_service.frontend.name
  role     = var.iam_binding_role
  members  = var.iam_binding_members
}

resource "google_cloud_run_domain_mapping" "frontend_domain_mapping" {
  location = var.location
  name     = var.frontend_domain
  metadata {
    namespace = var.project
  }
  spec {
    route_name = google_cloud_run_v2_service.frontend.name
  }
}
