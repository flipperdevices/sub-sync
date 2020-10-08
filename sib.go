package main

import (
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/sendinblue/APIv3-go-library/client"
	"github.com/sendinblue/APIv3-go-library/client/contacts"
	"github.com/sendinblue/APIv3-go-library/models"
)

func getSibClient(apiKey string) *client.SendinBlue {
	sib := client.NewHTTPClient(nil)
	sib.Transport.(*httptransport.Runtime).DefaultAuthentication = httptransport.APIKeyAuth("api-key", "header", apiKey)

	return sib
}

func createSibContact(sib *client.SendinBlue, email string) error {
	importParams := contacts.NewImportContactsParams()
	importParams.RequestContactImport = &models.RequestContactImport{
		FileBody: "email\n" + email,
		ListIds:  []int64{cfg.SendinblueListID},
	}
	_, err := sib.Contacts.ImportContacts(importParams, nil)
	if err != nil {
		return err
	}
	return nil
}
