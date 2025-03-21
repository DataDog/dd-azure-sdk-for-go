//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsupport

const (
	moduleName    = "armsupport"
	moduleVersion = "v1.0.0"
)

// CommunicationDirection - Direction of communication.
type CommunicationDirection string

const (
	CommunicationDirectionInbound  CommunicationDirection = "inbound"
	CommunicationDirectionOutbound CommunicationDirection = "outbound"
)

// PossibleCommunicationDirectionValues returns the possible values for the CommunicationDirection const type.
func PossibleCommunicationDirectionValues() []CommunicationDirection {
	return []CommunicationDirection{
		CommunicationDirectionInbound,
		CommunicationDirectionOutbound,
	}
}

// CommunicationType - Communication type.
type CommunicationType string

const (
	CommunicationTypePhone CommunicationType = "phone"
	CommunicationTypeWeb   CommunicationType = "web"
)

// PossibleCommunicationTypeValues returns the possible values for the CommunicationType const type.
func PossibleCommunicationTypeValues() []CommunicationType {
	return []CommunicationType{
		CommunicationTypePhone,
		CommunicationTypeWeb,
	}
}

// PreferredContactMethod - Preferred contact method.
type PreferredContactMethod string

const (
	PreferredContactMethodEmail PreferredContactMethod = "email"
	PreferredContactMethodPhone PreferredContactMethod = "phone"
)

// PossiblePreferredContactMethodValues returns the possible values for the PreferredContactMethod const type.
func PossiblePreferredContactMethodValues() []PreferredContactMethod {
	return []PreferredContactMethod{
		PreferredContactMethodEmail,
		PreferredContactMethodPhone,
	}
}

// SeverityLevel - A value that indicates the urgency of the case, which in turn determines the response time according to
// the service level agreement of the technical support plan you have with Azure. Note: 'Highest
// critical impact', also known as the 'Emergency - Severe impact' level in the Azure portal is reserved only for our Premium
// customers.
type SeverityLevel string

const (
	SeverityLevelCritical              SeverityLevel = "critical"
	SeverityLevelHighestcriticalimpact SeverityLevel = "highestcriticalimpact"
	SeverityLevelMinimal               SeverityLevel = "minimal"
	SeverityLevelModerate              SeverityLevel = "moderate"
)

// PossibleSeverityLevelValues returns the possible values for the SeverityLevel const type.
func PossibleSeverityLevelValues() []SeverityLevel {
	return []SeverityLevel{
		SeverityLevelCritical,
		SeverityLevelHighestcriticalimpact,
		SeverityLevelMinimal,
		SeverityLevelModerate,
	}
}

// Status - Status to be updated on the ticket.
type Status string

const (
	StatusClosed Status = "closed"
	StatusOpen   Status = "open"
)

// PossibleStatusValues returns the possible values for the Status const type.
func PossibleStatusValues() []Status {
	return []Status{
		StatusClosed,
		StatusOpen,
	}
}

// Type - The type of resource.
type Type string

const (
	TypeMicrosoftSupportSupportTickets Type = "Microsoft.Support/supportTickets"
	TypeMicrosoftSupportCommunications Type = "Microsoft.Support/communications"
)

// PossibleTypeValues returns the possible values for the Type const type.
func PossibleTypeValues() []Type {
	return []Type{
		TypeMicrosoftSupportSupportTickets,
		TypeMicrosoftSupportCommunications,
	}
}
