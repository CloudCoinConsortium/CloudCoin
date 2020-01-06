package main

//Transferer
//Transferes coins from the owners account to another remote user
// Version 6/20/2019
// Author: Sean Worthington & Sam Leary
// This code is property of RAIDA Tech.
// No permission to be used outside of RAIDA Tech.

// Sample use
// transferrer.exe -timeout=5 -transactionPath=C:\Bank\Log\transaction -logpath=C:\Bank\Logs -idpath=C:\path\to\ID\coin.stack -receiverID=example.skywallet.cc -fromtag=Envelope -tag=NewTag -amount=100
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

var raidahttp = &http.Client{
	Timeout: time.Duration(10) * time.Second,
}

func main() {

	var err error

	//create all the necessary Variables
	loc := time.FixedZone("UTC", 0) //set the timezone
	now := time.Now().In(loc)       //set the current time
	t := time.Now()                 //set the variable t
	var tagFilter string            //Filter the coins by tag
	var rootpath string             //Root Path
	var receiverID int              //Serial number of the receiver
	var receiverIDs string          //the receiver ID as a string (ip/SN/skywallet)
	var tag string                  //memo of the transaction
	var logpath string              //path to the log folder
	var idpath string               // path to the id Coin
	var JustShow bool               //
	var timeout int                 //how long before the requests time out
	var amount int                  //The number of coins to transfer.
	//seed the Random Servers
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

	//Gather 13 flags
	flag.BoolVar(&JustShow, "justshow", false, "Whether or not to transfer or just show the current skywallets coins (optional)") //optional                                                               //required
	flag.StringVar(&receiverIDs, "receiverID", "", "The receiver ID")                                                             //required                                                   //required
	flag.IntVar(&amount, "amount", 0, "The Number of coins to transfer ")
	flag.StringVar(&tag, "tag", "", "The memo of the coins being sent ")        //required
	flag.StringVar(&logpath, "logpath", "", "The memo of the coins being sent") //required
	flag.StringVar(&rootpath, "transactionPath", "", "A string that shows the path to the transaction log folder ")
	flag.StringVar(&idpath, "idpath", "", "The path to the ID coin")             //required
	flag.StringVar(&tagFilter, "fromtag", ".", "Filter the coins used by tag")   //required
	flag.IntVar(&timeout, "timeout", 5, "How long before a timeout in seconds.") //required

	raidahttp = &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}

	//parse flags
	flag.Parse()

	receiverID = ParseID(receiverIDs, t)
	if receiverID == 0 {

		ErrStop(12, errors.New("invalid receiver ID Supplied"), t)
	}
	total := amount
	if total <= 0 {
		ErrStop(1, errors.New("flag errors: negative or zero coin value requested"), t)
	}
	if tag == "" {
		ErrStop(1, errors.New("flag errors: Missing coin tag "), t)
	}
	if rootpath == "" || logpath == "" || idpath == "" {
		ErrStop(1, errors.New("flag errors: All File Paths not set"), t)
	}
	if tagFilter == "." {
		ErrStop(1, errors.New("flag errors: invalid fromTag"), t)
	}

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

	raidaURL1 := fmt.Sprintf("https://raida%d.cloudcoin.global/service/show?nn=1&sn=%s&an=%s&pan=%s&denomination=1", checkRaida1, idCoin.CloudCoin[0].SN, idCoin.CloudCoin[0].ANs[checkRaida1], idCoin.CloudCoin[0].ANs[checkRaida1])
	raidaURL2 := fmt.Sprintf("https://raida%d.cloudcoin.global/service/show?nn=1&sn=%s&an=%s&pan=%s&denomination=1", checkRaida2, idCoin.CloudCoin[0].SN, idCoin.CloudCoin[0].ANs[checkRaida2], idCoin.CloudCoin[0].ANs[checkRaida2])

	resp1, resp2 := listCoinsInWallet(raidaURL1, raidaURL2, t)

	// if os.Open returns an error then handle it

	// read our opened json as a byte array.
	response1Bytes, err := ioutil.ReadAll(resp1.Body)
	ErrStop(5, err, t)

	response2Bytes, err := ioutil.ReadAll(resp2.Body)
	ErrStop(5, err, t)

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
					ErrStop(7, err, t)
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
						ErrStop(7, err, t)
						worth := denomination(stringsn)
						values = append(values, worth)

					}

				}
			}
		}

	}
	if len(ShowResponses1.MESSAGES) == 0 {
		ErrStop(6, errors.New("No coins in this skywallet"), t)
	}

	if JustShow {
		var firstEntry []string
		listed := ""
		seperator := ","
		for i := 0; i < len(sn); i++ {

			Entry := fmt.Sprintf("%s : %d", sn[i], values[i])
			firstEntry = append(firstEntry, Entry)
			if i == len(ShowResponses1.MESSAGES)-1 {
				seperator = ""
			}
			listed = fmt.Sprintf("%s|%s|%s", listed, firstEntry[i], seperator)
		}
		fmt.Println("Finished Showing Coins")
		fmt.Printf("%s", listed)
		os.Exit(0)

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
		ErrStop(12, err, t)

	}

	/*TRANSFER COINS*/

	var sendURL = [25]string{}
	var URLData = []url.Values{}
	var NN1 = "1"

	toSN := fmt.Sprintf("%d", receiverID)
	for i := 0; i < 25; i++ {

		sendURL[i] = fmt.Sprintf("https://RAIDA%d.cloudcoin.global/service/transfer", i)
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
	var responses []string
	//var byteresponse []byte
	allresponses := ""
	allrequests := ""
	if tagFilter != "" {
		allrequests += "Using tag Filter:\r\n"
	}
	done := make(chan [3]string) // make the "done" channel

	for i := 0; i < 25; i++ {
		responses = append(responses, "no response") //makes two dime array
	} //end of i forloop
	//var passes int

	for i := 0; i < 25; i++ {
		go Transfer(done, sendURL[i], URLData[i], i, t)
	}

	for i := range responses {

		//Go Routine Responses
		results := <-done // receive from the channel
		responses[i] = results[0]
		elapsedTime := results[2]
		allresponses = fmt.Sprintf("%s |%s|%s\r\n\r\n", allresponses, responses[i], elapsedTime)
		//	allrequests = fmt.Sprintf("%s |%s|\r\n\r\n", allrequests, requests[i])

	} // end of i for loop

	responseResults, _ := readResponse(responses)

	goodResp := 0

	for i := range responseResults {
		if responseResults[i] == true {
			goodResp++
		}
	}

	fmt.Println(now.Format("Jan.2.06"))
	fmt.Println(goodResp)
	if goodResp >= 20 {
		path := fmt.Sprintf("%s", logpath)

		err = os.MkdirAll(path, os.ModePerm)
		ErrStop(13, err, t)
		//lpath := fmt.Sprintf("%s\\%s_requests.txt", path, tag)
		path = fmt.Sprintf("%s\\%s.Transfer.txt", path, now.Format("Jan.2.06"))

		WriteToFile(path, fmt.Sprintf("%s %s sent %s %d coins from tag \"%s\" to tag \"%s\"\r\n", now.Format("15:04"), idCoin.CloudCoin[0].SN, receiverIDs, amount, tagFilter, tag), t)
		path = fmt.Sprintf("%s", rootpath)
		err = os.MkdirAll(path, os.ModePerm)
		ErrStop(13, err, t)
		path = fmt.Sprintf("%s\\transactions.txt", path)

		WriteToFile(path, fmt.Sprintf("%s, %s,,%d,%s,%s\r\n", tagFilter, now.Format("2006-Jan-02 15-04 "), amount, "?", tag), t)

		fmt.Printf("{\"status\":\"success\",\"message\":\"Successfully sent %d coins to %s\",\" Execution Time = %s\",\"time\":\"%s\"}", amount, receiverIDs, time.Since(t), now.Format("2006-1-2 15:04:05"))

	} else {
		path := fmt.Sprintf("%s", logpath)

		err = os.MkdirAll(path, os.ModePerm)
		ErrStop(13, err, t)
		//lpath := fmt.Sprintf("%s\\%s_requests.txt", path, tag)
		path = fmt.Sprintf("%s\\%s.Transfer.txt", path, now.Format("Jan.2.06"))

		WriteToFile(path, fmt.Sprintf("%s %s Couldn't sent %s %d coins from tag \"%s\" to tag \"%s\"\r\n", now.Format("Jan"), idCoin.CloudCoin[0].SN, receiverIDs, amount, tagFilter, tag), t)
		path = fmt.Sprintf("%s", rootpath)
		err = os.MkdirAll(path, os.ModePerm)
		ErrStop(13, err, t)
		path = fmt.Sprintf("%s\\transactions.txt", path)
		WriteToFile(path, fmt.Sprintf("%s, %s, ,%d,%s,%s\r\n", tagFilter, now.String(), amount, "total", tag), t)
		fmt.Printf("{\"status\":\"failed\",\"message\":\"Couldn't fully Send\",\" Execution Time = %s\",\"time\":\"%s\"}", time.Since(t), now.Format("2006-1-2 15:04:05"))
	}

	//WriteToFile(lpath, allrequests, t)

} //end main

///

// H E L P E R  F U N C T I O N S

//WriteToFile data to a file
func WriteToFile(filename string, data string, t time.Time) error {

	//file, err := os.Create(filename)
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	ErrStop(14, err, t)

	_, err = f.Write([]byte(data))
	ErrStop(14, err, t)
	err = f.Close()
	ErrStop(15, err, t)
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
	response, err := raidahttp.PostForm(sendURL, URLData)
	if err != nil {
		fmt.Println("request failed for sendURL")
		responseText[0] = "{failed}"
		Request := fmt.Sprintf(" RAIDA %d")
		elapsed := time.Since(start)
		elapsedString := fmt.Sprintf("%v", elapsed)
		responseText[1] = Request
		responseText[2] = elapsedString
		done <- responseText
		return
	}
	//	defer response.Body.Close()
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

func listCoinsInWallet(raida1, raida2 string, t time.Time) (*http.Response, *http.Response) {

	check1, err := raidahttp.Get(raida1)
	ErrStop(4, err, t)

	check2, err := raidahttp.Get(raida2)
	ErrStop(4, err, t)

	return check1, check2
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

func IP4toInt(input net.IP, t time.Time) int64 {
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
