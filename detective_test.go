package detective

import (
  "reflect"
  "testing"
)

func TestSuspects(t *testing.T) {
  game := NewGame()
  expected := []string{
    "Charles B. Abbage",
    "Donald Duck Knuth",
    "Ada L. Ovelace",
    "Alan T. Uring",
    "Ivar J. Acobson",
    "Ras Mus Ler Dorf"}

  if !reflect.DeepEqual(game.Suspects, expected) {
    t.Errorf("expected %s got %s", expected, game.Suspects)
  }
}

func TestPlaces(t *testing.T) {
  game := NewGame()
  expected := []string{
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

  if !reflect.DeepEqual(game.Places, expected) {
    t.Errorf("expected %s got %s", expected, game.Places)
  }
}

func TestWeapons(t *testing.T) {
  game := NewGame()
  expected := []string{
    "Peixeira",
    "DynaTAC 8000X",
    "Trezoitão",
    "Trebuchet",
    "Maça",
    "Gládio"}

  if !reflect.DeepEqual(game.Weapons, expected) {
    t.Errorf("expected %s got %s", expected, game.Weapons)
  }
}

func TestNewCrime(t *testing.T) {
  game := NewGame()
  game.NewCrime()

  if len(game.Crime) == 0 {
    t.Error("expected a crime, got nothing")
  }
}

func TestWrongTheory(t *testing.T) {
  game := new(Game)
  game.Crime = []int{1, 3, 2} // Donald Duck Knuth usando um(a) trezoitão em Tokio

  theory := game.Theory(1, 1, 1)

  if !(theory == 1 || theory == 2 || theory == 3) {
    t.Errorf("expected in %s, got %s", []int{1, 2, 3}, theory)
  }
}

func TestTheoryWithRightPlace(t *testing.T) {
  game := new(Game)
  game.Crime = []int{1, 3, 2} // Donald Duck Knuth usando um(a) trezoitão em Tokio

  theory := game.Theory(3, 3, 3)

  if !(theory == 1 || theory == 3) {
    t.Errorf("expected in %s, got %s", []int{1, 3}, theory)
  }
}

func TestTheoryWithRightPlaceAndRightWeapon(t *testing.T) {
  game := new(Game)
  game.Crime = []int{1, 3, 2} // Donald Duck Knuth usando um(a) trezoitão em Tokio

  theory := game.Theory(5, 3, 2)

  if theory != 1 {
    t.Errorf("expected in %s, got %s", []int{1}, theory)
  }
}

func TestRightTheory(t *testing.T) {
  game := new(Game)
  game.Crime = []int{1, 3, 2} // Donald Duck Knuth usando um(a) trezoitão em Tokio

  theory := game.Theory(1, 3, 2)

  if theory != 0 {
    t.Errorf("expected %s, got %s", 0, theory)
  }
}

func TestShowTheory(t *testing.T) {
  game := NewGame()

  expected := "Teoria 0 -> Charles B. Abbage usando um(a) maça em San Francisco"
  message := game.ShowTheory(0, 2, 4)

  if message != expected {
    t.Errorf("expected %s, got %s", expected, message)
  }
}

func TestShowCrime(t *testing.T) {
  game := NewGame()
  game.Crime = []int{1, 3, 2}

  expected := "Crime solucionado -> Donald Duck Knuth usando um(a) trezoitão em Tokio"
  message := game.ShowCrime()

  if message != expected {
    t.Errorf("expected %s, got %s", expected, message)
  }
}
