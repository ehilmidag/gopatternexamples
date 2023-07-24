package main

type Doc struct {
	text string
}

type Scanner interface {
	Print(d Doc)
	Edit(d Doc)
	Fax(d Doc)
}

// yeni scannerimiz son teknoloji bunların hepsini yapabiliyor ama eski bir scannerimiz olsaydı ve fax özelliği olmasaydı
// bu yüzden bu interfaceleri ayırıp  bağımsız interfaceler seklinde koymalıyız.
type NewScanner struct {
}

func (n *NewScanner) Print(d Doc) {

}

func (n *NewScanner) Scan(d Doc) {

}

func (n *NewScanner) Fax(d Doc) {

}

func main() {

}
