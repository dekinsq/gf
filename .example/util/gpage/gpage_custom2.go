package main

import (
	"github.com/dekinsq/gf/frame/g"
	"github.com/dekinsq/gf/net/ghttp"
	"github.com/dekinsq/gf/os/gview"
	"github.com/dekinsq/gf/util/gpage"
)

// pageContent customizes the page tag name.
func pageContent(page *gpage.Page) string {
	page.NextPageTag = "NextPage"
	page.PrevPageTag = "PrevPage"
	page.FirstPageTag = "HomePage"
	page.LastPageTag = "LastPage"
	pageStr := page.FirstPage()
	pageStr += page.PrevPage()
	pageStr += page.PageBar()
	pageStr += page.NextPage()
	pageStr += page.LastPage()
	return pageStr
}

func main() {
	s := g.Server()
	s.BindHandler("/page/custom2/*page", func(r *ghttp.Request) {
		page := r.GetPage(100, 10)
		buffer, _ := gview.ParseContent(`
        <html>
            <head>
                <style>
                    a,span {padding:8px; font-size:16px;}
                    div{margin:5px 5px 20px 5px}
                </style>
            </head>
            <body>
                <div>{{.page}}</div>
            </body>
        </html>
        `, g.Map{
			"page": pageContent(page),
		})
		r.Response.Write(buffer)
	})
	s.SetPort(10000)
	s.Run()
}
