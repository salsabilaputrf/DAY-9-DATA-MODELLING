package main

import (
	"context"
	"data-modelling/connection"
	"fmt"
	"strconv"
	"time"

	// "log"
	"net/http"
	//"strconv"
	"text/template"
	//"time"

	"github.com/gorilla/mux"
)

var Data = map[string]interface{}{
	"Title":   "Personal Web",
	"IsLogin": true,
}


type dataProject struct {
	Id int
	ProjectName	string
	StartDate time.Time
	EndDate	time.Time
	Description	string
	Technologies []string
	Duration string
}

var Projects = []dataProject{
	
}

func main(){
	route := mux.NewRouter()

	connection.DatabaseConnect()

	// static folder
	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	// routing
	route.HandleFunc("/", home).Methods("GET").Name("home")
	// route.HandleFunc("/contact", contactMe).Methods("GET")
	// route.HandleFunc("/addProject", addProject).Methods("GET")
	// route.HandleFunc("/addProject", addProjectInput).Methods("POST")
	// route.HandleFunc("/detailProject/{id}", detailProject).Methods("GET")
	// route.HandleFunc("/deleteProject/{id}", deleteProject).Methods("GET")
	// route.HandleFunc("/editProject/{id}", editProject).Methods("GET")
	// route.HandleFunc("/editProjectInput/{id}", editProjectInput).Methods("POST")

	// port := 5000
	fmt.Println("Server is running on port 5000")
	http.ListenAndServe("localhost:5000", route)
}

func home(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var tmpl, err = template.ParseFiles("view/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}
	var result []dataProject

	rows, err := connection.Conn.Query(context.Background(), "SELECT * FROM tb_project ")
	if err != nil {
		fmt.Println(err.Error())
		return 
	}
	
	for rows.Next() {
		var each = dataProject{}

		var err = rows.Scan(&each.Id, &each.ProjectName, &each.StartDate, &each.EndDate, &each.Description, &each.Technologies)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		each.Duration = selisih(each.StartDate, each.EndDate)
		result = append(result, each)
	}

	respData := map[string]interface{}{
		"Data":  Data,
        "Projects": result,
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, respData)
}

func selisih(start time.Time, end time.Time)string{

	distance := end.Sub(start)

	// Menghitung durasi
	var duration string
	year := int(distance.Hours()/(12 * 30 * 24))
	if year != 0 {
		duration = strconv.Itoa(year) + " tahun"
	}else{
		month := int(distance.Hours()/(30 * 24))
		if month != 0 {
			duration = strconv.Itoa(month) + " bulan"
		}else{
			week := int(distance.Hours()/(7 * 24))
			if week != 0 {
				duration = strconv.Itoa(week) + " minggu"
			}else{
				day := int(distance.Hours()/(24))
				if day != 0 {
					duration = strconv.Itoa(day) + " hari"
				}
			}
		}
	}

	return duration
}

// func contactMe(w http.ResponseWriter, r *http.Request){
// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")

// 	var tmpl, err = template.ParseFiles("view/contact-form.html")
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte("message : " + err.Error()))
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	tmpl.Execute(w, Data)
// }

// func addProject(w http.ResponseWriter, r *http.Request){
// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")

// 	var tmpl, err = template.ParseFiles("view/addProject.html")
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte("message : " + err.Error()))
// 		return
// 	}

// 	respData := map[string]interface{} {
// 		"Data":  Data,
//         "Projects": Projects,
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	tmpl.Execute(w, respData)
// }

// func addProjectInput(w http.ResponseWriter, r *http.Request){
// 	err := r.ParseForm()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
	
// 	projectName := r.PostForm.Get("projectName")
// 	startDate := r.PostForm.Get("startDate")
// 	endDate := r.PostForm.Get("endDate")
// 	desc := r.PostForm.Get("desc")
// 	tech := r.Form["technologi"]

// 	// Menghitung durasi
// 	// Parsing string to time

// 	// Start Date
// 	startDateTime, _ := time.Parse("2006-01-02", startDate)

// 	// End Date
// 	endDateTime, _ := time.Parse("2006-01-02", endDate)
	
// 	// Perbedaan nya berupa : jam menit detik
// 	distance := endDateTime.Sub(startDateTime)

// 	// Menghitung durasi
// 	var duration string
// 	year := int(distance.Hours()/(12 * 30 * 24))
// 	if year != 0 {
// 		duration = strconv.Itoa(year) + " tahun"
// 	}else{
// 		month := int(distance.Hours()/(30 * 24))
// 		if month != 0 {
// 			duration = strconv.Itoa(month) + " bulan"
// 		}else{
// 			week := int(distance.Hours()/(7 * 24))
// 			if week != 0 {
// 				duration = strconv.Itoa(week) + " minggu"
// 			}else{
// 				day := int(distance.Hours()/(24))
// 				if day != 0 {
// 					duration = strconv.Itoa(day) + " hari"
// 				}
// 			}
// 		}
// 	}

// 	var newProject = dataProject {
// 		ProjectName: projectName,
// 		StartDate: startDate,
// 		EndDate: endDate,
// 		Description: desc,
// 		Technologies: tech,
// 		Duration: duration,
// 	}
	
// 	Projects = append(Projects, newProject)


// 	http.Redirect(w, r, "/", http.StatusMovedPermanently)
// }

// func detailProject(w http.ResponseWriter, r *http.Request){
// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")

// 	var tmpl, err = template.ParseFiles("view/detail-project.html")
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte("message : " + err.Error()))
// 		return
// 	}

// 	id, _ := strconv.Atoi(mux.Vars(r)["id"])

// 	ProjectDetail := dataProject{}

// 	for index, data := range Projects{
// 		if index == id {
// 			newStartDate, _ := time.Parse("2006-01-02", data.StartDate)
// 			newEndDate, _ := time.Parse("2006-01-02", data.EndDate)

// 			ProjectDetail = dataProject{
// 				Id: id,
// 				ProjectName: data.ProjectName,
// 				StartDate: newStartDate.Format("02 Jan 2006"),
// 				EndDate: newEndDate.Format("02 Jan 2006"),
// 				Description: data.Description,
// 				Technologies: data.Technologies,
// 				Duration: data.Duration,
// 		}
// 		}
// 	}


// 	respDataDetail := map[string]interface{}{
// 		"Data": Data,
// 		"ProjectDetail": ProjectDetail,
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	tmpl.Execute(w, respDataDetail)
// }

// func deleteProject(w http.ResponseWriter, r *http.Request){
// 	id, _ := strconv.Atoi(mux.Vars(r)["id"])

// 	Projects = append(Projects[:id], Projects[id+1:]...)
	
// 	http.Redirect(w, r, "/", http.StatusFound)
// }

// func editProject(w http.ResponseWriter, r *http.Request){
// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")

// 	var tmpl, err = template.ParseFiles("view/editProject.html")
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte("message : " + err.Error()))
// 		return
// 	}

// 	id, _ := strconv.Atoi(mux.Vars(r)["id"])

// 	ProjectDetail := dataProject{}

// 	for index, data := range Projects{
// 		if index == id {
// 			ProjectDetail = dataProject{
// 				Id: id,
// 				ProjectName: data.ProjectName,
// 				StartDate: data.StartDate,
// 				EndDate: data.EndDate,
// 				Description: data.Description,
// 				Technologies: data.Technologies,
// 				Duration: data.Duration,
// 		}
// 		}
// 	}

// 	respData := map[string]interface{} {
// 		"Data":  Data,
//         "ProjectDetail": ProjectDetail,
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	tmpl.Execute(w, respData)
// }

// func editProjectInput(w http.ResponseWriter, r *http.Request){
// 	err := r.ParseForm()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
	
// 	projectName := r.PostForm.Get("projectName")
// 	startDate := r.PostForm.Get("startDate")
// 	endDate := r.PostForm.Get("endDate")
// 	desc := r.PostForm.Get("desc")
// 	tech := r.Form["technologi"]

// 	// Menghitung durasi
// 	// Parsing string to time

// 	// Start Date
// 	startDateTime, _ := time.Parse("2006-01-02", startDate)

// 	// End Date
// 	endDateTime, _ := time.Parse("2006-01-02", endDate)
	
// 	// Perbedaan nya berupa : jam menit detik
// 	distance := endDateTime.Sub(startDateTime) 

// 	// Menghitung durasi
// 	var duration string
// 	year := int(distance.Hours()/(12 * 30 * 24))
// 	if year != 0 {
// 		duration = strconv.Itoa(year) + " tahun"
// 	}else{
// 		month := int(distance.Hours()/(30 * 24))
// 		if month != 0 {
// 			duration = strconv.Itoa(month) + " bulan"
// 		}else{
// 			week := int(distance.Hours()/(7 * 24))
// 			if week != 0 {
// 				duration = strconv.Itoa(week) + " minggu"
// 			}else{
// 				day := int(distance.Hours()/(24))
// 				if day != 0 {
// 					duration = strconv.Itoa(day) + " hari"
// 				}
// 			}
// 		}
// 	}

// 	var newProject = dataProject {
// 		ProjectName: projectName,
// 		StartDate: startDate,
// 		EndDate: endDate,
// 		Description: desc,
// 		Technologies: tech,
// 		Duration: duration,
// 	}

// 	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	
// 	Projects[id] = newProject

// 	http.Redirect(w, r, "/", http.StatusMovedPermanently)
// }

