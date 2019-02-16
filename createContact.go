package createContact

import (
	"MyModel/MyAPI"
	"MyModel/db"
	"MyModel/user"
	"MyModel/utils"
	"errors"
	"log"
	"runtime/debug"

	"github.com/astaxie/beego"
)

type CreateContact struct {
	beego.Controller
}

func (c *CreateContact) Get() {
	c.Ctx.Output.Body([]byte("Method Not Allowed."))
	return
}

func (c *CreateContact) Post() {
	log.Println(beego.AppConfig.String("loglevel"), "Info", "Create Contact Page Start")
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
			log.Println(beego.AppConfig.String("loglevel"), "Info", "Create Contact Page Fail")
		} else {
			c.Ctx.Output.Body([]byte("Contact CreatedSuccesfully"))
			log.Println(beego.AppConfig.String("loglevel"), "Info", "Create Contact Page Success")
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

	Email := c.Input().Get("Email")
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Email - ", Email)

	Address := c.Input().Get("Address")
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Address - ", Address)

	Mobile := c.Input().Get("Mobile")
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Mobile - ", Mobile)

	if ContactName == "" || Email == "" || Address == "" || Mobile == "" {
		log.Println(beego.AppConfig.String("loglevel"), "Error", "Please enter all fields")
		err = errors.New("Feilds can't be blank")
		return
	}

	//check for duplicate contact
	_, err = ssb.GetContactDetails(ContactName, Email, "", book_id)
	if err != nil {
		if err.Error() == "No record found." {
			_, err = db.Db.Exec("INSERT INTO lookup.contacts (name,email,mobile,address,book_id) VALUES ($1, $2, $3, $4, $5)", ContactName, Email, Mobile, Address, book_id)
			if err != nil {
				log.Println(beego.AppConfig.String("loglevel"), "Error", err)
				err = errors.New("Unable to Create Contact")
				return
			}
		}
	} else {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("Contact already exists.")
		return
	}
}
