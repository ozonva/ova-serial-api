package flusher_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"ova-serial-api/internal/flusher"
	mock_repo "ova-serial-api/internal/mocks"
	"ova-serial-api/internal/model"
)

var _ = Describe("Flusher", func() {
	var (
		mockCtrl    *gomock.Controller
		mockRepo    *mock_repo.MockRepo
		testFlusher flusher.Flusher
	)

	serials := []model.Serial{
		{UserID: 1, Title: "Friends", Genre: "comedy", Year: 1994, Seasons: 10},
		{UserID: 2, Title: "Game of Thrones", Genre: "fantasy", Year: 2011, Seasons: 8},
		{UserID: 3, Title: "Breaking Bad", Genre: "criminal", Year: 2008, Seasons: 5},
		{UserID: 4, Title: "The Big Bang Theory", Genre: "comedy", Year: 2007, Seasons: 12},
	}

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockRepo = mock_repo.NewMockRepo(mockCtrl)
	})

	Context("Flush with no errors", func() {
		It("should return nil for empty slice", func() {
			testFlusher = flusher.NewFlusher(
				1,
				mockRepo,
			)
			zeroSlice := make([]model.Serial, 0)
			Expect(testFlusher.Flush(zeroSlice)).Should(BeNil())
		})

		It("should return nil for zero chunkSize", func() {
			testFlusher = flusher.NewFlusher(
				0,
				mockRepo,
			)
			mockRepo.EXPECT().AddEntities(gomock.Len(1)).Times(4).Return(nil)
			Expect(testFlusher.Flush(serials)).Should(BeNil())
		})

		It("should return nil for chunkSize equal to 1", func() {
			testFlusher = flusher.NewFlusher(
				1,
				mockRepo,
			)
			mockRepo.EXPECT().AddEntities(gomock.Len(1)).Times(4).Return(nil)
			Expect(testFlusher.Flush(serials)).Should(BeNil())
		})

		It("should return nil for chunkSize equal to 2", func() {
			testFlusher = flusher.NewFlusher(
				2,
				mockRepo,
			)
			mockRepo.EXPECT().AddEntities(gomock.Len(2)).Times(2).Return(nil)
			Expect(testFlusher.Flush(serials)).Should(BeNil())
		})

		It("should return nil for chunkSize equal to slice length minus 1", func() {
			testFlusher = flusher.NewFlusher(
				3,
				mockRepo,
			)
			gomock.InOrder(
				mockRepo.EXPECT().AddEntities(gomock.Len(3)).Return(nil),
				mockRepo.EXPECT().AddEntities(gomock.Len(1)).Return(nil),
			)
			Expect(testFlusher.Flush(serials)).Should(BeNil())
		})

		It("should return nil for chunkSize equal to slice length", func() {
			testFlusher = flusher.NewFlusher(
				4,
				mockRepo,
			)
			mockRepo.EXPECT().AddEntities(gomock.Len(4)).Return(nil)
			Expect(testFlusher.Flush(serials)).Should(BeNil())
		})
	})

	Context("Flush with errors", func() {
		testError := errors.New("test error 123")

		It("should return chunk for error in the first chunk", func() {
			testFlusher = flusher.NewFlusher(
				3,
				mockRepo,
			)
			gomock.InOrder(
				mockRepo.EXPECT().AddEntities(gomock.Len(3)).Return(testError),
				mockRepo.EXPECT().AddEntities(gomock.Len(1)).Return(nil),
			)
			Expect(testFlusher.Flush(serials)).To(Equal(serials[:3]))
		})

		It("should return chunk for error in the second chunk", func() {
			testFlusher = flusher.NewFlusher(
				3,
				mockRepo,
			)
			gomock.InOrder(
				mockRepo.EXPECT().AddEntities(gomock.Len(3)).Return(nil),
				mockRepo.EXPECT().AddEntities(gomock.Len(1)).Return(testError),
			)
			Expect(testFlusher.Flush(serials)).To(Equal(serials[3:]))
		})

		It("should return chunk for error in both chunks", func() {
			testFlusher = flusher.NewFlusher(
				3,
				mockRepo,
			)
			gomock.InOrder(
				mockRepo.EXPECT().AddEntities(gomock.Len(3)).Return(testError),
				mockRepo.EXPECT().AddEntities(gomock.Len(1)).Return(testError),
			)
			Expect(testFlusher.Flush(serials)).To(Equal(serials))
		})
	})
})
