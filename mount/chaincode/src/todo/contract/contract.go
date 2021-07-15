package contract

import (
	"encoding/json"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"fmt"
)

type SmartContract struct{
	contractapi.Contract
}

type Todo struct{
	ID string `json:"ID"`
	Title string `json:"title"`
	Done bool `json:"done"`
	Creator string `json:"creator"`
	Org string `json:"org"`
}

//create todo
func(s *SmartContract)CreateTodo(ctx contractapi.TransactionContextInterface, id string, title string, creator string, org string) error {
	//check if id exist
	if exist, _ := s.IsExist(ctx,id); exist == true{
		return fmt.Errorf("todo  %s already exist", id);
	}
	todo := Todo{
		ID: id,
		Title: title,
		Done: false,
		Creator: creator,
		Org: org,
	}
	todoJSON, err := json.Marshal(todo);

	if err != nil{
		return err;
	}
	return ctx.GetStub().PutState(id, todoJSON);
}

func (s *SmartContract) Fortune(ctx contractapi.TransactionContextInterface) error {
	s.CreateTodo(ctx, "TODO0", "a" , "Anupam", "none");
	s.CreateTodo(ctx, "TODO1", "b" , "Anupam", "none");
	s.CreateTodo(ctx, "TODO2", "c" , "Anupam", "none");
	s.CreateTodo(ctx, "TODO3", "d" , "Anupam", "none");
	s.CreateTodo(ctx, "TODO4", "e" , "Anupam", "none");
	s.CreateTodo(ctx, "TODO5", "f" , "Anupam", "none");
	return nil;
}

//read todo
func(s *SmartContract)ReadTodo(ctx contractapi.TransactionContextInterface, id string) (Todo, error){
	//is exist
	if exist, _:=s.IsExist(ctx,id); exist == false{
		return Todo{}, fmt.Errorf("todo  %s does not exist", id);
	}

	todoJSON, err := ctx.GetStub().GetState(id);
	if err != nil {
		return Todo{},err;
	}
	
	var todo Todo;
	json.Unmarshal(todoJSON, &todo);
	return todo, nil;
}


//update todo
func(s *SmartContract)UpdateTodo(ctx contractapi.TransactionContextInterface, id string, title string, creator string ,org string)error{
	//is exist
	if exist, _:=s.IsExist(ctx,id); exist == false{
		return fmt.Errorf("todo  %s does not exist", id);
	}

	//is owner
	todot, err := s.ReadTodo(ctx, id);
	if err != nil {
		return err;
	}

	if todot.Creator != creator{
		return fmt.Errorf("you are not the creator of this todo");
	}
	
	todo := Todo{
		ID: id,
		Title: title,
	}
	todoJSON, err := json.Marshal(todo);

	if err != nil{
		return err;
	}
	return ctx.GetStub().PutState(id, todoJSON);
}

//delete todo
func(s *SmartContract)DeleteTodo(ctx contractapi.TransactionContextInterface, id string)error{
	//is exist
	if exist, _:=s.IsExist(ctx,id); exist == false{
		return fmt.Errorf("todo  %s does not exist", id);
	}
	return ctx.GetStub().DelState(id);
}

//set status
func(s *SmartContract)SetDone(ctx contractapi.TransactionContextInterface, id string)error{
	//is exist
	if exist, _:=s.IsExist(ctx,id); exist == false{
		return fmt.Errorf("todo  %s does not exist", id);
	}

	todo, err := s.ReadTodo(ctx, id);
	if err != nil {
		return err;
	}
	if todo.Done{
		fmt.Errorf("todo is already complete");
	}
	//check if true already
	todo.Done = true;
	todoJSON, err := json.Marshal(todo);
	if err != nil {
		return err;
	}
	return ctx.GetStub().PutState(id, todoJSON);
}
//retrieve all
func(s *SmartContract)GetAllTodo(ctx contractapi.TransactionContextInterface)([]*Todo,error){
	
	result, err := ctx.GetStub().GetStateByRange("","");
	if err != nil{
		return nil, err;
	}
	defer result.Close()
	var todos []*Todo
	for result.HasNext(){
		todoJSON, err:= result.Next()
		if err != nil{
			return nil, err;
		}
		var todo Todo
		json.Unmarshal(todoJSON.Value, &todo)

		todos = append(todos, &todo)
	}
	return todos, nil;
}

//retrieve all
func(s *SmartContract)IsExist(ctx contractapi.TransactionContextInterface, id string)(bool, error){
	todoJSON, err:=ctx.GetStub().GetState(id)
	if err != nil{
		return false, err;
	}
	if todoJSON != nil{
		return true, nil
	}
	return false, nil;
}
