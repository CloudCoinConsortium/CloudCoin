package main

//transfer_with_change.go
//Transferes coins from the owners account to another remote user even if the sender does not have the correct change.
// Version 1/6/2020
// Author: Sean Worthington & Sam Leary
// This code is property of RAIDA Tech.
// No permission to be used outside of RAIDA Tech.

// Sample use
// transfer.exe -amount=250 -server=18 -rootpath=C:\Bank -logto=C:\Bank\Logs\Transfer -id=C:\Users\userX\CloudCoinWallet\CloudCoin\Accounts\Userx\ID\idcoin.stack -receiversn=xxxxxxx -tagfilter="Spending" -tag="ExportedCoins" -justShow=false
// EXIT CODES:
// 1 Errors with flags
// 2 Raida ID is invalid (needs to be from 0-24)
// 3 The Receiver ID is invalid
// 4 Could not retrieve URL
// 5 Could not read show response
// 6 No Coins In this skywallet
// 7 Could not convert Sn to INT
// 8 not enough 1s
// 9 not enough 5s
// 10 not enough 25s
// 11 not enough 100s
// 12 not enough 250s
// 13 Can not Find/Create New Folders
// 14 error writing to file
// 15
// 16
// 17
// 18
// 19
// 20 error getting body response
// 0 Passing

/*
Send error back if raida could not be contaced. THen the progrm can try another RAIDA.
Transfer with change not just transfer
SHould require location of the ID coin.
Error if no ID coin.
Change RAIDA for showing coins, if the Raida Proided si bd.
Don't ask them for the RAIDA id. Just randomny guess one.

Add Documentation
Add a PHP page as an example

*/

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
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

// S T R U C T U R E S
//jsonChange is the struct that defines the json response from the raida
type jsonChange struct {
	Server  string   `json:"server"`
	Status  string   `json:"status"`
	Owner   string   `json:"owner"`
	D1      []string `json:"d1"`
	D5      []string `json:"d5"`
	D25     []string `json:"d25"`
	D100    []string `json:"d100"`
	Message string   `json:"message"`
	Version string   `json:"version"`
	Time    string   `json:"time"`
}

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

// M A I N  F U N C T I O N

func main() {

	var err error
	//create Variables

	loc := time.FixedZone("UTC", 0) //set the timezone
	now := time.Now().In(loc)       //set the current time
	t := time.Now()                 //set the variable t
	var tagFilter string            //Filter the coins by tag
	var receiveServer int           //example: Raida0.cloudcoin.global
	var rootpath string             //Root Path
	var receiverID int              //Serial number of the receiver
	var receiverAccount string      //the account of the receiver
	var intOnes int                 // number of ones to send
	var intFives int                // number of fives to send
	var intTwentyfives int          // number of twentyfives to send
	var intHundreds int             // number of hundreds to send
	var intTwohundredfifties int    // number of twohundredfifties to send
	var tag string                  //memo of the transaction
	var logpath string              //path to the log folder
	var idpath string               // path to the id Coin
	var JustShow bool
	//Gather 13 flags
	flag.BoolVar(&JustShow, "justshow", false, "Whether or not to transfer or just show the current skywallets coins (optional)") //optional
	flag.IntVar(&receiveServer, "server", -1, "The Server ID")                                                                    //required
	flag.StringVar(&receiverAccount, "receiversn", "", "The receiver envelope id")                                                //required
	flag.StringVar(&rootpath, "rootpath", "", "The path to the Root Directory")                                                   //required
	flag.IntVar(&intOnes, "1s", 0, "The total Number of ones sent")                                                               //optional
	flag.IntVar(&intFives, "5s", 0, "The total Number of fives sent")                                                             //optional
	flag.IntVar(&intTwentyfives, "25s", 0, "The total Number of twentyfives sent")                                                //optional
	flag.IntVar(&intHundreds, "100s", 0, "The total Number of hundreds sent")                                                     //optional
	flag.IntVar(&intTwohundredfifties, "250s", 0, "The total Number of twofifties sent")                                          //optional
	flag.StringVar(&tag, "tag", "", "The memo of the coins being sent ")                                                          //required
	flag.StringVar(&logpath, "logto", "", "The memo of the coins being sent")                                                     //required
	flag.StringVar(&idpath, "id", "", "The path to the ID coin")                                                                  //required
	flag.StringVar(&tagFilter, "tagfilter", "", "Filter the coins used by tag (optional)")                                        //optional
	//parse flags
	flag.Parse()

	receiverID = ParseID(receiverAccount, t)

	total := intOnes + intFives*5 + intTwentyfives*25 + intHundreds*100 + intTwohundredfifties*250
	if flag.NArg() > 6 {
		flag.Usage()
		ErrStop(1, errors.New("flag errors: Missing Too many Flags"), t)
	}
	if total <= 0 {
		ErrStop(1, errors.New("flag errors: negative or zero coin value requested"), t)
	}
	if receiveServer == -1 || receiverID == -1 {
		ErrStop(1, errors.New("flag errors: receiver Information incorrect"), t)
	}
	if tag == "" {
		ErrStop(1, errors.New("flag errors: Missing coin tag "), t)
	}
	if rootpath == "" || logpath == "" || idpath == "" {
		ErrStop(1, errors.New("flag errors: All File Paths not set"), t)
	}

	/* Validate inputs */
	//Validate recieveServer
	checkServer := fmt.Sprintf("https://Raida%d.cloudcoin.global", receiveServer)

	if !isValidURL(checkServer, t) {
		errURL := errors.New("Invalid URl" + checkServer)
		ErrStop(2, errURL, t)
		//	var message2 = "The server was an invalid uri."

	} //end if

	//Validate reciever id
	if receiverID <= 0 {
		ErrStop(3, errors.New("The Receiver ID is invalid"), t)
		//	var message3 = "The reciever id was not an int."

	}

	//open and read the ID Coin
	coinbyte, err := os.Open(idpath)
	bytesRead, err := ioutil.ReadAll(coinbyte)
	var idCoin Stack
	json.Unmarshal(bytesRead, &idCoin)

	showURL := fmt.Sprintf("https://raida%d.cloudcoin.global/service/show?nn=1&sn=%s&an=%s&pan=%s&denomination=1", receiveServer, idCoin.CloudCoin[0].SN, idCoin.CloudCoin[0].ANs[receiveServer], idCoin.CloudCoin[0].ANs[receiveServer])

	//	fmt.Println("show URL: " + showURL + "\n")

	resp, err := http.Get(showURL)
	ErrStop(4, err, t)
	// if os.Open returns an error then handle it

	// read our opened json as a byte array.
	responseBytes, err := ioutil.ReadAll(resp.Body)
	ErrStop(5, err, t)

	var sn []string // an array of serial numbers
	var values []int
	var ShowResponses ShowResponse
	json.Unmarshal(responseBytes, &ShowResponses)
	if len(ShowResponses.MESSAGES) == 0 {
		ErrStop(6, errors.New("No coins in this skywallet"), t)
	}
	//fmt.Println(tagFilter)

	if tagFilter == "" {
		//build the arrays for the Send Request
		for i := 0; i < len(ShowResponses.MESSAGES); i++ {

			sn = append(sn, ShowResponses.MESSAGES[i].SN)
			stringsn, err := strconv.Atoi(sn[i])
			ErrStop(7, err, t)
			worth := denomination(stringsn)
			values = append(values, worth)
		} //end for every an
	} else {
		for i := 0; i < len(ShowResponses.MESSAGES); i++ {
			if len(ShowResponses.MESSAGES) == 0 {
				ErrStop(6, errors.New("No Coins found with that Tag"), t)
			}
			if ShowResponses.MESSAGES[i].TAG == tagFilter {
				sn = append(sn, ShowResponses.MESSAGES[i].SN)
				stringsn, err := strconv.Atoi(ShowResponses.MESSAGES[i].SN)
				ErrStop(7, err, t)
				worth := denomination(stringsn)
				values = append(values, worth)
			}

		}
	}

	if JustShow {
		var firstEntry []string
		listed := ""
		seperator := ","
		for i := 0; i < len(ShowResponses.MESSAGES); i++ {
			Entry := fmt.Sprintf("%s : %d", sn[i], values[i])
			firstEntry = append(firstEntry, Entry)
			if i == len(ShowResponses.MESSAGES)-1 {
				seperator = ""
			}
			listed = fmt.Sprintf("%s|%s|%s", listed, firstEntry[i], seperator)
		}
		fmt.Println("Finished Showing Coins")
		fmt.Printf("%s", listed)
		os.Exit(0)

	}

	//Now get lists of all the coins that will be sent
	//Loop through all coin names. If they are needed add them to the pile
	//If they are not needed do not.
	unusedFives := []string{}
	unusedTwentyFives := []string{}
	unusedHundreds := []string{}
	unusedTwoFifties := []string{}
	var snToSend = []string{}
	var denominations = []int{}
	for i, d := range values {
		switch d {
		case 250:
			if intTwohundredfifties > 0 { //if a 250 note is needed
				snToSend = append(snToSend, sn[i])
				denominations = append(denominations, 250)
				intTwohundredfifties-- //reduce the ones neeeded by one
			} else {
				unusedTwoFifties = append(unusedTwoFifties, sn[i])
			} //end if needed greater than zero
			break
		case 100:
			if intHundreds > 0 { //if a 100 note is needed
				snToSend = append(snToSend, sn[i])
				denominations = append(denominations, 100)
				intHundreds-- //reduce the ones neeeded by one
			} else {
				unusedHundreds = append(unusedHundreds, sn[i])
			} //end if needed greater than zero
			break
		case 25:
			if intTwentyfives > 0 { //if a 25 note is needed
				snToSend = append(snToSend, sn[i])
				denominations = append(denominations, 25)
				intTwentyfives-- //reduce the ones neeeded by one
			} else {
				unusedTwentyFives = append(unusedTwentyFives, sn[i])
			} //end if needed greater than zero
			break
		case 5:
			if intFives > 0 { //if a 5 note is needed
				snToSend = append(snToSend, sn[i])
				denominations = append(denominations, 5)
				intFives-- //reduce the ones neeeded by one
			} else {
				unusedFives = append(unusedFives, sn[i])
			} //end if needed greater than zero
			break
		case 1:
			if intOnes > 0 { //if a one note is needed
				snToSend = append(snToSend, sn[i])
				denominations = append(denominations, 1)
				intOnes-- //reduce the ones neeeded by one
			} //end if needed greater than zero
			break
		} // end switch

		// Do we have all notes we need?

	} //end for each file in the folder

	missingCoins := intOnes + (intFives * 5) + (intTwentyfives * 25) + (intHundreds * 100) + (intTwohundredfifties * 250)
	//fmt.Print(missingCoins)
	changeSns := []string{"empty"}
	denom := 0
	changeStack := []string{}
	coin1, _ := strconv.Atoi(snToSend[0])
	fmt.Println(denomination(coin1))
	coin1, _ = strconv.Atoi(snToSend[1])
	fmt.Println(denomination(coin1))

	if missingCoins > 0 {

		if missingCoins > 100 {
			//snToSend[0] = unusedTwoFifties[0]
			//removeIndex(snToSend, 0)
			changeSns[0] = unusedTwoFifties[0]
			denom = 250
		} else if missingCoins > 25 {
			//snToSend[0] = unusedHundreds[0]
			changeSns[0] = unusedHundreds[0]
			//removeIndex(snToSend, 0)
			denom = 100
		} else if missingCoins > 5 {
			//snToSend[0] = unusedTwentyFives[0]
			//removeIndex(snToSend, 0)
			changeSns[0] = unusedTwentyFives[0]
			denom = 25

		} else {
			denom = 5
			//snToSend[0] = unusedFives[0]
			//removeIndex(snToSend, 0)
			changeSns[0] = unusedFives[0]
			//	fmt.Printf("%v\r\n", snToSend)
		}
		coin1, _ = strconv.Atoi(snToSend[0])
		fmt.Println(denomination(coin1))
		coin1, _ = strconv.Atoi(snToSend[1])
		fmt.Println(denomination(coin1))

		change := showChange("2", denom, 0, 5, 10)

		changeAmount := 0

		//		fmt.Printf("needs: 1-%d 5-%d 25-%d 100-%d  ", intOnes, intFives, intTwentyfives, intHundreds)
		for changeAmount < denom {

			for intOnes > 0 {
				changeStack = append(changeStack, change[0][rand.Intn(len(change[0]))])
				changeAmount++
				intOnes--
			}
			for intFives > 0 {
				changeStack = append(changeStack, change[1][rand.Intn(len(change[1]))])
				changeAmount += 5
				intFives--
			}

			for intTwentyfives > 0 {
				changeStack = append(changeStack, change[2][rand.Intn(len(change[2]))])
				changeAmount += 25
				intTwentyfives--
			}
			for intHundreds > 0 {
				changeStack = append(changeStack, change[3][rand.Intn(len(change[3]))])
				changeAmount += 100
				intHundreds--
			}

			if denom-changeAmount > 100 {
				changeStack = append(changeStack, change[3][rand.Intn(len(change[3]))])
				changeAmount += 100
			} else if denom-changeAmount > 25 {
				changeStack = append(changeStack, change[2][rand.Intn(len(change[2]))])
				changeAmount += 25
			} else if denom-changeAmount > 5 {
				changeStack = append(changeStack, change[1][rand.Intn(len(change[1]))])
				changeAmount += 5
			} else {
				changeStack = append(changeStack, change[0][rand.Intn(len(change[0]))])
				changeAmount++
			}

			//			fmt.Println("remaining needed", denom-changeAmount)
		}
		//		fmt.Println("change amount = ", changeAmount, "\r\n chang serialNumbers:", changeStack)
	}

	//	os.Exit(200)
	//Now we have a list of coins that we can send.

	/*TRANSFER COINS*/

	var sendURL = [25]string{}
	var URLData = []url.Values{}
	var NN1 = "1"

	toSN := fmt.Sprintf("%d", receiverID)
	for i := 0; i < 25; i++ {

		sendURL[i] = fmt.Sprintf("https://RAIDA%d.cloudcoin.global/service/transfer_with_change", i)
		URLData = append(URLData, url.Values{})
		//URLData[i].Set("tag", tag)
		URLData[i].Set("nn", NN1)
		URLData[i].Set("sn", idCoin.CloudCoin[0].SN)
		URLData[i].Set("an", idCoin.CloudCoin[0].ANs[i])
		URLData[i].Set("pan", idCoin.CloudCoin[0].ANs[i])
		URLData[i].Set("to_sn", toSN)
		URLData[i].Set("denomination", "1")

		for j := 0; j < len(snToSend); j++ {
			if j == 0 {
				//	URLData[i].Set("paysns[]", snToSend[j])
				URLData[i].Set("sns[]", snToSend[j])
				URLData[i].Set("nns[]", "1")
				//fmt.Println(snToSend[j])
			} else {
				//URLData[i].Add("paysns[]", snToSend[j])
				URLData[i].Add("sns[]", snToSend[j])
				URLData[i].Add("nns[]", "1")

			}
		} //end for each file

		if missingCoins > 0 {

			for j := 0; j < len(changeStack); j++ {
				if j == 0 {
					URLData[i].Set("chsns[]", changeStack[j])
				} else {
					URLData[i].Add("chsns[]", changeStack[j])
				}

			}
		}
		URLData[i].Set("paysns[]", changeSns[0])
		totalString := fmt.Sprintf("%d", total)
		URLData[i].Set("payment_required", totalString)
		URLData[i].Set("public_change_maker", "2")
		URLData[i].Set("payment_envelope", tag)
		//URLData[i].Set("paysns[]", changeSns[0])

		//	fmt.Printf("\r\ntest %s\r\n", changeSns[0])
	} //end for each raida request

	//Send our requests
	var responses [25]string
	var requests [25]string
	//var byteresponse []byte
	allresponses := ""
	allrequests := ""
	if tagFilter != "" {
		allrequests += "Using tag Filter:\r\n"
	}
	done := make(chan [3]string) // make the "done" channel
	for i := range responses {
		responses[i] = "no response" //makes two dime array
	} //end of i forloop
	//var passes int

	for i := range responses {
		go Transfer(done, sendURL[i], URLData[i], i, t)
	}

	for i := range responses {

		//Go Routine Responses
		results := <-done // receive from the channel
		responses[i] = results[0]
		requests[i] = results[1]
		elapsedTime := results[2]
		allresponses = fmt.Sprintf("%s |%s|%s\r\n\r\n", allresponses, responses[i], elapsedTime)
		allrequests = fmt.Sprintf("%s |%s|\r\n\r\n", allrequests, requests[i])

	} // end of i for loop

	path := fmt.Sprintf("%s\\%d_sent_to_%d", logpath, total, receiverID)

	err = os.MkdirAll(path, os.ModePerm)
	ErrStop(13, err, t)
	lpath := fmt.Sprintf("%s\\%s_requests.txt", path, tag)
	path = fmt.Sprintf("%s\\%s_responses.txt", path, tag)

	WriteToFile(path, allresponses, t)
	WriteToFile(lpath, allrequests, t)

	fmt.Printf("{\"status\":\"success\",\"message\":\" Execution Time = %s\",\"time\":\"%s\"}", time.Since(t), now.Format("2006-1-2 15:04:05"))
} //end main

// H E L P E R  F U N C T I O N S

//WriteToFile data to a file
func WriteToFile(filename string, data string, t time.Time) error {

	//file, err := os.Create(filename)
	err := ioutil.WriteFile(filename, []byte(data), 0666)
	ErrStop(14, err, t)

	return err
} //end of write to file

//isValidURL checks to see if URL is valid
func isValidURL(toTest string, t time.Time) bool {
	_, err := url.ParseRequestURI(toTest)
	ErrStop(4, err, t)
	return true
} // end is Valid URL

//Transfer sends a http post to the raida
func Transfer(done chan [3]string, sendURL string, URLData url.Values, raidaID int, t time.Time) {
	var responseText [3]string
	start := time.Now()
	//fmt.Printf("\nSentUrl: %v\nUrlData:%v\r\n\r\n", sendURL, URLData)
	response, err := http.PostForm(sendURL, URLData)
	ErrStop(20, err, t)

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	u, _ := url.Parse(sendURL)
	u.RawQuery = URLData.Encode()
	Request := fmt.Sprintf(" RAIDA %d: %v", raidaID, u)
	elapsed := time.Since(start)
	elapsedString := fmt.Sprintf("%v", elapsed)
	responseText[0] = string(body)
	responseText[1] = Request
	responseText[2] = elapsedString
	done <- responseText

} //end Send

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

//fail exits the program with only a error number
func fail(i int) {
	fmt.Printf("{\"status\":\"fail\",\"message\":\"error %d\"}", i)
	os.Exit(i)
}

//ErrStop checks if there is an error, and if so stops the program pringing the error message and the error number
func ErrStop(i int, err error, t time.Time) {
	if err != nil {
		fmt.Printf("{\"status\":\"fail\",\"message\":\"error %d. %s.  %v\"}", i, fmt.Sprintf("%s", err), time.Since(t))
		os.Exit(i)
	}
}

//ParseID parses the supplied id into a usable serial number
func ParseID(parseID string, t time.Time) int {

	if _, err := strconv.Atoi(parseID); err == nil {

		result, _ := strconv.Atoi(parseID)
		return result

	} else {

		addr := net.ParseIP(parseID)
		if addr == nil {
			ips, err := net.LookupIP(parseID)
			ErrStop(12, err, t)
			parseID = strings.Replace(ips[0].String(), "1.", "0.", 1)
			input := net.ParseIP(parseID)
			ipOut := IP4toInt(input, t)
			return int(ipOut)

		} else {
			if strings.HasPrefix(parseID, "1.") {
				parseID = strings.Replace(parseID, "1.", "0.", 1)
				input := net.ParseIP(parseID)
				ipOut := IP4toInt(input, t)

				return int(ipOut)
			}
		}
		return 0
	}
	return 0
}

//convert the Ip to a serial number
func IP4toInt(input net.IP, t time.Time) int64 {
	//ip := net.IP(input)

	//ip, _, err := net.ParseCIDR(input)
	//ErrStop(32, err, t)
	IPv4Int := big.NewInt(0)
	IPv4Int.SetBytes(input.To4())
	return IPv4Int.Int64()
}

func showChange(sn string, denom int, server ...int) [4][]string {
	rand.Seed(time.Now().UTC().UnixNano())
	changeArray := [4][]string{}
	serverArray := [3]int{}

	if len(server) > 0 {
		serverArray[0] = server[0]
		fmt.Println("set server 1 to ", serverArray[0])
	} else {
		serverArray[0] = rand.Intn(24)
	}

	if len(server) > 1 {
		serverArray[1] = server[1]
		fmt.Println("set server 2 to ", serverArray[1])
	} else {
		serverArray[1] = rand.Intn(24)
	}

	if len(server) > 2 {
		serverArray[2] = server[2]
		fmt.Println("set server 3 to ", serverArray[2])
	} else {
		serverArray[2] = rand.Intn(24)
	}

	for serverArray[0] == serverArray[1] {
		serverArray[1] = rand.Intn(24)

		for serverArray[1] == serverArray[2] {
			serverArray[2] = rand.Intn(24)

			for serverArray[2] == serverArray[0] {
				serverArray[0] = rand.Intn(24)
			}

		}

	}

	//	fmt.Printf("%v", serverArray)
	done := make(chan [2]string)
	//var changeURL string

	for i := 0; i < 3; i++ {
		go sendRequest(done, denom, sn, serverArray[i])
		//		fmt.Printf("sent request %d\r\n", i)
	}

	var change []jsonChange

	for i := 0; i < 3; i++ {
		results := <-done
		//fmt.Printf("time Elapsed for request %d: %v\r\n", i+1, results[1])
		change = append(change, readChange(results[0]))

	}

	var d1 []string
	var d5 []string
	var d25 []string
	var d100 []string

	if len(change[0].D1) > 0 || len(change[1].D1) > 0 {
		d1 = intersection(change[0].D1, change[1].D1)
		d1 = intersection(change[2].D1, d1)
	} else {
		d1 = change[0].D1
	}

	if len(change[0].D5) > 0 || len(change[1].D5) > 0 {
		d5 = intersection(change[0].D5, change[1].D5)
		d5 = intersection(change[2].D5, d5)
	} else {
		d5 = change[0].D5
	}

	if len(change[0].D25) > 0 || len(change[1].D25) > 0 {
		d25 = intersection(change[0].D25, change[1].D25)
		d25 = intersection(change[2].D25, d25)
	} else {
		d25 = change[0].D25
	}

	if len(change[0].D100) > 0 || len(change[1].D100) > 0 {
		d100 = intersection(change[0].D100, change[1].D100)
		d100 = intersection(change[2].D100, d100)
	} else {
		d100 = change[0].D100
	}

	changeArray[0] = d1
	changeArray[1] = d5
	changeArray[2] = d25
	changeArray[3] = d100

	return changeArray
}

//send the request for change.
func sendRequest(done chan [2]string, denom int, sn string, serverID int) {
	var responseText [2]string

	changeURL := fmt.Sprintf("https://raida%d.cloudcoin.global/service/show_change?sn=%s&denomination=%d", serverID, sn, denom)

	start := time.Now()
	response, _ := http.Get(changeURL)

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	//	Request := fmt.Sprintf(" RAIDA %d: %v", serverID, changeURL)
	elapsed := time.Since(start)
	elapsedString := fmt.Sprintf("%v", elapsed)
	responseText[0] = string(body)
	responseText[1] = elapsedString
	done <- responseText
}

//read the change and return useable data
func readChange(returnResponse string) jsonChange {
	raw := []byte(returnResponse)
	var change jsonChange
	_ = json.Unmarshal(raw, &change)
	return change
}

//get the intersection of two Arrays
func intersection(a, b []string) (c []string) {
	m := make(map[string]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}
	return
}

func removeIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
