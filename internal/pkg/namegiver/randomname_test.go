package namegiver

import (
	"strings"
	"testing"

	"github.com/dema501/randomjoke/internal/pkg/request"
)

func TestRandomNameGenerator(t *testing.T) {
	payload := `{"results":[{"gender":"female","name":{"title":"Miss","first":"Anna","last":"Li"},"location":{"street":{"number":2843,"name":"Victoria Road"},"city":"Auckland","state":"Hawke'S Bay","country":"New Zealand","postcode":18818,"coordinates":{"latitude":"66.0360","longitude":"23.2780"},"timezone":{"offset":"-3:00","description":"Brazil, Buenos Aires, Georgetown"}},"email":"anna.li@example.com","login":{"uuid":"17354e04-1d61-49cc-9f96-429e3c047d6d","username":"happyelephant216","password":"monalisa","salt":"2yeMYXu8","md5":"871778b1aed5706eef30c568b0b13934","sha1":"2295e43d183261248cd6fa6e6779ffc5190931a5","sha256":"b82466d130c2073da6d37ebbae6996f579a746caf28e0d7cffd03043aaab7315"},"dob":{"date":"1975-12-12T14:49:29.505Z","age":45},"registered":{"date":"2017-10-28T08:51:13.207Z","age":3},"phone":"(702)-484-8280","cell":"(065)-659-6612","id":{"name":"","value":null},"picture":{"large":"https://randomuser.me/api/portraits/women/19.jpg","medium":"https://randomuser.me/api/portraits/med/women/19.jpg","thumbnail":"https://randomuser.me/api/portraits/thumb/women/19.jpg"},"nat":"NZ"}],"info":{"seed":"3fb6f9b0a28ec60c","results":1,"page":1,"version":"1.3"}}`

	sa := &request.FakeSuperAgent{
		Body: strings.NewReader(payload),
	}

	rn := New(sa)
	if err := rn.Generate(); err != nil {
		t.Errorf("Expected NoError %v", err)
	}

	firstName, lastName := rn.GetName()

	if firstName != "Anna" {
		t.Errorf("Expected firstName Anna but get %v", firstName)
	}

	if lastName != "Li" {
		t.Errorf("Expected lastName Anna but get %v", firstName)
	}
}
