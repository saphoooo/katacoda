C'est le moment où skaffold entre en jeu.

Commençons pas télécharger skaffold :

`curl -Lo skaffold https://storage.googleapis.com/skaffold/releases/v1.3.1/skaffold-linux-amd64 && chmod +x skaffold && sudo mv skaffold /usr/local/bin`{{execute}}

Ensuite, initialisons skaffold :

`skaffold init`{{execute}}

Validez la configuration.