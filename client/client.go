package main

import (
	"context"
	"log"
	"time"

	article "github.com/ulascansenturk/grpc-go/genpb/article"
	"google.golang.org/grpc"
)

func main() {

	ctx := context.Background()

	dialCtx, cleanUp := context.WithTimeout(ctx, time.Second*10)

	defer cleanUp()

	conn, err := grpc.DialContext(dialCtx, "localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	client := article.NewArticleServiceClient(conn)

	ac, err := client.CreateArticle(ctx, &article.Article{
		Name:        "First Article",
		Description: "This is the first article",
		Author:      "Ulas",
		Status:      1,
	})

	if err != nil {
		log.Fatalf("could not create article: %v", err)
	}
	log.Printf("Created article: ID: %v", ac.GetId())

	//delete article
	_, err = client.DeleteArticle(ctx, &article.DeleteArticleRequest{Id: ac.GetId()})
	if err != nil {
		log.Fatalf("could not delete article: %s, %v", ac.GetId(), err)
	}

}
