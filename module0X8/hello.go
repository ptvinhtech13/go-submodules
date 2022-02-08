package module01

import "rsc.io/quote"

func Hello() string {
    return "module 21-" + quote.Hello()
}