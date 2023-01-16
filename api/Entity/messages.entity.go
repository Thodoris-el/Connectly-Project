package entity

type FacebookMessage struct {
	Object string      `json:"object"`
	Entry  []EntryType `json:"entry"`
}

type EntryType struct {
	ID        string          `json:"id"`
	Time      int64           `json:"time"`
	Messaging []MessagingType `json:"messaging"`
}

type MessagingType struct {
	Sender    SenderType    `json:"sender"`
	Recipient RecipientType `json:"recipient"`
	Timestamp int64         `json:"timestamp"`
	Message   MessageType   `json:"message"`
}

type SenderType struct {
	ID string `json:"id"`
}

type RecipientType struct {
	ID string `json:"id"`
}

type MessageType struct {
	Mid         string           `json:"mid"`
	Text        string           `json:"text,omitempty"`
	Attachments []AttachmentType `json:"attachments,omitempty"`
}

type AttachmentType struct {
	Type    string      `json:"type,omitempty"`
	Payload PayloadType `json:"payload,omitempty"`
}

type PayloadType struct {
	URL     string      `json:"url,omitempty"`
	Product ProductType `json:"product,omitempty"`
	Title   string      `json:"title,omitempty"`
}

type ProductType struct {
	ID          string `json:"id,omitempty"`
	Retailer_ID string `json:"retailer_id,omitempty"`
	Image_URL   string `json:"image_url,omitempty"`
	Title       string `json:"title,omitempty"`
	Subtitle    string `json:"subtitle,omitempty"`
}

type MessangeSent struct {
	Messaging_Type string         `json:"messaging_type"`
	Recipient      RecipientType  `json:"recipient"`
	Message        ResMessageType `json:"message"`
}

type ResMessageType struct {
	Text          string           `json:"text"`
	Quick_Replies []QuickReplyType `json:"quick_replies"`
}

type QuickReplyType struct {
	Content_Type string `json:"content_type"`
	Title        string `json:"title"`
	Payload      string `json:"payload"`
	Image        string `json:"image_url"`
}
