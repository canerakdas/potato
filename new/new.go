package new

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"github.com/canerakdas/potato/color"
	"strings"
	"text/template"
)

type dir struct {
	Name          string
	NodePackage   string
	Source        string
	Components    string
	Stylesheets   string
	DefaultHtml   string
	DefaultCss    string
	DefaultJS     string
	DefaultApp    string
	WebpackConfig string
	BabelConfig   string
}

var tpl *template.Template

func init() {
	gopath := os.Getenv("GOPATH")

	root := filepath.Join(gopath, "/src/github.com/canerakdas/potato")

	tpl = template.Must(template.ParseGlob(filepath.Join(root, "templates", "*")))
}

func Create(name string) {
	t := dir{
		Name:          name,
		Source:        "source",
		Stylesheets:   "stylesheets",
		Components:    "components",
		NodePackage:   "package.json",
		DefaultHtml:   "index.html",
		DefaultCss:    "default.css",
		DefaultJS:     "index.js",
		DefaultApp:    "app.js",
		WebpackConfig: "webpack.config.js",
		BabelConfig:   ".babelrc",
	}

	fmt.Println(color.Bright, "[1/4]", color.Reset, " Creating folder structure")

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	var dirPath = dir{
		Name:        filepath.Join(pwd, t.Name),
		Source:      filepath.Join(pwd, t.Name, t.Source),
		Stylesheets: filepath.Join(pwd, t.Name, t.Source, t.Stylesheets),
		Components:  filepath.Join(pwd, t.Name, t.Source, t.Components),
	}

	createFolder(dirPath.Name)
	createFolder(dirPath.Source)
	createFolder(dirPath.Stylesheets)
	createFolder(dirPath.Components)

	fmt.Println(color.Bright, "[2/4]", color.Reset, " Creating files")

	createTemplate(dirPath.Name, t.NodePackage, t)
	createTemplate(dirPath.Name, t.WebpackConfig, t)
	createTemplate(dirPath.Name, t.BabelConfig, dir{})
	createTemplate(dirPath.Stylesheets, t.DefaultCss, dir{})
	createTemplate(dirPath.Source, t.DefaultHtml, t)
	createTemplate(dirPath.Source, t.DefaultJS, t)
	createTemplate(dirPath.Source, t.DefaultApp, t)
	fmt.Println(color.Bright, "[3/4]", color.Reset, " Installing npm packages")

	app := "npm"
	args := []string{
		"i",
		"-D",
		"@babel/core",
		"@babel/preset-env",
		"@babel/preset-react",
		"babel-loader",
		"css-loader",
		"html-loader",
		"html-webpack-plugin",
		"react",
		"react-dom",
		"style-loader",
		"webpack",
		"webpack-cli",
		"webpack-dev-server",
	}

	for i, v := range args {
		if i > 1 {
			fmt.Println(
				color.FgGreen, "\t Package ", color.Reset,
				color.Underscore, v, color.Reset)
		}
	}

	cmd := exec.Command(app, args...)
	cmd.Dir = dirPath.Name

	_, err = cmd.Output()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(color.Bright, "[4/4] ", color.Reset, t.Name, " successfully created!")

	text := make([]string, 32-len(t.Name))
	text[0] = t.Name
	fmt.Println(" +------------------------------------+")
	fmt.Println(" | cd", strings.Join(text[:], " "), "|")
	fmt.Println(" | npm run start                      |")
	fmt.Println(" +------------------------------------+")
}

func createTemplate(path string, name string, template dir) {
	fullPath := filepath.Join(path, name)
	file, err := os.Create(fullPath)

	if err != nil {
		log.Fatal(err)
	}

	err = tpl.ExecuteTemplate(file, name, template)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(
		color.FgGreen, "\t Create ", color.Reset,
		color.Underscore, fullPath, color.Reset)

	defer file.Close()
}

func createFolder(path string) {
	err := os.Mkdir(path, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(
		color.FgGreen, "\t Create ", color.Reset,
		color.Underscore, path, color.Reset)
}
