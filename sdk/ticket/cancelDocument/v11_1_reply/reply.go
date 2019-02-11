package Ticket_CancelDocumentReply_v11_1 // trcanr111

//import "encoding/xml"

type TicketCancelDocumentReply struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TRCANR_11_1_1A Ticket_CancelDocumentReply"`

	// for each processed document, up to 20, contains the information expected in CTS for further processing and a canned message number to be displayed to the user.
	TransactionResults []*TransactionResults `xml:"transactionResults"` // maxOccurs="20"
}

type TransactionResults struct {
	// Response analysis details
	ResponseDetails *ResponseAnalysisDetailsTypeI `xml:"responseDetails"`

	// Sequence number of processed document
	SequenceNumberDetails *ItemNumberTypeI `xml:"sequenceNumberDetails,omitempty"` // minOccurs="0"

	// the ticket numbers of the modified document, inclusive of check digit.
	TicketNumbers *TicketNumberTypeI `xml:"ticketNumbers,omitempty"` // minOccurs="0"

	// Describe the error
	ErrorGroup *ErrorGroupType `xml:"errorGroup,omitempty"` // minOccurs="0"

	// settelment authorization code number
	SacNumber *ReferenceInformationTypeI `xml:"sacNumber,omitempty"` // minOccurs="0"
}

//
// Complex structs
//

type ApplicationErrorDetailType struct {
	// Code identifying the data validation error condition.
	// xmlType: AlphaNumericString_Length1To5
	ErrorCode string `xml:"errorCode"`

	// Identification of a code list.
	// xmlType: AlphaNumericString_Length1To3
	ErrorCategory string `xml:"errorCategory,omitempty"` // minOccurs="0"

	// Code identifying the agency responsible for a code list.
	// xmlType: AlphaNumericString_Length1To3
	ErrorCodeOwner string `xml:"errorCodeOwner,omitempty"` // minOccurs="0"
}

type ApplicationErrorInformationType struct {
	// Application error details.
	ErrorDetails *ApplicationErrorDetailType `xml:"errorDetails"`
}

type ErrorGroupType struct {
	// The details of error/warning code.
	ErrorOrWarningCodeDetails *ApplicationErrorInformationType `xml:"errorOrWarningCodeDetails"`

	// The desciption of warning or error.
	ErrorWarningDescription *FreeTextInformationType `xml:"errorWarningDescription,omitempty"` // minOccurs="0"
}

type FreeTextDetailsType struct {
	// Value will be 4
	// xmlType: AlphaNumericString_Length1To3
	TextSubjectQualifier string `xml:"textSubjectQualifier"`

	// Manual source : M
	// xmlType: AlphaNumericString_Length1To3
	Source string `xml:"source"`

	// Mutually defined : ZZZ
	// xmlType: AlphaNumericString_Length1To3
	Encoding string `xml:"encoding"`
}

type FreeTextInformationType struct {
	// free text to detail the error
	FreeTextDetails *FreeTextDetailsType `xml:"freeTextDetails"`

	// Free text and message sequence numbers of the remarks.
	// xmlType: AlphaNumericString_Length1To199
	FreeText []string `xml:"freeText"` // maxOccurs="99"
}

type ItemNumberIdentificationTypeI struct {
	// document sequence number
	// xmlType: AlphaNumericString_Length1To35
	Number string `xml:"number,omitempty"` // minOccurs="0"

	// value used: FROM or TO
	// xmlType: AlphaNumericString_Length1To3
	Type string `xml:"type,omitempty"` // minOccurs="0"
}

type ItemNumberTypeI struct {
	// range of sequence numbers
	ItemNumberDetails []*ItemNumberIdentificationTypeI `xml:"itemNumberDetails"` // maxOccurs="2"
}

type ReferenceInformationTypeI struct {
	// reference details
	ReferenceDetails *ReferencingDetailsTypeI `xml:"referenceDetails"`
}

type ReferencingDetailsTypeI struct {
	// type
	// xmlType: AlphaNumericString_Length1To3
	Type string `xml:"type"`

	// value
	// xmlType: AlphaNumericString_Length1To14
	Value string `xml:"value,omitempty"` // minOccurs="0"
}

type ResponseAnalysisDetailsTypeI struct {
	// Tell what kind or response is handled, X for cancel.
	// xmlType: AlphaString_Length1To1
	ResponseType string `xml:"responseType"`

	// Tell if response was successful or not. O if success, N if error.
	// xmlType: AlphaString_Length1To1
	StatusCode string `xml:"statusCode"`
}

type TicketNumberDetailsTypeI struct {
	// document number
	// xmlType: AlphaNumericString_Length1To35
	Number string `xml:"number"`
}

type TicketNumberTypeI struct {
	// document identifier
	DocumentDetails *TicketNumberDetailsTypeI `xml:"documentDetails"`

	// status
	// xmlType: AlphaNumericString_Length1To3
	Status string `xml:"status,omitempty"` // minOccurs="0"
}
