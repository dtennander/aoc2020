package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewLine(t *testing.T) {
	// Given
	text := "11-14 c: hccccccccccscp"
	// When
	line, err := NewLine(text)
	// Then
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, Line{
		letter:   'c',
		min:      11,
		max:      14,
		password: "hccccccccccscp",
	}, line, "Parsed line does not match")
}

func TestLine_PositionRuleValid(t *testing.T) {
	type fields struct {
		letter   rune
		min      int64
		max      int64
		password string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "",
			fields: fields{
				letter:   'c',
				min:      11,
				max:      14,
				password: "hccccccccccscp",
			},
			want: true,
		}, {
			name: "",
			fields: fields{
				letter:   'h',
				min:      3,
				max:      13,
				password: "zzhnhnjhhkplhhwph",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Line{
				letter:   tt.fields.letter,
				min:      tt.fields.min,
				max:      tt.fields.max,
				password: tt.fields.password,
			}
			if got := l.FollowsPositionRule(); got != tt.want {
				t.Errorf("FollowsPositionRule() = %v, want %v", got, tt.want)
			}
		})
	}
}
