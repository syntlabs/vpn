package main

import (
	"http"
	"fmt"
)

func frontCallback(w http.ResponseWriter, r *http.Request) {
	changeLayoutCallBack := "hello"
	fmt.Fprintf(w, `<script>localStorage.setItem("changeLayoutCallBack", "%s");</script>`, changeLayoutCallBack)
}