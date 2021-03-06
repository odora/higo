package main

import (
	"context"
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	proto "higo/h_micro/demo_two/proto"
	"os"
)

type logWrapper struct {
	client.Client
}

func (l *logWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	md, _ := metadata.FromContext(ctx)
	fmt.Printf("[Log Wrapper] ctx : %v service : %s method: %s\n", md, req.Service(), req.Method())
	return l.Client.Call(ctx, req, rsp)
}

func logWrap(c client.Client) client.Client {
	return &logWrapper{c}
}

func logHandlerWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		fmt.Printf("[Log Wrapper] Before serving request method : %v\n", req.Method())
		err := fn(ctx, req, rsp)
		fmt.Println("[Log Wrapper] After serving request")
		return err
	}
}

type Greeter struct {
}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func runClient(service micro.Service) {
	greeter := proto.NewGreeterService("greeter", service.Client())

	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "John"})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp.Greeting)
}

func main() {
	service := micro.NewService(
		micro.Client(client.NewClient(client.Wrap(logWrap))),
		micro.Server(server.NewServer(server.WrapHandler(logHandlerWrapper))),
		micro.Name("greeter"),
		micro.Version("latest"),
		micro.Metadata(map[string]string{
			"type": "helloworld",
		}),

		micro.Flags(cli.BoolFlag{
			Name:  "run_client",
			Usage: "Launch the client",
		}),
	)

	service.Init(
		micro.Action(func(i *cli.Context) {
			if i.Bool("run_client") {
				runClient(service)
				os.Exit(0)
			}
		}),
	)

	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
