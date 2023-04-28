//Em Go, os testes são realizados através de arquivos com o sufixo _test.go. O pacote de teste deve ter o mesmo nome do pacote que está sendo testado, mas com o sufixo _test.
//Os testes são escritos em funções com o prefixo Test, seguido do nome da função que está sendo testada. Essas funções recebem um único argumento do tipo *testing.T, que é usado para reportar falhas nos testes.

package estudos

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

//import (
//	"github.com/stretchr/testify/assert"
//	"testing"
//	"time"
//)

package campaign

import (
"github.com/stretchr/testify/assert"
"testing"
"time"
)

// Teste de unidade de forma manual!!!
//func TestNewCampaign(t *testing.T) {
//	name := "Campaing X"
//	content := "body"
//	contacts := []string{"lucascafe@gmail.com", "Renan@gmail.com"}
//
//	campaign := NewCampaign(name, content, contacts)
//
//	if campaign.ID != "1" {
//		t.Errorf("Expected 1")
//	} else if campaign.Name != name {
//		t.Errorf("Expected correct name")
//	} else if campaign.Content != content {
//		t.Errorf("expected correct content")
//	} else if len(campaign.Contacts) != len(contacts) {
//		t.Errorf("Expected correct contacts")
//	}
//}

// Usando a lib Testify
func TestNewCampaign(t *testing.T) {
	assert := assert.New(t)
	name := "Campaing X"
	content := "body"
	contacts := []string{"lucascafe@gmail.com", "Renan@gmail.com"}

	campaign, _ := NewCampaign(name, content, contacts)
	assert.Equal(campaign.ID, "1")
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))

}

func Test_NewCampaign_IDisNotNill(t *testing.T) {
	assert := assert.New(t)
	name := "Campaing X"
	content := "body"
	contacts := []string{"lucascafe@gmail.com", "Renan@gmail.com"}

	campaign, _ := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.ID)
}

func Test_NewCampaign_createdAt(t *testing.T) {
	assert := assert.New(t)
	name := "Campaing X"
	content := "body"
	contacts := []string{"lucascafe@gmail.com", "Renan@gmail.com"}
	now := time.Now().Add(-time.Minute)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.Greater(campaign.CreatedOn, now)
}

//func Test_NewCampaign_MustValidateName(t *testing.T)  {
//	assert := assert.New(t)
//	name := "Campaing X"
//	content := "body"
//	contacts := []string{"lucascafe@gmail.com", "Renan@gmail.com"}
//	now := time.Now().Add(-time.Minute)
//	assert := assert.New(t)
//
//	_, err := NewCampaign(name, content, contacts)
//}
