package utils_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"ova-serial-api/internal/model"

	"ova-serial-api/internal/utils"
)

var _ = Describe("Map", func() {
	Context("Invert string int map", func() {
		It("should return correct result for empty map", func() {
			zeroMap := make(map[string]int)
			zeroMapRes := make(map[int]string)
			Expect(utils.InvertStrIntMap(zeroMap)).To(Equal(zeroMapRes))
		})

		It("should return correct result for not empty map", func() {
			testMap := map[string]int{"a": 1, "b": 2, "c": 3}
			testRes := map[int]string{1: "a", 2: "b", 3: "c"}
			Expect(utils.InvertStrIntMap(testMap)).To(Equal(testRes))
		})
	})

	Context("Convert Serial slice to map", func() {
		It("should return correct result for empty slice", func() {
			zeroSlice := make([]model.Serial, 0)
			zeroRes := make(map[uint64]model.Serial)
			Expect(utils.SerialSliceToMap(zeroSlice)).To(Equal(zeroRes))
		})

		It("should return correct result for not empty slice", func() {
			serial1 := model.Serial{UserID: 0, Title: "Test1", Genre: "genre1", Year: 2000, Seasons: 3}
			serial2 := model.Serial{UserID: 1, Title: "Test2", Genre: "genre2", Year: 2001, Seasons: 2}
			serial3 := model.Serial{UserID: 2, Title: "Test3", Genre: "genre3", Year: 2002, Seasons: 1}

			testSlice := []model.Serial{serial1, serial2, serial3}
			testRes := map[uint64]model.Serial{0: serial1, 1: serial2, 2: serial3}

			Expect(utils.SerialSliceToMap(testSlice)).To(Equal(testRes))
		})
	})
})
