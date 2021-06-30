# Projet-Forum

Le projet consistait a créer un forum web fonctionnel avec différents condition, comme créer des posts, des commentaires, pouvoir ce connecter a l'aide d'identifiant , liker et filtrer par catégories ou posts.

# Installation du projet

Avant de lancer le forum plusieurs installation sont nécéssaires :

## _Installation de sqlite_ :  
> sudo apt update  
sudo apt install gcc

> go get github.com/mattn/go-sqlite3  
export CGO_ENABLED=1

## _Installation de uuid_ :
> go get github.com/google/uuid

## _Installation de Bcrypt_ :
> go get golang.org/x/crypto/bcrypt

# Déployer le forum

Exécuter la commande 
```bash 
go run server.go
```

# Structure du projet


# Architecture du projet