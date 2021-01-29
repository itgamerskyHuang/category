package handler

import (
	"context"

	pb "github.com/itgamerskyHuang/category/proto/category"
)

type Category struct{}

func (c *Category) CreateCategory(context.Context, *pb.CreateRequest, *pb.CreateResponse) error {

}
func (c *Category) UpdataCategory(context.Context, *pb.UpdataRequest, *pb.UpdataResponse) error {

}
func (c *Category) DelectCategory(context.Context, *pb.DelectRequest, *pb.DelectResponse) error {

}
func (c *Category) FindCategoryByName(context.Context, *pb.FindByNameRequest, *pb.PropertyResponse) error {

}
func (c *Category) FindCategoryByID(context.Context, *pb.FindByIDRequest, *pb.PropertyResponse) error {

}
func (c *Category) FindCategoryByLevel(context.Context, *pb.FindByLevelRequest, *pb.PropertyResponse) error {

}
func (c *Category) FindCategoryByParent(context.Context, *pb.FindByParentRequest, *pb.PropertyResponse) error {

}
func (c *Category) FindAllCategory(context.Context, *pb.FindAllRequest, *pb.FindAllResponse) error {

}
