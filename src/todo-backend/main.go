package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// TodoSvc mock service implementation
var TodoSvc *MockTodoService

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	origins, found := os.LookupEnv("ALLOWED_ORIGINS")
	if !found {
		origins = "*"
	}

	TodoSvc = NewMockTodoService()
	mux := http.NewServeMux()

	mux.Handle("/todos", commonHandlers(todoHandler, origins))
	mux.Handle("/todos/", commonHandlers(todoHandler, origins))

	todoURL := "http://"
	hostname, _ := os.Hostname()
	if "todomvc.go" != hostname {
		hostname = GetOutboundIP().String()
	}
	todoURL += hostname + ":" + port + "/todos/"
	log.Printf("Server is ready to handle requests at %q\n", todoURL)
	log.Printf("Allowed origins %s\n", origins)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

// GetOutboundIP will get preferred outbound IP of this machine/container.
// Using the UDP protocol, it does not have handshake nor a connection.
// The target does not need be there and you will receive the outbound IP
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "1.1.1.1:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func addURLToTodos(r *http.Request, todos ...*Todo) {
	scheme := "https"
	if r.TLS != nil {
		scheme = "https"
	}
	baseURL := scheme + "://" + r.Host + "/todos/"
	for _, todo := range todos {
		todo.URL = baseURL + strconv.Itoa(todo.ID)
	}
}

func todoHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	key := ""
	if len(parts) > 2 {
		key = parts[2]
	}

	switch r.Method {
	case "GET":
		if len(key) == 0 {
			todos, err := TodoSvc.GetAll()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			addURLToTodos(r, todos...)
			json.NewEncoder(w).Encode(todos)
		} else {
			id, err := strconv.Atoi(key)
			if err != nil {
				http.Error(w, "Invalid Id", http.StatusBadRequest)
				return
			}
			todo, err := TodoSvc.Get(id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if todo == nil {
				http.NotFound(w, r)
				return
			}
			addURLToTodos(r, todo)
			json.NewEncoder(w).Encode(todo)
		}
	case "POST":
		if len(key) > 0 {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		todo := Todo{
			Completed: false,
		}
		err := json.NewDecoder(r.Body).Decode(&todo)
		if err != nil {
			http.Error(w, err.Error(), 422)
			return
		}
		err = TodoSvc.Save(&todo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		addURLToTodos(r, &todo)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(todo)
	case "PATCH":
		id, err := strconv.Atoi(key)
		if err != nil {
			http.Error(w, "Invalid Id", http.StatusBadRequest)
			return
		}
		var todo Todo
		err = json.NewDecoder(r.Body).Decode(&todo)
		if err != nil {
			http.Error(w, err.Error(), 422)
			return
		}
		todo.ID = id

		err = TodoSvc.Save(&todo)
		if err != nil {
			if strings.ToLower(err.Error()) == "not found" {
				http.NotFound(w, r)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		addURLToTodos(r, &todo)
		json.NewEncoder(w).Encode(todo)
	case "DELETE":
		if len(key) == 0 {
			TodoSvc.DeleteAll()
		} else {
			id, err := strconv.Atoi(key)
			if err != nil {
				http.Error(w, "Invalid Id", http.StatusBadRequest)
				return
			}
			err = TodoSvc.Delete(id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		w.WriteHeader(http.StatusNoContent)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
}
