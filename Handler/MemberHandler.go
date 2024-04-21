package Handler

import (
	"SD_2024_Project/Controller"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func UpdateMemberHandler(w http.ResponseWriter, r *http.Request) {
	// Ambil nilai parameter ID dari URL
	r.ParseForm()
	id := r.Form.Get("id")
	idInt, _ := strconv.Atoi(id)
	user := Controller.MemberSearch(idInt)

	if r.Method == "GET" {
		// Menampilkan form inputan
		tmpl := template.Must(template.ParseFiles(
			"View/MemberUpdate.html"))
		//fmt.Println("tmpl ", tmpl)
		if err := tmpl.Execute(w, user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if r.Method == "POST" {
		// Handle form submission
		r.ParseForm()
		firstName := r.Form.Get("namaDepan")
		lastName := r.Form.Get("namaBelakang")
		noTelp := r.Form.Get("noTelp")

		// Memanggil controller untuk insert data
		Controller.MemberUpdate(idInt, firstName, lastName, noTelp)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

}
func DeleteMemberHandler(w http.ResponseWriter, r *http.Request) {
	// Ambil nilai parameter ID dari URL
	r.ParseForm()
	id := r.Form.Get("id")
	fmt.Println("ID yang dihapus : ", id)
	//conversion string to int
	idInt, _ := strconv.Atoi(id)

	Controller.MemberDelete(idInt)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	//memanggil memberView.html dengan tamplate

	tmpl := template.Must(template.ParseFiles(
		"View/MemberView.html"))
	users := Controller.MembersView()
	// Menampilkan data ke template HTML
	if err := tmpl.Execute(w, users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func InsertHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Menampilkan form inputan
		tmpl := template.Must(template.ParseFiles(
			"C:\\Users\\kurniawan\\GolandProjects\\SD_2024_Project\\View\\MemberInsert.html"))
		//fmt.Println("tmpl ", tmpl)
		if err := tmpl.Execute(w, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if r.Method == "POST" {
		// Handle form submission
		r.ParseForm()
		firstName := r.Form.Get("firstName")
		lastName := r.Form.Get("lastName")
		noTelp := r.Form.Get("phone")

		// Memanggil controller untuk insert data
		Controller.MemberInsert(firstName, lastName, noTelp)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
}
