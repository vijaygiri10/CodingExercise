package lib

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"

	"CodingExercise/shared/log"

	"github.com/gorilla/mux"
)

//WriteResponse ...
func WriteResponse(write http.ResponseWriter, response interface{}, statusCode int) {
	write.Header().Set("Content-type", "application/json")
	write.WriteHeader(statusCode)
	json.NewEncoder(write).Encode(response)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "indexHandler")
}

func healthStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "healthStatus")
}

//createStudent ...
func createStudent(w http.ResponseWriter, r *http.Request) {
	var msg Students
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		log.Error(r.Context(), "createStudent json NewDecoder err: ", err)
		WriteResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := msg.Insert(r.Context()); err != nil {
		WriteResponse(w, "unable to store student record: "+err.Error(), http.StatusInternalServerError)
		return
	}

	WriteResponse(w, "succesfully created record", http.StatusAccepted)
}

func deleteStudent(w http.ResponseWriter, r *http.Request) {
	var id student
	if id = student(strings.TrimSpace(mux.Vars(r)["id"])); id == "" {
		WriteResponse(w, "id is required", http.StatusBadRequest)
		return
	}
	err := id.deleteStudent(r.Context())
	if err != nil {
		WriteResponse(w, "unable to delete record: "+err.Error(), http.StatusInternalServerError)
		return
	}
	WriteResponse(w, "succesfully delete record ", http.StatusOK)
}

func getStudent(w http.ResponseWriter, r *http.Request) {
	var id student
	if id = student(strings.TrimSpace(mux.Vars(r)["id"])); id == "" {
		WriteResponse(w, "id is required", http.StatusBadRequest)
		return
	}

	std, err := id.getStudent(r.Context())
	if err != nil {
		WriteResponse(w, "unable to fetch record: "+err.Error(), http.StatusInternalServerError)
		return
	}

	WriteResponse(w, std, http.StatusOK)
}

func getAllStudents(w http.ResponseWriter, r *http.Request) {
	assignments, err := fetchAllStudent(r.Context())
	if err != nil {
		WriteResponse(w, "unable to fetch Students record: ", http.StatusOK)
		return
	}

	sort.SliceStable(assignments, func(i, j int) bool {
		return assignments[i].UpdatedAT.Before(assignments[j].UpdatedAT)
	})

	WriteResponse(w, assignments, http.StatusOK)
}

func createAssignment(w http.ResponseWriter, r *http.Request) {
	var msg Assignments
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		log.Error(r.Context(), "createAssignment json NewDecoder err: ", err)
		WriteResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := msg.Insert(r.Context()); err != nil {
		log.Error(r.Context(), "createAssignment json NewDecoder err: ", err)
		WriteResponse(w, "Unable to Store Assignments Record: "+err.Error(), http.StatusInternalServerError)
		return
	}
	WriteResponse(w, "successfully created Assignments", http.StatusOK)
}

func createScoreStudent(w http.ResponseWriter, r *http.Request) {
	var msg StudentAssignmentsScore
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		log.Error(r.Context(), "scoreStudent json NewDecoder err: ", err)
		WriteResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := msg.Insert(r.Context()); err != nil {
		log.Error(r.Context(), "scoreStudent json NewDecoder err: ", err)
		WriteResponse(w, "Unable to Store Assignments Record: "+err.Error(), http.StatusInternalServerError)
		return
	}
	WriteResponse(w, "successfully created Assignments", http.StatusOK)
}

func getAssignmentsList(w http.ResponseWriter, r *http.Request) {
	assignments, err := fetchAssignmentsList(r.Context())
	if err != nil {
		WriteResponse(w, "unable to Students record: ", http.StatusOK)
		return
	}
	WriteResponse(w, assignments, http.StatusOK)
}

func updateScoreStudent(w http.ResponseWriter, r *http.Request) {
	var msg StudentAssignmentsScore
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		log.Error(r.Context(), "updateScoreStudent json NewDecoder err: ", err)
		WriteResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := msg.Update(r.Context()); err != nil {
		log.Error(r.Context(), "updateScoreStudent json NewDecoder err: ", err)
		WriteResponse(w, "unable to update Student Score: "+err.Error(), http.StatusInternalServerError)
		return
	}
	WriteResponse(w, "Sucessfully update Student Score", http.StatusOK)
}
