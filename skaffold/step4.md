C'est le moment où skaffold entre en jeu.

Tout d'abord, prenons un malin plaisir à supprimer ce Dockerfile comme un vestige du passé :

`rm -fr Dockerfile`

Skaffold, pour faire court, est un petit outil de CI/CD très simple d'accès. Il nous permet, en une seul commande, de faire le build (CI), puis le déploiement (CD) de notre application.

Commençons pas télécharger skaffold :

`curl -Lo skaffold https://storage.googleapis.com/skaffold/releases/v1.4.0/skaffold-linux-amd64 && chmod +x skaffold && sudo mv skaffold /usr/local/bin`{{execute}}

Ensuite, initialisons skaffold dans le répertoire où se trouve nos fichiers :

`skaffold init --skip-build`{{execute}}

Et validez la configuration qu'il vous propose.

Nous allons maintenant éditer cette configuration afin d'utiliser les buildpacks. C'est très simple : dans le fichier `skaffold.yaml` ajoutez la partie build :

```
cat << EOF >> skaffold.yaml
build:
  tagPolicy:
    sha256: {}
  artifacts:
  - image: loto
    buildpack:
      builder: "cloudfoundry/cnb:tiny"
EOF
```{{copy}}

Afin que buildpack puisse automatiquement reconnaître le runtime que nous utilisons (ici Go), nous devons générer le fichier de dépendances.

`cd ..`{{execute}} 

`go mod init loto`{{execute}}

`mv go.mod loto/`{{execute}}

`cd loto`{{execute}}

Tout les éléments sont maintenant en place, il est temps de se jeter à l'eau :

`skaffold run`{{execute}}

C'est tout. Il ne nous reste plus qu'à regarder l'application tourner dans Kubernetes :

`kubectl get pod -l app=loto`{{execute}}

`kubectl logs -l app=loto`{{execute}}

