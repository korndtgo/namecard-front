package context_wrapper

import (
	"context"
	"errors"
	"fmt"

	"google.golang.org/grpc/metadata"
)

type ConTextWrapperResponse struct {
	TxId        string
	UserId      string
	DeviceID    string
	DeviceModel string
	DeviceOs    string
}

func ExtractGrpcContext(ctx context.Context) (ConTextWrapperResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("cannot extract ctx")
		return ConTextWrapperResponse{}, errors.New("cannot extract ctx")
	}

	res := ConTextWrapperResponse{}
	if len(md["txid"]) > 0 {
		res.TxId = md["txid"][0]
	}

	if len(md["userid"]) > 0 {
		res.UserId = md["userid"][0]
	}

	if len(md["deviceid"]) > 0 {
		res.DeviceID = md["deviceid"][0]
	}

	if len(md["deviceos"]) > 0 {
		res.DeviceOs = md["deviceos"][0]
	}

	if len(md["devicemodel"]) > 0 {
		res.DeviceModel = md["devicemodel"][0]
	}

	return res, nil
}

func StandardizedGrpcContext(ctx context.Context) context.Context {

	g, err := ExtractGrpcContext(ctx)
	if err != nil {
		fmt.Println("standardized grpc context error")
	}

	metaDataList := []string{}

	// must use lowercase with this lib

	// deviceId
	metaDataList = append(metaDataList, "deviceid")
	metaDataList = append(metaDataList, g.DeviceID)

	// deviceModel
	metaDataList = append(metaDataList, "devicemodel")
	metaDataList = append(metaDataList, g.DeviceModel)

	// deviceOs
	metaDataList = append(metaDataList, "deviceos")
	metaDataList = append(metaDataList, g.DeviceOs)

	// userId
	metaDataList = append(metaDataList, "userid")
	metaDataList = append(metaDataList, g.UserId)

	// txid
	metaDataList = append(metaDataList, "txid")
	metaDataList = append(metaDataList, g.TxId)

	ctxt := metadata.NewOutgoingContext(
		context.Background(),
		metadata.Pairs(metaDataList...),
	)

	return ctxt
}
