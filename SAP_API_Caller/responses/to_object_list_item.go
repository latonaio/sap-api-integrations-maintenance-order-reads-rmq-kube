package responses

type ToObjectListItem struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			MaintenanceOrder            string `json:"MaintenanceOrder"`
			MaintenanceOrderObjectList  string `json:"MaintenanceOrderObjectList"`
			MaintenanceObjectListItem   int    `json:"MaintenanceObjectListItem"`
			Equipment                   string `json:"Equipment"`
			MaintenanceNotification     string `json:"MaintenanceNotification"`
			Assembly                    string `json:"Assembly"`
			Material                    string `json:"Material"`
			SerialNumber                string `json:"SerialNumber"`
			UniqueItemIdentifier        string `json:"UniqueItemIdentifier"`
			FunctionalLocation          string `json:"FunctionalLocation"`
			MaintObjectListItemSequence string `json:"MaintObjectListItemSequence"`
		} `json:"results"`
	} `json:"d"`
}
