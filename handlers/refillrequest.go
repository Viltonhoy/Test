package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type js struct {
	Id  int
	Sum int
}

type njs struct {
	Id1 int
	Id2 int
	Sum int
}

func (u *UserOne) Ref(w http.ResponseWriter, r *http.Request) {
	var typ js

	str, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(str, &typ)

	if err != nil {
		fmt.Println("Error", err)
		http.Error(w, "Empty request body", http.StatusBadRequest)
		return
	}

	if u.Id == typ.Id {
		u.Account = u.Account + typ.Sum
	}

	js, err := json.Marshal(u.Account)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Write(js)

}

func (u *UserOne) Minus(w http.ResponseWriter, r *http.Request) {
	var typ js

	str, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(str, &typ)

	if err != nil {
		fmt.Println("Error", err)
		http.Error(w, "Empty request body", http.StatusBadRequest)
		return
	}

	if u.Id == typ.Id && u.Account >= typ.Sum {
		u.Account = u.Account - typ.Sum
	} else {
		return
	}

	js, err := json.Marshal(u.Account)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Write(js)

}

func (u *TwoUsers) Pere(w http.ResponseWriter, r *http.Request) {
	var typ njs

	str, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(str, &typ)

	if err != nil {
		fmt.Println("Error", err)
		http.Error(w, "Empty request body", http.StatusBadRequest)
		return
	}

	u.First.Account = u.First.Account - typ.Sum
	u.Second.Account = u.Second.Account + typ.Sum

	js, err := json.Marshal(u)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Write(js)
}
