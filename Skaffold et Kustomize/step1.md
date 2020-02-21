```
cat << EOF >> deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: kustomize-test
  labels:
    app: kustomize-test
spec:
  replicas: 1
  selector:
    matchLabels:
      app: kustomize-test
  template:
    metadata:
      labels:
        app: kustomize-test
    spec:
      containers:
      - name: kustomize-test
        image: not/a/valid/image
EOF
```{{execute}}

```
cat << EOF >> kustomization.yaml
resources:
  - deployment.yaml
patches:
  - patch.yaml
EOF
```{{execute}}

```
cat << EOF >> patch.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: kustomize-test
spec:
  template:
    spec:
      containers:
      - name: kustomize-test
        image: docker.io/saphoooo/myapp:dirty
EOF
```{{execute}}

```
cat << EOF >> skaffold.yaml
apiVersion: skaffold/v2alpha4
kind: Config
metadata:
  name: myapp
deploy:
  kustomize: {}
EOF
```{{execute}}

```
curl -s "https://raw.githubusercontent.com/\
kubernetes-sigs/kustomize/master/hack/install_kustomize.sh"  | bash && chmod +x kustomize && sudo mv kustomize /usr/local/bin
```{{execute}}

```
curl -Lo skaffold https://storage.googleapis.com/skaffold/releases/v1.4.0/skaffold-linux-amd64 && chmod +x skaffold && sudo mv skaffold /usr/local/bin
```{{execute}}

```
skaffold run
```{{execute}}

```
kubectl get po
```{{execute}}