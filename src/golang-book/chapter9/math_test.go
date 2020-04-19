package math_test

import (
	math "./" // import golang-book/chapter9/math package
	"testing"
)

type testpair struct {
	values  []float64
	average float64
}

var tests = []testpair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
}

/*
* Tests are identified by starting a function with the word Test and taking
* one argument of type *testing.T.
 */
func TestAverage(t *testing.T) {
	for _, pair := range tests {
		v := math.Average(pair.values)
		if v != pair.average {
			t.Error(
				"For", pair.values,
				"expected", pair.average,
				"got", v,
			)
		}
	}

}
