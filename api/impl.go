package api

import (
	"encoding/json"
	"net/http"
)

// ensure that we've conformed to the 'ServerInterface' with a compile-time check
var _ServerInterface = (*Server)(nil)

type Server struct{}

func NewServer() Server {
	return Server{}
}

// (GET /members)
func (Server) GetMembers(w http.ResponseWriter, r *http.Request) {
	members := Members{
		{
			Id:        "1",
			FirstName: "Wolfgang",
			LastName:  "Mattuttis",
		},
		{
			Id:        "2",
			FirstName: "Simone",
			LastName:  "Faber",
		},
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(members)
}

func (Server) GetMember(w http.ResponseWriter, r *http.Request, id int64) {
	member := Member{
		Id:        "1",
		FirstName: "Wolfgang",
		LastName:  "Mattuttis",
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(member)
}
