package storage

import (
	"testing"
)

func TestListStorage(t *testing.T) {

	listStorage := NewListStorage()

	listStorage.LeftPush("mylist", "item1")
	listStorage.LeftPush("mylist", "item2")
	listStorage.RightPush("mylist", "item3")
	listStorage.RightPush("mylist", "item4")

	expectedList := []string{"item1", "item2", "item3", "item4"}

	actualList, err := listStorage.ListRange("mylist", 0, 3)

	if err != nil {
		t.Fatalf("Error inesperado en LRange: %v", err)
	}

	if !equalSlices(expectedList, actualList) {
		t.Errorf("Estado incorrecto de la lista: esperado %v, obtenido %v", expectedList, actualList)
	}

	item, err := listStorage.ListIndex("mylist", 2)
	if err != nil {
		t.Fatalf("Error inesperado en LIndex: %v", err)
	}
	if item != "item3" {
		t.Errorf("Elemento incorrecto en el índice 2: esperado %v, obtenido %v", "item3", item)
	}

	expectedRange := []string{"item2", "item3", "item4"}

	actualRange, err := listStorage.ListRange("mylist", 1, 3)

	if err != nil {
		t.Fatalf("Error inesperado en LRange: %v", err)
	}

	if !equalSlices(expectedRange, actualRange) {
		t.Errorf("Rango incorrecto: esperado %v, obtenido %v", expectedRange, actualRange)
	}

	_, err = listStorage.ListIndex("mylist", 10)

	if err == nil {
		t.Errorf("Se esperaba un error para índice fuera de rango, pero no se obtuvo")
	}

	_, err = listStorage.ListRange("mylist", -1, 10)
	if err == nil {
		t.Errorf("Se esperaba un error para rango inválido, pero no se obtuvo")
	}
}

// Función auxiliar para comparar slices
func equalSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
