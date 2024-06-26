package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"
	"unicode/utf8"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

const (
	AZERTY = "Azerty"
	QWERTY = "Qwerty"
)

type Config struct {
	LastOpenedDocument string `json:"lastOpenDocument"`
	KeyboardLang       string `json:"keyboardLang"`
	FontScale          uint8  `json:"FontScale"`
}

func LoadConfig(saveLocation string) Config {
	content, err := os.ReadFile(path.Join(saveLocation, "config.json"))

	if err != nil {
		id, _ := gonanoid.New()
		return Config{
			LastOpenedDocument: id + ".txt",
		}
	}

	var config Config

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	err = json.Unmarshal(content, &config)
	if err != nil {
		fmt.Println("Unmarshal error:", err)
	}

	if config.FontScale == 0 {
		config.FontScale = 3 // Set default font scaling to Large
	}

	return config
}

func SaveConfig(config Config, saveLocation string) {
	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	content, _ := json.Marshal(config)

	err := os.MkdirAll(saveLocation, 0777)
	if err != nil {
		fmt.Println("MkdirAll error:", err)
	}

	err = os.WriteFile(path.Join(saveLocation, "config.json"), []byte(content), 0777)
	if err != nil {
		fmt.Println("WriteFile error:", err)
	}
}

func IsLetter(s string) bool {
	return !strings.Contains(s, "KEY")
}

func InsertAt(text string, insert string, index int) string {
	if index == LenString(text) {
		return text + insert
	}
	runeText := []rune(text)
	return string(append(runeText[:index], append([]rune(insert), runeText[index:]...)...))
}

func DeleteAt(text string, index int) string {
	runeText := []rune(text)
	return string(append(runeText[:index-1], runeText[index:]...))
}

func LenString(s string) int {
	return utf8.RuneCountInString(s)
}
