package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/* Model for Course */
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

/* We will use Slice as a database */
var courses []Course

/* Helper */
func isEmpty(C *Course) bool {
	return C.CourseName == ""
}

func main() {

}

/* Controllers */
/* Serve Home */
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Build API</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")
	/* Get the course id from the url */
	params := mux.Vars(r)              /* Similar to r.URL.Query() */
	courseId := params.Get("courseid") /* Get courseid */

	/* If course id matches, course will be returned */
	for _, course := range courses {
		if course.CourseId == courseId {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found for given ID")
}
