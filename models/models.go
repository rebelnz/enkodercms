package models

import (
	"time"
)

type Base struct {
	Id         int64
	Created_At time.Time
	Updated_At time.Time
}

type Account struct {
	Base
	Password  string
	Firstname string
	Lastname  string
	Email     string
	Avatar    string
	Priv      int64
}

type Accounts []Account

type Subscriber struct {
	Base
	RandomId   string
	Firstname  string
	Lastname   string
	Email      string
	Newsletter int64
}

type Subscribers []Subscriber

// not in DB
type Meta struct {
	Sitename    string
	Title       string
	Nav         string
	Description string
}

type Page struct {
	Base
	Slug       string
	Title      string
	Content    string
	Metatags   string
	Order      int64
	Priv       int64
	Status     int64
	Parent     int64
	Navtype    string
	Navorder   int64
	Subpages   Pages
	Subwidgets *Widgets
}

type Pages []Page

type Media struct {
	Base
	Title    string
	Filename string
}

type Medias []Media

type Widget struct {
	Base
	Title      string
	Content    string
	Url        string
	Filename   string
	Widgetsize string
}

type Widgets []Widget

// for map -- warning int/string
type Coords struct {
	Latitude  string
	Longitude string
}

type Settings struct {
	Base
	Email       string
	Address     string
	Street      string
	Suburb      string
	City        string
	Code        string
	Contact     string
	Twitter     string
	Facebook    string
	Linkedin    string
	Description string
	Keywords    string
	Ganalytics  string
	Smtp        string
}

type Slider struct {
	Base
	Title    string
	Content  string
	Filename string
	Url      string
}

type Sliders []Slider

type Post struct {
	Base
	Slug       string
	Account_id int64
	Title      string
	Content    string
	Filename   string
	Tags       string
	Priv       int64
	Status     int64
}

type Posts []Post

// type Item struct {
// 	Base
// 	Account_id int64
// 	Slug       string
// 	Parent     int64
// 	Navtype    string
// 	Navorder   int64
// 	Type       string
// 	Title      string
// 	Content    string
// 	Link       string
// 	Filename   string
// 	Widgetsize string
// 	Tags       string
// 	Metatags   string
// 	Order      int64
// 	Priv       int64
// 	Status     int64
// 	SubItems   Items
// }

// type Items []Item
