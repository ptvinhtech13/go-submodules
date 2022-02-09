package moduleX2

import "rsc.io/quote"

func Hello() string {
    return "moduleX2-" + quote.Hello()
}