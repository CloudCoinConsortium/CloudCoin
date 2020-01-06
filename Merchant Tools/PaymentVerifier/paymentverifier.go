//paymentverifier.go or changeverifier.go

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
	"strconv"
)

// Author: Samuel Leary & Sean Worthington
//Version 8-12-19
// This code is property of RAIDA Tech.
// No permission to be used outside of RAIDA Tech.
// Sample use
//C:\CloudCoin\paymentVerifier.exe -timeout=5 -getfrom=sean4 -payment=100 -refundto=1371486 -rootpath="C:\CloudCoin\Accounts\Change" -idpath="C:\CloudCoin\Accounts\Change\ID\1.CloudCoin.1.2..stack"
// EXIT CODES:
// 1  error with command line argument 3 cannot split 1 coins
// 2  error with command line argument 3 invalid amount of coins
// 3  error cannot open Keys file
// 4  error Cannot read Keys File
// 5 error cannot unmarshal json
// 6  error cannot create log
// 7  error cannot create log request
// 12 error Not Enough Good Replies
// 15 error with command line flags: missing flags

//CloudCoin struct for type cloudcoin
type CloudCoin struct {
	NN  string   `json:"nn"`
	SN  string   `json:"sn"`
	ANs []string `json:"an"`
}

//Stack Creates a stack of CloudCoin
type Stack struct {
	CloudCoin []CloudCoin `json:"cloudcoin"`
}



var Receive string
var totalCoinsSent string
var refundTo string
var rootPath string
var idPath string
var timeout string
var newtag string



func init() {
	//get argument flags
	flag.StringVar(&Receive, "getfrom", "", "The receiver envelope id")
	flag.StringVar(&totalCoinsSent, "payment", "", "A total amount of the coins sent")
	flag.StringVar(&refundTo, "refundto", "", "The account that would receive a refund")
	flag.StringVar(&rootPath, "rootpath", "", "The path to the Root Directory")
	flag.StringVar(&idPath, "idpath", "", "Path to the ID Coin")
	flag.StringVar(&timeout, "timeout", "", "Time allowed to call on the RAIDA.")
	flag.StringVar(&newtag, "newtag", "", "The envelope that the sent money will be moved to.")
}



func main() {
	loc := time.FixedZone("UTC", 0)
	now := time.Now().In(loc)
	t := time.Now()


	flag.Parse()

	//set up default variables
	var err error
	goodReplies := 0

	//open File and use its data to create URLKeys array
	var cloudcoin Stack
	//fmt.Println(idPath)

	timeout_int, err := strconv.Atoi(timeout)
	ErrStop(25, err, t)

	jsonfile, err := os.Open(idPath)
	ErrStop(3, err, t)

	byteValue, _ := ioutil.ReadAll(jsonfile)
	ErrStop(4, err, t)

	err = json.Unmarshal(byteValue, &cloudcoin)
	ErrStop(5, err, t)

	urlKeys := make([]string, 25)
	for i := 0; i < 25; i++ {
		currentKey := cloudcoin.CloudCoin[0].ANs[i]
		urlKeys[i] = currentKey
	}
 var nn = cloudcoin.CloudCoin[0].NN
 var sn = cloudcoin.CloudCoin[0].SN
	//create channels to receive responses
	done := make(chan string)
	requests := make(chan string)

	//create all 25 requests
	for i := 0; i < 25; i++ {
		//create go routines to Send Url Data
		go SendURL(nn, sn, timeout_int, done, requests, refundTo, urlKeys, totalCoinsSent, Receive, newtag, i, t)
	}
	logRequest := ""
	logResponse := ""
	returnRequest := make([]string, 0)
	returnResponse := make([]string, 0)
	//retrieve all necessary responses
	for i := 0; i < 25; i++ {
		//receive responses
		returnResponse = append(returnResponse, <-done)
		returnRequest = append(returnRequest, <-requests)
		if i == 0 {
			logResponse = fmt.Sprintf("|%s|", returnResponse[0])
			logRequest = fmt.Sprintf("|%s|", returnRequest[0])

		} else {
			logResponse = fmt.Sprintf("%s\r\n|%s|", logResponse, returnResponse[i])
			logRequest = fmt.Sprintf("%s\r\n|%s|", logRequest, returnRequest[i])

		}
	} // end for loop

	err = WriteToLog(rootPath, "response.txt", logResponse, Receive)
	ErrStop(6, err, t)

	err = WriteToLog(rootPath, "request.txt", logRequest, Receive)
	ErrStop(7, err, t)

	for i := 0; i < len(returnResponse); i++ {
		stringreader := strings.FieldsFunc(returnResponse[i], func(r rune) bool {
			if r == ':' || r == ',' {
				return true
			}
			return false
		})

		if 3 < len(stringreader) {
			fmt.Println(stringreader[3] +" " + stringreader[6])

			if stringreader[3] == "\"pass\"" {
				goodReplies++
			}
		} else {
			fmt.Println("outofrange")
		}
	}
	//if there is enough good responses print success and end
	if goodReplies > 19 {
		//print success
		fmt.Printf("{\"status\":\"success\",\"message\":\"%d good replies. Execution Time = %s\",\"time\":\"%s\"}", goodReplies, time.Since(t), now.Format("2006-1-2 15:04:05"))

	} else {
	//	fmt.Println(goodReplies)
		fmt.Printf("{\"server\":\"Change\",\"status\":\"fail\",\"message\":\"Could not make verifie payment. Did not Receive Coins. Execution Time  = %s\",\"time\":\"%s\"}", time.Since(t), now.Format("2006-1-2 15:04:05"))
		os.Exit(12)
	}

}

//ErrStop takes an error, and if it is not null will exit the program
func ErrStop(num int, err error, t time.Time) {
	if err != nil {
		Newerr := fmt.Sprintf("%v", err)
		fmt.Printf("{\"status\":\"fail\",\"message\":\"error %d. %s.  %v\"}", num, Newerr, time.Since(t))
		os.Exit(num)
	}
}

//SendURL sends data to the corresponding url, and sends back a response.
func SendURL(nn string, sn string, timeout int, done chan string, request chan string, refund string, keys []string, coinsSent string, tag string, newtag string, index int, t time.Time) {

	sendURL := fmt.Sprintf("https://RAIDA%d.cloudcoin.global/service/rename_tag", index)
	URLData := url.Values{}
	URLData.Set("nn", nn)
	URLData.Add("sn", sn)
	URLData.Add("an", keys[index])
	URLData.Add("pan", keys[index])
	intSN, _ := strconv.Atoi(sn)
	intD := Denomination( intSN )
	URLData.Add("denomination", strconv.Itoa(intD) )
	URLData.Add("if_total", coinsSent)
	URLData.Add("tag", tag)
	URLData.Add("if_total_wrong_return_to_number", refund)
	URLData.Add("new_tag", newtag)

	u, _ := url.Parse(sendURL)
	u.RawQuery = URLData.Encode()
	Request := fmt.Sprintf("%v", u)
	body :="RAIDA " + strconv.Itoa(index) + " timed out." //set it as a fail to beging with
	//start := time.Now()

	//use get to recieve response from RAIDA
	//tout := int64(timeout)
	var client = &http.Client{
	  Timeout: time.Duration(timeout) * time.Second,
	}
  response, err := 	client.Get(Request)
	if err == nil {
		bodybytes, _ :=ioutil.ReadAll(response.Body)
		body  = string(bodybytes)
	}// end if
//	response, err := http.Get(Request)
	//ErrStop(3, err, t)
	//defer response.Body.Close()

	//print send times to console
	//elapsed := time.Since(start)
	//fmt.Printf(" Sending took %s\n", elapsed)
	var responseText = fmt.Sprintf("%s\r\n", string(body))
	done <- responseText
	request <- Request
}//end send url

//WriteToLog Writes the text to a file
func WriteToLog(Path string, filepath string, text string, tag string) error {
	logPath := fmt.Sprintf("%s\\Logs\\%s", Path, tag)
	err := os.MkdirAll(logPath, 0666)
	logPath = fmt.Sprintf("%s\\%s_%s", logPath, tag, filepath)
	err = ioutil.WriteFile(logPath, []byte(text), 0666)
	return err
}
//Denomination determines the denomination of the coin
func Denomination(sn int) int {
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
