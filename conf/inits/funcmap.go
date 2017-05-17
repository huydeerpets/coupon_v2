package inits

import (
	"html/template"
	"net/url"
	"reflect"
	"strings"
	"time"

	"github.com/beego/i18n"
	"ttc/cps/utils"
	"github.com/mattn/go-runewidth"

	"github.com/astaxie/beego"
)

func init() {

	beego.AddFuncMap("i18n", i18n.Tr)

	beego.AddFuncMap("i18nvn", func(format string, args ...interface{}) string {
		return i18n.Tr("vn-VN", format, args...)
	})

	beego.AddFuncMap("datenow", func(format string) string {
		return time.Now().Add(time.Duration(7) * time.Hour).Format(format)
	})

	beego.AddFuncMap("dateformatVst", func(in time.Time) string {
		in = in.Add(time.Duration(7) * time.Hour)
		return in.Format("2006/01/02 15:04")
	})

	beego.AddFuncMap("qescape", func(in string) string {
		return url.QueryEscape(in)
	})

	beego.AddFuncMap("nl2br", func(in string) string {
		return strings.Replace(in, "\n", "<br>", -1)
	})

	beego.AddFuncMap("tostr", func(in interface{}) string {
		return utils.ToStr(reflect.ValueOf(in).Interface())
	})

	beego.AddFuncMap("first", func(in interface{}) interface{} {
		return reflect.ValueOf(in).Index(0).Interface()
	})

	beego.AddFuncMap("last", func(in interface{}) interface{} {
		s := reflect.ValueOf(in)
		return s.Index(s.Len() - 1).Interface()
	})

	beego.AddFuncMap("truncate", func(in string, length int) string {
		return runewidth.Truncate(in, length, "...")
	})

	beego.AddFuncMap("noname", func(in string) string {
		if in == "" {
			return "(nothing)"
		}
		return in
	})

	beego.AddFuncMap("cleanurl", func(in string) string {
		return strings.Trim(strings.Trim(in, " "), "　")
	})

	beego.AddFuncMap("append", func(data map[interface{}]interface{}, key string, value interface{}) template.JS {
		if _, ok := data[key].([]interface{}); !ok {
			data[key] = []interface{}{value}
		} else {
			data[key] = append(data[key].([]interface{}), value)
		}
		return template.JS("")
	})

	beego.AddFuncMap("appendmap", func(data map[interface{}]interface{}, key string, name string, value interface{}) template.JS {
		v := map[string]interface{}{name: value}

		if _, ok := data[key].([]interface{}); !ok {
			data[key] = []interface{}{v}
		} else {
			data[key] = append(data[key].([]interface{}), v)
		}
		return template.JS("")
	})
}
