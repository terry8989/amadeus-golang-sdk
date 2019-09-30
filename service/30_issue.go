package service

import (
	"github.com/tmconsulting/amadeus-golang-sdk/v2/client"
	"github.com/tmconsulting/amadeus-golang-sdk/v2/structs/docIssuance/issueTicket/v09.1"
)

func (s *service) DocIssuanceIssueTicket(query *DocIssuance_IssueTicket_v09_1.Request) (*DocIssuance_IssueTicket_v09_1.Response, *client.ResponseSOAPHeader, error) {
	return s.sdk.DocIssuanceIssueTicketV091(query)
}
