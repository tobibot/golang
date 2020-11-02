package learninggo

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"strings"
	"testing"

	"gopkg.in/xmlpath.v1"
)

var (
	requestBody = `<?xml version="1.0" encoding="utf-8"?>
					<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema">
						<soap:Body>
							<SendeRequestNode xmlns="...">
								<dataContainer>
									<Data>
										UEsDBBQAAAAIANl9YlFBW9QQChIAAGcaAAAvACQANTIxMTExMTAwXzk….. [ganz viel base64-Zeugs]
									</Data>
								</dataContainer>
								<bsnr>521111100</bsnr>
								<lanr>999999901</lanr>
								<containerGuid>e202fa65-4a3d-449c-8cb0-41d6eef445bf</containerGuid>
								<systemOid>f3ef36ec-540f-e474-a4f5-30f7827df6f8</systemOid>
								<timeSend>02.11.2020 15:46:51</timeSend>
								<isTest>true</isTest>
							</SendeRequestNode>
						</soap:Body>
					</soap:Envelope>`
)

func TestXML(t *testing.T) {

	path := xmlpath.MustCompile("/Envelope/Body/SendeRequestNode/dataContainer")
	root, err := xmlpath.Parse(strings.NewReader(requestBody))

	if err != nil {
		log.Fatal(err)
	}

	var got string

	want := "UEsDBBQAAAAIANl9YlFBW9QQChIAAGcaAAAvACQANTIxMTExMTAwXzk….. [ganz viel base64-Zeugs]"

	if value, ok := path.String(root); ok {
		got = value
	}

	got = strings.TrimSpace(got)
	got = strings.Trim(got, "\t")

	if want != got {
		t.Errorf("Wanted  %s\nbut got %s", want, got)
	}

}

// Here with un/-marshalling from/to struct
type Data struct {
	XMLName    xml.Name    `xml:"data" json:"-"`
	PersonList []XMLPerson `xml:"person" json:"people"`
}

type XMLPerson struct {
	XMLName   xml.Name `xml:"person" json:"-"`
	Firstname string   `xml:"firstname" json:"firstname"`
	Lastname  string   `xml:"lastname" json:"lastname"`
	Address   *Address `xml:"address" json:"address,omitempty"`
}

type Address struct {
	City  string `xml:"city" json:"city,omitempty"`
	State string `xml:"state" json:"state,omitempty"`
}

func TestXMLByParsingToStruct(t *testing.T) {
	rawXMLData := "<data><person><firstname>Nic</firstname><lastname>Raboy</lastname><address><city>San Francisco</city><state>CA</state></address></person><person><firstname>Maria</firstname><lastname>Raboy</lastname></person></data>"
	var data Data
	xml.Unmarshal([]byte(rawXMLData), &data)
	jsonData, _ := json.Marshal(data)

	fmt.Println(string(jsonData))
}
