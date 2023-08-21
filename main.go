package main

import (
	"fmt"
	htgotts "github.com/hegedustibor/htgo-tts"
	"github.com/hegedustibor/htgo-tts/handlers"
	"github.com/hegedustibor/htgo-tts/voices"
	"github.com/therecipe/qt/widgets"
	"os"
)

/*var TextToSpeak = "GPT4All is an ecosystem to train and deploy powerful and customized" +
" large language models that run locally on consumer grade CPUs." /* +
"The goal is simple - be the best instruction tuned assistant-style language model that" +
" any person or enterprise can freely use, distribute and build on. " +
"A GPT4All model is a 3GB - 8GB file that you can download and plug into the GPT4All" +
" open-source ecosystem software. Nomic AI supports and maintains this software " +
"ecosystem to enforce quality and security alongside spearheading the effort to " +
"allow any person or enterprise to easily train and deploy their own on-edge large language models."*/
/*
var (
	centralLayout       *widgets.QGridLayout
	centralLayoutRow    int
	centralLayoutColumn int
)*/

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)
	w := widgets.NewQWidget(nil, 0)
	wL := widgets.NewQVBoxLayout2(w)
	//ic := widgets.NewQHBoxLayout()
	iw := widgets.NewQTextEdit(w)
	btn := widgets.NewQPushButton2("Speak", w)
	wL.AddWidget(iw, 0, 0)
	wL.AddWidget(btn, 0, 0)
	btn.ConnectClick(Speak())
	//wL.AddLayout(wL, 0)
	/*for i := 0; i < 25; i++ {
		l := widgets.NewQHBoxLayout()
		for j := 0; j < 10; j++ {
			l.AddWidget(widgets.NewQLabel2(fmt.Sprint(i, ":", j), w, 0), 0, 0)
		}
		wL.AddLayout(l, 0)
	}*/
	w.Show()
	widgets.QApplication_Exec()
}

/*func addWidget(widget widgets.QWidget_ITF) {

	if centralLayoutColumn > 6 {
		centralLayoutColumn = 0
		centralLayoutRow++
	}

	wrappedWidget := widgets.NewQGroupBox2(widget.QWidget_PTR().WindowTitle(), nil)
	wrappedWidgetLayout := widgets.NewQVBoxLayout2(wrappedWidget)
	wrappedWidgetLayout.AddWidget(widget, 0, core.Qt__AlignCenter)
	wrappedWidget.SetFixedSize2(250, 250)

	centralLayout.AddWidget2(wrappedWidget, centralLayoutRow, centralLayoutColumn, core.Qt__AlignCenter)

	centralLayoutColumn++
}*/

func Speak() {

	speech := htgotts.Speech{Folder: "audio", Language: voices.English, Handler: &handlers.Native{}}
	if err := speech.Speak("Hello"); err != nil {
		fmt.Printf("ERROR: %s", err)
		//return err
	}
	//return nil
}
