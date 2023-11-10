package main

import (
	"reflect"
	"testing"
)

func TestGroupBreedsByCountry(t *testing.T) {
	catBreeds := []CatBreed{
		{Breed: "Breed1", Country: "Country1", Origin: "Origin1", Coat: "Coat1", Pattern: "Pattern1"},
		{Breed: "Breed2", Country: "Country1", Origin: "Origin2", Coat: "Coat2", Pattern: "Pattern2"},
	}

	expectedResult := map[string][]map[string]string{
		"Country1": {{
			"breed":   "Breed1",
			"origin":  "Origin1",
			"coat":    "Coat1",
			"pattern": "Pattern1",
		}, {
			"breed":   "Breed2",
			"origin":  "Origin2",
			"coat":    "Coat2",
			"pattern": "Pattern2",
		}},
	}

	result := groupBreedsByCountry(catBreeds)

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Expected %v, but got %v", expectedResult, result)
	}
}

func TestSortBreedsByLength(t *testing.T) {
	groupedBreeds := map[string][]map[string]string{
		"Country1": {{
			"breed":   "LongBreedName",
			"origin":  "Origin1",
			"coat":    "Coat1",
			"pattern": "Pattern1",
		}, {
			"breed":   "Short",
			"origin":  "Origin2",
			"coat":    "Coat2",
			"pattern": "Pattern2",
		}},
	}

	expectedResult := map[string][]map[string]string{
		"Country1": {{
			"breed":   "Short",
			"origin":  "Origin2",
			"coat":    "Coat2",
			"pattern": "Pattern2",
		}, {
			"breed":   "LongBreedName",
			"origin":  "Origin1",
			"coat":    "Coat1",
			"pattern": "Pattern1",
		}},
	}

	sortBreedsByLength(groupedBreeds)

	if !reflect.DeepEqual(groupedBreeds, expectedResult) {
		t.Errorf("Expected %v, but got %v", expectedResult, groupedBreeds)
	}
}
