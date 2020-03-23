package main

import (
	"github.com/ConquestSolutions/apiv2.examples/conquest_api"
	"github.com/ConquestSolutions/apiv2.examples/go-swagger/client/asset_service"
	"github.com/ConquestSolutions/apiv2.examples/go-swagger/models"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/require"
	"testing"
)

// TestAssetCRUD is a complete demonstration of how to
// - Create an Object
// - Retrieve an Object
// - Update an Object
// - Delete an Object
func TestAssetCRUD(t *testing.T) {

	api, err := conquest_api.New()
	require.NoError(t, err)

	// Create a Facility (ParentID = 0)

	newDescription := "TestAssetCRUD"

	createAssetCommand := asset_service.NewCreateAssetParams().WithBody(&models.ConquestAPICreateAssetCommand{
		Proposed:         true,
		AssetDescription: newDescription,
	})
	createAssetResult, err := api.AssetService.CreateAsset(createAssetCommand)
	require.NoError(t, err)
	facilityId := createAssetResult.GetPayload()

	// Retrieve the object for viewing or editing

	getAssetRequest := asset_service.NewGetAssetParams().WithValue(facilityId)
	getAssetResponse, err := api.AssetService.GetAsset(getAssetRequest)
	require.NoError(t, err)

	lastEdit := getAssetResponse.GetPayload().EditDate
	originalAsset := getAssetResponse.GetPayload().Record
	require.Equal(t, newDescription, originalAsset.AssetDescription)

	// Make an edit

	now := strfmt.NewDateTime()
	newDescription += " (updated!)"
	updatedAsset := *originalAsset // shallow copy
	updatedAsset.AssetDescription = newDescription
	updatedAsset.YearCreated = &models.ConquestAPITimestampValue{Value: now}

	updateAssetCommand := asset_service.NewUpdateAssetParams().WithBody(&models.ConquestAPIAssetRecordChangeSet{
		AssetID: facilityId,
		Changes: []string{
			"AssetDescription",
			"YearCreated",
		},
		LastEdit: lastEdit,
		Original: originalAsset,
		Updated:  &updatedAsset,
	})
	updateAssetResult, err := api.AssetService.UpdateAsset(updateAssetCommand)
	require.NoError(t, err)
	require.True(t, updateAssetResult.GetPayload().Accepted)

	// Retrieve the updated object to verify the change

	getAssetRequest = asset_service.NewGetAssetParams().WithValue(facilityId)
	getAssetResponse, err = api.AssetService.GetAsset(getAssetRequest)
	require.NoError(t, err)
	retrievedAsset := getAssetResponse.GetPayload().Record
	require.Equal(t, newDescription, retrievedAsset.AssetDescription)
	require.Equal(t, now, retrievedAsset.YearCreated)

	// Unset the year created

	updatedAsset = *retrievedAsset // shallow copy
	updateAssetCommand = asset_service.NewUpdateAssetParams().WithBody(&models.ConquestAPIAssetRecordChangeSet{
		AssetID: facilityId,
		Changes: []string{
			"YearCreated",
		},
		LastEdit: lastEdit,
		Original: retrievedAsset,
		Updated:  &updatedAsset,
	})
	updateAssetResult, err = api.AssetService.UpdateAsset(updateAssetCommand)
	require.NoError(t, err)
	require.True(t, updateAssetResult.GetPayload().Accepted)

	// Retrieve the updated object and verify the change

	getAssetRequest = asset_service.NewGetAssetParams().WithValue(facilityId)
	getAssetResponse, err = api.AssetService.GetAsset(getAssetRequest)
	require.NoError(t, err)
	retrievedAsset = getAssetResponse.GetPayload().Record
	require.Nil(t, retrievedAsset.YearCreated)

	// Clean up our test data

	deleteAssetCommand := asset_service.NewDeleteAssetParams().WithBody(&models.ConquestAPIDeleteAssetCommand{
		AssetID: facilityId,
	})
	_, err = api.AssetService.DeleteAsset(deleteAssetCommand)
	require.NoError(t, err)

	// Confirm it's gone

	getAssetRequest = asset_service.NewGetAssetParams().WithValue(facilityId)
	_, err = api.AssetService.GetAsset(getAssetRequest)
	require.Error(t, err)
}