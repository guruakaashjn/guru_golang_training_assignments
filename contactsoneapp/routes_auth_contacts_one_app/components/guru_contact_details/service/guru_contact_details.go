package guru_contact_details

import (
	"github.com/google/uuid"
)

// var ContactDetailsId = 3

type ContactDetails struct {
	ContactDetailsId uuid.UUID
	TypeName         string
	TypeValue        string
	IsActive         bool
}

func NewContactDetails(TypeName string, TypeValue string) *ContactDetails {

	var newObjectOfContactDetails = &ContactDetails{
		ContactDetailsId: uuid.New(),
		TypeName:         TypeName,
		TypeValue:        TypeValue,
		IsActive:         true,
	}
	return newObjectOfContactDetails
}

func CreateContactDetails(TypeName, TypeValue string) *ContactDetails {
	return NewContactDetails(TypeName, TypeValue)
}

func (cd *ContactDetails) UpdateContactDetails(keyName string, keyValue string) *ContactDetails {
	cd.TypeName = keyName
	cd.TypeValue = keyValue
	return cd
}

func (cd *ContactDetails) UpdateContactDetailsObject(contactDetailsTempObj *ContactDetails) *ContactDetails {
	if contactDetailsTempObj.TypeName != "" && contactDetailsTempObj.TypeName != cd.TypeName {
		cd.TypeName = contactDetailsTempObj.TypeName
	}
	if contactDetailsTempObj.TypeValue != "" && contactDetailsTempObj.TypeValue != cd.TypeValue {
		cd.TypeValue = contactDetailsTempObj.TypeValue
	}
	return cd
}

// func (cd *ContactDetails) ReadContactDetails() {
// 	fmt.Printf("\nContact Details")
// 	fmt.Printf("\nType : %s and Type Value : %s", cd.TypeName, cd.TypeValue)
// }

// func (cd *ContactDetails) ReadContactDetails() (readContactDetails string) {
// 	readContactDetails += "Contact Details" +
// 		"\nContact Details Id: " + cd.ContactDetailsId.String() +
// 		"\nType : " + cd.TypeName +
// 		" and Type Value : " + cd.TypeValue +
// 		"\nIsActive : " + strconv.FormatBool(cd.IsActive)
// 	return readContactDetails
// }

func (cd *ContactDetails) ReadContactDetails() (bool, *ContactDetails) {
	if cd.IsActive {
		return true, cd
	}
	return false, cd
}

func (cd *ContactDetails) DeleteContactDetails() *ContactDetails {
	// cd.TypeName = ""
	// cd.TypeValue = ""
	cd.IsActive = false
	return cd
}

func (cd *ContactDetails) GetType() string {
	return cd.TypeName
}
func (cd *ContactDetails) GetTypeValue() string {
	return cd.TypeValue
}

func (cd *ContactDetails) GetContactDetailsId() uuid.UUID {
	return cd.ContactDetailsId
}
