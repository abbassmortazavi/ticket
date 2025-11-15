package socket

import "sync"

type Manager struct {
	rooms map[string]*Room
	lock  sync.RWMutex
}

func NewManager() *Manager {
	return &Manager{
		rooms: make(map[string]*Room),
	}
}

// GetRoom returns a room by name, creates if it doesn't exist
func (m *Manager) GetRoom(name string) *Room {
	m.lock.RLock()
	defer m.lock.RUnlock()

	if room, exists := m.rooms[name]; exists {
		return room
	}
	room := NewRoom(name)
	m.rooms[name] = room
	go room.Run()
	return room
}
func (m *Manager) RemoveRoom(name string) {
	m.lock.Lock()
	defer m.lock.Unlock()
	if room, exists := m.rooms[name]; exists {
		room.Close()
		delete(m.rooms, name)
	}
}

// GetRoomStats returns room statistics
func (m *Manager) GetRoomStats() map[string]int {
	m.lock.RLock()
	defer m.lock.RUnlock()
	stats := make(map[string]int)
	for name, room := range m.rooms {
		stats[name] = room.ClientCount()
	}
	return stats

}

// GetAllRooms returns all room names
func (m *Manager) GetAllRooms() []string {
	m.lock.RLock()
	defer m.lock.RUnlock()
	rooms := make([]string, 0, len(m.rooms))
	for room := range m.rooms {
		rooms = append(rooms, room)
	}
	return rooms
}
