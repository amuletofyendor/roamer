package dice

import(
  "math/rand"
)

const (
  Heads int = 1
  Tails int = 2
)

type Dice struct {
  Rolls int
  Value int
  Modifier int
}

func (d Dice) Roll() int {
  result := 0

  for i := 0; i < d.Rolls; i++ {
    result += rand.Intn(d.Value) + 1
  }

  return result + d.Modifier
}

func CoinToss() int {
  return Dice{1, 2, 0}.Roll()
}

func Oned4() int {
  return Dice{1, 4, 0}.Roll()
}

func Twod4() int {
  return Dice{2, 4, 0}.Roll()
}

func Threed4() int {
  return Dice{3, 4, 0}.Roll()
}

func Fourd4() int {
  return Dice{4, 4, 0}.Roll()
}

func Oned6() int {
  return Dice{1, 6, 0}.Roll()
}

func Oned8() int {
  return Dice{1, 8, 0}.Roll()
}

func Oned10() int {
  return Dice{1, 8, 0}.Roll()
}

func Oned20() int {
  return Dice{1, 20, 0}.Roll()
}
