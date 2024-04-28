### A Temporary Email Package for Go

This package provides functionalities to interact with temporary email services in Go. You can create a temporary email address, retrieve emails, get email details, and more.

#### Example Usage:

```go
package main

import (
	"fmt"
	"github.com/Mixtre/tempmail"
)

func main() {
```
This is the main function where the code execution begins.

```go
    // Create a temporary email
	mail, err := tempmail.TempMail("example", tempmail.MailtoPlus)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
```
Here, we create a temporary email using the `TempMail` function from the `tempmail` package. We provide an alias ("example") and a domain (`tempmail.MailtoPlus`). If successful, `mail` will hold the temporary email address.

```go
    // Get mails for the temporary email
	emails := mail.GetMails()
	if emails == nil {
		fmt.Println("Error fetching emails")
		return
	}
```
We retrieve the emails for the temporary email address using the `GetMails` method. If there's an error or no emails are fetched, an error message is printed.

```go
    // Print subject of each email
	for _, email := range emails.MailList {
		fmt.Println("Subject:", email.Subject)
	}
```
We loop through each email in the fetched list and print its subject.

```go
    // Get details of a specific email
	emailDetail := mail.GetMail(emails.MailList[0].MailID)
	if emailDetail == nil {
		fmt.Println("Error fetching email details")
		return
	}
```
We retrieve the details of a specific email using the `GetMail` method with the ID of the first email in the list. If there's an error or the details cannot be fetched, an error message is printed.

```go
    // Print email details
	fmt.Println("From:", emailDetail.From)
	fmt.Println("Subject:", emailDetail.Subject)
}
```
Finally, we print the details of the fetched email, such as the sender and subject.

#### Supported Domains:

- mailto.plus
- fexpost.com
- fexbox.org
- mailbox.in.ua
- rover.info
- chitthi.in
- fextemp.com
- any.pink
- merepost.com

Feel free to contribute and extend this package to support more functionalities or temporary email services!
