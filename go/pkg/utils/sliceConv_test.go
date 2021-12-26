package utils

import (
	"reflect"
	"testing"
)


var	intSliceEmpty = []int{}
var	intSlice123 = []int{1,2,3}

var	stringSliceEmpty = []string{}
var	stringSlice123 = []string{"1", "2", "3"}


func TestNewIntSliceFromStringSliceEmpty(t *testing.T) {
	expected := intSliceEmpty
	result, err := NewIntSliceFromStringSlice(stringSliceEmpty)
	if err != nil {
		t.Errorf("%v",err)
	}

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("\tgot:\t%v\n\t\twanted:\t%v\n", result, expected)
	}
}

func TestNewIntSliceFromStringSlice123(t *testing.T) {
	expected := intSlice123
	result, err := NewIntSliceFromStringSlice(stringSlice123)
	if err != nil {
		t.Errorf("%v",err)
	}

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("\tgot:\t%v\n\t\twanted:\t%v\n", result, expected)
	}
}