package api

import (
	"encoding/json"
	"insider-assesment/pkg/app"
	"insider-assesment/pkg/domain/message"
	"log"
	"net/http"
	"sync"
)

const address = ":4300"

type Server struct {
	*app.ServiceProvider
	*app.OptionProvider
	http.Handler
	mux           *http.ServeMux
	Worker        *app.Worker
	workerMu      sync.Mutex
	workerRunning bool
}

func NewServer(
	serviceProvider *app.ServiceProvider,
	optionProvider *app.OptionProvider,
	worker *app.Worker,
) *Server {
	return &Server{
		ServiceProvider: serviceProvider,
		OptionProvider:  optionProvider,
		mux:             http.NewServeMux(),
		Worker:          worker,
	}
}

func (s *Server) Run() {
	s.mux.HandleFunc("/worker", s.handleWorkerControl)
	s.mux.HandleFunc("/messages", s.handleMessages)

	log.Println("Starting worker...")
	go s.Worker.Start()
	s.workerRunning = true

	log.Println("Listening on", address)
	if err := http.ListenAndServe(address, s.mux); err != nil {
		panic(err)
	}
}

func (s *Server) handleWorkerControl(w http.ResponseWriter, r *http.Request) {
	s.workerMu.Lock()
	defer s.workerMu.Unlock()

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if s.workerRunning {
		s.Worker.Stop()
		s.workerRunning = false
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Worker stopped successfully"))
	} else {
		go s.Worker.Start()
		s.workerRunning = true
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Worker started successfully"))
	}
}

func (s *Server) handleMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := r.Context()

	messages, err := s.PostgresService.GetMessagesByStatus(ctx, message.Sent, -1)
	if err != nil {
		log.Println("Error fetching messages:", err)

		http.Error(w, "Failed to fetch messages", http.StatusInternalServerError)
		return
	}

	if len(messages) == 0 {
		http.Error(w, "No messages found", http.StatusNotFound)
		return
	}

	response := newMessageListResponse(messages)

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(response); err != nil {
		log.Println("Error encoding messages:", err)
		http.Error(w, "Failed to encode messages", http.StatusInternalServerError)
	}
}
