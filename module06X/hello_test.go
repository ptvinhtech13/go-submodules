package module06X

import "testing"

func TestHello(t *testing.T) {
    want := "module 2-Hello, world."
    if got := Hello(); got != want {
        t.Errorf("Hello() = %q, want %q", got, want)
    }
}