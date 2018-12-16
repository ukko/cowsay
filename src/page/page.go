package page

type Page struct {
	PageTitle       string
	PageAuthor      string
	PageDescription string
	PageContent     string
	PageGenerated   string
}

func New() *Page {
	return &Page{
		PageTitle:       "Fortune",
		PageAuthor:      "Max Kamashev",
		PageDescription: "Cow say fortunes",
	}
}
