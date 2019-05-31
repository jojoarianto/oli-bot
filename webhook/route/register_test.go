package route

import (
	"strings"
	"testing"
)

func TestRegister(t *testing.T) {

	t.Run("Jika parameter tidak ada", func(t *testing.T) {

		input := "/register"
		inputArr := strings.Fields(input)
		expected := "Parameter yang anda inputkan salah"

		if got := Register(inputArr); got != expected {
			t.Errorf("Register() = %v, want %v", got, expected)
		}
	
	})

	t.Run("Jika parameter ada password salah", func(t *testing.T) {

		input := "/register panitia_2019@venolfkhunair.com wrongpassword id_line group"
		inputArr := strings.Fields(input)
		expected := "login gagal"

		if got := Register(inputArr); got != expected {
			t.Errorf("Register() = %v, want %v", got, expected)
		}
	
	})

}
