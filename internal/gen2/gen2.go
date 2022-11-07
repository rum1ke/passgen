package gen2

import (
	"math/rand"
	"time"
)

const (
	symbolsNum int = 5
	numbersNum int = 10
	lettersNum int = 26
)

const (
	symbols int = iota + 1
	numbers
	highLetters
	lowLetters
)

type gen struct {
	symbols     [symbolsNum]byte
	numbers     [numbersNum]byte
	highLetters [lettersNum]byte
	lowLetters  [lettersNum]byte

	currentGroup int
	lastGroup    int
}

func NewGen() *gen {
	var g gen
	var pos, posEnd byte
	var i int

	g.symbols = [symbolsNum]byte{'@', '#', '$', '%', '&'}

	pos, posEnd = '0', '9'
	for i = 0; pos <= posEnd; i++ {
		g.numbers[i] = pos
		pos++
	}

	pos, posEnd = 'A', 'Z'
	for i = 0; pos <= posEnd; i++ {
		g.highLetters[i] = pos
		pos++
	}

	pos, posEnd = 'a', 'z'
	for i = 0; pos <= posEnd; i++ {
		g.lowLetters[i] = pos
		pos++
	}

	return &g
}

func (g *gen) GetNewPass(length int) string {
	pass := make([]byte, length, length)
	var char byte

	for i := 0; i < length; i++ {
		g.chooseNewGroup()

		if g.currentGroup == symbols {
			g.currentGroup = symbols
			char = g.getNewSymbol()
		} else if g.currentGroup == numbers {
			g.currentGroup = numbers
			char = g.getNewNumber()
		} else if g.currentGroup == highLetters {
			g.currentGroup = highLetters
			char = g.getNewHighLetters()
		} else if g.currentGroup == lowLetters {
			g.currentGroup = lowLetters
			char = g.getNewLowLetters()
		}

		if char == 0 {
			i--
			continue
		}

		pass[i] = char
		g.lastGroup = g.currentGroup
	}

	return string(pass)
}

func (g *gen) chooseNewGroup() {
	for {
		g.currentGroup = getRandomNumber(symbols, lowLetters)

		if g.currentGroup == g.lastGroup {
			continue
		}

		return
	}
}

func (g *gen) getNewSymbol() byte {
	index := getRandomNumber(0, symbolsNum-1)
	char := g.symbols[index]
	g.symbols[index] = 0
	return char
}

func (g *gen) getNewNumber() byte {
	index := getRandomNumber(0, numbersNum-1)
	char := g.numbers[index]
	g.numbers[index] = 0
	return char
}

func (g *gen) getNewHighLetters() byte {
	index := getRandomNumber(0, lettersNum-1)
	char := g.highLetters[index]
	g.highLetters[index] = 0
	return char
}

func (g *gen) getNewLowLetters() byte {
	index := getRandomNumber(0, lettersNum-1)
	char := g.lowLetters[index]
	g.lowLetters[index] = 0
	return char
}

func getRandomNumber(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max+1-min) + min
}
