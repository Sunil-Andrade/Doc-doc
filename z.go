package main

// type Operation struct {
// 	ClientID string `json:"clientId"`
// 	Type     string `json:"type"`
// 	Text     string `json:"text"`
// 	Sequence int64  `json:"sequence"`
// }

// var upgrader = websocket.Upgrader{
// 	CheckOrigin: func(r *http.Request) bool {
// 		return true
// 	},
// }

// var (
// 	clients   = make(map[*websocket.Conn]bool)
// 	clientsMu sync.Mutex

// 	sequence int64
// )

// func doc(db *pgx.Conn) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {

// 		conn, err := upgrader.Upgrade(w, r, nil)
// 		if err != nil {
// 			log.Println("upgrade error:", err)
// 			return
// 		}

// 		defer conn.Close()

// 		clientsMu.Lock()
// 		clients[conn] = true
// 		clientsMu.Unlock()

// 		log.Println("client connected")

// 		defer func() {
// 			clientsMu.Lock()
// 			delete(clients, conn)
// 			clientsMu.Unlock()

// 			log.Println("client disconnected")
// 		}()

// 		for {
// 			var op Operation

// 			err := conn.ReadJSON(&op)
// 			if err != nil {
// 				log.Println("read error:", err)
// 				return
// 			}

// 			// Server assigns the operation order.
// 			sequence++
// 			op.Sequence = sequence

// 			// Save operation in PostgreSQL.
// 			_, err = db.Exec(
// 				context.Background(),
// 				`INSERT INTO operations
// 				(client_id, operation_type, text, sequence)
// 				VALUES ($1, $2, $3, $4)`,
// 				op.ClientID,
// 				op.Type,
// 				op.Text,
// 				op.Sequence,
// 			)

// 			if err != nil {
// 				log.Println("database insert error:", err)
// 				continue
// 			}

// 			// Convert operation back to JSON.
// 			data, err := json.Marshal(op)
// 			if err != nil {
// 				log.Println("marshal error:", err)
// 				continue
// 			}

// 			// Send operation to all connected clients.
// 			broadcast(data)
// 		}
// 	}
// }

// func broadcast(data []byte) {
// 	clientsMu.Lock()
// 	defer clientsMu.Unlock()

// 	for conn := range clients {
// 		err := conn.WriteMessage(
// 			websocket.TextMessage,
// 			data,
// 		)

// 		if err != nil {
// 			log.Println("write error:", err)
// 		}
// 	}
// }
