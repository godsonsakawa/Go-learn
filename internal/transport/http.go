package transport

import (
	"encoding/json"
	"log"
	"my-first-api/internal/todo"
	"net/http"
)

// TodoItem represents a single todo item with its description. It is used for JSON data exchange in the API.
type TodoItem struct {
	Item string `json:"item"`
}

// Server represents an HTTP server using ServeMux for request routing.
type Server struct {
	mux *http.ServeMux
}

// NewServer initializes and returns a new Server instance configured with the provided todo.Service for handling requests.
func NewServer(todoSvc *todo.Service) *Server {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /todo", func(w http.ResponseWriter, r *http.Request) {
		//  _, err := w.Write([]byte("Hello World")) We are taking the ResponseWriter which is part of the HandleFunc, we use the writer to write "Hello World" back to the user.

		todoItems, err := todoSvc.GetAll()
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		b, err := json.Marshal(todoItems)
		if err != nil {
			log.Println(err)
		}
		_, err = w.Write(b)
		if err != nil {
			log.Println(err)
		}
	})

	mux.HandleFunc("POST /todo", func(writer http.ResponseWriter, request *http.Request) {
		var t TodoItem
		err := json.NewDecoder(request.Body).Decode(&t)
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		err = todoSvc.Add(t.Item)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		writer.WriteHeader(http.StatusCreated)
		return
	})

	mux.HandleFunc("GET /search", func(writer http.ResponseWriter, request *http.Request) {
		query := request.URL.Query().Get("q")
		if query == "" {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		results, err := todoSvc.Search(query)
		if err != nil {
			log.Println(err.Error())
			writer.WriteHeader(http.StatusInternalServerError)
		}
		// Convert the search results into JSON and write them to the response
		jsonResults, err := json.Marshal(results)
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, err = writer.Write(jsonResults)
		if err != nil {
			log.Println(err)
			return
		}

	})

	return &Server{
		mux: mux,
	}
}

// Add a serve func

// Serve starts the HTTP server on port 8080 using the configured ServeMux for routing and handles incoming requests.
func (s *Server) Serve() error {
	return http.ListenAndServe(":8080", s.mux)
}
