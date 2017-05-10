# Juego del Manotazo

¿Cúal es la probabilidad de tirar todas las cartas en el juego del manotazo sin que el numero que se pronuncia corresponda alguna de las cartas que se van lanzando desde un mazo mezclado?.

Este programa responde a esa pregunta con una simulación usando metodos de montecarlo. Es decir prueba varias barajas aleatorias de n x m elementos donde n es el numero de cartas diferentes (símbolos) y m es el número de cartas de un símbolo en particular. se comienza la cuenta desde un símbolo hasta el último volviendo a comenzar en ese orden stablecido. Cuando el simbolo de la cuenta corresponde al que sale en la baraja se termina esa simulacion en particular y se devuelve el número de cartas que pasaron antes de este evento.


Para ejecutar este programa:

```
  go run main.go [opciones]

```

para ver las opciones ejecutar primero son ningun argumento.

