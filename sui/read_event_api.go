package sui

import (
	"context"
	"encoding/json"
	"github.com/block-vision/sui-go-sdk/httpconn"
	"github.com/block-vision/sui-go-sdk/models"
	"github.com/block-vision/sui-go-sdk/sui_error"
	"github.com/tidwall/gjson"
)

type IReadEventFromSuiAPI interface {
	GetEventsByEventType(ctx context.Context, req models.GetEventsByEventTypeRequest, opts ...interface{}) (models.GetEventsByEventTypeResponse, error)
	GetEventsByModule(ctx context.Context, req models.GetEventsByModuleRequest, opts ...interface{}) (models.GetEventsByModuleResponse, error)
	GetEventsByObject(ctx context.Context, req models.GetEventsByObjectRequest, opts ...interface{}) (models.GetEventsByObjectResponse, error)
	GetEventsByOwner(ctx context.Context, req models.GetEventsByOwnerRequest, opts ...interface{}) (models.GetEventsByOwnerResponse, error)
	GetEventsBySender(ctx context.Context, req models.GetEventsBySenderRequest, opts ...interface{}) (models.GetEventsBySenderResponse, error)
	GetEventsByTransaction(ctx context.Context, req models.GetEventsByTransactionRequest, opts ...interface{}) (models.GetEventsByTransactionResponse, error)
}

type suiReadEventFromSuiImpl struct {
	conn *httpconn.HttpConn
}

//GetEventsByEventType
func (s *suiReadEventFromSuiImpl) GetEventsByEventType(ctx context.Context, req models.GetEventsByEventTypeRequest, opts ...interface{}) (models.GetEventsByEventTypeResponse, error) {
	var rsp models.GetEventsByEventTypeResponse
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "sui_getEventsByEventType",
		Params: []interface{}{
			req.EventType,
			req.Count,
			req.StartTime,
			req.EndTime,
		},
	})
	if err != nil {
		return models.GetEventsByEventTypeResponse{}, err
	}
	if !gjson.ValidBytes(respBytes) {
		return models.GetEventsByEventTypeResponse{}, sui_error.ErrInvalidJson
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp.Result)
	if err != nil {
		return models.GetEventsByEventTypeResponse{}, err
	}
	return rsp, nil
}

func (s *suiReadEventFromSuiImpl) GetEventsByModule(ctx context.Context, req models.GetEventsByModuleRequest, opts ...interface{}) (models.GetEventsByModuleResponse, error) {
	var rsp models.GetEventsByModuleResponse
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "sui_getEventsByModule",
		Params: []interface{}{
			req.Package,
			req.Module,
			req.Count,
			req.StartTime,
			req.EndTime,
		},
	})
	if err != nil {
		return models.GetEventsByModuleResponse{}, err
	}
	if !gjson.ValidBytes(respBytes) {
		return models.GetEventsByModuleResponse{}, sui_error.ErrInvalidJson
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp.Result)
	if err != nil {
		return models.GetEventsByModuleResponse{}, err
	}
	return rsp, nil
}

func (s *suiReadEventFromSuiImpl) GetEventsByObject(ctx context.Context, req models.GetEventsByObjectRequest, opts ...interface{}) (models.GetEventsByObjectResponse, error) {
	var rsp models.GetEventsByObjectResponse
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "sui_getEventsByObject",
		Params: []interface{}{
			req.Object,
			req.Count,
			req.StartTime,
			req.EndTime,
		},
	})
	if err != nil {
		return models.GetEventsByObjectResponse{}, err
	}
	if !gjson.ValidBytes(respBytes) {
		return models.GetEventsByObjectResponse{}, sui_error.ErrInvalidJson
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp.Result)
	if err != nil {
		return models.GetEventsByObjectResponse{}, err
	}
	return rsp, nil
}

func (s *suiReadEventFromSuiImpl) GetEventsByOwner(ctx context.Context, req models.GetEventsByOwnerRequest, opts ...interface{}) (models.GetEventsByOwnerResponse, error) {
	var rsp models.GetEventsByOwnerResponse
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "sui_getEventsByOwner",
		Params: []interface{}{
			req.Owner,
			req.Count,
			req.StartTime,
			req.EndTime,
		},
	})
	if err != nil {
		return models.GetEventsByOwnerResponse{}, err
	}
	if !gjson.ValidBytes(respBytes) {
		return models.GetEventsByOwnerResponse{}, sui_error.ErrInvalidJson
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp.Result)
	if err != nil {
		return models.GetEventsByOwnerResponse{}, err
	}
	return rsp, nil
}

func (s *suiReadEventFromSuiImpl) GetEventsBySender(ctx context.Context, req models.GetEventsBySenderRequest, opts ...interface{}) (models.GetEventsBySenderResponse, error) {
	var rsp models.GetEventsBySenderResponse
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "sui_getEventsBySender",
		Params: []interface{}{
			req.Sender,
			req.Count,
			req.StartTime,
			req.EndTime,
		},
	})
	if err != nil {
		return models.GetEventsBySenderResponse{}, err
	}
	if !gjson.ValidBytes(respBytes) {
		return models.GetEventsBySenderResponse{}, sui_error.ErrInvalidJson
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp.Result)
	if err != nil {
		return models.GetEventsBySenderResponse{}, err
	}
	return rsp, nil
}

func (s *suiReadEventFromSuiImpl) GetEventsByTransaction(ctx context.Context, req models.GetEventsByTransactionRequest, opts ...interface{}) (models.GetEventsByTransactionResponse, error) {
	var rsp models.GetEventsByTransactionResponse
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "sui_getEventsByTransaction",
		Params: []interface{}{
			req.Digest,
		},
	})
	if err != nil {
		return models.GetEventsByTransactionResponse{}, err
	}
	if !gjson.ValidBytes(respBytes) {
		return models.GetEventsByTransactionResponse{}, sui_error.ErrInvalidJson
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp.Result)
	if err != nil {
		return models.GetEventsByTransactionResponse{}, err
	}
	return rsp, nil
}