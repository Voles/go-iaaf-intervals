package interval_test

import (
	interval "github.com/Voles/go-iaaf-intervals"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseInterval_Simple(t *testing.T) {
	actual, err := interval.Parse("1 x 300 (max)")

	expected := interval.Set{
		Repeats: 1,
		Repetitions: []interval.Repetition{
			{
				Repeats:  1,
				Distance: 300,
				Pace:     "max",
				Recovery: "",
			},
		},
		Recovery: "",
	}

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestParseInterval_Repetition_With_Recovery(t *testing.T) {
	actual, err := interval.Parse("10 x 400 (72”) [2’]")

	expected := interval.Set{
		Repeats: 1,
		Repetitions: []interval.Repetition{
			{
				Repeats:  10,
				Distance: 400,
				Pace:     "72”",
				Recovery: "2’",
			},
		},
		Recovery: "",
	}

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestParseInterval_Multiple_Repetitions(t *testing.T) {
	actual, err := interval.Parse("2 x 500 (300/48”, 200/max) [8’] [15’] 8 x 200 (35”) [1’]")

	expected := interval.Set{
		Repeats: 1,
		Repetitions: []interval.Repetition{
			{
				Repeats:  2,
				Distance: 500,
				Pace:     "300/48”, 200/max",
				Recovery: "8’",
			},
			{
				Repeats:  8,
				Distance: 200,
				Pace:     "35”",
				Recovery: "1’",
			},
		},
		Recovery: "15’",
	}

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestParseInterval_Grouped_Sets(t *testing.T) {
	actual, err := interval.Parse("2 x {1 x 500 (1500) [1’] 1 x 700 (1500) [30”] 1 x 300 (max)} [12’]")

	expected := interval.Set{
		Repeats: 2,
		Repetitions: []interval.Repetition{
			{
				Repeats:  1,
				Distance: 500,
				Pace:     "1500",
				Recovery: "1’",
			},
			{
				Repeats:  1,
				Distance: 700,
				Pace:     "1500",
				Recovery: "30”",
			},
			{
				Repeats:  1,
				Distance: 300,
				Pace:     "max",
				Recovery: "",
			},
		},
		Recovery: "12’",
	}

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestParseInterval_Sets(t *testing.T) {
	actual, err := interval.Parse("3 x 4 x 300 (3000) [100m r/o & 5’]")

	expected := interval.Set{
		Repeats: 3,
		Repetitions: []interval.Repetition{
			{
				Repeats:  4,
				Distance: 300,
				Pace:     "3000",
				Recovery: "100m r/o",
			},
		},
		Recovery: "5’",
	}

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}
