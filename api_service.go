package aml

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

type status string

const (
	statusOK    status = "ok"
	statusError status = "error"
)

type UploadBatchRequest struct {
	// Define the fields for the upload batch request
}

type UploadBatchResponse struct {
	// Define the fields for the upload batch response
}

type DownloadBatchRequest struct {
	// Define the fields for the download batch request
}

type DownloadBatchResponse struct {
}

type CheckNameRequest struct {
	NonEnglishName *string `json:"NonEnglishName,omitempty"` // 非英文姓名，無資料為空字串
	EnglishName    *string `json:"EnglishName,omitempty"`    // 英文姓名，無資料為空字串
	DOB            *string `json:"DOB,omitempty"`            // 生日 (YYYY/MM/DD，如：2017/09/01)，無資料為空字串
	Nationality    *string `json:"Nationality,omitempty"`    // 國籍，ISO 國家代碼(如臺灣為 TW，美國為 US) ，無資料為空字串
}

type rawCheckNameResponse struct {
	Status status `json:"status"`
	Error  string `json:"error,omitempty"` // 失敗原因，當 status 為 error 時回傳
	CheckNameResponse
}

type CheckNameResponse struct {
	RCScore int `json:"RCScore"` // 相似度分數，會回傳RC>=91 或 RC=0，如果上傳結果失敗，回傳-1。
}

type rawQueryOverResponse struct {
	Status status `json:"status"`
	QueryOverResponse
}

type QueryOverResponse struct {
	Remain int `json:"remain"` // 累計至前一日剩餘筆數
	Over   int `json:"over"`   // 累計至前一日超額查詢筆數
}

type Service interface {
	// UploadBatch 整批上傳檔案
	UploadBatch(ctx context.Context, request *UploadBatchRequest, opts ...RequestOption) (*UploadBatchResponse, *Response, error)

	DownloadBatch(ctx context.Context, request *DownloadBatchRequest, opts ...RequestOption) (*DownloadBatchResponse, *Response, error)

	CheckName(ctx context.Context, request *CheckNameRequest, opts ...RequestOption) (*CheckNameResponse, *Response, error)

	// QueryOver 查詢單筆線上查詢剩餘與超額筆數
	QueryOver(ctx context.Context, opts ...RequestOption) (*QueryOverResponse, *Response, error)
}

var _ Service = (*Client)(nil)

func (c *Client) UploadBatch(ctx context.Context, request *UploadBatchRequest, opts ...RequestOption) (*UploadBatchResponse, *Response, error) {
	// TODO implement me
	panic("implement me")
}

func (c *Client) DownloadBatch(ctx context.Context, request *DownloadBatchRequest, opts ...RequestOption) (*DownloadBatchResponse, *Response, error) {
	// TODO implement me
	panic("implement me")
}

func (c *Client) CheckName(ctx context.Context, request *CheckNameRequest, opts ...RequestOption) (*CheckNameResponse, *Response, error) {
	req, err := c.NewRequest(ctx, http.MethodPost, "amlupload/api/NameCheck", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var rawResponse rawCheckNameResponse
	resp, err := c.Do(req, &rawResponse)
	if err != nil {
		return nil, resp, err
	}
	if rawResponse.Status == statusError {
		return nil, resp, errors.New(rawResponse.Error)
	}
	return &rawResponse.CheckNameResponse, resp, nil
}

func (c *Client) QueryOver(ctx context.Context, opts ...RequestOption) (*QueryOverResponse, *Response, error) {
	req, err := c.NewRequest(ctx, http.MethodGet, "amlupload/api/queryOver", nil, opts)
	if err != nil {
		return nil, nil, err
	}

	var rawResponse rawQueryOverResponse
	resp, err := c.Do(req, &rawResponse)
	if err != nil {
		return nil, resp, err
	}
	if rawResponse.Status != statusOK {
		return nil, resp, fmt.Errorf("unexpected status: %s", rawResponse.Status)
	}
	return &rawResponse.QueryOverResponse, resp, nil
}
