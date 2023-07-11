package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"login/controller"
)

func main() {
	// Called HandleFunction (to run program)
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
		}
	}()
	HandleFunction()

}

func HandleFunction() {
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	router := mux.NewRouter().StrictSlash(true)
	router = router.PathPrefix("/api/v1/contactApp").Subrouter()

	router.HandleFunc("/login", controller.Login).Methods("POST")

	//userRouter := router.PathPrefix("/user").Subrouter()
	// Made routes (users)
	// userRouter.HandleFunc("/", controller.CreateUser).Methods("POST")
	// userRouter.HandleFunc("/{id}", controller.GetUserById).Methods("GET")
	// userRouter.HandleFunc("/", controller.GetAllUsers).Methods("GET")
	// userRouter.HandleFunc("/{id}", controller.UpdateUserById).Methods("PUT")
	// userRouter.HandleFunc("/{id}", controller.DeleteUserById).Methods("DELETE")

	// contactRouter := userRouter.PathPrefix("/{userid}/contacts").Subrouter()

	// contactRouter.HandleFunc("/", controller.CreateContact).Methods("POST")
	// contactRouter.HandleFunc("/{id}", controller.GetContactById).Methods("GET")
	// contactRouter.HandleFunc("/", controller.GetAllContacts).Methods("GET")
	// contactRouter.HandleFunc("/{id}", controller.UpdateContact).Methods("PUT")
	// contactRouter.HandleFunc("/{id}", controller.DeleteContact).Methods("DELETE")

	// contactInfoRouter := userRouter.PathPrefix("/{userid}/contactinfo").Subrouter()

	// contactInfoRouter.HandleFunc("/", controller.CreateContactInfo).Methods("POST")
	// contactInfoRouter.HandleFunc("/{id}", controller.GetContactInfoById).Methods("GET")
	// contactInfoRouter.HandleFunc("/", controller.GetAllContactInfo).Methods("GET")
	// contactInfoRouter.HandleFunc("/{id}", controller.UpdateContactInfo).Methods("PUT")
	// contactInfoRouter.HandleFunc("/{id}", controller.DeleteContactInfo).Methods("DELETE")

	log.Printf("Server Live on localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(originsOk, headersOk, methodsOk)(router)))

}
