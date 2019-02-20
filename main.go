package main

import (
	"crypto/sha1"
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/xushikuan/microlog"
	"io"
	"net/http"
	"sillyhat-wechat/common"
	"sort"
	"strings"
)

const (
	moduleName = "sillyhat-blog"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	log.SetLevel(log.InfoLevel)
	log.SetModuleName(moduleName)
}

func main() {
	log.Info("---------- start server ----------")
	router := mux.NewRouter()
	router.Use(loggingMiddleware)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	router.HandleFunc("/", WeChatValidateUrl)
	log.Fatal(http.ListenAndServe(":18001", router))
}

func WeChatValidateUrl(writer http.ResponseWriter, request *http.Request) {
	_ = request.ParseForm()
	if !validateUrl(writer, request) {
		log.Println("Wechat Service: this http request is not from Wechat platform!")
		return
	}
	log.Println("Wechat Service: validateUrl Ok!")
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//fmt.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func makeSignature(timestamp, nonce string) string {
	sl := []string{common.TOKEN, timestamp, nonce}
	sort.Strings(sl)
	s := sha1.New()
	_, _ = io.WriteString(s, strings.Join(sl, ""))
	return fmt.Sprintf("%x", s.Sum(nil))
}

func validateUrl(w http.ResponseWriter, r *http.Request) bool {
	timestamp := strings.Join(r.Form["timestamp"], "")
	nonce := strings.Join(r.Form["nonce"], "")
	signatureGen := makeSignature(timestamp, nonce)
	signatureIn := strings.Join(r.Form["signature"], "")
	if signatureGen != signatureIn {
		return false
	}
	echostr := strings.Join(r.Form["echostr"], "")
	_, _ = fmt.Fprintf(w, echostr)
	return true
}

func GetToken(writer http.ResponseWriter, request *http.Request) {

}
