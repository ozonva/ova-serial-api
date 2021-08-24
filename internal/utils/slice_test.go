package utils_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"ova-serial-api/internal/model"
	"ova-serial-api/internal/utils"
)

var _ = Describe("Slice", func() {
	Context("Split int slice into chunks", func() {
		slice := []int{5, 1, 2, 7, -1}

		It("should return correct result for empty slice", func() {
			zeroSlice := make([]int, 0)
			Expect(utils.SplitSlice(zeroSlice, 1)).To(Equal(make([][]int, 0)))
		})

		It("should return correct result for zero chunkSize", func() {
			Expect(utils.SplitSlice(slice, 0)).To(Equal([][]int{{5}, {1}, {2}, {7}, {-1}}))
		})

		It("should return correct result for chunkSize equal to 1", func() {
			Expect(utils.SplitSlice(slice, 1)).To(Equal([][]int{{5}, {1}, {2}, {7}, {-1}}))
		})

		It("should return correct result for chunkSize equal to 2", func() {
			Expect(utils.SplitSlice(slice, 2)).To(Equal([][]int{{5, 1}, {2, 7}, {-1}}))
		})

		It("should return correct result for chunkSize equal to slice length minus 1", func() {
			Expect(utils.SplitSlice(slice, 4)).To(Equal([][]int{{5, 1, 2, 7}, {-1}}))
		})

		It("should return correct result for chunkSize equal to slice length", func() {
			Expect(utils.SplitSlice(slice, 5)).To(Equal([][]int{{5, 1, 2, 7, -1}}))
		})

		It("should return correct result for chunkSize equal to slice length plus 1", func() {
			Expect(utils.SplitSlice(slice, 6)).To(Equal([][]int{{5, 1, 2, 7, -1}}))
		})
	})

	Context("Split Serial slice into chunks", func() {
		serial1 := model.Serial{UserID: 1, Title: "Friends", Genre: "comedy", Year: 1994, Seasons: 10}
		serial2 := model.Serial{UserID: 2, Title: "Game of Thrones", Genre: "fantasy", Year: 2011, Seasons: 8}
		serial3 := model.Serial{UserID: 3, Title: "Breaking Bad", Genre: "criminal", Year: 2008, Seasons: 5}
		serial4 := model.Serial{UserID: 4, Title: "The Big Bang Theory", Genre: "comedy", Year: 2007, Seasons: 12}

		slice := []model.Serial{
			serial1,
			serial2,
			serial3,
			serial4,
		}

		It("should return correct result for empty slice", func() {
			zeroSlice := make([]model.Serial, 0)
			Expect(utils.SplitSerialSlice(zeroSlice, 1)).To(Equal(make([][]model.Serial, 0)))
		})

		It("should return correct result for zero chunkSize", func() {
			Expect(utils.SplitSerialSlice(slice, 0)).To(Equal([][]model.Serial{
				{serial1}, {serial2}, {serial3}, {serial4},
			}))
		})

		It("should return correct result for chunkSize equal to 1", func() {
			Expect(utils.SplitSerialSlice(slice, 1)).To(Equal([][]model.Serial{
				{serial1}, {serial2}, {serial3}, {serial4},
			}))
		})

		It("should return correct result for chunkSize equal to 2", func() {
			Expect(utils.SplitSerialSlice(slice, 2)).To(Equal([][]model.Serial{
				{serial1, serial2}, {serial3, serial4},
			}))
		})

		It("should return correct result for chunkSize equal to slice length minus 1", func() {
			Expect(utils.SplitSerialSlice(slice, 3)).To(Equal([][]model.Serial{
				{serial1, serial2, serial3}, {serial4},
			}))
		})

		It("should return correct result for chunkSize equal to slice length", func() {
			Expect(utils.SplitSerialSlice(slice, 4)).To(Equal([][]model.Serial{slice}))
		})

		It("should return correct result for chunkSize equal to slice length plus 1", func() {
			Expect(utils.SplitSerialSlice(slice, 5)).To(Equal([][]model.Serial{slice}))
		})
	})

	Context("Filter int slice", func() {
		forbiddenValues := []int{-2, 1, 2, 5}

		It("should return correct result for empty slice", func() {
			zeroSlice := make([]int, 0)
			Expect(utils.FilterIntSlice(zeroSlice, forbiddenValues)).To(Equal(make([]int, 0)))
		})

		It("should return correct result with empty forbidden values", func() {
			slice := []int{1, 2, 3}
			Expect(utils.FilterIntSlice(slice, []int{})).To(Equal(slice))
		})

		It("should return correct result with no forbidden values in slice", func() {
			slice := []int{0, 3, 4, 6}
			Expect(utils.FilterIntSlice(slice, forbiddenValues)).To(Equal(slice))
		})

		It("should return correct result with 1 forbidden value in slice", func() {
			slice := []int{1, 3, 4, 6}
			Expect(utils.FilterIntSlice(slice, forbiddenValues)).To(Equal([]int{3, 4, 6}))
		})

		It("should return correct result with 2 forbidden values in slice", func() {
			slice := []int{0, 2, 4, 5}
			Expect(utils.FilterIntSlice(slice, forbiddenValues)).To(Equal([]int{0, 4}))
		})

		It("should return correct result with all values forbidden in slice", func() {
			slice := []int{2, -2}
			Expect(utils.FilterIntSlice(slice, forbiddenValues)).To(Equal(make([]int, 0)))
		})
	})
})
