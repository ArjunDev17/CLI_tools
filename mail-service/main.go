package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
	"gopkg.in/gomail.v2"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	InitCronScheduler()
	select {}
}

func InitCronScheduler() {

	c := cron.New(cron.WithSeconds())
	_, err := c.AddFunc("@every 5s", func() {
		SendEmailJob()
		os.Exit(0) // Terminate the process after the job executes
	})
	if err != nil {
		fmt.Println("Error adding cron function:", err)
		return
	}
	c.Start()
	fmt.Println("Cron scheduler started. The job will execute within 5 seconds.")
}

type EmailRequest struct {
	Recipients []string `json:"recipients"`
	Subject    string   `json:"subject"`
	LinkText   string   `json:"link_text"`
	LinkURL    string   `json:"link_url"`
}

func SendEmailJob() {
	url := os.Getenv("S3_URL")
	if url == "" {
		fmt.Println("No S3 URL found in environment variables")
		return
	}
	Link_s3 := url

	recipients := os.Getenv("RECIPIENTS")
	if recipients == "" {
		fmt.Println("No recipients found in environment variables")
		return
	}
	emailRecipients := strings.Split(recipients, ",")

	emailReq := EmailRequest{
		Recipients: emailRecipients,
		Subject:    "Daily Sales Report",
		LinkText:   "",
		LinkURL:    Link_s3,
	}

	err := sendEmail(emailReq)
	if err != nil {
		fmt.Println("Error sending email:", err)
		return
	}
	modifiedURL, err := IncrementDateInURL(url)
	if err != nil {
		fmt.Println("Error modifying URL:", err)
		return
	}
	fmt.Println("Modified URL:", modifiedURL)
	err = UpdateEnvVariable("S3_URL", modifiedURL)
	if err != nil {
		fmt.Println("Error updating S3_URL in environment file:", err)
		return
	}
	fmt.Println("S3_URL successfully updated in environment file.")
}

func sendEmail(req EmailRequest) error {
	yesterday := time.Now().AddDate(0, 0, -1)
	yesterdayFormatted := yesterday.Format("02/01/2006")

	body := fmt.Sprintf(`
		<p>Hi,</p>
		<p>Please find attached the report for yesterday (%s).</p>
		<p>You can access the report using the following link: <a href="%s">Access the Report</a></p>
		<p>Kindly review it at your earliest convenience and let me know if you have any questions or require further assistance.</p>
		<p>Thank you.<br>Best regards,<br>Arjun Singh <br><a href="mailto:sarjun@falcasolutions.com">sarjun@falcasolutions.com</a></p>
	`, yesterdayFormatted, req.LinkURL)

	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("SENDER_EMAIL"))
	m.SetHeader("To", req.Recipients...)
	m.SetHeader("Subject", req.Subject)
	m.SetBody("text/html", body)

	port, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		return fmt.Errorf("Invalid SMTP port: %v", err)
	}

	d := gomail.NewDialer(os.Getenv("SMTP_SERVER"), port, os.Getenv("SMTP_USERNAME"), os.Getenv("SMTP_PASSWORD"))

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("Failed to send email: %v", err)
	}

	return nil
}

func IncrementDateInURL(url string) (string, error) {
	re := regexp.MustCompile(`(\d{4}_\d{2}_\d{2})`)
	dateStr := re.FindString(url)
	if dateStr == "" {
		return "", fmt.Errorf("no date found in URL")
	}
	date, err := time.Parse("2006_01_02", dateStr)
	if err != nil {
		return "", fmt.Errorf("invalid date format: %v", err)
	}
	newDate := date.AddDate(0, 0, 1)
	newDateStr := newDate.Format("2006_01_02")
	modifiedURL := re.ReplaceAllString(url, newDateStr)
	return modifiedURL, nil
}

func UpdateEnvVariable(key, value string) error {
	envFile := ".env"
	content, err := ioutil.ReadFile(envFile)
	if err != nil {
		return fmt.Errorf("error reading .env file: %v", err)
	}
	lines := strings.Split(string(content), "\n")
	for i, line := range lines {
		if strings.HasPrefix(line, key+"=") {
			lines[i] = fmt.Sprintf("%s=%s", key, value)
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(envFile, []byte(output), 0644)
	if err != nil {
		return fmt.Errorf("error writing to .env file: %v", err)
	}
	return nil
}
