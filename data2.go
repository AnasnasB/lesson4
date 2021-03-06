package main

import "fmt"
import "encoding/json"
import "io/ioutil"
import "reflect"

type Action struct {
	A string `json:"action"`
	ObjName string `json:"object"`
}

type Teacher struct {
	ID string  `json:"id"`
	Salary float64 `json:"salary"`
	Subject string `json:"subject"`
	Classroom []string `json:"classroom"`
	P Person `json:"person"`
}

type Student struct{
	ID string  `json:"id"`
	Class string `json:"class"`
	P Person `json:"person"`
}

type Staff struct{
	ID string  `json:"id"`
	Phone string `json:"phonenumber"`
	Classroom string `json:"classroom"`
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

type DeleteStudent struct {
	D Data `json:"data"`
}
type CreateStudent struct {
	S Student `json:"data"`
}
type UpdateStudent struct {
	S Student `json:"data"`
}
type ReadStudent struct {
	D Data `json:"data"`
}

type DeleteStaff struct {
	D Data `json:"data"`
}
type CreateStaff struct {
	S Staff `json:"data"`
}
type UpdateStaff struct {
	S Staff `json:"data"`
}
type ReadStaff struct {
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


func (s Staff) GetCreateAction() DefinedAction {
	return &CreateStaff{}
}
func (s Staff) GetUpdateAction() DefinedAction {
	return &UpdateStaff{}
}
func (s Staff) GetReadAction() DefinedAction {
	return &ReadStaff{}
}
func (s Staff) GetDeleteAction() DefinedAction {
	return &DeleteStaff{}
}


func (s Student) GetCreateAction() DefinedAction {
	return &CreateStudent{}
}
func (s Student) GetUpdateAction() DefinedAction {
	return &UpdateStudent{}
}
func (s Student) GetReadAction() DefinedAction {
	return &ReadStudent{}
}
func (s Student) GetDeleteAction() DefinedAction {
	return &DeleteStudent{}
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

func (action *CreateStudent) GetFromJSON (rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (action *UpdateStudent) GetFromJSON (rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (action *ReadStudent) GetFromJSON (rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (action *DeleteStudent) GetFromJSON (rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}


func (action *CreateStaff) GetFromJSON (rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (action *UpdateStaff) GetFromJSON (rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (action *ReadStaff) GetFromJSON (rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (action *DeleteStaff) GetFromJSON (rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
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

func (action *CreateStaff) Process(list[]GeneralObject)[]GeneralObject{
	var s Staff
	s.ID = action.S.ID
	s.Phone = action.S.Phone
	s.Classroom = action.S.Classroom
	s.P.Name = action.S.P.Name
	s.P.Surname = action.S.P.Surname
	s.P.PersonalCode = action.S.P.PersonalCode
	list = append(list,s)
	fmt.Print(s)
	return list
}

func (action *UpdateStaff) Process(list[]GeneralObject)[]GeneralObject{
	var s Staff
	s.ID = action.S.ID
	s.Phone = action.S.Phone
	s.Classroom = action.S.Classroom
	s.P.Name = action.S.P.Name
	s.P.Surname = action.S.P.Surname
	s.P.PersonalCode = action.S.P.PersonalCode
	for i := 0; i<len(list); i++ {
		if reflect.TypeOf(list[i]) == reflect.TypeOf(s){
			if list[i].(Staff).ID == action.S.ID{
				list[i] = s
			}
		}
	}
	return list
}

func (action *ReadStaff) Process(list[]GeneralObject)[]GeneralObject{
	var s Staff
	for i := 0; i<len(list); i++ {
		if reflect.TypeOf(list[i]) == reflect.TypeOf(s){
			if list[i].(Staff).ID == action.D.ID{
				fmt.Print(list[i].(Staff))
			}
		}
	}
	return list
}

func (action *DeleteStaff) Process(list[]GeneralObject)[]GeneralObject{
	var s Staff
	for i := 0; i<len(list); i++ {
		if reflect.TypeOf(list[i]) == reflect.TypeOf(s){
			if list[i].(Staff).ID == action.D.ID{
				Delete(list, i)
			}
		}
	}
	return list
}


func (action *CreateStudent) Process(list[]GeneralObject)[]GeneralObject{
	var s Student 
	s.ID = action.S.ID
	s.Class = action.S.Class
	s.P.Name = action.S.P.Name
	s.P.Surname = action.S.P.Surname
	s.P.PersonalCode = action.S.P.PersonalCode
	list = append(list,s)
	return list
}
func (action *UpdateStudent) Process(list[]GeneralObject)[]GeneralObject{
	var s Student 
	s.ID = action.S.ID
	s.Class = action.S.Class
	s.P.Name = action.S.P.Name 
	s.P.Surname = action.S.P.Surname
	s.P.PersonalCode = action.S.P.PersonalCode
	for i := 0; i<len(list); i++ {
		if reflect.TypeOf(list[i]) == reflect.TypeOf(s){
			if list[i].(Student).ID == action.S.ID{
				list[i] = s
			}
		}
	}
	return list
}

func (action *ReadStudent) Process(list[]GeneralObject)[]GeneralObject{
	var s Student
	for i := 0; i<len(list); i++ {
		if reflect.TypeOf(list[i]) == reflect.TypeOf(s){
			if list[i].(Student).ID == action.D.ID{
				fmt.Print(list[i].(Student))
			}
		}
	}
	return list
}

func (action *DeleteStudent) Process(list[]GeneralObject)[]GeneralObject{
	var s Student
	for i := 0; i<len(list); i++ {
		if reflect.TypeOf(list[i]) == reflect.TypeOf(s){
			if list[i].(Student).ID == action.D.ID{
				Delete(list, i)
			}
		}
	}
	return list
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
	for i := 0; i<len(list); i++ {
		if reflect.TypeOf(list[i]) == reflect.TypeOf(t){
			if list[i].(Teacher).ID == action.T.ID{
				list[i] = t
			}
		}
	}
	return list
}
func (action *ReadTeacher) Process(list[]GeneralObject)[]GeneralObject{
	var t Teacher 
	for i := 0; i<len(list); i++  {
		if reflect.TypeOf(list[i]) == reflect.TypeOf(t){
			if list[i].(Teacher).ID == action.D.ID{
				fmt.Print(list[i].(Teacher))
			}
		}
	}
	return list
}
func (action *DeleteTeacher) Process(list[]GeneralObject)[]GeneralObject{
	var t Teacher 
	for i := 0; i<len(list); i++ {
		if reflect.TypeOf(list[i]) == reflect.TypeOf(t){
			if list[i].(Teacher).ID == action.D.ID{
				Delete(list, i)
			}
		}
	}
	return list
}

func Delete(list[]GeneralObject ,s int) ([]GeneralObject){
	return append(list[:s], list[s+1:]...)
}


func main() {
	list:= make([]GeneralObject,0)
	text, err := ioutil.ReadFile("untitled2.json")
	if err != nil {
		fmt.Println("Error",err)
		return
	}
	inp := make (map[string]interface{})
	err = json.Unmarshal(text, &inp)
	if err != nil {
		fmt.Println("Error---",err)
		return
	}
	var act Action
	act.A = string(inp["action"].(string))
	act.ObjName = string(inp["object"].(string))
	var obj GeneralObject
	switch act.ObjName {
		case "Teacher":
			obj = &Teacher{}
		case "Student":
			obj = &Student{}
		case "Staff":
			obj = &Staff{}
	}
	var toDo DefinedAction
	switch act.A { 
		case "create":
			toDo = obj.GetCreateAction()
		case "update":
			toDo = obj.GetUpdateAction()
		case "read":
			toDo = obj.GetReadAction()
	}
	
	toDo.GetFromJSON(text)
	
	list = toDo.Process(list)
	
	fmt.Println(list)
	
}
