# monorepo

## Build container image

```
docker build \
  -t project-1/api \
  -f ./projects/project-1/application/Dockerfile \
  ./projects/project-1/application
```

## Deploy to local

### Docker for mac

```
kubectl kustomize ./deployments/k8s/overlays/local/docker-for-mac
kubectl apply -k ./deployments/k8s/overlays/local/docker-for-mac
```
