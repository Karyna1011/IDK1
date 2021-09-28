package main
import (
	"encoding/json"
	"fmt"

	"io/ioutil"
	"log"
	"net/http"

)

type PageContent struct{
	Info [] Data `json:"data"`
	Included [] string `json:"included"`
	Links Linkstruct `json:"links"`
}

type Linkstruct struct{
	Next string `json:"next"`
	Self string `json:"self"`
}

type Data struct {
	Id string `json:"id"`
	TypeData string `json:"type"`
	Attributes AttributesStruct `json:"attributes"`
}
type AttributesStruct struct{
	Value ValueStruct `json:"value"`
	U32 int `json:"u32"`
}
type ValueStruct struct{
	TypeValue TypeStruct `json:"type"`
}
type TypeStruct struct{
	Value int `json:"value"`
	Name string `json:"name"`
}

type Item struct {
	Pages [] PageContent
}
func main() {
	MyLink:="http://localhost:8000/_/api/v3/key_values"
	length:=1
	Datan:=PageContent{}
	response, err := http.Get(MyLink)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()



	dataInBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("failed to read json file, error: %v", err)
		return
	}
	
	err0:= json.Unmarshal([]byte(dataInBytes), &Datan)
	if err0 != nil {
		fmt.Println(err0)
		return
	}

	fmt.Println("Self=",Datan.Links.Self)
	fmt.Println("Next=",Datan.Links.Next)

	length=len(Datan.Info)
	MyLink="http://localhost:8000/_/api"+Datan.Links.Next

	for i := 0; length!=0 ; i++ {
		        fmt.Println("Data=",Datan.Info[0])
			response, err := http.Get(MyLink)
			if err != nil {
				log.Fatal(err)
			}
			defer response.Body.Close()



			dataInBytes, err := ioutil.ReadAll(response.Body)
			if err != nil {
				fmt.Printf("failed to read json file, error: %v", err)
				return
			}
			
			error:= json.Unmarshal([]byte(dataInBytes), &Datan)
			if error != nil {
				fmt.Println(err0)
				return
			}

			fmt.Println("Self=",Datan.Links.Self)
			fmt.Println("Next=",Datan.Links.Next)

			length=len(Datan.Info)
			MyLink="http://localhost:8000/_/api"+Datan.Links.Next}



}
