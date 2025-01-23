package handlers

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/google/uuid"
    "fasms/internal/models"
    "fasms/internal/db"
    log "github.com/sirupsen/logrus"
    "strconv"
    "time"
)

func CreateScheme(w http.ResponseWriter, r *http.Request) {
    var scheme models.Scheme
    if err := json.NewDecoder(r.Body).Decode(&scheme); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    scheme.ID = uuid.New().String()
    if err := db.CreateScheme(scheme); err != nil {
        http.Error(w, "Insertion error", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(scheme)
}

func GetSchemes(w http.ResponseWriter, r *http.Request) {
    schemes, err := db.GetSchemes()
    if err != nil {
        http.Error(w, "Schemes not found", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(schemes)
}
func GetSchemesForApplicant(w http.ResponseWriter, r *http.Request) {

    applicantID := r.URL.Query().Get("applicant")
    log.Info("applicant_id : " + applicantID)

    applicant, err:= db.GetApplicantByID(applicantID)
    if err != nil {
        http.Error(w, "Applicant not found", http.StatusNotFound)
        return
    } 
    noOfSchoolGoingKids := 0
    if (*applicant.Household != ""){
        noOfSchoolGoingKids = GetSchoolGoingKids(applicant.ID)
        log.Info("noOfSchoolGoingKids : " + strconv.Itoa(noOfSchoolGoingKids))
    }
    
    schemes, err := db.GetSchemes()
    if err != nil {
        http.Error(w, "Schemes not found", http.StatusInternalServerError)
        return
    }
    for _, scheme := range schemes {
        switch scheme.Name {
        case "Retrenchment Assistance Scheme (families)":
            benefits, _ := db.GetBenefitsByScheme( scheme.ID)
            if *applicant.EmploymentStatus == "unemployed" && noOfSchoolGoingKids > 0  && benefits !=nil {
                applicant.Benefit = append(applicant.Benefit, benefits... )
            }
        case "Retrenchment Assistance Scheme":
            benefits, _ := db.GetBenefitsByScheme( scheme.ID)
            if *applicant.EmploymentStatus == "unemployed" && benefits !=nil {
                applicant.Benefit = append(applicant.Benefit, benefits...)
            }

        default:
        }
    }

   



    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(applicant)
}

func GetSchoolGoingKids(applicantID string) int {
    noOfSchoolGoingKids := 0
    listOfDependents, err := db.GetDependentsOfApplicant(applicantID)
    if err != nil {
        return 0
    } 
    

    // Get today's date
	now := time.Now()

	// Calculate the start of 12 years ago and the end of 7 years ago
	start12YearsAgo := time.Date(now.Year()-12, time.January, 1, 0, 0, 0, 0, time.UTC)
	end7YearsAgo := time.Date(now.Year()-7, time.December, 31, 23, 59, 59, 0, time.UTC)

	//var filteredDependent []models.Dependent

	for _, dependent := range listOfDependents {
		// Parse the DateOfBirth string
		dob, err := time.Parse("02-01-2006", dependent.DateOfBirth)
		if err != nil {
			log.Error("Error parsing date for Dependent " + dependent.DateOfBirth)
			continue
		}

		// Check if the student's DOB falls in the range
		if dob.After(start12YearsAgo) && dob.Before(end7YearsAgo) {
			noOfSchoolGoingKids+=1
            log.Info("Name: " + dependent.Name + " DateOfBirth: " + dependent.DateOfBirth + " Schoolgoing : yes")
		} else {
            log.Info("Name: " + dependent.Name + " DateOfBirth: " + dependent.DateOfBirth + " Schoolgoing : no")
        }
	}
    
    return noOfSchoolGoingKids
}

func UpdateScheme(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
    var updatedScheme models.Scheme
    if err := json.NewDecoder(r.Body).Decode(&updatedScheme); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    if err := db.UpdateScheme(id, updatedScheme); err != nil {
        http.Error(w, "Scheme not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(updatedScheme)
}

func DeleteScheme(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
    if err := db.DeleteScheme(id); err != nil {
        http.Error(w, "Scheme not found", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
