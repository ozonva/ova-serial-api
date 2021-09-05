package api_test

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	api "ova-serial-api/internal/api"
	mock "ova-serial-api/internal/mocks"
	"ova-serial-api/internal/model"
	"ova-serial-api/internal/repo"
	desc "ova-serial-api/pkg/ova-serial-api"
)

var _ = Describe("Internal/Api/Service", func() {
	var (
		repoMock *mock.MockRepo
	)

	BeforeEach(func() {
		ctrl := gomock.NewController(GinkgoT())
		repoMock = mock.NewMockRepo(ctrl)
		defer ctrl.Finish()
	})

	Describe("Create serial", func() {
		Context("all is ok", func() {
			It("should return serial", func() {
				apiInstance := api.NewSerialAPI(repoMock)

				serial := model.Serial{
					UserID:  1,
					Title:   "Title1",
					Genre:   "genre1",
					Year:    2000,
					Seasons: 2,
				}

				repoMock.EXPECT().AddEntity(gomock.Eq(serial)).Return(int64(1), nil)

				ctx := context.Background()

				request := desc.CreateSerialRequestV1{
					UserId:  serial.UserID,
					Title:   serial.Title,
					Genre:   serial.Genre,
					Year:    serial.Year,
					Seasons: serial.Seasons,
				}

				response, err := apiInstance.CreateSerialV1(ctx, &request)

				Expect(response.Id).To(Equal(int64(1)))
				Expect(err).To(BeNil())
			})
		})

		Context("repo returns error", func() {
			It("should return internal error", func() {
				apiInstance := api.NewSerialAPI(repoMock)

				serial := model.Serial{
					UserID:  1,
					Title:   "Title2",
					Genre:   "genre2",
					Year:    2002,
					Seasons: 3,
				}

				repoMock.EXPECT().AddEntity(gomock.Eq(serial)).Return(int64(0), errors.New("test error"))

				ctx := context.Background()

				request := desc.CreateSerialRequestV1{
					UserId:  serial.UserID,
					Title:   serial.Title,
					Genre:   serial.Genre,
					Year:    serial.Year,
					Seasons: serial.Seasons,
				}

				_, err := apiInstance.CreateSerialV1(ctx, &request)

				Expect(err.Error()).To(Equal("rpc error: code = Internal desc = internal error"))
			})
		})
	})

	Describe("Get serial", func() {
		Context("all is ok", func() {
			It("should return serial", func() {
				apiInstance := api.NewSerialAPI(repoMock)

				serial := model.Serial{
					ID:      1,
					UserID:  1,
					Title:   "Title2",
					Genre:   "genre2",
					Year:    2002,
					Seasons: 3,
				}

				repoMock.EXPECT().GetEntity(gomock.Eq(serial.ID)).Return(&serial, nil)

				ctx := context.Background()

				request := desc.GetSerialRequestV1{
					Id: serial.ID,
				}

				response, err := apiInstance.GetSerialV1(ctx, &request)

				Expect(response.Serial.Id).To(Equal(serial.ID))
				Expect(response.Serial.UserId).To(Equal(serial.UserID))
				Expect(response.Serial.Title).To(Equal(serial.Title))
				Expect(response.Serial.Genre).To(Equal(serial.Genre))
				Expect(response.Serial.Year).To(Equal(serial.Year))
				Expect(response.Serial.Seasons).To(Equal(serial.Seasons))
				Expect(err).To(BeNil())
			})
		})

		Context("repo returns error", func() {
			It("should return error", func() {
				apiInstance := api.NewSerialAPI(repoMock)

				serial := model.Serial{
					ID:      1,
					UserID:  1,
					Title:   "Title2",
					Genre:   "genre2",
					Year:    2002,
					Seasons: 3,
				}

				repoMock.EXPECT().GetEntity(gomock.Eq(serial.ID)).Return(nil, errors.New("test error"))

				ctx := context.Background()

				request := desc.GetSerialRequestV1{
					Id: serial.ID,
				}

				_, err := apiInstance.GetSerialV1(ctx, &request)

				Expect(err.Error()).To(Equal("rpc error: code = Internal desc = internal error"))
			})
		})

		Context("row not found", func() {
			It("should return error", func() {
				apiInstance := api.NewSerialAPI(repoMock)

				serial := model.Serial{
					ID:      1,
					UserID:  1,
					Title:   "Title2",
					Genre:   "genre2",
					Year:    2002,
					Seasons: 3,
				}

				repoMock.EXPECT().GetEntity(gomock.Eq(serial.ID)).Return(nil, &repo.NotFound{})

				ctx := context.Background()

				request := desc.GetSerialRequestV1{
					Id: serial.ID,
				}

				_, err := apiInstance.GetSerialV1(ctx, &request)

				Expect(err.Error()).To(Equal("rpc error: code = NotFound desc = not found"))
			})
		})
	})

	Describe("List serials", func() {
		Context("all is ok", func() {
			It("should return serials", func() {
				apiInstance := api.NewSerialAPI(repoMock)

				serial := model.Serial{
					ID:      1,
					UserID:  1,
					Title:   "Title2",
					Genre:   "genre2",
					Year:    2002,
					Seasons: 3,
				}

				gomock.InOrder(
					repoMock.EXPECT().ListEntities(uint64(1), uint64(0)).Return([]model.Serial{serial}, nil),
				)

				ctx := context.Background()

				request := desc.ListSerialsRequestV1{
					Limit:  1,
					Offset: 0,
				}

				response, err := apiInstance.ListSerialsV1(ctx, &request)

				Expect(response.Serials[0].Id).To(Equal(serial.ID))
				Expect(response.Serials[0].UserId).To(Equal(serial.UserID))
				Expect(response.Serials[0].Title).To(Equal(serial.Title))
				Expect(response.Serials[0].Genre).To(Equal(serial.Genre))
				Expect(response.Serials[0].Year).To(Equal(serial.Year))
				Expect(response.Serials[0].Seasons).To(Equal(serial.Seasons))
				Expect(err).To(BeNil())
			})
		})

		Context("repo returns error", func() {
			It("should return error", func() {
				apiInstance := api.NewSerialAPI(repoMock)

				gomock.InOrder(
					repoMock.EXPECT().ListEntities(uint64(1), uint64(0)).Return(nil, errors.New("test error")),
				)

				ctx := context.Background()

				request := desc.ListSerialsRequestV1{
					Limit:  1,
					Offset: 0,
				}

				_, err := apiInstance.ListSerialsV1(ctx, &request)

				Expect(err.Error()).To(Equal("rpc error: code = Internal desc = internal error"))
			})
		})
	})

	Describe("Delete serial", func() {
		Context("all is ok", func() {
			It("should not return error", func() {
				apiInstance := api.NewSerialAPI(repoMock)

				serial := model.Serial{
					ID:      1,
					UserID:  1,
					Title:   "Title2",
					Genre:   "genre2",
					Year:    2002,
					Seasons: 3,
				}

				repoMock.EXPECT().RemoveEntity(gomock.Eq(serial.ID)).Return(nil)

				ctx := context.Background()

				request := desc.RemoveSerialRequestV1{
					Id: serial.ID,
				}

				_, err := apiInstance.RemoveSerialV1(ctx, &request)

				Expect(err).To(BeNil())
			})
		})
	})

	Context("repo returns error", func() {
		It("should return error", func() {
			apiInstance := api.NewSerialAPI(repoMock)

			serial := model.Serial{
				ID:      1,
				UserID:  1,
				Title:   "Title1",
				Genre:   "genre1",
				Year:    2000,
				Seasons: 2,
			}

			repoMock.EXPECT().RemoveEntity(gomock.Eq(serial.ID)).Return(errors.New("test error"))

			ctx := context.Background()

			request := desc.RemoveSerialRequestV1{
				Id: serial.ID,
			}

			_, err := apiInstance.RemoveSerialV1(ctx, &request)

			Expect(err.Error()).To(Equal("rpc error: code = Internal desc = internal error"))
		})
	})

	Context("row not found", func() {
		It("should return error", func() {
			apiInstance := api.NewSerialAPI(repoMock)

			serial := model.Serial{
				ID:      1,
				UserID:  1,
				Title:   "Title1",
				Genre:   "genre1",
				Year:    2000,
				Seasons: 2,
			}

			repoMock.EXPECT().RemoveEntity(gomock.Eq(serial.ID)).Return(&repo.NotFound{})

			ctx := context.Background()

			request := desc.RemoveSerialRequestV1{
				Id: serial.ID,
			}

			_, err := apiInstance.RemoveSerialV1(ctx, &request)

			Expect(err.Error()).To(Equal("rpc error: code = NotFound desc = not found"))
		})
	})
})
