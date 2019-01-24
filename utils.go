package ics

import (

	// "io/ioutil"
	"io/ioutil"
	"strings"
	// "errors"

	"net/http"
	"os"
	"regexp"
	"sync"
)

var o sync.Once
var mutex *sync.Mutex
var idCounter int

// if RepeatRuleApply is true , the rrule will create new objects for the repeated events
var RepeatRuleApply bool

// max of the rrule repeat for single event
var MaxRepeats int

//  unixtimestamp
const uts = "1136239445"

//ics date time format
const IcsFormat = "20060102T150405Z"

// Y-m-d H:i:S time format
const YmdHis = "2006-01-02 15:04:05"

// ics date format ( describes a whole day)
const IcsFormatWholeDay = "20060102"

// downloads the calendar before parsing it
func downloadFromUrl(url string) ([]byte, error) {

	// get the URL
	response, err := http.Get(url)

	if err != nil {
		return []byte{}, err
	}
	// close the response body
	defer response.Body.Close()

	// copy response body to []byte
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []byte{}, err
	}

	//return the bytes that contains the info
	return body, nil
}

func stringToByte(str string) []byte {
	return []byte(str)
}

// removes newlines and cutset from given string
func trimField(field, cutset string) string {
	re, _ := regexp.Compile(cutset)
	cutsetRem := re.ReplaceAllString(field, "")
	return strings.TrimRight(cutsetRem, "\r\n")
}

//  checks if file exists
func fileExists(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil
}

func parseDayNameToIcsName(day string) string {
	var dow string
	switch day {
	case "Mon":
		dow = "MO"
		break
	case "Tue":
		dow = "TU"
		break
	case "Wed":
		dow = "WE"
		break
	case "Thu":
		dow = "TH"
		break
	case "Fri":
		dow = "FR"
		break
	case "Sat":
		dow = "ST"
		break
	case "Sun":
		dow = "SU"
		break
	default:
		// fmt.Println("DEFAULT :", start.Format("Mon"))
		dow = ""
		break
	}
	return dow
}
