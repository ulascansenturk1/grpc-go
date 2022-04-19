package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/protobuf/types/known/emptypb"

	genpb "github.com/ulascansenturk/grpc-go/genpb/article"
	"google.golang.org/grpc"
)

func main() {
	//ctx := context.Background()
	//Getting port from env
	port := os.Getenv("PORT")
	// if port is empty, use 8080
	if port == "" {
		port = "8080"
	}
	//get address from env
	addr := os.Getenv("LISTEN_ADDR")

	listenAddr := fmt.Sprintf("%s:%s", addr, port)

	lis, err := net.Listen("tcp", listenAddr)

	if err != nil {
		log.Fatalf("failed to listen on : %s:%v", listenAddr, err)
	}

	grpcServer := grpc.NewServer()

	articleServer := new(myArticleServer)

	genpb.RegisterArticleServiceServer(grpcServer, articleServer)

	fmt.Println("Server is up and running on port: ", port)

	log.Fatal(grpcServer.Serve(lis))

}

type myArticleServer struct{}

func (m myArticleServer) CreateArticle(ctx context.Context, a *genpb.Article) (*genpb.Article, error) {
	log.Printf("Create Article %v", a)
	return &genpb.Article{}, nil
}

func (m myArticleServer) DeleteArticle(ctx context.Context, request *genpb.DeleteArticleRequest) (*emptypb.Empty, error) {

	log.Printf("Delete Article %v", request)
	return &emptypb.Empty{}, nil

}

func (m myArticleServer) GetArticle(ctx context.Context, request *genpb.GetArticleRequest) (*genpb.GetArticleResponse, error) {
	log.Printf("Get Article %v", request)
	return &genpb.GetArticleResponse{}, nil
}

func (m myArticleServer) GetArticleList(ctx context.Context, request *genpb.GetArticleListRequest) (*genpb.ArticleList, error) {
	log.Printf("Get Article List %v", request)
	return &genpb.ArticleList{}, nil
}
