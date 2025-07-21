// Package websockets maneja la comunicación en tiempo real entre el servidor y los clientes
// utilizando WebSockets para permitir comunicación bidireccional.
package websockets

import "github.com/gorilla/websocket"

// Client representa un cliente conectado al servidor WebSocket.
// Cada cliente tiene una conexión única y un canal para enviar mensajes.
type Client struct {
	hub      *Hub            // Referencia al Hub central que maneja todos los clientes
	id       string          // Identificador único del cliente (actualmente no se usa)
	socket   *websocket.Conn // Conexión WebSocket activa con el cliente
	outbound chan []byte     // Canal para enviar mensajes al cliente de forma asíncrona
}

// NewClient crea una nueva instancia de Client.
// Parámetros:
//   - hub: Referencia al Hub que gestionará este cliente
//   - socket: Conexión WebSocket establecida con el cliente
// Retorna un puntero a la nueva instancia de Client
func NewClient(hub *Hub, socket *websocket.Conn) *Client {
	return &Client{
		hub:      hub,
		socket:   socket,
		outbound: make(chan []byte), // Crea un canal buffered para mensajes salientes
	}
}

// Write es el método principal que maneja el envío de mensajes al cliente.
// Ejecuta en una goroutine separada y escucha continuamente el canal outbound.
// Cuando recibe un mensaje, lo envía al cliente a través de la conexión WebSocket.
// Cuando el canal se cierra (no hay más mensajes), envía un mensaje de cierre.
func (c *Client) Write() {
	// Itera sobre todos los mensajes que llegan al canal outbound
	for message := range c.outbound {
		// Envía cada mensaje como texto al cliente WebSocket
		c.socket.WriteMessage(websocket.TextMessage, message)
	}
	// Cuando el canal se cierra, notifica al cliente que la conexión terminará
	c.socket.WriteMessage(websocket.CloseMessage, []byte{})
}
