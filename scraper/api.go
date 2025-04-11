package scraper

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
)

// Remove Illegal Characters from search term.
func Sanitise(unsanitised string) string {
	sanitisationpattern := "[^a-zA-Z가-힣]"
	re := regexp.MustCompile(sanitisationpattern)
	sanitised := re.ReplaceAllString(unsanitised, "")
	return sanitised
}

// Fetch JSON Data from URL.
func Fetch(url string) (map[string]interface{}, error) {
	// Create HTTP Request
	request, errorreq := http.NewRequest(http.MethodGet, url, nil)
	if errorreq != nil {
		msg := fmt.Sprintf("cannot create HTTP request: %v", errorreq)
		return nil, errors.New(msg)
	}
	request.Header.Add("Referer", "https://korean.dict.naver.com/koendict/") // Trick to get JSON Data.

	// Send HTTP Request.
	client := &http.Client{}
	resp, errordo := client.Do(request)
	if errordo != nil {
		msg := fmt.Sprintf("cannot fetch URL %q: %v", url, errordo)
		return nil, errors.New(msg)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		msg := fmt.Sprintf("unexpected http GET status: %s", resp.Status)
		return nil, errors.New(msg)
	}

	// Decode JSON Data.
	var result map[string]interface{}
	errordecode := json.NewDecoder(resp.Body).Decode(&result)
	if errordecode != nil {
		msg := fmt.Sprintf("cannot decode JSON: %v", errordecode)
		return nil, errors.New(msg)
	}

	return result, nil
}

// Format Search Term into Naver Dictionary Entry Url.
func GetEntryUrl(searchterm string) (string, error) {
	hostname := "https://korean.dict.naver.com"
	// Note: Korean Search Term MUST be utf-8 encoded!
	Url, err := url.Parse(hostname)
	if err != nil {
		msg := fmt.Sprintf("cannot parse URL: %v", err)
		return "", errors.New(msg)
	}

	Url.Path += "/api3/koen/search"
	parameters := url.Values{}
	parameters.Add("query", searchterm)
	parameters.Add("m", "mobile")
	parameters.Add("range", "entrySearch")
	Url.RawQuery = parameters.Encode()

	entryurl := Url.String()
	return entryurl, nil
}

// Get Entry Information from Naver Dictionary
func GetEntryInfo(searchterm string) (map[string]interface{}, error) {
	entryurl, errentryurl := GetEntryUrl(searchterm)
	if errentryurl != nil {
		return nil, errentryurl
	}
	entryinfo, errentryinfo := Fetch(entryurl)
	if errentryinfo != nil {
		return nil, errentryinfo
	}
	return entryinfo, nil
}

// Get the Entry ID of the Search Term from the Entry Information.
func GetEntryId(entryinfo map[string]interface{}) (string, error) {
	// Check Entry
	// Equivalent to searchInfo.searchResultMap.searchResultListMap.WORD.items[0].entryId;
	searchresultmap, errorsearchresultmap := entryinfo["searchResultMap"].(map[string]interface{})
	if !errorsearchresultmap {
		return "", errors.New("cannot find searchResultMap in entryinfo")
	}

	searchresultlistmap, errorsearchresultlistmap := searchresultmap["searchResultListMap"].(map[string]interface{})
	if !errorsearchresultlistmap {
		return "", errors.New("cannot find searchResultListMap in searchresultmap")
	}

	word, errorword := searchresultlistmap["WORD"].(map[string]interface{})
	if !errorword {
		return "", errors.New("cannot find WORD in searchresultlistmap")
	}

	items, erroritems := word["items"].([]interface{})
	if !erroritems || len(items) == 0 {
		return "", errors.New("cannot find items in word")
	}

	item, erroritem := items[0].(map[string]interface{})
	if !erroritem {
		return "", errors.New("cannot find items in word")
	}

	entryid, errorentryid := item["entryId"].(string)
	if !errorentryid {
		return "", errors.New("cannot find entryId in items")
	}

	return entryid, nil
}

// Format Entry Id into Naver Dictionary Search Url.
func GetSearchUrl(entryid string) (string, error) {
	hostname := "https://korean.dict.naver.com"
	// Note: Korean Search Term MUST be utf-8 encoded!
	Url, err := url.Parse(hostname)
	if err != nil {
		msg := fmt.Sprintf("cannot parse URL: %v", err)
		return "", errors.New(msg)
	}

	Url.Path += "/api/platform/koen/entry"
	parameters := url.Values{}
	parameters.Add("entryId", entryid)
	Url.RawQuery = parameters.Encode()

	searchurl := Url.String()
	return searchurl, nil
}

// Get Search Information from Naver Dictionary
func GetSearchInfo(entryid string) (map[string]interface{}, error) {
	searchurl, errsearchurl := GetSearchUrl(entryid)
	if errsearchurl != nil {
		return nil, errsearchurl
	}
	searchinfo, errsearchinfo := Fetch(searchurl)
	if errsearchinfo != nil {
		return nil, errsearchinfo
	}
	return searchinfo, nil
}

// Scrape Entry Information from Naver Dictionary. (Public API)
func GetEntryInfoRaw(searchterm string) (map[string]interface{}, error) {
	sanitised := Sanitise(searchterm)
	if sanitised == "" {
		welcomemsg := "Welcome to NaverDict Bot! Please enter a Korean word to search (e.g. 나무)."
		return nil, errors.New(welcomemsg)
	}
	entryinfo, errentryinfo := GetEntryInfo(sanitised)
	if errentryinfo != nil {
		return nil, errentryinfo
	}
	return entryinfo, nil
}

// Scrape Search Information from Naver Dictionary. (Public API)
func GetSearchInfoRaw(searchterm string) (map[string]interface{}, error) {
	entryinfo, errentryinfo := GetEntryInfoRaw(searchterm)
	if errentryinfo != nil {
		return nil, errentryinfo
	}
	entryid, errentryid := GetEntryId(entryinfo)
	if errentryid != nil {
		return nil, errentryid
	}
	searchinfo, errsearchinfo := GetSearchInfo(entryid)
	if errsearchinfo != nil {
		return nil, errsearchinfo
	}
	return searchinfo, nil
}

// Scrape Naver Dictionary from a Search Term. (Public API)
func Get(searchterm string) (DictInfo, error) {
	searchinfo, errsearchinfo := GetSearchInfoRaw(searchterm)
	if errsearchinfo != nil {
		return DictInfo{}, errsearchinfo
	}
	dictinfo, errscrape := Scrape(searchinfo)
	if errscrape != nil {
		return DictInfo{}, errscrape
	}
	return dictinfo, nil
}

// Scrape Naver Dictionary from a Search Term. (Public API)
func GetMessage(searchterm string) (string, error) {
	dictinfo, err := Get(searchterm)
	if err != nil {
		return "", err
	}
	message := Buildmessage(dictinfo)
	return message, nil
}
