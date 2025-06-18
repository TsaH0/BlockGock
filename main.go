package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

// so i will create a outline of what i want
// so a record which has details of what is sent from (string),who(string) and purpose(string) and Amount (int)
// then i will create a hash which is going to be the id for the record and using this id i will create a block, which consists of
// hash(sha256) and this is going to be a sha256 hash,previous hash, block position , the nonce and timestamps this is going to be a struct
// now i will have a blockchain and i will use it to give the difficulty of the blockchain and have a array of blocks but a pointer to the *Block array,
//
// functions overview
// I will have a the structs defined
// first define the mux router and use it to get the record data and create a hash of md5 and assign the id to it ; use it to create a block out of it, by creating a hash of previous data and converting to the sha256 hash, i will give it to a function which keeps on increasing the nonce until, i get the final answer according to the prefix and difficulty-get the blockchain and add to the blockchain
type Block struct {
	Pos           string    `json:"id"`
	Hash          string    `json:"hash"`
	Previous_hash string    `json:"previousHash"`
	Nonce         int       `json:"nonce"`
	Data          string    `json:"data"`
	Timestamps    time.Time `json:"time"`
}

type Record struct {
	Id     string `json:"id"`
	From   string `json:"from"`
	To     string `json:"to"`
	Amount int    `json:"amount"`
	Policy string `json:"Policy"`
}

type Blockchain struct {
	difficulty int
	chain      []*Block
}

func CreateBlockchain(difficulty int) Blockchain {
	Genesis := Block{
		Previous_hash: "0", Hash: "0", Timestamps: time.Now(),
	}
	return Blockchain{
		difficulty: difficulty, chain: []*Block{&Genesis}, //the way i understand is that we return back the address of & Genesis and the original type is *Block
	}
}
func isValid(hash string, difficulty int) bool {

	return strings.HasPrefix(hash, strings.Repeat("0", difficulty))

}
func (b *Block) generateHash(data_for_block string, difficulty int) string {
	hash := ""
	nonce := 0
	if isValid(hash, difficulty) {
		h := sha256.New()
		io.WriteString(h, data_for_block+strconv.Itoa(nonce))
		hash = fmt.Sprint("%x", h.Sum(nil))
		nonce += 1
	}
	b.Nonce = nonce
	b.Hash = hash

}
func (bc *Blockchain) createBlock(data Record) *Block {
	//extract all the fields and get the value of the hash, sha256 hash
	data_to_write, _ := json.Marshal(data)
	data_for_block := strconv.Itoa(data.Amount) + data.From + data.To + data.Policy + data.Id
	prevBlock := bc.chain[len(bc.chain)-1]
	//this is a bit wierd, so how to go about this ,
	b := Block{
		Data: string(data_to_write), Timestamps: time.Now(),Previous_hash: prevBlock.Hash,
	}

	b.generateHash(data_for_block, bc.difficulty)
	return &b
	//how to get the the values after now then
}

// so we finally created a blockchain
func (bc *Blockchain) createRecord(w http.ResponseWriter, r *http.Request) {
	var data Record
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		log.Println("There was an error in the createRecord Handler.The Error is :%v", err)
		w.Write([]byte("There was an error from our side"))
		w.WriteHeader(http.StatusInternalServerError)
	}
	h := md5.New()
	io.WriteString(h, data.From+data.To+data.Policy+strconv.Itoa(data.Amount))
	//now get the hash in terms of bytes and convert to a string
	hash := fmt.Sprint("%x", h.Sum(nil))
	//so i got the hash
	//now i wish
	data.Id = hash
	bl := bc.createBlock(data)
	bc.chain = append(bc.chain, bl)

	resp, _ := json.Marshal(data)
	//but now i wish to use this
	w.Write(resp)
	w.WriteHeader(http.StatusOK)
	//else now that we have the data we wish to parse it from

}
func (bc Blockchain)valid(){
	
}
func (bc Blockchain)displayBlockchain(){
	var bool ifValid=bc.valid()
	if ifValid{
		//this code will display if valid
	}
	else{
		//this code runs if not valid
	}
}
func main() {
	r := mux.NewRouter()
	bc := CreateBlockchain(3)

	r.HandleFunc("/createRecord", bc.createRecord).Methods("POST")
	r.HandleFunc("/display", bc.displayBlockchain).Methods("GET")
	log.Println("Listening on PORT 8000...")
	log.Fatal(http.ListenAndServe(":8000", r))

}
