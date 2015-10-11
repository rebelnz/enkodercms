package handlers

import (
	"bitbucket.org/enkdr/enkoder/config"
	m "bitbucket.org/enkdr/enkoder/models"
	p "bitbucket.org/enkdr/enkoder/process"
	"bitbucket.org/enkdr/enkoder/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"strconv"
)

var adminFuncMap = template.FuncMap{
	"unescaped":  func(x string) template.HTML { return template.HTML(x) },
	"fileicon":   utils.FileIcon,
	"formattime": utils.FormatTime,
	"checkit":    utils.CheckIt, // returns false if str doesnt match
}

func adminDataMap(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	// get all data for pages
	account := new(m.Account)
	pages := new(m.Pages)
	posts := new(m.Posts)
	parentPages := new(m.Pages)
	widgets := new(m.Widgets)
	medias := new(m.Medias)
	sliders := new(m.Sliders)
	accounts := new(m.Accounts)

	data := make(map[string]interface{})

	data["Sitename"] = config.Conf().Sitename
	data["Fqdn"] = config.Conf().Fqdn
	data["Account"] = account.GetAccount(w, r)
	data["Pages"] = pages.GetPages()
	data["Posts"] = posts.GetPosts()
	data["ParentPages"] = parentPages.GetTopPages("")
	data["Widgets"] = widgets.GetWidgets() // REFACTOR -- should be able to use common query
	data["Sliders"] = sliders.GetSliders()
	data["Medias"] = medias.GetMedia()
	data["Settings"] = m.GetSettings()
	data["Accounts"] = accounts.GetAccounts(w, r)

	return data
}

func AdminIndex(w http.ResponseWriter, r *http.Request) {
	Data := adminDataMap(w, r)
	subscribers := new(m.Subscribers)
	Data["Subscribers"] = subscribers.GetSubscribers()
	Data["Meta"] = m.Meta{Title: "Admin", Nav: "admin"}
	t := template.Must(template.New("index.html").Funcs(adminFuncMap).ParseFiles("templates/admin/index.html", "templates/admin/admin.html"))
	if err := t.Execute(w, Data); err != nil {
		fmt.Println(err)
		return
	}
}

func AdminGeneric(item string) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		title := fmt.Sprintf("Admin %s", item)
		Data := adminDataMap(w, r)
		Data["Meta"] = m.Meta{Title: title, Nav: item}
		t := template.Must(template.New("index.html").Funcs(adminFuncMap).ParseFiles("templates/admin/index.html", "templates/admin/"+item+".html"))
		t.Execute(w, Data)
	}
	return http.HandlerFunc(fn)
}

func AdminNewPost(w http.ResponseWriter, r *http.Request) {
	Data := adminDataMap(w, r)
	Data["Meta"] = m.Meta{Title: "Admin Post", Nav: "posts"}
	t := template.Must(template.New("index.html").Funcs(adminFuncMap).ParseFiles("templates/admin/index.html", "templates/admin/newpost.html"))
	if err := t.Execute(w, Data); err != nil {
		fmt.Println(err)
		return
	}
}

func AdminNewPostProcess(w http.ResponseWriter, r *http.Request) {
	Data := adminDataMap(w, r)
	Data["Meta"] = m.Meta{Title: "Admin New Post", Nav: "posts"}
	var pid string
	if r.FormValue("pid") != "" {
		pid = r.FormValue("pid")
	}
	message, err := p.NewPostP(w, r, pid)
	if err != nil {
		Data["Message"] = message
		Data["Pid"] = pid
		Data["Title"] = r.FormValue("title")
		Data["Content"] = r.FormValue("content")
		Data["Tags"] = r.FormValue("tags")
		Data["Status"], _ = strconv.ParseInt(r.FormValue("status"), 10, 64)
		t, _ := template.ParseFiles("templates/admin/index.html", "templates/admin/newpost.html")
		if err := t.Execute(w, Data); err != nil {
			fmt.Println(err)
			return
		}
	} else {
		http.Redirect(w, r, "/admin/posts", 302)
	}
	return
}

func AdminEditPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	post := new(m.Post)
	p := post.GetPostById(id) // widget to edit
	Data := adminDataMap(w, r)
	Data["Meta"] = m.Meta{Title: "Admin Edit Post", Nav: "posts"}
	Data["Pid"] = p.Id
	Data["Title"] = p.Title
	Data["Content"] = p.Content
	Data["Filename"] = p.Filename
	Data["Tags"] = p.Tags
	Data["Status"] = p.Status
	t := template.Must(template.New("index.html").Funcs(adminFuncMap).ParseFiles("templates/admin/index.html", "templates/admin/newpost.html"))
	if err := t.Execute(w, Data); err != nil {
		fmt.Println(err)
		return
	}
	return
}

func AdminNewSlider(w http.ResponseWriter, r *http.Request) {
	Data := adminDataMap(w, r)
	Data["Meta"] = m.Meta{Title: "Admin Slider", Nav: "newslider"}
	t, _ := template.ParseFiles("templates/admin/index.html", "templates/admin/newslider.html")
	if err := t.Execute(w, Data); err != nil {
		fmt.Println(err)
		return
	}
}

func AdminNewSliderProcess(w http.ResponseWriter, r *http.Request) {
	Data := adminDataMap(w, r)
	Data["Meta"] = m.Meta{Title: "Admin New Slider", Nav: "newslider"}
	var sid string
	if r.FormValue("sid") != "" {
		sid = r.FormValue("sid")
	}
	message, err := p.NewSliderP(w, r, sid)
	if err != nil {
		Data["Message"] = message
		Data["Sid"] = sid
		Data["Title"] = r.FormValue("title")
		Data["Content"] = r.FormValue("content")
		Data["Url"] = r.FormValue("url")
		t, _ := template.ParseFiles("templates/admin/index.html", "templates/admin/newslider.html")
		if err := t.Execute(w, Data); err != nil {
			fmt.Println(err)
			return
		}
	} else {
		http.Redirect(w, r, "/admin/sliders", 302)
	}
	return
}

func AdminEditSlider(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	slider := new(m.Slider)
	sid := slider.GetSliderById(id) // slider to edit
	Data := adminDataMap(w, r)
	Data["Meta"] = m.Meta{Title: "Admin Edit Slider", Nav: "newslider"}
	Data["Sid"] = sid.Id
	Data["Title"] = sid.Title
	Data["Content"] = sid.Content
	Data["Filename"] = sid.Filename
	Data["Url"] = sid.Url
	t := template.Must(template.New("index.html").Funcs(adminFuncMap).ParseFiles("templates/admin/index.html", "templates/admin/newslider.html"))
	if err := t.Execute(w, Data); err != nil {
		fmt.Println(err)
		return
	}
	return
}

func AdminNewAccount(w http.ResponseWriter, r *http.Request) {
	Data := adminDataMap(w, r)
	Data["Meta"] = m.Meta{Title: "Admin New Account", Nav: "newaccount"}
	t, _ := template.ParseFiles("templates/admin/index.html", "templates/admin/newaccount.html")
	if err := t.Execute(w, Data); err != nil {
		fmt.Println(err)
		return
	}
}

func AdminNewAccountProcess(w http.ResponseWriter, r *http.Request) {
	Data := adminDataMap(w, r)
	Data["Meta"] = m.Meta{Title: "Admin New Account", Nav: "newaccount"}
	var accid string
	if r.FormValue("aid") != "" {
		accid = r.FormValue("aid")
	}
	processed, message := p.NewAccountP(w, r, accid) // ./process.go an empty var here doesnt matter
	if processed == true {                           // no errors
		http.Redirect(w, r, "/admin/accounts", 302)
	} else { // re-populate form fields and show message(s)
		Data["Aid"] = accid
		Data["Message"] = message
		Data["Firstname"] = r.FormValue("firstname")
		Data["Lastname"] = r.FormValue("lastname")
		Data["Email"] = r.FormValue("email")
		t, _ := template.ParseFiles("templates/admin/index.html", "templates/admin/newaccount.html")
		if err := t.Execute(w, Data); err != nil {
			fmt.Println(err)
			return
		}
	}
	return
}

func AdminEditAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	account := new(m.Account)
	acc := account.GetAccountById(w, r, id) // account to edit
	Data := adminDataMap(w, r)
	Data["Meta"] = m.Meta{Title: "Admin Edit Account", Nav: "newaccount"}
	Data["Aid"] = id
	Data["Firstname"] = acc.Firstname
	Data["Lastname"] = acc.Lastname
	Data["Email"] = acc.Email
	t, _ := template.ParseFiles("templates/admin/index.html", "templates/admin/newaccount.html")
	if err := t.Execute(w, Data); err != nil {
		fmt.Println(err)
		return
	}
	return
}

func AdminPages(w http.ResponseWriter, r *http.Request) {
	Data := adminDataMap(w, r)
	// declare TopPages as slice of pages
	// for each toppage get child pages (where page id == parent)
	// add it to topPages.Subpages
	var TopPages []m.Page
	tp := new(m.Pages)
	topPages := tp.GetAllTopPages()
	for _, topPage := range topPages {
		sp := new(m.Pages)
		subPages := sp.GetSubPages(topPage.Id)
		for _, subPage := range subPages {
			topPage.Subpages = append(topPage.Subpages, subPage)
		}
		// before we leave the loop
		// append to TopPages (slice of Pages declared above)
		TopPages = append(TopPages, topPage)
	}
	Data["TopPages"] = TopPages
	Data["SpecPages"] = []string{"contact", "home", "map"} // dont show delete link
	Data["Meta"] = m.Meta{Title: "Admin Pages", Nav: "pages"}
	t := template.Must(template.New("index.html").Funcs(adminFuncMap).ParseFiles("templates/admin/index.html", "templates/admin/pages.html"))
	if err := t.Execute(w, Data); err != nil {
		fmt.Println(err)
		return
	}
}

func AdminNewPage(w http.ResponseWriter, r *http.Request) {
	// parent drop down uses parent pages for now (subpages only 1 level deep)
	Data := adminDataMap(w, r)
	Data["Meta"] = m.Meta{Title: "Admin New Page", Nav: "newpage"}
	t, _ := template.ParseFiles("templates/admin/index.html", "templates/admin/newpage.html")
	if err := t.Execute(w, Data); err != nil {
		fmt.Println(err)
		return
	}
}

func AdminNewPageProcess(w http.ResponseWriter, r *http.Request) {
	Data := adminDataMap(w, r)
	Data["Meta"] = m.Meta{Title: "Admin New Page", Nav: "newpage"}
	pid := ""
	if r.FormValue("pid") != "" {
		pid = r.FormValue("pid")
	}
	processed, message := p.NewPageP(w, r, pid) // ./process.go
	if processed == true {                      // no errors
		http.Redirect(w, r, "/admin/pages", 302)
	} else { // re-populate form fields and show message(s)
		Data["Pid"] = pid
		Data["Message"] = message
		Data["Title"] = r.FormValue("title")
		Data["Content"] = r.FormValue("content")
		Data["Status"] = r.FormValue("status")
		Data["Parent"] = r.FormValue("parent")
		t, _ := template.ParseFiles("templates/admin/index.html", "templates/admin/newpage.html")
		if err := t.Execute(w, Data); err != nil {
			fmt.Println(err)
			return
		}
	}
	return
}

func AdminEditPage(w http.ResponseWriter, r *http.Request) {
	Data := adminDataMap(w, r)
	vars := mux.Vars(r)
	slug := vars["slug"]
	page := new(m.Page)
	p := page.GetPage(w, r, slug)
	subW := []string{}
	for _, v := range *p.Subwidgets {
		subW = append(subW, v.Title)
	}
	Data["Meta"] = m.Meta{Title: "Admin Edit Page", Nav: "newpage"}
	Data["Title"] = p.Title
	Data["Content"] = p.Content
	Data["Metatags"] = p.Metatags
	Data["Status"] = p.Status
	Data["Parent"] = p.Parent
	Data["Subwidgets"] = p.Subwidgets
	Data["Pid"] = p.Id
	Data["SubW"] = subW
	t := template.Must(template.New("index.html").Funcs(adminFuncMap).ParseFiles("templates/admin/index.html", "templates/admin/editpage.html"))
	if err := t.Execute(w, Data); err != nil {
		fmt.Println(err)
		return
	}
	return
}

func AdminNavigation(w http.ResponseWriter, r *http.Request) {
	Data := adminDataMap(w, r)
	Data["TopNavPages"] = m.GetNavPages("topnav")
	Data["SideNavPages"] = m.GetNavPages("sidenav")
	Data["FooterNavPages"] = m.GetNavPages("footernav")
	Data["Meta"] = m.Meta{Title: "Admin Settings", Nav: "navigation"}
	t, _ := template.ParseFiles("templates/admin/index.html", "templates/admin/navigation.html")
	if err := t.Execute(w, Data); err != nil {
		fmt.Println(err)
		return
	}
}

func AdminNewWidgetProcess(w http.ResponseWriter, r *http.Request) {
	Data := adminDataMap(w, r)
	Data["Meta"] = m.Meta{Title: "Admin New Widget", Nav: "newwidget"}
	var wid string
	if r.FormValue("wid") != "" {
		wid = r.FormValue("wid")
	}
	message, err := p.NewWidgetP(w, r, wid)
	if err != nil {
		Data["Message"] = message
		Data["Wid"] = wid
		Data["Title"] = r.FormValue("title")
		Data["Content"] = r.FormValue("content")
		Data["Url"] = r.FormValue("url")
		Data["Widgetsize"] = r.FormValue("widgetsize")
		t, _ := template.ParseFiles("templates/admin/index.html", "templates/admin/newwidget.html")
		if err := t.Execute(w, Data); err != nil {
			fmt.Println(err)
			return
		}
	} else {
		http.Redirect(w, r, "/admin/widgets", 302)
	}
	return
}

func AdminNewWidget(w http.ResponseWriter, r *http.Request) {
	Data := adminDataMap(w, r)
	Data["Meta"] = m.Meta{Title: "Admin New Widget", Nav: "newwidget"}
	Data["Widgetsize"] = "" // to prevent error - .Widgetsize value is needed for editwidget - same template
	t := template.Must(template.New("index.html").Funcs(adminFuncMap).ParseFiles("templates/admin/index.html", "templates/admin/newwidget.html"))
	if err := t.Execute(w, Data); err != nil {
		fmt.Println(err)
		return
	}
}

func AdminEditWidget(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	widget := new(m.Widget)
	wid := widget.GetWidgetById(id) // widget to edit
	Data := adminDataMap(w, r)
	Data["Meta"] = m.Meta{Title: "Admin Edit Widget", Nav: "newwidget"}
	Data["Wid"] = wid.Id
	Data["Title"] = wid.Title
	Data["Content"] = wid.Content
	Data["Filename"] = wid.Filename
	Data["Url"] = wid.Url
	Data["Widgetsize"] = wid.Widgetsize
	t := template.Must(template.New("index.html").Funcs(adminFuncMap).ParseFiles("templates/admin/index.html", "templates/admin/newwidget.html"))
	if err := t.Execute(w, Data); err != nil {
		fmt.Println(err)
		return
	}
	return
}

func AdminNewMedia(w http.ResponseWriter, r *http.Request) {
	Data := adminDataMap(w, r)
	Data["Meta"] = m.Meta{Title: "Admin Upload", Nav: "upload"}
	t, _ := template.ParseFiles("templates/admin/index.html", "templates/admin/newmedia.html")
	if err := t.Execute(w, Data); err != nil {
		fmt.Println(err)
		return
	}
}

func AdminNewMediaProcess(w http.ResponseWriter, r *http.Request) {
	Data := adminDataMap(w, r)
	Data["Meta"] = m.Meta{Title: "Admin Upload", Nav: "upload"}
	message, err := p.NewMediaP(w, r)
	if err != nil {
		Data["Title"] = r.FormValue("title")
		Data["Message"] = message
	} else {
		http.Redirect(w, r, "/admin/medias", 302)
	}
	t, _ := template.ParseFiles("templates/admin/index.html", "templates/admin/newmedia.html")
	if err := t.Execute(w, Data); err != nil {
		fmt.Println(err)
		return
	}
}

func AdminDeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	item := vars["item"]
	id := vars["id"]
	err := m.DeleteItem(w, r, item, id)
	if err != nil {
		fmt.Println("item could not be deleted")
	}
	url := fmt.Sprintf("/admin/%ss", item) // redirect to %s page
	// unless item == account then redirect to login page - delete session
	http.Redirect(w, r, url, 302)
	return
}

// ajax update nav order
func AdminUpdateNavOrder(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	navtype := r.FormValue("navtype")
	navorder := r.Form["navorder[]"]
	for k, v := range navorder {
		k++
		v, _ := strconv.Atoi(v)
		m.UpdateNavOrder(k, v, navtype)
	}
	return
}

// ajax get/update map coords
func AdminGetMap(w http.ResponseWriter, r *http.Request) {
	coords := m.GetMap()
	json.NewEncoder(w).Encode(coords)
	return
}

func AdminUpdateMap(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	lat := r.FormValue("latitude")
	lon := r.FormValue("longitude")
	m.UpdateMap(lat, lon)
	return
}

// ajax update page status
func AdminUpdatePageStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	status, _ := strconv.Atoi(vars["status"])
	pageid, _ := strconv.Atoi(vars["pageid"])
	m.UpdatePageStatus(status, pageid)
	url := "/admin/pages"
	http.Redirect(w, r, url, 302)
	return
}

// ajax get/update settings
func AdminGetSettings(w http.ResponseWriter, r *http.Request) {
	settings := m.GetSettings()
	json.NewEncoder(w).Encode(settings)
	return
}

func AdminUpdateSettings(w http.ResponseWriter, r *http.Request) {
	m.UpdateSettings(r)
	return
}
