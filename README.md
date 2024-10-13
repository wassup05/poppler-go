# poppler-go

An easy to use and intuitive PDF library using Go and Poppler

## Features

* Implementation of many poppler-glib public C API methods such as
    * Loading Documents (encrypted as well) (from path, os.File, bytes)
    * Getting Document Info such as metadata, page count, attachment count, PDF version, permissions, keywords, subject, title, author, creation date, modification date etc
    * Setting Document Info such as metadata, title, author, subject, keywords, title, author, craetion date, modification date and many more
    * Getting Pages from a Document by their index or label.
        * Getting PageInfo such as page text, size, thumbnail, images, attachments, annotations and so on
        * Searching text, getting text from a certain region, selecting text etc
        * Rendering a Page to Cairo surfaces for printing and viewing.
        * Getting Media on a page such as videos and saving them.
    * Dealing with forms.
    * Getting all the info about fonts, layers, index in a document.
    * Saving the modified Document at a path
    * Rendering to a PostScript File

* Integration with [go-cairo](https://github.com/ungerik/go-cairo)

# Poppler
[Poppler](https://poppler.freedesktop.org/) is a library for rendering PDF files and examining, modifying their structure.
This repository contains Go Bindings for the stable C api of poppler (glib) to examine and modify the structure of PDF files and render it's pages to [Cairo](https://www.cairographics.org/) contexts.

## Why?
I once thought about making a tui app for viewing PDF's in go and hence went on a search for PDF rendering, manipulation libraries and saw that poppler was one of the many choices available... But it's go bindings [cheggaaa/go-poppler](https://github.com/cheggaaa/go-poppler) and multitude of it's forks although quite suitable for simple tasks was not as complete as it's other language counterparts, So I thought about making a more complete version of it.

Although there are some edges here and there, I plan on making them as smooth as possible.

# Requirements

* [Go](https://go.dev/)
* poppler library
* cairo library

For Debian/Ubuntu
```bash
sudo apt install libpoppler-dev libcairo2-dev
```

For Fedora
```bash
sudo dnf install poppler-glib-devel cairo-devel
```

For Arch
```bash
sudo pacman -S poppler poppler-data poppler-glib cairo
```

Official source: [poppler](https://poppler.freedesktop.org/), [cairo](https://www.cairographics.org/)

## Examples of common tasks

Run the following command to add the library to your `go.mod` and look at the examples below 
```bash
go get github.com/wassup05/poppler-go
```

* Dealing with PDF metadata

```go
package demo

import (
    "fmt"
    "github.com/wassup05/poppler-go"
)

func main(){
    path := "/path/to/pdf"
    
    doc, err := poppler.NewDocFromFilename(path)
    defer doc.Close()

    if err != nil {
        fmt.Printf("Error opening PDF: %s\n",err.Message)
    }

    fmt.Printf("page count:%d\n",doc.GetPageCount())
    fmt.Printf("author:%s\n",doc.GetAuthor())
    fmt.Printf("creator:%s\n",doc.GetCreator())
    fmt.Printf("keywords:%s\n",doc.GetKeywords())

    doc.SetTitle("blahblah")

    fmt.Printf("title:%s\n",doc.GetTitle()) // blahblah

    info := doc.GetInfo()

    fmt.Printf("info:%#v\n",info) // a convinience method to get all the useful info in one go as a struct
}
```

* Getting the pages of a document and it's info

```go
package demo

import (
    "fmt"
    "github.com/wassup05/poppler-go"
)

func main(){
    path := "/path/to/pdf"
    
    doc, err := poppler.NewDocFromFilename(path)
    defer doc.Close()

    if err != nil {
        fmt.Printf("Error opening PDF: %s\n",err.Message)
    }

    page := doc.GetPage(2)
    defer page.Close()

    label := page.GetLabel() // gets the label of the page
    height, width := page.GetSize() // gets the height, width of the page
    text := page.GetText() // gets all the text on the page

    fmt.Printf("label:%s\nheight,width: %f, %f\ntext: %s\n",label,height,width,text)

    pageInfo := page.GetInfo() // convinience method

    fmt.Printf("Page Info:%#v\n",pageInfo)
}

```

* Rendering a PDF page to a SVG

```go
package demo

import (
    "github.com/ungerik/go-cairo"
    "github.com/wassup05/poppler-go"
)

func main(){
	doc, err := poppler.NewDocFromFilename("abs/path/to/pdf")
    defer doc.Close()

	if err != nil {
        fmt.Printf("Error opening PDF: %s\n",err.Message)
	}

	page := doc.GetPage(0) // get the first page
    defer page.Close()

	h, w := page.GetSize()

	surface := cairo.NewSVGSurface("test.svg", w, h, cairo.SVG_VERSION_1_2)

	page.RenderToSurface(surface)

	surface.Destroy() // a svg with name `test.svg` with the contents of the 1st page will be created
}
```

* Dealing with attachments

```go
package demo

import (
    "fmt"
    "github.com/wassup05/poppler-go"
)

func main(){
	doc, err := poppler.NewDocFromFilename("abs/path/to/pdf")
    defer doc.Close()

	if err != nil {
        fmt.Printf("Error opening PDF: %s\n",err.Message)
	}

    attachments := doc.GetAttachments()

    for i, a := range attachments {
        fmt.Printf("name:%s\n",a.GetName())
        fmt.Printf("description:%s\n",a.GetDescription())

        saveErr := a.Save(fmt.Sprintf("demo-%d",i)) // saves the attachment to the path

        if saveErr != nil {
            // handle the error
        }
    }
}
```

## Missing Features that ARE present in poppler

* Advanced Interaction with form fields, annotations and page structures, Which will be implemented soon
* Extensive Documentation

# Credits
* [go-pdfium](https://github.com/klippa-app/go-pdfium) - PDF's for testing
* [go-poppler](https://github.com/cheggaaa/go-poppler) - For inspiartion.

# Contributing and bug reports.
All contributions are very much welcome, For Bug Reports open an Issue with enough information so that they are debuggable. 

There are many functions which are quite easy to implement, documentation for which can be found [here](ihttps://poppler.freedesktop.org/api/glib/) please go through them and send the contributions via a pull request.

For functions which are very common while handling PDF's I have made a separate file [extras.go](./poppler/extras.go) which can contain functions that are not necessarily provided by poppler but a combination of many functions that will make our lives a lot easier. You can contribute to these as well.

# License

GNU General Public License, Version 3.0 (GPL v3.0)
