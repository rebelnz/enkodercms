package utils

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"html/template"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
	"unicode"
)

var MimeTypesImg = map[string]bool{
	"image/jpeg":   true,
	"image/png":    true,
	"image/gif":    true,
	"image/bmp":    true,
	"image/tiff":   true,
	"image/x-icon": true}

var MimeTypesAllowed = map[string]bool{
	"image/jpeg":                                                        true,
	"image/png":                                                         true,
	"image/gif":                                                         true,
	"image/bmp":                                                         true,
	"image/tiff":                                                        true,
	"image/x-icon":                                                      true,
	"application/msword":                                                true,
	"text/plain":                                                        true,
	"application/pdf":                                                   true,
	"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet": true}

type JSONResponse map[string]interface{}

func (r JSONResponse) String() (s string) {
	j, err := json.Marshal(r)
	if err != nil {
		s = ""
		return
	}
	s = string(j)
	return
}

func NewPassword() string {
	pwd, err := exec.Command("pwgen", "-1").Output()
	if err != nil {
		fmt.Println(err)
	}
	pwdn := fmt.Sprintf("%s", pwd)
	return strings.Replace(pwdn, "\n", "", -1) // pwgen adds a \n
}

// GenerateSlug converts a string into a lowercase dasherized slug
// For example: GenerateSlug("My new page") returns "my-new-page"
func GenerateSlug(str string) (slug string) {
	return strings.Map(func(r rune) rune {
		switch {
		case r == ' ', r == '-':
			return '-'
		case r == '_', unicode.IsLetter(r), unicode.IsDigit(r):
			return r
		default:
			return -1
		}
		return -1
	}, strings.ToLower(strings.TrimSpace(str)))
}

// function to check if string is in slice of strings
// used to test value of checkboxes
// see template/admin/editpage for example
func CheckIt(x string, subW []string) bool {
	for _, b := range subW {
		if b == x {
			return true
		}
	}
	return false
}

// function to time function execution speed
func Timed(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
	// use like this
	// defer timed(time.Now(), "functionName")
}

// find uploaded filetype -- check fontawesome for icon name
// http://fontawesome.github.io/Font-Awesome/icons/#file-type
func FileIcon(fileName string) (icon string) {
	fn := filepath.Ext(fileName)
	switch fn {
	case ".pdf":
		return "pdf"
	case ".txt":
		return "text"
	case ".doc", ".docx":
		return "word"
	default:
		return ""
	}
}

// this changes by css framework
func WidgetCss(size string) (class string) {
	switch size {
	case "small":
		return "small"
	case "medium":
		return "medium"
	case "large":
		return "large"
	default:
		return "small"
	}
}

func FormatTime(t time.Time) string {
	format := "Mon Jan 2 15:04:05"
	return t.Format(format)
}

func RandStr(strSize int, randType string) string {
	var dictionary string
	if randType == "alphanum" {
		dictionary = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	}
	if randType == "alpha" {
		dictionary = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	}
	if randType == "number" {
		dictionary = "0123456789"
	}
	var bytes = make([]byte, strSize)
	rand.Read(bytes)
	for k, v := range bytes {
		bytes[k] = dictionary[v%byte(len(dictionary))]
	}
	return string(bytes)
}

func PostSummary(s string) string {
	if strings.Contains(s, "|") {
		split := strings.Split(s, "|")
		return split[1]
	}
	return ""
}

func PostContent(s string) template.HTML {
	if strings.Contains(s, "|") {
		return template.HTML(strings.Replace(s, "|", "", -1))
	}
	return template.HTML(s)
}
