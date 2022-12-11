package pagemaker

import (
	"fmt"
	"os"
)

var PageMaker = pageMaker{}

type pageMaker struct{}

func (p pageMaker) MakeWelcomePage(mailAddr, fileName string) {
	mailProps := Database.GetProperties("maildata")

	username := mailProps.GetPropaty(mailAddr)

	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	htmlWriter := NewHtmlWriter(file)

	htmlWriter.Title("welcome to " + username + "'s page!")
	htmlWriter.Paragraph(username + "のページへようこそ。")
	htmlWriter.Paragraph("メール待っていますね。")
	htmlWriter.Mailto(mailAddr, username)
	htmlWriter.Close()

	fmt.Println("complete")
}
