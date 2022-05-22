package CommModels

import (
	"github.com/beego/beego/v2/core/validation"
	"time"
)

type BaseModel struct {
	Id        int       `json:"id" orm:"auto"`
	CreatedAt time.Time `json:"createdAt" orm:"auto_now_add;"`
	UpdatedAt time.Time `json:"updatedAt" orm:"auto_now_update;type(datetime)"`
	DeletedAt time.Time `json:"deletedAt" orm:"auto_now_delete;type(datetime)"`
}

// Valid 自定义校验
func (base *BaseModel) Valid(v *validation.Validation) {

}
