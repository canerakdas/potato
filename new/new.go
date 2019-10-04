package new

import (
	"fmt"
	"potato/color"
	"log"
	"os"
	"os/exec"
	"path/filepath"
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
	WebpackConfig string
	BabelConfig   string
}

var tpl *template.Template

func init() {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	root := filepath.Join(gopath,"/src/github.com/canerakdas/potato")

	tpl = template.Must(template.ParseGlob(filepath.Join(root, "templates", "*")))
}

func Create(name string) {
	t := dir{
		Name:          name,
		Source:        "source",
		Stylesheets:   "stylesheets",
		NodePackage:   "package.json",
		DefaultHtml:   "index.html",
		DefaultCss:    "default.css",
		DefaultJS:     "main.js",
		WebpackConfig: "webpack.config.js",
		BabelConfig:   ".babelrc",
	}

	fmt.Println(color.Bright,"INFO    ",color.Reset,"▶ Creating React App...")

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	var dirPath = dir{
		Name:        filepath.Join(pwd, t.Name),
		Source:      filepath.Join(pwd, t.Name, t.Source),
		Stylesheets: filepath.Join(pwd, t.Name, t.Source, t.Stylesheets),
	}

	createFolder(dirPath.Name)
	createFolder(dirPath.Source)
	createFolder(dirPath.Stylesheets)

	createTemplate(dirPath.Name, t.NodePackage, t)
	createTemplate(dirPath.Name, t.WebpackConfig, t)
	createTemplate(dirPath.Name, t.BabelConfig, dir{})
	createTemplate(dirPath.Stylesheets, t.DefaultCss, dir{})
	createTemplate(dirPath.Source, t.DefaultHtml, t)
	createTemplate(dirPath.Source, t.DefaultJS, t)
	fmt.Println(color.Bright,"INFO    ",color.Reset,"▶ Installing npm packages")

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

	cmd := exec.Command(app, args...)
	cmd.Dir = dirPath.Name

	stdout, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(color.Bright,"SUCCESS ",color.Reset,"▶ New React App successfully created!")
	fmt.Println(color.Bright,"NPM LOGS",color.Reset)
	fmt.Println(string(stdout))
}

func createTemplate(path string, name string, template dir) {
	fullPath := filepath.Join(path,name)
	file, err := os.Create(fullPath)

	if err != nil {
		log.Fatal(err)
	}

	err = tpl.ExecuteTemplate(file, name, template)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(
		"    ",color.FgGreen, "Create ",color.Reset,
		color.Underscore,fullPath,color.Reset)

	defer file.Close()
}

func createFolder(path string) {
	err := os.Mkdir(path, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(
		"    ",color.FgGreen, "Create ",color.Reset,
		color.Underscore,path,color.Reset)
}
