C'est le moment où skaffold entre en jeu.

Commençons pas télécharger skaffold :

`curl -Lo skaffold https://storage.googleapis.com/skaffold/releases/v1.3.1/skaffold-linux-amd64 && chmod +x skaffold && sudo mv skaffold /usr/local/bin`{{execute}}

Ensuite, initialisons skaffold :

`skaffold init`{{execute}}

Et validez la configuration.

Nous allons maintenant l'éditer afin d'utiliser buildpack au lieu de Dockerfile. C'est très simple, remplacer dans le fichier `skaffold.yaml` la patie build

```
build:
  artifacts:
  - image: loto
```

par celle-ci

```
build:
  tagPolicy:
    sha256: {}
  artifacts:
  - image: loto
    buildpack:
      builder: "cloudfoundry/cnb:tiny"
```{{copy}}

Afin que buildpack puisse reconnaître le runtime que nous utilisons (ici Go), nous devons générer le fichier de dépendances.

`cd ..`{{execute}} 

`go mod init loto`{{execute}}

`mv go.mod loto/`{{execute}}

`cd loto`{{execute}}

Maintenant noous sommes prêt à lancer notre déploiement !

`skaffold run`{{execute}}

Plus qu'à regarder l'application tourner dans Kubernetes :

`kubectl get pod -l app=loto`{{execute}}

`kubectl logs -l app=loto`{{execute}}

