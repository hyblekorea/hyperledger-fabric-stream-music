package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type SmartContract struct {}

type Wallet struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	Token string `json:"token"`
}

type Music struct {
	Title    string `json:"title"`
	Singer   string `json:"singer"`
	Price    string `json:"price"`
	WalletID    string `json:"walletid"`
}

type MusicKey struct {
	Key string
	Idx int
}

func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) pb.Response {
	function, args := APIstub.GetFunctionAndParameters()
	
	if function == "initWallet" {
		return s.initWallet(APIstub)
	} else if function == "getWallet" {
		return s.getWallet(APIstub, args)
	} else if function == "setMusic" {
		return s.setMusic(APIstub, args)
	} else if function == "getAllMusic" {
		return s.getAllMusic(APIstub)
	} else if function == "purchaseMusic" {
		return s.purchaseMusic(APIstub, args)
	}
	fmt.Println("Please check your function : " + function)
	return shim.Error("Unknown function")
}

func (s *SmartContract) initWallet(APIstub shim.ChaincodeStubInterface) pb.Response {

	//Declare wallets
	seller := Wallet{Name: "Hyper", ID: "1Q2W3E4R", Token: "100"}
	customer := Wallet{Name: "Ledger", ID: "5T6Y7U8I", Token: "200"}

	// Convert seller to []byte
	SellerasJSONBytes, _ := json.Marshal(seller)
	err := APIstub.PutState(seller.ID, SellerasJSONBytes)
	if err != nil {
		return shim.Error("Failed to create asset " + seller.Name)
	}
	// Convert customer to []byte
	CustomerasJSONBytes, _ := json.Marshal(customer)
	err = APIstub.PutState(customer.ID, CustomerasJSONBytes)
	if err != nil {
		return shim.Error("Failed to create asset " + customer.Name)
	}

	return shim.Success(nil)
}

func (s *SmartContract) getWallet(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {

	walletAsBytes, err := APIstub.GetState(args[0])
	if err != nil {
		fmt.Println(err.Error())
	}

	wallet := Wallet{}
	json.Unmarshal(walletAsBytes, &wallet)

	var buffer bytes.Buffer
	buffer.WriteString("[")
	bArrayMemberAlreadyWritten := false

	if bArrayMemberAlreadyWritten == true {
		buffer.WriteString(",")
	}
	buffer.WriteString("{\"Name\":")
	buffer.WriteString("\"")
	buffer.WriteString(wallet.Name)
	buffer.WriteString("\"")

	buffer.WriteString(", \"ID\":")
	buffer.WriteString("\"")
	buffer.WriteString(wallet.ID)
	buffer.WriteString("\"")

	buffer.WriteString(", \"Token\":")
	buffer.WriteString("\"")
	buffer.WriteString(wallet.Token)
	buffer.WriteString("\"")

	buffer.WriteString("}")
	bArrayMemberAlreadyWritten = true
	buffer.WriteString("]\n")

	return shim.Success(buffer.Bytes())

}

func generateKey(APIstub shim.ChaincodeStubInterface, key string) []byte {

	var isFirst bool = false

	musickeyAsBytes, err := APIstub.GetState(key)
	if err != nil {
		fmt.Println(err.Error())
	}

	musickey := MusicKey{}
	json.Unmarshal(musickeyAsBytes, &musickey)
	var tempIdx string
	tempIdx = strconv.Itoa(musickey.Idx)
	fmt.Println(musickey)
	fmt.Println("Key is " + strconv.Itoa(len(musickey.Key)))
	if len(musickey.Key) == 0 || musickey.Key == "" {
		isFirst = true
		musickey.Key = "MS"
	}
	if !isFirst {
		musickey.Idx = musickey.Idx + 1
	}

	fmt.Println("Last MusicKey is " + musickey.Key + " : " + tempIdx)

	returnValueBytes, _ := json.Marshal(musickey)

	return returnValueBytes
}

func (s *SmartContract) setMusic(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}
	
	var musickey = MusicKey{}
	json.Unmarshal(generateKey(APIstub, "latestKey"), &musickey)
	keyidx := strconv.Itoa(musickey.Idx)
	fmt.Println("Key : " + musickey.Key + ", Idx : " + keyidx)

	var music = Music{Title: args[0], Singer: args[1], Price: args[2], WalletID: args[3]}
	musicAsJSONBytes, _ := json.Marshal(music)

	var keyString = musickey.Key + keyidx
	fmt.Println("musickey is " + keyString)

	err := APIstub.PutState(keyString, musicAsJSONBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to record music catch: %s", musickey))
	}

	musickeyAsBytes, _ := json.Marshal(musickey)
	APIstub.PutState("latestKey", musickeyAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) getAllMusic(APIstub shim.ChaincodeStubInterface) pb.Response {
	
	// Find latestKey
	musickeyAsBytes, _ := APIstub.GetState("latestKey")
	musickey := MusicKey{}
	json.Unmarshal(musickeyAsBytes, &musickey)
	idxStr := strconv.Itoa(musickey.Idx + 1)

	var startKey = "MS0"
	var endKey = musickey.Key + idxStr
	fmt.Println(startKey)
	fmt.Println(endKey)

	resultsIter, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIter.Close()
	
	var buffer bytes.Buffer
	buffer.WriteString("[")
	bArrayMemberAlreadyWritten := false
	for resultsIter.HasNext() {
		queryResponse, err := resultsIter.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")
		
		buffer.WriteString(", \"Record\":")
		
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]\n")
	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) purchaseMusic(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	var tokenFromKey, tokenToKey int // Asset holdings
	var musicprice int // Transaction value
	var err error

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	musicAsBytes, err := APIstub.GetState(args[1])
	if err != nil {
		return shim.Error(err.Error())
	}

	music := Music{}
	json.Unmarshal(musicAsBytes, &music)
	musicprice, _ = strconv.Atoi(music.Price)

	SellerAsBytes, err := APIstub.GetState(music.WalletID)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if SellerAsBytes == nil {
		return shim.Error("Entity not found")
	}
	seller := Wallet{}
	json.Unmarshal(SellerAsBytes, &seller)
	tokenToKey, _ = strconv.Atoi(seller.Token)

	CustomerAsBytes, err := APIstub.GetState(args[0])
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if CustomerAsBytes == nil {
		return shim.Error("Entity not found")
	}

	customer := Wallet{}
	json.Unmarshal(CustomerAsBytes, &customer)
	tokenFromKey, _ = strconv.Atoi(string(customer.Token))

	customer.Token = strconv.Itoa(tokenFromKey - musicprice)
	seller.Token = strconv.Itoa(tokenToKey + musicprice)
	updatedCustomerAsBytes, _ := json.Marshal(customer)
	updatedSellerAsBytes, _ := json.Marshal(seller)
	APIstub.PutState(args[0], updatedCustomerAsBytes)
	APIstub.PutState(music.WalletID, updatedSellerAsBytes)

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	buffer.WriteString("{\"Customer Token\":")
	buffer.WriteString("\"")
	buffer.WriteString(customer.Token)
	buffer.WriteString("\"")

	buffer.WriteString(", \"Seller Token\":")
	buffer.WriteString("\"")
	buffer.WriteString(seller.Token)
	buffer.WriteString("\"")

	buffer.WriteString("}")
	buffer.WriteString("]\n")

	return shim.Success(buffer.Bytes())
}

func main() {

	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
