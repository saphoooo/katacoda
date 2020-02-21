Vous commencez le développement d'une toute nouvelle application révolutionnaire (je vous le souhaite en tout cas), et vous êtes impatient de la voir tourner !



```
cat << EOF >> deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: myapp
  labels:
    app: myapp
spec:
  replicas: 1
  selector:
    matchLabels:
      app: myapp
  template:
    metadata:
      labels:
        app: myapp
    spec:
      containers:
      - name: myapp
          imagePullPolicy: Always
          image: docker.io/saphoooo/myapp:dirty
          resources:
          requests:
            memory: "64Mi"
            cpu: "100m"
          limits:
            memory: "128Mi"
            cpu: "200m"

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
  name: myapp
spec:
  template:
    spec:
      replicas: 1
      containers:
      - name: myapp
      resources:
        requests:
          memory: "128Mi"
          cpu: "250m"
        limits:
          memory: "256Mi"
          cpu: "500m"
        
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