package transport

import (
	"encoding/json"
	"github.com/omotolani98/goland-task-api/internal/todo"
	"log"
	"net/http"
)

type TodoItem struct {
	//ID   int    `json:"id"`
	Item string `json:"item"`
}

type Server struct {
	mux *http.ServeMux
}

func NewServer(todoService *todo.Service) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /todo", func(writer http.ResponseWriter, request *http.Request) {
		todos, err := todoService.GetAll()
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		todoResp, err := json.Marshal(todos)
		if err != nil {
			return
		}
		_, _ = writer.Write(todoResp)
		writer.WriteHeader(http.StatusOK)
		return
	})

	mux.HandleFunc("POST /todo", func(writer http.ResponseWriter, request *http.Request) {
		var t TodoItem
		err := json.NewDecoder(request.Body).Decode(&t)
		if err != nil {
			log.Print(err)
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		err = todoService.Add(t.Item)
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

		result, err := todoService.Search(query)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		body, err := json.Marshal(result)

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, err = writer.Write(body)
		writer.WriteHeader(http.StatusOK)

		if err != nil {
			log.Println(err)
			return
		}
	})

	if err := http.ListenAndServe(":8070", mux); err != nil {
		log.Fatal(err)
	}

	return &Server{mux: mux}
}

func (s *Server) ServeHTTP() error {
	return http.ListenAndServe(":8070", s.mux)
}
