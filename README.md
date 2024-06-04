# Learning GO + TailwindCSS config

1. Init project
`go mod init my-go-project`
2. Create directory structure

```
my-go-project/
├── go.mod
├── main.go
├── static/
│   └── css/
│       ├── styles.css
│       └── tailwind.css
├── templates/
│   └── index.html
├── package.json
```

3. Init npm project and install tailwind

`npm init -y`
`npm install tailwindcss`
`npx tailwindcss init`


4. Create TailwindCSS entry file

```
`/* static/css/styles.css */`
@tailwind base;
@tailwind components;
@tailwind utilities;
```
5. Configure **package.json** to build and watch scripts

```  "scripts": {
    "build:css": "tailwindcss build static/css/styles.css -o static/css/tailwind.css",
    "watch": "tailwindcss build -o static/css/tailwind.css --watch"
  },
```

6. Configure tailwind.config.js to serve the html templates

Basically here its serving the content and some added fonts :)

```
/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./templates/*.html"],
  theme: {
    extend: {
      fontFamily: {
        sans: ["Iosevka Aile Iaso", "sans-serif"],
        mono: ["Iosevka Curly Iaso", "monospace"],
        serif: ["Iosevka Etoile Iaso", "serif"],
    },},
  },
  plugins: [],
}
```


7. Basic template

```
<!-- templates/index.html -->
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Go Project with Tailwind CSS</title>
    <link href="/static/css/tailwind.css" rel="stylesheet">
</head>
<body>
    <div class="container mx-auto">
        <h1 class="text-teal-400 text-4xl font-bold text-center mt-8">Hello, Tailwind CSS!</h1>
        <div class="bg-blue-500 text-white font-bold py-2 px-4 rounded">Button</div>
    </div>
</body>
</html>
```
8. Basic web server serving the static files

```
package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
    // Serve static files
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    // Render template
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        tmpl, err := template.ParseFiles("templates/index.html")
        if err != nil {
            log.Fatal(err)
        }
        err = tmpl.Execute(w, nil)
        if err != nil {
            log.Fatal(err)
        }
    })

    // Start server
    log.Println("Listening on :8888")
    err := http.ListenAndServe(":8888", nil)
    if err != nil {
        log.Fatal(err)
    }
}
```

9. Build static files
`npm run build:css`

10. Run the server :)

`go run main.go`
