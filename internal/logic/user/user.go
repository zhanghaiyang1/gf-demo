package user

import "context"

type sUser struct{}


func New() *sUser{
	return &sUser{}
}
//生成用户
func (s *sUser) Create(ctx context.Context)(err error){

	return nil
}