**TempMail Package**
=====================

The `tempmail` package provides a simple way to interact with temporary email services. It allows you to create a temporary email address, retrieve emails, and extract information from them.

**Getting Started**
-------------------

To use the `tempmail` package, you need to import it in your Go program:
```go
import "github.com/Mixtre/tempmail"
```
**Creating a Temporary Email Address**
-------------------------------------

To create a temporary email address, you need to specify an alias and a domain. The `TempMail` function returns a `*Mail` struct that represents the temporary email address:
```go
func TempMail(alias string, domain string) (*Mail, error)
```
* `alias`: The alias for the temporary email address (e.g. "myalias")
* `domain`: The domain for the temporary email address (e.g. `tempmail.MailtoPlus`, `tempmail.FexpostCom`, etc.)

Available domains:

* `tempmail.MailtoPlus`
* `tempmail.FexpostCom`
* `tempmail.FexboxOrg`
* `tempmail.MailboxInUa`
* `tempmail.RoverInfo`
* `tempmail.ChitthiIn`
* `tempmail.FextempCom`
* `tempmail.AnyPink`
* `tempmail.MerepostCom`

Example:
```go
mail, err := tempmail.TempMail("myalias", tempmail.MailtoPlus)
if err!= nil {
    log.Fatal(err)
}
```
**Printing the Temporary Email Address**
--------------------------------------

You can print the temporary email address using the `String()` method of the `Mail` struct:
```go
fmt.Println("Temporary Email Address:", mail.String())
```
This will output the temporary email address in the format `alias@domain`.

**Retrieving Emails**
--------------------

To retrieve emails, you can use the `GetMails` method of the `Mail` struct. It returns a `MailResponse` struct that contains a list of emails:
```go
func (mail *Mail) GetMails() *MailResponse
```
No parameters are required for this function.

Example:
```go
mailResponse := mail.GetMails()
if mailResponse == nil {
    log.Fatal("Error retrieving emails")
}
```
**Getting Email IDs**
---------------------

To get the IDs of the emails, you can iterate over the `MailList` field of the `MailResponse` struct:
```go
for _, email := range mailResponse.MailList {
    fmt.Println("Email ID:", email.MailID)
}
```
**Extracting Information from Emails**
-------------------------------------

You can extract information from emails using the `GetMail` method of the `Mail` struct. It returns a `MailDetail` struct that contains information about the email:
```go
func (mail *Mail) GetMail(id int) *MailDetail
```
* `id`: The ID of the email to retrieve (e.g. 123)

Example:
```go
mailDetail := mail.GetMail(123)
if mailDetail == nil {
    log.Fatal("Error retrieving email")
}
```
You can extract the following information from the `MailDetail` struct:

* `Subject`: The subject of the email
* `From`: The sender of the email
* `To`: The recipient of the email
* `Text`: The text content of the email
* `Html`: The HTML content of the email
* `Attachments`: A list of attachments
* `Date`: The date the email was received

**Printing the Email Content**
-----------------------------

You can print the email content using the `String()` method of the `MailDetail` struct:
```go
fmt.Println("Email Content:")
fmt.Println("Subject:", mailDetail.Subject)
fmt.Println("From:", mailDetail.From)
fmt.Println("To:", mailDetail.To)
fmt.Println("Text:", mailDetail.Text)
fmt.Println("Html:", mailDetail.Html)
fmt.Println("Attachments:")
for _, attachment := range mailDetail.Attachments {
    fmt.Println("Attachment ID:", attachment.AttachmentID)
    fmt.Println("Attachment Name:", attachment.Name)
}
```
This will output the email content in a human-readable format.

**Getting Attachment IDs**
-------------------------

To get the IDs of the attachments, you can iterate over the `Attachments` field of the `MailDetail` struct:
```go
for _, attachment := range mailDetail.Attachments {
    fmt.Println("Attachment ID:", attachment.AttachmentID)
}
```
**Extracting Attachments**
-------------------------

You can extract attachments from emails using the `GetAttachmentLink` method of the `Mail` struct. It returns a URL that you can use to download the attachment:
```go
func (mail *Mail) GetAttachmentLink(AttachmentID int, attachments []Attachment, MailID int) string
```
* `AttachmentID`: The ID of the attachment to retrieve (e.g. 123)
* `attachments`: A list of attachments for the email
* `MailID`: The ID of the email that the attachment belongs to (e.g. 123)

Example:
```go
attachmentLink := mail.GetAttachmentLink(123, mailDetail.Attachments, 123)
if attachmentLink == "" {
    log.Fatal("Error retrieving attachment link")
}
fmt.Println("Attachment Link:", attachmentLink)
```
This will output the URL of the attachment.

