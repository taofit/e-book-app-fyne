package navList

import (
	"fmt"

	"github.com/taofit/e-book-fyne/internal/articles"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

const PreferenceCurrentArticle = "currentArticle"

type NavSectionList struct {
	list *widget.List
}
type TappableLabel struct {
	widget.Label
	listItemID widget.ListItemID
	parent     *NavSectionList
}

func NewTappableLabel(parent *NavSectionList, text string) *TappableLabel {
	tpbLabel := &TappableLabel{
		Label:      widget.Label{Text: text},
		parent:     parent,
		listItemID: -1,
	}
	tpbLabel.ExtendBaseWidget(tpbLabel)

	return tpbLabel
}

func (t *TappableLabel) SetListItemID(id widget.ListItemID) {
	t.listItemID = id
}

func (t *TappableLabel) Tapped(pe *fyne.PointEvent) {
	t.parent.list.Unselect(t.listItemID) //need to unselect the listItemID first, as if it is selected, it won't run the code in Select func
	t.parent.list.Select(t.listItemID)
}

func (section *NavSectionList) MakeMblNav(
	setArticle func(article articles.Article),
	setSubList func(listTitle string, list fyne.CanvasObject),
	articlesForSubject []string,
	loadPrevious bool,
) fyne.CanvasObject {
	curApp := fyne.CurrentApp()

	section.list = &widget.List{
		Length: func() int {
			return len(articlesForSubject)
		},
		CreateItem: func() fyne.CanvasObject {
			return NewTappableLabel(section, "single articles")
		},
		UpdateItem: func(id widget.ListItemID, item fyne.CanvasObject) {
			subjectName := articlesForSubject[id]
			listTitle := articles.Articles[subjectName].Title
			item.(*TappableLabel).SetText(listTitle)
			item.(*TappableLabel).SetListItemID(id)
		},
		OnSelected: func(id widget.ListItemID) {
			subjectName := articlesForSubject[id]
			listTitle := articles.Articles[subjectName].Title
			articlesForSubject, ok := articles.ArticleIndex[subjectName]
			if ok {
				subNavSection := NavSectionList{}
				list := subNavSection.MakeMblNav(setArticle, setSubList, articlesForSubject, false)
				setSubList(listTitle, list)
			} else {
				if a, ok := articles.Articles[subjectName]; ok {
					curApp.Preferences().SetInt(PreferenceCurrentArticle, id)
					setArticle(a)
				}
			}
		},
	}

	if loadPrevious {
		currentPref := curApp.Preferences().IntWithFallback(PreferenceCurrentArticle, 0)
		section.list.Select(currentPref)
	}

	return section.list
}

func (section *NavSectionList) MakeNav(
	setArticle func(article articles.Article),
	setSubList func(listTitle string, list fyne.CanvasObject),
	articlesForSubject []string,
	gotoParentLevel func(articlesForSubject []string),
	loadPrevious bool,
) fyne.CanvasObject {
	curApp := fyne.CurrentApp()

	section.list = &widget.List{
		Length: func() int {
			parentList := getNoRootParentList(articlesForSubject)
			if len(parentList) > 0 {
				gotoParentLevel(parentList)
			}
			return len(articlesForSubject)
		},
		CreateItem: func() fyne.CanvasObject {
			return NewTappableLabel(section, "single articles")
		},
		UpdateItem: func(id widget.ListItemID, item fyne.CanvasObject) {
			subjectName := articlesForSubject[id]
			listTitle := articles.Articles[subjectName].Title
			item.(*TappableLabel).SetText(listTitle)
			item.(*TappableLabel).SetListItemID(id)
		},
		OnSelected: func(id widget.ListItemID) {
			subjectName := articlesForSubject[id]
			listTitle := articles.Articles[subjectName].Title
			aclsForSubject, ok := articles.ArticleIndex[subjectName]
			if ok {
				subNavSection := NavSectionList{}
				list := subNavSection.MakeNav(setArticle, setSubList, aclsForSubject, gotoParentLevel, false)
				setSubList(listTitle, list)
			} else {
				if a, ok := articles.Articles[subjectName]; ok {
					curApp.Preferences().SetInt(PreferenceCurrentArticle, id)
					gotoParentLevel(articlesForSubject)
					setArticle(a)
				}
			}
		},
	}

	if loadPrevious {
		currentPref := curApp.Preferences().IntWithFallback(PreferenceCurrentArticle, 0)
		section.list.Select(currentPref)
	}

	return section.list
}

func findParentList(subjectName string) []string {
	var key string
	var parentKey string
	for k, subList := range articles.ArticleIndex {
		if contains(subList, subjectName) {
			key = k
			break
		}
	}
	if key != "" {
		for k, subList := range articles.ArticleIndex {
			if contains(subList, key) && k != articles.RootSubjectsKey {
				parentKey = k
				break
			}
		}
	}

	aclsForSubject, ok := articles.ArticleIndex[parentKey]
	if ok {
		return aclsForSubject
	}
	return []string{}
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func getNoRootParentList(articlesForSubject []string) []string {
	if len(articlesForSubject) > 0 {
		parentLevelList := findParentList(articlesForSubject[0])
		if len(parentLevelList) > 0 {
			return parentLevelList
		}
	}
	return []string{}
}
