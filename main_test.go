package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrSum(t *testing.T) {

	testTable := []struct {
		numbers  []int
		expected int
	}{
		{
			numbers:  []int{4, 2, 4, 6, 12},
			expected: 28,
		},
		{
			numbers:  []int{4, 12, -4, -6, 12},
			expected: 18,
		},
		{
			numbers:  []int{},
			expected: 0,
		},
	}

	for _, testCase := range testTable {
		result := ArrSum(testCase.numbers)
		assert.Equal(t, testCase.expected, result)
	}

}

func TestMax(t *testing.T) {
	testTable := []struct {
		numbers  []int
		expected int
	}{
		{
			numbers:  []int{1, 2, 3, 6},
			expected: 6,
		},
		{
			numbers:  []int{4, 12, -4, -6, 12},
			expected: 12,
		},
		{
			numbers:  []int{0},
			expected: 0,
		},
	}

	for _, testCase := range testTable {
		result := Max(testCase.numbers)
		assert.Equal(t, testCase.expected, result)
	}
}

func TestMin(t *testing.T) {
	testTable := []struct {
		numbers  []int
		expected int
	}{
		{
			numbers:  []int{1, 2, 3, 6},
			expected: 1,
		},
		{
			numbers:  []int{4, 12, -4, -6, 12},
			expected: -6,
		},
		{
			numbers:  []int{0},
			expected: 0,
		},
	}

	for _, testCase := range testTable {
		result := Min(testCase.numbers)
		assert.Equal(t, testCase.expected, result)
	}
}

func TestMaxNumSum(t *testing.T) {
	testTable := []struct {
		numbers  []int
		expected int
	}{
		{
			numbers:  []int{1, 2, 3, 6},
			expected: 11,
		},
		{
			numbers:  []int{4, 12, -4, -6, 12},
			expected: 24,
		},
		{
			numbers:  []int{0},
			expected: 0,
		},
	}

	for _, testCase := range testTable {
		result := MaxNumSum(testCase.numbers)
		assert.Equal(t, testCase.expected, result)
	}
}
func TestMinNumSum(t *testing.T) {
	testTable := []struct {
		numbers  []int
		expected int
	}{
		{
			numbers:  []int{1, 2, 3, 6},
			expected: 6,
		},
		{
			numbers:  []int{4, 12, -4, -6, 12},
			expected: 6,
		},
		{
			numbers:  []int{0},
			expected: 0,
		},
	}

	for _, testCase := range testTable {
		result := MinNumSum(testCase.numbers)
		assert.Equal(t, testCase.expected, result)
	}
}
