package main

import (
	"fmt"
	"log"
	"encoding/json"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)

type Student struct{
	ID string `json:"id"`
	Usn string `json:"usn"`
	Name string `json:"name"`
	Department *Department `json:"department"`
}

type Department struct{
	Dept_id string `json:"dept_id"`
	Dept_name string `json:"dept_name"`
}

var students []Student

func getStudents(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

func deleteStudent(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range students {
		if item.ID == params["id"]{
			students = append(students[:index], students[index+1:]...) //Main logic of the Delete the student
			break
		}
	}
}

func getStudentsByID(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _,item := range students {
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return 
		}
	}
}

func createStudent(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var student Student
	_ = json.NewDecoder(r.Body).Decode(&student)
	student.ID = strconv.Itoa(rand.Intn(100000))
	students = append(students, student)
	json.NewEncoder(w).Encode(student)
}


func updateStudent(w http.ResponseWriter, r *http.Request){
	//which type we have define first
	w.Header().Set("Content-Type", "application/json")
	// then we have to take the params id where that will be in the "http://localhost/8000/students/1234/"
	params := mux.Vars(r)
	// go through the loop in the students 
	for index, item := range students{
		// check for the params id 
		if item.ID == params["id"]{
			students = append(students[:index], students[index+1:]...)				// here , if id is matched then we have to first delete the matched id,
			var student Student														// and then the new data is added to the same id 
			_ = json.NewDecoder(r.Body).Decode(&student)
			student.ID = params["id"]
			students = append(students, student)
			json.NewEncoder(w).Encode(student)
			return
		}
	}

}

func main(){
	r := mux.NewRouter()

	students = append(students, Student{ID: "1", Usn: "23cd055", Name: "Tejas", Department: &Department{Dept_id: "ds01", Dept_name: "Data Science"}})
	students = append(students, Student{ID: "2", Usn: "23cd020", Name: "Lord", Department: &Department{Dept_id: "ds01", Dept_name: "Data Science"}})


	r.HandleFunc("/students/", getStudents).Methods("GET")
	r.HandleFunc("/students/{id}/", getStudentsByID).Methods("GET")
	r.HandleFunc("/students/", createStudent).Methods("POST")
	r.HandleFunc("/students/{id}/", deleteStudent).Methods("DELETE")
	r.HandleFunc("/students/{id}/", updateStudent).Methods("PUT")


	fmt.Printf("Starting the server at Port 8000\n")
	log.Fatal(http.ListenAndServe(":8000",r))
}

