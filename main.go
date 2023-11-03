package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"github.com/shomali11/slacker"
	"strconv"
	"time"
	"github.com/joho/godotenv"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent){
	for event:=range analyticsChannel{
		fmt.Println("Command Event")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main(){
	e := godotenv.Load()
	if e != nil {
		log.Fatal("Error loading .env file")
	}


	bot:= slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"),os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my yob is <year>",&slacker.CommandDefinition{
		Description: "yob claculator",
		// Examples: "my yob is 2020",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter){
			year:= request.Param("year")
			yob,err:=strconv.Atoi(year)
			if err!=nil{
				println("error")
			}
			age:= time.Now().Year()-yob
			r:=fmt.Sprintf("age is %d", age)
			response.Reply(r)
		},
	})

	ctx, cancel:=context.WithCancel(context.Background())
	defer cancel()
	err :=bot.Listen(ctx)
	if err !=nil{
		log.Fatal(err)
	}
}