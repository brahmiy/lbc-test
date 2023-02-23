package services

type IStatsService interface {
	Hit(request []byte)
	GetTopRequest() ([]byte, int64)
}
