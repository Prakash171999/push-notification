package main

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/readytowork-org/backend-RnD/config"
)

func main() {

	_, ctx, client := config.SetupFirebase()
	// sendToToken(app)
	// sendToTopic(ctx, client)
	sendMultiClients(ctx, client)
}

func sendToToken(app *firebase.App) {
	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	registrationToken := "d3adVu0tMDyBUTPkgc_l-0:APA91bHjb6-wWkT1ABGSasFqxrsOR3AdfcTjLc8b7f7yukWLt32GS4UA5XdIwZ8p98oOLp-CBcyuYaCYdEPRji_f2WSXO9JKb7XPjotm_3bdkk-7hJyxJS8JuUHt82xzGGJ6Aacy0QWb"

	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: "Firebase Notification",
			Body:  "This is firebase notification",
		},
		Token: registrationToken,
	}

	response, err := client.Send(ctx, message)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Successfully sent message:", response)
}

func sendToTopic(ctx context.Context, client *messaging.Client) {
	topic := "highScores"

	message := &messaging.Message{
		Data: map[string]string{
			"score": "850",
			"time":  "2:45",
		},
		Topic: topic,
	}

	response, err := client.Send(ctx, message)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Successfully sent message:", response)
}

func sendMultiClients(ctx context.Context, client *messaging.Client) {

	registrationTokens := []string{
		"d3adVu0tMDyBUTPkgc_l-0:APA91bHjb6-wWkT1ABGSasFqxrsOR3AdfcTjLc8b7f7yukWLt32GS4UA5XdIwZ8p98oOLp-CBcyuYaCYdEPRji_f2WSXO9JKb7XPjotm_3bdkk-7hJyxJS8JuUHt82xzGGJ6Aacy0QWb",
		"f3AdfhQi6VAIuB6hq39s31:APA91bFYwHrKSjdV9eUjHFGeupwcPuOjv1X4Z5aKrbz15jU3kmqHr8_0y4welxmGkTmYxsnnxQeQ0cq0reEcgDpzllr2XeRMb8OaFX7IZEOiOc54Q4-1KSF06bIS8otQVa3WQt359aIh",
	}
	message := &messaging.MulticastMessage{
		Notification: &messaging.Notification{
			Title: "Firebase Notification Multi",
			Body:  "This is firebase notification",
		},
		Tokens: registrationTokens,
	}

	br, err := client.SendMulticast(context.Background(), message)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%d messages were sent successfully\n", br.SuccessCount)
}
