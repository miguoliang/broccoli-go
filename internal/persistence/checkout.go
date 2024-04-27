package persistence

type Checkout struct {
	Model
	UserID         string `json:"user_id"`
	SubscriptionID string `json:"subscription_id"`
	IdempotencyKey string `json:"idempotency_key"`
}
