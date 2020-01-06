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

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

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
	flag.IntVar(&receiverID, "receiversn", -1, "The receiver envelope id")                                                        //required
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
	checkServer := fmt.Sprintf("https://Raida%d.cloudcoin.global", receiverID)

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

	fmt.Println("show URL: " + showURL + "\n")

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
	fmt.Println(tagFilter)

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
	var snToSend = []string{}
	var denominations = []int{}
	for i, d := range values {
		switch d {
		case 1:
			if intOnes > 0 { //if a one note is needed
				snToSend = append(snToSend, sn[i])
				denominations = append(denominations, 1)
				intOnes-- //reduce the ones neeeded by one
			} //end if needed greater than zero
			break
		case 5:
			if intFives > 0 { //if a one note is needed
				snToSend = append(snToSend, sn[i])
				denominations = append(denominations, 5)
				intFives-- //reduce the ones neeeded by one
			} //end if needed greater than zero
			break
		case 25:
			if intTwentyfives > 0 { //if a one note is needed
				snToSend = append(snToSend, sn[i])
				denominations = append(denominations, 25)
				intTwentyfives-- //reduce the ones neeeded by one
			} //end if needed greater than zero
			break
		case 100:
			if intHundreds > 0 { //if a one note is needed
				snToSend = append(snToSend, sn[i])
				denominations = append(denominations, 100)
				intHundreds-- //reduce the ones neeeded by one
			} //end if needed greater than zero
			break
		case 250:
			if intTwohundredfifties > 0 { //if a one note is needed
				snToSend = append(snToSend, sn[i])
				denominations = append(denominations, 250)
				intTwohundredfifties-- //reduce the ones neeeded by one
			} //end if needed greater than zero
			break
		} // end switch
		// Do we have al notes we need?
		if intOnes == 0 && intFives == 0 && intTwentyfives == 0 && intHundreds == 0 && intTwohundredfifties == 0 {
			break
		} //End if all needs met

	} //end for each file in the folder
	if intOnes > 0 {
		message3 := "There were not enough ones to fill the request."
		err := fmt.Errorf("error with inventory: %v", message3)
		ErrStop(8, err, t)
	} //End if needs not met

	if intFives > 0 {
		message4 := "There were not enough fives to fill the request."
		err := fmt.Errorf("error with inventory: %v", message4)
		ErrStop(9, err, t)
	} //End if needs not met

	if intTwentyfives > 0 {
		message5 := "There were not enough 25s to fill the request."
		err := fmt.Errorf("error with inventory: %v", message5)
		ErrStop(10, err, t)
	} //End if needs not met

	if intHundreds > 0 {
		message6 := "There were not enough 100s to fill the request."
		err := fmt.Errorf("error with inventory: %v", message6)
		ErrStop(11, err, t)
	} //End if needs not met

	if intTwohundredfifties > 0 {
		message7 := "There were not enough 250s to fill the request."
		err := fmt.Errorf("error with inventory: %v", message7)
		ErrStop(12, err, t)
	} //End if needs not met

	//Now we have a list of coins that we can send. Read, send and delete.

	/*TRANSFER COINS*/

	var sendURL = [25]string{}
	var URLData = []url.Values{}
	var NN1 = "1"

	toSN := fmt.Sprintf("%d", receiverID)
	for i := 0; i < 25; i++ {

		sendURL[i] = fmt.Sprintf("https://RAIDA%d.cloudcoin.global/service/transfer_with_change", i)
		URLData = append(URLData, url.Values{})
		URLData[i].Set("tag", tag)
		URLData[i].Set("nn", NN1)
		URLData[i].Set("sn", idCoin.CloudCoin[0].SN)
		URLData[i].Set("an", idCoin.CloudCoin[0].ANs[i])
		URLData[i].Set("pan", idCoin.CloudCoin[0].ANs[i])
		URLData[i].Set("to_sn", toSN)
		URLData[i].Set("denomination", "1")

		for j := 0; j < len(snToSend); j++ {
			if j == 0 {
				URLData[i].Set("sns[]", snToSend[j])
				URLData[i].Set("nns[]", "1")
			} else {
				URLData[i].Add("sns[]", snToSend[j])
				URLData[i].Add("nns[]", "1")
			}
		} //end for each file

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
	//	fmt.Printf("\nSentUrl: %v\nUrlData:%v", sendURL, URLData)
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
