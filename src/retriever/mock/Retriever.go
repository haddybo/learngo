package mock

import "fmt"

type Retriever struct {
	Content string
}

func (r *Retriever) String() string {
	return fmt.Sprintf("Retriever:{Contents=%s}", r.Content)
}

func (r *Retriever) Post(url string, form map[string]string) string {
	r.Content = form["contents"]
	return "ok"
}

func (r *Retriever) Get(rul string) string {
	return r.Content
}

