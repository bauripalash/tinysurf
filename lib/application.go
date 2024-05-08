package lib

import (

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)



type Application struct {
	app *tview.Application
	addressBar *tview.InputField
	viewpanel *tview.TextView
	currentAdress string
	container *tview.Flex
	rowContainer *tview.Flex
}


func NewApplication() Application {
	return Application{
		app: tview.NewApplication().EnableMouse(true),
		addressBar: tview.NewInputField(),
		viewpanel: tview.NewTextView(),
		container: tview.NewFlex(),
		rowContainer: tview.NewFlex().SetDirection(tview.FlexRow),
		currentAdress: "",
	}
}

func (a * Application) Init() {
	a.addressBar = a.addressBar.SetLabel("?")
	a.addressBar.SetBorder(true)


	a.addressBar = a.addressBar.SetDoneFunc(func(key tcell.Key) {

		a.searchKey()
	})

	a.viewpanel = a.viewpanel.SetDynamicColors(true)
	a.viewpanel = a.viewpanel.SetRegions(true)
	a.viewpanel = a.viewpanel.SetChangedFunc(func() {
		a.app.Draw() 
	})

	a.viewpanel.SetBorder(true)
	a.rowContainer = a.rowContainer.AddItem(a.addressBar, 3, 1, true)
	a.rowContainer = a.rowContainer.AddItem(a.viewpanel, 0, 5, false)

	a.container.AddItem(a.rowContainer, 0, 1, true)

}

func (a * Application) Run() {
	if err := a.app.SetRoot(a.container, true).Run(); err != nil {
		panic(err)
	}
}

func (a * Application) searchKey() {
	search := a.addressBar.GetText()
	a.currentAdress = search

	a.viewpanel =  a.viewpanel.SetText(FetchPage(search))

}
