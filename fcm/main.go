package main

import (
	"context"
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()

	app, err := InitFirebase("fcm.json")
	if err != nil {
		log.Fatalf("init error: %v", err)
	}

	err = SendNotification(
		ctx,
		app,
		"cgBhOCuoAE20vYkzqFlF2K:APA91bH8ur-5IceOeSh0kN93EXNvl5MzCAnS6Edu2ubLxA4H541sFZokvTmMKxoYqWcb02RwEFHiMGCS6PZvAj0h_cSr05_ub5uVTcvfcE9g7449FA9ekXs",
		"Hello 👋",
		"This is a test notification",
		map[string]string{
			"type": "test",
		},
	)

	if err != nil {
		log.Fatalf("send error: %v", err)
	}
}

// InitFirebase initializes Firebase App
func InitFirebase(serviceAccountPath string) (*firebase.App, error) {
	jsonBytes, err := os.ReadFile(serviceAccountPath)
	if err != nil {
		return nil, fmt.Errorf("error reading service account file: %w", err)
	}

	creds, err := google.CredentialsFromJSON(context.Background(), jsonBytes,
		"https://www.googleapis.com/auth/firebase.messaging",
	)
	if err != nil {
		return nil, fmt.Errorf("error creating credentials: %w", err)
	}

	conf := &firebase.Config{ProjectID: creds.ProjectID}
	app, err := firebase.NewApp(context.Background(), conf,
		option.WithTokenSource(creds.TokenSource),
	)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %w", err)
	}
	return app, nil
}

// SendNotification sends a push notification to a single FCM token
func SendNotification(ctx context.Context, app *firebase.App, fcmToken string, title string, body string, data map[string]string) error {
	client, err := app.Messaging(ctx)
	if err != nil {
		return fmt.Errorf("failed to get messaging client: %w", err)
	}

	message := &messaging.Message{
		Token: fcmToken,
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Data: data,
	}

	response, err := client.Send(ctx, message)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	fmt.Printf("Successfully sent message: %s\n", response)
	return nil
}
