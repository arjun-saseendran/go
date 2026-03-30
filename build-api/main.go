package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for course
type Course struct {
	CourseId    string  `json:"courseid`
	CourseName  string  `json:"coursename`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake db
var courses []Course

// check fields middleware
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	r := mux.NewRouter()
	// seeding
	courses = append(courses, Course{CourseId: "3", CourseName: "Golang", CoursePrice: 4999, Author: &Author{Fullname: "Arjun Saseendran", Website: "go.dev"}})
	courses = append(courses, Course{CourseId: "2", CourseName: "Node", CoursePrice: 2999, Author: &Author{Fullname: "Aswini Ramachadran", Website: "node.dev"}})

	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// create server
	log.Fatal(http.ListenAndServe(":3000", r))

}

// controller
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Home page</h1>"))
}

// courses
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

// course
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// get id
	params := mux.Vars(r)
	// find match id
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
	return
}

// create course
func createOneCourse(w http.ResponseWriter, r *http.Request) {
	// set header
	w.Header().Set("Content-Type", "application/json")
	// hanlde empty body
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please add some data!")
	}
	// create couse
	var newCourse Course
	_ = json.NewDecoder(r.Body).Decode(&newCourse)

	if newCourse.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside json")
		return
	}
	// generate unique id
	rand.Seed(time.Now().UnixNano())
	// add id to new course
	newCourse.CourseId = strconv.Itoa(rand.Intn(100))
	// Handle duplicate
	for _, course := range courses {
		if course.CourseName == newCourse.CourseName {
			json.NewEncoder(w).Encode("Course already exist!")
			return
		}
	}
	// add new course
	courses = append(courses, newCourse)
	json.NewEncoder(w).Encode(newCourse)
	return

}

// Update data
func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	// Set header
	w.Header().Set("Content-Type", "application/json")
	// get id from parmas
	params := mux.Vars(r)
	// find id and update content
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index:+1]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

}

// Delete data
func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	// Set Header
	w.Header().Set("Content-Type", "application/json")
	// Get id from params
	params := mux.Vars(r)
	// find id and delete
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(courses)

}
