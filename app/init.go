package app

import (
	"github.com/robfig/revel"
	"strings"
)

func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.ActionInvoker,           // Invoke the action.
	}

	// register startup functions with OnAppStart
	// ( order dependent )
	// revel.OnAppStart(InitDB())
	// revel.OnAppStart(FillCache())
}

// TODO turn this into revel.HeaderFilter
// should probably also have a filter for CSRF
// not sure if it can go in the same filter or not
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	// Add some common security headers
	SetMimeType(c)
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}

var ExtensionMap = map[string]string{
	"png":  "image/png",
	"jpg":  "image/jpeg",
	"gif":  "image/gif",
	"js":   "application/javascript; charset=utf-8",
	"css":  "text/css; charset=utf-8",
	"xml":  "text/xml; charset=utf-8",
	"html": "text/html; charset=utf-8",
}

func SetMimeType(c *revel.Controller) {
	path := c.Request.RequestURI
	if strings.LastIndex(path, ".") >= 0 {
		extension := strings.ToLower(path[strings.LastIndex(path, ".")+1:])
		if mime, ok := ExtensionMap[extension]; ok {
			c.Response.Out.Header().Set("Content-Type", mime)
			if strings.Index(mime, "image/") != -1 {
				c.Response.Out.Header().Set("Cache-Control", "public")
			}
		}
	}
}
