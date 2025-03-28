package scraper

import (
	"testing"
)

func TestSanitiseIllegal(t *testing.T) {
	unsanitised := "％＆＊（@＆％＊（안#％@＆＊％＊（＆@＆％녕！@￥#％@％@￥＊＆@（（＆@％（123124141341{}:><<>"
	want := "안녕"
	got := Sanitise(unsanitised)
	if got != want {
		t.Errorf("Sanitise(%q) = %q; want %q", unsanitised, got, want)
	}
}

func TestSanitiseIdentity(t *testing.T) {
	unsanitised := "하하"
	want := "하하"
	got := Sanitise(unsanitised)
	if got != want {
		t.Errorf("Sanitise(%q) = %q; want %q", unsanitised, got, want)
	}
}

func TestSanitiseEmpty(t *testing.T) {
	unsanitised := ""
	want := ""
	got := Sanitise(unsanitised)
	if got != want {
		t.Errorf("Sanitise(%q) = %q; want %q", unsanitised, got, want)
	}
}

func TestFetchValid(t *testing.T) {
	url := "http://korean.dict.naver.com/api3/koen/search?m=mobile&query=%EC%95%88%EB%85%95&range=entrySearch"
	got, error := Fetch(url)
	if error != nil {
		t.Errorf("Fetch(%q) = %q; want no error", url, error)
	}
	if got == nil {
		t.Errorf("Fetch(%q) = %q; want not nil", url, got)
	}
}

func TestFetchInvalid(t *testing.T) {
	url := "INVALID"
	_, error := Fetch(url)
	if error == nil {
		t.Errorf("Fetch(%q) = %q; want error", url, error)
	}
}

func TestGetEntryUrl(t *testing.T) {
	searchterm := "안녕"
	want := "https://korean.dict.naver.com/api3/koen/search?m=mobile&query=%EC%95%88%EB%85%95&range=entrySearch"
	got, error := GetEntryUrl(searchterm)
	if error != nil {
		t.Errorf("GetEntryUrl(%q) = %q; want no error", searchterm, error)
	}
	if got != want {
		t.Errorf("GetEntryUrl(%q) = %q; want %q", searchterm, got, want)
	}
}

func TestGetEntryInfoValid(t *testing.T) {
	searchterm := "안녕"
	entryinfo, error := GetEntryInfo(searchterm)
	if error != nil {
		t.Errorf("GetEntryInfo(%q) = %q; want no error", searchterm, error)
	}
	if entryinfo == nil {
		t.Errorf("GetEntryInfo(%q) = %q; want not nil", searchterm, entryinfo)
	}
}

func TestGetEntryId(t *testing.T) {
	entryinfo := map[string]interface{}{
		"searchResultMap": map[string]interface{}{
			"searchResultListMap": map[string]interface{}{
				"WORD": map[string]interface{}{
					"items": []interface{}{
						map[string]interface{}{
							"entryId": "12345ASDFSDGfsg6",
						},
					},
				},
			},
		},
	}
	want := "12345ASDFSDGfsg6"
	got, error := GetEntryId(entryinfo)
	if error != nil {
		t.Errorf("GetEntryId(%q) = %q; want no error", entryinfo, error)
	}
	if got != want {
		t.Errorf("GetEntryId(%q) = %q; want %q", entryinfo, got, want)
	}
}

func TestGetSearchUrl(t *testing.T) {
	entryid := "ac75d1845900457bbda2fdbc4fbaac05" //사랑
	want := "https://korean.dict.naver.com/api/platform/koen/entry?entryId=ac75d1845900457bbda2fdbc4fbaac05"
	got, error := GetSearchUrl(entryid)
	if error != nil {
		t.Errorf("GetSearchUrl(%q) = %q; want no error", entryid, error)
	}
	if got != want {
		t.Errorf("GetSearchUrl(%q) = %q; want %q", entryid, got, want)
	}
}

func TestGetSearchInfoValid(t *testing.T) {
	entryid := "ac75d1845900457bbda2fdbc4fbaac05" //사랑
	searchinfo, errorsearchinfo := GetSearchInfo(entryid)
	if errorsearchinfo != nil {
		t.Errorf("GetSearchInfo(%q) = %q; want no error", entryid, errorsearchinfo)
	}
	if searchinfo == nil {
		t.Errorf("GetSearchInfo(%q) = %q; want not nil", entryid, searchinfo)
	}
	searchterm, errorsearchterm := GetTitle(searchinfo)
	if errorsearchterm != nil {
		t.Errorf("GetSearchInfo(%q) = %q; want no error", entryid, errorsearchterm)
	}
	if searchterm != "사랑" {
		t.Errorf("GetSearchInfo(%q) = %q; want not empty", entryid, searchterm)
	}
}

func TestGetValid(t *testing.T) {
	searchterm := "안녕"
	got, error := Get(searchterm)
	if error != nil {
		t.Errorf("Get(%q) = %q; want no error", searchterm, error)
	}
	got_searchterm := got.Title
	want := "안녕"
	if got_searchterm != want {
		t.Errorf("Get(%q) = %q; want %q", searchterm, got, want)
	}
}

func TestGetSparse(t *testing.T) {
	searchterm := "저출새" // Test Inference Skill
	got, error := Get(searchterm)
	if error != nil {
		t.Errorf("Get(%q) = %q; want no error", searchterm, error)
	}
	got_searchterm := got.Title
	want := "저출생"
	if got_searchterm != want {
		t.Errorf("Get(%q) = %q; want %q", searchterm, got, want)
	}
}

func TestGetInvalid(t *testing.T) {
	searchterm := "gfgdg!@#"
	_, error := Get(searchterm)
	if error == nil {
		t.Errorf("Get(%q) = %q; want error", searchterm, error)
	}
}
