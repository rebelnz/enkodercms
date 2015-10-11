package handlers

import (
	"bitbucket.org/enkdr/enkoder/config"
	"bitbucket.org/enkdr/enkoder/middleware"
	m "bitbucket.org/enkdr/enkoder/models"
	p "bitbucket.org/enkdr/enkoder/process"
	u "bitbucket.org/enkdr/enkoder/utils"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"io/ioutil"
	"net/http"
	"path"
	"time"
)

var funcMap = template.FuncMap{
	"unescaped":   func(x string) template.HTML { return template.HTML(x) },
	"fileicon":    u.FileIcon,
	"formattime":  u.FormatTime,
	"widgetcss":   u.WidgetCss,
	"postsummary": u.PostSummary,
	"postcontent": u.PostContent,
}

func dataMap(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	posts := new(m.Posts)

	data := make(map[string]interface{})
	data["TopNavPages"] = m.GetNavPages("topnav")
	data["Posts"] = posts.GetPosts(1)
	// data["SideNavPages"] = getNavPages("sidenav") 	// if needed.. design specific
	data["FooterNavPages"] = m.GetNavPages("footernav")
	account := new(m.Account)

	data["Account"] = account.GetAccount(w, r)
	data["Settings"] = m.GetSettings()
	data["Sitename"] = config.Conf().Sitename
	data["Fqdn"] = config.Conf().Fqdn
	tFormat := "2006"
	data["Time"] = time.Now().Format(tFormat)
	return data
}

// how to pass in args to handler -- see routes for arg example
func Index(format string) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		sliders := new(m.Sliders)
		page := new(m.Page)
		p := page.GetPage(w, r, "home")

		Data := dataMap(w, r)
		Data["Sliders"] = sliders.GetSliders()
		Data["Meta"] = m.Meta{Title: config.Conf().Sitename}
		Data["Page"] = p
		Data["Subwidgets"] = p.Subwidgets

		t := template.Must(template.New("index.html").Funcs(funcMap).ParseFiles("templates/index.html", "templates/home.html"))
		t.Execute(w, Data)
	}
	return http.HandlerFunc(fn)
}

func Login(w http.ResponseWriter, r *http.Request) {
	Data := dataMap(w, r)
	Data["Meta"] = m.Meta{Title: "Login", Nav: "login"}
	t, _ := template.ParseFiles("templates/index.html", "templates/login.html")
	if err := t.Execute(w, Data); err != nil {
		fmt.Println(err)
		return
	}
}

func LoginProcess(w http.ResponseWriter, r *http.Request) {
	Data := dataMap(w, r)
	Data["Meta"] = m.Meta{Title: "Login", Nav: "login"}
	message := p.LoginP(w, r) // ./process.go
	if len(message) == 0 {    // no errors
		http.Redirect(w, r, "/admin", 302)
	} else {
		Data["Message"] = message
		t, _ := template.ParseFiles("templates/index.html", "templates/login.html")
		if err := t.Execute(w, Data); err != nil {
			fmt.Println(err)
			return
		}
	}
	return
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := middleware.Store.Get(r, config.Conf().SessionName)
	delete(session.Values, "UserID")
	delete(session.Values, "Authorized")
	session.Options.MaxAge = -1
	_ = session.Save(r, w)
	http.Redirect(w, r, "/login", 302)
	return
}

func ContactProcess(w http.ResponseWriter, r *http.Request) {
	settings := m.GetSettings()
	r.ParseForm()
	email := r.FormValue("email")
	message := r.FormValue("message")
	subject := fmt.Sprintf("Website Enquiry from: %s", email)
	w.Header().Set("Content-Type", "application/json")
	u.SendMail("info@enkoder.com.au", settings.Email, subject, message)
	fmt.Fprint(w, u.JSONResponse{"success": true, "message": "Email Sent"})
	return
}

func ResetProcess(w http.ResponseWriter, r *http.Request) {
	response, err := p.ResetPasswordP(w, r)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, u.JSONResponse{"success": true, "message": response})
	return
}

func ShowPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["slug"]
	if vars["slug"] == "home" { // use Index fumc for homepage
		http.Redirect(w, r, "/", 302)
		return
	}
	page := new(m.Page)
	p := page.GetPage(w, r, slug)

	if p.Status != 1 { // either page is draft or does not exist...
		NotFound(w, r)
		return
	}

	Data := dataMap(w, r)
	Data["Page"] = p
	Data["Subwidgets"] = p.Subwidgets
	Data["Meta"] = m.Meta{Title: p.Title, Nav: p.Slug}

	// if templates/slug.html exists -- use that otherwise default tmpl is page
	pageTmpl := "page"
	tmpls, _ := ioutil.ReadDir("templates")
	for _, t := range tmpls {
		if t.Name() == slug+".html" {
			pageTmpl = slug
		}
	}
	t := template.Must(template.New("index.html").Funcs(funcMap).ParseFiles("templates/index.html", "templates/"+pageTmpl+".html"))
	if err := t.Execute(w, Data); err != nil {
		fmt.Println(err)
		return
	}
}

func ShowPosts(w http.ResponseWriter, r *http.Request) {
	Data := dataMap(w, r)
	Data["Meta"] = m.Meta{Title: "News", Nav: "news"}
	t := template.Must(template.New("index.html").Funcs(funcMap).ParseFiles("templates/index.html", "templates/posts.html"))
	if err := t.Execute(w, Data); err != nil {
		fmt.Println(err)
		return
	}
}

func ShowPost(w http.ResponseWriter, r *http.Request) {
	Data := dataMap(w, r)
	vars := mux.Vars(r)
	slug := fmt.Sprintf("news/%s/%s", vars["date"], vars["slug"])
	post := new(m.Post)
	p := post.GetPost(w, r, slug)
	if p.Status != 1 { // either post is draft or does not exist...
		NotFound(w, r)
		return
	}
	Data["Post"] = p
	Data["Meta"] = m.Meta{Title: p.Title, Nav: p.Slug}

	// funcMap := template.FuncMap{
	// 	"unescaped":  func(x string) template.HTML { return template.HTML(x) },
	// 	"fileicon":   u.FileIcon,
	// 	"formattime": u.FormatTime,
	// }

	t := template.Must(template.New("index.html").Funcs(funcMap).ParseFiles("templates/index.html", "templates/post.html"))
	if err := t.Execute(w, Data); err != nil {
		fmt.Println(err)
		return
	}
}

func SearchProcess(w http.ResponseWriter, r *http.Request) {
	Data := dataMap(w, r)
	Data["Meta"] = m.Meta{Title: "Search"}
	_, results := p.SearchP(w, r)
	Data["Results"] = results
	t, _ := template.ParseFiles("templates/index.html", "templates/search.html")
	if err := t.Execute(w, Data); err != nil {
		fmt.Println(err)
		return
	}
}

func SubscribeProcess(w http.ResponseWriter, r *http.Request) {
	Data := dataMap(w, r)
	Data["Meta"] = m.Meta{Title: "Subscribe"}
	message, err := p.SubscribeP(w, r)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprint(w, u.JSONResponse{"message": message})
	return
}

func UnSubscribeProcess(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	randomid, _ := vars["randomid"]
	message, _ := p.UnSubscribeP(randomid)
	Data := dataMap(w, r)
	Data["Message"] = message
	Data["Meta"] = m.Meta{Title: "Unsubscribe"}
	t, _ := template.ParseFiles("templates/index.html", "templates/unsubscribe.html")
	if err := t.Execute(w, Data); err != nil {
		fmt.Println(err)
		return
	}
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Page Not Found\n")
	Data := make(map[string]interface{})
	Data["Meta"] = m.Meta{Title: "Page Not Found"}
	t, _ := template.ParseFiles("templates/index.html", "templates/404.html")
	if err := t.Execute(w, Data); err != nil {
		fmt.Println(err)
		return
	}
}

func ServeFileHandler(res http.ResponseWriter, req *http.Request) {
	fname := path.Base(req.URL.Path)
	http.ServeFile(res, req, "static/"+fname)
}
