package one_away

import (
	"testing"

	"github.com/subfuzion/coding-problems-in-go/test"
)

// F is the type of function that implements the solution
type F func(string, string) bool

// Test defines a test case
type Test struct {
	A        string
	B        string
	Category string
	Expected bool
}

// Tests is a test suite
var Tests = []Test{
	{
		"abc",
		"abd",
		"one replacement",
		true,
	},
	{
		"abc",
		"abb",
		"one replacement",
		true,
	},
	{
		"bbc",
		"abc",
		"one replacement",
		true,
	},
	{
		"abc",
		"abcd",
		"one insert",
		true,
	},
	{
		"ac",
		"abc",
		"one insert",
		true,
	},
	{
		"bc",
		"abc",
		"one insert",
		true,
	},
	{
		"abc",
		"bc",
		"one insert",
		true,
	},
	{
		"abc",
		"ac",
		"one insert",
		true,
	},
	{
		"abc",
		"ab",
		"one insert",
		true,
	},
	{
		"ab",
		"abcd",
		"more than one away insert",
		false,
	},
	{
		"abcde",
		"ab",
		"more than one away insert",
		false,
	},
	{
		"abcd",
		"abbb",
		"more than one away replace",
		false,
	},
	{
		"abc",
		"xyz",
		"more than one away replace",
		false,
	},
	{
		"abc",
		"abc",
		"equal strings",
		true,
	},
}

func TestSolution(t *testing.T) {
	for _, f := range Solutions {
		t.Run(test.GetFileFuncName(f), func(t *testing.T) {
			for _, test := range Tests {
				if ok := isOneEditAway(test.A, test.B); ok != test.Expected {
					t.Errorf("%s: '%s', '%s'", test.Category, test.A, test.B)
				}
			}
		})
	}
}

// TODO - add your own solutions to the list below
// Solutions maps descriptive names to solution implementations
var Solutions = []F{
	isOneEditAway,
	// isOneEditAway1,
	// ...
}

//TODO - implement (and update the Solutions list, above)
//nolint:deadcode,unused //golangci-lint
/******************************************************************************
One Away: There are three types of edits that can be performed on strings:
insert a character, remove a character, or replace a character.
Given two strings, write a function to check if they are one edit (or zero
edits) away.
Hints (Cracking the Coding Interview 6): #23, #97, #130
******************************************************************************/
func isOneEditAway1(a, b string) bool {
	return false
}
