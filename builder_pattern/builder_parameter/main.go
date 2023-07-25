package main

import "strings"

// bu çalışmada email objesini kullanıcıya manupile ettirmek yerine builderler vasıtasıyla mailimizi oluşturduk

type email struct {
	from    string
	subject string
	body    string
}
type EmailBuilder struct {
	email email
}

func (eb *EmailBuilder) From(from string) *EmailBuilder {
	if !strings.Contains(from, "@") {
		panic("Email should contain @ character")
	}
	eb.email.from = from
	return eb
}
func (eb *EmailBuilder) Subject(subject string) *EmailBuilder {
	eb.email.subject = subject
	if len(eb.email.subject) == 0 {
		eb.email.subject = "unknown subject"
	}
	return eb
}
func (eb *EmailBuilder) Body(body string) *EmailBuilder {
	eb.email.body = body
	return eb
}
func SendMail(email *email) {

}

type build func(*EmailBuilder)

func SendEmailz(build build) {
	builder := EmailBuilder{}
	build(&builder)
	SendMail(&builder.email)
}
func main() {
	SendEmailz(func(builder *EmailBuilder) {
		builder.From("hilmi").Subject("hello").Body("its meeeee")
	})
}
