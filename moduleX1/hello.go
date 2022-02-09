package moduleX1

import "rsc.io/quote"

func Hello() string {
    return "moduleX1-" + quote.Hello()
}