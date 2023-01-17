package main

import (
	"github.com/fanke15/tool/api/handler"
	"github.com/fanke15/tool/pkg/lib"
	"github.com/kataras/iris/v12"
	"github.com/yosssi/ace"
	"html/template"
)

func main() {
	lib.NewBolt()
	lib.InitBolt().Save("service", []byte("tool"))

	app := iris.New()
	app.HandleDir("/assets", "./web/static")
	app.Favicon("./web/static/favicon.ico")

	//-----------------------WEB-----------------------
	webPage := app.Party("page")

	// 注册方法
	funcMap := template.FuncMap{
		"ConvertToJson": func(s string) string {
			return s
		},
	}

	{
		webPage.Get("/dashboard", func(ctx iris.Context) {
			tpl, err := ace.Load("./web/static/template/index", "./web/static/template/convertToJson", &ace.Options{DynamicReload: true})
			if err != nil {
				return
			}
			_ = tpl.Execute(ctx.ResponseWriter(), map[string]interface{}{
				"title": "tool",
			})
		})
		webPage.Get("/dashboard/tojson", func(ctx iris.Context) {
			tpl, err := ace.Load("./web/static/template/index", "./web/static/template/convertToJson", &ace.Options{
				DynamicReload: true,
				FuncMap:       funcMap})
			if err != nil {
				return
			}
			_ = tpl.Execute(ctx.ResponseWriter(), map[string]interface{}{
				"title": "tool",
			})
		})
		webPage.Get("/dashboard/dict", func(ctx iris.Context) {
			tpl, err := ace.Load("./web/static/template/index", "./web/static/template/dict", &ace.Options{DynamicReload: true})
			if err != nil {
				return
			}
			_ = tpl.Execute(ctx.ResponseWriter(), map[string]interface{}{
				"title": "tool",
			})
		})
	}

	//-----------------------API-----------------------
	webApi := app.Party("api")
	{
		ch := handler.ConvertHandler{}
		webApi.Post("/convert/tojson", ch.ToJson)

		dict := handler.DictHandler{}
		webApi.Get("/dict/fields", dict.GetDickFieldList)
	}

	app.Run(iris.Addr(":8080"), iris.WithoutInterruptHandler)
}
