Vous commencez le développement d'une toute nouvelle application révolutionnaire (je vous le souhaite en tout cas), et vous êtes impatient de la voir tourner !

Mais bon, avant toute chose, il faut faire un build du code dans un container, puis le déployer. Fort heureusement, nous avons déjà eu l'occasion d'aborder tout ça dans [Débutez avec Skaffold](https://www.katacoda.com/saphoooo/scenarios/d%C3%A9butez-avec-skaffold), il ne s'agira donc que d'un rappel.

## Préparation

Pour nous lancer, nous avons besoin de deux-trois choses, et pour commencer, un fichier de déploiement pour Kubernetes :

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
        image: myapp
        resources:
          requests:
            memory: "64Mi"
            cpu: "100m"
          limits:
            memory: "128Mi"
            cpu: "200m"
EOF
```{{execute}}

De Skaffold, bien entendu !

```
curl -Lo skaffold https://storage.googleapis.com/skaffold/releases/v1.4.0/skaffold-linux-amd64 && chmod +x skaffold && sudo mv skaffold /usr/local/bin
```{{execute}}

Et nous allons de ce pas l'initialiser, faisant un `skip-build` afin de pouvoir le customiser et utiliser les buildpacks.

```
skaffold init --skip-build
```{{execute}}
hello wold

Insérons la partie build à `skaffold.yaml` :

```
cat << EOF >> skaffold.yaml
build:
  tagPolicy:
    sha256: {}
  artifacts:
  - image: myapp
    buildpack:
      builder: "cloudfoundry/cnb:tiny"
EOF
```{{execute}}

A ce stade, Skaffold a maintenant tout ce qu'il faut pour gérer notre build et notre déploiement :

```
skaffold run
```{{execute}}

```
kubectl get po
```{{execute}}

That's life!