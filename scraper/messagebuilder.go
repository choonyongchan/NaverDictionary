package scraper

import (
	"strings"
)

func Buildsentence(prefix string, components []string) string {
	// Filter
	filtered := make([]string, 0, len(components))
	for _, str := range components {
		if str != "" {
			filtered = append(filtered, str)
		}
	}
	// Join
	sentence := strings.Join(filtered, " ")
	if sentence == "" {
		return ""
	}
	return prefix + sentence
}

func Buildmessage(dictinfo DictInfo) string {
	// Return empty string if all fields are empty
	if dictinfo.Topik == "" && dictinfo.Importance == "" && dictinfo.Title == "" &&
		dictinfo.Hanja == "" && dictinfo.Endef == "" && dictinfo.Pronun == "" &&
		dictinfo.Partspeech == "" && dictinfo.Meanings == "" {
		return ""
	}

	parts := []string{
		Buildsentence("", []string{dictinfo.Topik, dictinfo.Importance}),
		Buildsentence("", []string{dictinfo.Title, dictinfo.Hanja}),
		Buildsentence("", []string{dictinfo.Endef}),
		Buildsentence("----------\nPronunciation:\nroma ", []string{dictinfo.Pronun}),
		"----------",
		Buildsentence("", []string{dictinfo.Partspeech}),
		Buildsentence("", []string{dictinfo.Meanings}),
	}

	filtered := make([]string, 0, len(parts))
	for _, str := range parts {
		if str != "" {
			filtered = append(filtered, str)
		}
	}

	return strings.Join(filtered, "\n")
}
