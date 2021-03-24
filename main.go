package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/russross/blackfriday"
)

func getLayoutStart(isIntro bool) string {
	intro := ""
	if isIntro {
		intro = `<p class="site-intro">Hi there, it's me and I ship code for a coffee. Find me on:</p>
				<ul style="list-style-type: none;">
					<li>
						<i class="fa fa-github fa-1x" aria-hidden="true"></i>
						<a href="https://github.com/puertigris">GitHub</a>
					</li>
					<li>
						<i class="fa fa-twitter fa-1x" aria-hidden="true"></i>
                    	<a href="https://twitter.com/quannv132">Twitter</a>
					</li>
					<li>
						<i class="fa fa-rss fa-1x" aria-hidden="true"></i>
                    	<a href="/rss">RSS</a>
					</li>
            	</ul>
				`
	}
	return `<!DOCTYPE html>
			<html>
			  <head>
			  <meta charset="utf-8">
			  <meta http-equiv="X-UA-Compatible" content="IE=edge">
			  <meta name="viewport" content="width=device-width, initial-scale=1">
			  <title>QUAN NGUYEN</title>
			  <link rel="stylesheet" href="/assets/styte.css">
    		  <link rel="stylesheet" href="assets/styte.css">
			  <link rel="alternate" type="application/rss+xml" title="puertigris" href="/rss.xml">
			  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css">
			</head>
			
			  <body>
				<div class="page-content">
					<div class="home">
						<section class="site-header">
							<h1 class="smallcap"><a class="site-title" href="/">QUAN NGUYEN</a></h1>
							` + getNav() + `
							` + intro + `
						</section>
					`
}

func getNav() string {
	return `<p class="site-nav"><a href="/">Home</a> / <a href="/about">About</a> / <a href="/game/tic-tac-toe">Game</a></p>`
}

func getLayoutEnd() string {
	return `
			<div class="copyright">
        <p>&copy; 2021 <a href="/"><strong>QUAN NGUYEN</strong></a></p>
      </div>
    </div>
    
  </body>
</html>`
}

func getPostPage(title string, date string) string {
	return `<!DOCTYPE html>
			<html>
			  <head>
			  <meta charset="utf-8">
			  <meta http-equiv="X-UA-Compatible" content="IE=edge">
			  <meta name="viewport" content="width=device-width, initial-scale=1">
			  <title>` + title + ` - Quan Nguyen</title>
			  <link rel="stylesheet" href="/assets/styte.css">
			  <link rel="stylesheet" href="assets/styte.css">
			  <link rel="alternate" type="application/rss+xml" title="TextLog" href="/rss.xml">
			</head>
			<body>
				<div class="page-content">
				  <article class="post" itemscope itemtype="http://schema.org/BlogPosting">
					<header class="post-header">
						` + getNav() + `
						<h1 class="post-title" itemprop="name headline">` + title + `</h1>
						<p class="post-meta">
				  			<time datetime="2017-01-15T00:00:00+00:00" itemprop="datePublished">` + date + `</time>
						</p>
			  		</header>
			  		<div class="post-content" itemprop="articleBody">`
}

func getFile(f string) []byte {
	b, err := ioutil.ReadFile(f)

	if err != nil {
		panic(err)
	}

	return b
}

func getDir(dir string) []os.FileInfo {
	p, err := ioutil.ReadDir(dir)

	if err != nil {
		panic(err)
	}

	return p
}

func writeFile(fileName string, b bytes.Buffer) {
	err := ioutil.WriteFile(fileName+".html", b.Bytes(), 0644)

	if err != nil {
		panic(err)
	}
}

func writeCSSFile(fileName string, b bytes.Buffer) {
	err := ioutil.WriteFile(fileName+".css", b.Bytes(), 0644)

	if err != nil {
		panic(err)
	}
}

func writeRSSFile(fileName string, b bytes.Buffer) {
	err := ioutil.WriteFile(fileName+".xml", b.Bytes(), 0644)

	if err != nil {
		panic(err)
	}
}

func getSiteTitle() string {
	return strings.Split(string(getFile("_sections/header.md")), "\n")[0][2:]
}

func getPostMeta(fi os.FileInfo) (string, string, string) {
	id := fi.Name()[:len(fi.Name())-3]
	date := fi.Name()[0:10]
	title := strings.Split(string(getFile("_posts/"+fi.Name())), "\n")[0]
	title = strings.ReplaceAll(title, "[comment]: <> (", "")
	title = strings.ReplaceAll(title, ")", "")
	return id, date, title
}

func convertDate(date string) string {
	myDate, _ := time.Parse("2006-01-02", date)
	return fmt.Sprintf("%s", myDate.Format("January 02, 2006"))
}

func getAboutMeta(fi os.FileInfo) (string, string) {
	id := fi.Name()[:len(fi.Name())-3]
	title := strings.Split(string(getFile("_about/"+fi.Name())), "\n")[0][2:]

	return id, title
}

func writeIndex() {
	var b bytes.Buffer
	b.WriteString(getLayoutStart(true))
	writePostsSection(&b)
	b.WriteString(getLayoutEnd())
	writeFile("index", b)
}

func writePostsSection(b *bytes.Buffer) {
	b.WriteString(`<section>
							<ul class="post-list">`)

	posts := getDir("_posts")
	for i := len(posts) - 1; i >= 0; i-- {
		_, date, title := getPostMeta(posts[i])
		dateFolder := strings.ReplaceAll(date, "-", "/")
		path := "/posts/" + dateFolder + "/" + strings.ReplaceAll(title, " ", "-")

		b.WriteString("<li><a class=\"post-link\" href=\"" + path + "\">" + title + "</a> <time datetime=\"2017-01-15T00:00:00+00:00\">" + convertDate(date) + "</time></li>\n")
	}
	b.WriteString(`
					</div>
				</ul>
		</section>`)
}

func writePosts() {
	posts := getDir("_posts")

	for i := 0; i < len(posts); i++ {
		_, date, title := getPostMeta(posts[i])
		var b bytes.Buffer
		b.WriteString(getPostPage(title, convertDate(date)))
		b.Write(blackfriday.MarkdownCommon(getFile("_posts/" + posts[i].Name())))
		b.WriteString(getLayoutEnd())

		dateFolder := strings.ReplaceAll(date, "-", "/")
		dir := "posts/" + dateFolder + "/" + strings.ReplaceAll(title, " ", "-")
		os.MkdirAll(dir, 0755)
		writeFile(dir+"/index", b)
	}
}

func writeCSS() {
	var b bytes.Buffer
	b.WriteString(getCSS())
	writeCSSFile("assets/styte", b)
}

func writePostsPage() {
	posts := getDir("_posts")
	var b bytes.Buffer

	b.WriteString(getLayoutStart(false))
	b.WriteString("<h1>All posts</h1>\n")
	b.WriteString("<nav class=\"posts\"><ul>\n")

	for i := len(posts) - 1; i >= 0; i-- {
		_, date, title := getPostMeta(posts[i])

		dateFolder := strings.ReplaceAll(date, "-", "/")
		path := "/posts/" + dateFolder + "/" + strings.ReplaceAll(title, " ", "-")

		b.WriteString("<li><a class=\"post-link\" href=\"" + path + "\">" + title + "</a><span class=\"date\">" + convertDate(date) + "</span></li>\n")
	}

	b.WriteString(getLayoutEnd())
	writeFile("pages/index", b)
}

func writeRSS() {
	posts := getDir("_posts")
	var b bytes.Buffer
	b.WriteString("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n")
	b.WriteString("<rss version=\"2.0\" xmlns:atom=\"http://www.w3.org/2005/Atom\">\n")
	b.WriteString("<channel>\n")
	b.WriteString("<title>Nauk17</title>\n")
	b.WriteString("<description>For the Future</description>\n")
	b.WriteString("<link>https://puerrtigris.github.io/</link>\n")
	for i := len(posts) - 1; i >= 0; i-- {
		_, date, title := getPostMeta(posts[i])

		dateFolder := strings.ReplaceAll(date, "-", "/")
		path := "posts/" + dateFolder + "/" + strings.ReplaceAll(title, " ", "-")
		b.WriteString("<item>\n")
		b.WriteString("<title>" + title + "</title>\n")
		b.WriteString("<link>https://puerrtigris.github.io/" + path + "</link>\n")
		b.WriteString("</item>\n")
	}
	b.WriteString("</channel>\n")
	b.WriteString("</rss>")
	writeRSSFile("rss", b)

}

func writeAbout() {
	pages := getDir("_about")

	for i := 0; i < len(pages); i++ {
		_, _ = getAboutMeta(pages[i])

		var b bytes.Buffer
		b.WriteString(getLayoutStart(false))
		b.WriteString(`<section>`)
		b.Write(blackfriday.MarkdownCommon(getFile("_about/" + pages[i].Name())))
		b.WriteString(`</section>`)
		b.WriteString(getLayoutEnd())

		writeFile("about/index", b)
	}
}

func createFilesAndDirs() {
	os.MkdirAll("_sections", 0755)
	os.MkdirAll("_posts", 0755)
	os.MkdirAll("_about", 0755)

	os.MkdirAll("posts", 0755)
	os.MkdirAll("assets", 0755)
	os.MkdirAll("about", 0755)
	os.MkdirAll("images", 0755)
	os.MkdirAll("pages", 0755)
}

func main() {
	createFilesAndDirs()
	writeCSS()
	writeIndex()
	writePosts()
	writePostsPage()
	writeAbout()
	writeRSS()
}

func getCSS() string {
	return `
	blockquote, body, dd, dl, figure, h1, h2, h3, h4, h5, h6, hr, ol, p, pre, ul {
		margin: 0;
		padding: 0
	}
	
	body {
		font: 300 14px/1.5 -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica, Arial, sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol";
    	font-family: "Overpass Mono", monospace;		
		color: #333;
		background-color: #fff
	}
	
	blockquote, dl, figure, h1, h2, h3, h4, h5, h6, ol, p, pre, table, ul {
		margin-bottom: 1rem
	}
	
	img {
		max-width: 100%;
		vertical-align: middle
	}
	
	figure > img {
		display: block
	}
	
	figcaption {
		font-size: .875rem
	}
	
	ol, ul {
		margin-left: 2rem
	}
	
	li > ol, li > ul {
		margin-bottom: 0
	}
	
	h1, h2, h3, h4, h5, h6 {
		font-weight: 300
	}
	
	a {
		color: #0645ad;
		text-decoration: none
	}
	
	a:hover {
		text-decoration: underline
	}
	
	blockquote {
		color: #6a6a6a;
		padding-left: 1rem;
		border-left: 2px solid #eaeaea;
		font-style: italic;
		font-size: .875rem
	}
	
	blockquote > :last-child {
		margin-bottom: 0
	}
	
	code, pre {
		background-color: #fcfcfc
	}
	
	code {
		padding: 1px 5px;
		font-family: Inconsolata, Monaco, Consolas, monospace;
		color: #f14e32
	}
	
	pre {
		padding: 8px 12px;
		overflow-x: auto;
		border: 1px solid #eaeaea
	}
	
	pre > code {
		border: 0;
		padding-right: 0;
		padding-left: 0;
		tab-size: 4;
		color: inherit
	}
	
	table {
		width: 100%;
		max-width: 100%;
		border-collapse: separate;
		border-spacing: 0;
		table-layout: fixed
	}
	
	td, th {
		padding: 0.5rem;
		line-height: inherit
	}
	
	th {
		text-align: left;
		vertical-align: bottom;
		border-bottom: 2px solid #eaeaea
	}
	
	td {
		vertical-align: top;
		border-bottom: 1px solid #eaeaea
	}
	
	hr {
		border: none;
		border-top: 1px solid #f7f7f7;
		margin: 2rem auto
	}
	
	.underlined {
		flex: 1;
		text-decoration: none;
		background-image: linear-gradient(to right, #ff0 0, #ff0 100%);
		background-position: 0 1.2em;
		background-size: 0 100%;
		background-repeat: no-repeat;
		transition: background .5s
	}
	
	.underlined:hover {
		background-size: 100% 100%
	}
	
	.underlined--thin {
		background-image: linear-gradient(to right, #000 0, #000 100%)
	}
	
	.underlined--thick {
		background-position: 0 -0.1em
	}
	
	.underlined--offset {
		background-position: 0 0.2em;
		box-shadow: inset 0 -.5em 0 0 white
	}
	
	.underlined--gradient {
		background-position: 0 -0.1em;
		background-image: linear-gradient(to right, #ff0 0, #90ee90 100%)
	}
	
	.underlined--reverse {
		background-position: 100% -0.1em;
		transition: background 1s;
		background-image: linear-gradient(to right, #ff0 0, #ff0 100%)
	}
	
	.site-header {
		border-bottom: 1px solid #eaeaea;
		margin-top: -2rem;
		max-width: 48rem
	}
	
	.site-header p {
		font-size: .875rem
	}
	
	.site-header .site-intro {
		font-size: 0.9rem
	}
	
	.smallcap {
		font-size: 1.5rem;
		font-weight: bold
	}
	
	.smallcap a, .smallcap a:hover {
		text-decoration: none;
		letter-spacing: 2px;
		background-color: #333;
		color: #eee;
		padding-left: 0.5rem;
		padding-right: 0.5rem;
		padding-top: 0.2rem;
		padding-bottom: 0.2rem
	}
	
	.page-content {
		position: relative;
		padding: 2rem 1.5rem;
		margin: 1rem auto;
		box-sizing: border-box;
		max-width: 48rem
	}
	
	.home section + section {
		margin-top: 2rem;
		max-width: 48rem
	}

	.post-list {
		margin-left: 4em;
	}
	
	.post-list > li {
		margin-bottom: .5rem;
		margin-left: -2rem
	}
	
	.post-list > li a {
		color: #333;
		text-decoration: none;
		font-weight: normal
	}
	
	.post-list > li a:hover {
		color: #0645ad;
		text-decoration: underline
	}
	
	.post-list > li time {
		font-size: .875rem;
		color: #aaa;
		display: inline-block
	}
	
	@media screen and (max-width: 600px) {
		.post-list > li time {
			display: block;
			font-size: .875rem
		}
	}
	
	.tag-title {
		color: #0645ad
	}
	
	.post-header {
		margin-bottom: 2rem
	}
	
	.post-title {
		font-size: 2rem;
		letter-spacing: -1px;
		line-height: 1.2;
		margin-bottom: 0.5rem;
		font-weight: bold
	}
	
	.post-meta {
		font-size: .875rem;
		font-family: Inconsolata, Monaco, Consolas, monospace;
		color: #aaa
	}
	
	.post-meta a, .post-meta a:visited {
		color: #6a6a6a
	}
	
	.post-meta .tags a, .post-meta .tags a:visited {
		background: #eaeaea;
		padding: 0.1rem 0.5rem
	}
	
	.post-content {
		margin-bottom: 2rem;
		font-weight: normal
	}
	
	.post-content h1, .post-content h2, .post-content h3, .post-content h4, .post-content h5, .post-content h6 {
		margin-top: 2rem;
		font-weight: normal
	}
	
	.post-content h1, .post-content h2 {
		font-size: 2rem
	}
	
	.post-content h3 {
		font-size: 1.5rem
	}
	
	.post-content h4 {
		font-size: 1.25rem
	}
	
	.post-content h5, .post-content h6 {
		font-size: 1rem
	}
	
	.copyright {
		margin-top: 2rem;
		font-size: .875rem;
		font-family: Inconsolata, Monaco, Consolas, monospace
	}
	
	.copyright p {
		color: #aaa
	}
	
	.copyright p a, .copyright p a:visited {
		color: #6a6a6a
	}
	
	.highlight {
		background-color: #fcfcfc;
		color: #586e75
	}
	
	.highlight .c {
		color: #93a1a1
	}
	
	.highlight .err {
		color: #586e75
	}
	
	.highlight .g {
		color: #586e75
	}
	
	.highlight .k {
		color: #859900
	}
	
	.highlight .l {
		color: #586e75
	}
	
	.highlight .n {
		color: #586e75
	}
	
	.highlight .o {
		color: #859900
	}
	
	.highlight .x {
		color: #cb4b16
	}
	
	.highlight .p {
		color: #586e75
	}
	
	.highlight .cm {
		color: #93a1a1
	}
	
	.highlight .cp {
		color: #859900
	}
	
	.highlight .c1 {
		color: #93a1a1
	}
	
	.highlight .cs {
		color: #859900
	}
	
	.highlight .gd {
		color: #2aa198
	}
	
	.highlight .ge {
		color: #586e75;
		font-style: italic
	}
	
	.highlight .gr {
		color: #dc322f
	}
	
	.highlight .gh {
		color: #cb4b16
	}
	
	.highlight .gi {
		color: #859900
	}
	
	.highlight .go {
		color: #586e75
	}
	
	.highlight .gp {
		color: #586e75
	}
	
	.highlight .gs {
		color: #586e75;
		font-weight: bold
	}
	
	.highlight .gu {
		color: #cb4b16
	}
	
	.highlight .gt {
		color: #586e75
	}
	
	.highlight .kc {
		color: #cb4b16
	}
	
	.highlight .kd {
		color: #268bd2
	}
	
	.highlight .kn {
		color: #859900
	}
	
	.highlight .kp {
		color: #859900
	}
	
	.highlight .kr {
		color: #268bd2
	}
	
	.highlight .kt {
		color: #dc322f
	}
	
	.highlight .ld {
		color: #586e75
	}
	
	.highlight .m {
		color: #2aa198
	}
	
	.highlight .s {
		color: #2aa198
	}
	
	.highlight .na {
		color: #586e75
	}
	
	.highlight .nb {
		color: #B58900
	}
	
	.highlight .nc {
		color: #268bd2
	}
	
	.highlight .no {
		color: #cb4b16
	}
	
	.highlight .nd {
		color: #268bd2
	}
	
	.highlight .ni {
		color: #cb4b16
	}
	
	.highlight .ne {
		color: #cb4b16
	}
	
	.highlight .nf {
		color: #268bd2
	}
	
	.highlight .nl {
		color: #586e75
	}
	
	.highlight .nn {
		color: #586e75
	}
	
	.highlight .nx {
		color: #586e75
	}
	
	.highlight .py {
		color: #586e75
	}
	
	.highlight .nt {
		color: #268bd2
	}
	
	.highlight .nv {
		color: #268bd2
	}
	
	.highlight .ow {
		color: #859900
	}
	
	.highlight .w {
		color: #586e75
	}
	
	.highlight .mf {
		color: #2aa198
	}
	
	.highlight .mh {
		color: #2aa198
	}
	
	.highlight .mi {
		color: #2aa198
	}
	
	.highlight .mo {
		color: #2aa198
	}
	
	.highlight .sb {
		color: #93a1a1
	}
	
	.highlight .sc {
		color: #2aa198
	}
	
	.highlight .sd {
		color: #586e75
	}
	
	.highlight .s2 {
		color: #2aa198
	}
	
	.highlight .se {
		color: #cb4b16
	}
	
	.highlight .sh {
		color: #586e75
	}
	
	.highlight .si {
		color: #2aa198
	}
	
	.highlight .sx {
		color: #2aa198
	}
	
	.highlight .sr {
		color: #dc322f
	}
	
	.highlight .s1 {
		color: #2aa198
	}
	
	.highlight .ss {
		color: #2aa198
	}
	
	.highlight .bp {
		color: #268bd2
	}
	
	.highlight .vc {
		color: #268bd2
	}
	
	.highlight .vg {
		color: #268bd2
	}
	
	.highlight .vi {
		color: #268bd2
	}
	
	.highlight .il {
		color: #2aa198
	}
	`
}
