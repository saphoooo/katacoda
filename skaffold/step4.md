C'est le moment où skaffold entre en jeu.

Repenons depuis le début : nous allons créer un second répertoire :

`cd ~`{{execute}}

`mkdir skaffold`

`cp main.go skaffold`

`cd skaffold`

`skaffold init --skip-build=false`{{execute}}