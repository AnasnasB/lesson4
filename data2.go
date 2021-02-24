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
	P Person `json:"person"`
}
type Person struct {
	Name string `json:"name"`
	Surname string `json:"surname"`
	PersonalCode string `json:"personalCode"`
	}
	
type DeleteTeacher struct {
	D Data `json:"data"`
}
type CreateTeacher struct {
	T Teacher `json:"data"`
}
type UpdateTeacher struct {
	T Teacher `json:"data"`
}
type ReadTeacher struct {
	D Data `json:"data"`
}
type Data struct {
		ID string `json:"id"`
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
	t.P.Name = action.T.P.Name
	t.P.Surname = action.T.P.Surname
	t.P.PersonalCode = action.T.P.PersonalCode
	list = append(list,t)
	return list
}
func (action *UpdateTeacher) Process(list[]GeneralObject)[]GeneralObject{
	var t Teacher
	t.ID = action.T.ID
	t.Salary = action.T.Salary
	t.Subject = action.T.Subject
	t.Classroom = action.T.Classroom
	t.P.Name = action.T.P.Name
	t.P.Surname = action.T.P.Surname
	t.P.PersonalCode = action.T.P.PersonalCode
	for _,u := range list {
		if u.(Teacher).ID == action.T.ID{
			u = t
		}
	}
	return list
}
func (action *ReadTeacher) Process(list[]GeneralObject)[]GeneralObject{
	for _,u := range list {
		if u.(Teacher).ID == action.D.ID{
			fmt.Print(u.(Teacher))
		}
	}
	return list
}
func (action *DeleteTeacher) Process(list[]GeneralObject)[]GeneralObject{
	for i := 0; i<len(list); i++ {
		if list[i].(Teacher).ID == action.D.ID{
			Delete(list, i)
		}
	}
	return list
}

func Delete(list[]GeneralObject ,s int) ([]GeneralObject){
	return append(list[:s], list[s+1:]...)
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
