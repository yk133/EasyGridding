package proto

type HandlerGetHistoryReq struct {
	DateStart int `json:"date_start"`
	DateEnd   int `json:"date_end"`
	UserId    int `json:"user_id"`
}
