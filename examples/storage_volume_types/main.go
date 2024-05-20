package main

import (
	"fmt"
	"log"

	"github.com/gomorpheus/morpheus-go-sdk"
)

func main() {
	client := morpheus.NewClient("https://yourmorpheus.com")
	client.SetUsernameAndPassword("username", "password")
	resp, err := client.Login()
	if err != nil {
		fmt.Println("LOGIN ERROR: ", err)
	}
	fmt.Println("LOGIN RESPONSE:", resp)

	// List node types
	req := &morpheus.Request{}
	storageVolumeTypeResponse, err := client.ListStorageVolumeTypes(req)
	if err != nil {
		log.Fatal(err)
	}
	result := storageVolumeTypeResponse.Result.(*morpheus.ListStorageVolumeTypesResult)
	log.Println(result)
}
