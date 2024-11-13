package tglogger

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/wmuga/tglogger/internal/models"
)

type tgLogger struct {
	notImplementedLogger
	topic string
	chats map[int64]struct{}
	bot   *tgbotapi.BotAPI

	mx *sync.RWMutex
}

func (t *tgLogger) AddChat(chatID int64) error {
	t.mx.Lock()
	defer t.mx.Unlock()
	t.chats[chatID] = struct{}{}
	return t.printToChat(chatID, fmt.Sprint("Connected to ", chatID))
}

func (t *tgLogger) RemoveChat(chatID int64) error {
	t.mx.Lock()
	defer t.mx.Unlock()
	delete(t.chats, chatID)
	return nil
}

func (t *tgLogger) Fatal(message string, fields ...models.Field) {
	err := t.Print("error", message, fields...)
	if err != nil {
		log.Println(err)
	}
	os.Exit(1)
}

func (t *tgLogger) Error(message string, fields ...models.Field) error {
	return t.Print("error", message, fields...)
}

func (t *tgLogger) Info(message string, fields ...models.Field) error {
	return t.Print("info", message, fields...)
}

func (t *tgLogger) Print(level string, message string, fields ...models.Field) error {
	msg := strings.Builder{}
	msg.WriteString(
		fmt.Sprint("Topic: ", t.topic, "\nTime: ", time.Now().UTC(),
			"\nLevel: ", level, "\nMessage: ", message))

	for _, f := range fields {
		msg.WriteRune('\n')
		msg.WriteString(f.String())
	}

	msgStr := msg.String()

	t.mx.RLock()
	defer t.mx.RUnlock()
	for k := range t.chats {
		if err := t.printToChat(k, msgStr); err != nil {
			return err
		}
	}
	return nil
}

func (t *tgLogger) printToChat(chatID int64, message string) error {
	msg := tgbotapi.NewMessage(chatID, message)
	_, err := t.bot.Send(msg)
	if err == nil {
		return nil
	}
	return fmt.Errorf("can't send send message to chat: %d err: %w", chatID, err)
}

func NewTgLogger(ctx context.Context, topic, token string, chatIDs ...int64) (models.TgLogger, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	logger := &tgLogger{
		notImplementedLogger: notImplementedLogger{},
		chats:                map[int64]struct{}{},
		topic:                topic,
		bot:                  bot,
		mx:                   &sync.RWMutex{},
	}

	for _, chat := range chatIDs {
		logger.chats[chat] = struct{}{}
	}

	_, err = bot.Request(tgbotapi.DeleteWebhookConfig{DropPendingUpdates: true})
	if err != nil {
		return nil, err
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	upds := bot.GetUpdatesChan(u)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case upd := <-upds:
				// Только сообщения
				if upd.Message == nil {
					continue
				}

				switch upd.Message.Text {
				case "/start":
					if err := logger.AddChat(upd.Message.Chat.ID); err != nil {
						fmt.Println(err)
					}
				case "/end":
					if err := logger.RemoveChat(upd.Message.Chat.ID); err != nil {
						fmt.Println(err)
					}
				}
			}
		}
	}()

	return logger, nil
}
