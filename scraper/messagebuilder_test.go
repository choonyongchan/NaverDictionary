package scraper

import (
	"testing"
)

var emptydictinfo = DictInfo{
	Topik:      "",
	Importance: "",
	Title:      "",
	Hanja:      "",
	Endef:      "",
	Pronun:     "",
	Partspeech: "",
	Meanings:   "",
}

var completedictinfo = DictInfo{
	Topik:      "(TOPIK Elementary)",
	Importance: "★★",
	Title:      "강아지",
	Hanja:      "奮發",
	Endef:      "1.puppy 2.small dog 3.young dog",
	Pronun:     "[gang-a-ji] [강아지]",
	Partspeech: "명사",
	Meanings:   "TestMeaning\nTestMeaning2",
}

func TestBuildsentence(t *testing.T) {
	// Test empty input
	result := Buildsentence("", []string{})
	if result != "" {
		t.Errorf("Expected empty string, got %s", result)
	}

	// Test single input
	result = Buildsentence("", []string{"Hello"})
	if result != "Hello" {
		t.Errorf("Expected Hello, got %s", result)
	}

	// Test multiple inputs
	result = Buildsentence("", []string{"Hello", "World"})
	if result != "Hello World" {
		t.Errorf("Expected Hello World, got %s", result)
	}
}

func TestBuildsentenceEmptyValid(t *testing.T) {
	// Test empty input
	result := Buildsentence("", []string{""})
	if result != "" {
		t.Errorf("Expected empty string, got %s", result)
	}
}

func TestBuildsentenceSingleValid(t *testing.T) {
	// Test single input
	result := Buildsentence("", []string{"Hello"})
	if result != "Hello" {
		t.Errorf("Expected Hello, got %s", result)
	}
}

func TestBuildsentenceMultipleValid(t *testing.T) {
	// Test multiple inputs
	result := Buildsentence("", []string{"Hello", "World"})
	if result != "Hello World" {
		t.Errorf("Expected Hello World, got %s", result)
	}
}

func TestBuildmessageEmptyValid(t *testing.T) {
	// Test empty input
	result := Buildmessage(emptydictinfo)
	if result != "" {
		t.Errorf("Expected empty string, got %s", result)
	}
}

func TestBuildmessageCompleteValid(t *testing.T) {
	// Test complete input
	result := Buildmessage(completedictinfo)
	expected := "(TOPIK Elementary) ★★\n강아지 奮發\n1.puppy 2.small dog 3.young dog\n----------\nPronunciation:\nroma [gang-a-ji] [강아지]\n----------\n명사\nTestMeaning\nTestMeaning2"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
