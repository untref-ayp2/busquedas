package busquedas

// BusquedaBinaria busca un elemento en un arreglo de elementos comparables.
// Si lo encuentra devuelve true y la posición en la que se encuentra.
// Si no lo encuentra devuelve false y -1.
// Se necesita que el arreglo esté ordenado, y por lo tanto que los elementos del arreglo soporten orden total
// (es decir, que sean comparables por igual, menor y mayor).
// Se debe pasar una función de comparación como parámetro.
// Complejidad: O(log n)
func BusquedaBinaria[T comparable](arreglo []T, elemento T, comp func(T, T) int) (bool, int) {
	izq := 0
	der := len(arreglo) - 1
	for izq <= der {
		medio := (izq + der) / 2
		if arreglo[medio] == elemento {
			return true, medio
		} else if comp(arreglo[medio], elemento) < 0 {
			der = medio - 1
		} else {
			izq = medio + 1
		}
	}
	return false, -1
}
