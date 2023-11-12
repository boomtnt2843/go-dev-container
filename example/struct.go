package example

import (
	"encoding/json"
	"fmt"
	"log"
)

type Address struct {
	Address 	string
	Postcode 	string
}

type UserProfile struct {
	Firstname 	string 		`json:"firstname"`
	Lastname 	string		`json:"lastname"`
	Age 		int
	Height		float32
	Address 	Address
	Bill		struct{
		BillAddress	string
	}
}

func (u UserProfile) ToFullDecs() string {
	return fmt.Sprintf("%s %s", u.Firstname, u.Lastname)
}

func main() {
	user := map[string]string{}

	user["username"] = "boomTnT"
	user["password"] = "xxxxxx"

	fmt.Println(user)
	fmt.Println(user["username"])

	userProfile := UserProfile{
		Firstname: 	"boomphat",
		Lastname: 	"pimpim",
		Age: 		21,
		Address: 	Address{
			Address: 	"abc",
			Postcode: 	"123333",
		},
	} 

	fmt.Println(userProfile)
	fmt.Println(userProfile.ToFullDecs())
	byteTxtJson, err := json.MarshalIndent(userProfile, "", "  ")
	if err != nil {
		fmt.Println("err")
	}
	fmt.Println(string(byteTxtJson))

	dataJson := `{
		"firstname": "boomphat",
		"lastname": "pimpim",
		"Age": 21,
		"Height": 0,
		"Address": {
		  "Address": "abc",
		  "Postcode": "123333"
		},
		"Bill": {
		  "BillAddress": ""
		}
	  }`
	var boomProfile UserProfile
	err = json.Unmarshal([]byte(dataJson), &boomProfile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(boomProfile)
}
