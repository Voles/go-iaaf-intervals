package interval

import (
	"errors"
	"regexp"
	"strconv"
)

type Repetition struct {
	Repeats  int
	Distance int
	Pace     string
	Recovery string
}

func (repetition Repetition) TotalDistance() int {
	return repetition.Repeats * repetition.Distance
}

type Set struct {
	Repeats     int
	Repetitions []Repetition
	Recovery    string
}

func (set Set) TotalDistance() int {
	var result int

	for i := 0; i < len(set.Repetitions); i++ {
		result += set.Repetitions[i].TotalDistance()
	}

	return set.Repeats * result
}

func Parse(notation string) (Set, error) {
	var res Set
	var err error

	res, err = loadGroupedSets(notation)
	if err != nil {
		res, err = loadSet(notation)
	}
	if err != nil {
		res, err = loadRepetitionSet(notation)
	}
	if err != nil {
		return res, err
	}

	return res, nil
}

func loadGroupedSets(notation string) (Set, error) {
	var result Set

	grouped_sets_regex := regexp.MustCompile(`(\d+)\s*x\s*\{(.*?)\}\s*\[(.*?)\]`)
	match := grouped_sets_regex.FindAllStringSubmatch(notation, -1)

	if len(match) < 1 {
		return result, errors.New("not a grouped set")
	}

	repeats, _ := strconv.Atoi(match[0][1])
	repetitions := match[0][2]
	recovery := match[0][3]

	result, _ = loadRepetitionSet(repetitions)
	result.Repeats = repeats
	result.Recovery = recovery

	return result, nil
}

func loadSet(notation string) (Set, error) {
	var result Set

	sets_regex := regexp.MustCompile(`(?:(\d+)\s*x\s*)?(\d+\s*x\s*.*)`)
	match := sets_regex.FindAllStringSubmatch(notation, -1)

	if len(match) < 1 {
		return result, errors.New("not a set")
	}

	var repeats int
	if len(match[0][1]) > 0 {
		repeats, _ = strconv.Atoi(match[0][1])
	} else {
		repeats = 1
	}
	repetitions := match[0][0]
	set, _ := loadRepetitionSet(repetitions)
	set.Repeats = repeats

	return set, nil
}

func loadRepetitionSet(notation string) (Set, error) {
	var result Set

	repetition_regex := regexp.MustCompile(`(\d+)\s*x\s*(\d+)\s*\((.*?)\)\s*(?:\[(.*?)(?:\s*&\s*(.*?))?\]\s*)?(?:\[(.*?)\])?`)
	match := repetition_regex.FindAllStringSubmatch(notation, -1)

	if len(match) < 1 {
		return result, errors.New("no set")
	}

	var recoveryBetweenSets string
	if len(match[0][6]) > 0 {
		recoveryBetweenSets = match[0][6]
	} else if len(match[0][5]) > 0 {
		recoveryBetweenSets = match[0][5]
	} else {
		recoveryBetweenSets = ""
	}

	return Set{1, loadRepetitionsFromMatch(match), recoveryBetweenSets}, nil
}

func loadRepetitionsFromMatch(match [][]string) []Repetition {
	var result []Repetition

	for i := 0; i < len(match); i++ {
		repeats, _ := strconv.Atoi(match[i][1])
		distance, _ := strconv.Atoi(match[i][2])
		pace := match[i][3]
		recovery := match[i][4]

		result = append(result, Repetition{
			Repeats:  repeats,
			Distance: distance,
			Pace:     pace,
			Recovery: recovery,
		})
	}

	return result
}
