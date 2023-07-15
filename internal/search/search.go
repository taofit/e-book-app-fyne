package search

import (
	"reflect"
	"runtime"
	"strings"

	"github.com/taofit/e-book-fyne/internal/articles"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type articleWithKey struct {
	key string
	articles.Article
}

type resultItem struct {
	key  string
	text string
}

func MakeSearchEntry(
	setSearchResult func(resultCnt fyne.CanvasObject, input string),
	setArticle func(article articles.Article)) *fyne.Container {
	input := widget.NewEntry()
	input.SetPlaceHolder("enter keyword")
	searchBtn := widget.NewButtonWithIcon("", theme.SearchIcon(), func() {
		resultList := makeResultList(input.Text, searchKeyword, setArticle)
		setSearchResult(resultList, input.Text)
	})
	content := container.NewBorder(
		nil,
		nil,
		nil,
		searchBtn,
		input,
	)

	return content
}

func searchKeyword(word string) []resultItem {
	if !isValidLength(word) {
		return []resultItem{}
	}
	workers := 2 * runtime.GOMAXPROCS(0)
	jobs := make(chan articleWithKey)
	subResults := make(chan []resultItem)

	for w := 0; w < workers; w++ {
		go worker(jobs, subResults, word)
	}
	for key, acl := range articles.Articles {
		jobs <- articleWithKey{key, acl}
	}

	close(jobs)
	collectedResult := []resultItem{}
	for w := 0; w < workers; w++ {
		collectedResult = append(collectedResult, <-subResults...)
	}

	return collectedResult
}

func worker(jobs <-chan articleWithKey, subResults chan<- []resultItem, word string) {
	resultsArr := []resultItem{}
	for job := range jobs {
		text := job.GetFileContent()
		runeText := []rune(text)
		runeWord := []rune(word)

		i := containWord(runeText, runeWord)
		if i > -1 {
			startIdx := i - 10
			if startIdx < 0 {
				startIdx = 0
			}
			endIdx := i + len(runeWord) + 10
			if endIdx > len(runeText) {
				endIdx = len(runeText)
			}
			resultStr := string(runeText[startIdx:endIdx])
			resultsArr = append(resultsArr, resultItem{job.key, resultStr})
		}
	}
	subResults <- resultsArr
}

func containWord(text []rune, word []rune) int {
	lengthText := len(text)
	lengthWord := len(word)
	if lengthWord > lengthText {
		return -1
	}

	for i := 0; i <= lengthText-lengthWord; i++ {
		if reflect.DeepEqual(text[i:i+lengthWord], word) {
			return i
		}
	}

	return -1
}

func isValidLength(word string) bool {
	if strings.TrimSpace(word) == "" || len(strings.TrimSpace(word)) > 20 {
		return false
	}
	return true
}
