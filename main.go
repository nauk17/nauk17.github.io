package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"github.com/russross/blackfriday"
)

func getLayoutStart(title string) string {
	return `<!DOCTYPE html>
	<html>
		<head>
			<meta charset="utf-8">
			<meta name="viewport" content="width=device-width, initial-scale=1">
			<link rel="icon" href="/images/icon.jpg" type="image/gif" sizes="16x16">
			<link href="https://fonts.googleapis.com/css?family=IBM+Plex+Sans:300,400,400i,500" rel="stylesheet">
			<link href="https://fonts.googleapis.com/css?family=IBM+Plex+Mono:400" rel="stylesheet">
			<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css">
			<link href="assets/styte.css" rel="stylesheet">
			<link href="/assets/styte.css" rel="stylesheet">
			<title>` + title + `</title>
			<script>
				document.addEventListener('DOMContentLoaded', function(event) {
					if (localStorage.getItem('theme') === 'dark') {
						setDarkTheme();
					} else {
						setLightTheme();
					}
				});

				function toggleTheme(event) {
					event.preventDefault();

					if (document.body.className === 'dark') {
						setLightTheme();
					} else {
						setDarkTheme();
					}
				}

				function setLightTheme() {
					document.body.className = 'light';
					document.getElementsByClassName('toggle-theme')[0].children[0].innerHTML = '<i class="fa fa-moon-o"></i>';
					localStorage.setItem('theme', 'light');
				}

				function setDarkTheme() {
						document.body.className = 'dark';
						document.getElementsByClassName('toggle-theme')[0].children[0].innerHTML = '<i class="fa fa-sun-o"></i>';
						localStorage.setItem('theme', 'dark');
				}
			</script>
		</head>
		<body>
			<div class="header"> 
				<nav class="nav"> 
					<div class="nav-content"> 
						<h1 class="logo">
							<a href="/" style="font-size:29px;font-weight:bold;color:rgba(51,51,51,0.8);">Nauk<span>17</span></a>
						</h1> 
						<ul class="navbar"> 
							<li><a href="/about">About me</a></li> 
							<li><a href="/rss.xml" target="_blank">RSS</a></li> 
						</ul>
					</div> 
				</nav> 
			</div>
			<div class="container">
			`
}

func getLayoutEnd() string {
	return `
			</div>
			<footer>
				<small>© 2020 Quan Nguyen</small>
				<small class="toggle-theme">
					<a href="#" onclick="toggleTheme(event)">Dark</a>
				</small>
			</footer>
		</body>
	</html>`
}

func getCSS() string {
	return `
	body {
		font-family: "Overpass Mono",monospace;
		line-height: 1.875;
		font-weight: 400;
		font-size: 16px;
	}

	body.light {
		background-color: #fdffff;
		color: rgba(51,51,51,0.8);
	}

	body.dark {
		background-color: #141c2b;
		color: #d9dce4;
	}

	.header {
		border-bottom: 1px solid #d9d9d9;
		margin-bottom: .75em;
		min-height: 80px;
		top: 0;
		left: 0;
		position: fixed;
		width: 100%;
		background:#FFF;
	}

	.navbar {
		float: right;
		margin: 0;
		position: relative;
		padding: 0;
		pointer-events: all;
		cursor: pointer;
	}

	.logo {
		float: left;
		margin: 0 0 1em 0;
		cursor: pointer;
		letter-spacing: 0.8px;
		font-size: 20px;
		line-height: 28px;
		font-weight: 300;
	}

	.navbar li {
		display: inline-block;
		padding: 0 .6em;
	}

	@media (max-width: 1023.98px) {
		.nav-content {
			margin: 28px auto 48px auto;
		}
	}

	@media (min-width: 1024px) {
		.nav-content {
			margin: 48px auto;
		}
	}

	.nav-content {
		margin: auto;
		padding: 1.5em;
		margin-left: auto;
		margin-right: auto;
		max-width: 800px;
		font-weight: normal;
	}

	/* h1, h2, h3 {
	 	font-weight: 300;
	 }*/

	h1 {
		font-size: 26.79296875px;
		margin-top: 24px;
	}

	h2 {
		margin-bottom: 0;
		line-height: 1.2em;
		margin-top: 1em;
	}

	h3 {
		font-size: 19px;
		margin-top: 40px;
	}

	body.light h1,
	body.light h2,
	body.light h3 {
		color: rgba(51,51,51,0.8);
	}

	body.dark h1,
	body.dark h2,
	body.dark h3 {
		color: #c8ddff;
	}

	@media (max-width: 1023.98px) {
		.container {
			margin: 28px auto 48px auto;
		}
	}

	@media (min-width: 1024px) {
		.container {
			margin: 48px auto;
		}
	}

	.container {
		margin-left: 80px;
		padding: 1.5em;
		margin-left: auto;
		margin-right: auto;
		max-width: 800px;
		font-weight: normal;
	}

	nav ul {
		list-style-type: none;
		padding: 0;
	}

	nav li {
		margin-left: 25px;
	}

	nav li .date {
		float: right;
		width: 104px;
		margin-top: 0;
		font-size: 0.8em;
		color: #777777;
		font-style: italic;
	}

	.year {
		margin-bottom: 0;
	}

	.all-posts {
		font-size: 13.47368421052632px;
		margin: 16px 0 32px 0;
	}

	/* a {
		text-decoration: none;
	} */

	a {
		position: relative;
		color: #000;
		text-decoration: none;
	}

	a:hover {
		color: #0086B3;
	}

	a::before {
		content: "";
		position: absolute;
		width: 100%;
		height: 2px;
		bottom: 0;
		left: 0;
		background-color:#0086B3;
		visibility: hidden;
		transform: scaleX(0);
		transition: all 0.3s ease-in-out 0s;
	}

	:hover::before {
		visibility: visible;
		transform: scaleX(1);
	}

	/* a:hover {
		border-bottom: 3px solid #0086B3;
	} */

	body.light a {
		color: #000;
	}

	body.light .navbar li > a {
		font-weight:bold;
		color:rgba(51,51,51,0.8);
	}

	body.dark a {
		color: #7da7ef;
	}
	pre {
		overflow: auto;
		padding: 0.25rem 0.75rem;
		margin-bottom: 32px;
	}

	body.light pre {
		background-color: #f4f7ff;
	}

	body.dark pre {
		background-color: #1a2231;
	}

	code {
		font-size: 0.875em;
		font-family: 'IBM Plex Mono', monospace;
	}

	table {
		border-collapse: collapse;
		width: 100%;
		margin-bottom: 32px;
	}

	@media (max-width: 1023.98px) {
		table {
			font-size: 14px;
		}
	}

	@media (min-width: 1024px) {
		table {
			font-size: 15px;
		}
	}

	body.light tr {
		border-bottom: 0.5px solid #bdc5d8;
	}

	body.dark tr {
		border-bottom: 0.5px solid #424957;
	}

	th {
		text-align: left;
		font-weight: 500;
		padding: 12px;
		white-space: nowrap;
	}

	td {
		padding: 12px;
		white-space: nowrap;
	}
	footer { 
		align-items: center; 
		text-align: center;
	} 

	footer small {
		display: inline;
	}
	`
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
	title := strings.Split(string(getFile("_posts/"+fi.Name())), "\n")[0][2:]

	return id, date, title
}

func getPageMeta(fi os.FileInfo) (string, string) {
	id := fi.Name()[:len(fi.Name())-3]
	title := strings.Split(string(getFile("_pages/"+fi.Name())), "\n")[0][2:]

	return id, title
}

func writeIndex() {
	var b bytes.Buffer
	b.WriteString(getLayoutStart(getSiteTitle()))
	b.Write(blackfriday.MarkdownCommon(getFile("_sections/header.md")))
	writePostsSection(&b)
	b.WriteString(getLayoutEnd())
	writeFile("index", b)
}

func writePostsSection(b *bytes.Buffer) {
	b.WriteString("<nav class=\"posts\"><ul>")

	posts := getDir("_posts")
	var years []string
	for i := len(posts) - 1; i >= 0; i-- {
		_, date, _ := getPostMeta(posts[i])
		y := strings.Split(date, "-")[0]
		i := sort.Search(len(years), func(i int) bool { return y <= years[i] })
		if i < len(years) && years[i] == y {
			continue
		}
		years = append(years, y)
	}
	// limit := int(math.Max(float64(len(posts))-5, 0))
	for _, year := range years {
		b.WriteString("<h1 class=\"year\">" + year + "</h1>")
		for i := len(posts) - 1; i >= 0; i-- {
			_, date, title := getPostMeta(posts[i])

			y := strings.Split(date, "-")[0]

			if y != year {
				continue
			}

			dateFolder := strings.ReplaceAll(date, "-", "/")
			path := "/posts/" + dateFolder + "/" + strings.ReplaceAll(title, " ", "-")

			b.WriteString("<li><a class=\"post-link\" href=\"" + path + "\">" + title + "</a><span class=\"date\">" + date + "</span></li>\n")
		}
	}

	b.WriteString("</ul></nav><p class=\"all-posts\"><a href=\"all-posts.html\">All posts</a></p>")
}

// func writePagesSection(b *bytes.Buffer) {
// 	b.WriteString("<h2>Pages</h2><nav class=\"pages\"><ul>")

// 	pages := getDir("_pages")

// 	for i := 0; i < len(pages); i++ {
// 		id, title := getPageMeta(pages[i])

// 		b.WriteString("<li><a href=\"pages/" +
// 			id + ".html\">" +
// 			title + "</a></li>\n")
// 	}

// 	b.WriteString("</ul></nav>")
// }

func writePosts() {
	posts := getDir("_posts")

	for i := 0; i < len(posts); i++ {
		_, date, title := getPostMeta(posts[i])

		var b bytes.Buffer
		b.WriteString(getLayoutStart(title + " – " + getSiteTitle()))
		// b.WriteString("<p><a href=\"../index.html\">←</a></p>")
		b.WriteString("<p class=\"date\">" + date + "</p>")
		b.Write(blackfriday.MarkdownCommon(getFile("_posts/" + posts[i].Name())))
		// b.WriteString("<p><a href=\"../index.html\">←</a></p>")
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

	b.WriteString(getLayoutStart("All posts – " + getSiteTitle()))
	// b.WriteString("<p><a href=\"index.html\">←</a></p>")
	b.WriteString("<h1>All posts</h1>\n")
	b.WriteString("<nav class=\"posts\"><ul>\n")

	for i := len(posts) - 1; i >= 0; i-- {
		_, date, title := getPostMeta(posts[i])

		dateFolder := strings.ReplaceAll(date, "-", "/")
		path := "/posts/" + dateFolder + "/" + strings.ReplaceAll(title, " ", "-")

		b.WriteString("<li><a class=\"post-link\" href=\"" + path + "\">" + title + "</a><span class=\"date\">" + date + "</span></li>\n")

		// b.WriteString("<li><span class=\"date\">" + date +
		// 	"</span><a href=\"posts/" +
		// 	id + ".html\">" +
		// title + "</a></li>\n")
	}

	// b.WriteString("</ul></nav><p><a href=\"index.html\">←</a></p>")
	b.WriteString(getLayoutEnd())
	writeFile("all-posts", b)
}

func writeRSS() {
	posts := getDir("_posts")
	var b bytes.Buffer
	b.WriteString("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n")
	b.WriteString("<rss version=\"2.0\" xmlns:atom=\"http://www.w3.org/2005/Atom\">\n")
	b.WriteString("<channel>\n")
	b.WriteString("<title>Nauk17</title>\n")
	b.WriteString("<description>For the Future</description>\n")
	b.WriteString("<link>https://nauk17.github.io/</link>\n")
	for i := len(posts) - 1; i >= 0; i-- {
		_, date, title := getPostMeta(posts[i])

		dateFolder := strings.ReplaceAll(date, "-", "/")
		path := "posts/" + dateFolder + "/" + strings.ReplaceAll(title, " ", "-")
		b.WriteString("<item>\n")
		b.WriteString("<title>" + title + "</title>\n")
		b.WriteString("<link>https://nauk17.github.io/" + path + "</link>\n")
		b.WriteString("</item>\n")
	}
	b.WriteString("</channel>\n")
	b.WriteString("</rss>")
	writeRSSFile("rss", b)

}

// func writePages() {
// 	pages := getDir("_pages")

// 	for i := 0; i < len(pages); i++ {
// 		fileName, title := getPageMeta(pages[i])

// 		var b bytes.Buffer
// 		b.WriteString(getLayoutStart(title + " – " + getSiteTitle()))
// 		// b.WriteString("<p><a href=\"../index.html\">←</a></p>")
// 		b.Write(blackfriday.MarkdownCommon(getFile("_pages/" + pages[i].Name())))
// 		// b.WriteString("<p><a href=\"../index.html\">←</a></p>")
// 		b.WriteString(getLayoutEnd())

// 		writeFile("pages/"+fileName, b)
// 	}
// }

func writeAbout() {
	pages := getDir("_about")

	for i := 0; i < len(pages); i++ {
		_, title := getPageMeta(pages[i])

		var b bytes.Buffer
		b.WriteString(getLayoutStart(title + " – " + getSiteTitle()))
		b.Write(blackfriday.MarkdownCommon(getFile("_about/" + pages[i].Name())))
		b.WriteString(getLayoutEnd())

		writeFile("about/index", b)
	}
}

func createFilesAndDirs() {
	os.MkdirAll("_sections", 0755)
	os.MkdirAll("_posts", 0755)
	os.MkdirAll("_pages", 0755)
	os.MkdirAll("_about", 0755)

	os.MkdirAll("posts", 0755)
	os.MkdirAll("pages", 0755)
	os.MkdirAll("assets", 0755)
	os.MkdirAll("about", 0755)
	os.MkdirAll("images", 0755)
}

func main() {
	createFilesAndDirs()
	writeCSS()
	writeIndex()
	writePosts()
	writePostsPage()
	// writePages()
	writeAbout()
	writeRSS()
}
