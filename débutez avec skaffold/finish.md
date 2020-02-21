Skaffold offre quelques autres options intéressantes, mais pour une première prise en main nous avons atteind notre objectif : avoir notre code déployé automatiquement dans Kubernetes.

Si skaffold ne nous affranchi pas de devoir décider de la manière de faire le build de notre conteneur (Dockerfile, Jib, Kaniko, Buildpack), ni de devoir créer nos fichiers de déploiements, il s'occupe à merveille de tout le reste.

Parce que maintenant que le boilerplate est en place, vous pouvez simplement modifier votre code, l'enregistrer, faire un `git commit`, et exécuter de nouveau `skaffold run` pour qu'il soit déployer dans Kubernetes. Vous pouvez même exécuter `skaffold dev` pour faire cette action en continu !