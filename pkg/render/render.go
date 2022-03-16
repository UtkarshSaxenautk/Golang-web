package render

import (
	"bytes"
	"github.com/utkarshsaxenautk/pkg/config"
	"github.com/utkarshsaxenautk/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

var app *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	app = a
}
func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	/*tc, err := CreateTemplateCache()
	if err != nil {
		log.Fatal("Error in Creating Cache template : ", err)
	}*/
	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		var er error
		tc, er = CreateTemplateCache()
		if er != nil {
			log.Fatal("Problem in creating template : ", er)
		}
	}

	t, exist := tc[tmpl]
	if !exist {
		log.Fatal(tmpl, " doesnt exist : ", exist)

	}
	buf := new(bytes.Buffer)
	td = AddDefaultData(td)
	_ = t.Execute(buf, td)
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Fatal("Error in writing template to browser : ", err)
	}
	/*_, err := CreateTemplateCache()
	if err != nil {
		log.Fatal("Error getting cache : ", err)
	}*/

	/*parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Fatal("Error in Parsing : ", err)
		return
	}*/
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	mycache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return mycache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		//log.Println("Page is currently : ", page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return mycache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.tmpl")

		if err != nil {
			return mycache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return mycache, err
			}
			//log.Println(ts)
		}
		mycache[name] = ts
	}
	return mycache, nil
}
