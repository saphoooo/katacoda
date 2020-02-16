## Kubernetes

Il est temps de déployer notre application dans Kubernetes.

Commençons par un peu de ménage :

`docker rm -f loto`{{execute}}

`docker image prune`{{execute}}

Pour déployer notre image, il faut écrire un fichier de déploiement en yaml. En se rendant sur la [documentation de Kubernetes](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/), nous pouvons trouver un exemple de déploiement :

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

Nous allons modifier ce fichier pour nos besoin. Créez le fichier `deploiement-loto.yaml` pour y placer le contenu suivant :

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

Ici nous sommes confronté à un problème : notre image est locale à la machine, et minikube ne peut pas accéder aux images de notre machine...

Quelles possibilités cela nous laisse-t-il ?

- Pousser l'image dans une Registry publique (dockerhub.io, gcr.io, quay.io, ...) ou privée pour pouvoir y accéder depuis minikube ; pas génial pour faire de l'expérimentation locale.
- Utiliser les [Buildpacks](https://buildpacks.io/). C'est de loin la solution que j'affectionne le plus : non seuelement nous allons pouvoir créer notre image directement dans minikube, mais en plus nous allons pouvoir nous oter la peine d'avoir à créer un Dockerfile.

Vous ne me croyez pas ?
