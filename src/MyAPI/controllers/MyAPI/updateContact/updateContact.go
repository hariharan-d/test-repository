package updateContact

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

type UpdateContact struct {
	beego.Controller
}

type DeleteContact struct {
	beego.Controller
}

func (c *UpdateContact) Get() {
	c.Ctx.Output.Body([]byte("Method Not Allowed."))
	return
}

func (c *UpdateContact) Post() {
	log.Println(beego.AppConfig.String("loglevel"), "Info", "Update Contact Page Start")
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
			log.Println(beego.AppConfig.String("loglevel"), "Info", "Update Contact Page Fail")
		} else {
			c.Ctx.Output.Body([]byte("Contact Updated Succesfully"))
			log.Println(beego.AppConfig.String("loglevel"), "Info", "Update Contact Page Success")
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

	book_id, err := user.Authenticate(UserName, Password) //authenticate user
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("Unable to Authenticate user")
		return
	}
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "BookID - ", book_id)

	ContactID := c.Input().Get("ContactID")
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "ContactID - ", ContactID)

	if ContactID == "" { //contact ID is mandatory to update
		err = errors.New("Please provide ContactID to perform an Update")
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		return
	}

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

	//check if contact exists
	_, err = ssb.GetContactDetails("", "", ContactID, book_id)
	if err == nil {
		_, err = db.Db.Exec("UPDATE lookup.contacts SET name=$1, email=$2, mobile=$3, address=$4 WHERE id=$5", ContactName, Email, Mobile, Address, ContactID)
		if err != nil {
			log.Println(beego.AppConfig.String("loglevel"), "Error", err)
			err = errors.New("Unable to Update Contact")
			return
		}
	} else {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("Contact does not exists.")
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		return
	}
}

func (c *DeleteContact) Post() {
	log.Println(beego.AppConfig.String("loglevel"), "Info", "Delete Contact Page Start")
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
			log.Println(beego.AppConfig.String("loglevel"), "Info", "Update Contact Page Fail")
		} else {
			c.Ctx.Output.Body([]byte("Contact Deleted Succesfully"))
			log.Println(beego.AppConfig.String("loglevel"), "Info", "Update Contact Page Success")
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

	ContactID := c.Input().Get("ContactID")
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "ContactID - ", ContactID)

	if ContactID == "" { //contact ID is mandatory to delete
		err = errors.New("Please provide ContactID to perform a Delete")
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		return
	}

	//check if contact exists
	_, err = ssb.GetContactDetails("", "", ContactID, book_id)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("Unable to find Contact")
		return
	} else {
		_, err = db.Db.Exec("DELETE FROM lookup.contacts WHERE id=$1", ContactID)
		if err != nil {
			log.Println(beego.AppConfig.String("loglevel"), "Error", err)
			err = errors.New("Unable to Delete Contact")
			return
		}
	}

}
