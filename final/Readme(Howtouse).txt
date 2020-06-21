
GENERAL DESCRIPTION
=============================================
This is a simple time tracker that makes it possible to track time of the specified activity.
The program uses html page interface for communication with the user. To use the program one
should use any web browser.

++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
!!!!! ONE SHOULD MODIFY 'userbase.go' FILE TO SET USERNAME AND PASSWORD TO ACCSESS TO DATABASE FROM NEW PC.

THIS PIECE OF CODE SHOULD BE MODIFIED:

//This is for database connect, should be changed in case of database transfer to new PC
const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "12345"
	dbname = "final"
)
++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

The program can do the following actions:

1) Create name of the Group of tasks;
2) Create Task and assign in to selected Group;
3) Create Timeframe and assign it to specified Task;

4) Get data about selected Group;
5) Get data about selected Task;
6) Get data about timeframes;

7) Delete Timeframe by uuid;
8) Delete Task by name;
9) Delete Group by name;

10) Update data for specified Task (except timeframes);
11) Update data for specified Group (except timeframes).

All data is stored in Postgres database. Since each user has it own setup parameters for Postgres
(namely host name, user password) in some circumstances the end user will have to setup necessary data
for connection of the program with database. The description of this process will be placed below.

INITIAL PROGRAM SETUP
==========================================================
The program uses two side packages:


1) "github.com/google/uuid"
2) "github.com/lib/pq"

that must be installed for correct work of the program. 

It is highly recommended to place program to '\go\src\' folder since there can be an issue with packages.

INITIAL DATABASE SETUP
=========================================================
The program is distributed with dummy Postgres database, so it can be loaded from root folder.
Database name is 'final.sql'. It works correctly only with Postgres. 

However, if something goes wrong, the database can easily be created by using the next data.

#This is a brief description of how to create database in case postgres dump is not ok.
#The command below can be directly copied to Postgres Shell in order to create database valid for the program.
#If something goes wrong, one should modify 'userbase.go' file, namely 'const' piece, including:'host', 'port', 'user', 'password', 'dbname'
#since by default this block is set up for original author's PC. 

1). CREATE DATABASE final;     #(We create database in Postgres).

2). \c final;                  #(We connect to database in Postgres).

3). CREATE TABLE Tasks(id SERIAL PRIMARY KEY, name TEXT UNIQUE, uuid TEXT UNIQUE);

4). CREATE TABLE Groups(id SERIAL PRIMARY KEY, name TEXT UNIQUE, uuid TEXT UNIQUE);

5). CREATE TABLE Timeframes(id SERIAL PRIMARY KEY, start TEXT, stop TEXT, uuid TEXT UNIQUE);

6). CREATE TABLE GroupsToTasks(id SERIAL PRIMARY KEY, group_id TEXT, task_id TEXT);

7). CREATE TABLE TasksToTimeframes(id SERIAL PRIMARY KEY, task_id TEXT, timeframe_id TEXT);

#The commands above create 'final' database with all necessary tables.


HOW TO USE THE PROGRAM
============================================================
The program uses html interface to communicate both with end user and Postgres database. 
To run the program, one should do the following steps:

1) Run 'main.go' file
2) In any web browser type: 'localhost:8080'

This will run the 'new.html' file with control buttons. 

CONTROL BUTTONS
==========================================================
1) POST Group -  creates new Group with 'Name of the group:' with unique uuid. Name of the group should be unique. 
2) POST Task - creates new Task with 'Name of the task:' with unique uuid and assigns it to the 'Name of the group:'. Name of the task should be unique. 
                    The group for the task should exist in database. Each group can have any number of tasks.
3) POST Timeframe - creates new Timeframe with 'Time of start' and 'Time of stop' with unique uuid. Timeframe is assigned to the 'Name of the task:'. 
                    The task for the timeframe should exist in database. Each task can have any number of timeframes.

4) GET Groups - if 'Get the groups:' is blank, gives all data about every group in database. If 'Get the groups:' possess name of the existing group, it gives all data about this group.
5) GET Tasks - if 'Get the tasks:' is blank, gives all data about every task in database. If 'Get the tasks:' possess name of the existing task, it gives all data about this task.
6) GET Tifeframes - gives a list of all timeframes with uuid, for user convenience. 

7) DELETE Timeframe - delete the Timeframe from database by uuid. It deletes the Timeframe from all tables and destroys any relations.
8) DELETE Task - delete the Task from database by name. It deletes the Task from all tables and destroys any relations. 
9) DELETE Group - delete the Group from database by name. It deletes the Group from all tables and destroys any relations.

10) PUT task - updates data of the Task that already exists in database. If one does not want to change the selected data, the corresponding field should left blank. 
11) PUT group - updates data of the Group that already exists in database. If one does not want to change the selected data, the corresponding field should left blank.

(Example: we have task with name "Reading". We want to change it to new name, like "Communicating with book", but we want to leave the rest data the same. 
          We type "Reading" in 'Task name:' field, we type new name "Communicating with book" in 'New task name:', we leave 'New group assignment:' and 'New task uuid:' blank. 
          The program will change the name of the task, but leave the Task group and Task uuid as it was)

STUFF
===========================================
The program handles major errors by using panic, so it continue to work even if there is a lot of panic in the console. One should ignore it.

  
