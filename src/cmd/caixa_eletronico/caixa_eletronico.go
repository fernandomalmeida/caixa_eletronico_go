package main

// import "fmt"

var tipo_notas = [...]int {2, 5, 10, 20, 50, 100}

func getNumeroNotas(saldo int) (numero_notas int, troco int)  {
  numero_notas = 0
  for i := len(tipo_notas)-1; i >= 0; i-- {
    numero_notas += saldo/tipo_notas[i]
    saldo %= tipo_notas[i]
  }

  troco = saldo

  return
}
