package action

import (
	"ants/crawler"
	"ants/http"
	"ants/node"
)

type RpcBase struct {
	NodeInfo *node.NodeInfo
	Result   bool
}

type LeftMeInRequest struct {
	RpcBase
}

// if result is true, it is master ;if false,close the client and connect to nodeinfo
type LeftMeInResponse struct {
	RpcBase
}

//
type DistributeRequest struct {
	RpcBase
	Request *http.Request
}
type DistributeReqponse struct {
	RpcBase
}

// report to master the result of crawl request
type ReportRequest struct {
	RpcBase
	ScrapeResult *crawler.ScrapeResult
}

type ReportResponse struct {
	RpcBase
}

// stop the node when all request is finished
type StopRequest struct {
	RpcBase
}
type StopResponse struct {
	RpcBase
}
