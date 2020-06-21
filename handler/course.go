package handler

import (
	"course/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func getHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	courseId := r.URL.Query().Get("id")
	id, _ := strconv.ParseUint(courseId, 10, 64)
	course, err := model.Get(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	resp, _ := json.Marshal(course)
	_, _ = w.Write(resp)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	body, _ := ioutil.ReadAll(r.Body)
	var toCreate model.Course
	_ = json.Unmarshal(body, &toCreate)
	result, err := model.Create(toCreate)
	if err != nil {
		log.Println("Create course failed")
		_, _ = w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		log.Println("Course ", result.Name, "created")
	}
	response, err := json.Marshal(result)
	_, _ = w.Write(response)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	courseId := r.URL.Query().Get("id")
	id, _ := strconv.ParseUint(courseId, 10, 64)
	err := model.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getHandler(w, r)
	case "POST":
		postHandler(w, r)
	case "DELETE":
		deleteHandler(w, r)
	}
}

func AllHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	all, err := model.All()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	var body []byte
	if len(all) != 0 {
		body, _ = json.Marshal(all)
	} else {
		body = []byte("[]")
	}
	_, _ = w.Write(body)
}
