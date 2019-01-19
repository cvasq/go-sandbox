// A usage example of the gorilla/mux http URL router

package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs/xid"
)

const listeningPort = "8080"

type person struct {
	id        string
	firstName string
	lastName  string
	age       int
}

var allPeople = []person{
	{id: "bh0kkf47uh3blslo4t80", firstName: "Homer", lastName: "Simpson", age: 36},
	{id: "bh0kkj47uh3bkgjuukqg", firstName: "Marge", lastName: "Simpson", age: 34},
	{id: "bh0kkjc7uh3bj7jstv7g", firstName: "Bart", lastName: "Simpson", age: 10},
	{id: "bh0kkn47uh3bien15vsg", firstName: "Lisa", lastName: "Simpson", age: 8},
	{id: "bh0kknk7uh3bhj2k4r1g", firstName: "Maggie", lastName: "Simpson", age: 1},
}

// Lists each person in allPeople slice
func listPeople(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	for _, p := range allPeople {
		fmt.Fprintf(w, "Active User:\n\tID: %v\n\tName: %v %v\n\tAge: %v\n",
			p.id,
			p.firstName,
			p.lastName,
			p.age)
	}
}

// Adds a new person to the allPeople slice
func addPerson(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)

	// Convert age to int
	ageInt, _ := strconv.Atoi(mux.Vars(r)["age"])

	// Generate unique ID for new user
	guid := xid.New()

	allPeople = append(allPeople, person{
		id:        guid.String(),
		firstName: mux.Vars(r)["first"],
		lastName:  mux.Vars(r)["last"],
		age:       ageInt,
	})
	fmt.Fprintf(w, "Added %v to the user database!\n", mux.Vars(r)["first"])
}

// Retrieves a user's name by ID
func getUserByID(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)

	idRequested := mux.Vars(r)["id"]

	for _, v := range allPeople {
		if v.id == idRequested {
			fmt.Fprintf(w, "User found:\n\tName: %v %v\n\tID: %v\n", v.firstName, v.lastName, v.id)
			return
		}
	}
	fmt.Fprintf(w, "User id %v not found\n", idRequested)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	message := `Endpoint Usage Examples:

// List all users
curl http://localhost:8080/list-users

// Get user name by ID
curl http://localhost:8080/get-user/id/bh0kkf47uh3blslo4t80 

// Add new user
curl -X POST "http://localhost:8080/add-user/?first=Bill&last=Nye&age=35"
`
	fmt.Fprintln(w, message)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", rootHandler)

	r.HandleFunc("/list-users", listPeople)

	r.HandleFunc("/get-user/id/{id}", getUserByID)

	// Matches URL with query values that uses POST method
	r.HandleFunc("/add-user/", addPerson).Methods("POST").Queries(
		"first", "{first}",
		"last", "{last}",
		"age", "{age}",
	)

	log.Println("Listening for HTTP connections on port:", listeningPort)
	err := http.ListenAndServe(":"+listeningPort, r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
