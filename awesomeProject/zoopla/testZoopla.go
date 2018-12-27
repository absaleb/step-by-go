package zoopla

import (
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	webClientTimeout       = 3000
	ZooplaClientTimeout    = 3000
	ZooplaApiKey           = "qh46j22xzsx79byk84vb9tms"
	ZooplaListingUrl       = "https://realtime-listings-api.webservices.zpg.co.uk/sandbox/v1/"
	ZooplaListingHeaderUrl = "https://realtime-listings.webservices.zpg.co.uk/docs/v1.2/schemas/"
)

type ZooplaMethod int

const (
	Branch_update  ZooplaMethod = 0
	Listing_update ZooplaMethod = 1
	Listing_delete ZooplaMethod = 2
	Listing_list   ZooplaMethod = 3
)

func (z ZooplaMethod) String() string {
	names := [...]string{
		"branch/update",
		"listing/update",
		"listing/delete",
		"listing/list"}

	if z < 0 || z > 3 {
		return ""
	}

	return names[z]
}

type ZooplaArea struct {
	NumberOfSales7Year    string `json:"number_of_sales_7year"`
	AverageSoldPrice7Year string `json:"average_sold_price_7year"`
	NumberOfSales5Year    string `json:"number_of_sales_5year"`
	NumberOfSales3Year    string `json:"number_of_sales_3year"`
	AverageSoldPrice1Year string `json:"average_sold_price_1year"`
	NumberOfSales1Year    string `json:"number_of_sales_1year"`
	Turnover              string `json:"turnover"`
	PricesURL             string `json:"prices_url"`
	AverageSoldPrice3Year string `json:"average_sold_price_3year"`
	AverageSoldPrice5Year string `json:"average_sold_price_5year"`
}

type ZooplaAverageAreaSoldPrice struct {
	Country     string       `json:"country"`
	ResultCount string       `json:"result_count"`
	Longitude   float64      `json:"longitude"`
	AreaName    string       `json:"area_name"`
	Street      string       `json:"street"`
	Town        string       `json:"town"`
	Latitude    float64      `json:"latitude"`
	County      string       `json:"county"`
	Areas       []ZooplaArea `json:"areas"`
	BoundingBox struct {
		LongitudeMin string `json:"longitude_min"`
		LatitudeMin  string `json:"latitude_min"`
		LongitudeMax string `json:"longitude_max"`
		LatitudeMax  string `json:"latitude_max"`
	} `json:"bounding_box"`
	Postcode string `json:"postcode"`
}

func (z ZooplaAverageAreaSoldPrice) String() string {
	return fmt.Sprintf("Country:%s, AreaName:%s, Street:%s, Town:%s, County:%s,\n Areas count:%d", z.Country, z.AreaName, z.Street, z.Town, z.County, len(z.Areas))
}

func GetAverageAreaSoldPrice() (string, error) {
	addr := "https://api.zoopla.co.uk/api/v1/average_area_sold_price.js"
	request, err := http.NewRequest("GET", addr, nil)
	if err != nil {
		return "", err
	}

	q := request.URL.Query()

	q.Add("postcode", "w12")
	q.Add("output_type", "county")
	q.Add("area_type", "streets")
	q.Add("api_key", ZooplaApiKey)

	request.URL.RawQuery = q.Encode()

	bytes, err := getBytes(request)
	if err != nil {
		return "", err
	}

	// api response json parsing
	var data ZooplaAverageAreaSoldPrice
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return "", err
	}

	return data.String(), nil
}

func getBytes(request *http.Request) ([]byte, error) {
	timeout := time.Duration(time.Duration(webClientTimeout) * time.Millisecond)
	client := &http.Client{Timeout: timeout}
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

	return []byte(resp_body), nil
}

func sendBytes() error {
	return nil
}

func SendListing(method ZooplaMethod, data []byte) error {
	addr := fmt.Sprintf("%s%s.json", ZooplaListingUrl, method)

	req, err := http.NewRequest("POST", addr, bytes.NewBuffer([]byte(data)))
	if err != nil {
		return err
	}

	headerValue := fmt.Sprintf("application/json; profile=%s%s.json", ZooplaListingHeaderUrl, method)
	req.Header.Set("Content-Type", headerValue)

	if method == Listing_update {
		etagValue := getEtagValue(data)
		req.Header.Set("ZPG-Listing-ETag", etagValue)
	}

	client := &http.Client{}
	client.Timeout = ZooplaClientTimeout * time.Millisecond

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))

	return nil
}

func getEtagValue(data []byte) string {
	h := sha1.New()
	h.Write(data)
	sha := base64.URLEncoding.EncodeToString(h.Sum(nil))
	return sha
}

func getJSON(method ZooplaMethod) *string {
	switch method {
	case Branch_update:
		return getBranchUpdateJSON()
	case Listing_delete:
		return getListingListJSON()
	case Listing_delete:
		return getListingListJSON()
	default:
		return nil

	}
}

func getBranchUpdateJSON() *string {
	location := Location{
		PropertyNumberOrName: PROPERTY_NAME,
		StreetName:           STREET_NAME,
		TownOrCity:           TOWN,
		PostalCode:           POSTCODE,
		CountryCode:          COUNTRY_CODE}

	jsn := BranchUpdate{
		BranchReference: BRANCH_REF,
		BranchName:      BRANCH_NAME,
		Email:           EMAIL,
		Website:         WEBSITE,
		Location:        &location}

	b, err := json.Marshal(jsn)
	if err != nil {
		return nil
	}

	result := string(b)
	return &result
}

func getListingListJSON() *string {
	jsn := ListingList{
		BranchReference: BRANCH_REF}

	b, err := json.Marshal(jsn)
	if err != nil {
		return nil
	}

	result := string(b)
	return &result
}
