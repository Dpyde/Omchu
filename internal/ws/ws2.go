package ws

// import (
//     "net/http"
//     //"github.com/Dpyde/Omchu/internal/entity/ws"
//     "github.com/gofiber/fiber/v2"
//     "github.com/gorilla/websocket"
// )

// type Handler struct {
//     hub *Hub
// }

// func NewHandler(h *Hub) *Handler {
//     return &Handler{
//         hub: h,
//     }
// }

// type CreateRoomReq struct {
//     ID   string `json:"id"`
//     Name string `json:"name"`
// }

// func (h *Handler) CreateRoom(c *fiber.Ctx) error {
//     var req CreateRoomReq
//     if err := c.BodyParser(&req); err != nil {
//         return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
//     }

//     h.hub.Rooms[req.ID] = &Room{
//         ID:      req.ID,
//         Name:    req.Name,
//         Clients: make(map[string]*Client),
//     }

//     return c.Status(fiber.StatusOK).JSON(req)
// }

// var upgrader = websocket.Upgrader{
//     ReadBufferSize:  1024,
//     WriteBufferSize: 1024,
//     CheckOrigin: func(r *http.Request) bool {
//         return true
//     },
// }

// func (h *Handler) JoinRoom(c *fiber.Ctx) error {
//     conn, err := upgrader.Upgrade(c.Context().Response.BodyWriter(), c.Context().Request().(*http.Request), nil)
//     if err != nil {
//         return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
//     }

//     roomID := c.Params("roomId")
//     clientID := c.Query("userId")
//     username := c.Query("username")

//     cl := &ws.Client{
//         Conn:     conn,
//         Message:  make(chan *Message, 10),
//         ID:       clientID,
//         RoomID:   roomID,
//         Username: username,
//     }

//     m := &Message{
//         Content:  "A new user has joined the room",
//         RoomID:   roomID,
//         Username: username,
//     }

//     h.hub.Register <- cl
//     h.hub.Broadcast <- m

//     go cl.WriteMessage()
//     cl.ReadMessage(h.hub)

//     return nil
// }

// type RoomRes struct {
//     ID   string `json:"id"`
//     Name string `json:"name"`
// }

// func (h *Handler) GetRooms(c *fiber.Ctx) error {
//     rooms := make([]RoomRes, 0)

//     for _, r := range h.hub.Rooms {
//         rooms = append(rooms, RoomRes{
//             ID:   r.ID,
//             Name: r.Name,
//         })
//     }

//     return c.Status(fiber.StatusOK).JSON(rooms)
// }

// type ClientRes struct {
//     ID       string `json:"id"`
//     Username string `json:"username"`
// }

// func (h *Handler) GetClients(c *fiber.Ctx) error {
//     var clients []ClientRes
//     roomId := c.Params("roomId")

//     if _, ok := h.hub.Rooms[roomId]; !ok {
//         clients = make([]ClientRes, 0)
//         return c.Status(fiber.StatusOK).JSON(clients)
//     }

//     for _, c := range h.hub.Rooms[roomId].Clients {
//         clients = append(clients, ClientRes{
//             ID:       c.ID,
//             Username: c.Username,
//         })
//     }

//     return c.Status(fiber.StatusOK).JSON(clients)
// }