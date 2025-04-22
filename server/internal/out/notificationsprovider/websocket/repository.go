package websocket

type NotificationsProvider struct{}

func New() *NotificationsProvider {
	return &NotificationsProvider{}
}
