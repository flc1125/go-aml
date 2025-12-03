package aml

import (
	"context"
	"net/http"
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
	// Define the fields for the download batch response
}

type CheckNameRequest struct {
	// Define the fields for the name check request
}

type CheckNameResponse struct {
	// Define the fields for the name check response
}

type QueryOverResponse struct {
	Remain int `json:"remain"`
	Over   int `json:"over"`
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
	// TODO implement me
	panic("implement me")
}

func (c *Client) QueryOver(ctx context.Context, opts ...RequestOption) (*QueryOverResponse, *Response, error) {
	req, err := c.NewRequest(ctx, http.MethodGet, "amlupload/api/queryOver", nil, opts)
	if err != nil {
		return nil, nil, err
	}

	var response QueryOverResponse
	resp, err := c.Do(req, &response)
	if err != nil {
		return nil, resp, err
	}
	return &response, resp, nil
}
