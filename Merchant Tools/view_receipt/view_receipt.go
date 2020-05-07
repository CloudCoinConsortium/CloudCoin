package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

// S T R U C T U R E S
//jsonResponse is the struct that defines the json response from the raida
type jsonResponse struct {
	Server        string  `json:"server"`
	TotalReceived int     `json:"total_received"`
	SerialNumbers string  `json:"serial_numbers"`
	Version       string  `json:"version"`
	Time          string  `json:"time"`
	ExecutionTime float64 `json:"execution_time"`
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

var expAmount int
var logpath string
var timeout int
var guid string
var account string

var raidahttp = &http.Client{
	Timeout: time.Duration(10) * time.Second,
}

func main() {

	t := time.Now() //set the time variable t
	//get flags
	flag.IntVar(&expAmount, "expected_amount", 0, "The amount of coins to be"+
		"transferred to the account")
	flag.StringVar(&logpath, "logpath", "", "The path to the folder to stor the"+
		" receipt")
	flag.IntVar(&timeout, "timeout", 10, "how many seconds will wait before timing out")
	flag.StringVar(&guid, "guidtag", "default", "the guid of the included payment")
	flag.StringVar(&account, "account", "default", "the id of the receiving account.")
	flag.Parse()
	//convert output to correct user id
	accountID := parseID(account, t)

	done := make(chan string)
	//create and send Requests
	for i := 0; i < 25; i++ {
		go SendURL(i, accountID, guid, done)
	}

	returnResponse := make([]string, 25)
	jsonResponse := make([]jsonResponse, 25)
	pass := 0
	pick := 0

	for i := 0; i < 25; i++ {
		//receive finished responses from done channel
		returnResponse[i] = <-done
		//fmt.Println(returnResponse[i])
		//confirm that the data is a success
		if !strings.Contains(returnResponse[i], "timed out") {
			jsonResponse[i] = convertToJSON(returnResponse[i])
			if jsonResponse[i].SerialNumbers != "" {
				if jsonResponse[i].TotalReceived == expAmount {
					pass++
					pick = i

					//fmt.Println(pass)
				}
			}
		}

	}
	time := t.Format("15:04:05")
	logText := fmt.Sprintf("start responses %s {", time)
	if pass >= 20 {
		fmt.Println(returnResponse[pick])

		//create log text
		for i := 0; i < len(returnResponse); i++ {
			logText = fmt.Sprintf("%s\r\n %s", logText, returnResponse[i])
		}

	} else {
		logText = fmt.Sprintf("%s\r\n Receipt Could not be verified\r\n", logText)
	}
	logText = fmt.Sprintf("%s\r\n} end responses %s\r\n", logText, time)
	//right logtext to filepath
	writeLog(logText, t)
}

//parses a skywallet id, or ip address and converts to a useable user id.
func parseID(parseID string, t time.Time) int {

	if _, err := strconv.Atoi(parseID); err == nil {

		result, _ := strconv.Atoi(parseID)
		return result

	} else {

		addr := net.ParseIP(parseID)
		if addr == nil {
			ips, err := net.LookupIP(parseID)
			errStop(12, err, t)
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

func errStop(i int, err error, t time.Time) {
	if err != nil {
		fmt.Printf("{\"status\":\"fail\",\"message\":\"error %d. %s.  %v\"}", i, fmt.Sprintf("%s", err), time.Since(t))
		os.Exit(i)
	}
}

//SendURL sends data to the corresponding url, and sends back a response.
func SendURL(index, account int, tag string, done chan string) {
	sendURL := fmt.Sprintf("https://RAIDA%d.cloudcoin.global/service/view_receipt", index)

	//fmt.Println(sendURL)

	//create get request
	URLData := url.Values{}
	var raidahttp = &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
	URLData.Set("account", strconv.Itoa(account))
	URLData.Set("tag", tag)

	u, _ := url.Parse(sendURL)
	u.RawQuery = URLData.Encode()
	Request := fmt.Sprintf("%v", u)
	//create timeout response
	body := "RAIDA " + strconv.Itoa(index) + " timed out." //set it as a fail to begin with

	//send request
	response, err := raidahttp.Get(Request)
	if err == nil {
		//capture results
		bodybytes, _ := ioutil.ReadAll(response.Body)
		body = string(bodybytes)
	} // end if

	//format results
	var responseText = fmt.Sprintf("%s\r\n", string(body))
	//send results to done channel
	done <- responseText

} //end send url

//read the change and return useable data
func convertToJSON(returnResponse string) jsonResponse {
	raw := []byte(returnResponse)
	var response jsonResponse
	_ = json.Unmarshal(raw, &response)
	return response
}

func writeLog(text string, t time.Time) {
	//convert date and time to proper format
	time := t.Format("Jan.2.2006")
	filename := fmt.Sprintf("%s\\%s.view_receipt.txt", logpath, time)
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := f.Write([]byte(text)); err != nil {
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
