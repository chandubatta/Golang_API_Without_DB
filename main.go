package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Models for Course-File
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrise int     `json:"prise"`
	Author      *Author `json:"author"`
}
type Author struct {
	Fullname string `json:"fullname"`
	WebSite  string `json:"website"`
}

// Fake DB
var courses []Course

// Middleware , helper-file
func (c *Course) isempty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {
	fmt.Println("Building API's Without DataBase")
	r := mux.NewRouter()

	//Seeding
	courses = append(courses, Course{CourseId: "3", CourseName: "Golang", CoursePrise: 99, Author: &Author{Fullname: "Chandu", WebSite: "chandubatta.netlify.app"}})

	courses = append(courses, Course{CourseId: "6", CourseName: "Reactjs", CoursePrise: 89, Author: &Author{Fullname: "chandubatta", WebSite: "chandubatta.netlify.app"}})

	//Routing
	r.HandleFunc("/", ServeHome).Methods("GET")
	r.HandleFunc("/get_all_courses", get_all_courses).Methods("GET")
	r.HandleFunc("/get_one_course/{id}", get_one_course).Methods("GET")
	r.HandleFunc("/creat_course", creat_course).Methods("POST")
	r.HandleFunc("/course/{id}", update_course).Methods("PUT")
	r.HandleFunc("/course/{id}", delete_course).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000", r))
}

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Chandu Online Study</h1>"))
}

func get_all_courses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Takaing All Courses")
	json.NewEncoder(w).Encode(courses)
}

func get_one_course(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One course by id")
	w.Header().Set("content-type", "application/json")

	params := mux.Vars(r)

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course Found")
	return
}

func creat_course(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "applicatioan/json")

	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// what about - {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.isempty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	//TODO: check only if title is duplicate
	// loop, title matches with course.coursename, JSON

	// generate unique id, string
	// append course into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}

func update_course(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update Method")
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	for index, value := range courses {
		if value.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
		}
		var course Course
		_ = json.NewDecoder(r.Body).Decode(&course)
		course.CourseId = params["id"]
		courses = append(courses, course)
		json.NewEncoder(w).Encode(course)
		return
	}
}

func delete_course(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete One record")
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	for index, val := range courses {
		if val.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			_ = json.NewEncoder(w).Encode(params)
			_ = json.NewEncoder(w).Encode("Successfully Deleted")
			break
		}
	}
}
