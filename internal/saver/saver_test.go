package saver_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	mock "ova-serial-api/internal/mocks"
	"ova-serial-api/internal/model"
	"ova-serial-api/internal/saver"
	"time"
)

var _ = Describe("Saver", func() {
	var (
		mockCtrl    *gomock.Controller
		mockFlusher *mock.MockFlusher
	)

	serial1 := model.Serial{UserID: 0, Title: "Test1", Genre: "genre1", Year: 2000, Seasons: 4}
	serial2 := model.Serial{UserID: 1, Title: "Test2", Genre: "genre2", Year: 2001, Seasons: 3}

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockFlusher = mock.NewMockFlusher(mockCtrl)
	})

	Context("Create instance", func() {
		It("should create new instance and flush by timeout", func() {
			var testEntities []model.Serial

			saver := saver.NewSaver(10, mockFlusher, 1)
			Expect(saver).NotTo(BeNil())

			saver.Init()

			mockFlusher.EXPECT().Flush(gomock.Len(1)).
				DoAndReturn(func(entities []model.Serial) []model.Serial {
					testEntities = append(testEntities, entities...)
					return nil
				}).
				AnyTimes()

			saver.Save(serial1)

			time.Sleep(2 * time.Second)

			Expect(testEntities).To(Equal([]model.Serial{serial1}))
		})
	})

	Context("Save", func() {
		It("should save with free space", func() {
			var testEntities []model.Serial
			saver := saver.NewSaver(3, mockFlusher, 1)
			saver.Init()

			mockFlusher.EXPECT().Flush(gomock.Len(2)).
				DoAndReturn(func(entities []model.Serial) []model.Serial {
					testEntities = append(testEntities, entities...)
					return nil
				}).
				Times(1)

			// Add 2 entities
			saver.Save(serial1)
			saver.Save(serial2)

			Expect(testEntities).To(BeNil())
		})
	})

	Context("Close", func() {
		It("should not flush without data", func() {
			var testEntities []model.Serial
			saver := saver.NewSaver(10, mockFlusher, 5)
			saver.Init()

			mockFlusher.EXPECT().Flush(gomock.Any()).
				DoAndReturn(func(entities []model.Serial) []model.Serial {
					testEntities = append(testEntities, entities...)
					return []model.Serial{serial1}
				}).
				Times(1)

			saver.Close()

			Expect(testEntities).To(BeNil())
		})
	})
})
