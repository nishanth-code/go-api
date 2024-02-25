package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

//fake db

var courses []Course

func (c *Course) isEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""


}

func main() {

}

//controller

func serveHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>hello welcome to  my frst go api</h1>"))
}

func getAllcourses(w http.ResponseWriter, r *http.Request){
	fmt.Println("get all courses")
	w.Header().Set("content-type","application/json")
	json.NewEncoder(w).Encode(courses)
}

func getCourse(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type","application/json")
	params := mux.Vars(r)

	for _,course := range courses{
		if course.CourseId == params["id"]{
			json.NewEncoder(w).Encode(course)
			return
		}

	}
	json.NewEncoder(w).Encode("invalid id")
	return

}

func CreateOneCourse(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type","application/json")
	if r.Body ==nil {
		json.NewEncoder(w).Encode("please send some data")

	}

	var course Course
	_=json.NewDecoder(r.Body).Decode(&course)


	if course.isEmpty(){
		json.NewEncoder(w).Encode("empty json")
		return


	}
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
} 
func updateCourse(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type","application/json")
	params := mux.Vars(r)
	var courseToUpdate Course
	_=json.NewDecoder(r.Body).Decode(&courseToUpdate)
	for index,course := range courses{
		if course.CourseId == params["id"]{
			courses = append(courses[:index],courses[index+1:]... )
			courseToUpdate.CourseId = params["id"]
			courses = append(courses, courseToUpdate)
			json.NewEncoder(w).Encode(courseToUpdate)
			return

		}
	}
	json.NewEncoder(w).Encode("invalid id")




}