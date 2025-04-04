package framework

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetAge(name string) (int, error) {
	type Age struct {
		Age int
	}
	age := Age{}

	client := http.Client{}
	res, err := client.Get("https://api.agify.io/?name=" + name)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	json.NewDecoder(res.Body).Decode(&age)
	return age.Age, err
}

func GetGender(name string) (string, error) {
	type Gender struct {
		Gender string //`json:"age"`
	}
	gender := Gender{}

	client := http.Client{}
	res, err := client.Get("https://api.genderize.io/?name=" + name)
	if err != nil {
		log.Println(err)
		return "", err
	}

	json.NewDecoder(res.Body).Decode(&gender)
	return gender.Gender, err
}

func GetRace(name string) (string, error) {
	type Country struct {
		Country_id string `json:"country_id"`
		//Probability float64 `json:"probability"`
	}
	type Race struct {
		Race []Country `json:"country"`
	}
	race := Race{}

	client := http.Client{}
	res, err := client.Get("https://api.nationalize.io/?name=" + name)
	if err != nil {
		log.Println(err)
		return "", err
	}

	body, err := io.ReadAll(res.Body)
	fmt.Print(json.Unmarshal(body, &race))
	//log.Println(json.NewDecoder(res.Body).Decode(&race))
	return race.Race[0].Country_id, err
}
