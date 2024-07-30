package skatepark_api

import (
	"encoding/json"
	"net/http"
<<<<<<< HEAD:controller/health.go

	"github.com/derekpedersen/skatepark-api-go/service"
	log "github.com/sirupsen/logrus"
=======
>>>>>>> master:src/health.go
)

const (
	// HealthyMessage indicates the service is healthy
	HealthyMessage = "Johnny 5 Healthy!"

	// AliveMessage indicates the service is alive
	AliveMessage = "Johnny 5 Alive!"

	// ReadyMessage indicates the service is ready
	ReadyMessage = "Johnny 5 Ready!"
)

// swagger:model healthMessage
type HealthMessage struct {
	Message string `json:"message"`
}

// NewHealthMessage creates a new HealthMessage struct
func NewHealthMessage(str string) *HealthMessage {
	return &HealthMessage{
		Message: str,
	}
}

// HealthService interface
type HealthService interface {
	GetHealthyMessage() *HealthMessage
	GetAliveMessage() *HealthMessage
	GetReadyMessage() *HealthMessage
}

// HealthServiceImpl struct
type HealthServiceImpl struct {
}

// NewHealthService creates a new HealthService
func NewHealthService() HealthService {
	return &HealthServiceImpl{}
}

// GetHealthyMessage gets the alive message
func (a *HealthServiceImpl) GetHealthyMessage() *HealthMessage {
	return NewHealthMessage(HealthyMessage)
}

// GetAliveMessage gets the alive message
func (a *HealthServiceImpl) GetAliveMessage() *HealthMessage {
	return NewHealthMessage(AliveMessage)
}

// GetReadyMessage gets the ready message
func (a *HealthServiceImpl) GetReadyMessage() *HealthMessage {
	return NewHealthMessage(ReadyMessage)
}

// HealthAPIController interface
type HealthAPIController interface {
	GetAliveMessage(w http.ResponseWriter, r *http.Request)
	GetReadyMessage(w http.ResponseWriter, r *http.Request)
	GetHealthyMessage(w http.ResponseWriter, r *http.Request)
}

// HealthAPIControllerImpl struct
type HealthAPIControllerImpl struct {
	svc           service.HealthService
	skateparksSvc service.SkateparksService
}

// NewHealthAPIController creates a new HealthAPIController
func NewHealthAPIController(svc service.HealthService, skateparksSvc service.SkateparksService) HealthAPIController {

	return &HealthAPIControllerImpl{
		svc:           svc,
		skateparksSvc: skateparksSvc,
	}
}


// GetAliveMessage controller
func (api *HealthAPIControllerImpl) GetAliveMessage(w http.ResponseWriter, r *http.Request) {
	alive := api.svc.GetAliveMessage()

	_, err := api.skateparksSvc.GetSkateparks()
	if err != nil {
		log.Errorf("Error in getting skateparks:\n %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(*alive)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// GetReadyMessage controller
func (api *HealthAPIControllerImpl) GetReadyMessage(w http.ResponseWriter, r *http.Request) {
	alive := api.svc.GetReadyMessage()

	_, err := api.skateparksSvc.GetSkateparks()
	if err != nil {
		log.Errorf("Error in getting skateparks:\n %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(*alive)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// GetHealthyMessage controller
func (api *HealthAPIControllerImpl) GetHealthyMessage(w http.ResponseWriter, r *http.Request) {
	alive := api.svc.GetHealthyMessage()

	_, err := api.skateparksSvc.GetSkateparks()
	if err != nil {
		log.Errorf("Error in getting skateparks:\n %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(*alive)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
