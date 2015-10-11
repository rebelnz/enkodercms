package routes

import (
	h "bitbucket.org/enkdr/enkoder/handlers"
	"net/http"
	"time"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	Protected   bool
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{"index", "GET", "/", false, h.Index(time.RFC1123)}, // how to pass arg to handler
	Route{"index", "GET", "/home", false, h.Index(time.RFC1123)},

	Route{"favicon", "GET", "/favicon.ico", false, h.ServeFileHandler},

	Route{"login", "GET", "/login", false, h.Login},
	Route{"login", "POST", "/login", false, h.LoginProcess},
	Route{"logout", "GET", "/logout", false, h.Logout},
	Route{"reset", "POST", "/reset", false, h.ResetProcess},

	Route{"search", "POST", "/search", false, h.SearchProcess},
	Route{"subscribe", "POST", "/subscribe", false, h.SubscribeProcess},
	Route{"unsubscribe", "GET", "/unsubscribe/{randomid}", false, h.UnSubscribeProcess},

	Route{"contact", "POST", "/contact", false, h.ContactProcess},

	Route{"admin", "GET", "/admin", true, h.AdminIndex},
	Route{"settings", "GET", "/admin/settings", true, h.AdminGeneric("settings")},

	Route{"accounts", "GET", "/admin/accounts", true, h.AdminGeneric("accounts")},
	Route{"newaccount", "GET", "/admin/newaccount", true, h.AdminNewAccount},
	Route{"editaccount", "GET", "/admin/editaccount/{id}", true, h.AdminEditAccount},
	Route{"newaccount", "POST", "/admin/newaccount", true, h.AdminNewAccountProcess},

	Route{"sliders", "GET", "/admin/sliders", true, h.AdminGeneric("sliders")},
	Route{"newslider", "GET", "/admin/newslider", true, h.AdminNewSlider},
	Route{"newslider", "POST", "/admin/newslider", true, h.AdminNewSliderProcess},
	Route{"editslider", "GET", "/admin/editslider/{id}", true, h.AdminEditSlider},

	Route{"pages", "GET", "/admin/pages", true, h.AdminPages},
	Route{"newpage", "GET", "/admin/newpage", true, h.AdminNewPage},
	Route{"editpage", "GET", "/admin/editpage/{slug}", true, h.AdminEditPage},
	Route{"newpage", "POST", "/admin/newpage", true, h.AdminNewPageProcess},

	Route{"widgets", "GET", "/admin/widgets", true, h.AdminGeneric("widgets")},
	Route{"newwidget", "GET", "/admin/newwidget", true, h.AdminNewWidget},
	Route{"newwidget", "POST", "/admin/newwidget", true, h.AdminNewWidgetProcess},
	Route{"editwidget", "GET", "/admin/editwidget/{id}", true, h.AdminEditWidget},

	Route{"medias", "GET", "/admin/medias", true, h.AdminGeneric("medias")},
	Route{"upload", "GET", "/admin/newmedia", true, h.AdminNewMedia},
	Route{"upload", "POST", "/admin/newmedia", true, h.AdminNewMediaProcess},
	Route{"deleteitem", "GET", "/admin/deleteitem/{item}/{id}", true, h.AdminDeleteItem},

	Route{"posts", "GET", "/admin/posts", true, h.AdminGeneric("posts")},
	Route{"newpost", "GET", "/admin/newpost", true, h.AdminNewPost},
	Route{"newpost", "POST", "/admin/newpost", true, h.AdminNewPostProcess},
	Route{"editpost", "GET", "/admin/editpost/{id}", true, h.AdminEditPost},

	Route{"navigation", "GET", "/admin/navigation", true, h.AdminNavigation},

	Route{"map", "GET", "/admin/map", true, h.AdminGeneric("map")},

	// Route{"newsletters", "GET", "/admin/newsletters", true, h.AdminNewsletters},
	// Route{"newsletter", "GET", "/admin/newsletter/{id}", true, h.AdminNewsletter},
	// Route{"newnewsletter", "GET", "/admin/newnewsletter", true, h.AdminNewNewsletter},
	// Route{"newnewsletter", "POST", "/admin/newnewsletter", true, h.AdminNewNewsletterProcess},
	// Route{"editnewsletter", "GET", "/admin/editnewsletter/{id}", true, h.AdminEditNewsletter},

	// AJAX //
	Route{"getsettings", "GET", "/admin/ajax/getsettings", false, h.AdminGetSettings},
	Route{"updatesettings", "GET", "/admin/ajax/updatesettings", true, h.AdminUpdateSettings},

	Route{"updatemap", "POST", "/admin/ajax/updatemap", true, h.AdminUpdateMap},
	Route{"getmap", "GET", "/admin/ajax/getmap", false, h.AdminGetMap},

	Route{"updatenavorder", "POST", "/admin/ajax/navorder", true, h.AdminUpdateNavOrder},

	Route{"updatepagestatus", "GET", "/admin/ajax/updatepagestatus/{status}/{pageid}", true, h.AdminUpdatePageStatus},

	// Route{"updatesettings", "GET", "/admin/ajax/updatesettings", true, h.AdminUpdateSettings},
	// Route{"checkaccount","GET","/admin/JSON/checkaccount",true,h.AdminCheckAccount,},

	Route{"news", "GET", "/news/{date}/{slug}", false, h.ShowPost},
	Route{"news", "GET", "/news", false, h.ShowPosts},
	Route{"page", "GET", "/{slug}", false, h.ShowPage},
}
