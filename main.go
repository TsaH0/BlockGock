package main

import("fmt"
	"io"
	"strings"
	"strconv"
	"crypto/sha256"
	"crypto/md5"
	"github.com/gorilla/mux"
	"log"
	"encoding/json"
	"net/http"
	"time")
type Block struct{
	hash string `json:"hash"`
	previous_hash string `json:"prevhash"`
	timestamps time.Time `json:"timestamps"`
	difficulty int
	data Record `json:"data"`
	id int `json:"id"`
}
type BlockChain struct{
	chain []Block
	difficulty int
	GenesisBlock Block
}
type Record struct{
	transactions map[string]interface{} `json:"transactions"`

}
func CreateBlockChain(difficulty int )BlockChain{
	Genesis:=Block{
		hash:"0",timestamps:time.Now(),id:0,
	}
	return BlockChain{
	[]Block{Genesis},difficulty,Genesis}

}
func createRecord(w http.ResponseWriter, r *http.Request){
	record:=map[string]interface{
		from:from,to:to,amount:amount
	}
}
func main(){
	r:=mux.NewRouter()

	var BlockGock Blockchain=CreateBlockChain(3)
	r.HandleFunc("/createRecord",createRecord)
	log.Println("The Server has started at PORT 8080")
	log.Fatal(http.ListenAndServe(":8080",r))
}
