package main

//Transferer
//Transferes coins from the owners account to another remote user
// Version 3/31/2020
// Author: Sean Worthington & Sam Leary
// This code is property of RAIDA Tech.
// No permission to be used outside of RAIDA Tech.

// Sample use
// transferrer.exe go run transferer2.go -timeout=5 -transactionPath=C:\temp -logpath=C:\temp -idpath=C:\temp\coin.stack -amount=100 -receiverID=sean.cloudcoin.global -fromtag="" -tag="From Change"
// EXIT CODES:
// 1 Errors with flags
// 2 Raida ID is invalid (needs to be from 0-24)
// 3 The Receiver ID is invalid
// 4 Could not retrieve URL
// 5 Could not read show response
// 6 No Coins In this skywallet
// 7 Could not convert Sn to INT
// 12 not enough coins in wallet/envelop
// 13 Can not Find/Create New Folders
// 14 error writing to file
// 15 Cannot open file
// 16
// 17
// 18
// 19
// 20 error getting body response
// 0 Passing

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"math/big"
	"math/rand"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// G L O B A L   V A R I A B L E S
var err error          //an error holder
var tagFilter string   //Filter the coins by tag
var rootpath string    //Root Path
var receiverID int     //Serial number of the receiver
var receiverIDs string //the receiver ID as a string (ip/SN/skywallet)
var tag string         //memo of the transaction
var logpath string     //path to the log folder
var idpath string      // path to the id Coin
//var JustShow bool      //Can be used to just show the currenc skywallet coins (optinal)
var timeout int //how long before the requests time out
var amount int  //The number of coins to transfer.
var idCoin Stack

// S T R U C T U R E S

//Stack Defines a stack of cloudcoins
type Stack struct {
	CloudCoin []CloudCoin `json:"cloudcoin"`
}

//ShowResponse A struct that defines json response
type ShowResponse struct {
	SERVER   string         `json:"server"`
	STATUS   string         `json:"status"`
	MESSAGES []ShowMessages `json:"message"`
	VERSION  string         `json:"version"`
	TIME     string         `json:"time"`
	NN       string         `json:"n"`
	SN       string         `json:"sn"`
}

//ShowMessages a struct that defines a message
type ShowMessages struct {
	SN      string `json:"sn"`
	TAG     string `json:"tag"`
	CREATED string `json:"created"`
}

//CloudCoin a struct that defines each individual coin
type CloudCoin struct {
	NN  string   `json:"nn"`
	SN  string   `json:"sn"`
	ANs []string `json:"an"`
}

//SendResponse creates place for server and status
type SendResponse struct {
	Server string `json:"server"`
	Status string `json:"status"`
	SN     string `json:"sn"`
}

//G L O B A L  M E T H O D S
var raidahttp = &http.Client{
	Timeout: time.Duration(10) * time.Second,
}

// M A I N  F U N C T I O N

func main() {

	//0. Get Flags
	getFlags()

	//.00 Validate Flags
	validateFlags()

	//1. Show Balance
	loadIDCoin()
	needChange := verifyBalance()

	//2. Show Change


	//3. Make Change

	//4. Call Transfer

} //end main

// Gets the flags and sets their value
func getFlags() {
	//flag.BoolVar(&JustShow, "justshow", false, "Whether or not to transfer or just show the current skywallets coins (optional)") //optional
	flag.StringVar(&receiverIDs, "receiverID", "", "The receiver ID") //required
	flag.IntVar(&amount, "amount", 0, "The Number of coins to transfer ")
	flag.StringVar(&tag, "tag", "", "The memo of the coins being sent ")        //required
	flag.StringVar(&logpath, "logpath", "", "The memo of the coins being sent") //required
	flag.StringVar(&rootpath, "transactionPath", "", "A string that shows the path to the transaction log folder ")
	flag.StringVar(&idpath, "idpath", "", "The path to the ID coin")             //required
	flag.StringVar(&tagFilter, "fromtag", ".", "Filter the coins used by tag")   //required
	flag.IntVar(&timeout, "timeout", 5, "How long before a timeout in seconds.") //required
	flag.Parse()
} // End getFlags

//Validate Flags
func validateFlags() {
	//Validate inputs
	total := amount
	if total < 1 {
		ErrStop(1, errors.New("flag errors: negative or zero coin value requested"))
	}
	receiverID = ParseID(receiverIDs)
	if receiverID == 0 {
		ErrStop(12, errors.New("invalid receiver ID Supplied"))
	}
	if tag == "" {
		ErrStop(1, errors.New("flag errors: Missing coin tag "))
	}
	if rootpath == "" || logpath == "" || idpath == "" {
		ErrStop(1, errors.New("flag errors: All File Paths not set"))
	}
	if tagFilter == "." {
		ErrStop(1, errors.New("flag errors: invalid fromTag"))
	}
	if receiverID < 1 {
		ErrStop(3, errors.New("The Receiver ID is invalid"))
		//	var message3 = "The reciever id was not an int."
	}

} //end validate flags

func loadIDCoin() {
	coinbyte, err := os.Open(idpath)
	ErrStop(15, err)
	bytesRead, err := ioutil.ReadAll(coinbyte)
	ErrStop(13, err)
	json.Unmarshal(bytesRead, &idCoin)
} //End load id coin

func verifyBalance() bool {

	//Pick two RAIDA by random and see if there is enough funds to pay
	randRaida1, randRaida2 := randRaida()
	raidaURL1 := fmt.Sprintf("https://raida%d.cloudcoin.global/service/show?nn=1&sn=%s&an=%s&pan=%s&denomination=1", randRaida1, idCoin.CloudCoin[0].SN, idCoin.CloudCoin[0].ANs[randRaida1], idCoin.CloudCoin[0].ANs[randRaida1])
	raidaURL2 := fmt.Sprintf("https://raida%d.cloudcoin.global/service/show?nn=1&sn=%s&an=%s&pan=%s&denomination=1", randRaida2, idCoin.CloudCoin[0].SN, idCoin.CloudCoin[0].ANs[randRaida2], idCoin.CloudCoin[0].ANs[randRaida2])

	resp1, resp2 := listCoinsInWallet(raidaURL1, raidaURL2)

	// read our opened json as a byte array.
	response1Bytes, err := ioutil.ReadAll(resp1.Body)
	ErrStop(5, err)

	response2Bytes, err := ioutil.ReadAll(resp2.Body)
	ErrStop(5, err)

	var sn []string // an array of serial numbers
	var values []int

	var ShowResponses1 ShowResponse
	var ShowResponses2 ShowResponse
	json.Unmarshal(response1Bytes, &ShowResponses1)
	json.Unmarshal(response2Bytes, &ShowResponses2)

	for i := 0; i < len(ShowResponses1.MESSAGES); i++ {
		currentValue := ShowResponses1.MESSAGES[i].SN
		if tagFilter == "" {

			for ix := 0; ix < len(ShowResponses2.MESSAGES); ix++ {

				if currentValue == ShowResponses2.MESSAGES[ix].SN {

					sn = append(sn, ShowResponses1.MESSAGES[i].SN)
					stringsn, err := strconv.Atoi(sn[len(sn)-1])
					ErrStop(7, err)
					worth := denomination(stringsn)
					values = append(values, worth)

				}

			}
		} else {
			currentTag := ShowResponses1.MESSAGES[i].TAG
			if currentTag == tagFilter {
				for ix := 0; ix < len(ShowResponses2.MESSAGES); ix++ {
					if currentValue == ShowResponses2.MESSAGES[ix].SN {
						sn = append(sn, ShowResponses1.MESSAGES[i].SN)
						stringsn, err := strconv.Atoi(sn[len(sn)-1])
						ErrStop(7, err)
						worth := denomination(stringsn)
						values = append(values, worth)

					}

				}
			}
		}

	}
	if len(ShowResponses1.MESSAGES) == 0 {
		ErrStop(6, errors.New("No coins in this skywallet"))
	}

	tempAmount := amount

	var snToSend = []string{}
	var denominations = []int{}

	for i := 0; tempAmount != 0; i++ {

		placement := len(values) - 1 - i

		if placement < 0 {
			break
		}

		if values[placement] == 250 {

			if tempAmount >= 250 {
				tempAmount -= values[placement]
				snToSend = append(snToSend, sn[placement])
				denominations = append(denominations, 250)
			}
		}
		if values[placement] == 100 {

			if tempAmount >= 100 {
				tempAmount -= values[placement]
				snToSend = append(snToSend, sn[placement])
				denominations = append(denominations, 100)
			}
		}

		if values[placement] == 25 {

			if tempAmount >= 25 {
				tempAmount -= values[placement]
				snToSend = append(snToSend, sn[placement])
				denominations = append(denominations, 25)
			}
		}

		if values[placement] == 5 {
			if tempAmount >= 5 {
				tempAmount -= values[placement]
				snToSend = append(snToSend, sn[placement])
				denominations = append(denominations, 5)
			}
		}
		if values[placement] == 1 {

			tempAmount -= values[placement]
			snToSend = append(snToSend, sn[placement])
			denominations = append(denominations, 1)

		}

	}

	if tempAmount > 0 {
		message7 := "There were not enough coins to fulfill the request."
		err := fmt.Errorf("error with inventory: %v", message7)
		ErrStop(12, err)

	}
} //end verify balance

func randRaida() (int, int) {

	//seed  Random Servers
	rand.Seed(time.Now().UnixNano())

	var checkRaida1 = rand.Intn(25) //set the first raida to check
	var checkRaida2 = rand.Intn(25) //set the second raida to check

	for checkRaida1 == checkRaida2 {
		checkRaida1--
		checkRaida2++

		if checkRaida1 < 0 {
			checkRaida1 = 0
			checkRaida1 += rand.Intn(3)
		} else if checkRaida1 > 24 {
			checkRaida1 = 24
			checkRaida1 -= rand.Intn(3)
		}

		if checkRaida2 < 0 {
			checkRaida1 = 0
			checkRaida1 += rand.Intn(3)
		} else if checkRaida2 > 24 {
			checkRaida2 = 24
			checkRaida2 -= rand.Intn(3)

		}
	}
	return checkRaida1, checkRaida2
} // end rand raida

//ErrStop checks if there is an error, and if so stops the program pringing the error message and the error number
func ErrStop(i int, err error) {
	if err != nil {
		fmt.Printf("{\"status\":\"fail\",\"message\":\"error %d. %s.  %v\"}", i, fmt.Sprintf("%s", err), time.Since(time.Now()))
		os.Exit(0)
	}
}

func listCoinsInWallet(raida1, raida2 string) (*http.Response, *http.Response) {

	response1, err := raidahttp.Get(raida1)
	ErrStop(4, err)

	response2, err := raidahttp.Get(raida2)
	ErrStop(4, err)

	return response1, response2

}

//ParseID parses the supplied id into a usable serial number
func ParseID(parseID string) int {

	if _, err := strconv.Atoi(parseID); err == nil {

		result, _ := strconv.Atoi(parseID)
		return result

	} else {

		addr := net.ParseIP(parseID)
		if addr == nil {
			ips, err := net.LookupIP(parseID)
			ErrStop(12, err)
			parseID = strings.Replace(ips[0].String(), "1.", "0.", 1)
			input := net.ParseIP(parseID)
			ipOut := IP4toInt(input)
			return int(ipOut)

		} else {
			if strings.HasPrefix(parseID, "1.") {
				parseID = strings.Replace(parseID, "1.", "0.", 1)
				input := net.ParseIP(parseID)
				ipOut := IP4toInt(input)

				return int(ipOut)
			}
		}
		return 0
	}
	return 0
}

func IP4toInt(input net.IP) int64 {
	//ip := net.IP(input)

	//ip, _, err := net.ParseCIDR(input)
	//ErrStop(32, err, t)
	IPv4Int := big.NewInt(0)
	IPv4Int.SetBytes(input.To4())
	return IPv4Int.Int64()
}

func readResponse(responses []string) ([]bool, []int) {
	//create necessary arrays
	Passes := make([]bool, 25)
	sortedArray := make([]int, 25)
	//loop through and read each response: and determine if it was a success or a fail

	for i := 0; i < len(responses); i++ {

		stringreader := strings.FieldsFunc(responses[i], func(r rune) bool {
			if r == ':' || r == ',' {
				return true
			}
			return false
		})

		if len(stringreader) < 3 {

			Passes[sortedArray[i]] = false

			return Passes, sortedArray
		}

		stringreader[1] = strings.Replace(stringreader[1], "RAIDA", "", -1)
		stringreader[1] = strings.Replace(stringreader[1], "\"", "", -1)

		if stringreader[3] == "\"pass\"" {
			sortedArray[i], _ = strconv.Atoi(stringreader[1])

			Passes[sortedArray[i]] = true
		} else {
			sortedArray[i], _ = strconv.Atoi(stringreader[1])
			Passes[sortedArray[i]] = false
		}

	}
	return Passes, sortedArray
}

//denomination determines the denomination of the coin
func denomination(sn int) int {
	var returnInt int
	returnInt = 0

	if sn >= 1 && sn <= 2097152 {
		returnInt = 1
	} else if sn <= 4194304 {
		returnInt = 5
	} else if sn <= 6291456 {
		returnInt = 25
	} else if sn <= 14680064 {
		returnInt = 100
	} else if sn <= 16777216 {
		returnInt = 250
	}
	return returnInt

} //end func denomination
