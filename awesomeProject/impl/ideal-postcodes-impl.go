package impl

import (
	"encoding/json"
	"errors"
	"gitlab.okta-solutions.com/mashroom/backend/common/errs"
	"gitlab.okta-solutions.com/mashroom/backend/verification"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	IdealPostcodesApiKeyDefault   = "ak_jncyoqr449EwIiyNK1T3CAJrlspQC" //test api_key
	IdealPostcodesApiKeyEnv       = "IDEAL_POSTCODES_API_KEY"
	IdealPostcodesQueryUrlDefault = "https://api.ideal-postcodes.co.uk/v1/addresses"
	IdealPostcodesQueryUrlEnv     = "IDEAL_POSTCODES_QUERY_URL"
	IdealPostcodesUrlDefault      = "https://api.ideal-postcodes.co.uk/v1/postcodes"
	IdealPostcodesUrlEnv          = "IDEAL_POSTCODES_URL"
	IdealPostcodesTimeoutDefault  = 3000                      //milliseconds
	IdealPostcodesTimeoutEnv      = "IDEAL_POSTCODES_TIMEOUT" //milliseconds
)

var (
	IdealPostcodesApiKey   = IdealPostcodesApiKeyDefault
	IdealPostcodesUrl      = IdealPostcodesUrlDefault
	IdealPostcodesQueryUrl = IdealPostcodesQueryUrlDefault
	IdealPostcodesTimeout  = IdealPostcodesTimeoutDefault
)

func VerifyPostcodeImpl(field *verification.VerifyPostcodeRequest) (*verification.VerifyPostcodeResult, error) {
	postcode := (*field).Postcode
	addr := strings.TrimRight(IdealPostcodesUrl, "/") + "/" + postcode

	request, err := http.NewRequest("GET", addr, nil)
	if (err != nil) {
		return nil, err
	}

	q := request.URL.Query()
	q.Add("api_key", IdealPostcodesApiKey)
	request.URL.RawQuery = q.Encode()

	timeout := time.Duration(time.Duration(IdealPostcodesTimeout) * time.Millisecond)
	client := &http.Client{Timeout: timeout,}

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("provider return error code " + resp.Status)
	}

	/////////////////////////////json parsing
	var city string
	var addresses []string
	err = parseVerifyPostcode([]byte(resp_body), &city, &addresses)
	if err != nil {
		return nil, err
	}

	var result verification.VerifyPostcodeResult
	result.City = city
	result.AddressLine = addresses
	return &result, err
}

func parseVerifyPostcode(bytes []byte, city *string, addresses *[]string) error {
	var result map[string]interface{}
	json.Unmarshal(bytes, &result)

	if (result == nil) {
		return errors.New("no data response") //no data
	}

	code := result["code"].(float64)
	if code != 2000 {
		codeStr := strconv.FormatFloat(code, 'f', 0, 64)
		return errors.New("provider return error code: " + codeStr + " and message: " + result["message"].(string))
	}

	list := result["result"].([]interface{})
	if (list == nil) {
		return errors.New("no data response") //no data
	}
	if (len(list) == 0) {
		return errors.New("no data response") //no data
	}

	firstObj := list[0].(map[string]interface{})
	*city = firstObj["post_town"].(string)
	for _, resp := range list {
		respObj := resp.(map[string]interface{})
		line1 := respObj["line_1"].(string)
		line2 := respObj["line_2"].(string)
		line3 := respObj["line_3"].(string)

		addressLine := strings.TrimSpace(line1) + " " + strings.TrimSpace(line2) + " " + strings.TrimSpace(line3)

		*addresses = append(*addresses, addressLine)
	}

	return nil
}

func VerifyPostcodeQueryImpl(field *verification.VerifyPostcodeQueryRequest) (*verification.VerifyPostcodeResult, error) {
	query := (*field).QueryLine
	if len(query) == 0 {
		return nil, errs.NilRequest
	}

	words := strings.Fields(query)
	if (len(words) < 2) {
		return nil, errors.New("query contains less than 2 words")
	}

	var addresses []string
	var city string
	err := verifyPostcodeQueryProc(query, &city, &addresses)
	if err != nil {
		return nil, err
	}

	var result verification.VerifyPostcodeResult
	result.City = city
	result.AddressLine = addresses
	return &result, err
}

func verifyPostcodeQueryProc(query string, city *string, addresses *[]string) error {
	addr := strings.TrimRight(IdealPostcodesQueryUrl, "/")
	request, err := http.NewRequest("GET", addr, nil)
	if (err != nil) {
		return err
	}

	q := request.URL.Query()
	q.Add("api_key", IdealPostcodesApiKey)
	q.Add("query", query)
	request.URL.RawQuery = q.Encode()

	timeout := time.Duration(time.Duration(IdealPostcodesTimeout) * time.Millisecond)
	client := &http.Client{Timeout: timeout,}
	resp, err := client.Do(request)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New("provider return error code " + resp.Status)
	}

	/////////////////////////////json parsing
	err = parseVerifyPostcodeQuery([]byte(resp_body), city, addresses)
	return err
}

func parseVerifyPostcodeQuery(body []byte, city *string, addresses *[]string) error {
	var data map[string]interface{}
	json.Unmarshal(body, &data)
	if (data == nil) {
		return nil
	}

	code := data["code"].(float64)
	if code != 2000 {
		return errors.New("provider return data with error code " + data["code"].(string))
	}

	result := data["result"].(map[string]interface{})
	list := result["hits"].([]interface{})
	if (list == nil) {
		return nil
	}
	if (len(list) == 0) {
		return nil
	}

	*city = list[0].(map[string]interface{})["post_town"].(string)
	for _, resp := range list {
		respObj := resp.(map[string]interface{})
		town := respObj["post_town"].(string)

		if (*city != town) {
			return errors.New("need to expand the query")
		}

		line1 := respObj["line_1"].(string)
		line2 := respObj["line_2"].(string)
		line3 := respObj["line_3"].(string)

		addressLine := strings.TrimSpace(line1) + " " + strings.TrimSpace(line2) + " " + strings.TrimSpace(line3)
		*addresses = append(*addresses, addressLine)
	}

	return nil
}
