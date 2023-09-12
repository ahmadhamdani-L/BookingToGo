package nationality

import (
	"bookingtogo/internal/dto"
	"bookingtogo/internal/factory"
	"encoding/json"
	"net/http"
	
)

type handler struct {
	service *service
}
func NewHandler(f *factory.Factory) *handler {
	service := NewService(f)
	return &handler{service}
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
  
	payload := new(dto.NationalityGetRequest)
    
    result, err := h.service.Find(w, r, payload)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
	
    responseData, err := json.Marshal(result)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    _, err = w.Write(responseData)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
  
	payload := new(dto.NationalityCreateRequest)

    err := json.NewDecoder(r.Body).Decode(payload)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    result, err := h.service.Create(w, r, payload)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
	
    responseData, err := json.Marshal(result)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    _, err = w.Write(responseData)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

