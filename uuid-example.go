package main

import (
	"fmt"

	"github.com/satori/go.uuid"
)

func main() {
	// Creating UUID Version 1
	uuidstr1, err := uuid.NewV1()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return
	}
	fmt.Printf("UUIDv1:		%s\n", uuidstr1)

	// Creating UUID Version 2
	// uuid.DomainPerson
	// uuid.DomainGroup
	uuidstr2, err := uuid.NewV2(uuid.DomainGroup)
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return
	}
	fmt.Printf("UUIDv2:		%s\n", uuidstr2)

	// Creating UUID Version 3
	uuidstr3 := uuid.NewV3(uuid.NamespaceDNS, "google.com")
	fmt.Printf("UUIDv3-DNS:	%s\n", uuidstr3)

	uuidstr3 = uuid.NewV3(uuid.NamespaceURL, "google.com")
	fmt.Printf("UUIDv3-URL:	%s\n", uuidstr3)

	uuidstr3 = uuid.NewV3(uuid.NamespaceOID, "Hello World!")
	fmt.Printf("UUIDv3-OID:	%s\n", uuidstr3)

	uuidstr3 = uuid.NewV3(uuid.NamespaceX500, "Hello World!")
	fmt.Printf("UUIDv3-X500:	%s\n", uuidstr3)

	// Creating UUID Version 4
	uuidstr4, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return
	}
	fmt.Printf("UUIDv4: 	%s\n", uuidstr4)

	// Creating UUID Version 5
	uuidstr5 := uuid.NewV5(uuid.NamespaceDNS, "google.com")
	fmt.Printf("UUIDv5-DNS:	%s\n", uuidstr5)

	uuidstr5 = uuid.NewV5(uuid.NamespaceURL, "google.com")
	fmt.Printf("UUIDv5-URL:	%s\n", uuidstr5)

	uuidstr5 = uuid.NewV5(uuid.NamespaceOID, "Hello World!")
	fmt.Printf("UUIDv5-OID:	%s\n", uuidstr5)

	uuidstr5 = uuid.NewV5(uuid.NamespaceX500, "Hello World!")
	fmt.Printf("UUIDv5-X500:	%s\n", uuidstr5)
}
