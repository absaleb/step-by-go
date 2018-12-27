package zoopla

type ListingDelete struct {
	DeletionReason string `json:"deletion_reason,omitempty"`
	ListingReference string `json:"listing_reference"`
}

