package ex2

import "testing"

type addTest struct {
	inp, ans uint32
	valid    bool
}

var addTests = []addTest{
	{0, 0, true},
	{1, 1, true},
	{2, 1, true},
	{3, 2, true},
	{4, 3, true},
	{5, 5, true},
	{6, 8, true},
	{7, 13, true},
	{8, 21, true},
	{9, 34, true},
	{10, 55, true},
	{101, 0, false},
	{200, 0, false},
}

func TestExercise(t *testing.T) {
	for _, test := range addTests {
		got, ok := Exercise(test.inp)
		if ok != nil && test.valid == true {
			t.Errorf("Incorrect validation for valid input %d", test.inp)
		} else if ok == nil && test.valid == false {
			t.Errorf("Incorrectly acepted invalid input %d", test.inp)
		} else if ok == nil && got != test.ans {
			t.Errorf("Incorrect calculation for input %d (expected %d, got %d)", test.inp, test.ans, got)
		}
	}

}
