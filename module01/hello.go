package hello

import "rsc.io/quote"

func Hello() string {
    return "module 1-" + quote.Hello()
}