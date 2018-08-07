package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"github.com/amstee/Go-Cake-Factory/models"
	"strconv"
	"sort"
	"strings"
)

// SORT THE CAKE LIST BY RANK AND ALPHABETICAL
func SortCakes(i, j int) bool {
	if models.Cakes[i].Rank == models.Cakes[j].Rank {
		max := len(models.Cakes[i].Title)
		if max > len(models.Cakes[j].Title) {
			max = len(models.Cakes[j].Title)
		}
		s1 := strings.ToLower(models.Cakes[i].Title)
		s2 := strings.ToLower(models.Cakes[j].Title)

		for idx := 0; idx < max; idx++ {
			if s1[idx] != s2[idx] {
				return s1[idx] < s2[idx]
			}
		}
		return len(s1) < len(s2)
	}
	return models.Cakes[i].Rank > models.Cakes[j].Rank
}

func GetCakes(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	cakes, err := json.Marshal(models.Cakes); if err != nil {
		w.WriteHeader(500)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(cakes)
	}
}

func GetCake(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	i, err := strconv.Atoi(ps.ByName("id")); if err != nil || i < 0 || i >= len(models.Cakes) {
		w.WriteHeader(400)
		return
	}
	cakes, err := json.Marshal(models.Cakes[i]); if err != nil {
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(cakes)
}

func AddCake(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)

	var cake models.Cake

	err := decoder.Decode(&cake); if err != nil {
		w.WriteHeader(400)
	}
	models.Cakes = append(models.Cakes, cake)
	sort.Slice(models.Cakes, SortCakes)
	w.WriteHeader(http.StatusOK)
}

func DeleteCake(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	i, err := strconv.Atoi(ps.ByName("id")); if err != nil || i < 0 || i >= len(models.Cakes) {
		w.WriteHeader(400)
		return
	}
	models.Cakes = append(models.Cakes[:i], models.Cakes[i+1:]...)
	w.WriteHeader(http.StatusOK)
}

func ResetCakes(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	models.Cakes = nil
	jsonFile, err := os.Open("config/cakes.json"); if err != nil {
		fmt.Println("Cake file not found")
		os.Exit(1)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &models.Cakes)
	jsonFile.Close()
	sort.Slice(models.Cakes, SortCakes)
	w.WriteHeader(http.StatusOK)
}

func main() {
	router := httprouter.New()
	router.GET("/cakes", GetCakes)
	router.GET("/cakes/:id", GetCake)
	router.POST("/cakes", AddCake)
	router.DELETE("/cakes/:id", DeleteCake)
	router.PUT("/cakes/init", ResetCakes)

	jsonFile, err := os.Open("config/cakes.json"); if err != nil {
		fmt.Println("Cake file not found")
		os.Exit(1)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &models.Cakes)
	jsonFile.Close()
	sort.Slice(models.Cakes, SortCakes)
	log.Fatal(http.ListenAndServe(":5000", router))
}