C'est le moment où skaffold entre en jeu.

Commençons pas télécharger skaffold :

`curl -Lo skaffold https://storage.googleapis.com/skaffold/releases/v1.3.1/skaffold-linux-amd64 && chmod +x skaffold && sudo mv skaffold /usr/local/bin`{{execute}}

Ensuite, initialisons skaffold :

`skaffold init`{{execute}}

Et validez la configuration.

Nous allons maintenant l'éditer afin d'utiliser buildpack au lieu de Dockerfile. C'est très simple, remplacer la patie build

```
build:
  artifacts:
  - image: loto
```

par celle-ci

```
  tagPolicy:
    sha256: {}
  artifacts:
  - image: loto
    buildpack:
      builder: "cloudfoundry/cnb:tiny"
```

C'est tout. Prêt à lancer un déploiement ?

`skaffold run`{{execute}}

Plus qu'à regarder l'application tourner dans Kubernetes :

`kubectl get pod -l app=loto`{{execute}}

`kubectl logs -l app=loto`{{execute}}

