package main

import (
	"context"
	"fmt"
	"log"
	blogpb "swiggy/grpc/blog/blogpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm a client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("coul not connet %f", err)
	}
	defer cc.Close()
	c := blogpb.NewBlogServiceClient(cc)
	doUnary(c)
}

func doUnary(c blogpb.BlogServiceClient) {
	fmt.Println("Creating New blog")
	blog := &blogpb.Blog{
		AuthorId: "ishan",
		Title:    "my first blog",
		Content:  "Content of first Blog",
	}
	res, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatalf("Error while calling blog RPC %v", err)
	}
	log.Printf("Blog has been craeted: %v", res.Blog)
}
