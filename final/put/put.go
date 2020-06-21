package put

import (
	"final/structs"
	"final/userbase"
	"fmt"
	"net/http"
	"os"
)

// This function puts task to database
func PutTask(w http.ResponseWriter, r *http.Request){

	// Connect to database
	db, err := userbase.Connect()
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	defer db.Close()

	// We take all data from form input form
	r.ParseForm()

	taskNameOld := r.Form["put_tasks"][0]
	taskNameNew := r.Form["task_name_new"][0]
	groupNameNew := r.Form["task_group_new"][0]
	taskUUIDNew := r.Form["task_uuid_new"][0]

	if taskNameOld == ""{
		fmt.Fprintf(w, "The name of the task is empty")
		return
	}

	// Check if task in Tasks table
	check, _ := userbase.ExistTaskName(db, taskNameOld)
	if check == "false"{
		fmt.Fprintf(w, "The task '" +taskNameOld+"' does not exist")
		return
	}

	// Fill in data before change
	var groupUUIDOld string

	taskUUIDOld := userbase.ReadTaskID(db, taskNameOld)
	row := userbase.ReadGroupTask(db, taskUUIDOld)

	err = row.Scan(&groupUUIDOld)
	if err != nil {
		panic(err)
	}

	groupNameOld := userbase.ReadGroupName(db, groupUUIDOld)

	// We check new name field: if blank - we do not change it
	if taskNameNew == ""{
		taskNameNew = taskNameOld
	}else{
		// Check if task in Tasks table
		check, _ = userbase.ExistTaskName(db, taskNameNew)
		if check == "true"{
			fmt.Fprintf(w, "The task '" +taskNameNew+"' already exist")
			return
		}
	}

	// Check group field: if blank - we do not change it
	if groupNameNew == ""{

		groupNameNew = groupNameOld

	}else{
		// Check if group in Groups table
		check, _ := userbase.ExistGroupName(db, groupNameNew)
		if check == "false"{
			fmt.Fprintf(w, "The group '" +groupNameNew+"' does not exist")
			return
		}
	}

	var groupUUIDNew string
	groupUUIDNew = userbase.ReadGroupID(db, groupNameNew)

	// Check UUID field: if blank - we do not change it
	if taskUUIDNew == ""{

		taskUUIDNew = taskUUIDOld

	}else{
		// Check if task UUID in task table
		check, _ := userbase.ExistTaskUUID(db, taskUUIDNew)
		if check == "true"{
			fmt.Fprintf(w, "The task with UUID: '" +taskUUIDNew+"' already exist")
			return
		}
	}

	// We Put data to database
	err = userbase.UpdateTaskTableName(db, taskNameOld, taskNameNew)
	if err != nil {
		panic(err)
	}

	err = userbase.UpdateTaskTableUUID(db, taskUUIDOld, taskUUIDNew)
	if err != nil {
		panic(err)
	}

	err = userbase.UpdateGTTRelationGroup(db, taskUUIDOld, groupUUIDNew)
	if err != nil {
		panic(err)
	}

	err = userbase.UpdateGTTTaskUUID(db, taskUUIDOld, taskUUIDNew)
	if err != nil {
		panic(err)
	}

	err = userbase.UpdateGTTTaskUUID(db, taskUUIDOld, taskUUIDNew)
	if err != nil {
		panic(err)
	}

	err = userbase.UpdateTTTTaskUUID(db, taskUUIDOld, taskUUIDNew)
	if err != nil {
		panic(err)
	}

	// We read data about task
	row = userbase.ReadTask(db, taskNameNew)
	var tk structs.Task
	var gt structs.GroupTask

	// We read Tasks table to task struct
	err = row.Scan(&tk.Id, &tk.Name, &tk.Uuid)
	if err != nil {
		panic(err)
	}

	// We read GroupToTask table to groupTask struct
	row = userbase.ReadGroupTask(db, tk.Uuid)
	err = row.Scan(&gt.Group_id)
	if err != nil {
		panic(err)
	}

	tk.Group = userbase.ReadGroupName(db, gt.Group_id)

	// We read all tasks for group
	rows, _ := userbase.ReadTaskTimeframes(db, tk.Uuid)

	for rows.Next() {

		var timeframe_id string
		var task_start string
		var task_stop string
		var ti structs.Interval

		err = rows.Scan(&timeframe_id)
		if err != nil {
			panic(err)
		}

		task_start, task_stop = userbase.ReadTimeframeStartStop(db, timeframe_id)

		ti.Start = append(ti.Start, task_start)
		ti.Stop = append(ti.Stop, task_stop)

		tk.Timeframes = append(tk.Timeframes, ti)
	}
	fmt.Fprintf(w, "%+v\n", tk)
}

// This function puts group to database
func PutGroup(w http.ResponseWriter, r *http.Request){

	// Connect to database
	db, err := userbase.Connect()
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	defer db.Close()

	// We take all data from form input form
	r.ParseForm()

	groupNameOld := r.Form["put_groups"][0]
	groupNameNew := r.Form["group_name_new"][0]
	groupUUIDNew := r.Form["group_uuid_new"][0]

	if groupNameOld == ""{
		fmt.Fprintf(w, "The name of the group is empty")
		return
	}

	// Check if group in Groups table
	check, _ := userbase.ExistGroupName(db, groupNameOld)
	if check == "false"{
		fmt.Fprintf(w, "The group '" +groupNameOld+"' does not exist")
		return
	}

	// Read an old group UUID
	var groupUUIDOld string
	groupUUIDOld = userbase.ReadGroupID(db, groupNameOld)

	// We check new name field: if blank - we do not change it
	if groupNameNew == ""{
		groupNameNew = groupNameOld
	}else{
		// Check if group in Groups table
		check, _ = userbase.ExistGroupName(db, groupNameNew)
		if check == "true"{
			fmt.Fprintf(w, "The group '" +groupNameNew+"' already exist")
			return
		}
	}

	// We check new UUID field: if blank - we do not change it
	if groupUUIDNew == ""{
		groupUUIDNew = groupUUIDOld
	}else{
		// Check if group UUID in Groups table
		check, _ = userbase.ExistGroupUUID(db, groupUUIDNew)
		if check == "true"{
			fmt.Fprintf(w, "The group uuid: '"+groupUUIDNew+"' already exist")
		}
	}

	// We put data to database
	err = userbase.UpdateGroupTableName(db, groupNameOld, groupNameNew)
	if err != nil {
		panic(err)
	}

	err = userbase.UpdateGroupTableUUID(db, groupUUIDOld, groupUUIDNew)
	if err != nil {
		panic(err)
	}

	err =userbase.UpdateGTTGroupUUID(db, groupUUIDOld, groupUUIDNew)
	if err != nil {
		panic(err)
	}

	// We read data about group
	row := userbase.ReadGroup(db, groupNameNew)
	var gr structs.Group

	err = row.Scan(&gr.Id, &gr.Name, &gr.Uuid)
	if err != nil {
		panic(err)
	}

	// We read all tasks for group
	rows, _ := userbase.ReadGroupTasks(db, gr.Uuid)

	for rows.Next(){

		var task_id string
		var task_name string

		err = rows.Scan(&task_id)
		if err != nil {
			panic(err)
		}

		task_name = userbase.ReadTaskName(db, task_id)

		gr.Tasks = append(gr.Tasks, task_name)
	}

	fmt.Fprintf(w, "%+v\n", gr)
}
