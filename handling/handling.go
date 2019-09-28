package handling

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jcasado94/nats-points/invalidator"
	"github.com/jcasado94/nats-points/mongo/drivers"
)

type Handling struct {
	driver *drivers.MongoDriver
}

func NewHandling() (Handling, error) {
	md, err := drivers.NewMongoDriver()
	if err != nil {
		return Handling{}, err
	}
	return Handling{
		driver: &md,
	}, nil
}

func (h *Handling) HandleInvalidation(w http.ResponseWriter, r *http.Request) {
	countryName := r.URL.Query().Get("country")
	if countryName == "" {
		http.Error(w, "no country name", http.StatusBadRequest)
		return
	}
	invalidator, err := invalidator.NewInvalidator(h.driver)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = invalidator.InvalidateAllCountryArticles(countryName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	log.Printf("%s articles invalidated", countryName)
}

func (h *Handling) HandleTagArticles(w http.ResponseWriter, r *http.Request) {
	countryName := r.URL.Query().Get("country")
	if countryName == "" {
		http.Error(w, "no country name", http.StatusBadRequest)
		return
	}
	articles, err := h.driver.GetAllCountryResultArticlesTagged(countryName)
	if err != nil {
		http.Error(w, fmt.Sprintf("couldn't retreive all tagged articles for %s", countryName), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(articles)
}

func (h *Handling) HandleArticles(w http.ResponseWriter, r *http.Request) {
	countryName := r.URL.Query().Get("country")
	if countryName == "" {
		http.Error(w, "no country name", http.StatusBadRequest)
		return
	}
	articles, err := h.driver.GetAllCountryResultArticles(countryName)
	if err != nil {
		http.Error(w, fmt.Sprintf("couldn't retreive all articles for %s", countryName), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(articles)
}

func (h *Handling) HandleInformation(w http.ResponseWriter, r *http.Request) {
	countryName := r.URL.Query().Get("country")
	if countryName == "" {
		http.Error(w, "no country name", http.StatusBadRequest)
		return
	}
	info, err := h.driver.GetCountryInformation(countryName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(info)
}
