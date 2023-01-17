package entity

//Entities as returned from FB and as needed to be sent to FB

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
	Sender             SenderType    `json:"sender"`
	Recipient          RecipientType `json:"recipient"`
	Timestamp          int64         `json:"timestamp"`
	Message            MessageType   `json:"message,omitempty"`
	Messaging_Feedback MesFeedType   `json:"messaging_feedback,omitempty"`
}

type SenderType struct {
	ID string `json:"id"`
}

type RecipientType struct {
	ID string `json:"id"`
}

type MessageType struct {
	Mid         string           `json:"mid,omitempty"`
	Text        string           `json:"text,omitempty"`
	Attachments []AttachmentType `json:"attachments,omitempty"`
	QuickReply  QuickReplyType   `json:"quick_reply,omitempty"`
}

type MesFeedType struct {
	FeedbackScreens []FeScType `json:"feedback_screens,omitempty"`
}

type FeScType struct {
	ScreenID  int         `json:"screen_id,omitempty"`
	Questions QuesTypeRes `json:"questions,omitempty"`
}

type QuesTypeRes struct {
	Myquestion1 MyQuestionType `json:"myquestion1,omitempty"`
}

type MyQuestionType struct {
	Type     string          `json:"type,omitempty"`
	Payload  string          `json:"payload,omitempty"`
	FollowUp FollowUpTypeRes `json:"follow_up"`
}

type FollowUpTypeRes struct {
	Type    string `json:"type"`
	Payload string `json:"payload"`
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

type MessangeSentTemplate struct {
	Recipient RecipientType          `json:"recipient"`
	Message   ResMessageTypeTemplate `json:"message"`
}

type ResMessageType struct {
	Text          string           `json:"text,omitempty"`
	Quick_Replies []QuickReplyType `json:"quick_replies,omitempty"`
}

type ResMessageTypeTemplate struct {
	Text       string                     `json:"text,omitempty"`
	Attachment AttachmentTypeSendTemplate `json:"attachment,omitempty"`
}

type QuickReplyType struct {
	Content_Type string `json:"content_type,omitempty"`
	Title        string `json:"title,omitempty"`
	Payload      string `json:"payload,omitempty"`
	Image        string `json:"image_url,omitempty"`
}

type AttachmentTypeSendTemplate struct {
	Type    string                  `json:"type,omitempty"`
	Payload PayloadTypeSendTemplate `json:"payload,omitempty"`
}

type PayloadTypeSendTemplate struct {
	Template_Type    string              `json:"template_type"`
	Title            string              `json:"title"`
	Subtitle         string              `json:"subtitle"`
	Button_Title     string              `json:"button_title"`
	Feedbavk_Screens []FeedbackType      `json:"feedback_screens"`
	Business_Privacy BusinessPrivacyType `json:"business_privacy"`
	Expires_In_Days  int                 `json:"expires_in_days"`
}

type FeedbackType struct {
	Questions []QuestionType `json:"questions"`
}

type BusinessPrivacyType struct {
	Url string `json:"url"`
}

type QuestionType struct {
	ID           string       `json:"id"`
	Type         string       `json:"type"`
	Title        string       `json:"title"`
	Score_Label  string       `json:"score_label"`
	Score_Option string       `json:"score_option"`
	FollowUp     FollowUpType `json:"follow_up"`
}

type FollowUpType struct {
	Type        string `json:"type"`
	Placeholder string `json:"placeholder"`
}
