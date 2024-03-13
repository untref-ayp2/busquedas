package busquedas

// BusquedaLineal busca un elemento en un arreglo de enteros.
// Si lo encuentra devuelve true y la posición en la que se encuentra.
// Si no lo encuentra devuelve false y -1.
// Los elementos del arreglo deben ser comparables por igual solamente.
// Complejidad: O(n)
func BusquedaLineal[T comparable](arreglo []T, elemento T) (bool, int) {
	for i, v := range arreglo {
		if v == elemento {
			return true, i
		}
	}
	return false, -1
}

