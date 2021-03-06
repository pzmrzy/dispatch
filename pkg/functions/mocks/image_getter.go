// Code generated by mockery v1.0.0

// CLOSE THIS FILE AS QUICKLY AS POSSIBLE

package mocks

import context "context"

import mock "github.com/stretchr/testify/mock"
import v1 "github.com/vmware/dispatch/pkg/api/v1"

// ImageGetter is an autogenerated mock type for the ImageGetter type
type ImageGetter struct {
	mock.Mock
}

// GetImage provides a mock function with given fields: ctx, organizationID, imageName
func (_m *ImageGetter) GetImage(ctx context.Context, organizationID string, imageName string) (*v1.Image, error) {
	ret := _m.Called(ctx, organizationID, imageName)

	var r0 *v1.Image
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *v1.Image); ok {
		r0 = rf(ctx, organizationID, imageName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Image)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, organizationID, imageName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
