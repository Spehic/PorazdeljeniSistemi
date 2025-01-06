package main

import (
		"fmt"
		"flag"
		"cetrtadomaca/storage"
		"cetrtadomaca/protobufStorage"
		"context"
		"net"
		"os"
		"time"
		"google.golang.org/grpc"
		"google.golang.org/protobuf/types/known/emptypb"
		"google.golang.org/grpc/credentials/insecure"
)

func odjemalec() {
	time.Sleep(1 * time.Second)

	// za pisanje rabimo glavo
	urlHead := fmt.Sprintf("%v:%v",inUrl, port)
	fmt.Printf("gRPC client connecting to (head) %v\n", urlHead)
	connHead, err := grpc.Dial(urlHead, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer connHead.Close()

	contextCRUDHead, cancelHead := context.WithTimeout(context.Background(), time.Second)
	defer cancelHead()

	grpcClientHead := protobufStorage.NewCRUDClient(connHead)

	todo1 := protobufStorage.Todo{Task: "Naredi IRZ", Completed: false}
	fmt.Println("1. Put:", todo1.Task)
	if _, err := grpcClientHead.Put(contextCRUDHead, &todo1); err != nil {
		panic(err)
	}
	fmt.Println("1 done")

	todo2 := protobufStorage.Todo{Task: "Nauči se devops", Completed: true}
	fmt.Println("2. Put:", todo2.Task)
	if _, err := grpcClientHead.Put(contextCRUDHead, &todo2); err != nil {
		panic(err)
	}
	fmt.Println("2 done")

	todo3 := protobufStorage.Todo{Task: "Tek", Completed: true}
	fmt.Println("3. Put:", todo3.Task)
	if _, err := grpcClientHead.Put(contextCRUDHead, &todo3); err != nil {
		panic(err)
	}
	fmt.Println("3 done")


	// za branje rabimo rep

	urlTail := fmt.Sprintf("%v:%v",inUrl, port + maxId - 2)
	fmt.Printf("gRPC client connecting to (tail) %v\n", urlTail)
	connTail, err := grpc.Dial(urlTail, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer connTail.Close()

	contextCRUDTail, cancelTail := context.WithTimeout(context.Background(), time.Second)
	defer cancelTail()

	grpcClientTail := protobufStorage.NewCRUDClient(connTail)

	fmt.Println("1. Getting:", todo3.Task)
	retObj, err := grpcClientTail.Get(contextCRUDTail, &todo3)
	if err != nil {
		panic(err)
	}
	fmt.Println("1 get done:", retObj.Todos)

	fmt.Println("2. Get all")
	retAllObj, err := grpcClientTail.Get(contextCRUDTail, &protobufStorage.Todo{Task: "", Completed: true})
	if err != nil {
		panic(err)
	}
	fmt.Println("2 get done:", retAllObj.Todos)

}

var id int
var maxId int
var port int
var inUrl string

func main () {
	// preberemo argumente iz ukazne vrstice
	urlPtr := flag.String("s", "localhost", "server URL")
	portPtr := flag.Int("p", 9876, "port number")
	idPtr := flag.Int("id", 0, "id of process")
	nPtr := flag.Int("n", 1, "amount of processes")
	flag.Parse()

	

	url := fmt.Sprintf("%v:%v", *urlPtr, *portPtr + *idPtr)
	id = *idPtr
	maxId = *nPtr
	inUrl = *urlPtr
	port = *portPtr

	fmt.Println("Hello, world!", *urlPtr, *idPtr, *portPtr, *nPtr, url)

	if *idPtr == *nPtr - 1 {
		odjemalec()
	} else {
		server( url )
	}

}

func server(url string) {
		// pripravimo strežnik gRPC
		grpcServer := grpc.NewServer()

		// pripravimo strukturo za streženje metod CRUD na shrambi TodoStorage
		crudServer := NewServerCRUD()
	
		// streženje metod CRUD na shrambi TodoStorage povežemo s strežnikom gRPC
		protobufStorage.RegisterCRUDServer(grpcServer, crudServer)
	
		// izpišemo ime strežnika
		hostName, err := os.Hostname()
		fmt.Println(hostName)
		if err != nil {
			panic(err)
		}
		// odpremo vtičnico
		listener, err := net.Listen("tcp", url)
		if err != nil {
			panic(err)
		}
		

		fmt.Printf("gRPC server listening at %v%v\n", hostName, url)
		// začnemo s streženjem
		if err := grpcServer.Serve(listener); err != nil {
			panic(err)
		}
	return
}



// struktura za strežnik CRUD za shrambo TodoStorage
type serverCRUD struct {
	protobufStorage.UnimplementedCRUDServer
	todoStore *storage.TodoStorage
}

// pripravimo nov strežnik CRUD za shrambo TodoStorage
func NewServerCRUD() *serverCRUD {
	todoStorePtr := storage.NewTodoStorage()
	return &serverCRUD{protobufStorage.UnimplementedCRUDServer{}, todoStorePtr}
}

func (s *serverCRUD) Put(ctx context.Context, in *protobufStorage.Todo) (*emptypb.Empty, error) {
	fmt.Println("Called, started put procedure on", inUrl,port+id)
	var urlNext, urlPrev string
	var connNext, connPrev *grpc.ClientConn
	var errs error // Declare err to handle connection errors

	var contextCRUDNext, contextCRUDPrev  context.Context
	var cancelNext, cancelPrev context.CancelFunc

	var grpcClientNext, grpcClientPrev protobufStorage.CRUDClient

	// ni rep
	if id != maxId-2 {
		urlNext = fmt.Sprintf("%v:%v", inUrl, port+id+1)
		connNext, errs = grpc.Dial(urlNext, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if errs != nil {
			fmt.Printf("Connection failesd (%v): %v\n", urlNext, errs)
		} else {
			fmt.Printf("Server connected to %v\n", urlNext)
		}
		defer connNext.Close() // Ensure connection is closed when no longer needed
		
		contextCRUDNext, cancelNext = context.WithTimeout(context.Background(), time.Second)
		defer cancelNext()

		// vzpostavimo vmesnik gRPC
		grpcClientNext = protobufStorage.NewCRUDClient(connNext)
	}

	// ni glava
	if id != 0 {
		urlPrev = fmt.Sprintf("%v:%v", inUrl, port+id-1)
		connPrev, errs = grpc.Dial(urlPrev, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if errs != nil {
			fmt.Printf("Connection failesd (%v): %v\n", urlPrev, errs)
		} else {
			fmt.Printf("Server connected to %v\n", urlPrev)
		}
		defer connPrev.Close() // Ensure connection is closed when no longer needed

		contextCRUDPrev, cancelPrev = context.WithTimeout(context.Background(), time.Second)
		defer cancelPrev()

		// vzpostavimo vmesnik gRPC
		grpcClientPrev = protobufStorage.NewCRUDClient(connPrev)
	}

	var ret struct{}
	err := s.todoStore.Put(&storage.Todo{Task: in.Task, Completed: in.Completed}, &ret)

	// 
	if id == maxId - 2 {
		err := s.todoStore.Commit(&storage.Todo{Task: in.Task, Completed: in.Completed}, &ret)
		if err != nil {
			fmt.Printf("Commit failed\n")
			return &emptypb.Empty{}, err
		}
		if id != 0 {
			fmt.Println("Calling commit (from tail) to",  urlPrev, "Task:", in.Task)
			if _, err := grpcClientPrev.Commit(contextCRUDPrev, &protobufStorage.Todo{Task: in.Task, Completed: in.Completed}); err != nil {
				panic(err)
			}
			fmt.Println("Commited", in.Task)
		}

	} else {
		fmt.Println("Calling Put to:",urlNext, "Task:", in.Task)
		if _, err := grpcClientNext.Put(contextCRUDNext, &protobufStorage.Todo{Task: in.Task, Completed: in.Completed}); err != nil {
			panic(err)
		}
		fmt.Println("Done put:", in.Task)

	}
	return &emptypb.Empty{}, err
}

func (s *serverCRUD) Get(ctx context.Context, in *protobufStorage.Todo) (*protobufStorage.TodoStorage, error) {
	fmt.Println("Getting", in.Task)
	if id != maxId - 2 {
		return &protobufStorage.TodoStorage{}, nil
	}
	
	dict := make(map[string](storage.Todo))
	err := s.todoStore.Get(&storage.Todo{Task: in.Task, Completed: in.Completed}, &dict)
	pbDict := protobufStorage.TodoStorage{}
	for k, v := range dict {
		pbDict.Todos = append(pbDict.Todos, &protobufStorage.Todo{Task: k, Completed: v.Completed})
	}
	fmt.Println("Returning", pbDict.Todos, err)
	return &pbDict, err
}

func (s *serverCRUD) Commit(ctx context.Context, in *protobufStorage.Todo) (*emptypb.Empty, error) {
	var ret struct{}
	err := s.todoStore.Commit(&storage.Todo{Task: in.Task, Completed: in.Completed}, &ret)
	if id == 0 {
		return &emptypb.Empty{}, err
	}
	
	urlPrev := fmt.Sprintf("%v:%v", inUrl, port+id-1)
	connPrev, errs := grpc.Dial(urlPrev, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if errs != nil {
		fmt.Printf("Connection failesd (%v): %v\n", urlPrev, errs)
	} else {
		fmt.Printf("Server connected to %v\n", urlPrev)
	}
	defer connPrev.Close() // Ensure connection is closed when no longer needed

	contextCRUDPrev, cancelPrev := context.WithTimeout(context.Background(), time.Second)
	defer cancelPrev()

	// vzpostavimo vmesnik gRPC
	grpcClientPrev := protobufStorage.NewCRUDClient(connPrev)
	fmt.Println("Calling commit to",  urlPrev, "Task:", in.Task)
	if _, err := grpcClientPrev.Commit(contextCRUDPrev, &protobufStorage.Todo{Task: in.Task, Completed: in.Completed}); err != nil {
		panic(err)
	}
	fmt.Println("Commited:", in.Task)
	return &emptypb.Empty{}, nil
}
	