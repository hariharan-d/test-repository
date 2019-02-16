package routers

import (
	MyAPIcreateContact "MyAPI/controllers/MyAPI/createContact"
	MyAPIsearchContact "MyAPI/controllers/MyAPI/searchContact"
	MyAPIupdateContact "MyAPI/controllers/MyAPI/updateContact"

	"MyAPI/controllers/error"
	"MyModel/db"

	"github.com/astaxie/beego"
)

func init() {
	if db.Init() != nil {
		return
	}
	beego.ErrorController(&error.Error{})
	beego.Router(beego.AppConfig.String("MyAPI_SEARCH_CONTACT_PATH"), &MyAPIsearchContact.ViewContact{})
	beego.Router(beego.AppConfig.String("MyAPI_CREATE_CONTACT_PATH"), &MyAPIcreateContact.CreateContact{})
	beego.Router(beego.AppConfig.String("MyAPI_UPDATE_CONTACT_PATH"), &MyAPIupdateContact.UpdateContact{})
	beego.Router(beego.AppConfig.String("MyAPI_DELETE_CONTACT_PATH"), &MyAPIupdateContact.DeleteContact{})
}
