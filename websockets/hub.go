// Package websockets gestiona las conexiones WebSocket múltiples y la comunicación
// entre el servidor y todos los clientes conectados.
package websockets

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// upgrader es una instancia global que permite convertir conexiones HTTP regulares
// en conexiones WebSocket. CheckOrigin permite conexiones desde cualquier origen
// (en producción deberías validar los orígenes permitidos por seguridad).
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Permitir conexiones desde cualquier origen
	},
}

// Hub es el centro de comunicaciones que gestiona todos los clientes WebSocket conectados.
// Coordina el registro, desregistro y la comunicación entre clientes.
type Hub struct {
	clients    []*Client    // Lista de todos los clientes actualmente conectados
	register   chan *Client // Canal para registrar nuevos clientes que se conectan
	unregister chan *Client // Canal para desregistrar clientes que se desconectan
	mutex      *sync.Mutex  // Mutex para proteger el acceso concurrente a la lista de clientes
}

// NewHub crea una nueva instancia de Hub.
// Inicializa todos los campos necesarios:
// - Un slice vacío para clientes
// - Canales para registro y desregistro de clientes
// - Un mutex para proteger el acceso concurrente
func NewHub() *Hub {
	return &Hub{
		clients:    make([]*Client, 0), // Crea un slice vacío de clientes
		register:   make(chan *Client), // Canal para registrar nuevos clientes
		unregister: make(chan *Client), // Canal para desregistrar clientes
		mutex:      &sync.Mutex{},      // Mutex para protección de concurrencia
	}
}

// WebSocketHandler maneja nuevas conexiones WebSocket entrantes.
// Pasos que realiza:
// 1. Convierte la conexión HTTP a WebSocket usando el upgrader
// 2. Crea un nuevo cliente para esa conexión
// 3. Registra el cliente en el Hub
// 4. Inicia una goroutine para manejar los mensajes del cliente
func (h *Hub) WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	// Intenta convertir la conexión HTTP a WebSocket
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Could not upgrade websocket connection", http.StatusInternalServerError)
		return
	}

	// Crea un nuevo cliente con la conexión WebSocket
	client := NewClient(h, socket)
	// Envía el cliente al canal de registro para que el Hub lo añada a la lista
	h.register <- client

	// Inicia una goroutine para manejar el envío de mensajes a este cliente
	go client.Write()
}

// Run es el bucle principal del Hub que maneja el registro y desregistro de clientes.
// Utiliza un select statement para escuchar en ambos canales de forma no bloqueante.
// Este método debe ejecutarse en una goroutine separada ya que es un bucle infinito.
func (h *Hub) Run() {
	log.Println("🌀 WebSocket Hub is running...")
	for {
		select {
		// Cuando llega un nuevo cliente para registrar
		case client := <-h.register:
			log.Println("➕ Nuevo cliente registrado:", client.socket.RemoteAddr())
			h.onConnect(client) // Llama al método para manejar la conexión del cliente
		// Cuando llega un cliente para desregistrar
		case client := <-h.unregister:
			log.Println("💥 Cliente desregistrado:", client.socket.RemoteAddr())
			h.onDisconnect(client) // Llama al método para manejar la desconexión del cliente
		}
	}
}

// onConnect maneja la lógica cuando un nuevo cliente se conecta al Hub.
// Registra el cliente en la lista de clientes conectados y puede realizar otras
// acciones como enviar un mensaje de bienvenida.
func (h *Hub) onConnect(client *Client) {
	h.mutex.Lock()         // Bloquea para acceso exclusivo
	defer h.mutex.Unlock() // Asegura que se desbloquee al final

	// Añade el nuevo cliente a la lista
	client.id = client.socket.RemoteAddr().String()
	h.clients = append(h.clients, client)
}

// onDisconnect maneja la lógica cuando un cliente se desconecta del Hub.
// Elimina el cliente de la lista de clientes conectados y puede realizar otras
// acciones como notificar a otros clientes.
func (h *Hub) onDisconnect(client *Client) {
	h.mutex.Lock()         // Bloquea para acceso exclusivo
	defer h.mutex.Unlock() // Asegura que se desbloquee al final

	// Busca el cliente en la lista y lo elimina
	for i, c := range h.clients {
		if c == client {
			// Elimina el cliente usando slicing de Go
			h.clients = append(h.clients[:i], h.clients[i+1:]...)
			break
		}
	}
}

func (h *Hub) SendMessageToClients(message any, ignore *Client) {
	jsonMessage, _ := json.Marshal(message)

	for _, client := range h.clients {
		if client != ignore { // No enviar al cliente que envió el mensaje
			client.outbound <- jsonMessage // Enviar el mensaje al canal outbound del cliente
		}
	}
}
