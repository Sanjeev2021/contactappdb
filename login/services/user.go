package services

import (
	"errors"
	"time"

	//"contactapp/contact"
	//contactinfo "contactapp/contactInfo"

	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	//IsAdmin          bool
	usersCreatedByMe []User
	Username         string
	//ContactCreatedByMe []contact.Contact
	//ContactInfo        []contactinfo.ContactInfo
}

var usersCreatedByMe []User

func findUser(userSlice []User, username string) (*User, bool) {
	for i := 0; i < len(userSlice); i++ {
		if userSlice[i].ID.String() == username {
			return &userSlice[i], true
		}

	}
	return nil, false
}

// func NewAdmin(firstName, lastName, username string) *User {
// 	return &User{
// 		ID:        uuid.NewV4(),
// 		FirstName: firstName,
// 		LastName:  lastName,
// 		Username:  username,
// 		IsAdmin:   true,
// 	}
// }

func (u *User) NewUser(firstName, lastName, username string) (*User, error) {
	// if !u.IsAdmin {
	// 	return nil, errors.New(u.FirstName + "is not authorized to create a user")
	// }
	_, isUserExisit := findUser(u.usersCreatedByMe, username)
	if isUserExisit {
		return nil, errors.New("user already doest not exist")
	}

	newUser := &User{
		ID:        uuid.NewV4(),
		FirstName: firstName,
		LastName:  lastName,
		Username:  username,
		//IsAdmin:   false,
	}
	u.usersCreatedByMe = append(u.usersCreatedByMe, *newUser)
	return newUser, nil
}

func (u *User) ReadNewUser() ([]User, error) {
	// if !u.IsAdmin {
	// 	return nil, errors.New(u.FirstName + "not admin")
	// }
	return u.usersCreatedByMe, nil
}

func (u *User) UpdatedUser(id string, field string, value string) (*User, error) {
	userToUpdate, isUserExisit := FindUserById(id)
	if isUserExisit != nil {
		return nil, errors.New("user already doest not exist")
	}
	switch field {
	case "firstName":
		userToUpdate.FirstName = value
		return userToUpdate, nil
	case "lastName":
		userToUpdate.LastName = value
		return userToUpdate, nil

	default:
		return nil, errors.New("field is not present")

	}
}

func (u *User) DeleteUser(id interface{}) error {

	strID, ok := id.(string)
	if !ok {
		return errors.New("invalid id type")
	}

	_, isUserExist := findUser(usersCreatedByMe, strID)
	if !isUserExist {
		return errors.New("user already does not exist")
	}
	for i := 0; i < len(usersCreatedByMe); i++ {
		if usersCreatedByMe[i].ID.String() == id {

			usersCreatedByMe = append(usersCreatedByMe[:i], usersCreatedByMe[i+1:]...)
			return nil
		}
	}
	return errors.New("User does not exist")
}

func FindUserById(ID string) (*User, error) {
	for i := 0; i < len(usersCreatedByMe); i++ {
		if usersCreatedByMe[i].ID.String() == ID {
			return &usersCreatedByMe[i], nil
		}
	}
	return nil, errors.New("no user found")
}

func CreateUser(firstName, lastName string, IsAdmin bool, username string) (*User, error) {

	newUser := User{
		ID:        uuid.NewV4(),
		FirstName: firstName,
		LastName:  lastName,
		//IsAdmin:   IsAdmin,
		Username: username,
	}
	usersCreatedByMe = append(usersCreatedByMe, newUser)
	return &newUser, nil
}

func GetAllUser() (*[]User, error) {
	if len(usersCreatedByMe) == 0 {
		return nil, errors.New("no users")
	}
	return &usersCreatedByMe, nil
}

// func (u *User) CreateContact(contactname string, contacttype string, contactvalue interface{}) *contact.Contact {
// 	newcontact := contact.NewContact(contactname, contacttype, contactvalue)
// 	u.ContactCreatedByMe = append(u.ContactCreatedByMe, newcontact)
// 	return &newcontact
// }

// func FindContactById(id string) (*contact.Contact, error) {
// 	for i := 0; i < len(usersCreatedByMe); i++ {
// 		for j := 0; j < len(usersCreatedByMe[i].ContactCreatedByMe); j++ {
// 			if usersCreatedByMe[i].ContactCreatedByMe[j].ID.String() == id {
// 				return &usersCreatedByMe[i].ContactCreatedByMe[j], nil
// 			}
// 		}
// 	}
// 	return nil, errors.New("contact nhi hai")
// }

// func (u *User) UpdateContact(id, field, value string) (*contact.Contact, error) {
// 	contactfound, err := contact.FindContactById(id, u.ContactCreatedByMe)
// 	if err != nil {
// 		return nil, err
// 	}
// 	switch field {
// 	case "ContactName":
// 		contactfound.ContactName = value
// 	case "ContactType":
// 		contactfound.ContactType = value
// 	case "ContactValue":
// 		contactfound.ContactValue = value
// 	}
// 	return contactfound, nil

// }

// func (u *User) DeleteContact(id string) (*contact.Contact, error) {
// 	Contactotdelete, err := contact.FindContactById(id, u.ContactCreatedByMe)
// 	if err != nil {
// 		return nil, errors.New("User does not exist")
// 	}
// 	for i := 0; i < len(u.ContactCreatedByMe); i++ {
// 		if u.ContactCreatedByMe[i].ID.String() == id {
// 			u.ContactCreatedByMe = append(u.ContactCreatedByMe[:i], u.ContactCreatedByMe[i+1:]...)
// 			return Contactotdelete, nil
// 		}
// 	}
// 	return nil, errors.New("contact not found")
// }

// func GetAllContacts() (*[]contact.Contact, error) {
// 	var contacts []contact.Contact
// 	if len(usersCreatedByMe) == 0 {
// 		return nil, errors.New("no users")
// 	}

// 	for i := 0; i < len(usersCreatedByMe); i++ {
// 		for j := 0; j < len(usersCreatedByMe[i].ContactCreatedByMe); j++ {
// 			contacts = append(contacts, usersCreatedByMe[i].ContactCreatedByMe[j])
// 		}
// 	}
// 	if len(contacts) == 0 {
// 		return nil, errors.New("no contacts found")
// 	}
// 	return &contacts, nil
// }

// func (u *User) CreateContactInfo(contactinfotype, contactinfovalue string) (*contactinfo.ContactInfo, error) {
// 	contactinfo := contactinfo.NewContactInfo(contactinfotype, contactinfovalue)
// 	u.ContactInfo = append(u.ContactInfo, contactinfo)
// 	return &contactinfo, nil
// }

// func (u *User) UpdateContactInfo(id, field, value string) (*contactinfo.ContactInfo, error) {
// 	contactinfound, err := contactinfo.FindContactInfo(u.ContactInfo, id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	switch field {
// 	case "contactinfotype":
// 		contactinfound.ContactInfoType = value
// 	case "contactinfovalue":
// 		contactinfound.ContactInfoValue = value
// 	}
// 	return contactinfound, nil
// }

// func (u *User) DeleteContact(id string) (*contact.Contact, error) {
// 	Contactotdelete, err := contact.FindContactById(id, u.ContactCreatedByMe)
// 	if err != nil {
// 		return nil, errors.New("User does not exist")
// 	}
// 	for i := 0; i < len(u.ContactCreatedByMe); i++ {
// 		if u.ContactCreatedByMe[i].ID.String() == id {
// 			u.ContactCreatedByMe = append(u.ContactCreatedByMe[:i], u.ContactCreatedByMe[i+1:]...)
// 			return Contactotdelete, nil
// 		}
// 	}
// 	return nil, errors.New("contact not found")
// }

// func (u *User) DeleteContactInfo(id string) (*contactinfo.ContactInfo, error) {
// 	contactinfodelete, err := contactinfo.FindContactInfo(u.ContactInfo, id)
// 	if err != nil {
// 		return nil, errors.New("User does not exist")
// 	}
// 	for i := 0; i < len(u.ContactInfo); i++ {
// 		if u.ContactInfo[i].ID.String() == id {
// 			u.ContactInfo = append(u.ContactInfo[:i], u.ContactInfo[:i+1]...)
// 			return contactinfodelete, nil
// 		}
// 	}
// 	return nil, errors.New("Contact not found")

// }

// func GetAllContactInfo() (*[]contactinfo.ContactInfo, error) {
// 	var contactinfos []contactinfo.ContactInfo
// 	if len(usersCreatedByMe) == 0 {
// 		return nil, errors.New("no users")
// 	}

// 	for i := 0; i < len(usersCreatedByMe); i++ {
// 		for j := 0; j < len(usersCreatedByMe[i].ContactInfo); j++ {
// 			contactinfos = append(contactinfos, usersCreatedByMe[i].ContactInfo[j])
// 		}
// 	}
// 	if len(contactinfos) == 0 {
// 		return nil, errors.New("no contactinfos found")
// 	}
// 	return &contactinfos, nil
// }

// func FindContactInfoById(id string) (*contactinfo.ContactInfo, error) {
// 	for i := 0; i < len(usersCreatedByMe); i++ {
// 		for j := 0; j < len(usersCreatedByMe[i].ContactInfo); j++ {
// 			if usersCreatedByMe[i].ContactInfo[j].ID.String() == id {
// 				return &usersCreatedByMe[i].ContactInfo[j], nil
// 			}
// 		}
// 	}
// 	return nil, errors.New("contactinfo nhi hai")
// }

// creating jwt key
//var jwtKey = []byte("secret_key")

func (u *User) GenerateToken() (string, error) {
	// Create the JWT claims, which includes the username and expiry time
	claims := jwt.MapClaims{
		"username": u.Username,
		"exp":      time.Now().Add(time.Minute * 30).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
