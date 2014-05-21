# Coding Dojo

Descubra o assassino [http://dojopuzzles.com](http://dojopuzzles.com/problemas/exibe/descubra-o-assassino/) feito em Golang

## Testes e cobertura
```bash
go test -coverprofile=c.out | go tool cover -func=c.out
```

## Exemplo de solução
```go
package main

import (
  "fmt"
  dojo "github.com/marcosinger/coding-dojo-detective"
)

func main() {
  game := dojo.NewGame()
  game.NewCrime()

  s := 0
  p := 0
  w := 0

  for {
    t := game.Theory(s, p, w)
    fmt.Println(game.ShowTheory(s, p, w))

    if t == 0 {
      fmt.Println(game.ShowCrime())
      break
    }

    if t == 1 {
      s++
    }

    if t == 2 {
      p++
    }

    if t == 3 {
      w++
    }
  }
}
```
