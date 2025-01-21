package unit

import (
	"testing"

	"github.com/Parichatx/dd.git/entity"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestID(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`Star is required`, func(t *testing.T) {
		mem := entity.Mem{
			Star:  "",
			Email: "ff@gg.com",
		}

		ok, err := govalidator.ValidateStruct(mem)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Star is required"))
	})

	t.Run(`Star is invalid`, func(t *testing.T) {
		mem := entity.Mem{
			Star:  "fff",
			Email: "ff@gg.com",
		}

		ok, err := govalidator.ValidateStruct(mem)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Star is invalid"))
	})
	t.Run(`OK`, func(t *testing.T) {
		mem := entity.Mem{
			Star:  "AAA11",
			Email: "ff@gg.com",
		}

		ok, err := govalidator.ValidateStruct(mem)
		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())

	})

	t.Run(`Email is invalid`, func(t *testing.T) {
		mem := entity.Mem{
			Star:  "AAA11",
			Email: "ffom",
		}

		ok, err := govalidator.ValidateStruct(mem)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Email is invalid"))
	})
}
