package controllers

import (
	"fmt"
	internalHttp "github.com/wmrodrigues/go-pokedex/internal/http"
	"net/http"
)

// HomeIndexHandler
func HomeIndexHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	var items []map[string]interface{}

	items = make([]map[string]interface{}, 0)
	for i := 1; i < 151; i++ {
		item := make(map[string]interface{})
		item["Id"] = i
		item["Name"] = fmt.Sprintf("Pokémon %v", i)
		items = append(items, item)
	}

	result := map[string]interface{}{
		"Items": items,
	}

	err = internalHttp.Render("/web/templates/home/index.html", w, result)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
