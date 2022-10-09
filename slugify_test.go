package slugify

import (
	"testing"
	"strings"
)

func TestConvertSpacesIntoHyphens(t *testing.T) {
	s := Slugify("all lower case")
	if s != "all-lower-case" {
		t.Errorf("all spaces must be converted to hyphen")
	}
}

func TestConvertSingleQuotasIntoHyphens(t *testing.T) {
	s := Slugify("sensorario's branch")
	if s != "sensorario-s-branch" {
		t.Errorf("all single quotas must be converted to hyphen")
	}
}

func TestRemovNewLines(t *testing.T) {
	s := Slugify("sensorario's branch\n")
	if s != "sensorario-s-branch" {
		t.Errorf("all new lines should be removed")
	}
}

func TestLowercaseSentence(t *testing.T) {
	s := Slugify("My name is sensorario")
	if s != "my-name-is-sensorario" {
		t.Errorf(strings.Join([]string{"Slug  should be 'my-name-is-sensorario' instead of", s}, " "))
	}
}

func TestTrimSlug(t *testing.T) {
	s := Slugify("My name is sensorario     ")
	if s != "my-name-is-sensorario" {
		t.Errorf(strings.Join([]string{"Slug  should be 'my-name-is-sensorario' instead of", s}, " "))
	}
}
