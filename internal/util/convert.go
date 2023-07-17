package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httputil"
	"regexp"
	"time"

	// "campaign-service/internal/entity"
	"strconv"
)

//ConvertToReader ...
func ConvertToReader(data interface{}) (*bytes.Reader, error) {
	dataBytes, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	return bytes.NewReader(dataBytes), nil
}

//ConvertJSONByteToStruct ...
func ConvertJSONByteToStruct(data []byte, target interface{}) error {
	return json.Unmarshal(data, target)
}

//ConvertStructToJSONByte ...
func ConvertStructToJSONByte(data interface{}) ([]byte, error) {
	dataBytes, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	return dataBytes, nil
}

//ConvertStructToJSONString ...
func ConvertStructToJSONString(data interface{}) string {
	strJson, err := json.Marshal(data)
	if err != nil {
		return err.Error()
	}

	return string(strJson)
}

//ConvertRequestToString ...
func ConvertRequestToString(req *http.Request) (string, error) {
	data, err := httputil.DumpRequest(req, true)

	if err != nil {
		return "", err
	}

	return string(data), nil
}

//ConvertResponseToString ...
func ConvertResponseToString(resp *http.Response) (string, error) {
	data, err := httputil.DumpResponse(resp, true)

	if err != nil {
		return "", err
	}

	return string(data), nil
}

var (
	Struct2ProtoModelErrDestType = errors.New("destination wrong type")
	Struct2ProtoModelTitle       = "json="
	Struct2ProtoModelTitle2      = "name="
	Struct2ProtoModelRegex, _    = regexp.Compile(Struct2ProtoModelTitle + `\w+`)
	Struct2ProtoModelRegex2, _   = regexp.Compile(Struct2ProtoModelTitle2 + `\w+`)
)

//ConvertStringToInt ...
func ConvertStringToInt(strString string) (int, error) {
	intVar, err := strconv.Atoi(strString)
	if err != nil {
		return 0, err
	}

	return intVar, nil
}

//ConvertIntToString ...
func ConvertIntToString(strInt int) string {
	strVar := strconv.Itoa(strInt)
	return strVar
}

//UniqueInvestmentPlanList ...
// func UniqueInvestmentPlanList(investmentPlanList []*entity.InvestmentPlanList) []*entity.InvestmentPlanList {
// 	unique := make([]*entity.InvestmentPlanList, 0)
// loop:
// 	for _, v := range investmentPlanList {
// 		for i, u := range unique {
// 			if v.ID == u.ID {
// 				unique[i] = v
// 				continue loop
// 			}
// 		}
// 		unique = append(unique, v)
// 	}
// 	return unique
// }
//FindMinAndMax...
func FindMinAndMax(a []int) (min int, max int) {
	if len(a) <= 0 {
		return 0, 0
	}
	min = a[0]
	max = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}

//contains
func CheckContains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

//ConvertPointerToString ...
func ConvertPointerToString(str *string) string {
	if str == nil {
		return ""
	}

	return *str
}

//ConvertStringToPointer ...
func ConvertStringToPointer(str string) *string {
	if str == "" {
		return nil
	}

	return &str
}

//ConvertPointerToFloat64 ...
func ConvertPointerToFloat64(float *float64) float64 {
	if float == nil {
		return 0.0
	}

	return *float
}

//ConvertFloat64ToPointer ...
func ConvertFloat64ToPointer(float float64) *float64 {
	if float == 0 {
		return nil
	}

	return &float
}

//FormatAccountNumber
func FormatAccountNumber(isHidden bool, accNo string) string {
	newAccNo := "-"
	if len(accNo) >= 10 {
		if isHidden {
			return fmt.Sprintf("XXX-X-%v-X", accNo[4:9])
		} else {
			//show last 4 digits
			return fmt.Sprintf("%v", accNo[6:10])
		}
	}
	return newAccNo
}

//ConvertToMonthPeriod
func ConvertToMonthPeriod(openDate string, date string) int {
	totalYear := 0
	totalMonth := 0
	totalDate := 0

	current := openDate
	currentDate, _ := strconv.Atoi(current[6:8])
	currentMonth, _ := strconv.Atoi(current[4:6])
	currentYear, _ := strconv.Atoi(current[0:4])

	day, _ := strconv.Atoi(date[6:8])
	month, _ := strconv.Atoi(date[4:6])
	year, _ := strconv.Atoi(date[0:4])
	if currentYear > year {
		totalYear = (currentYear - year) * 12
		if currentMonth < month {
			totalYear = totalYear - (month - currentMonth)
		}
	}

	if currentMonth > month {
		totalMonth = currentMonth - month
	}
	if currentDate < day {
		totalDate = totalDate - 1
	}

	monthPeriod := totalMonth + totalYear + totalDate
	if monthPeriod < 0 {
		monthPeriod = 0
	}
	return monthPeriod
}

//ConvertStringToDate ...
func ConvertStringToDate(date string) time.Time {
	t, _ := time.Parse("20060102", date)
	return t
}

//ConvertDateToString ...
func ConvertDateToString(date time.Time) string {
	strDate := date.Format("20060102")
	return strDate
}

//ConvertPointerInt32ToInt ...
func ConvertPointerInt32ToInt(Int32Value *int32) int {
	if Int32Value == nil {
		return 0
	} else {
		sampleString := fmt.Sprint(*Int32Value)
		if convertToInt, err := strconv.Atoi(sampleString); err == nil {
			return convertToInt
		} else {
			return 0
		}
	}
}
