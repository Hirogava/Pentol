package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool { return true },
}

var n int

type MessageHistory struct {
	mu sync.RWMutex
	messages []string
}

func handleWS(ctx context.Context, mh *MessageHistory, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil { log.Println("upgrade:", err); return }
	defer conn.Close()

	conn.SetReadLimit(512)
	conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	conn.SetPongHandler(func(string) error {
		conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})

	pingTicker := time.NewTicker(50 * time.Second)
	defer pingTicker.Stop()

	history := getMessageHistory(mh)
	for _, msg := range history {
		err := conn.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			log.Println("write:", err)
			return
		}
	}

	for {
		select {
		case <-ctx.Done():
            log.Println("closing connection from global shutdown")
            return
		case <-r.Context().Done():
			return
		default:
			mt, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			n++
			log.Printf("%d: %s",n, msg)
			saveMessage(string(msg), mh)

			if err = conn.WriteMessage(mt, msg); err != nil {
				log.Println("write:", err)
				return
			}
		}

		select {
		case <-pingTicker.C:
			conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		default:
		}
	}
}

func saveMessage(msg string, mh *MessageHistory) {
	mh.mu.Lock()
	defer mh.mu.Unlock()

	if len(mh.messages) >= 10 {
		mh.messages = mh.messages[1:]
	}

	mh.messages = append(mh.messages, msg)
}

func getMessageHistory(mh *MessageHistory) []string {
	mh.mu.RLock()
	defer mh.mu.RUnlock()
	history := make([]string, len(mh.messages))
	copy(history, mh.messages)
	return history
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c
		log.Println("Shutting down gracefully...")
		cancel()
	}()

	var mh MessageHistory
	
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handleWS(ctx, &mh, w, r)
	})

	srv := &http.Server{Addr: ":8080"}

	go func() {
		log.Println("📡 ws://localhost:8080/ws")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe: %v", err)
		}
	}()

	<-ctx.Done()

	log.Println("Graceful shutdown initiated...")
	shutdownCtx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Println("Server gracefully stopped")
}