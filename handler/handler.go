package handler

import (
	//"fmt"
	"html/template"
	"log"
	"net/http"

	//"strings"
	"github.com/JohnDirewolf/capstone/maze"
	"github.com/JohnDirewolf/capstone/shared/types"
)

// HANDLERS
func Root(w http.ResponseWriter, r *http.Request) {
	pageData := struct {
		Title string
	}{
		Title: "Root of Capstone Project!",
	}

	pageTemplate, err := template.ParseFiles("templates/shared/base.html", "templates/shared/header.html", "templates/index.html")
	if err != nil {
		log.Printf("Handler, Root, Error accessing HTML file: %v", err)
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500: Unable to find page."))
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	err = pageTemplate.Execute(w, pageData)
	if err != nil {
		//Too late to do any real error handling, just log the error.
		log.Printf("Handler, Root, Error executing page: %v", err)
	}
}

func Game(w http.ResponseWriter, r *http.Request) {
	//This takes the action encoded in r and updates and displays the new updated page.
	//To prevent caching, add the following header to the response "Cache-Control: no-store or Cache-Control: no-cache, no-store, must-revalidate"

	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	//Get the action path parameter from the request.
	action := r.URL.Query().Get("action")
	switch types.UrlAction(action) {
	case types.Start:
		start(w)
	case types.End:
		end(w)
	case types.North, types.South, types.West, types.East:
		move(action, w)
	case types.Get:
		get(w)
	default:
		//Unknown action
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusBadRequest)
		var pageBody string = "400: Unknown action."
		w.Write([]byte(pageBody))
	}
}

// ACTIONS
func start(w http.ResponseWriter) {
	//Initialize the Maze
	maze.Init()
	generateRoom(w, true)
}

func end(w http.ResponseWriter) {
	startPageData := types.PageData{
		Title: "Maze Runner End!",
	}

	pageTemplate, err := template.ParseFiles("templates/shared/base.html", "templates/shared/header.html", "templates/end.html")
	if err != nil {
		log.Printf("Handler, Game, end, Error accessing HTML file: %v", err)
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500: Unable to find page."))
		return
	}

	err = pageTemplate.Execute(w, startPageData)
	if err != nil {
		//Too late to do any real error handling, just log the error.
		log.Printf("Handler, Game, end, Error executing page: %v", err)
	}
}

func move(direction string, w http.ResponseWriter) {
	maze.Move(direction)
	generateRoom(w, false)
}

func get(w http.ResponseWriter) {
	maze.GetItems()
	generateRoom(w, false)
}

func generateRoom(w http.ResponseWriter, isStart bool) {
	var pageInfo types.PageData
	pageTemplate, err := template.ParseFiles("templates/shared/base.html", "templates/shared/header.html", "templates/maze.html")
	if err != nil {
		log.Printf("Handler, Game, move, Error accessing HTML file: %v", err)
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500: Unable to find page."))
		return
	}

	pageInfo = maze.GetPageInfo()
	//Start has a few extra things that need to be set.
	if isStart {
		pageInfo.Description = template.HTML("The entrace slams shut behind you. You will have to look for a different exit!<br />") + pageInfo.Description
		pageInfo.Instructions = maze.GetInstructions()
	}

	err = pageTemplate.Execute(w, pageInfo)
	if err != nil {
		//Too late to do any real error handling, just log the error.
		log.Printf("Handler, Game, move, Error executing page: %v", err)
	}
}
