package slugify

import (
	"testing"
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
