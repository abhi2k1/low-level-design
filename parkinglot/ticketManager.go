package main

type TicketManager interface {
	IssueTicket(vehicle Vehicle, parkingSpot ParkingSpot) Ticket
	GetTicket(id int) Ticket
	RemoveTicket(id int) bool
}

type ticketManager struct {
	tickets []Ticket
}

func NewTicketManager() *ticketManager {
	return &ticketManager{}
}

func (t *ticketManager) IssueTicket(vehicle Vehicle, parkingSpot ParkingSpot) Ticket {
	ticket := NewTicket(len(t.tickets), vehicle, parkingSpot)
	t.tickets = append(t.tickets, ticket)
	return ticket
}

func (t *ticketManager) GetTicket(id int) Ticket {
	if id >= len(t.tickets) {
		return nil
	}

	return t.tickets[id]
}

func (t *ticketManager) RemoveTicket(id int) bool {
	if id >= len(t.tickets) {
		return false
	}
	t.tickets[id] = nil
	return true
}
