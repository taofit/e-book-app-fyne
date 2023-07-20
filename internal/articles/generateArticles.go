package articles

import (
	"embed"
	"encoding/json"
	"log"

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
	// 	"chronology":     {"重要纪事", "", ""},
	// 	"biography":      {"法尊传记", "", ""},
	// 	"youth":          {"年少峥嵘", "", ""},
	// 	"monastery":      {"祖庭狮吼", "", ""},
	// 	"dharmaFounding": {"华藏开宗", "", ""},
	// 	"virtues":        {"圣德圣行", "", ""},
	// }

	ArticleIndex = map[string][]string{}
	// ArticleIndex = map[string][]string{
	// 	"rootSubjects": {"biography", "chronology"},
	// 	"chronology": {"youth", "monastery", "dharmaFounding", "virtues"},
	// }
)

func PopulateArticles() {
	var rootSubjects []string
	var subjects = loadArticles()

	for _, subject := range subjects {
		rootSubjects = append(rootSubjects, subject.Key)
	}
	processArticles(subjects)
	ArticleIndex["rootSubjects"] = rootSubjects
}

func loadArticles() []Subject {
	var subjects = []Subject{}
	var articleIdxFilePath = assetsFolder + "/" + articleIdxFile

	jsonArticleList, err := articlesFolder.ReadFile(articleIdxFilePath)
	if err != nil {
		log.Printf("Could not load json file: %s\n", err)
		return []Subject{}
	}
	err = json.Unmarshal([]byte(jsonArticleList), &subjects)

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

func generateText(content string) fyne.CanvasObject {
	ricTxt := widget.NewRichTextWithText(content)
	ricTxt.Wrapping = fyne.TextWrapWord
	ricTxt.Scroll = container.ScrollBoth

	return ricTxt
}
