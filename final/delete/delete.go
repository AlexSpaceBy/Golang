package delete

import (
	"final/userbase"
	"fmt"
	"net/http"
	"os"
)

// Delete Timeframe from database
func DeteleTimeframe(w http.ResponseWriter, r *http.Request){

	// Connect to database
	db, err := userbase.Connect()
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	defer db.Close()

	// We read timeframe uuid
	r.ParseForm()

	timeID := r.Form["del_timeframes"][0]

	// Check if the Timeframe uuid is empty
	if timeID == ""{
		fmt.Fprintf(w, "The uuid of the Timeframe is empty")
		return
	}

	// Check if timeframe in Timeframes table
	check, _ := userbase.ExistTimeframeUUID(db, timeID)
	if check == "false"{
		fmt.Fprintf(w, "The timeframe '" +timeID+"' does not exist")
		return
	}

	err = userbase.DeleteTimeframe(db, timeID)
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}

	err = userbase.DeleteTasksToTimeframes(db, timeID)
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Fprintf(w, "Timeframe uuid: "+timeID+" deleted")

}

// Delete Task from database
func DeleteTask(w http.ResponseWriter, r *http.Request){

	// Connect to database
	db, err := userbase.Connect()
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	defer db.Close()

	// We read which task to delete
	r.ParseForm()
	taskName := r.Form["del_tasks"][0]

	// Check if the Task name is empty
	if taskName == ""{
		fmt.Fprintf(w, "The name of the task is empty")
		return
	}

	// Check if task in Tasks table
	check, _ := userbase.ExistTaskName(db, taskName)
	if check == "false"{
		fmt.Fprintf(w, "The task '" +taskName+"' does not exist")
		return
	}

	taskID := userbase.ReadTaskID(db, taskName)

	// Delete task from Tasks table
	err = userbase.DeleteTask(db, taskID)
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}

	// Delete task from GroupsToTasks table
	err = userbase.DeleteGroupsToTasks(db, taskID)
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}

	// Delete all timeframes from Timeframes table for specified task
	rows, _ := userbase.ReadTaskTimeframes(db, taskID)
	for rows.Next(){

		var timeframeID string

		err = rows.Scan(&timeframeID)
		if err != nil {
			panic(err)
		}

		err = userbase.DeleteTimeframe(db, timeframeID)
		if err != nil {
			panic(err)
		}
	}

	// Delete task from tasksToTimeframes table
	err = userbase.DeleteTasksTimeframes(db, taskID)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "Task "+taskName+" uuid: "+taskID+" deleted")
}

// Delete group from database
func DeleteGroup(w http.ResponseWriter, r *http.Request){

	// Connect to database
	db, err := userbase.Connect()
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	defer db.Close()

	// We read which group to delete
	r.ParseForm()
	groupName := r.Form["del_groups"][0]

	// Check if the name is empty
	if groupName == ""{
		fmt.Fprintf(w, "The name of the group is empty")
		return
	}

	// Check if group in Groups table
	check, _ := userbase.ExistGroupName(db, groupName)
	if check == "false"{
		fmt.Fprintf(w, "The group '" +groupName+"' does not exist")
		return
	}

	groupID := userbase.ReadGroupID(db, groupName)

	// We delete group from Groups table
	err = userbase.DeleteGroup(db, groupID)

	// We read all uuid of each task in the group
	rows, _ := userbase.ReadGroupTasks(db, groupID)
	for rows.Next(){

		var taskID string

		err = rows.Scan(&taskID)
		if err != nil {
			panic(err)
		}

		// Delete task from Tasks table
		err = userbase.DeleteTask(db, taskID)
		if err != nil{
			fmt.Println(err)
			os.Exit(1)
		}

		// Delete all timeframes from Timeframes table for specified task
		rowsTask, _ := userbase.ReadTaskTimeframes(db, taskID)
		for rowsTask.Next(){

			var timeframeID string

			err = rowsTask.Scan(&timeframeID)
			if err != nil {
				panic(err)
			}

			err = userbase.DeleteTimeframe(db, timeframeID)
			if err != nil {
				panic(err)
			}
		}

		// Delete task from tasksToTimeframes table
		err = userbase.DeleteTasksTimeframes(db, taskID)
		if err != nil {
			panic(err)
		}
	}

	// Delete group from groupsToTasks
	err = userbase.DeleteGroupsTasks(db, groupID)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "Group "+groupName+" uuid: "+groupID+" deleted")

}