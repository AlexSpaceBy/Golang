package  main

import (
	"final/delete"
	"final/get"
	"final/post"
	"final/put"
	"net/http"
)

// For initial page
func index(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "new.html")
}

func main(){

	// Index page handler
	http.HandleFunc("/", index)

	// POST handlers
	http.HandleFunc("/group_handler", post.PostGroups)
	http.HandleFunc("/task_handler", post.PostTasks)
	http.HandleFunc("/timeframe_handler", post.PostTimeframes)

	// GET handlers
	http.HandleFunc("/group_getter", get.GetGroups)
	http.HandleFunc("/task_getter", get.GetTasks)
	http.HandleFunc("/timeframe_getter", get.GetTimeframes)

	// DELETE handlers
	http.HandleFunc("/timeframe_deleter", delete.DeteleTimeframe)
	http.HandleFunc("/task_deleter", delete.DeleteTask)
	http.HandleFunc("/group_deleter", delete.DeleteGroup)

	// PUT handler
	http.HandleFunc("/task_putter", put.PutTask)
	http.HandleFunc("/group_putter", put.PutGroup)

	http.ListenAndServe(":8080", nil)

}