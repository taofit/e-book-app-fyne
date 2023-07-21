package navTree

import (
	"github.com/taofit/e-book-fyne/internal/articles"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

const preferenceCurrentTreeArticle string = "current article"

func MakeNav(
	setArticle func(article articles.Article),
	loadPrevious bool,
) fyne.CanvasObject {
	curApp := fyne.CurrentApp()

	tree := &widget.Tree{
		ChildUIDs: func(uid widget.TreeNodeID) []widget.TreeNodeID {
			return articles.ArticleIndex[uid]
		},
		IsBranch: func(uid widget.TreeNodeID) bool {
			children, ok := articles.ArticleIndex[uid]

			return ok && len(children) > 0
		},
		CreateNode: func(branch bool) fyne.CanvasObject {
			return widget.NewLabel("article Widgets")
		},
		UpdateNode: func(uid widget.TreeNodeID, branch bool, node fyne.CanvasObject) {
			a, ok := articles.Articles[uid]
			if !ok {
				fyne.LogError("Missing article panel: "+uid, nil)
				return
			}
			node.(*widget.Label).SetText(a.Title)
		},
		OnSelected: func(uid widget.TreeNodeID) {
			if a, ok := articles.Articles[uid]; ok {
				curApp.Preferences().SetString(preferenceCurrentTreeArticle, uid)
				setArticle(a)
			}
		},
	}

	if loadPrevious {
		currentPref := curApp.Preferences().StringWithFallback(preferenceCurrentTreeArticle, "forward")
		tree.Select(currentPref)
	}

	return tree
}
