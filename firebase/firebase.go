package firebase

import (
	"firebase.google.com/go"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

func FirebaseApp() (*firebase.App, error) {
	opt := option.WithCredentialsFile("./firebase_sdk_key.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		logrus.Fatal(fmt.Errorf("unable to setup firebase app"), err)
		return nil, err
	}

	return app, nil
}
