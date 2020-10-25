package main

import (
	"fmt"
	"os"
	"sort"
)

// OrdenAscendente ....
type OrdenAscendente []string

func (a OrdenAscendente) Len() int           { return len(a) }
func (a OrdenAscendente) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a OrdenAscendente) Less(i, j int) bool { return a[i] < a[j] }

func main() {
	var str string
	var words []string

	fmt.Print("\nIngresa todas las cadenas que gustes, separadas por enter. Para salir, ingresa 'salir'.\n\n")
	for {
		fmt.Scan(&str)

		if str == "salir" {
			break
		} else {
			words = append(words, str)
		}
	}

	// Sort ascendente y escritura en txt.
	sort.Sort(OrdenAscendente(words))
	file, err := os.Create("ascendente.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	for _, v := range words {
		file.WriteString(v + "\n")
	}

	// Sort en descendente y escritura en txt.
	sort.Sort(sort.Reverse(sort.StringSlice(words)))
	file2, err := os.Create("descendente.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file2.Close()

	for _, v := range words {
		file2.WriteString(v + "\n")
	}
}
