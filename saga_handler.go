package artiefact

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	c "github.com/hirokazu/artiefact-backend/constants"
	"github.com/hirokazu/artiefact-backend/model"
)

// SagaApp is app for User resource
type SagaApp struct {
	*App
}

type StartSagaRequest struct {
	UserID    int64
	Latitude  float64
	Longitude float64
}

type StartSagaResponse struct {
	SagaID int64
}

type StartChapterRequest struct {
	SagaID    int64
	Latitude  float64
	Longitude float64
}

type StartChapterResponse struct {
	ChapterID int64
}

type EndChapterRequest struct {
	ChapterID int64
}

type EndChapterResponse struct {
	ChapterID int64
}

type LocationBatchRequest struct {
	ChapterID  int64
	Longitudes []int
	Latitudes  []int
}

type LocationBatchResponse struct {
	Response string
}

// BeginSagaHandler begins a saga
func (app *SagaApp) BeginSagaHandler(w http.ResponseWriter, r *http.Request) {
	// Begin Database
	tx, err := app.DB.Begin()
	if err != nil {
		e := &Error{
			Status: http.StatusInternalServerError,
			Type:   c.ErrorDBFailedToBegin,
			Detail: err.Error(),
		}
		json.NewEncoder(w).Encode(e)
		return
	}
	defer tx.Rollback()

	var request StartSagaRequest
	decoder := json.NewDecoder(r.Body)

	err = decoder.Decode(&request)
	if err != nil {
		fmt.Println(err)
	}

	//

	saga := &model.Saga{
		ID:                 request.UserID,
		BeginDate:          time.Now(),
		StartingLatitudes:  request.Latitude,
		StartingLongitudes: request.Longitude,
	}

	err = saga.Create(tx)
	if err != nil {
		fmt.Println(err)
	}

	err = tx.Commit()
	if err != nil {
		e := &Error{
			Status: http.StatusInternalServerError,
			Type:   c.ErrorDBFailedToCommit,
			Detail: err.Error(),
		}
		json.NewEncoder(w).Encode(e)
		return
	}

	// Create Response
	response := StartSagaResponse{
		SagaID: saga.ID,
	}

	json.NewEncoder(w).Encode(response)
}

// BeginChapterHander begins a saga
func (app *SagaApp) BeginChapterHandler(w http.ResponseWriter, r *http.Request) {
	// Begin Database
	tx, err := app.DB.Begin()
	if err != nil {
		e := &Error{
			Status: http.StatusInternalServerError,
			Type:   c.ErrorDBFailedToBegin,
			Detail: err.Error(),
		}
		json.NewEncoder(w).Encode(e)
		return
	}
	defer tx.Rollback()

	var request StartChapterRequest
	decoder := json.NewDecoder(r.Body)

	err = decoder.Decode(&request)
	if err != nil {
		fmt.Println(err)
	}

	chapter := &model.Chapter{
		SagaID:             request.SagaID,
		BeginDate:          time.Now(),
		StartingLatitudes:  request.Latitude,
		StartingLongitudes: request.Longitude,
	}

	err = chapter.Create(tx)
	if err != nil {
		fmt.Println(err)
	}

	err = tx.Commit()
	if err != nil {
		e := &Error{
			Status: http.StatusInternalServerError,
			Type:   c.ErrorDBFailedToCommit,
			Detail: err.Error(),
		}
		json.NewEncoder(w).Encode(e)
		return
	}

	// Create Response
	response := StartChapterResponse{
		ChapterID: chapter.ID,
	}

	json.NewEncoder(w).Encode(response)
}

// EndChapterHandler ends a saga
func (app *SagaApp) EndChapterHandler(w http.ResponseWriter, r *http.Request) {
	// Begin Database
	tx, err := app.DB.Begin()
	if err != nil {
		e := &Error{
			Status: http.StatusInternalServerError,
			Type:   c.ErrorDBFailedToBegin,
			Detail: err.Error(),
		}
		json.NewEncoder(w).Encode(e)
		return
	}
	defer tx.Rollback()

	var request EndChapterRequest
	decoder := json.NewDecoder(r.Body)

	err = decoder.Decode(&request)
	if err != nil {
		fmt.Println(err)
	}

	err = model.EndChapter(tx, request.ChapterID)
	if err != nil {
		fmt.Println(err)
	}

	err = tx.Commit()
	if err != nil {
		e := &Error{
			Status: http.StatusInternalServerError,
			Type:   c.ErrorDBFailedToCommit,
			Detail: err.Error(),
		}
		json.NewEncoder(w).Encode(e)
		return
	}

	// Create Response
	response := StartChapterResponse{
		ChapterID: request.ChapterID,
	}

	json.NewEncoder(w).Encode(response)
}

// TrackingBatchHandler records the batch tracking data
func (app *SagaApp) TrackingBatchHandler(w http.ResponseWriter, r *http.Request) {
	// Begin Database
	tx, err := app.DB.Begin()
	if err != nil {
		e := &Error{
			Status: http.StatusInternalServerError,
			Type:   c.ErrorDBFailedToBegin,
			Detail: err.Error(),
		}
		json.NewEncoder(w).Encode(e)
		return
	}
	defer tx.Rollback()

	var request LocationBatchRequest
	decoder := json.NewDecoder(r.Body)

	err = decoder.Decode(&request)
	if err != nil {
		fmt.Println(err)
	}
	var latitudes [128]int
	var longitudes [128]int

	copy(latitudes[:], request.Latitudes)
	copy(longitudes[:], request.Longitudes)

	trackingBatch := &model.TrackingBatch{
		Chapter:    request.ChapterID,
		Latitudes:  latitudes,
		Longitudes: longitudes,
	}

	err = trackingBatch.Create(tx)
	if err != nil {
		fmt.Println(err)
	}

	err = tx.Commit()
	if err != nil {
		e := &Error{
			Status: http.StatusInternalServerError,
			Type:   c.ErrorDBFailedToCommit,
			Detail: err.Error(),
		}
		json.NewEncoder(w).Encode(e)
		return
	}

	// Create Response
	response := LocationBatchResponse{
		Response: "ok",
	}

	json.NewEncoder(w).Encode(response)
}
