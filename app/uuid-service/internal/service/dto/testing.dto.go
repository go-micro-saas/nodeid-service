package dto

import "google.golang.org/protobuf/types/known/emptypb"

var (
	TestdataDto testdata
)

type testdata struct{}

func (s *testdata) ToBoXxx() interface{} {
	return emptypb.Empty{}
}
