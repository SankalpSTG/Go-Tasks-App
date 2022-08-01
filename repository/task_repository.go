package repository

import (
	"database/sql"
	"fmt"
	"goapp/dao"
	"goapp/db"
)
var conn *sql.DB
func InsertTask(task *dao.Task) (int64, error) {
	curTask := *task
	res, err := conn.Exec("INSERT INTO tasks (task, created_at) VALUES (?, ?)", curTask.TaskName, curTask.CreatedAt)
	if(err != nil){
		return 0, err
	}
	SerialID, err := res.LastInsertId()
	if(err != nil){
		fmt.Print(err)
	}
	return SerialID, nil
}
func GetTasks(isCompleted *bool) ([]dao.Task, error)  {
	var results *sql.Rows
	if(isCompleted != nil){
		results, _ = conn.Query("SELECT serial_id, task, created_at, updated_at, is_completed, is_deleted FROM tasks WHERE is_deleted = 0 AND is_completed = ?", isCompleted)
	}else{
		results, _ = conn.Query("SELECT serial_id, task, created_at, updated_at, is_completed, is_deleted FROM tasks WHERE is_deleted = 0")
	}
	rows := []dao.Task{}
	for results.Next() {
		var r dao.Task
		err := results.Scan(&r.SerialID, &r.TaskName, &r.CreatedAt, &r.UpdatedAt, &r.IsCompleted, &r.IsDeleted)
		if(err != nil){
			fmt.Println(err)
			//Error Handling
		}else{
			rows = append(rows, r)
		}
	}
	return rows, nil
}
func UpdateTask(task *dao.Task) (error) {
	curTask := *task
	_, err := conn.Exec("UPDATE tasks SET task = ?, updated_at = ?, is_completed = ?, is_deleted = ? WHERE serial_id = ?", curTask.TaskName, curTask.UpdatedAt, curTask.IsCompleted, curTask.IsDeleted, curTask.SerialID)
	if(err != nil){
		return err
	}
	return nil
}
func SetDatabase(){
	fmt.Println("Setting Database")
	dbConn, _ := db.GetDatabase()
	conn = dbConn
}
func CloseDatabase(){
	fmt.Println("Closing Database")
	if(conn != nil){
		defer conn.Close()
	}
}