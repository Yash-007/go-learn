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

// types
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"-"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake db
var courses []Course

// middleware or helper
func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

// controllers

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<h1>Welcome to API by Yash</h1>`))
}

func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func GetOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")

	params := mux.Vars(r)

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode(fmt.Sprintf("No course found with the given id: %v", params["id"]))
}

func CreateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send course data")
	}

	var course Course
	json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside the JSON")
		return
	}

	for _, val := range courses {
		if val.CourseName == course.CourseName {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": "course with the given name already exists",
			})
			return
		}
	}

	rand.NewSource(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))

	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func UpdateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			course.CourseId = params["id"]
			json.NewDecoder(r.Body).Decode(&course)
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": "id not found",
	})
}

func DeleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for i, val := range courses {
		if val.CourseId == params["id"] {
			courses = append(courses[:i], courses[i+1:]...)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message": "course deleted successfully",
			})
			return
		}
	}
	w.WriteHeader(400)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": "id not found",
	})
}

func main() {

	courses = append(courses,
		Course{CourseName: "Reactjs", CourseId: "1", CoursePrice: 234, Author: &Author{Fullname: "Yash", Website: "react.com"}},
		Course{CourseName: "Nodejs", CourseId: "2", CoursePrice: 299, Author: &Author{Fullname: "Yash", Website: "node.com"}},
		Course{CourseName: "Nextjs", CourseId: "3", CoursePrice: 399, Author: &Author{Fullname: "Yash", Website: "nextjs.com"}},
	)
	// fmt.Printf("%#v", courses[0])
	r := mux.NewRouter()

	r.HandleFunc("/", ServeHome).Methods("GET")
	r.HandleFunc("/courses", GetAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", GetOneCourse).Methods("GET")
	r.HandleFunc("/course", CreateOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", UpdateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", DeleteOneCourse).Methods("DELETE")

	fmt.Println("Server running fine")
	log.Fatal(http.ListenAndServe(":4000", r))
}
