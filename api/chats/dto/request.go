package dto

type Chats struct {
	CustomerUUID string `json:"customer_uuid,omitempty" validate:"required,uuid4"`
	AgentUUID    string `json:"agent_uuid,omitempty" validate:"required,uuid4"`
	SenderName   string `json:"sender_name,omitempty"`
	Message      string `json:"message,omitempty" validate:"required"`
	IsCSChat     bool   `json:"is_cs_chat,omitempty"`
}

type Type string

const (
	Agent    Type = "agent"
	Customer Type = "customer"
)

type ChatSender struct {
	UUID  string `json:"uuid,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Type  Type   `json:"type,omitempty"`
}
