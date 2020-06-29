package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/rohsa/prime-concurrent-go/objects"
	"github.com/rohsa/prime-concurrent-go/rendoring"
	"net/http"
	"path/filepath"
	"strings"
	"time"
)

var viewTemplate = filepath.Join("view")

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	rendoring.RenderTemplate(w, viewTemplate, nil)
}

func DateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		decoder := json.NewDecoder(r.Body)
		var dateObject objects.Date
		err := decoder.Decode(&dateObject)
		if err != nil {
			http.Error(w, "Error parsing request", http.StatusInternalServerError)
		}
		values := strings.Split(dateObject.Date, "T")
		if len(values) > 0 {
			dateString := values[0]
			var timeString string
			if len(values) > 1 {
				timeString = values[1]
			}
			dateLayout := "2006-01-02 15:04"
			dateValue, err := time.Parse(dateLayout, fmt.Sprintf("%s %s", dateString, timeString))
			if err != nil {
				http.Error(w, "Invalid input format", http.StatusUnprocessableEntity)
				return
			}
			numberOfDays := time.Date(dateValue.Year(), dateValue.Month() + 1, 0, 0, 0, 0, 0, time.Local).Day()
			responseJson := objects.Date{Date:dateObject.Date, LastDayOfMonth:numberOfDays}
			responseBytes, err := json.Marshal(responseJson)
			if err != nil {
				// Return internal server error
				http.Error(w, "Empty date cannot be processed", http.StatusInternalServerError)
				return
			}
			_, err = fmt.Fprintf(w, "%s\n", responseBytes)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			http.Error(w, "Invalid input", http.StatusUnprocessableEntity)
			return
		}
	}
}
