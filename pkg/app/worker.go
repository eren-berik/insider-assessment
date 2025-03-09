package app

import (
	"context"
	"insider-assessment/pkg/domain/message"
	"log"
	"sync"
	"time"
)

type Worker struct {
	postgres  message.PostgresService
	cache     message.CacheService
	sender    *MessageSender
	stopChan  chan struct{}
	runningMu sync.Mutex
	running   bool
	config    OptionProvider
}

func NewWorker(service *ServiceProvider, config *OptionProvider) *Worker {
	return &Worker{
		postgres: service.PostgresService,
		cache:    service.CacheService,
		sender:   &service.MessageSenderService,
		stopChan: make(chan struct{}),
		config:   *config,
	}
}

func (w *Worker) Start() {
	w.runningMu.Lock()
	if w.running {
		w.runningMu.Unlock()
		log.Println("Worker is already running")
		return
	}
	w.running = true
	w.stopChan = make(chan struct{})
	w.runningMu.Unlock()

	ctx := context.Background()
	ticker := time.NewTicker(w.config.TriggerTime)
	defer ticker.Stop()

	log.Println("Worker started")

	for {
		select {
		case <-ticker.C:
			log.Println("Fetching unsent messages...")
			messages, err := w.postgres.GetMessagesByStatus(ctx, message.Pending, w.config.BatchSize)
			if err != nil {
				log.Println("Error fetching messages:", err)
				continue
			}

			if len(messages) == 0 {
				log.Println("No messages found")
			}

			go w.processMessages(messages)

		case <-w.stopChan:
			log.Println("Worker stopped")
			w.runningMu.Lock()
			w.running = false
			w.runningMu.Unlock()
			return
		}
	}
}

func (w *Worker) Stop() {
	w.runningMu.Lock()
	defer w.runningMu.Unlock()

	if !w.running {
		log.Println("Worker is not running")
		return
	}

	close(w.stopChan)
	w.running = false
	log.Println("Worker stopping...")
}

func (w *Worker) processMessages(msgs []*message.Message) {
	ctx := context.Background()

	for _, msg := range msgs {
		success, res := w.sender.SendMessage(*msg)

		if success {
			err := w.postgres.UpdateMessageStatus(ctx, msg.ID(), message.Sent)
			if err != nil {
				log.Println("Error updating message status:", err)
			} else {
				log.Printf("Message %d sent successfully!\n", msg.ID())
			}

			cacheMsg := message.NewMessageCache(res.MessageId, time.Now())

			err = w.cache.Save(ctx, cacheMsg)
			if err != nil {
				log.Println("Error saving to cache:", err)
			}
		} else {
			err := w.postgres.UpdateMessageStatus(ctx, msg.ID(), message.Failed)
			if err != nil {
				log.Println("Error updating message status:", err)
			} else {
				log.Printf("Message %d sent failed!\n", msg.ID())
			}

			cacheMsg := message.NewMessageCache(res.MessageId, time.Now())

			err = w.cache.Save(ctx, cacheMsg)
			if err != nil {
				log.Println("Error saving to cache:", err)
			}
		}
	}
}
