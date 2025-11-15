package html

import (
	"net/http"
	"strings"

	"github.com/spf13/viper"
)

// Render for Chi framework
func Render(w http.ResponseWriter, name string, data map[string]interface{}) {
	if data == nil {
		data = make(map[string]interface{})
	}

	// Add common template data
	data["app_name"] = viper.GetString("APPNAME")

	// Auto-append .html if not present
	tmpl := name
	if !strings.HasSuffix(name, ".html") {
		tmpl = name + ".html"
	}
	RenderTemplate(w, tmpl, data)
}
