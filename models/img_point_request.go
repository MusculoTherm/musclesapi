package models

type ImagePointRequest struct {
	Name   string `json:"name"`
	X      int64  `json:"x"`
	Y      int64  `json:"y"`
	Radius int64  `json:"r"`
}

func (req *ImagePointRequest) MirrorToResponse() ImagePointResponse {
	return ImagePointResponse{
		Name:   req.Name,
		X:      req.X,
		Y:      req.Y,
		Radius: req.Radius,
	}
}

func MirrorImagePointRequestsToResponses(reqs []ImagePointRequest) []ImagePointResponse {
	out := make([]ImagePointResponse, len(reqs))
	for i := 0; i < len(out); i++ {
		out[i] = reqs[i].MirrorToResponse()
	}
	return out
}
