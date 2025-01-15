package test

import (
	"main/entity"
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestValid(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`Success`, func(t *testing.T) {
		product := entity.Products{
			Name:  "test",
			Price: 500.00,
			SKU:   "ABC12345",
		}

		ok, err := govalidator.ValidateStruct(product)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())

	})
}
