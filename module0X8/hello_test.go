package module0X8

import "testing"

func TestHello(t *testing.T) {
    want := "module 1-Hello, world."
    if got := Hello(); got != want {
        t.Errorf("Hello() = %q, want %q", got, want)
    }
}