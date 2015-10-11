package process

import (
	"bitbucket.org/enkdr/enkoder/config"
	"bitbucket.org/enkdr/enkoder/dbconn"
	"bitbucket.org/enkdr/enkoder/middleware"
	m "bitbucket.org/enkdr/enkoder/models"
	u "bitbucket.org/enkdr/enkoder/utils"
	"code.google.com/p/go.crypto/bcrypt"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var db = dbconn.NewDB()

func LoginP(w http.ResponseWriter, r *http.Request) (message []string) {
	account := m.Account{}
	// message = append(message,"") // handler directs to admin on successful login - maybe return a boolean as well?
	email := r.FormValue("email")
	password := r.FormValue("password")
	if email == "" { // no email
		message = append(message, "Please enter your email")
		return message
	}
	if password == "" { // no password
		message = append(message, "Please enter your password")
		return message
	}
	err := db.QueryRowx(`SELECT email, password, id FROM account WHERE email = $1`, email).StructScan(&account)
	if err != nil {
		if err == sql.ErrNoRows { // email doesnt exists in db
			message = append(message, "We couldn't find that email")
			return message
		}
		// message = append(message,"problem with db.QueryRowx")
	}
	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil {
		message = append(message, "Incorrect password - please try again")
	} else {
		// success - create session
		session, err := middleware.Store.Get(r, config.Conf().SessionName)
		if err != nil {
			fmt.Println(err)
		}
		session.Options = &sessions.Options{
			Path: "/",
		}
		session.Values["UserID"] = account.Id
		session.Values["Authorized"] = 1
		session.Save(r, w)
	}
	return message
}

func NewAccountP(w http.ResponseWriter, r *http.Request, accid string) (processed bool, message []string) {
	processed = false
	firstname := r.FormValue("firstname")
	lastname := r.FormValue("lastname")
	email := r.FormValue("email")
	password := r.FormValue("password")
	confirm := r.FormValue("confirm")
	re := regexp.MustCompile(".+@.+\\..+")
	matched := re.Match([]byte(email))
	if matched == false {
		message = append(message, "Please enter a valid email address")
	}
	if strings.TrimSpace(firstname) == "" {
		message = append(message, "Please enter a first name")
	}
	if strings.TrimSpace(lastname) == "" {
		message = append(message, "Please enter a last name")
	}
	if strings.TrimSpace(password) == "" { // no password
		message = append(message, "Please enter a password")
	}
	if password != confirm {
		message = append(message, "Passwords do not match")
	}

	// file is not mandatory - if empty - set to empty string
	_, filename := uploadFile(r)

	if len(message) == 0 {
		if accid != "" {
			var avatar string // get avatar - use val if filename hasn't changed
			err := db.QueryRow(`SELECT avatar FROM account WHERE id = $1`, accid).Scan(&avatar)
			if err != nil {
				fmt.Println(err)
			}
			if len(filename) == 0 { // no filename from form - use "avatar" from db
				filename = avatar
			} else {
				m.DeleteFile(avatar) // delete the avatar - will be inserted as filename below
			}
			pwd, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
			editAccount := `UPDATE account SET firstname = $1, lastname= $2, email = $3, password = $4, avatar = $5, updated_at = $6 WHERE id = $7`
			db.MustExec(editAccount, firstname, lastname, email, pwd, filename, time.Now(), accid)
			processed = true
			message = append(message, "Account updated")
			return processed, message
		} else {
			var id int
			err := db.QueryRow(`SELECT id FROM account WHERE email = $1`, email).Scan(&id)
			if err != nil {
				if err == sql.ErrNoRows { // user doesnt exists in db
					pwd, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
					addAccount := `INSERT INTO account (email, password, firstname, lastname, avatar) VALUES ($1, $2, $3, $4, $5)`
					db.MustExec(addAccount, email, pwd, firstname, lastname, filename)
					processed = true
					message = append(message, "User created")
					return processed, message
				}
			}
		}
	}
	message = append(message, "Email already exists")
	return processed, message
}

func NewPageP(w http.ResponseWriter, r *http.Request, pid string) (processed bool, message []string) {

	// these page titles are reserved
	exclude := map[string]bool{
		"admin":       true,
		"news":        true,
		"login":       true,
		"logout":      true,
		"unsubscribe": true,
		"reset":       true}

	processed = false
	r.ParseForm()
	title := r.FormValue("title")
	content := r.FormValue("content")
	metatags := r.FormValue("metatags")
	widgets := r.Form["widgets"]
	status, _ := strconv.ParseInt(r.FormValue("status"), 10, 64)
	parent, _ := strconv.ParseInt(r.FormValue("parent"), 10, 64)
	slug := u.GenerateSlug(r.FormValue("title"))
	if strings.TrimSpace(title) == "" {
		message = append(message, "Please enter a title")
		return processed, message
	}
	if exclude[strings.ToLower(title)] {
		message = append(message, "Title is reserved")
		return processed, message
	}
	if pid != "" { // editing page
		editPage := `UPDATE page SET title = $1, content = $2, metatags = $3, slug= $4, status = $5, parent = $6, updated_at = $7 WHERE id = $8`
		db.MustExec(editPage, title, content, metatags, slug, status, parent, time.Now(), pid)
		purgePageWidget := `DELETE FROM page_widget WHERE page_id = $1`
		db.MustExec(purgePageWidget, pid)
		for _, v := range widgets {
			insertPageWidgets := `INSERT INTO page_widget (page_id, widget_id) VALUES ($1, $2)`
			db.MustExec(insertPageWidgets, pid, v)
		}
		processed = true
		message = append(message, "Page updated")
		return processed, message
	}
	if len(message) == 0 {
		var slg string
		err := db.QueryRow(`SELECT slug FROM page WHERE slug = $1`, slug).Scan(&slg)
		if err != nil {
			if err == sql.ErrNoRows { // slug doesnt exists in db

				addPage := `INSERT INTO page (title, content, metatags, slug, status, parent) VALUES ($1, $2, $3, $4, $5, $6)`
				db.MustExec(addPage, title, content, metatags, slug, status, parent)

				// need new page.id for page_widget query
				// slug is unique enough...
				q := fmt.Sprintf(`SELECT id FROM page WHERE slug = '%s'`, slug)
				var pageid int
				err := db.QueryRow(q).Scan(&pageid)
				if err != nil {
					fmt.Println(err)
				}
				for _, v := range widgets {
					insertPageWidgets := `INSERT INTO page_widget (page_id, widget_id) VALUES ($1, $2)`
					db.MustExec(insertPageWidgets, pageid, v)
				}

				processed = true
				message = append(message, "Page created")
				return processed, message
			}
		}
		if len(slg) > 0 || pid == "" {
			message = append(message, "Page title already exists")
		}
	}
	return processed, message
}

func NewMediaP(w http.ResponseWriter, r *http.Request) (message []string, err error) {

	title := r.FormValue("title")
	r.ParseMultipartForm(32 << 20)
	if strings.TrimSpace(title) == "" {
		message = append(message, "Please enter a title")
		return message, errors.New("No title")
	}
	err, filename := uploadFile(r)
	if err != nil {
		message = append(message, "Problem with upload")
		return message, errors.New("problem upload")
	}
	addMedia := `INSERT INTO media (title, filename) VALUES ($1, $2)`
	db.MustExec(addMedia, title, filename)
	return message, nil
}

func NewWidgetP(w http.ResponseWriter, r *http.Request, wid string) (message []string, err error) {
	title := r.FormValue("title")
	url := r.FormValue("url")
	content := r.FormValue("content")
	widgetsize := r.FormValue("widgetsize")
	r.ParseMultipartForm(32 << 20)
	if strings.TrimSpace(title) == "" {
		message = append(message, "Please enter a title")
		return message, errors.New("No title")
	}
	_, filename := uploadFile(r)
	if len(message) == 0 {
		if wid != "" {
			var oldfile string // get oldfile - use val if filename hasn't changed
			err := db.QueryRow(`SELECT filename FROM widget WHERE id = $1`, wid).Scan(&oldfile)
			if err != nil {
				fmt.Println(err)
			}
			if len(filename) == 0 { // no filename from form - use "oldfile" from db
				filename = oldfile
			} else {
				m.DeleteFile(oldfile) // delete the oldfile - will be inserted as filename below
			}
			editWidget := `UPDATE widget SET title = $1, content= $2, filename = $3, url = $4, widgetsize = $5, updated_at = $6 WHERE id = $7`
			db.MustExec(editWidget, title, content, filename, url, widgetsize, time.Now(), wid)
			message = append(message, "Widget updated")
			return message, nil
		} else {
			addWidget := `INSERT INTO widget (title, content, filename, url, widgetsize) VALUES ($1, $2, $3, $4, $5)`
			db.MustExec(addWidget, title, content, filename, url, widgetsize)
			message = append(message, "Widget created")
			return message, nil
		}
	}
	message = append(message, "Widget created")
	return message, nil
}

func NewPostP(w http.ResponseWriter, r *http.Request, pid string) (message []string, err error) {
	title := r.FormValue("title")
	content := r.FormValue("content")
	tags := r.FormValue("tags")
	accid, _ := strconv.ParseInt(r.FormValue("accid"), 10, 64)
	status, _ := strconv.ParseInt(r.FormValue("status"), 10, 64)
	if strings.TrimSpace(title) == "" {
		message = append(message, "Please enter a title")
		return message, errors.New("No title")
	}

	t := time.Now()
	slug := fmt.Sprintf("news/%s/%s", t.Format("2006-Jan-02"), u.GenerateSlug(r.FormValue("title")))

	r.ParseMultipartForm(32 << 20)
	_, filename := uploadFile(r)
	if len(message) == 0 {
		if pid != "" {
			var oldfile string // get oldfile - use val if filename hasn't changed
			err := db.QueryRow(`SELECT filename FROM post WHERE id = $1`, pid).Scan(&oldfile)
			if err != nil {
				fmt.Println(err)
			}
			if len(filename) == 0 { // no filename from form - use "oldfile" from db
				filename = oldfile
			} else {
				m.DeleteFile(oldfile) // delete the oldfile - will be inserted as filename below
			}
			editPost := `UPDATE post SET title = $1, content = $2, filename = $3, tags = $4, status = $5, slug = $6, updated_at = $7 WHERE id = $8`
			db.MustExec(editPost, title, content, filename, tags, status, slug, time.Now(), pid)
			message = append(message, "Post updated")
			return message, nil
		} else {
			addPost := `INSERT INTO post (title, content, filename, tags, status, slug, account_id) VALUES ($1, $2, $3, $4, $5, $6, $7)`
			db.MustExec(addPost, title, content, filename, tags, status, slug, accid)
			message = append(message, "Post created")
			return message, nil
		}
	}
	message = append(message, "Post created")
	return message, nil
}

func NewSliderP(w http.ResponseWriter, r *http.Request, sid string) (message []string, err error) {
	title := r.FormValue("title")
	content := r.FormValue("content")
	url := r.FormValue("url")
	r.ParseMultipartForm(32 << 20)
	if strings.TrimSpace(title) == "" {
		message = append(message, "Please enter a heading")
		return message, errors.New("No heading")
	}
	_, filename := uploadFile(r)
	if len(message) == 0 {
		if sid != "" {
			var oldfile string // get oldfile - use val if filename hasn't changed
			err := db.QueryRow(`SELECT filename FROM slider WHERE id = $1`, sid).Scan(&oldfile)
			if err != nil {
				fmt.Println(err)
			}
			if len(filename) == 0 { // no filename from form - use "oldfile" from db
				filename = oldfile
			} else {
				m.DeleteFile(oldfile) // delete the oldfile - will be inserted as filename below
			}
			editSlider := `UPDATE slider SET title = $1, content= $2, filename = $3, url = $4, updated_at = $5 WHERE id = $6`
			db.MustExec(editSlider, title, content, filename, url, time.Now(), sid)
			message = append(message, "Slider updated")
			return message, nil
		} else {
			addSlider := `INSERT INTO slider (title, content, filename, url) VALUES ($1, $2, $3, $4)`
			db.MustExec(addSlider, title, content, filename, url)
			message = append(message, "Slider created")
			return message, nil
		}
	}
	message = append(message, "Slider added")
	return message, nil
}

////

func uploadFile(r *http.Request) (err error, filename string) {
	file, handler, err := r.FormFile("file")
	if err != nil {
		return err, ""
		fmt.Println(err)
	}
	if file == nil {
		return nil, ""
	}

	defer file.Close()
	filename = strings.Replace(strconv.FormatInt(time.Now().Unix(), 10)+handler.Filename, " ", "", -1)
	// filename = strings.Replace(handler.Filename, " ", "", -1)
	pth := "./static/uploads/"
	if handler != nil {
		mimetype := handler.Header["Content-Type"]
		if u.MimeTypesAllowed[mimetype[0]] {
			f, err := os.Create(pth + filename)
			if err != nil {
				fmt.Println(err)
				return err, ""
			}
			io.Copy(f, file)
			if u.MimeTypesImg[mimetype[0]] { // file is image
				makeThumbnail(pth, filename, mimetype[0])
			}
		}
	} else {
		err = errors.New("Filetype not allowed")
		fmt.Println(err)
		return err, ""
	}
	return nil, filename
}

func makeThumbnail(path, filename, mimetype string) error {

	file, err := os.Open(path + filename)
	if err != nil {
		return errors.New("could not open file.\n")
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return errors.New("could not decode file.\n")
	}
	m := resize.Resize(0, 80, img, resize.Lanczos3)

	out, err := os.Create(path + "thumb_" + filename)
	if err != nil {
		return errors.New("could not create thumbnail.\n")
	}
	defer out.Close()

	switch mimetype {
	case "image/jpeg", "image/jpg":
		err = jpeg.Encode(out, m, nil)
		if err != nil {
			return errors.New("could not encode (jpeg).\n")
		}
	case "image/png":
		err = png.Encode(out, m)
		if err != nil {
			return errors.New("could not encode (png).\n")
		}
	// // case "application/pdf": // not image, but application !
	// // 	fmt.Println(filetype)
	default:
		fmt.Println("unknown file type uploaded")
	}

	return nil
}

// TODO remove w?
func ResetPasswordP(w http.ResponseWriter, r *http.Request) (message []string, err error) {
	email := r.FormValue("email")
	re := regexp.MustCompile(".+@.+\\..+")
	matched := re.Match([]byte(email))
	if matched == false {
		message = append(message, "Email is invalid")
		return message, errors.New("email invalid")
	}
	var id int
	err = db.QueryRow(`SELECT id FROM account WHERE email = $1`, email).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows { // user doesnt exists in db
			message = append(message, "Email doesnt exist")
			return message, errors.New("Email doesnt exist")
		}
		message = append(message, "Query failed")
		return message, errors.New("Query failed")
	}
	newpwd := u.NewPassword()
	pwd, _ := bcrypt.GenerateFromPassword([]byte(newpwd), bcrypt.DefaultCost)
	updatePW := `UPDATE account SET password = $1, updated_at = $2 WHERE id = $3`
	db.MustExec(updatePW, pwd, time.Now(), id)

	msg := fmt.Sprintf("Here is your temporary password - be sure to login and choose a new password: %s\n", newpwd)

	u.SendMail("info@enkoder.com.au", email, "Password Updated", msg)
	message = append(message, "Account updated with new password")
	return message, nil
}

func SearchP(w http.ResponseWriter, r *http.Request) (message []string, pages m.Pages) {
	search := r.FormValue("search")
	err := db.Select(&pages, "select * from page where to_tsvector('english',content) @@ to_tsquery('english',$1);", search)
	if err != nil {
		if err == sql.ErrNoRows {
			message = append(message, "No results")
			return message, nil
		}
	}
	return message, pages
}

func SubscribeP(w http.ResponseWriter, r *http.Request) (message string, err error) {
	email := r.FormValue("email")
	fname := r.FormValue("firstname")
	lname := r.FormValue("lastname")
	randomid := u.RandStr(32, "alphanum")

	if len(email) == 0 {
		return "please enter an email", errors.New("no email value")
	}

	var subscriber int
	db.QueryRow(`SELECT id FROM subscriber WHERE email = $1`, email).Scan(&subscriber)
	if subscriber != 0 {
		return "email already registered", errors.New("subscriber email already registered")
	}

	addSubscriber := `INSERT INTO subscriber (email, firstname, lastname, newsletter, randomid) VALUES ($1,$2,$3,$4,$5)`
	db.MustExec(addSubscriber, email, fname, lname, 1, randomid)
	return "subscriber added", nil
}

func UnSubscribeP(randomid string) (message string, err error) {
	purgeSubscriber := `DELETE FROM subscriber WHERE randomid = $1`
	_, err = db.Exec(purgeSubscriber, randomid)
	if err != nil {
		return "Problem Unsubscribing", errors.New("Exec failed")
	}
	return "Unsubscribe successful", nil
}
