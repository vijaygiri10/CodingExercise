package lib

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//GetRoutes ...
func GetRoutes() http.Handler {
	// Creating Mux Router Object
	router := mux.NewRouter().StrictSlash(true)

	//Registering HTTP End Point with Mux Router
	router.HandleFunc("/", indexHandler).Methods(http.MethodGet)
	router.HandleFunc("/health_status", healthStatus).Methods(http.MethodGet)
	router.HandleFunc("/create/student", createStudent).Methods(http.MethodPost)
	router.HandleFunc("/delete/student/{id}", deleteStudent).Methods(http.MethodDelete)
	router.HandleFunc("/get/students", getAllStudents).Methods(http.MethodGet)
	router.HandleFunc("/get/students/{id}", getStudent).Methods(http.MethodGet)

	router.HandleFunc("/create/assignment", createAssignment).Methods(http.MethodPost)
	router.HandleFunc("/get/assignments", getAssignmentsList).Methods(http.MethodGet)
	router.HandleFunc("/create/student/score", createScoreStudent).Methods(http.MethodPost)
	router.HandleFunc("/update/student/score", updateScoreStudent).Methods(http.MethodPut)
	// router.HandleFunc("/open/{email_id}/{account_id}/{url_index}/{url_digest}", EmailOpen).Methods(http.MethodGet)

	allowedHeaders := []string{"X-Requested-With", "Content-Type", "Authorization"}
	allowedMethods := []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodHead, http.MethodOptions, http.MethodPatch}
	allowedOrigins := []string{"*"}

	return handlers.CORS(handlers.AllowedHeaders(allowedHeaders), handlers.AllowedMethods(allowedMethods), handlers.AllowedOrigins(allowedOrigins))(router)

}
