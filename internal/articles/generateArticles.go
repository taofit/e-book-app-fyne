package articles

import (
	"embed"
	"encoding/json"
	"log"
	"unicode/utf8"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Article struct {
	Title, Intro string
	FilePath     string
}
type ContentArticle struct {
	Title string `json:"title"`
	File  string `json:"file"`
}

type Subject struct {
	Key             string    `json:"key"`
	Title           string    `json:"title"`
	TableOfContents []Subject `json:"tableOfContents"`
}

const assetsFolder = "assets"
const articleIdxFile = "articles_index.json"

var (
	//go:embed assets/*
	articlesFolder embed.FS

	Articles = map[string]Article{}
	// Articles = map[string]Article{
	// 	"chronology":     {"important event", "", ""},
	// 	"biography":      {"biography", "", ""},
	// 	"youth":          {"earlier age", "", ""},
	// 	"monastery":      {"lion roar", "", ""},
	// 	"dharmaFounding": {"init the dharma", "", ""},
	// 	"virtues":        {"saint", "", ""},
	// }

	ArticleIndex = map[string][]string{}
	// ArticleIndex = map[string][]string{
	// 	"rootSubjects": {"biography", "chronology"},
	// 	"chronology": {"youth", "monastery", "dharmaFounding", "virtues"},
	// }
	RootSubjectsKey = ""
)

func PopulateArticles() {
	var rootSubjects []string
	var subjects = loadArticles()
	for _, subject := range subjects {
		rootSubjects = append(rootSubjects, subject.Key)
	}
	processArticles(subjects)
	ArticleIndex[RootSubjectsKey] = rootSubjects
}

func loadArticles() []Subject {
	var subjects = []Subject{}
	var articleIdxFilePath = assetsFolder + "/" + articleIdxFile

	jsonArticleList, err := articlesFolder.ReadFile(articleIdxFilePath)
	if err != nil {
		log.Printf("Could not load json file: %s\n", err)
		return []Subject{}
	}
	err = json.Unmarshal(jsonArticleList, &subjects)

	if err != nil {
		log.Printf("Could not unmarshal json: %s\n", err)
		return []Subject{}
	}

	return subjects
}

func processArticles(subjects []Subject) {
	for _, subject := range subjects {
		key := subject.Key
		if subject.hasChildren() {
			processArticles(subject.TableOfContents)
			subject.populateArticlesUnderSubject(key)
		} else {
			filePath := assetsFolder + "/" + key + ".txt"
			Articles[key] = Article{Title: subject.Title, Intro: "", FilePath: filePath}
		}
	}
}

func (subject Subject) populateArticlesUnderSubject(key string) {
	var subjectTableOfContent = []string{}
	for _, contentArticle := range subject.TableOfContents {
		fileName := contentArticle.Key
		subjectTableOfContent = append(subjectTableOfContent, fileName)
	}
	ArticleIndex[key] = subjectTableOfContent
	Articles[key] = Article{Title: subject.Title, Intro: "", FilePath: ""}
}

func (subject Subject) hasChildren() bool {
	return len(subject.TableOfContents) > 0
}

func (a Article) LoadFile(_ fyne.Window) fyne.CanvasObject {
	return generateText(a.GetFileContent())
}

func (a Article) GetFileContent() string {
	input, err := articlesFolder.ReadFile(a.FilePath)
	if err != nil {
		return ""
	}
	return string(input)
}

func (a Article) HasFile() bool {
	return a.FilePath != ""
}

func (a Article) ShortenTitle(length ...int) string {
	lenTemp := 30
	if len(length) > 0 {
		lenTemp = length[0]
	}
	if utf8.RuneCountInString(a.Title) > lenTemp {
		return a.Title[0:lenTemp] + "..."
	}

	return a.Title
}

func generateText(content string) fyne.CanvasObject {
	ricTxt := widget.NewRichTextWithText(content)
	ricTxt.Wrapping = fyne.TextWrapWord
	ricTxt.Scroll = container.ScrollBoth

	return ricTxt
}
