package service

type VideoService struct {
}

func NewVideoService() *VideoService {
	return &VideoService{}
}

func (s *VideoService) ProcessVideo(inputPath string) error {
	// TODO: 实现视频处理逻辑
	return nil
}
