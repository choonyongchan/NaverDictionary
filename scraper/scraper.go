package scraper

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// Scrape TOPIK level
func GetTopik(searchinfo map[string]interface{}) (string, error) {
	// Equivalent to searchInfo.entry.entry_level ?? ""
	entry, errorentry := searchinfo["entry"].(map[string]interface{})
	if !errorentry {
		return "", errors.New("Cannot find searchResultMap in searchinfo")
	}
	entrylevel, errorentrylevel := entry["entry_level"].(string)
	if !errorentrylevel {
		return "", errors.New("Cannot find entryLevel in entry")
	}

	containsone := strings.Contains(entrylevel, "1")
	if containsone {
		return "(TOPIK Elementary)", nil
	}
	containstwo := strings.Contains(entrylevel, "2")
	if containstwo {
		return "(TOPIK Intermediate)", nil
	}
	return "", nil
}

// Scrape Importance Stars
func GetImportance(searchinfo map[string]interface{}) (string, error) {
	// Importance is ranked from 0-3 stars.
	// Equivalent to searchInfo.entry.entry_importance ?? 0
	// BEWARE: numstars is a float64, not an int.
	entry, errorentry := searchinfo["entry"].(map[string]interface{})
	if !errorentry {
		return "", errors.New("Cannot find searchResultMap in searchinfo")
	}
	entryimportancefloat, errorentryimportance := entry["entry_importance"].(float64)
	if !errorentryimportance {
		return "", errors.New("Cannot find entryImportance in entry")
	}
	entryimportance := int(entryimportancefloat)
	if entryimportance < 0 || entryimportance > 3 {
		return "", errors.New("Importance is out of range.")
	}
	stars := strings.Repeat("★", entryimportance)
	return stars, nil
}

// Scrape Title
func GetTitle(searchinfo map[string]interface{}) (string, error) {
	// Equivalent to searchInfo.entry.members[0].entry_name ?? ""
	entry, errorentry := searchinfo["entry"].(map[string]interface{})
	if !errorentry {
		return "", errors.New("Cannot find searchResultMap in searchinfo")
	}
	members, errormembers := entry["members"].([]interface{})
	if !errormembers || len(members) == 0 {
		return "", errors.New("Cannot find members in entry")
	}
	member, errormember := members[0].(map[string]interface{})
	if !errormember {
		return "", errors.New("Cannot find member in members")
	}
	entryname, errorentryname := member["entry_name"].(string)
	if !errorentryname {
		return "", errors.New("Cannot find entryName in member")
	}
	return entryname, nil
}

// Scrape Hanja
func GetHanja(searchinfo map[string]interface{}) (string, error) {
	// Equivalent to searchInfo.entry.members[0].origin_language ?? ""
	entry, errorentry := searchinfo["entry"].(map[string]interface{})
	if !errorentry {
		return "", errors.New("Cannot find searchResultMap in searchinfo")
	}
	members, errormembers := entry["members"].([]interface{})
	if !errormembers || len(members) == 0 {
		return "", errors.New("Cannot find members in entry")
	}
	member, errormember := members[0].(map[string]interface{})
	if !errormember {
		return "", errors.New("Cannot find member in members")
	}
	originlanguage, errororiginlanguage := member["origin_language"].(string)
	if !errororiginlanguage {
		return "", errors.New("Cannot find originLanguage in member")
	}
	return originlanguage, nil
}

// Scrape English Definition
func GetEnDef(searchinfo map[string]interface{}) (string, error) {
	// Equivalent to searchInfo.entry.primary_mean ?? ""
	entry, errorentry := searchinfo["entry"].(map[string]interface{})
	if !errorentry {
		return "", errors.New("Cannot find searchResultMap in searchinfo")
	}
	primarymean, errorprimarymean := entry["primary_mean"].(string)
	if !errorprimarymean {
		return "", errors.New("Cannot find primaryMean in entry")
	}

	endefs := strings.Split(primarymean, `|||`)
	endefsNumbered := make([]string, len(endefs))
	for i, def := range endefs {
		endefsNumbered[i] = fmt.Sprintf("%d.%s", i+1, def)
	}
	def := strings.Join(endefsNumbered, " ")
	return def, nil
}

// Scrape Pronunciation
func GetPronun(searchinfo map[string]interface{}) (string, error) {
	// Equivalent to searchInfo.entry.members[0].prons ?? []).map(({ show_pron_symbol }) => show_pron_symbol).filter((p) => p);
	entry, errorentry := searchinfo["entry"].(map[string]interface{})
	if !errorentry {
		return "", errors.New("Cannot find searchResultMap in searchinfo")
	}
	members, errormembers := entry["members"].([]interface{})
	if !errormembers || len(members) == 0 {
		return "", errors.New("Cannot find members in entry")
	}
	member, errormember := members[0].(map[string]interface{})
	if !errormember {
		return "", errors.New("Cannot find member in members")
	}
	prons, errorprons := member["prons"].([]interface{})
	if !errorprons || len(prons) == 0 {
		return "", errors.New("Cannot find prons in member")
	}
	if len(prons) < 2 {
		return "", errors.New("Cannot find two pronunciations in prons")
	}

	firstprons, errorfirstprons := prons[0].(map[string]interface{})
	if !errorfirstprons {
		return "", errors.New("Cannot find first prons in prons")
	}
	engpronun, errorengpronun := firstprons["show_pron_symbol"].(string)
	if !errorengpronun {
		return "", errors.New("Cannot find ShowPronSymbol in first prons")
	}

	secondprons, errorsecondprons := prons[1].(map[string]interface{})
	if !errorsecondprons {
		return "", errors.New("Cannot find second prons in prons")
	}
	korproun, errorkorproun := secondprons["show_pron_symbol"].(string)
	if !errorkorproun {
		return "", errors.New("Cannot find ShowPronSymbol in second prons")
	}
	pronun := fmt.Sprintf("[%s] [%s]", engpronun, korproun)
	return pronun, nil
}

// Scrape Part of Speech
func GetPartSpeech(searchinfo map[string]interface{}) (string, error) {
	// Equivalent to searchInfo.entry.means[0].part.part_ko_name ?? ""
	entry, errorentry := searchinfo["entry"].(map[string]interface{})
	if !errorentry {
		return "", errors.New("Cannot find searchResultMap in searchinfo")
	}
	means, errormeans := entry["means"].([]interface{})
	if !errormeans || len(means) == 0 {
		return "", errors.New("Cannot find means in entry")
	}
	mean, errormean := means[0].(map[string]interface{})
	if !errormean {
		return "", errors.New("Cannot find mean in means")
	}
	part, errorpart := mean["part"].(map[string]interface{})
	if !errorpart {
		return "", errors.New("Cannot find part in mean")
	}
	partKoName, errorpartKoName := part["part_ko_name"].(string)
	if !errorpartKoName {
		return "", errors.New("Cannot find partKoName in part")
	}
	return partKoName, nil
}

// Scrape Description
func GetMeaning(meaningitem map[string]interface{}, idx int) (string, error) {
	// Step 1. Get Meaning
	meaning, errormeaning := meaningitem["show_mean"].(string)
	if !errormeaning {
		return "", errors.New("Cannot find ShowMean in meaningitem")
	}
	meaningstr := fmt.Sprintf("%d.%s", idx+1, meaning)

	// Step 2. Get English and Korean Description
	descitem, errordescitem := meaningitem["description_json"].(string)
	if !errordescitem {
		return "", errors.New("Cannot find DescriptionJSON in meaningitem")
	}

	var result map[string]interface{}
	errordecode := json.Unmarshal([]byte(descitem), &result)
	if errordecode != nil {
		return "", errors.New("Cannot decode JSON")
	}

	// Step 2.1 Get English Description
	enstr, errorenstr := result["en"].(string)
	if !errorenstr {
		return "", errors.New("Cannot find En in result")
	}

	// Step 2.2 Get Korean Description
	kostr, errorkostr := result["ko"].(string)
	if !errorkostr {
		return "", errors.New("Cannot find Ko in result")
	}

	// Step 3. Get Example Sentence
	examples, errorexamples := meaningitem["examples"].([]interface{})
	if !errorexamples || len(examples) == 0 {
		return "", errors.New("Cannot find Examples in meaningitem")
	}
	example, errorexample := examples[0].(map[string]interface{})
	if !errorexample {
		return "", errors.New("Cannot find example in examples")
	}
	originexample, errororiginexample := example["origin_example"].(string)
	if !errororiginexample {
		return "", errors.New("Cannot find OriginExample in example")
	}
	examplestr := fmt.Sprintf("|| %s", originexample)

	// Step 4. Combine Meaning, Description, and Example
	description := fmt.Sprintf("%s\n%s\n%s\n%s", meaningstr, enstr, kostr, examplestr)
	return description, nil
}

// Scrape Meanings of the Word
func GetMeanings(searchinfo map[string]interface{}) (string, error) {
	// Equivalent to searchInfo.entry.means.map(GetMeanings);
	entry, errorentry := searchinfo["entry"].(map[string]interface{})
	if !errorentry {
		return "", errors.New("Cannot find searchResultMap in searchinfo")
	}
	means, errormeans := entry["means"].([]interface{})
	if !errormeans || len(means) == 0 {
		return "", errors.New("Cannot find means in entry")
	}
	meaningsfmted := make([]string, len(means))
	var errormeaning error // Dont have to redefine variable.
	for i, meaning := range means {
		meaningfmt := meaning.(map[string]interface{})
		meaningsfmted[i], errormeaning = GetMeaning(meaningfmt, i)
		if errormeaning != nil {
			return "", errormeaning
		}
	}
	meaning := strings.Join(meaningsfmted, "\n\n")
	return meaning, nil
}

// Scrape Dictionary
func Scrape(searchinfo map[string]interface{}) (DictInfo, error) {
	topik, errortopik := GetTopik(searchinfo)
	if errortopik != nil {
		topik = ""
	}
	importance, errorimportance := GetImportance(searchinfo)
	if errorimportance != nil {
		importance = ""
	}
	title, errortitle := GetTitle(searchinfo)
	if errortitle != nil {
		title = ""
	}
	hanja, errorhanja := GetHanja(searchinfo)
	if errorhanja != nil {
		hanja = ""
	}
	endef, errorendef := GetEnDef(searchinfo)
	if errorendef != nil {
		endef = ""
	}
	pronun, errorpronun := GetPronun(searchinfo)
	if errorpronun != nil {
		pronun = ""
	}
	partspeech, errpartspeech := GetPartSpeech(searchinfo)
	if errpartspeech != nil {
		partspeech = ""
	}
	meanings, errmeanings := GetMeanings(searchinfo)
	if errmeanings != nil {
		meanings = ""
	}
	dictinfo := DictInfo{topik, importance, title, hanja, endef, pronun, partspeech, meanings}
	return dictinfo, nil
}
