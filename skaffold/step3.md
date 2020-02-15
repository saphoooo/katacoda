## Kubernetes

Maintenant que notre image est prête, il est temps de la déployer dans Kubernetes.

Commençons par un peu de ménage :

`docker rm -f loto`{{execute}}

`docker image prune`{{execute}}

Pour déployer notre image, il faut écrire un fichier yaml de déploiement. En se rendant sur la page de [documentation de Kubernetes](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/), nous pouvons trouver un exemple de déploiement :

```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
  labels:
    app: nginx
spec:
  replicas: 3
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.7.9
        ports:
        - containerPort: 80
```

Nous allons le modifier pour nos besoin. Créez le fichier `deploiement-loto.yaml` pour y placer le contenu suivant :

```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: loto-deployment
  labels:
    app: loto
spec:
  replicas: 1
  selector:
    matchLabels:
      app: loto
  template:
    metadata:
      labels:
        app: loto
    spec:
      containers:
      - name: loto
        image: loto
```{{copy}}

Ici nous avons remplacé toutes les occurences de nginx par loto, et passé le nombre de réplicas à 1, rien de plus.

A présent, tentons de le déployer :

`kubectl apply -f deploiement-loto.yaml`{{execute}}

`curl -Lo skaffold https://storage.googleapis.com/skaffold/releases/v1.3.1/skaffold-linux-amd64 && chmod +x skaffold && sudo mv skaffold /usr/local/bin`{{execute}}

```
apiVersion: skaffold/v2alpha3
kind: Config
metadata:
  name: loto
build:
  artifacts:
  - image: loto
    buildpack:
      builder: "heroku/buildpacks"
deploy:
  kubectl:
    manifests:
    - loto-deploy.yaml
```{{copy}}

