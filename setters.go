package mailyak

// To sets a list of recipient addresses.
//
// You can pass one or more addresses to this method, all of which are viewable to the recipients.
//
//	mail.To("dom@itsallbroken.com", "another@itsallbroken.com")
//
// or pass a slice of strings:
//
//	tos := []string{
//		"one@itsallbroken.com",
//		"two@itsallbroken.com"
//	}
//
//	mail.To(tos...)
func (m *MailYak) To(addrs ...string) {
	m.toAddrs = []string{}

	for _, addr := range addrs {
		trimmed := m.trimRegex.ReplaceAllString(addr, "")
		if trimmed == "" {
			continue
		}

		m.toAddrs = append(m.toAddrs, trimmed)
	}
}

// Bcc sets a list of blind carbon copy (BCC) addresses.
//
// You can pass one or more addresses to this method, none of which are viewable to the recipients.
//
//	mail.Bcc("dom@itsallbroken.com", "another@itsallbroken.com")
//
// or pass a slice of strings:
//
//	bccs := []string{
//		"one@itsallbroken.com",
//		"two@itsallbroken.com"
//	}
//
// 	mail.Bcc(bccs...)
func (m *MailYak) Bcc(addrs ...string) {
	m.bccAddrs = []string{}

	for _, addr := range addrs {
		trimmed := m.trimRegex.ReplaceAllString(addr, "")
		if trimmed == "" {
			continue
		}

		m.bccAddrs = append(m.bccAddrs, trimmed)
	}
}

// WriteBccHeader writes the BCC header to the MIME body when true. Defaults to
// false.
//
// This is usually required when writing the MIME body to an email API such as
// Amazon's SES, but can cause problems when sending emails via a SMTP server.
//
// Specifically, RFC822 says:
//
// 		Some  systems  may choose to include the text of the "Bcc" field only in the
// 		author(s)'s  copy,  while  others  may also include it in the text sent to
// 		all those indicated in the "Bcc" list.
//
// This ambiguity can result in some SMTP servers not stripping the BCC header
// and exposing the BCC addressees to recipients. For more information, see:
//
// 		https://github.com/domodwyer/mailyak/issues/14
//
func (m *MailYak) WriteBccHeader(shouldWrite bool) {
	m.writeBccHeader = shouldWrite
}

// Cc sets a list of carbon copy (CC) addresses.
//
// You can pass one or more addresses to this method, which are viewable to the other recipients.
//
//	mail.Cc("dom@itsallbroken.com", "another@itsallbroken.com")
//
// or pass a slice of strings:
//
//	ccs := []string{
//		"one@itsallbroken.com",
//		"two@itsallbroken.com"
//	}
//
// 	mail.Cc(ccs...)
func (m *MailYak) Cc(addrs ...string) {
	m.ccAddrs = []string{}

	for _, addr := range addrs {
		trimmed := m.trimRegex.ReplaceAllString(addr, "")
		if trimmed == "" {
			continue
		}

		m.ccAddrs = append(m.ccAddrs, trimmed)
	}
}

// From sets the sender email address.
//
// Users should also consider setting FromName().
func (m *MailYak) From(addr string) {
	m.fromAddr = m.trimRegex.ReplaceAllString(addr, "")
}

// FromName sets the sender name.
//
// If set, emails typically display as being from:
//
// 		From Name <sender@example.com>
//
func (m *MailYak) FromName(name string) {
	m.fromName = m.trimRegex.ReplaceAllString(name, "")
}

// ReplyTo sets the Reply-To email address.
//
// Setting a ReplyTo address is optional.
func (m *MailYak) ReplyTo(addr string) {
	m.replyTo = m.trimRegex.ReplaceAllString(addr, "")
}

// Subject sets the email subject line.
func (m *MailYak) Subject(sub string) {
	m.subject = m.trimRegex.ReplaceAllString(sub, "")
}

// SetHeader sets a custom header field name and value.
// It is up to the caller to ensure that custom headers do not conflict with built-in headers.
func (m *MailYak) SetHeader(name, value string) {
	m.customHeaders[name] = value
}

// DeleteHeader removes a custom header field.
func (m *MailYak) DeleteHeader(name string) {
	delete(m.customHeaders, name)
}

// ClearHeaders removes all custom header fields.
func (m *MailYak) ClearHeaders() {
	m.customHeaders = make(map[string]string)
}
