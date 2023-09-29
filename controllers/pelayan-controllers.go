package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	
	"github.com/gorilla/mux"
	"github.com/josepnapitupulu/API_Pelayan/models"
	"github.com/josepnapitupulu/API_Pelayan/utils"
)

var NewPelayan models.Pelayan
// var tmpl * template.Template
// func init(){
// 	tmpl = template.Must(template.ParseFiles("jemaat.html"))
// }

// func Index(w http.ResponseWriter, r *http.Request){
// 	temp, _ := template.ParseFiles("views/jemaat.html")
// 	temp.Execute(w, nil)
// }

// func Create(w http.ResponseWriter, r *http.Request){
// 	temp, _ := template.ParseFiles("views/createJemaat.html")
// 	temp.Execute(w, nil)
// }

func GetPelayan(w http.ResponseWriter, r *http.Request) {
	newPelayans := models.GetAllPelayans()
	res, _ := json.Marshal(newPelayans)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetPelayanById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pelayanId := vars["pelayanId"]
	NIK, err := strconv.ParseInt(pelayanId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	pelayanDetails, _ := models.GetPelayanbyId(NIK)
	res, _ := json.Marshal(pelayanDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreatePelayan(w http.ResponseWriter, r *http.Request) {
	CreatePelayan := &models.Pelayan{}
	utils.ParseBody(r, CreatePelayan)
	b := CreatePelayan.CreatePelayan()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeletePelayan(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pelayanId := vars["pelayanId"]
	ID, err := strconv.ParseInt(pelayanId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	pelayan := models.DeletePelayan(ID)
	res, _ := json.Marshal(pelayan)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdatePelayan(w http.ResponseWriter, r *http.Request) {
    var updatePelayan = &models.Pelayan{}
    utils.ParseBody(r, updatePelayan)
    vars := mux.Vars(r)
    pelayanId := vars["pelayanId"]
    ID, err := strconv.ParseInt(pelayanId, 0, 0)
    if err != nil {
        fmt.Println("error while parsing")
    }
    pelayanDetails, db := models.GetPelayanbyId(ID)
    if updatePelayan.NIK != "" {
        pelayanDetails.NIK = updatePelayan.NIK
    }
    if updatePelayan.NamaPelayan != "" {
        pelayanDetails.NamaPelayan = updatePelayan.NamaPelayan
    }
    if updatePelayan.Tahbisan != "" {
        pelayanDetails.Tahbisan = updatePelayan.Tahbisan
    }
    if updatePelayan.Pendidikan != "" {
        pelayanDetails.Pendidikan = updatePelayan.Pendidikan
    }
    if updatePelayan.BidangPendidikan != "" {
        pelayanDetails.BidangPendidikan = updatePelayan.BidangPendidikan
    }
	if updatePelayan.Pekerjaan != "" {
        pelayanDetails.Pekerjaan = updatePelayan.Pekerjaan
    }
	if updatePelayan.Alamat != "" {
        pelayanDetails.Alamat = updatePelayan.Alamat
    }
    db.Save(&pelayanDetails)
    res, _ := json.Marshal(pelayanDetails)
    w.Header().Set("Content-Type", "pkglication/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}