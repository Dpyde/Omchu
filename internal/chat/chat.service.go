package chat

type ChatService interface {
	GetChat(userId string) ([]ExtendedChat, error)
}

type chatServiceImpl struct {
	repo ChatRepository
}

func NewChatService(repo ChatRepository) ChatService {
	return &chatServiceImpl{repo: repo}
}

func (s *chatServiceImpl) GetChat(userId string) ([]ExtendedChat, error) {
	chats, err := s.repo.FindById(userId)
	if err != nil {
		return []ExtendedChat{}, err
	}
	return chats, nil
}
