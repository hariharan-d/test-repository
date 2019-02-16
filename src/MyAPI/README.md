Project Title:
ContactBook API by Hariharan

Getting Started:
This is a stateless API application which uses simple authentication for each serive request.

Prerequisites:
*).Requires Beego Framework.
*).A database with following schemas/tables :

	1). Contacts Table to store contacts:
		CREATE TABLE lookup.contacts
		(
		  id serial,
		  name character varying(50) NOT NULL,
		  email character varying(100) NOT NULL,
		  mobile character varying(20) NOT NULL,
		  address character varying(200) NOT NULL,
		  book_id smallint,
		  CONSTRAINT contact_pk PRIMARY KEY (id),
		  CONSTRAINT contack_fk FOREIGN KEY (book_id)
		      REFERENCES lookup."user" (id) MATCH SIMPLE
		      ON UPDATE NO ACTION ON DELETE NO ACTION
		)
		WITH (
		  OIDS=FALSE
		);
		ALTER TABLE lookup.contacts
		  OWNER TO postgres;

	2). User table:
		CREATE TABLE lookup."user"
		(
		  id serial,
		  username character varying(50) NOT NULL,
		  password character varying(20) NOT NULL,
		  CONSTRAINT user_pk PRIMARY KEY (id)
		)
		WITH (
		  OIDS=FALSE
		);
		ALTER TABLE lookup."user"
		  OWNER TO postgres;



Installing: 
*). Configuration changes (/MyAPI/conf/app.conf)
	1). Provide Database configuration details such as :
		i). example:	
		DBType=postgres
		DBIP=127.0.0.1
		DBPort=5432
		DBName=ContactBook
		DBUsername=postgres
		DBPassword=postgresPassword
	2).Provide the desirable HTTP port:
		i). HTTPPort=6003
	3). HTTPS:
		i). EnableHTTPS=false
		ii).HTTPSCertFile=../../key/PortalSSBAdmin/dev/public_key.cer
		iii).HTTPSKeyFile=../../key/PortalSSBAdmin/dev/private_key.cer
	4). Provide run mode and log level:
		i). RunMode = dev
		ii). loglevel=Debug
*). Build the App using:
	go build go build MyAPI.go

Running the tests:
1). Run the test using the following command:
	go test -run NameOfTest

Built With:
Beego Framework - The web framework used

Authors
Hariharan Durai.

