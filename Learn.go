package main
import (
	"encoding/json"
	"fmt"

	"io/ioutil"
	"log"
	"net/http"

)

type Foo struct{
	Info [] Data `json:"data"`
	Included [] string `json:"included"`
	Links Linksstruct `json:"links"`
}

type Linksstruct struct{
	Next string `json:"next"`
	Self string `json:"self"`
}

type Data struct {
	Id string `json:"id"`
	Type1 string `json:"type"`
	Attributes FirstStruc `json:"attributes"`
}
type FirstStruc struct{
	Value SecondStruc `json:"value"`
	U32 int `json:"u32"`
}
type SecondStruc struct{
	Type2 ThirdStruc `json:"type"`
}
type ThirdStruc struct{
	Value int `json:"value"`
	Name string `json:"name"`
}

type Item struct {
	Foos [] Foo
}
func main() {
	mylink:="http://localhost:8000/_/api/v3/key_values"
	for i := 0; i < 15; i++ {
	response, err := http.Get(mylink)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()



	dataInBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("failed to read json file, error: %v", err)
		return
	}


	//Array:=Item{}
    Datan:=Foo{}


	//json.Unmarshal([]byte(dataInBytes), &Datan)
	//fmt.Println(Datan)

	err0:= json.Unmarshal([]byte(dataInBytes), &Datan)
	if err0 != nil {
		fmt.Println(err0)
		return
	}

	fmt.Println("Self=",Datan.Links.Self)
	fmt.Println("Next=",Datan.Links.Next)
	mylink="http://localhost:8000/_/api"+Datan.Links.Next}

	

}

