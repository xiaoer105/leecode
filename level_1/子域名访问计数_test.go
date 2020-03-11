package level_1

import (
	"log"
	"testing"
)

func TestSubdomainVisits(t *testing.T) {
	cpdomains :=[]string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}
	log.Println(SubdomainVisits(cpdomains))
}
