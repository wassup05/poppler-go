package tests

import (
	"testing"

	"github.com/wassup05/poppler-go/poppler"
)

func TestNewDoc(t *testing.T) {
	t.Run("new single page document from path", func(t *testing.T) {
		doc, err := poppler.NewDocFromFilename(cwd + "/test.pdf")
		if err != nil {
			t.Fatalf("Could not open document test.pdf... error: %v\n", err.Message)
		}

		t.Cleanup(func() {
			doc.Close()
		})
	})

	t.Run("new multipage document from path", func(t *testing.T) {
		doc, err := poppler.NewDocFromFilename(cwd + "/test_multipage.pdf")
		if err != nil {
			t.Fatalf("Could not open document test_multipage.pdf... error: %v\n", err.Message)
		}

		t.Cleanup(func() {
			doc.Close()
		})
	})

	t.Run("new document from path with password", func(t *testing.T) {
		doc, err := poppler.NewDocFromFilenameWithPass(cwd+"/password_test123.pdf", "test123")
		if err != nil {
			t.Fatalf("Could not open document password_test123.pdf... error: %v\n", err.Message)
		}

		t.Cleanup(func() {
			doc.Close()
		})
	})
}

func TestInfo(t *testing.T) {
	t.Run("getting info of a document", func(t *testing.T) {
		doc, err := poppler.NewDocFromFilename(cwd + "/hello_world.pdf")

		t.Cleanup(func() {
			doc.Close()
		})

		if err != nil {
			t.Fatalf("Could not open document hello_word.pdf... error:%v\n", err.Message)
		}

		const (
			author   string = "johndoe"
			producer        = "producer"
			title           = "title"
			creator         = "me"
			subject         = "subject"
			keywords        = "keywords"
		)

		doc.SetTitle(title)
		doc.SetAuthor(author)
		doc.SetProducer(producer)
		doc.SetCreator(creator)
		doc.SetKeywords(keywords)
		doc.SetSubject(subject)

		info := doc.GetInfo()

		if info == nil {
			t.Fatalf("Could not get info\n")
		}
	})
}
