package addressbook

import "github.com/kevinsnydercodes/protobuf-example/protobuf/compiled/go/address_book"

var (
	// AddressBook is an example of an address book created in memory using
	// the code generated from our Protobuf schema.
	AddressBook = &address_book.AddressBook{
		People: []*address_book.Person{
			&address_book.Person{
				Id:    1,
				Name:  "John Doe",
				Email: "jdoe@email.com",
				Phones: []*address_book.PhoneNumber{
					&address_book.PhoneNumber{
						Number: "205-555-5555",
						Type:   address_book.PhoneType_HOME,
					},
				},
				IsBlocked: false,
			},
			&address_book.Person{
				Id:    2,
				Name:  "Susan Boyle",
				Email: "sboyle@email.com",
				Phones: []*address_book.PhoneNumber{
					&address_book.PhoneNumber{
						Number: "415-555-5555",
						Type:   address_book.PhoneType_MOBILE,
					},
				},
				Tags: map[string]string{
					"Relationship": "Long time friend",
				},
				IsBlocked: false,
			},
		},
	}
)
