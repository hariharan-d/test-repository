package viewContact

import (
	"MyModel/MyAPI"
	"MyModel/user"
	"MyModel/utils"
	"errors"
	"runtime/debug"

	"log"

	"encoding/json"

	"github.com/astaxie/beego"
)

type ViewContact struct {
	beego.Controller
}

func (c *ViewContact) Get() {
	c.Ctx.Output.Body([]byte("Method Not Allowed."))
	return

	return
}

type Page struct {
	PageNumber int
	Result     []ssb.Contact
}

func (c *ViewContact) Post() {
	log.Println(beego.AppConfig.String("loglevel"), "Info", "ViewContact POST Page Start")
	pip := c.Ctx.Input.IP()
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Client IP - ", pip)
	var err error
	defer func() {
		if l_exception := recover(); l_exception != nil {
			stack := debug.Stack()
			log.Println(beego.AppConfig.String("loglevel"), "Exception", string(stack))
			c.TplName = "error/error.html"
		}
		if err != nil {
			c.Ctx.Output.Body([]byte(err.Error()))
			log.Println(beego.AppConfig.String("loglevel"), "Info", "ViewContact POST Page Fail")
		} else {
			c.TplName = "admin/searchContact/searchContact.html"
			log.Println(beego.AppConfig.String("loglevel"), "Info", "ViewContact POST Page Success")
		}
		return
	}()
	utils.SetHTTPHeader(c.Ctx)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("System is unable to process your request.Please contact customer care")
		return
	}

	UserName := c.Input().Get("UserName")
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "UserName - ", UserName)

	Password := c.Input().Get("Password")
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Password - ", Password)

	if UserName == "" || Password == "" {
		err = errors.New("Please provide Username and Password")
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		return
	}

	book_id, err := user.Authenticate(UserName, Password) //authenticate user
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("Unable to Authenticate user")
		return
	}
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "BookID - ", book_id)

	ContactName := c.Input().Get("ContactName")
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "ContactName - ", ContactName)
	c.Data["ContactName"] = ContactName

	Email := c.Input().Get("Email")
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Email - ", Email)
	c.Data["Email"] = Email

	result, err := ssb.GetContactDetails(ContactName, Email, "", book_id) //get contacts list
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		return
	}

	//form a structure to send a page with 10 contacts
	var times int
	modulus := len(result) % 10
	if modulus != 0 {
		times = (len(result) / 10) + 1
	} else {
		times = (len(result) / 10)
	}

	var PageResult []Page

	for i := 1; i <= times; i++ {
		var tmp Page
		tmp.PageNumber = i
		if i != times {
			tmp.Result = result[(i-1)*10 : ((i)*10)+1] //take only 10 elements per page
		} else {
			tmp.Result = result[(i-1)*10:] // teke rest of the elements
		}
		PageResult = append(PageResult, tmp) //append to send in result
	}

	byt, err := json.Marshal(&PageResult)

	log.Println(beego.AppConfig.String("loglevel"), "Debug", "PageResult - ", PageResult)
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "byt - ", string(byt))

	c.Ctx.Output.Body(byt)

	return
}
