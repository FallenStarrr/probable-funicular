package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)





func main() {
	type Article struct {
		Id int `json: "id"`
		Title string `json: "title"`
		Body string `json: "body"`
		UserId int `json: "userId"`
}


   article := Article{}

	  url := "https://jsonplaceholder.typicode.com/posts"
    r := chi.NewRouter()
    r.Use(middleware.Logger)
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			    res, _ := http.Get("https://jsonplaceholder.typicode.com/posts/101")
					re, _ := ioutil.ReadAll(res.Body)
					w.Write(re)
		})
    r.Post("/", func(w http.ResponseWriter, r *http.Request) {


		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
		}



		var unmarshalErr error
		unmarshalErr = json.Unmarshal(reqBody, article)
		if err != nil {
			fmt.Println(unmarshalErr)
		}

		postRequest, err := json.Marshal(article)
		if err != nil {
			fmt.Println(err)
		}

		req, err	:=  http.NewRequest(http.MethodPost, url, bytes.NewBuffer(postRequest))
		if err != nil {
     fmt.Println(err)
		}

		  rep, err := http.DefaultClient.Do(req)
			resp, err := ioutil.ReadAll(rep.Body)
			if err != nil {
				 fmt.Println(err)
			}

        w.Write([]byte(resp))
    })
    http.ListenAndServe(":3000", r)
}
