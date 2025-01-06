package handlers

import (
	"encoding/json"
	"net/http"
	"quickcache/cache"
)

func ShowCache(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cacheValues := cache.CacheData.GetValues()
	if cacheValues == nil {
		http.Error(w, `{"error": "Cache is empty"}`, http.StatusNotFound)
		return
	}
	response := struct {
		Message string   `json:"message"`
		Values  []string `json:"values"`
	}{
		Message: "Cache values retrieved successfully",
		Values:  cacheValues,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func AddCache(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		http.Error(w, `{"error": "Invalid method. Use POST."}`, http.StatusMethodNotAllowed)
		return
	}
	var realData requestData
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&realData)
	if err != nil {
		http.Error(w, `{"error": "Invalid JSON format"}`, http.StatusBadRequest)
		return
	}
	cache.CacheData.Set(realData.Value)
	response := struct {
		Message string `json:"message"`
		Value   string `json:"value"`
	}{
		Message: "Cache added successfully",
		Value:   realData.Value,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
	cache.CacheData.Traverse()

}

type requestData struct {
	Value string `json:"value"`
}
