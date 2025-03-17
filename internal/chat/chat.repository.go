package chat

type ChatRepository interface {
	FindById(userId string) ([]ExtendedChat, error)
}
