package socket

import "log"

// Package socket provides WebSocket functionality for real-time communication
// It includes room management, client handling, and message broadcasting

// GetOrCreateRoom gets an existing room or creates a new one
// This maintains backward compatibility with your existing code

// GetOrCreateRoom gets an existing room or creates a new one
func GetOrCreateRoom(name string) *Room {
	roomsLock.Lock()
	defer roomsLock.Unlock()

	if room, exists := rooms[name]; exists {
		log.Printf("Returning existing room: %s", name)
		return room
	}

	room := NewRoom(name)
	rooms[name] = room
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Room %s panicked: %v", name, r)
			}
		}()
		room.Run()
	}()
	log.Printf("Created new room: %s", name)
	return room
}

// RemoveRoom removes a room from the global rooms map
func RemoveRoom(name string) {
	roomsLock.Lock()
	defer roomsLock.Unlock()

	if room, exists := rooms[name]; exists {
		room.Close()
		delete(rooms, name)
		log.Printf("Removed room: %s", name)
	}
}

// GetRoomStats returns statistics for all rooms
func GetRoomStats() map[string]int {
	roomsLock.Lock()
	defer roomsLock.Unlock()

	stats := make(map[string]int)
	for name, room := range rooms {
		stats[name] = room.ClientCount()
	}

	return stats
}
