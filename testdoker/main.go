package main

import (
	"log"
	"net/http"
	"strings"
	"time"

	wkhtml "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func main() {
	http.HandleFunc("/api/v1/pdf", sendPDF)
	log.Println("server started on http://127.0.0.1:8081/api/v1/pdf")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}

func sendPDF(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request started :%s", r.URL.String())
	start := time.Now()
	pdfg, err := wkhtml.NewPDFGenerator()
	if err != nil {
		return
	}
	htmlStr := `
<html>
<body>
<table><tr><td width="20%">
    <img src="https://th.bing.com/th/id/R8791ad17e523902816f8705320ed391d?rik=e87QWBWR%2fAx84w&pid=ImgRaw" width="100px"></img>
</td>
    <td><h1 style="color:red;">This is an html from pdf to test color</h1></td></tr></table>
    <p> There is a sample text </p>
</body>
</html>
`
	pdfg.Title.Set("Sample")
	pdfg.AddPage(wkhtml.NewPageReader(strings.NewReader(htmlStr)))

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		w.Header().Add("Content-Type", "text/plain")
		w.WriteHeader(400)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	w.Header().Add("Content-Type", "application/pdf")
	w.WriteHeader(200)
	_, _ = w.Write(pdfg.Bytes())
	log.Printf("Document generation took: %v\n", time.Since(start))
}
