package userbase

import (
	"database/sql"
	"fmt"

	//We generate uuid
	"github.com/google/uuid"

	//We connect to postgres database
	_ "github.com/lib/pq"
)

//This is for database connect, should be changed in case of database transfer to new PC
const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "12345"
	dbname = "final"
)

// Generate unique uuid for database
func Newid()(id string){
	id = uuid.New().String()
	return id
}

//We connect to database
func Connect()(db *sql.DB, err error) {
	connect := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = sql.Open("postgres", connect)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db, err
}

// ========================= Write block ===========================
//We write data to Tasks table
func WriteTasks(db *sql.DB, name string, uuid string)(err error){
	sqlWrite := "INSERT INTO Tasks (name, uuid) VALUES ($1, $2)"

	_, err = db.Exec(sqlWrite, name, uuid)
	if err != nil {
		panic(err)
	}
	return err
}

//We write data to Group table
func WriteGroups(db *sql.DB, name string, uuid string)(err error){
	sqlWrite := "INSERT INTO Groups (name, uuid) VALUES ($1, $2)"

	_, err = db.Exec(sqlWrite, name, uuid)
	if err != nil {
		panic(err)
	}
	return err
}

//We write data to Timeframe table
func WriteTimeframes(db *sql.DB, start string, stop string, uuid string)(err error){
	sqlWrite := "INSERT INTO Timeframes (start, stop, uuid) VALUES ($1, $2, $3)"

	_, err = db.Exec(sqlWrite, start, stop, uuid)
	if err != nil {
		panic(err)
	}
	return err
}

//We connect groups to tasks through relation table
func WriteGroupsTasks(db *sql.DB, uuid_group string, uuid_task string)(err error){
	sqlWrite := "INSERT INTO GroupsToTasks (group_id, task_id) VALUES ($1, $2)"

	_, err = db.Exec(sqlWrite, uuid_group, uuid_task)
	if err != nil {
		panic(err)
	}
	return err
}

//We connect tasks to timeframes through relation table
func WriteTasksTimeframes(db *sql.DB, uuid_task string, uuid_timeframe string)(err error){
	sqlWrite := "INSERT INTO TasksToTimeframes (task_id, timeframe_id) VALUES ($1, $2)"

	_, err = db.Exec(sqlWrite, uuid_task, uuid_timeframe)
	if err != nil {
		panic(err)
	}
	return err
}

// ========================= Read block ========================
//We read all Groups in table
func ReadGroups(db *sql.DB)(rows *sql.Rows,err error){
	sqlRead := "SELECT * FROM Groups"

	rows, err = db.Query(sqlRead)
	if err != nil {
		panic(err)
	}

	return rows, err
}

// We read all data for specified group by group name
func ReadGroup(db *sql.DB, name string)(row *sql.Row){
	sqlRead := "SELECT id, name, uuid FROM groups WHERE name='"+name+"'"

	row = db.QueryRow(sqlRead)

	return row
}

// We read group name for selected uuid
func ReadGroupName(db *sql.DB, uuid string)(name string){
	sqlRead := "SELECT name FROM groups WHERE uuid='"+uuid+"'"

	row := db.QueryRow(sqlRead)
	err := row.Scan(&name)
	if err != nil {
		panic(err)
	}

	return name
}

// We read group uuid by using name
func ReadGroupID(db *sql.DB, name string)(uuid string){

	sqlRead := "SELECT uuid FROM Groups WHERE name='"+name+"'"

	row := db.QueryRow(sqlRead)
	err := row.Scan(&uuid)
	if err != nil {
		panic(err)
	}

	return uuid
}

// We read all tasks in the task table
func ReadTasks(db *sql.DB)(rows *sql.Rows, err error){
	sqlRead :="SELECT * FROM Tasks"

	rows, err = db.Query(sqlRead)
	if err != nil {
		panic(err)
	}

	return rows, err
}

// We read all data about task by task name
func ReadTask(db *sql.DB, name string)(row *sql.Row){

	sqlRead := "SELECT id, name, uuid FROM tasks WHERE name='"+name+"'"

	row = db.QueryRow(sqlRead)

	return row
}

// We read all about group with selected task uuid - to get task id to get task name for group
func ReadGroupTask(db *sql.DB, uuid string)(row *sql.Row){

	sqlRead := "SELECT group_id FROM GroupsToTasks WHERE task_id='"+uuid+"'"

	row = db.QueryRow(sqlRead)

	return row
}

// We read all data about timeframe by timeframe uuid
func ReadTimeframe(db *sql.DB, uuid string)(row *sql.Row){

	sqlRead := "SELECT id, start, stop, uuid FROM Timeframes WHERE uuid='"+uuid+"'"

	row = db.QueryRow(sqlRead)

	return row
}

// We read all timeframes in Timeframes table
func ReadTimeframes( db *sql.DB)(rows *sql.Rows, err error){

	sqlRead := "SELECT * FROM Timeframes"

	rows, err = db.Query(sqlRead)
	if err != nil {
		panic(err)
	}

	return rows, err
}

// We read all tasks for specified group from GroupsToTasks table
func ReadGroupTasks(db *sql.DB, group_id string)(rows *sql.Rows, err error){

	sqlRead := "SELECT DISTINCT task_id FROM GroupsTotasks WHERE group_id='"+group_id+"'"

	rows, err = db.Query(sqlRead)
	if err != nil {
		panic(err)
	}

	return rows, err
}

// We read task name by using task id
func ReadTaskName(db *sql.DB, uuid string)(name string){

	sqlRead := "SELECT name FROM tasks WHERE uuid='"+uuid+"'"

	row := db.QueryRow(sqlRead)
	err := row.Scan(&name)
	if err != nil {
		panic(err)
	}

	return name
}

func ReadTaskID(db *sql.DB, name string)(uuid string){

	sqlRead := "SELECT uuid FROM tasks WHERE name='"+name+"'"

	row := db.QueryRow(sqlRead)
	err := row.Scan(&uuid)
	if err != nil {
		panic(err)
	}

	return uuid
}

// We read all timeframes for specified task from TasksToTimeframes table
func ReadTaskTimeframes(db *sql.DB, task_id string)(rows *sql.Rows, err error){

	sqlRead := "SELECT DISTINCT timeframe_id FROM TasksToTimeframes WHERE task_id='"+task_id+"'"

	rows, err = db.Query(sqlRead)
	if err != nil {
		panic(err)
	}

	return rows, err
}

// We read start stop timeframe forspecified task by using uuid
func ReadTimeframeStartStop(db *sql.DB, uuid string)(start, stop string){

	sqlRead := "SELECT start, stop FROM Timeframes WHERE uuid='"+uuid+"'"

	row := db.QueryRow(sqlRead)
	err := row.Scan(&start, &stop)
	if err != nil {
		panic(err)
	}

	return start, stop
}

// ========================= Delete block ================================================

// Delete timeframe by using timeframe id from Timeframes table
func DeleteTimeframe(db *sql.DB, uuid string)(err error){

	sqlRead := "DELETE FROM Timeframes WHERE uuid='"+uuid+"'"

	_, err = db.Exec(sqlRead)

	return err
}

// Delete timeframe by using timeframe id from TasksToTimeframes table
func DeleteTasksToTimeframes(db *sql.DB, uuid string)(err error){

	sqlRead := "DELETE FROM TasksToTimeframes WHERE timeframe_id='"+uuid+"'"

	_, err = db.Exec(sqlRead)

	return err
}

// Delete task from Tasks table by using uuid
func DeleteTask(db *sql.DB, uuid string)(err error){

	sqlRead := "DELETE FROM Tasks WHERE uuid='"+uuid+"'"

	_, err = db.Exec(sqlRead)

	return err
}

// Delete task from GroupsToTasks table by using uuid
func DeleteGroupsToTasks(db *sql.DB, uuid string)(err error){

	sqlRead := "DELETE FROM GroupsToTasks WHERE task_id='"+uuid+"'"

	_, err = db.Exec(sqlRead)

	return err
}

// Delete task from TasksToTimeframes by using uuid
func DeleteTasksTimeframes(db *sql.DB, uuid string)(err error){

	sqlRead := "DELETE FROM TasksToTimeframes WHERE task_id='"+uuid+"'"

	_, err = db.Exec(sqlRead)

	return err
}

// Delete group from Groups by using uuid
func DeleteGroup(db *sql.DB, uuid string)(err error){

	sqlRead := "DELETE FROM Groups WHERE uuid='"+uuid+"'"

	_, err = db.Exec(sqlRead)

	return err
}

// Delete group from GroupsToTasks by using uuid
func DeleteGroupsTasks(db *sql.DB, uuid string)(err error){

	sqlRead := "DELETE FROM GroupsToTasks WHERE group_id='"+uuid+"'"

	_, err = db.Exec(sqlRead)

	return err
}

//======================== Check block ================================

// Check if Group exist by name
func ExistGroupName(db *sql.DB, name string)(result string, err error){

	sqlRead := "SELECT EXISTS(SELECT * FROM Groups WHERE name='"+name+"')"

	row := db.QueryRow(sqlRead)
	err = row.Scan(&result)

	return result, err
}


// Check if Group exist by uuid
func ExistGroupUUID(db *sql.DB, uuid string)(result string, err error){

	sqlRead := "SELECT EXISTS(SELECT * FROM Groups WHERE uuid='"+uuid+"')"

	row := db.QueryRow(sqlRead)
	err = row.Scan(&result)

	return result, err
}

// Check if Task exist by name
func ExistTaskName(db *sql.DB, name string)(result string, err error){

	sqlRead := "SELECT EXISTS(SELECT * FROM Tasks WHERE name='"+name+"')"

	row := db.QueryRow(sqlRead)
	err = row.Scan(&result)

	return result, err
}

// Check if Task exist by uuid
func ExistTaskUUID(db *sql.DB, uuid string)(result string, err error){

	sqlRead := "SELECT EXISTS(SELECT * FROM Tasks WHERE uuid='"+uuid+"')"

	row := db.QueryRow(sqlRead)
	err = row.Scan(&result)

	return result, err
}

func ExistTimeframeUUID(db *sql.DB, uuid string)( result string, err error){

	sqlRead := "SELECT EXISTS(SELECT * FROM Timeframes WHERE uuid='"+uuid+"')"

	row := db.QueryRow(sqlRead)
	err = row.Scan(&result)

	return result, err
}

// ======================= Update block ===================================

// Update task name in Tasks table
func UpdateTaskTableName(db *sql.DB, nameOld string, nameNew string)(err error){

	sqlRead := "UPDATE Tasks SET name='"+nameNew+"' WHERE name='"+nameOld+"'"

	_, err = db.Exec(sqlRead)

	return err
}

// Update task UUID in Tasks table
func UpdateTaskTableUUID(db *sql.DB, uuidOld string, uuidNew string)(err error){

	sqlRead := "UPDATE Tasks SET uuid='"+uuidNew+"' WHERE uuid='"+uuidOld+"'"

	_, err = db.Exec(sqlRead)

	return err
}

// Update task UUID in GroupsToTasks table
func UpdateGTTTaskUUID(db *sql.DB, uuidOld string, uuidNew string)(err error){

	sqlRead := "UPDATE GroupsToTasks SET task_id='"+uuidNew+"' WHERE task_id='"+uuidOld+"'"

	_, err = db.Exec(sqlRead)

	return err
}

// Update task UUID in TasksToTimeframes table
func UpdateTTTTaskUUID(db *sql.DB, uuidOld string, uuidNew string)(err error){

	sqlRead := "UPDATE TasksToTimeframes SET task_id='"+uuidNew+"' WHERE task_id='"+uuidOld+"'"

	_, err = db.Exec(sqlRead)

	return err
}

// Update relation between task and group in GroupsToTasks table
func UpdateGTTRelationGroup(db *sql.DB, taskIdOld string, groupIdNew string)(err error){

	sqlRead := "UPDATE GroupsToTasks SET group_id='"+groupIdNew+"' WHERE task_id='"+taskIdOld+"'"

	_, err = db.Exec(sqlRead)

	return err
}

// Update group name in Groups table
func UpdateGroupTableName(db *sql.DB, nameOld string, nameNew string)(err error){

	sqlRead := "UPDATE Groups SET name='"+nameNew+"' WHERE name='"+nameOld+"'"

	_, err = db.Exec(sqlRead)

	return err
}

// Update group UUID in Groups table
func UpdateGroupTableUUID(db *sql.DB, uuidOld string, uuidNew string)(err error){

	sqlRead := "UPDATE Groups SET uuid='"+uuidNew+"' WHERE uuid='"+uuidOld+"'"

	_, err = db.Exec(sqlRead)

	return err
}

// Update group UUID in GroupsToTasks table
func UpdateGTTGroupUUID(db *sql.DB, uuidOld string, uuidNew string)(err error){

	sqlRead := "UPDATE GroupsToTasks SET group_id='"+uuidNew+"' WHERE group_id='"+uuidOld+"'"

	_, err = db.Exec(sqlRead)

	return err
}
