package models

import (
	"bitbucket.org/enkdr/enkoder/config"
	"bitbucket.org/enkdr/enkoder/dbconn"
	"bitbucket.org/enkdr/enkoder/middleware"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"path"
	"time"
)

var db = dbconn.NewDB()

// Get account by session
func (a Account) GetAccount(w http.ResponseWriter, r *http.Request) Account {
	session, _ := middleware.Store.Get(r, config.Conf().SessionName)
	userid, _ := session.Values["UserID"]
	err := db.QueryRowx(`SELECT * FROM account WHERE id = $1`, userid).StructScan(&a)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err)
		http.Redirect(w, r, "/login", 302) // no user or any other error.. kick out
	}
	return a
}

func (a Account) GetAccountById(w http.ResponseWriter, r *http.Request, id string) Account {
	err := db.QueryRowx(`SELECT * FROM account WHERE id = $1`, id).StructScan(&a)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err)
	}
	return a
}

func (as Accounts) GetAccounts(w http.ResponseWriter, r *http.Request) Accounts {
	err := db.Select(&as, "SELECT * FROM account")
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err)
	}
	return as
}

func (p Page) GetPage(w http.ResponseWriter, r *http.Request, slug string) Page {
	err := db.QueryRowx(`SELECT * FROM page WHERE slug = $1`, slug).StructScan(&p)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err)
	}
	ws := new(Widgets)
	err = db.Select(ws, `SELECT w.* FROM page_widget pw INNER JOIN page p ON pw.page_id = p.id INNER JOIN widget w ON pw.widget_id = w.id WHERE p.id = $1`, p.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println(err)
		}
		fmt.Println("THIS:", err)
	}
	p.Subwidgets = ws
	return p
}

func (ps Pages) GetPages(status ...int) Pages {
	if len(status) > 0 {
		s := status[0]
		// err := db.Select(&ps, "SELECT * FROM page WHERE status = $1 ORDER BY id", s)
		err := db.Select(&ps, "SELECT * FROM page WHERE status = $1 order by id", s)
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Println(err)
			}
			fmt.Println(err)
		}
	} else {
		err := db.Select(&ps, "SELECT * FROM page ORDER BY id")
		if err != nil && err != sql.ErrNoRows {
			fmt.Println(err)
		}
	}
	return ps
}

// to get draft AND live pages leave out 'status' when calling function
func (ps Pages) GetTopPages(navtype string, status ...int) Pages {
	if len(navtype) == 0 {
		navtype = "topnav"
	}
	if len(status) > 0 {
		s := status[0] // status is first aprt of arg
		err := db.Select(&ps, "SELECT * FROM page WHERE status = $1 AND parent = $2 AND navtype = $3 ORDER BY navorder", s, 0, navtype)
		if err != nil && err != sql.ErrNoRows {
			fmt.Println(err)
		}
	} else {
		err := db.Select(&ps, "SELECT * FROM page WHERE parent = $1 AND navtype = $2 ORDER BY navorder", 0, navtype)
		if err != nil && err != sql.ErrNoRows {
			fmt.Println(err)
		}
	}
	return ps
}

func (ps Pages) GetAllTopPages() Pages {
	err := db.Select(&ps, "SELECT * FROM page WHERE parent = $1 ORDER BY id", 0)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err)
	}
	return ps
}

// to get draft AND live pages leave out 'status' when calling function
func (ps Pages) GetSubPages(parentId int64, status ...int) Pages {
	if len(status) > 0 {
		s := status[0]
		err := db.Select(&ps, "SELECT * FROM page WHERE status = $1 and parent = $2 ORDER BY id", s, parentId)
		if err != nil && err != sql.ErrNoRows {
			fmt.Println(err)
		}
	} else {
		err := db.Select(&ps, "SELECT * FROM page WHERE parent = $1 ORDER BY id", parentId)
		if err != nil && err != sql.ErrNoRows {
			fmt.Println(err)
		}
	}
	return ps
}

func GetNavPages(navtype string) (navPages []Page) {
	tp := new(Pages)
	var TopPages []Page
	topPages := tp.GetTopPages(navtype, 1) // 1 for only live pages
	for _, topPage := range topPages {
		sp := new(Pages)
		subPages := sp.GetSubPages(topPage.Id, 1)
		for _, subPage := range subPages {
			topPage.Subpages = append(topPage.Subpages, subPage)
		}
		// before we leave the loop
		// append to TopPages (slice of Pages declared above)
		TopPages = append(TopPages, topPage)
	}
	return TopPages
}

func (ps Posts) GetPosts(status ...int) Posts {
	if len(status) > 0 {
		s := status[0]
		err := db.Select(&ps, "SELECT * FROM post WHERE status = $1 order by id", s)
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Println(err)
			}
			fmt.Println(err)
		}
	} else {
		err := db.Select(&ps, "SELECT * FROM post ORDER BY id")
		if err != nil && err != sql.ErrNoRows {
			fmt.Println(err)
		}
	}
	return ps
}

func (p Post) GetPost(w http.ResponseWriter, r *http.Request, slug string) Post {
	err := db.QueryRowx(`SELECT * FROM post WHERE slug = $1`, slug).StructScan(&p)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err)
	}
	return p
}

func (post Post) GetPostById(id string) Post {
	err := db.QueryRowx(`SELECT * FROM post WHERE id = $1`, id).StructScan(&post)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err)
	}
	return post
}

func (ss Subscribers) GetSubscribers() Subscribers {
	err := db.Select(&ss, "SELECT * FROM subscriber")
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err)
	}
	return ss
}

func (widget Widget) GetWidgetById(id string) Widget {
	err := db.QueryRowx(`SELECT * FROM widget WHERE id = $1`, id).StructScan(&widget)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err)
	}
	return widget
}

func (ws Widgets) GetWidgets() Widgets {
	err := db.Select(&ws, "SELECT * FROM widget")
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err)
	}
	return ws
}

func (slider Slider) GetSliderById(id string) Slider {
	err := db.QueryRowx(`SELECT * FROM slider WHERE id = $1`, id).StructScan(&slider)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err)
	}
	return slider
}

func (ss Sliders) GetSliders() Sliders {
	err := db.Select(&ss, "SELECT * FROM slider")
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err)
	}
	return ss
}

func (ms Medias) GetMedia() Medias {
	err := db.Select(&ms, "SELECT * FROM media")
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err)
	}
	return ms
}

func DeleteItem(w http.ResponseWriter, r *http.Request, item, id string) (err error) {
	var fn string
	switch item {
	case "post", "media", "slider":
		q := fmt.Sprintf("SELECT filename FROM %s WHERE id = %s", item, id)
		err = db.QueryRow(q).Scan(&fn)
		// err = db.QueryRow(`SELECT filename FROM $1 WHERE id = $2`, item, id).Scan(&fn)
		if err != nil {
			fmt.Println(err)
		}
		DeleteFile(fn)
	case "widget":
		err = db.QueryRow(`SELECT filename FROM widget WHERE id = $1`, id).Scan(&fn)
		if err != nil {
			return
			fmt.Println(err)
		}
		// remove join - TODO - warning widget is attached to page(s)
		delJoin := fmt.Sprintf("DELETE FROM %s WHERE widget_id = %s", "page_widget", id)
		_, err = db.Exec(delJoin)
		if err != nil {
			fmt.Println(err)
		}
		DeleteFile(fn)
	case "page":
		updateChildren := fmt.Sprintf("UPDATE page SET parent = 0 WHERE parent = %s", id)
		_, err = db.Exec(updateChildren)
		if err != nil {
			fmt.Println(err)
		}
	case "account":
		err = db.QueryRow(`SELECT avatar FROM account WHERE id = $1`, id).Scan(&fn)
		if err != nil {
			fmt.Println(err)
		}
		DeleteFile(fn)
		session, _ := middleware.Store.Get(r, config.Conf().SessionName)
		delete(session.Values, "UserID")
		delete(session.Values, "Authorized")
		session.Options.MaxAge = -1
		_ = session.Save(r, w)
	}
	// should work for all "items"
	delItem := fmt.Sprintf("DELETE FROM %s WHERE id = %s", item, id)
	_, err = db.Exec(delItem)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func UpdateMap(lat, long string) {
	coords := `UPDATE coords SET latitude=$1,longitude=$2,updated_at=$3`
	db.MustExec(coords, lat, long, time.Now())
}

func UpdatePageStatus(status int, pageid int) {
	pstatus := `UPDATE page SET status=$1, updated_at=$2 WHERE id=$3`
	db.MustExec(pstatus, status, time.Now(), pageid)
}

func UpdateNavOrder(order, pageid int, navtype string) {
	navorder := `UPDATE page SET navorder=$1, navtype=$2 WHERE id=$3`
	db.MustExec(navorder, order, navtype, pageid)
}

func GetMap() Coords {
	coords := Coords{}
	err := db.QueryRowx(`SELECT latitude, longitude FROM coords`).StructScan(&coords)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err)
	}
	return coords
}

func UpdateSettings(r *http.Request) {
	settings := `UPDATE settings SET
	email=$1, 
    address=$2,
	street=$3,
	suburb=$4,
	city=$5,
	code=$6,
	contact=$7,
	twitter=$8,
	facebook=$9,
	linkedin=$10,
	description=$11,
	keywords=$12,
	ganalytics=$13,
	smtp=$14,
	updated_at=$15`

	db.MustExec(settings,
		r.URL.Query().Get("email"),
		r.URL.Query().Get("address"),
		r.URL.Query().Get("street"),
		r.URL.Query().Get("suburb"),
		r.URL.Query().Get("city"),
		r.URL.Query().Get("code"),
		r.URL.Query().Get("contact"),
		r.URL.Query().Get("twitter"),
		r.URL.Query().Get("facebook"),
		r.URL.Query().Get("linkedin"),
		r.URL.Query().Get("description"),
		r.URL.Query().Get("keywords"),
		r.URL.Query().Get("ganalytics"),
		r.URL.Query().Get("smtp"),
		time.Now())
	return
}

func GetSettings() Settings {
	settings := Settings{}
	err := db.QueryRowx(`SELECT * FROM settings`).StructScan(&settings)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err)
	}
	return settings
}

func DeleteFile(filename string) {
	if filename != "" {
		absdir, _ := os.Getwd()
		err := os.Remove(path.Join(absdir, "static", "uploads", filename))
		err = os.Remove(path.Join(absdir, "static", "uploads", "thumb_"+filename))
		if err != nil {
			fmt.Println(err)
		}
	}
	return
}
