package stat

import (
	"log"

	"github.com/user-xat/short-link/pkg/di"
	"github.com/user-xat/short-link/pkg/event"
)

type StatServiceDeps struct {
	EventBus       *event.EventBus
	StatRepository di.IStatRepository
}

type StatService struct {
	EventBus       *event.EventBus
	StatRepository di.IStatRepository
}

func NewStatService(deps StatServiceDeps) *StatService {
	return &StatService{
		EventBus:       deps.EventBus,
		StatRepository: deps.StatRepository,
	}
}

func (s *StatService) AddClick() {
	for msg := range s.EventBus.Subscribe() {
		if msg.Type == event.EventLinkVisited {
			id, ok := msg.Data.(uint)
			if !ok {
				log.Fatalln("Bad EventLinkVisited Data: ", msg.Data)
				continue
			}
			s.StatRepository.AddClick(id)
		}
	}
}
