package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/antchfx/xquery/html"
	"golang.org/x/net/html"
)

// /html/body/div/div/h2

func main() {
	response, _ := http.Get("https://biz.kaslyju.net")
	body, _ := ioutil.ReadAll(response.Body)
	/*htmlstr := `<?xml version="1.0" ?>
	  <html>
	  <head>
	   <title>this is a title666666</title>
	  </head>
	  <body>Hello,World</body>


	  </html>`*/
	//fmt.Println(string(body))
	fmt.Println("-------------------------------------------------------")
	//root, err := xml.Parse(strings.NewReader(string(body)))
	//if err != nil {
	//	panic(err)
	//}
	//title := xmlquery.FindOne(root, "//html/body/div/div/h2")
	//fmt.Println(title.InnerText())
	//xp := goxpath.MustParse("/html/body/div/div/h2")
	//t := xmltree.MustParseXML(bytes.NewBufferString(xml.Header + string(body)))
	//res := goxpath.MustExec(xp, t, nil)
	//fmt.Println(res[0])
	root, err := html.Parse(strings.NewReader(string(body)))
	if err != nil {
		panic(err)
	}
	//p := "//*[@id='loginForm']" + "/@id"
	//node := htmlquery.FindOne(root, "//*[@id='loginForm']/@id")
	node1 := htmlquery.Find(root, "//*[@id=\"loginForm\"]/@class")
	fmt.Println(node1[0].Attr[2].Val)
	//htmlquery.Find
	fmt.Println(len(node1))
	fmt.Println(htmlquery.InnerText(node1[0]))

}
