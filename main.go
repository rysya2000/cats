package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
)

const URL = "https://catfact.ninja/breeds"

type CatBreed struct {
	Breed   string `json:"breed"`
	Country string `json:"country"`
	Origin  string `json:"origin"`
	Coat    string `json:"coat"`
	Pattern string `json:"pattern"`
}

func main() {
	catBreeds, err := fetchCatBreeds()
	if err != nil {
		fmt.Printf("fetchCatBreeds err: %v\n", err)
		return
	}

	groupedBreeds := groupBreedsByCountry(catBreeds)
	sortBreedsByLength(groupedBreeds)

	resultJSON, err := json.MarshalIndent(groupedBreeds, "", "    ")
	if err != nil {
		fmt.Printf("marshaling err: %v\n", err)
		return
	}

	err = os.WriteFile("out.json", resultJSON, 0644)
	if err != nil {
		fmt.Printf("writeFile err: %v\n", err)
		return
	}

	fmt.Println("The result was successfully written to out.json")
}

func fetchCatBreeds() ([]CatBreed, error) {
	response, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var breedsResponse struct {
		Data []CatBreed `json:"data"`
	}

	err = json.Unmarshal(body, &breedsResponse)
	if err != nil {
		return nil, err
	}

	return breedsResponse.Data, nil
}

func groupBreedsByCountry(breeds []CatBreed) map[string][]map[string]string {
	groupedBreeds := make(map[string][]map[string]string)

	for _, breed := range breeds {
		breedDetails := map[string]string{
			"breed":   breed.Breed,
			"origin":  breed.Origin,
			"coat":    breed.Coat,
			"pattern": breed.Pattern,
		}
		groupedBreeds[breed.Country] = append(groupedBreeds[breed.Country], breedDetails)
	}

	return groupedBreeds
}

func sortBreedsByLength(groupedBreeds map[string][]map[string]string) {
	for _, breeds := range groupedBreeds {
		sort.Slice(breeds, func(i, j int) bool {
			return len(breeds[i]["breed"]) < len(breeds[j]["breed"])
		})
	}
}
