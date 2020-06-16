package gsheetsreader


import (
        "log"
        "golang.org/x/net/context"
        "golang.org/x/oauth2/google"
        "google.golang.org/api/sheets/v4"
)


func getService(f []byte) *sheets.Service{
        // Creates a google sheets service using gserviceaccount jwt credentials 

        config, err := google.JWTConfigFromJSON(f, "https://www.googleapis.com/auth/spreadsheets.readonly")

        if err != nil {
                log.Fatalf("Unable to parse client secret file to config: %v", err)
        }

        client := config.Client(context.Background())
        srv, err := sheets.New(client)
                
        if err != nil {
                log.Fatalf("Unable to retrieve Sheets client: %v", err)
        }

        return srv
}
