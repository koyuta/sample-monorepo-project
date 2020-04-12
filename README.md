# monorepo

### Docker for Mac

Run applications:

```
$ skaffold dev -p dfm
```

### Minikube

Enable ingress-nginx controller:

```
$ minikube addons enable ingress
```

If you want to see pod metricses, enable the metrics-server addon.

```
$ minikube addons enable metrics-server
```

Run applications:

```
$ skaffold dev -p minikube
```
