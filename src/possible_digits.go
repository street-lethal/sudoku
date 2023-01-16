package src

import (
	"fmt"
	"strconv"
	"strings"
)

type PossibleDigits map[int]bool

func (d *PossibleDigits) Multiple(other *PossibleDigits) *PossibleDigits {
	product := PossibleDigits{}
	for i := 1; i <= 9; i++ {
		if (*d)[i] && (*other)[i] {
			product[i] = true
		}
	}

	return &product
}

func (d *PossibleDigits) Add(i int) {
	(*d)[i] = true
}

func (d *PossibleDigits) Remove(i int) {
	(*d)[i] = false
}

func (d *PossibleDigits) String() string {
	array := []string{}
	for digit, possible := range *d {
		if possible {
			array = append(array, strconv.Itoa(digit))
		}
	}

	return fmt.Sprintf("[%s]", strings.Join(array, ","))
}
