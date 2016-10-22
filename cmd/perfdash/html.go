package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl = template.Must(template.New("index.html").Parse(indexHTML))

func index(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}

	var v struct {
		DebugVars *debugVars
		ID        int
	}

	u := req.URL.Query().Get("baseurl")
	if u == "" {
		u = *baseURL
	}

	user, pass, _ := req.BasicAuth()

	var err error
	v.ID, v.DebugVars, err = fetchDebugVars(u, user, pass)
	if err == errAuth {
		w.Header().Set("WWW-Authenticate", `Basic realm="perfdash"`)
		http.Error(w, "auth pls", 401)
		return
	} else if err != nil {
		http.Error(w, err.Error(), 500)
		log.Println(err)
		return
	}
}

const indexHTML = `
<h1>perfdash</h1>
Open dev tools to see the full /debug/vars data.
{{$id := .ID}}
{{range $k, $v := .DebugVars.Latency}}
	<img src="/histogram.png?name={{$k}}&id={{$id}}">
{{end}}
<script>
console.log("/debug/vars", {{.DebugVars.Raw}});
</script>
`
