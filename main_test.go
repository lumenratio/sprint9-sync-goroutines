package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Пишите тесты в этом файле
func Test_generateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
		want int
	}{
		{name: "Zero size", size: 0, want: 0},
		{name: "Negative size", size: -1, want: 0},
		{name: "Positive size: 1", size: 1, want: 1},
		{name: "Positive size: big num", size: 10_000, want: 10_000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.size)
			assert.Len(t, got, tt.want)
		})
	}
}

func Test_maximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "Zero len of slice", data: []int{}, want: 0},
		{name: "Send nil slice", data: nil, want: 0},
		{name: "One element in slice: zero", data: []int{0}, want: 0},
		{name: "One element in slice: positive number", data: []int{1}, want: 1},
		{name: "Small slice: positive numbers", data: []int{2, 5, 3, 1, 10, 8}, want: 10},
		{name: "Small slice: mixed numbers", data: []int{2, 5, -4, 3, -6, 1, 10, -5, 8}, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximum(tt.data)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_maxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "Zero len of slice", data: []int{}, want: 0},
		{name: "Send nil slice", data: nil, want: 0},
		{name: "One element in slice: zero", data: []int{0}, want: 0},
		{name: "One element in slice: positive number", data: []int{1}, want: 1},
		{name: "Small slice: positive numbers", data: []int{2, 5, 3, 1, 10, 8}, want: 10},
		{name: "Big slice: mixed numbers, 33 values", data: []int{188, 99, 25, 49, 144, 89, 75, 45, 75, 35, 133, 190,
			22, 156, 99, 185, 105, 183, 145, 105, 112, 180, 69, 54,
			104, 87, 155, 94, 45, 89, 142, 146, 133},
			want: 190},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxChunks(tt.data)
			assert.Equal(t, tt.want, got)
		})
	}
}
