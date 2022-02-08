package module02

import "rsc.io/quote"

func Hello() string {
    return "module 2-" + quote.Hello()
}