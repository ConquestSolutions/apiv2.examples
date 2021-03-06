package main

import (
	"fmt"
	"github.com/ConquestSolutions/apiv2.examples/conquest_api"
	"github.com/ConquestSolutions/apiv2.examples/go-swagger/client/asset_service"
	"github.com/ConquestSolutions/apiv2.examples/go-swagger/client/document_service"
	"github.com/ConquestSolutions/apiv2.examples/go-swagger/models"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"strings"
	"testing"
)

// TestAddDocument is a complete demonstration of how to add a document to an object
//
// This test is flaky due to how the ObjectKey is serialized with the default generated code in go-swagger.
//
// The work around is to make ObjectKey.TimestampValue a pointer, that way, if it's not set, it's not serialized.
func TestAddDocument(t *testing.T) {

	cfg, err := conquest_api.DefaultConfig()
	require.NoError(t, err)
	api, err := conquest_api.NewClient(*cfg)
	require.NoError(t, err)

	// Create a Facility (ParentID = 0)

	createAssetCommand := asset_service.NewCreateAssetParams().WithBody(&models.ConquestAPICreateAssetCommand{
		Proposed:         true,
		AssetDescription: "TestAddDocumentToAsset",
	})
	createAssetResult, err := api.AssetService.CreateAsset(createAssetCommand, nil)
	require.NoError(t, err)
	assetId := createAssetResult.GetPayload()

	fmt.Printf("Created asset with id %d", assetId)

	// Create a document record and get an upload link

	addDocumentCommand := document_service.NewAddDocumentParams().WithBody(&models.ConquestAPIAddDocumentCommand{
		// There's a bug here in how ObjectKey is serialised. Although we expect 'oneof' int32Value, stringValue or timestampValue; the go-swagger output looks like:
		//
		// "ObjectKey":{"int32Value":2610,"objectType":"ObjectType_Asset","timestampValue":"0001-01-01T00:00:00.000Z"}
		//
		// Instead of
		//
		// "ObjectKey":{"int32Value":2610,"objectType":"ObjectType_Asset"}
		//
		// This is an instance of the problem https://conquestsolutions.github.io/specification/optional_values/
		//
		// 'oneof' appears in OpenAPI 3.0, 'go-swagger' implements OpenAPI 2.0
		ObjectKey: &models.ConquestAPIObjectKey{
			ObjectType: models.ConquestAPIObjectTypeObjectTypeAsset,
			Int32Value: assetId,
		},
		DocumentDescription: "Test document",
		Address:             fmt.Sprintf("file://conquest_documents/Asset/%d/TestAddDocument.txt", assetId),
		ContentType:         "text/plain",
	})

	addDocumentResult, err := api.DocumentService.AddDocument(addDocumentCommand, nil)
	require.NoError(t, err)
	uploadInfo := addDocumentResult.GetPayload()

	// Upload the document
	// require.Equal(t, "PUT", uploadInfo.UploadMethod)
	fileData := ioutil.NopCloser(strings.NewReader("some text"))

	require.NoError(t, cfg.UploadFile(fileData, uploadInfo))
}
