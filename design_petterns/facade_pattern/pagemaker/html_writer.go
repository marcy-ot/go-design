package pagemaker

import "os"

type HtmlWriter struct {
	writer *os.File
}

func NewHtmlWriter(file *os.File) HtmlWriter {
	return HtmlWriter{
		writer: file,
	}
}

func (h HtmlWriter) Title(title string) {
	h.writer.WriteString("<html>")
	h.writer.WriteString("<head>")
	h.writer.WriteString("<title>" + title + "</title>")
	h.writer.WriteString("</head>")
	h.writer.WriteString("<body>\n")
	h.writer.WriteString("<h1>" + title + "</h1>\n")
}

func (h HtmlWriter) Paragraph(msg string) {
	h.writer.WriteString("<p>" + msg + "</p>\n")
}

func (h HtmlWriter) Link(href, caption string) {
	h.Paragraph("<a href=\"" + href + "\">" + caption + "</a>")
}

func (h HtmlWriter) Mailto(mailAddr, userName string) {
	h.Link("mailto:"+mailAddr, userName)
}

func (h HtmlWriter) Close() {
	h.writer.WriteString("</body>")
	h.writer.WriteString("</html>\n")
}
