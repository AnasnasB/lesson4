package main

import "fmt"
import "encoding/json"
import "io/ioutil"

type Action struct {
	Action string `json:"action"`
	ObjName string `json:"object"`
}

type Teacher struct {
	ID string  `json:"id"`
	Salary float64 `json:"salary"`
	Subject string `json:"subject"`
	Classroom []string `json:"classroom"`
	Persons struct {
		Name string `json:"name"`
		Surname string `json:"surname"`
		PersonalCode string `json:"personalCode"`
	} `json:"person"`
}
type DeleteTeacher struct {
	Data struct {
		ID string `json:"id"`
	} `json:"data"`
}
type CreateTeacher struct {
	T Teacher `json:"data"`
}
type UpdateTeacher struct {
	T Teacher `json:"data"`
}
type ReadTeacher struct {
	Data struct {
		ID string `json:"id"`
	} `json:"data"`
}


type DefinedAction interface {
	GetFromJSON([]byte)
	Process([]GeneralObject)[]GeneralObject
}
type GeneralObject interface {
	GetCreateAction() DefinedAction
	GetUpdateAction() DefinedAction
	GetReadAction() DefinedAction
	GetDeleteAction() DefinedAction
}


func (t Teacher) GetCreateAction() DefinedAction {
	return &CreateTeacher{}
}
func (t Teacher) GetUpdateAction() DefinedAction {
	return &UpdateTeacher{}
}
func (t Teacher) GetReadAction() DefinedAction {
	return &ReadTeacher{}
}
func (t Teacher) GetDeleteAction() DefinedAction {
	return &DeleteTeacher{}
}



func (action *CreateTeacher) GetFromJSON (rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (action *UpdateTeacher) GetFromJSON (rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (action *ReadTeacher) GetFromJSON (rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (action *DeleteTeacher) GetFromJSON (rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}


func (action CreateTeacher) Process(list[]GeneralObject)[]GeneralObject{
	var t Teacher 
	t.ID = action.T.ID
	t.Salary = action.T.Salary
	t.Subject = action.T.Subject
	t.Classroom = action.T.Classroom
	t.Persons.Name = action.T.Name
	t.Surname = action.T.Surname
	t.PersonalCode = action.T.PersonalCode
	list = append(list,t)
	return list
}
func (action *UpdateTeacher) Process(list[]GeneralObject)[]GeneralObject{
	return list
}
func (action *ReadTeacher) Process(list[]GeneralObject)[]GeneralObject{
	return list
}
func (action *DeleteTeacher) Process(list[]GeneralObject)[]GeneralObject{
	return list
}

func main() {
	list:= make([]GeneralObject,0)
	text, err := ioutil.ReadFile("untitled.json")
	if err != nil {
		fmt.Println("Error",err)
		return
	}
	
	var act Action
	err = json.Unmarshal(text, &act)
	if err != nil {
		fmt.Println("Error  ",err)
		return
	}
	var obj GeneralObject
	switch act.ObjName {
	case "Teacher":
		obj = &Teacher{}
	}
	var toDo DefinedAction
	switch act.Action {
	case "create":
		toDo = obj.GetCreateAction()
	case "update":
		toDo = obj.GetUpdateAction()
	case "read":
		toDo = obj.GetReadAction()
	}
	
	toDo.GetFromJSON(text)
	
	toDo.Process(list)
	fmt.Print(toDo)
}
