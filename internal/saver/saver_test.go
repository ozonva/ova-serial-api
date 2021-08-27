package saver_test

import (
	"errors"
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
	serial3 := model.Serial{UserID: 2, Title: "Test3", Genre: "genre3", Year: 2002, Seasons: 2}
	serial4 := model.Serial{UserID: 2, Title: "Test4", Genre: "genre4", Year: 2003, Seasons: 1}

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockFlusher = mock.NewMockFlusher(mockCtrl)
	})

	Context("Create instance", func() {
		It("should create new instance and flush by timeout", func() {
			var testEntities []model.Serial

			saver := saver.NewSaver(10, mockFlusher, 1)
			Expect(saver).NotTo(BeNil())

			mockFlusher.EXPECT().Flush(gomock.Len(1)).
				DoAndReturn(func(entities []model.Serial) []model.Serial {
					testEntities = append(testEntities, entities...)
					return nil
				}).
				AnyTimes()

			err := saver.Save(serial1)
			Expect(err).To(BeNil())

			time.Sleep(2 * time.Second)

			Expect(testEntities).To(Equal([]model.Serial{serial1}))
		})
	})

	Context("Save", func() {
		It("should save with free space", func() {
			var testEntities []model.Serial
			saver := saver.NewSaver(3, mockFlusher, 1)

			mockFlusher.EXPECT().Flush(gomock.Len(2)).
				DoAndReturn(func(entities []model.Serial) []model.Serial {
					testEntities = append(testEntities, entities...)
					return nil
				}).
				Times(1)

			// Add 2 entities
			err := saver.Save(serial1)
			Expect(err).To(BeNil())
			err = saver.Save(serial2)
			Expect(err).To(BeNil())

			Expect(testEntities).To(BeNil())
		})

		It("should flush with no free space", func() {
			var testEntities []model.Serial
			saver := saver.NewSaver(1, mockFlusher, 1)

			mockFlusher.EXPECT().Flush(gomock.Len(1)).
				DoAndReturn(func(entities []model.Serial) []model.Serial {
					testEntities = append(testEntities, entities...)
					return nil
				}).
				Times(1)

			// Add 2 entities
			err := saver.Save(serial1)
			Expect(err).To(BeNil())
			err = saver.Save(serial2)
			Expect(err).To(BeNil())

			Expect(testEntities).To(Equal([]model.Serial{serial1}))
		})

		It("should flush with unsaved data", func() {
			var testEntities []model.Serial
			saver := saver.NewSaver(2, mockFlusher, 1)

			mockFlusher.EXPECT().Flush(gomock.Len(2)).
				DoAndReturn(func(entities []model.Serial) []model.Serial {
					testEntities = entities
					return []model.Serial{serial1}
				}).
				Times(1)

			// Add 3 serials
			err := saver.Save(serial1)
			Expect(err).To(BeNil())
			err = saver.Save(serial2)
			Expect(err).To(BeNil())
			err = saver.Save(serial3)
			Expect(err).To(BeNil())

			// Should have an attempt to flush 2 of them
			Expect(testEntities).To(Equal([]model.Serial{serial1, serial2}))

			mockFlusher.EXPECT().Flush(gomock.Len(2)).
				DoAndReturn(func(entities []model.Serial) []model.Serial {
					testEntities = entities
					return nil
				}).
				Times(1)

			// Add 4th one
			err = saver.Save(serial4)
			Expect(err).To(BeNil())

			// Should have an attempt to flush the last and previously not flushed one
			Expect(testEntities).To(Equal([]model.Serial{serial1, serial3}))
		})

		It("should throw error when no free space after flush", func() {
			var testEntities []model.Serial
			saver := saver.NewSaver(1, mockFlusher, 1)

			mockFlusher.EXPECT().Flush(gomock.Len(1)).
				DoAndReturn(func(entities []model.Serial) []model.Serial {
					testEntities = append(testEntities, entities...)
					return []model.Serial{serial1}
				}).
				Times(1)

			err := saver.Save(serial1)
			Expect(err).To(BeNil())
			err = saver.Save(serial2)
			Expect(err).To(Equal(errors.New("no capacity in storage")))

			Expect(testEntities).To(Equal([]model.Serial{serial1}))
		})
	})

	Context("Close", func() {
		It("should not flush without data", func() {
			var testEntities []model.Serial
			saver := saver.NewSaver(10, mockFlusher, 5)

			mockFlusher.EXPECT().Flush(gomock.Any()).
				DoAndReturn(func(entities []model.Serial) []model.Serial {
					testEntities = append(testEntities, entities...)
					return []model.Serial{serial1}
				}).
				Times(1)

			saver.Close()

			Expect(testEntities).To(BeNil())
		})

		It("should flush with data", func() {
			var testEntities []model.Serial
			saver := saver.NewSaver(10, mockFlusher, 5)

			err := saver.Save(serial1)
			Expect(err).To(BeNil())

			mockFlusher.EXPECT().Flush(gomock.Any()).
				DoAndReturn(func(entities []model.Serial) []model.Serial {
					testEntities = append(testEntities, entities...)
					return []model.Serial{serial1}
				}).
				Times(1)

			saver.Close()

			Expect(testEntities).To(Equal([]model.Serial{serial1}))
		})
	})
})
