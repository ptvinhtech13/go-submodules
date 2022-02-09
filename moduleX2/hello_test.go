package moduleX2

import "testing"

func TestHello(t *testing.T) {
    want := "moduleX2-Hello, world."
    if got := Hello(); got != want {
        t.Errorf("Hello() = %q, want %q", got, want)
    }
}