package main

import (
  "fmt"
)

func main() {
  var saldo int
  fmt.Scanf("%d", &saldo)

  numero_notas, troco := getNumeroNotas(saldo)

  fmt.Printf("Numero de notas: %d, Troco: %d\n", numero_notas, troco)
}
