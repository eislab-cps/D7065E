package websocket

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/eislab-cps/buildingsim/pkg/store"
	ws "github.com/gorilla/websocket"
)

type Message struct {
	Type    string      `json:"type"`
	Data    interface{} `json:"data,omitempty"`
	Version int64       `json:"version,omitempty"`
}

type client struct {
	conn      *ws.Conn
	sessionID string
	send      chan []byte
}

type Hub struct {
	mu       sync.RWMutex
	clients  map[*client]bool
	sessions map[string]map[*client]bool // session id -> clients
	store    *store.MemoryStore
}

func Upgrader() ws.Upgrader {
	return ws.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
}

func NewHub(store *store.MemoryStore) *Hub {
	return &Hub{
		clients:  make(map[*client]bool),
		sessions: make(map[string]map[*client]bool),
		store:    store,
	}
}

func (h *Hub) HandleConnection(conn *ws.Conn, sessionID string) {
	c := &client{
		conn:      conn,
		sessionID: sessionID,
		send:      make(chan []byte, 256),
	}

	h.mu.Lock()
	h.clients[c] = true
	if h.sessions[sessionID] == nil {
		h.sessions[sessionID] = make(map[*client]bool)
	}
	h.sessions[sessionID][c] = true
	h.mu.Unlock()

	h.store.TouchSession(sessionID)

	go h.writePump(c)
	h.readPump(c)
}

func (h *Hub) readPump(c *client) {
	defer func() {
		h.removeClient(c)
		c.conn.Close()
	}()

	c.conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		h.store.TouchSession(c.sessionID)
		return nil
	})

	for {
		_, _, err := c.conn.ReadMessage()
		if err != nil {
			break
		}
		h.store.TouchSession(c.sessionID)
	}
}

func (h *Hub) writePump(c *client) {
	ticker := time.NewTicker(30 * time.Second)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case msg, ok := <-c.send:
			if !ok {
				c.conn.WriteMessage(ws.CloseMessage, []byte{})
				return
			}
			c.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if err := c.conn.WriteMessage(ws.TextMessage, msg); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if err := c.conn.WriteMessage(ws.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (h *Hub) removeClient(c *client) {
	h.mu.Lock()
	defer h.mu.Unlock()

	delete(h.clients, c)
	if clients, ok := h.sessions[c.sessionID]; ok {
		delete(clients, c)
		if len(clients) == 0 {
			delete(h.sessions, c.sessionID)
		}
	}
	close(c.send)
}

func (h *Hub) SendToSession(sessionID string, msg Message) {
	data, err := json.Marshal(msg)
	if err != nil {
		log.Printf("Failed to marshal WS message: %v", err)
		return
	}

	h.mu.RLock()
	clients := h.sessions[sessionID]
	h.mu.RUnlock()

	for c := range clients {
		select {
		case c.send <- data:
		default:
			// Client buffer full, skip
		}
	}
}

func (h *Hub) BroadcastToAll(msg Message) {
	data, err := json.Marshal(msg)
	if err != nil {
		log.Printf("Failed to marshal WS message: %v", err)
		return
	}

	h.mu.RLock()
	defer h.mu.RUnlock()

	for c := range h.clients {
		select {
		case c.send <- data:
		default:
		}
	}
}

func (h *Hub) StartPurger(interval time.Duration, maxAge time.Duration) {
	go func() {
		ticker := time.NewTicker(interval)
		for range ticker.C {
			purged := h.store.PurgeInactiveSessions(maxAge)
			for _, id := range purged {
				h.mu.Lock()
				if clients, ok := h.sessions[id]; ok {
					for c := range clients {
						c.conn.Close()
					}
					delete(h.sessions, id)
				}
				h.mu.Unlock()
				log.Printf("Purged inactive session: %s", id)
			}
		}
	}()
}
