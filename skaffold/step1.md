## Le code

Le code de votre application est dans votre *home*:

`ls ~`{{execute}}

C'est du Go.

`kubectl` et `minikube` sont déjà téléchargés, et `minikube` est déjà démarré (peut-être):

`minikube status`{{execute}}

Si `host`, `kubelet`, et `apiserver` ne sont pas encore `running`, il vous faudra encore un peu de partience.

## Dockerfile

Comment, à patir de notre code, allons-nous créer une image Docker ?

Tout d'abord, mettons par convention notre code dans un répertoire portant le nom de l'application que nous développons ; ici, c'est une application pour générer des grilles du loto :

`mkdir loto`{{execute}}

`mv main.go loto`{{execute}}

`cd loto`{{execute}}

Pour créer une image, la méthode la plus répandue à ce jour c'est d'utiliser un Dockerfile : un fichier qui décrit les étapes d'installation de notre application en somme.

Essayons de faire ça de manière empirique :

`nano Dockerfile`{{execute}}

```
FROM golang

ADD main.go

RUN GCO_ENABLE=0 go build -o loto

CMD loto
```{{copy}}

Nous allons utiliser `minikube` pour le build de l'image :

`eval $(minikube docker-env)`{{execute}}

`docker build -t loto .`{{execute}}