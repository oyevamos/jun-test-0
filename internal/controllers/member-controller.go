package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/oyevamos/jun-test-0.git/internal/models"
	"github.com/oyevamos/jun-test-0.git/internal/utils"
	"net/http"
	"strconv"
)

var NewMember models.Member

func CreateMember(w http.ResponseWriter, r *http.Request) {
	CreateMember := &models.Member{}
	utils.ParseBody(r, CreateMember)
	m := CreateMember.CreateMember()
	res, _ := json.Marshal(m)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetAllMembers(w http.ResponseWriter, r *http.Request) {
	newMembers := models.GetAllMembers()
	res, _ := json.Marshal(newMembers)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetMemberById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	memberId := vars["memberId"]
	ID, err := strconv.ParseInt(memberId, 0, 0)
	if err != nil {
		fmt.Println("error while Parsing")
	}
	memberDetails, _ := models.GetMemberById(ID)
	res, _ := json.Marshal(memberDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteMember(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	memberId := vars["memberId"]
	ID, err := strconv.ParseInt(memberId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	member := models.DeleteMember(ID)
	res, _ := json.Marshal(member)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateMember(w http.ResponseWriter, r *http.Request) {
	var updateMember = &models.Member{}
	utils.ParseBody(r, updateMember)
	vars := mux.Vars(r)
	memberId := vars["memberId"]
	ID, err := strconv.ParseInt(memberId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	memberDetails, db := models.GetMemberById(ID)
	if updateMember.Name != "" {
		memberDetails.Name = updateMember.Name
	}
	if updateMember.Email != "" {
		memberDetails.Email = updateMember.Email
	}
	if updateMember.Country != "" {
		memberDetails.Country = updateMember.Country
	}
	if updateMember.Color != "" {
		memberDetails.Color = updateMember.Color
	}
	db.Save(&memberDetails)
	res, _ := json.Marshal(memberDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
