package scraper

import (
	"testing"
)

// Get Keys of a Map for debugging.
func Keys(m map[string]interface{}) []string {
	keys := make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

var examplesearchinfo = map[string]interface{}{
	"entry": map[string]interface{}{
		"entry_level":      "1",
		"entry_importance": 2.0,
		"primary_mean":     "puppy|||small dog|||young dog",
		"members": []interface{}{
			map[string]interface{}{
				"entry_name":      "강아지",
				"origin_language": "奮發",
				"prons": []interface{}{
					map[string]interface{}{"show_pron_symbol": "gang-a-ji"},
					map[string]interface{}{"show_pron_symbol": "강아지"},
				},
			},
		},
		"means": []interface{}{
			map[string]interface{}{
				"part": map[string]interface{}{
					"part_ko_name": "명사",
				},
				"show_mean":        "강아지",
				"description_json": `{"en":"a puppy or young dog", "ko":"어린 개"}`,
				"examples": []interface{}{
					map[string]interface{}{
						"origin_example": "강아지가 귀엽다",
					},
				},
			},
		},
	},
}

func TestGetTopikExample(t *testing.T) {
	topik, error := GetTopik(examplesearchinfo)
	if error != nil {
		t.Errorf("GetTopik(%q) = %q; want no error", examplesearchinfo, error)
	}
	if topik != "(TOPIK Elementary)" {
		t.Errorf("Expected 1, got %s", topik)
	}
}

func TestGetTopikExhaustive(t *testing.T) {

	emptytopikdata := map[string]interface{}{
		"entry": map[string]interface{}{
			"entry_level": "",
		},
	}

	onetopikdata := map[string]interface{}{
		"entry": map[string]interface{}{
			"entry_level": "1",
		},
	}

	twotopikdata := map[string]interface{}{
		"entry": map[string]interface{}{
			"entry_level": "2",
		},
	}

	topikempty, errortopikempty := GetTopik(emptytopikdata)
	if errortopikempty != nil {
		t.Errorf("GetTopik(%q) = %q; want no error", emptytopikdata, errortopikempty)
	}
	if topikempty != "" {
		t.Errorf("Expected empty string, got %s", topikempty)
	}

	topikone, errortopikone := GetTopik(onetopikdata)
	if errortopikone != nil {
		t.Errorf("GetTopik(%q) = %q; want no error", onetopikdata, errortopikone)
	}
	if topikone != "(TOPIK Elementary)" {
		t.Errorf("Expected (TOPIK Elementary), got %s", topikone)
	}

	topiktwo, errortopiktwo := GetTopik(twotopikdata)
	if errortopiktwo != nil {
		t.Errorf("GetTopik(%q) = %q; want no error", twotopikdata, errortopiktwo)
	}
	if topiktwo != "(TOPIK Intermediate)" {
		t.Errorf("Expected (TOPIK Intermediate) , got %s", topiktwo)
	}
}

func TestGetImportanceExample(t *testing.T) {
	importance, error := GetImportance(examplesearchinfo)
	if error != nil {
		t.Errorf("GetImportance(%q) = %q; want no error", examplesearchinfo, error)
	}
	if importance != "★★" {
		t.Errorf("Expected 2 stars, got %s", importance)
	}
}

func TestGetImportanceExhaustive(t *testing.T) {
	oneimportancedata := map[string]interface{}{
		"entry": map[string]interface{}{
			"entry_importance": 1.0,
		},
	}
	twoimportancedata := map[string]interface{}{
		"entry": map[string]interface{}{
			"entry_importance": 2.0,
		},
	}
	threeimportancedata := map[string]interface{}{
		"entry": map[string]interface{}{
			"entry_importance": 3.0,
		},
	}

	importanceone, errorone := GetImportance(oneimportancedata)
	if errorone != nil {
		t.Errorf("GetImportance(%q) = %q; want no error", oneimportancedata, errorone)
	}
	if importanceone != "★" {
		t.Errorf("Expected 1 star, got %s", importanceone)
	}

	importancetwo, errortwo := GetImportance(twoimportancedata)
	if errortwo != nil {
		t.Errorf("GetImportance(%q) = %q; want no error", twoimportancedata, errortwo)
	}
	if importancetwo != "★★" {
		t.Errorf("Expected 2 stars, got %s", importancetwo)
	}

	importancethree, emptythree := GetImportance(threeimportancedata)
	if emptythree != nil {
		t.Errorf("GetImportance(%q) = %q; want no error", threeimportancedata, emptythree)
	}
	if importancethree != "★★★" {
		t.Errorf("Expected 3 stars, got %s", importancethree)
	}
}

func TestGetTitleExample(t *testing.T) {
	title, error := GetTitle(examplesearchinfo)
	if error != nil {
		t.Errorf("GetTitle(%q) = %q; want no error", examplesearchinfo, error)
	}
	if title != "강아지" {
		t.Errorf("Expected 강아지, got %s", title)
	}
}

func TestGetHanjaExample(t *testing.T) {
	hanja, error := GetHanja(examplesearchinfo)
	if error != nil {
		t.Errorf("GetHanja(%q) = %q; want no error", examplesearchinfo, error)
	}
	if hanja != "奮發" {
		t.Errorf("Expected 奮發, got %s", hanja)
	}
}

func TestGetEnDefExample(t *testing.T) {
	endef, error := GetEnDef(examplesearchinfo)
	if error != nil {
		t.Errorf("GetEnDef(%q) = %q; want no error", examplesearchinfo, error)
	}
	if endef != "1.puppy 2.small dog 3.young dog" {
		t.Errorf("Expected 1.puppy 2.small dog 3.young dog, got %s", endef)
	}
}

func TestGetPronunExample(t *testing.T) {
	pronun, error := GetPronun(examplesearchinfo)
	if error != nil {
		t.Errorf("GetPronun(%q) = %q; want no error", examplesearchinfo, error)
	}
	if pronun != "[gang-a-ji] [강아지]" {
		t.Errorf("Expected [gang-a-ji] [강아지], got %s", pronun)
	}
}

func TestGetPartSpeechExample(t *testing.T) {
	partspeech, error := GetPartSpeech(examplesearchinfo)
	if error != nil {
		t.Errorf("GetPartSpeech(%q) = %q; want no error", examplesearchinfo, error)
	}
	if partspeech != "명사" {
		t.Errorf("Expected 명사, got %s", partspeech)
	}
}

func TestGetMeaningsExample(t *testing.T) {
	meaning, error := GetMeanings(examplesearchinfo)
	if error != nil {
		t.Errorf("GetMeanings(%q) = %q; want no error", examplesearchinfo, error)
	}
	if meaning != "1.강아지\na puppy or young dog\n어린 개\n|| 강아지가 귀엽다" {
		t.Errorf("Expected 1.강아지\na puppy or young dog\n어린 개\n|| 강아지가 귀엽다, got %s", meaning)
	}
}

func TestScrapeExample(t *testing.T) {
	scraped, error := Scrape(examplesearchinfo)
	expected := DictInfo{
		Topik:      "(TOPIK Elementary)",
		Importance: "★★",
		Title:      "강아지",
		Hanja:      "奮發",
		Endef:      "1.puppy 2.small dog 3.young dog",
		Pronun:     "[gang-a-ji] [강아지]",
		Partspeech: "명사",
		Meanings:   "1.강아지\na puppy or young dog\n어린 개\n|| 강아지가 귀엽다",
	}
	if error != nil {
		t.Errorf("Scrape(%q) = %q; want no error", scraped, error)
	}
	if scraped != expected {
		t.Errorf("Expected %v, got %v", expected, scraped)
	}
}
