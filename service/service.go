package service

import (
	"evenApi/config"
	"evenApi/genproto"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

// type evenApiI interface {
// 	goEvent() genproto.EventServiceClient
// }

// type serviceManager struct {
// 	evenService genproto.EventServiceClient
// }

// func (s *serviceManager) goEvent() genproto.EventServiceClient {
// 	return s.evenService
// }

type ToDoClient interface {
	ToDo() genproto.EventServiceClient
}

type clientManager struct {
	clientToDo genproto.EventServiceClient
}

func (c clientManager) ToDo() genproto.EventServiceClient {
	return c.clientToDo
}

func NewserviceManager(cfg *config.Config) (ToDoClient, error) {
	//resolver.SetDefaultScheme("dns")    // what is it?

	// connectService, err := grpc.Dial(
	// 	fmt.Sprintf("%s:%d", cfg.TodoServiceHost, cfg.TodoServicePort),
	// 	grpc.WithTransportCredentials(insecure.NewCredentials()),
	// )
	// if err != nil {
	// 	return nil, err
	// }
	// serviceManager := &serviceManager{
	// 	evenService: genproto.NewEventServiceClient(connectService),
	// }

	// return serviceManager, nil

	connectService, err := grpc.Dial(fmt.Sprintf("%s:%d", cfg.EventServiceHost, cfg.EventServicePort), grpc.WithInsecure())
	log.Println("connection grpc dial")
	if err != nil {
		log.Println("error occure")
		return nil, err
	}
	log.Println("returning connected new client")
	clientManager := &clientManager{clientToDo: genproto.NewEventServiceClient(connectService)}

	return clientManager, nil

}

/*
connTodo, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.TodoServiceHost, conf.TodoServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	serviceManager := &serviceManager{
		todoService: pb.NewTodoServiceClient(connTodo),
	}

	return serviceManager, nil*/
