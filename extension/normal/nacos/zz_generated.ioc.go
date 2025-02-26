//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli, run 'iocli gen' to re-generate

package nacos

import (
	autowire "github.com/alibaba/ioc-golang/autowire"
	normal "github.com/alibaba/ioc-golang/autowire/normal"
	util "github.com/alibaba/ioc-golang/autowire/util"
	"github.com/nacos-group/nacos-sdk-go/v2/model"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

func init() {
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &impl_{}
		},
	})
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &Impl{}
		},
		ParamFactory: func() interface{} {
			var _ configInterface = &Config{}
			return &Config{}
		},
		ConstructFunc: func(i interface{}, p interface{}) (interface{}, error) {
			param := p.(configInterface)
			impl := i.(*Impl)
			return param.New(impl)
		},
	})
}

type configInterface interface {
	New(impl *Impl) (*Impl, error)
}
type impl_ struct {
	GetConfig_                func(param vo.ConfigParam) (string, error)
	PublishConfig_            func(param vo.ConfigParam) (bool, error)
	DeleteConfig_             func(param vo.ConfigParam) (bool, error)
	ListenConfig_             func(params vo.ConfigParam) (err error)
	CancelListenConfig_       func(params vo.ConfigParam) (err error)
	SearchConfig_             func(param vo.SearchConfigParm) (*model.ConfigPage, error)
	RegisterInstance_         func(param vo.RegisterInstanceParam) (bool, error)
	DeregisterInstance_       func(param vo.DeregisterInstanceParam) (bool, error)
	UpdateInstance_           func(param vo.UpdateInstanceParam) (bool, error)
	GetService_               func(param vo.GetServiceParam) (service model.Service, err error)
	GetAllServicesInfo_       func(param vo.GetAllServiceInfoParam) (model.ServiceList, error)
	SelectAllInstances_       func(param vo.SelectAllInstancesParam) ([]model.Instance, error)
	SelectInstances_          func(param vo.SelectInstancesParam) ([]model.Instance, error)
	SelectOneHealthyInstance_ func(param vo.SelectOneHealthInstanceParam) (*model.Instance, error)
	Subscribe_                func(param *vo.SubscribeParam) error
	Unsubscribe_              func(param *vo.SubscribeParam) (err error)
}

func (i *impl_) GetConfig(param vo.ConfigParam) (string, error) {
	return i.GetConfig_(param)
}

func (i *impl_) PublishConfig(param vo.ConfigParam) (bool, error) {
	return i.PublishConfig_(param)
}

func (i *impl_) DeleteConfig(param vo.ConfigParam) (bool, error) {
	return i.DeleteConfig_(param)
}

func (i *impl_) ListenConfig(params vo.ConfigParam) (err error) {
	return i.ListenConfig_(params)
}

func (i *impl_) CancelListenConfig(params vo.ConfigParam) (err error) {
	return i.CancelListenConfig_(params)
}

func (i *impl_) SearchConfig(param vo.SearchConfigParm) (*model.ConfigPage, error) {
	return i.SearchConfig_(param)
}

func (i *impl_) RegisterInstance(param vo.RegisterInstanceParam) (bool, error) {
	return i.RegisterInstance_(param)
}

func (i *impl_) DeregisterInstance(param vo.DeregisterInstanceParam) (bool, error) {
	return i.DeregisterInstance_(param)
}

func (i *impl_) UpdateInstance(param vo.UpdateInstanceParam) (bool, error) {
	return i.UpdateInstance_(param)
}

func (i *impl_) GetService(param vo.GetServiceParam) (service model.Service, err error) {
	return i.GetService_(param)
}

func (i *impl_) GetAllServicesInfo(param vo.GetAllServiceInfoParam) (model.ServiceList, error) {
	return i.GetAllServicesInfo_(param)
}

func (i *impl_) SelectAllInstances(param vo.SelectAllInstancesParam) ([]model.Instance, error) {
	return i.SelectAllInstances_(param)
}

func (i *impl_) SelectInstances(param vo.SelectInstancesParam) ([]model.Instance, error) {
	return i.SelectInstances_(param)
}

func (i *impl_) SelectOneHealthyInstance(param vo.SelectOneHealthInstanceParam) (*model.Instance, error) {
	return i.SelectOneHealthyInstance_(param)
}

func (i *impl_) Subscribe(param *vo.SubscribeParam) error {
	return i.Subscribe_(param)
}

func (i *impl_) Unsubscribe(param *vo.SubscribeParam) (err error) {
	return i.Unsubscribe_(param)
}

type ImplIOCInterface interface {
	GetConfig(param vo.ConfigParam) (string, error)
	PublishConfig(param vo.ConfigParam) (bool, error)
	DeleteConfig(param vo.ConfigParam) (bool, error)
	ListenConfig(params vo.ConfigParam) (err error)
	CancelListenConfig(params vo.ConfigParam) (err error)
	SearchConfig(param vo.SearchConfigParm) (*model.ConfigPage, error)
	RegisterInstance(param vo.RegisterInstanceParam) (bool, error)
	DeregisterInstance(param vo.DeregisterInstanceParam) (bool, error)
	UpdateInstance(param vo.UpdateInstanceParam) (bool, error)
	GetService(param vo.GetServiceParam) (service model.Service, err error)
	GetAllServicesInfo(param vo.GetAllServiceInfoParam) (model.ServiceList, error)
	SelectAllInstances(param vo.SelectAllInstancesParam) ([]model.Instance, error)
	SelectInstances(param vo.SelectInstancesParam) ([]model.Instance, error)
	SelectOneHealthyInstance(param vo.SelectOneHealthInstanceParam) (*model.Instance, error)
	Subscribe(param *vo.SubscribeParam) error
	Unsubscribe(param *vo.SubscribeParam) (err error)
}

func GetImpl(p *Config) (*Impl, error) {
	i, err := normal.GetImpl(util.GetSDIDByStructPtr(new(Impl)), p)
	if err != nil {
		return nil, err
	}
	impl := i.(*Impl)
	return impl, nil
}

func GetImplIOCInterface(p *Config) (ImplIOCInterface, error) {
	i, err := normal.GetImplWithProxy(util.GetSDIDByStructPtr(new(Impl)), p)
	if err != nil {
		return nil, err
	}
	impl := i.(ImplIOCInterface)
	return impl, nil
}
