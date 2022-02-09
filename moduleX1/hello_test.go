package moduleX1

import "testing"

func TestHello(t *testing.T) {
    want := "moduleX1-Hello, world."
    if got := Hello(); got != want {
        t.Errorf("Hello() = %q, want %q", got, want)
    }
}