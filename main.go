package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"sort"
	"time"
)

var (
	numCartas   int
	numSimbolos int
	numIntentos int
	verbose     bool
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
	flag.IntVar(&numCartas, "num-cartas", 80, "número de cartas en el mazo")
	flag.IntVar(&numSimbolos, "num-simbolos", 10, "número de diferentes tipos de cartas")
	flag.IntVar(&numIntentos, "num-intentos", 0, "cuántas veces se repetirá el experimento?")
	flag.BoolVar(&verbose, "v", false, "muestra informacion adicional")
	flag.Parse()

	if verbose {
		log.SetOutput(os.Stdout)

	} else {
		log.SetOutput(ioutil.Discard) // TODO mejorar rendimiento
	}
	log.SetFlags(0)

	if numIntentos == 0 {
		fmt.Println("número de intentos no válido, use -num-intentos")
		flag.PrintDefaults()
		os.Exit(-1)
	}
}

func main() {
	cartas := make([]byte, numCartas)
	cuenta := 0
	sum := 0

	tabla := make(map[int]int, numCartas)
	//resultados := make([]int, 1000)
	//for k, _ := range resultados {
	for k := 0; k < numIntentos; k++ {
		mezclarCartas(cartas, numSimbolos)
		log.Printf("mazo de cartas: %v\n", cartas)
		r := manotazo(cartas, numSimbolos)
		log.Printf("recoje %v cartas...\n", r)
		sum += r
		tabla[r]++
		if r == numCartas {
			log.Printf("intento %d, valor %d\n", k+1, r)
			cuenta++
		}
		//resultados[k] = r
	}
	//fmt.Printf("reultados: %v\n", resultados)
	fmt.Println("========RESULTADOS=========")
	mostrarTabla(tabla)
	fmt.Printf("se ganó el juego %d veces en %d intentos\n", cuenta, numIntentos)
	fmt.Printf("la probabilidad de ganar es: %f%% \n",
		float64(cuenta)/float64(numIntentos)*100)
	fmt.Printf("el promedio de cartas recogidas es de %f\n", float64(sum)/float64(numIntentos))

}

func mezclarCartas(cartas []byte, mod int) {
	for i, _ := range cartas {
		imod := i % mod
		j := rand.Intn(i + 1)
		cartas[i], cartas[j] = cartas[j], byte(imod)
	}
}

func manotazo(cartas []byte, mod int) (recoge int) {
	for i, carta := range cartas {
		imod := i % mod
		recoge++
		log.Printf("carta: %d, num: %d\n", carta, imod)
		if int(carta) == imod {
			break
		}

	}
	return recoge

}

func mostrarTabla(m map[int]int) {
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	fmt.Printf("|carta\t| valor\t|\n|-------|-------|\n")
	for _, k := range keys {
		fmt.Printf("|%v\t|%v\t|\n", k, m[k])
	}
}
