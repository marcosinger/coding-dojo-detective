package detective

import (
  "fmt"
  "math/rand"
  "strings"
  "time"
)

type Game struct {
  Suspects []string
  Places   []string
  Weapons  []string
  Crime    []int
  Tries    int
}

var (
  Weapons = []string{
    "Peixeira",
    "DynaTAC 8000X",
    "Trezoitão",
    "Trebuchet",
    "Maça",
    "Gládio"}

  Places = []string{
    "Redmond",
    "Palo Alto",
    "San Francisco",
    "Tokio",
    "Restaurante no Fim do Universo",
    "São Paulo",
    "Cupertino",
    "Helsinki",
    "Maida Vale",
    "Toronto"}

  Suspects = []string{
    "Charles B. Abbage",
    "Donald Duck Knuth",
    "Ada L. Ovelace",
    "Alan T. Uring",
    "Ivar J. Acobson",
    "Ras Mus Ler Dorf"}
)

func NewGame() *Game {
  return &Game{
    Suspects: Suspects,
    Places:   Places,
    Weapons:  Weapons}
}

func (g *Game) NewCrime() {
  rand.Seed(time.Now().UTC().UnixNano())
  g.Crime = []int{
    rand.Intn(len(g.Suspects)),
    rand.Intn(len(g.Places)),
    rand.Intn(len(g.Weapons))}
}

func (g *Game) Theory(suspect, place, weapon int) int {
  possibilities := make([]int, 0)

  // incrementa as tentativas
  g.Tries++

  // Suspeito errado
  if suspect != g.Crime[0] {
    possibilities = append(possibilities, 1)
  }

  // Lugar errado
  if place != g.Crime[1] {
    possibilities = append(possibilities, 2)
  }

  // Arma errada
  if weapon != g.Crime[2] {
    possibilities = append(possibilities, 3)
  }

  if len(possibilities) > 0 {
    // devolve um dos erros aleatóriamente
    rand.Seed(time.Now().UTC().UnixNano())
    return possibilities[rand.Intn(len(possibilities))]
  } else {
    // resposta correta
    return 0
  }
}

func (g *Game) ShowTheory(suspect, place, weapon int) string {
  return fmt.Sprintf("Teoria %d -> %s usando um(a) %s em %s",
    g.Tries, g.Suspects[suspect], strings.ToLower(g.Weapons[weapon]), g.Places[place])
}

func (g *Game) ShowCrime() string {
  return fmt.Sprintf("Crime solucionado -> %s usando um(a) %s em %s",
    g.Suspects[g.Crime[0]], strings.ToLower(g.Weapons[g.Crime[2]]), g.Places[g.Crime[1]])
}
