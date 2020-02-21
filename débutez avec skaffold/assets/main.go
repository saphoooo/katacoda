package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/pkg/errors"
)

// sliceDeNombres retourne un slice de 1 à max nombres
func sliceDeNombres(max int) []int {
	var grid []int
	for n := 0; n < max; n++ {
		grid = append(grid, n+1)
	}
	return grid
}

// maGrille créer une grille de la taille désirée
// (1 à 49 pour la grille, 1 à 10 pour le numéro chance)
// de max numéros (5 pour une grille standard)
func maGrille(taille, max int) []int {
	if max == 0 {
		errors.New("max must be greater than 1")
	}
	rand.Seed(time.Now().UnixNano())
	base := sliceDeNombres(max)
	var grille []int
	for i := 0; i < taille; i++ {
		n := rand.Intn(len(base))
		grille = append(grille, base[n])
		base = append(base[:n], base[n+1:]...)
	}
	return grille
}

func main() {
	for {
		fmt.Printf("ma grille = %v\n", maGrille(5, 49))
		fmt.Printf("mon numéro chance = %v\n", maGrille(1, 10))
		time.Sleep(time.Second * 30)
	}
}
