package post

import (
	"final/structs"
	"final/userbase"
	"fmt"
	"net/http"
	"os"
)


// Post Groups - name of the group should be unique
func PostGroups(w http.ResponseWriter, r *http.Request){

	// Connect to database
	db, err := userbase.Connect()
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	defer db.Close()

	// Take group name from HTML page
	r.ParseForm()

	newGroup := r.Form["groups"][0]


	// Check if the name is empty
	if newGroup == ""{
		fmt.Fprintf(w, "The name of the group is empty")
		return
	}

	// Check if group in Groups table
	check, _ := userbase.ExistGroupName(db, newGroup)
	if check == "true"{
		fmt.Fprintf(w, "The group '" +newGroup+"' already exist")
		return
	}

	newID := userbase.Newid()
	userbase.WriteGroups(db, newGroup, newID)
	row := userbase.ReadGroup(db, newGroup)

	var gr structs.Group
	err = row.Scan(&gr.Id, &gr.Name, &gr.Uuid)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "%+v\n", gr)
}

// Post Tasks - name of the task should be unique
func PostTasks(w http.ResponseWriter, r *http.Request){

	// Connect to database
	db, err := userbase.Connect()
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	defer db.Close()

	// Take the task name and group name to task
	r.ParseForm()

	newTask := r.Form["task_id"][0]
	taskGroup := r.Form["group_id"][0]

	// Check if the Task name is empty
	if newTask == ""{
		fmt.Fprintf(w, "The name of the task is empty")
		return
	}

	// Check if the Group name is empty
	if taskGroup == ""{
		fmt.Fprintf(w, "The name of the group is empty")
		return
	}

	// Check if group in Groups table
	check, _ := userbase.ExistGroupName(db, taskGroup)
	if check == "false"{
		fmt.Fprintf(w, "The group '" +taskGroup+"' does not exist")
		return
	}

	// Check if task in Tasks table
	check, _ = userbase.ExistTaskName(db, newTask)
	if check == "true"{
		fmt.Fprintf(w, "The task '" +newTask+"' already exist")
		return
	}

	// We write task to database
	newID := userbase.Newid()
	userbase.WriteTasks(db, newTask, newID)

	// We read group name to get group uuid
	row := userbase.ReadGroup(db, taskGroup)

	var gr structs.Group
	err = row.Scan(&gr.Id, &gr.Name, &gr.Uuid)
	if err != nil {
		panic(err)
	}

	// We write the connection between group uuid and task uuid
	userbase.WriteGroupsTasks(db, gr.Uuid, newID)

	var tk structs.Task
	row = userbase.ReadTask(db, newTask)
	err = row.Scan(&tk.Id, &tk.Name, &tk.Uuid)

	tk.Group = taskGroup

	fmt.Fprintf(w, "%+v\n", tk)
}

// Post Timeframes: timeframes should be in the specified format, name of the task should exist
func PostTimeframes(w http.ResponseWriter, r *http.Request){

	// Connect to database
	db, err := userbase.Connect()
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	defer db.Close()

	// Take task name, start time, stop time
	r.ParseForm()

	timeTask := r.Form["task_name"][0]
	timeStart := r.Form["start_time"][0]
	timeStop := r.Form["time_stop"][0]

	// Check if the fields are ok
	if timeTask == "" || timeStart=="" || timeStop==""{
		fmt.Fprintf(w, "Check the fields because they are empty")
		return
	}

	// Check if task in Tasks table
	check, _ := userbase.ExistTaskName(db, timeTask)
	if check == "false"{
		fmt.Fprintf(w, "The task '" +timeTask+"' does not exist")
		return
	}

	timeID := userbase.Newid()

	// We read task to get task uuid
	var tk structs.Task
	row := userbase.ReadTask(db, timeTask)
	err = row.Scan(&tk.Id, &tk.Name, &tk.Uuid)
	taskID := tk.Uuid

	// We write start time stop time to Timeframes, we link task to timeframe in TaskToTimeframes
	userbase.WriteTimeframes(db, timeStart, timeStop, timeID)
	userbase.WriteTasksTimeframes(db, taskID, timeID)

	// We read new timeframe from database
	var tf structs.Timeframe
	row = userbase.ReadTimeframe(db, timeID)
	err = row.Scan(&tf.Id, &tf.Start, &tf.Stop, &tf.Uuid)

	tf.Task = timeTask

	fmt.Fprintf(w, "%+v\n", tf)
}