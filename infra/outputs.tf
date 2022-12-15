output "backend-uri" {
  value = google_cloud_run_v2_service.backend.uri
}

output "frontend-uri" {
  value = google_cloud_run_v2_service.frontend.uri
}
