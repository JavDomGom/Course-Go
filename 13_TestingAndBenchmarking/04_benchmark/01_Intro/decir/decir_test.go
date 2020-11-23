package decir

import (
	"fmt"
	"testing"
)

func TestSaludar(t *testing.T) {
	s := Saludar("Javier")
	if s != "Bienvenido Javier" {
		t.Error("Expected", "Bienvenido Javier", ", got", s)
	}
}

func ExampleSaludar() {
	fmt.Println(Saludar("Javier"))
	// Output: Bienvenido Javier
}

func BenchmarkSaludar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Saludar("Javier")
	}
}
