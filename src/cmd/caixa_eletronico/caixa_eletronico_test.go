package main

import "testing"

func TestSaldoZero(t *testing.T) {
  numero_notas, _ := getNumeroNotas(0)

  if numero_notas != 0 {
    t.Error("Numero de notas diferente de 0,", numero_notas, "obtido.")
  }
}

func TestSaldoDois(t *testing.T) {
  numero_notas, _ := getNumeroNotas(2)

  if numero_notas != 1 {
    t.Error("Numero de notas diferente de 1,", numero_notas, "obtido.")
  }
}

func TestSaldoQuatro(t *testing.T) {
  numero_notas, _ := getNumeroNotas(4)

  if numero_notas != 2 {
    t.Error("Numero de notas diferente de 2,", numero_notas, "obtido.")
  }
}

func TestSaldoSeis(t *testing.T) {
  numero_notas, troco := getNumeroNotas(6)

  if numero_notas != 1 {
    t.Error("Numero de notas diferente de 1,", numero_notas, "obtido.")
  }
  if troco != 1 {
    t.Error("Troco diferente de 1,", troco, "obtido.")
  }
}

func TestSaldoOitentaESete(t *testing.T) {
  numero_notas, troco := getNumeroNotas(87)

  if numero_notas != 5 {
    t.Error("Numero de notas diferente de 5,", numero_notas, "obtido.")
  }
  if troco != 0 {
    t.Error("Troco diferente de 0,", troco, "obtido.")
  }
}
