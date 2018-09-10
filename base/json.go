package base

import "fmt"
import "encoding/json"

type Movie struct {
	Title string
	Year int `json:"released"`
	Color bool `json:"color,omitempty"`
	Actors []string `json:"actors"`
}

func Json() {

	var movie = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: true,
			Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	}
	data,err:=json.MarshalIndent(movie,"","   ")
	if err!=nil{
		fmt.Println("err",err)
	}
	fmt.Printf("%s\n",data)
}