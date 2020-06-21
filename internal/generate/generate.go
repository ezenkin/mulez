package generate

import (
	"fmt"
	"math/rand"
	"time"
)

type Data [][]string

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Table(count int, columns int) (t Data) {
	for i := 0; i < count/columns; i++ {
		r := generateRow(columns)
		t = append(t, r)
	}

	if rem := count % columns; rem > 0 {
		r := generateRow(rem)
		t = append(t, r)
	}
	return
}

func generateRow(count int) (res []string) {
	for j := 0; j < count; j++ {
		res = append(res, generateExercise())
	}
	return
}

func generateExercise() string {
	a, b := generateTwoNum(9)
	return fmt.Sprintf("%d x %d =", a, b)
}

func generateTwoNum(max int) (int, int) {
	return rand.Intn(max) + 1, rand.Intn(max) + 1
}
