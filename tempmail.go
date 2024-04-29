package tempmail

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Mail struct {
	Alias  string
	Domain string
}

type Attachment struct {
	AttachmentID int    `json:"attachment_id"`
	Name         string `json:"name"`
	Size         int    `json:"size"`
}

type MailDetail struct {
	Attachments []Attachment `json:"attachments"`
	Date        string       `json:"date"`
	From        string       `json:"from"`
	FromIsLocal bool         `json:"from_is_local"`
	FromMail    string       `json:"from_mail"`
	FromName    string       `json:"from_name"`
	Html        string       `json:"html"`
	IsTLS       bool         `json:"is_tls"`
	MailID      int          `json:"mail_id"`
	MessageID   string       `json:"message_id"`
	Result      bool         `json:"result"`
	Subject     string       `json:"subject"`
	Text        string       `json:"text"`
	To          string       `json:"to"`
}

type MailListItem struct {
	AttachmentCount     int    `json:"attachment_count"`
	FirstAttachmentName string `json:"first_attachment_name"`
	FromMail            string `json:"from_mail"`
	FromName            string `json:"from_name"`
	IsNew               bool   `json:"is_new"`
	MailID              int    `json:"mail_id"`
	Subject             string `json:"subject"`
	Time                string `json:"time"`
}

type MailResponse struct {
	Count    int            `json:"count"`
	FirstID  int            `json:"first_id"`
	LastID   int            `json:"last_id"`
	Limit    int            `json:"limit"`
	More     bool           `json:"more"`
	Result   bool           `json:"result"`
	MailList []MailListItem `json:"mail_list"`
}

const (
	MailtoPlus  = "mailto.plus"
	FexpostCom  = "fexpost.com"
	FexboxOrg   = "fexbox.org"
	MailboxInUa = "mailbox.in.ua"
	RoverInfo   = "rover.info"
	ChitthiIn   = "chitthi.in"
	FextempCom  = "fextemp.com"
	AnyPink     = "any.pink"
	MerepostCom = "merepost.com"
)

func (m *Mail) String() string {
	return fmt.Sprintf("%s@%s", m.Alias, m.Domain)
}

var httpClient = &http.Client{}

func TempMail(alias string, domain string) (*Mail, error) {
	if alias == "" {
		return nil, errors.New("alias can't be empty")
	}

	validDomains := []string{
		MailtoPlus,
		FexpostCom,
		FexboxOrg,
		MailboxInUa,
		RoverInfo,
		ChitthiIn,
		FextempCom,
		AnyPink,
		MerepostCom,
	}

	for _, v := range validDomains {
		if v == domain {
			return &Mail{Alias: alias, Domain: domain}, nil
		}
	}

	return nil, errors.New("invalid domain")
}

func (mail *Mail) GetMails() *MailResponse {
	resp, err := httpClient.Get(fmt.Sprintf("https://tempmail.plus/api/mails?email=%s@%s&limit=20&epin=", mail.Alias, mail.Domain))
	if err != nil {
		return nil
	}

	defer resp.Body.Close()

	var emails MailResponse
	err = json.NewDecoder(resp.Body).Decode(&emails)
	if err != nil {
		return nil
	}

	return &emails
}

func (mail *Mail) GetMail(id int) *MailDetail {
	resp, err := httpClient.Get(fmt.Sprintf("https://tempmail.plus/api/mails/%d?email=%s@%s&epin=", id, mail.Alias, mail.Domain))
	if err != nil {
		fmt.Println("Error fetching mail:", err)
		return nil
	}
	defer resp.Body.Close()

	var email MailDetail
	if err := json.NewDecoder(resp.Body).Decode(&email); err != nil {
		fmt.Println("Error decoding mail:", err)
		return nil
	}

	return &email
}

func (mail *Mail) GetSecretAddress() string {
	resp, err := httpClient.Get(fmt.Sprintf("https://tempmail.plus/api/box/hidden?email=%s@%s&epin=", mail.Alias, mail.Domain))
	if err != nil {
		fmt.Println("Error fetching mail:", err)
		return ""
	}

	defer resp.Body.Close()

	var hiddenmail map[string]interface{}

	if err := json.NewDecoder(resp.Body).Decode(&hiddenmail); err != nil {
		fmt.Println("Error decoding hidden mail:", err)
		return ""
	}

	return hiddenmail["email"].(string)
}

func (mail *Mail) GetAttachmentLink(AttachmentID int, attachments []Attachment, MailID int) string {
	for _, value := range attachments {
		if value.AttachmentID == AttachmentID {
			return fmt.Sprintf("https://tempmail.plus/api/mails/%d/attachments/%d?email=%s@%s&epin=", MailID, AttachmentID, mail.Alias, mail.Domain)
		}
	}
	return ""
}
