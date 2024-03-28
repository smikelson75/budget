package handlers

import (
	"budgetstoragelib/dto/requests"
	"budgetstoragelib/interfaces"
	"budgetstoragelib/mssql/repositories"
	"encoding/json"
	"fmt"
	"net/http"
)

func ReturnCategoryHandlers(storage interfaces.IServer) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		handlers := map[string]http.HandlerFunc{
			http.MethodGet:  getCategoryList(storage),
			http.MethodPost: getCategoryPost(storage),
			http.MethodDelete: getCategoryDelete(storage),
		}

		if handler, ok := handlers[r.Method]; ok {
			handler(w, r)
			return
		}

		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
	}
}

func getCategoryList(storage interfaces.IServer) func(w http.ResponseWriter, r *http.Request) {
	repo := repositories.NewCategoryRepository(storage)
	return func(w http.ResponseWriter, r *http.Request) {
		list, err := repo.ListCategories()
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		err = json.NewEncoder(w).Encode(Result{ Categories: list})
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("an internal server error occurred"))
		}
	}
}

func getCategoryPost(storage interfaces.IServer) func(w http.ResponseWriter, r *http.Request) {
	repo := repositories.NewCategoryRepository(storage)
	return func(w http.ResponseWriter, r *http.Request) {
		var request requests.CategoryRequest

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid request"))
			return
		}

		result, err := repo.Manage(request)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("an internal server error occurred"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		err = json.NewEncoder(w).Encode(&Result{ Categories: result})
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("an internal server error occurred"))
		}
	}
}

func getCategoryDelete(storage interfaces.IServer) func(w http.ResponseWriter, r *http.Request) {
	repo := repositories.NewCategoryRepository(storage)
	return func(w http.ResponseWriter, r *http.Request) {
		var request requests.CategoryRequest

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid request"))
			return
		}

		var deleteRequest requests.CategoryRequest = requests.CategoryRequest{
			User: request.User,
		}

		for _, category := range request.Categories {
			deleteRequest.Categories = append(deleteRequest.Categories, requests.CategoryDto{
				Id: category.Id,
				Remove: true,
			})
		}

		_, err = repo.Manage(request)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("an internal server error occurred"))
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
