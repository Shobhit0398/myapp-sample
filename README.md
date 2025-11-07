# myapp-sample (shobhitnagar)


This sample repo is a minimal Go web app with Prometheus metrics and Kubernetes manifests, plus a GitHub Actions workflow that builds, tests, pushes to Docker Hub and deploys to a Minikube cluster.


## What is included
- `main.go` — a simple HTTP server with `/` `/healthz` and `/metrics` endpoints
- `Dockerfile` — multi-stage build (static binary → distroless)
- `k8s/` — Deployment, Service, Ingress, ServiceMonitor
- `.github/workflows/ci-cd.yml` — CI/CD pipeline that pushes to Docker Hub and deploys to Kubernetes


## Quick start (local)
1. Clone this repo locally.
2. Build and run locally:
```bash
go build -o myapp .
./myapp
# open http://localhost:8080
