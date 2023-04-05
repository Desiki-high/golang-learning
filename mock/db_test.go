package mock

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang-learning/mock/mocks"
	"testing"
)

func TestGetFromDB(t *testing.T) {
	tests := []struct {
		name   string                            // 测试的名称
		input  string                            // 测试输入
		mock   func(m *mocks.MockDBMockRecorder) // 桩函数
		expect func(t *testing.T, value int)     // 测试过程
	}{
		{
			name:  "nil input",
			input: "",
			mock: func(m *mocks.MockDBMockRecorder) {
				m.Get(gomock.Eq("")).Return(0, errors.New("input nil")).Times(1)
			},
			expect: func(t *testing.T, value int) {
				assertions := assert.New(t)
				assertions.Equal(-1, value)
			},
		},
		{
			name:  "successful input",
			input: "Tom",
			mock: func(m *mocks.MockDBMockRecorder) {
				m.Get(gomock.Eq("Tom")).Return(100, nil).Times(1)
			},
			expect: func(t *testing.T, value int) {
				assertions := assert.New(t)
				assertions.Equal(100, value)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mocks.NewMockDB(ctrl)
			tc.mock(m.EXPECT())
			value := GetFromDB(m, tc.input)
			tc.expect(t, value)
		})
	}

	//ctrl := gomock.NewController(t)
	//defer ctrl.Finish()
	//
	//m := mocks.NewMockDB(ctrl)
	//
	//m.EXPECT().Get(gomock.Eq("Tom")).Return(100, errors.New("not exist"))

	// Does not make any assertions. Executes the anonymous functions and returns
	// its result when Bar is invoked with 99.
	//m.
	//	EXPECT().
	//	Get(gomock.Eq("Tom")).
	//	DoAndReturn(func(_ string) (int, error) {
	//		time.Sleep(1 * time.Second)
	//		return 101, errors.New("test")
	//	}).
	//	AnyTimes()

	// Does not make any assertions. Not "" return -10, err and call times is 1.
	//m.
	//	EXPECT().
	//	Get(gomock.Not("")).
	//	Return(-10, errors.New("nil")).
	//	Times(1)

	//if v := GetFromDB(m, "Tom"); v != -1 {
	//	t.Fatal("expected -1, but got", v)
	//}
}
