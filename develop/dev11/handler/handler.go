package handler

import (
	"encoding/json"
	"errors"
	"level_2/develop/dev11/model"
	"level_2/develop/dev11/service"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() {
	mux := http.NewServeMux()
	mux.HandleFunc("/create_event", h.createEventHandler)
	mux.HandleFunc("/update_event", h.updateEventHandler)
	mux.HandleFunc("/delete_event", h.deleteEventHandler)
	mux.HandleFunc("/events_for_day", h.eventsForDayHandler)
	mux.HandleFunc("/events_for_week", h.eventsForWeekHandler)
	mux.HandleFunc("/events_for_month", h.eventsForMonthHandler)
	http.Handle("/", loggingMiddleware(mux))
}

func (h *Handler) createEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	event, err := parseForm(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := h.service.EventService.CreateEvent(event)
	writeJSON(w, map[string]interface{}{"result": "event created", "id": id})
}

func (h *Handler) updateEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	event, err := parseForm(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.EventService.UpdateEvent(id, event); err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	writeJSON(w, map[string]string{"result": "event updated"})
}

func (h *Handler) deleteEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	if err := h.service.EventService.DeleteEvent(id); err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	writeJSON(w, map[string]string{"result": "event deleted"})
}

func (h *Handler) eventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	h.eventsForPeriodHandler(1, w, r)
}

func (h *Handler) eventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	h.eventsForPeriodHandler(7, w, r)
}

func (h *Handler) eventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	h.eventsForPeriodHandler(30, w, r)
}

func (h *Handler) eventsForPeriodHandler(days int, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		http.Error(w, "invalid user_id", http.StatusBadRequest)
		return
	}
	date, err := time.Parse("2006-01-02", r.URL.Query().Get("date"))
	if err != nil {
		http.Error(w, "invalid date format", http.StatusBadRequest)
		return
	}
	startDate := date
	endDate := date.AddDate(0, 0, days)
	events := h.service.EventService.GetEvents(userID, startDate, endDate)
	writeJSON(w, events)

}

func writeJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		log.Println("encoding error")
	}
}

func parseForm(r *http.Request) (model.Event, error) {
	err := r.ParseForm()
	if err != nil {
		return model.Event{}, err
	}
	userID, err := strconv.Atoi(r.FormValue("user_id"))
	if err != nil {
		return model.Event{}, errors.New("invalid user_id")
	}
	date, err := time.Parse("2006-01-02", r.FormValue("date"))
	if err != nil {
		return model.Event{}, errors.New("invalid date format")
	}
	title := r.FormValue("title")
	if title == "" {
		return model.Event{}, errors.New("title is required")
	}
	return model.Event{
		UserID: userID,
		Date:   date,
		Title:  title,
	}, nil
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		duration := time.Since(startTime)
		log.Printf("Completed %s %s in %v", r.Method, r.URL.Path, duration)
	})
}
