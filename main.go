package main

import (
	"fmt"
)

var teclas = map[int]string{
	0:      " ",
	1:      ".",
	11:     ",",
	111:    ":",
	1111:   ";",
	20:     "A",
	220:    "B",
	2220:   "C",
	22220:  "Á",
	30:     "D",
	330:    "E",
	3330:   "F",
	33330:  "É",
	40:     "G",
	440:    "H",
	4440:   "I",
	44440:  "Í",
	50:     "J",
	550:    "K",
	5550:   "L",
	60:     "M",
	660:    "N",
	6660:   "O",
	66660:  "Ó",
	666660: "Ñ",
	70:     "P",
	770:    "Q",
	7770:   "R",
	77770:  "S",
	80:     "T",
	880:    "U",
	8880:   "V",
	88880:  "Ú",
	90:     "W",
	990:    "X",
	9990:   "Y",
	99990:  "Z",
	2:      "a",
	22:     "b",
	222:    "c",
	2222:   "á",
	3:      "d",
	33:     "e",
	333:    "f",
	3333:   "é",
	4:      "g",
	44:     "h",
	444:    "i",
	4444:   "í",
	5:      "j",
	55:     "k",
	555:    "l",
	6:      "m",
	66:     "n",
	666:    "o",
	6666:   "ó",
	66666:  "ñ",
	7:      "p",
	77:     "q",
	777:    "r",
	7777:   "s",
	8:      "t",
	88:     "u",
	888:    "v",
	8888:   "ú",
	9:      "w",
	99:     "x",
	999:    "y",
	9999:   "z",
}

var frase = "Eliseo Damián Esteverena"

func main() {
	devolverNumeros()
}

// 1. LEÉ LA FRASE Y DEVUELVE UN SLICE CON LOS NÚMEROS CORRESPONDIENTES A CADA CARACTER.
// 2. ENVIA A LA FUNCIÓN devolverLetras() EL SLICE.
func devolverNumeros() {
	slice := []int{}
	// 1:
	for _, l := range frase {
		for n, ltr := range teclas {
			if string(l) == ltr {
				slice = append(slice, n)
			}
		}
	}
	fmt.Println(slice)
	// 2:
	devolverLetras(slice)
}

// DEVUELVE LA CADENA DE TEXTO A PARTIR DEL SLICE.
func devolverLetras(slice []int) {
	indice := len(slice) - 1       // CANTIDAD DE CARACTERES.
	a := ""                        // STRING VACIO.
	for i := 0; i <= indice; i++ { // SE EJECUTA LA MISMA CANTIDAD DE VECES COMO CANTIDAD DE CARACTERES.
		for k, n := range teclas {
			if slice[i] == k { // EVALÚA SI EL VALOR DEL SLICE ES IGUAL AL VALOR DEL ELEMENTE DEL MAPA QUE SE ESTÁ LEYENDO.
				a = a + n // SE AÑADE A LA VARIABLE 'a' EL CARACTER CORRESPONDIENTE.
			}
		}
	}
	fmt.Println(a) // SE IMPRIME LA VARIABLE CON LA FRASE 'TRADUCIDA'
}
