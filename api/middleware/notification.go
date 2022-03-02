package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/blood-bank-backend/api/routeutils"
	"gitlab.com/Aubichol/blood-bank-backend/chat/dto"
	storetoken "gitlab.com/Aubichol/blood-bank-backend/store/token"
	"gitlab.com/Aubichol/blood-bank-backend/user"
	"go.uber.org/dig"
	"google.golang.org/api/option"
)

//Draft stores notification related information
type Draft struct {
	Blocks []struct {
		Key               string        `json:"key"`
		Text              string        `json:"text"`
		Type              string        `json:"type"`
		Depth             int           `json:"depth"`
		InlineStyleRanges []interface{} `json:"inlineStyleRanges"`
		EntityRanges      []interface{} `json:"entityRanges"`
		Data              struct {
		} `json:"data"`
	} `json:"blocks"`
	EntityMap struct {
	} `json:"entityMap"`
}

//MessageNotification sends messages to firebase
type MessageNotification struct {
	userProfiler user.MyProfiler
	userToken    storetoken.Token
	//	userSessionVerifier user.SessionVerifier
}

//Middleware implments a middleware that checkes the user session for authorization
func (a *MessageNotification) Middleware(h http.Handler) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		r.Body = ioutil.NopCloser(bytes.NewBuffer(data))

		if err != nil {
			logrus.Error(err)
			return
		}

		reader := bytes.NewReader(data)

		message := dto.Message{}

		if err := message.FromReader(reader); err != nil {
			logrus.Error(err)
			routeutils.ServeError(w, err)
			return
		}

		var draft Draft
		json.Unmarshal([]byte(message.Message), &draft)
		TextToSend := draft.Blocks[0].Text

		selfID := r.Context().Value("userID").(string)
		friendID := chi.URLParam(r, "id")

		me, err := a.userProfiler.Me(selfID)

		if err != nil {
			fmt.Println(err)
			h.ServeHTTP(w, r)
			return
		}

		usertoken, err := a.userToken.FindByUser(friendID)

		if err != nil {
			h.ServeHTTP(w, r)
			return
		}

		opt := option.WithCredentialsFile("/root/src/src/gitlab.com/Aubichol/blood-bank-backend/firebase/fire.json")

		app, err := firebase.NewApp(context.Background(), nil, opt)

		if err != nil {
			fmt.Errorf("error initializing app: %v", err)
		}

		if err != nil {
			fmt.Println(err)
		}

		ctx := context.Background()
		meclient, err := app.Messaging(ctx)

		if err != nil {
			logrus.Error("error getting Messaging client: %v\n", err)
		}

		notification := &messaging.Notification{
			Title: me.FirstName + " sent you a message.",
			Body:  TextToSend,
		}

		for i := 0; i < len(usertoken); i++ {
			mes := &messaging.Message{
				Data: map[string]string{
					"sender":  me.FirstName,
					"message": message.Message,
				},
				Notification: notification,
				Token:        usertoken[i].Token,
			}

			response, err := meclient.Send(ctx, mes)

			if err != nil {
				logrus.Error(err)
			}
			// Response is a message ID string.
			fmt.Println("Successfully sent message:", response)
		}

		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(f)
}

//MessageNotificationParams provide parameters for notification middleware
type MessageNotificationParams struct {
	dig.In
	UserProfile user.MyProfiler
	Token       storetoken.Token
}

//MessageNotificationMiddleware returns an AuthMiddleware
func MessageNotificationMiddleware(params MessageNotificationParams) *MessageNotification {
	return &MessageNotification{
		userProfiler: params.UserProfile,
		userToken:    params.Token,
	}
}
