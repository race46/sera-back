package gdive

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"os"

	"google.golang.org/api/drive/v3"
)

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}
func saveToken(path string, token *oauth2.Token) {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}
func getClient(config *oauth2.Config) *http.Client {
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)

	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	println("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}
	return tok
}
func UploadToDrive(srv *drive.Service, path string, name string) (string, error) {
	file := drive.File{Name: name}
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	res, err := srv.Files.Create(&file).Media(f).Do()
	if err != nil {
		return "", err
	}
	return res.Id, nil
}
func PublicUrl(srv *drive.Service, fileID string) (string, error) {
	permission := &drive.Permission{
		Type:               "anyone",
		Role:               "reader",
		AllowFileDiscovery: true,
	}

	_, _ = srv.Permissions.Create(fileID, permission).Do()
	file, err := srv.Files.Get(fileID).Fields("webContentLink").Do()
	if err != nil {
		return "", err
	}
	url := file.WebContentLink
	return url, nil
}
func DeleteFile(srv *drive.Service, fileID string) error {
	err := srv.Files.Delete(fileID).Do()
	return err
}
func GetService() (*drive.Service, error) {
	ctx := context.Background()
	b, err := os.ReadFile("./gdive/credentials.json")
	config, err := google.ConfigFromJSON(b, drive.DriveScope)
	client := getClient(config)

	srv, err := drive.NewService(ctx, option.WithHTTPClient(client))
	return srv, err
}
