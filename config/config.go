package config

const (
	MinioEndpoint  = "localhost:9000"
	MinioAccessKey = "minioadmin"
	MinioSecretKey = "minioadmin"
	BucketName     = "documents"

	SegmentSize  = 50 * 1024 // 50 Кбайт
	TransportURL = "http://10.205.157.61:8080/receive"
	ServerPort   = ":8080"
)
