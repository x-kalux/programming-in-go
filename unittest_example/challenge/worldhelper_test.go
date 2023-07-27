package worldhelper

import (
	"regexp"
	"testing"
)

type pairString struct {
	arg      string
	expected string
}

var wordCensorTestCases = []pairString{
	{"Hi prayuth", "Hi *******"},
	{"I love Prayuth", "I love *******"},
	{"I love PraYuth", "I love *******"},
	{"I love prayutH", "I love *******"},
	{"I love PrAyuth", "I love *******"},
	{"Fuck up", "**** up"},
	{"FucK up", "**** up"},
	{"fUck up", "**** up"},
	{"fUCk up", "**** up"},
	{"Son of a bitch", "Son of a *****"},
	{"Son of a bitCh", "Son of a *****"},
	{"Son of a bItch", "Son of a *****"},
	{"Son of a BItch", "Son of a *****"},
	{"Piss off, Fuck you", "********, **** you"},
	{"Piss Off, Fuck you", "********, **** you"},
	{"Piss 0ff, Fuck you", "********, **** you"},
	{"PIss 0ff, Fuck you", "********, **** you"},
	{"P1ss Off, Fuck you", "********, **** you"},
	{"Plss Off, Fuck you", "********, **** you"},
}

func TestWordCensor(t *testing.T) {
	for _, testCase := range wordCensorTestCases {
		var want string = testCase.expected //Expected
		var got = WordCensor(testCase.arg)  //Actual
		if got != want {                    //Assert
			t.Errorf("\ninput: '%v', \n\twant: '%v', \n\tgot: '%v'\n", testCase.arg, want, got)
		}
	}
}

type pairIntInt struct {
	arg      int
	expected int
}

var buffaloNeckGeneratorTestCases = []pairIntInt{
	{-3, 3},
	{3, 3},
	{0, 1},
	{-30, 30},
	{30, 30},
	{-100, 100},
	{100, 100},
	{1000, 1000},
	{-798, 798},
	{1001, 1000},
	{2001, 1000},
	{3001, 1000},
	{4001, 1000},
	{5001, 1000},
	{6001, 1000},
	{7001, 1000},
	{8001, 1000},
	{-1001, 1000},
	{-2001, 1000},
	{-3001, 1000},
	{-4001, 1000},
	{-5001, 1000},
	{-6001, 1000},
	{-7001, 1000},
	{-8001, 1000},
}

func TestBuffaloNeckGenerator(t *testing.T) {
	var isBuffaloNeck = regexp.MustCompile(`^[ค]+$`).MatchString
	for _, testCase := range buffaloNeckGeneratorTestCases {
		var want int = testCase.expected                    //Expected
		var got = BuffaloNeckGenerator(testCase.arg)        //Actual
		if isBuffaloNeck(got) != true && want != len(got) { //Assert
			t.Errorf("\ninput: '%v', \n\twant: 'ค' %v characters, \n\tgot: 'ค' %v characters\n", testCase.arg, want, len(got))
		}
	}
}
