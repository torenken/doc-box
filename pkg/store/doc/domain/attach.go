package domain

func NewAttachment(docId, data string) Attachment {
	return Attachment{
		DocId: docId,
		Data:  data,
	}
}

type Attachment struct {
	DocId string `json:"docId"`
	Data  string `json:"data"`
	Url   string `json:"url"`
}
