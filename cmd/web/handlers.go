package main

import (
	"fmt"
	// "html/template"
	"net/http"
	"strconv"

	"github.com/jbresky/snippetbox/pkg/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	s, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	// for _, snippet := range s {
	// 	fmt.Fprintf(w, "%v\n", snippet)
	// }
	// data := &templateData{Snippets: s}

	app.render(w, r, "home.page.html", &templateData{
		Snippets: s,
	})

	// files := []string{
	// 	"./ui/html/home.page.html",
	// 	"./ui/html/base.layout.html",
	// 	"./ui/html/footer.partial.html",
	// }

	// ts, err := template.ParseFiles(files...)

	// if err != nil {
	// 	app.errorLog.Println(err.Error())
	// 	app.serverError(w, err)
	// 	return
	// }

	// err = ts.Execute(w, data)

	// if err != nil {
	// 	app.errorLog.Println(err.Error())
	// 	app.serverError(w, err)
	// 	return
	// }
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	s, err := app.snippets.Get(id)
	if err == models.ErrNoRecord {
		app.notFound(w)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "show.page.html", &templateData{
		Snippet: s,
	})

	// data := &templateData{Snippet: s}

	// files := []string{
	// 	"./ui/html/show.page.html",
	// 	"./ui/html/base.layout.html",
	// 	"./ui/html/footer.partial.html",
	// }

	// ts, err := template.ParseFiles(files...)

	// if err != nil {
	// 	app.errorLog.Println(err.Error())
	// 	app.serverError(w, err)
	// 	return
	// }

	// err = ts.Execute(w, data)
	// if err != nil {
	// 	app.serverError(w, err)
	// }

	// fmt.Fprintf(w, "%v", s)
}

func (app *application) createSnippetForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create.page.html", nil)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi"
	expires := "7"

	id, err := app.snippets.Insert(title, content, expires)

	if err != nil {
		app.serverError(w, err)
		return
	}

	// Change the redirect to use the new semantic URL style of /snippet/:id
	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)

}
