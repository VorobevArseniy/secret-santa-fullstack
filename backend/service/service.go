package service

import (
	"math/rand"
	"strconv"
	"strings"
)

type service struct{}

func New() *service {
	return &service{}
}

type player struct {
	id int64
}

type players struct {
	players []player
}

func (s *service) CreateSeed() map[string]string {
	prs := &players{
		[]player{
			{id: 1},
			{id: 2},
			{id: 3},
			{id: 4},
			{id: 195},
			{id: 300},
			{id: 7},
			{id: 52},
			{id: 5},
			{id: 13},
			{id: 90},
			{id: 42},
		},
	}

	prs.shufflePlayers()
	m := make(map[string]string)
	size := len(prs.players)

	for len(m) != int(size) {
		var gifter, receiver int
		first := 1
		randNum := first

		for i := 0; i < size; i++ {
			gifter = randNum
			receiver = lsg(randNum, size)
			g := strconv.FormatInt(prs.players[gifter].id, 10)
			r := strconv.FormatInt(prs.players[receiver].id, 10)
			m[g] = r
			randNum = receiver
		}
	}

	return m
}

func euclid(a, b int) int {
	if b > a {
		return euclid(b, a)
	} else {
		if !(a%b == 0) {
			return euclid(b, a%b)
		} else {
			return b
		}
	}
}

func (slice *players) shufflePlayers() {
	for i := len(slice.players) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slice.players[i], slice.players[j] = slice.players[j], slice.players[i]
	}
}

func lsg(num, MaxNum int) int {
	gen_num := 2

	for works := euclid(MaxNum, gen_num) == 1; !(works); {
		gen_num += 1
		works = euclid(MaxNum, gen_num) == 1
	}

	return (num + gen_num) % MaxNum
}

func (s *service) EncodeSeed(m map[string]string) string {
	var seed string

	for k, v := range m {
		seed += k + "&" + v + "."
	}

	return seed[:len(seed)-1]
}

const (
	dot = "."
	and = "&"
)

func (s *service) DecodeSeed(seed string) map[string]string {
	m := make(map[string]string)

	pairs := strings.FieldsFunc(seed, func(r rune) bool { return r == rune(dot[0]) })

	for _, val := range pairs {
		// 71&31
		s := strings.FieldsFunc(val, func(r rune) bool { return r == rune(and[0]) })
		m[s[0]] = s[1]
	}

	return m
}
