package proto

type HandlerGetHistoryReq struct {
	StartTime int `json:"start_time"`
	EndTime   int `json:"end_time"`
	UserId    int `json:"user_id"`
	Offset    int `json:"offset"`
	Limit     int `json:"limit"`
}

type HandlerCreateRecordReq struct {
	UserId    string  `json:"user_id"`
	Type      int     `json:"type"`
	Name      string  `json:"name"`
	BuyPrice  float64 `json:"buy_price"`
	WantPrice float64 `json:"want_price"`
	Status    int     `json:"status"` // 1 未完成, 2 已完成
}
