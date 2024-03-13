package tests

import (
	"testing"

	"github.com/untref-ayp2/busquedas"
)

func TestBusquedaLineal(t *testing.T) {
	arreglo := []int{1, 2, 3, 4, 5, 6, 7, 8}
	if ok, pos := busquedas.BusquedaLineal(arreglo, 1); !ok || pos != 0 {
		t.Error("El elemento 1 debería estar en la posición 0")
	}
	if ok, pos := busquedas.BusquedaLineal[int](arreglo, 2); !ok || pos != 1 {
		t.Error("El elemento 2 debería estar en la posición 1")
	}
	if ok, pos := busquedas.BusquedaLineal[int](arreglo, 3); !ok || pos != 2 {
		t.Error("El elemento 3 debería estar en la posición 2")
	}
	if ok, pos := busquedas.BusquedaLineal[int](arreglo, 4); !ok || pos != 3 {
		t.Error("El elemento 4 debería estar en la posición 3")
	}
	if ok, pos := busquedas.BusquedaLineal[int](arreglo, 5); !ok || pos != 4 {
		t.Error("El elemento 5 debería estar en la posición 4")
	}
	if ok, pos := busquedas.BusquedaLineal[int](arreglo, 6); !ok || pos != 5 {
		t.Error("El elemento 6 debería estar en la posición 5")
	}
	if ok, pos := busquedas.BusquedaLineal[int](arreglo, 7); !ok || pos != 6 {
		t.Error("El elemento 7 debería estar en la posición 6")
	}
	if ok, pos := busquedas.BusquedaLineal[int](arreglo, 8); !ok || pos != 7 {
		t.Error("El elemento 8 debería estar en la posición 7")
	}
	if ok, pos := busquedas.BusquedaLineal[int](arreglo, 9); ok || pos != -1 {
		t.Error("El elemento 9 no debería estar en el arreglo")
	}
}
