package main

import (
	"errors"
	"fmt"
	"strconv"
)

type rationals struct {
	sequence []int64
}

func (this rationals) rational(i int) string {
	return strconv.FormatInt(this.sequence[i], 10)
}

func add1(i int64, n int64, seq []int64) ([]int64, error) {
	if i < 0 {
		return nil, errors.New("i is less than 0")
	}
	if i > n {
		return seq, nil
	}
	seq = append(seq, i)
	i++
	return add1(i, n, seq)
}

func mltp3(i int64, n int64, seq []int64) ([]int64, error) {
	if i < 0 {
		return nil, errors.New("i is less than 0")
	}
	if i > n {
		return seq, nil
	}
	seq = append(seq, i)
	i *= 3
	return mltp3(i, n, seq)
}

func plusi(i int64, n int64, seq []int64) ([]int64, error) {
	if i < 0 {
		return nil, errors.New("i is less than 0")
	}
	if i > n {
		return seq, nil
	}
	seq = append(seq, i)
	i += i
	return plusi(i, n, seq)
}

func New(
	i int64,
	n int64,
	seq []int64,
	op func(
		int64,
		int64,
		[]int64,
	) ([]int64, error),
) rationals {
	sequence, err := op(i, n, seq)
	if err != nil {
		fmt.Println(err)
		panic(1)
	}
	return rationals{sequence}
}

func main() {
	i := int64(1)
	n := int64(1000)
	var sequence rationals

	sequence = New(i, n, make([]int64, 0), add1)
	fmt.Println(sequence.rational(5))
	sequence = New(i, n, make([]int64, 0), mltp3)
	fmt.Println(sequence.rational(5))
	sequence = New(i, n, make([]int64, 0), plusi)
	fmt.Println(sequence.rational(5))
}
