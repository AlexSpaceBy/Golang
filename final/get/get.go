package get

import (
	"database/sql"
	"final/structs"
	"final/userbase"
	"fmt"
	"net/http"
	"os"
)

// Get Groups: if text field is blank - we get all groups, if text field has name of the group - we get this group
// Done according to the task
func GetGroups(w http.ResponseWriter, r *http.Request){

	// We get name of the group from html page
	r.ParseForm()
	groupName := r.Form["get_groups"][0]

	// Connect to database
	db, err := userbase.Connect()
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	defer db.Close()

	if groupName == "" {

		// Read Groups from database
		rows, _ := userbase.ReadGroups(db)

		for rows.Next() {

			var gr structs.Group

			err = rows.Scan(&gr.Id, &gr.Name, &gr.Uuid)
			if err != nil {
				panic(err)
			}

			// We read all tasks for group
			rows, _ := userbase.ReadGroupTasks(db, gr.Uuid)

			for rows.Next() {

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
	}else{

		// Check if group in Groups table
		check, _ := userbase.ExistGroupName(db, groupName)
		if check == "false"{
			fmt.Fprintf(w, "The group '" +groupName+"' does not exist")
			return
		}

		// We read data about group
		row := userbase.ReadGroup(db, groupName)
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
}

// Get Tasks: if text field is blank - we get all tasks, if text field has name of the task - we get this task
// Done according to the task 
func GetTasks(w http.ResponseWriter, r *http.Request){

	// Connect to database
	db, err := userbase.Connect()
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	defer db.Close()

	// We read task name
	r.ParseForm()
	taskName := r.Form["get_tasks"][0]

	if taskName == "" {
		// Read Groups from database
		rows, _ := userbase.ReadTasks(db)

		var row *sql.Row

		for rows.Next() {

			var tk structs.Task
			var gt structs.GroupTask

			// We read Tasks table to task struct
			err = rows.Scan(&tk.Id, &tk.Name, &tk.Uuid)
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
	}else{

		// Check if task in Tasks table
		check, _ := userbase.ExistTaskName(db, taskName)
		if check == "false"{
			fmt.Fprintf(w, "The task '" +taskName+"' does not exist")
			return
		}

		// We read data about task
		row := userbase.ReadTask(db, taskName)
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
}

// Get Timeframes: we do it for convenience, since uuid is a 16 digit number
// Not in task
func GetTimeframes(w http.ResponseWriter, r *http.Request){

	// Connect to database
	db, err := userbase.Connect()
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	defer db.Close()

	// Read Timeframes from database
	rows, _ := userbase.ReadTimeframes(db)
	str := ""

	for rows.Next() {

		var tf structs.Timeframe

		err = rows.Scan(&tf.Id, &tf.Start, &tf.Stop, &tf.Uuid)
		if err != nil {
			panic(err)
		}

		str = str + tf.Id + " " + tf.Start + " " + tf.Stop + " " + tf.Uuid + "\n"
	}
	fmt.Fprintf(w, str)
}

