package service_test

import (
	"campaign-service/internal/entity"
	"campaign-service/internal/infrastructure/database"
	"campaign-service/internal/model"
	claim_item_model "campaign-service/internal/model/claim_item"
	claim_order_model "campaign-service/internal/model/claim_order"
	repository_mock "campaign-service/internal/repository/tests/mocks"
	"campaign-service/internal/service"
	"campaign-service/internal/tests"
	"campaign-service/internal/util/error_wrapper"
	campaignv1 "campaign-service/pkg/api/proto/campaign/v1"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
)

func TestCreateClaimOrder(t *testing.T) {
	testTools := tests.SetupSuite(t, tests.SetupOptions{})
	defer testTools.Teardown(t)

	type expected struct {
		tests.Expected
		order      claim_order_model.ClaimOrder
		claimItems []claim_item_model.ClaimItem
		address    *model.DeliveryAddress
	}

	mockClaimRepo := repository_mock.MockClaimRepository{
		GenerateOrderId: &[]string{"order1"}[0],
		GenerateClaimItemId: []string{
			"claim1",
			"claim2",
		},
	}

	mockTimeRepo := repository_mock.MockTimeRepositoy{
		CurrentTime: func() *time.Time {
			mockT, err := time.Parse(time.RFC3339, "2022-10-02T00:00:00+07:00")
			if err != nil {
				t.Fatal(err)
			}
			return &mockT
		}(),
	}

	suite := tests.Testcase{
		"001.Successfully": {
			Mock: &tests.Mock{
				MockFilesPath: []string{
					"./mocks/campaigns_and_item_rewards.json",
					"./mocks/user_point.json",
				},
				ClaimRepository: mockClaimRepo,
				TimeRepository:  mockTimeRepo,
			},
			Input: entity.CreateClaimOrderInput{
				UserId: "R0001",
				Req: &campaignv1.CreateClaimOrderRequest{
					CampaignId: "campaign01",
					Rewards: []*campaignv1.RewardQuantity{
						{
							RewardId: "reward01",
							Quantity: 1,
						},
						{
							RewardId: "reward02",
							Quantity: 1,
						},
					},
					Address: &campaignv1.Address{
						AddressNumber: &[]string{"11/12"}[0],
						Moo:           &[]string{"1"}[0],
						Building:      &[]string{"สูงส่ง"}[0],
						Floor:         &[]string{"10"}[0],
						Room:          &[]string{"1001"}[0],
						Soi:           &[]string{"ต้นทาง"}[0],
						Road:          &[]string{"เรียบทาง"}[0],
						SubDistrict:   &[]string{"ท่าจั่น"}[0],
						District:      &[]string{"เมือง"}[0],
						Province:      &[]string{"กทม"}[0],
						ZipCode:       &[]string{"10002"}[0],
						Country:       &[]string{"Thailand"}[0],
					},
				},
			},
			Expected: expected{
				order: claim_order_model.ClaimOrder{
					ID:          "order1",
					UserId:      "R0001",
					CampaignId:  "campaign01",
					Quantity:    2,
					Point:       3000,
					Status:      claim_order_model.CREATED,
					FullAddress: "บ้านเลขที่ 11/12 หมู่ที่ 1 หมู่บ้าน/อาคาร สูงส่ง ห้อง 1001 ชั้น 10 ซอย ต้นทาง ถนน เรียบทาง แขวง/ตำบล ท่าจั่น เขต/อำเภอ เมือง จังหวัด กทม รหัสไปรษณีย์ 10002",
				},
				claimItems: []claim_item_model.ClaimItem{
					{
						ID:           "claim1",
						ClaimOrderId: "order1",
						RewardId:     "reward01",
						Quantity:     1,
						Point:        1000,
						Status:       claim_item_model.CREATED,
					},
					{
						ID:           "claim2",
						ClaimOrderId: "order1",
						RewardId:     "reward02",
						Quantity:     1,
						Point:        2000,
						Status:       claim_item_model.CREATED,
					},
				},
				address: &model.DeliveryAddress{
					UserId:        "R0001",
					AddressNumber: &[]string{"11/12"}[0],
					Moo:           &[]string{"1"}[0],
					Building:      &[]string{"สูงส่ง"}[0],
					Floor:         &[]string{"10"}[0],
					Room:          &[]string{"1001"}[0],
					Soi:           &[]string{"ต้นทาง"}[0],
					Road:          &[]string{"เรียบทาง"}[0],
					SubDistrict:   &[]string{"ท่าจั่น"}[0],
					District:      &[]string{"เมือง"}[0],
					Province:      &[]string{"กทม"}[0],
					ZipCode:       &[]string{"10002"}[0],
					Country:       &[]string{"Thailand"}[0],
					FullAddress:   "บ้านเลขที่ 11/12 หมู่ที่ 1 หมู่บ้าน/อาคาร สูงส่ง ห้อง 1001 ชั้น 10 ซอย ต้นทาง ถนน เรียบทาง แขวง/ตำบล ท่าจั่น เขต/อำเภอ เมือง จังหวัด กทม รหัสไปรษณีย์ 10002",
				},
			},
		},

		"002.Insufficient Stock": {
			Mock: &tests.Mock{
				MockFilesPath: []string{
					"./mocks/campaigns_and_item_rewards.json",
					"./mocks/user_point.json",
					"./mocks/delivery_address.json",
				},
				ClaimRepository: mockClaimRepo,
				TimeRepository:  mockTimeRepo,
			},
			Input: entity.CreateClaimOrderInput{
				UserId: "R0001",
				Req: &campaignv1.CreateClaimOrderRequest{
					CampaignId: "campaign01",
					Rewards: []*campaignv1.RewardQuantity{
						{
							RewardId: "reward01",
							Quantity: 100,
						},
						{
							RewardId: "reward02",
							Quantity: 1,
						},
					},
					Address: &campaignv1.Address{
						Id: &[]string{"addressUserR0001"}[0],
					},
				},
			},
			Expected: expected{
				Expected: tests.Expected{
					Error: error_wrapper.GrpcErrorWrapper(
						codes.FailedPrecondition,
						error_wrapper.InternalError[error_wrapper.INSUFFICIENT_STOCK].Error(),
						error_wrapper.InternalError[error_wrapper.INSUFFICIENT_STOCK],
						error_wrapper.CreateGrpcErrorDetail(error_wrapper.INSUFFICIENT_STOCK),
					),
				},
				claimItems: []claim_item_model.ClaimItem{},
			},
		},

		"003.Insufficient Balance": {
			Mock: &tests.Mock{
				MockFilesPath: []string{
					"./mocks/campaigns_and_item_rewards.json",
					"./mocks/user_point.json",
					"./mocks/delivery_address.json",
				},
				ClaimRepository: mockClaimRepo,
				TimeRepository:  mockTimeRepo,
			},
			Input: entity.CreateClaimOrderInput{
				UserId: "R0002",
				Req: &campaignv1.CreateClaimOrderRequest{
					CampaignId: "campaign01",
					Rewards: []*campaignv1.RewardQuantity{
						{
							RewardId: "reward01",
							Quantity: 1,
						},
						{
							RewardId: "reward02",
							Quantity: 1,
						},
					},
					Address: &campaignv1.Address{
						Id: &[]string{"addressUserR0001"}[0],
					},
				},
			},
			Expected: expected{
				Expected: tests.Expected{
					Error: error_wrapper.GrpcErrorWrapper(
						codes.FailedPrecondition,
						error_wrapper.InternalError[error_wrapper.INSUFFICIENT_BALANCE].Error(),
						error_wrapper.InternalError[error_wrapper.INSUFFICIENT_BALANCE],
						error_wrapper.CreateGrpcErrorDetail(error_wrapper.INSUFFICIENT_BALANCE),
					),
				},
				claimItems: []claim_item_model.ClaimItem{},
			},
		},

		"004.1.Limitation Exceeded": {
			Mock: &tests.Mock{
				MockFilesPath: []string{
					"./mocks/campaigns_and_item_rewards.json",
					"./mocks/user_point.json",
					"./mocks/delivery_address.json",
					"./mocks/claimed_order.json",
				},
				ClaimRepository: mockClaimRepo,
				TimeRepository:  mockTimeRepo,
			},
			Input: entity.CreateClaimOrderInput{
				UserId: "R0001",
				Req: &campaignv1.CreateClaimOrderRequest{
					CampaignId: "campaign01",
					Rewards: []*campaignv1.RewardQuantity{
						{
							RewardId: "reward01",
							Quantity: 5,
						},
						{
							RewardId: "reward02",
							Quantity: 1,
						},
					},
					Address: &campaignv1.Address{
						Id: &[]string{"addressUserR0001"}[0],
					},
				},
			},
			Expected: expected{
				Expected: tests.Expected{
					Error: error_wrapper.GrpcErrorWrapper(
						codes.FailedPrecondition,
						error_wrapper.InternalError[error_wrapper.LIMITATION_EXCEEDED].Error(),
						error_wrapper.InternalError[error_wrapper.LIMITATION_EXCEEDED],
						error_wrapper.CreateGrpcErrorDetail(error_wrapper.LIMITATION_EXCEEDED),
					),
				},
				claimItems: []claim_item_model.ClaimItem{},
			},
		},

		"004.2.Limit Exceeded when item reward unlimit": {
			Mock: &tests.Mock{
				MockFilesPath: []string{
					"./mocks/campaigns_unlimit_item_but_limit_redeem.json",
					"./mocks/user_point.json",
					"./mocks/delivery_address.json",
					"./mocks/claimed_order.json",
				},
				ClaimRepository: mockClaimRepo,
				TimeRepository:  mockTimeRepo,
			},
			Input: entity.CreateClaimOrderInput{
				UserId: "R0001",
				Req: &campaignv1.CreateClaimOrderRequest{
					CampaignId: "campaign01",
					Rewards: []*campaignv1.RewardQuantity{
						{
							RewardId: "reward01",
							Quantity: 5,
						},
						{
							RewardId: "reward02",
							Quantity: 1,
						},
					},
					Address: &campaignv1.Address{
						Id: &[]string{"addressUserR0001"}[0],
					},
				},
			},
			Expected: expected{
				Expected: tests.Expected{
					Error: error_wrapper.GrpcErrorWrapper(
						codes.FailedPrecondition,
						error_wrapper.InternalError[error_wrapper.LIMITATION_EXCEEDED].Error(),
						error_wrapper.InternalError[error_wrapper.LIMITATION_EXCEEDED],
						error_wrapper.CreateGrpcErrorDetail(error_wrapper.LIMITATION_EXCEEDED),
					),
				},
				claimItems: []claim_item_model.ClaimItem{},
			},
		},

		"005.Create Order with exist address": {
			Mock: &tests.Mock{
				MockFilesPath: []string{
					"./mocks/campaigns_and_item_rewards.json",
					"./mocks/user_point.json",
					"./mocks/delivery_address.json",
				},
				ClaimRepository: mockClaimRepo,
				TimeRepository:  mockTimeRepo,
			},
			Input: entity.CreateClaimOrderInput{
				UserId: "R0001",
				Req: &campaignv1.CreateClaimOrderRequest{
					CampaignId: "campaign01",
					Rewards: []*campaignv1.RewardQuantity{
						{
							RewardId: "reward01",
							Quantity: 1,
						},
						{
							RewardId: "reward02",
							Quantity: 1,
						},
					},
					Address: &campaignv1.Address{
						Id: &[]string{"addressUserR0001"}[0],
					},
				},
			},
			Expected: expected{
				order: claim_order_model.ClaimOrder{
					ID:          "order1",
					UserId:      "R0001",
					CampaignId:  "campaign01",
					Quantity:    2,
					Point:       3000,
					Status:      "CREATED",
					FullAddress: "11/12 แสนสุข อาคาร:สูงส่ง ชั้น:10 ห้อง:1001 หมู่:1 ซอย:ต้นทาง ถนน:เรียบทาง ตำบล/แขวง:ท่าจั่น อำเภอ/เขต:เมือง กทม 10002 Thailand",
				},
				claimItems: []claim_item_model.ClaimItem{
					{
						ID:           "claim1",
						ClaimOrderId: "order1",
						RewardId:     "reward01",
						Quantity:     1,
						Point:        1000,
						Status:       claim_item_model.CREATED,
					},
					{
						ID:           "claim2",
						ClaimOrderId: "order1",
						RewardId:     "reward02",
						Quantity:     1,
						Point:        2000,
						Status:       claim_item_model.CREATED,
					},
				},
				address: &model.DeliveryAddress{
					UserId:        "R0001",
					AddressNumber: &[]string{"11/12"}[0],
					Moo:           &[]string{"1"}[0],
					Building:      &[]string{"สูงส่ง"}[0],
					Floor:         &[]string{"10"}[0],
					Room:          &[]string{"1001"}[0],
					Soi:           &[]string{"ต้นทาง"}[0],
					Road:          &[]string{"เรียบทาง"}[0],
					SubDistrict:   &[]string{"ท่าจั่น"}[0],
					District:      &[]string{"เมือง"}[0],
					Province:      &[]string{"กทม"}[0],
					ZipCode:       &[]string{"10002"}[0],
					Country:       &[]string{"Thailand"}[0],
					FullAddress:   "11/12 แสนสุข อาคาร:สูงส่ง ชั้น:10 ห้อง:1001 หมู่:1 ซอย:ต้นทาง ถนน:เรียบทาง ตำบล/แขวง:ท่าจั่น อำเภอ/เขต:เมือง กทม 10002 Thailand",
				},
			},
		},

		"006.Create Order with update address": {
			Mock: &tests.Mock{
				MockFilesPath: []string{
					"./mocks/campaigns_and_item_rewards.json",
					"./mocks/user_point.json",
					"./mocks/delivery_address.json",
				},
				ClaimRepository: mockClaimRepo,
				TimeRepository:  mockTimeRepo,
			},
			Input: entity.CreateClaimOrderInput{
				UserId: "R0001",
				Req: &campaignv1.CreateClaimOrderRequest{
					CampaignId: "campaign01",
					Rewards: []*campaignv1.RewardQuantity{
						{
							RewardId: "reward01",
							Quantity: 1,
						},
						{
							RewardId: "reward02",
							Quantity: 1,
						},
					},
					Address: &campaignv1.Address{
						Id:            &[]string{"addressUserR0001"}[0],
						AddressNumber: &[]string{"101/11"}[0],
					},
				},
			},
			Expected: expected{
				order: claim_order_model.ClaimOrder{
					ID:          "order1",
					UserId:      "R0001",
					CampaignId:  "campaign01",
					Quantity:    2,
					Point:       3000,
					Status:      "CREATED",
					FullAddress: "บ้านเลขที่ 101/11 หมู่ที่ 1 หมู่บ้าน/อาคาร สูงส่ง ห้อง 1001 ชั้น 10 ซอย ต้นทาง ถนน เรียบทาง แขวง/ตำบล ท่าจั่น เขต/อำเภอ เมือง จังหวัด กทม รหัสไปรษณีย์ 10002",
				},
				claimItems: []claim_item_model.ClaimItem{
					{
						ID:           "claim1",
						ClaimOrderId: "order1",
						RewardId:     "reward01",
						Quantity:     1,
						Point:        1000,
						Status:       claim_item_model.CREATED,
					},
					{
						ID:           "claim2",
						ClaimOrderId: "order1",
						RewardId:     "reward02",
						Quantity:     1,
						Point:        2000,
						Status:       claim_item_model.CREATED,
					},
				},
				address: &model.DeliveryAddress{
					UserId:        "R0001",
					AddressNumber: &[]string{"101/11"}[0],
					Moo:           &[]string{"1"}[0],
					Building:      &[]string{"สูงส่ง"}[0],
					Floor:         &[]string{"10"}[0],
					Room:          &[]string{"1001"}[0],
					Soi:           &[]string{"ต้นทาง"}[0],
					Road:          &[]string{"เรียบทาง"}[0],
					SubDistrict:   &[]string{"ท่าจั่น"}[0],
					District:      &[]string{"เมือง"}[0],
					Province:      &[]string{"กทม"}[0],
					ZipCode:       &[]string{"10002"}[0],
					Country:       &[]string{"Thailand"}[0],
					FullAddress:   "บ้านเลขที่ 101/11 หมู่ที่ 1 หมู่บ้าน/อาคาร สูงส่ง ห้อง 1001 ชั้น 10 ซอย ต้นทาง ถนน เรียบทาง แขวง/ตำบล ท่าจั่น เขต/อำเภอ เมือง จังหวัด กทม รหัสไปรษณีย์ 10002",
				},
			},
		},

		"007.Create Order on unreach claim phase should error": {
			Mock: &tests.Mock{
				MockFilesPath: []string{
					"./mocks/campaigns_and_item_rewards.json",
					"./mocks/user_point.json",
					"./mocks/delivery_address.json",
				},
				ClaimRepository: mockClaimRepo,
				TimeRepository:  mockTimeRepo,
			},
			Input: entity.CreateClaimOrderInput{
				UserId: "R0001",
				Req: &campaignv1.CreateClaimOrderRequest{
					CampaignId: "unreachCampaign",
					Rewards: []*campaignv1.RewardQuantity{
						{
							RewardId: "reward03",
							Quantity: 1,
						},
					},
					Address: &campaignv1.Address{
						Id: &[]string{"addressUserR0001"}[0],
					},
				},
			},
			Expected: expected{
				Expected: tests.Expected{
					Error: error_wrapper.GrpcErrorWrapper(
						codes.FailedPrecondition,
						error_wrapper.InternalError[error_wrapper.CLAIM_PHASE_NOT_RUNNING].Error(),
						error_wrapper.InternalError[error_wrapper.CLAIM_PHASE_NOT_RUNNING],
						error_wrapper.CreateGrpcErrorDetail(error_wrapper.CLAIM_PHASE_NOT_RUNNING),
					),
				},
			},
		},

		"008.Create Order on expried claim phase should error": {
			Mock: &tests.Mock{
				MockFilesPath: []string{
					"./mocks/campaigns_and_item_rewards.json",
					"./mocks/user_point.json",
					"./mocks/delivery_address.json",
				},
				ClaimRepository: mockClaimRepo,
				TimeRepository:  mockTimeRepo,
			},
			Input: entity.CreateClaimOrderInput{
				UserId: "R0001",
				Req: &campaignv1.CreateClaimOrderRequest{
					CampaignId: "expiredCampaign",
					Rewards: []*campaignv1.RewardQuantity{
						{
							RewardId: "reward04",
							Quantity: 1,
						},
					},
					Address: &campaignv1.Address{
						Id: &[]string{"addressUserR0001"}[0],
					},
				},
			},
			Expected: expected{
				Expected: tests.Expected{
					Error: error_wrapper.GrpcErrorWrapper(
						codes.FailedPrecondition,
						error_wrapper.InternalError[error_wrapper.CLAIM_PHASE_NOT_RUNNING].Error(),
						error_wrapper.InternalError[error_wrapper.CLAIM_PHASE_NOT_RUNNING],
						error_wrapper.CreateGrpcErrorDetail(error_wrapper.CLAIM_PHASE_NOT_RUNNING),
					),
				},
			},
		},

		"009.Create Order with incorrect reward": {
			Mock: &tests.Mock{
				MockFilesPath: []string{
					"./mocks/campaigns_and_item_rewards.json",
					"./mocks/user_point.json",
					"./mocks/delivery_address.json",
				},
				ClaimRepository: mockClaimRepo,
				TimeRepository:  mockTimeRepo,
			},
			Input: entity.CreateClaimOrderInput{
				UserId: "R0001",
				Req: &campaignv1.CreateClaimOrderRequest{
					CampaignId: "campaign01",
					Rewards: []*campaignv1.RewardQuantity{
						{
							RewardId: "reward03",
							Quantity: 1,
						},
					},
					Address: &campaignv1.Address{
						Id: &[]string{"addressUserR0001"}[0],
					},
				},
			},
			Expected: expected{
				Expected: tests.Expected{
					Error: error_wrapper.GrpcErrorWrapper(
						codes.FailedPrecondition,
						error_wrapper.InternalError[error_wrapper.NOT_FOUND].Error(),
						error_wrapper.InternalError[error_wrapper.NOT_FOUND],
						error_wrapper.CreateGrpcErrorDetail(error_wrapper.NOT_FOUND),
					),
				},
			},
		},

		"010.Test Unlimit on limit reward": {
			Mock: &tests.Mock{
				MockFilesPath: []string{
					"./mocks/campaigns_and_item_rewards_limit_item.json",
					"./mocks/user_point_unlimit_item.json",
				},
				ClaimRepository: mockClaimRepo,
				TimeRepository:  mockTimeRepo,
			},
			Input: entity.CreateClaimOrderInput{
				UserId: "R0001",
				Req: &campaignv1.CreateClaimOrderRequest{
					CampaignId: "campaign01",
					Rewards: []*campaignv1.RewardQuantity{
						{
							RewardId: "reward01",
							Quantity: 1,
						},
						{
							RewardId: "reward02",
							Quantity: 3,
						},
					},
					Address: &campaignv1.Address{
						AddressNumber: &[]string{"11/12"}[0],
						Moo:           &[]string{"1"}[0],
						Village:       &[]string{"แสนสุข"}[0],
						Building:      &[]string{"สูงส่ง"}[0],
						Floor:         &[]string{"10"}[0],
						Room:          &[]string{"1001"}[0],
						Soi:           &[]string{"ต้นทาง"}[0],
						Road:          &[]string{"เรียบทาง"}[0],
						SubDistrict:   &[]string{"ท่าจั่น"}[0],
						District:      &[]string{"เมือง"}[0],
						Province:      &[]string{"กทม"}[0],
						ZipCode:       &[]string{"10002"}[0],
						Country:       &[]string{"Thailand"}[0],
					},
				},
			},
			Expected: expected{
				order: claim_order_model.ClaimOrder{
					ID:          "order1",
					UserId:      "R0001",
					CampaignId:  "campaign01",
					Quantity:    4,
					Point:       1300,
					Status:      claim_order_model.CREATED,
					FullAddress: "บ้านเลขที่ 11/12 หมู่ที่ 1 หมู่บ้าน/อาคาร สูงส่ง ห้อง 1001 ชั้น 10 ซอย ต้นทาง ถนน เรียบทาง แขวง/ตำบล ท่าจั่น เขต/อำเภอ เมือง จังหวัด กทม รหัสไปรษณีย์ 10002",
				},
				claimItems: []claim_item_model.ClaimItem{
					{
						ID:           "claim1",
						ClaimOrderId: "order1",
						RewardId:     "reward01",
						Quantity:     1,
						Point:        1000,
						Status:       claim_item_model.CREATED,
					},
					{
						ID:           "claim2",
						ClaimOrderId: "order1",
						RewardId:     "reward02",
						Quantity:     3,
						Point:        300,
						Status:       claim_item_model.CREATED,
					},
				},
				address: &model.DeliveryAddress{
					UserId:        "R0001",
					AddressNumber: &[]string{"11/12"}[0],
					Moo:           &[]string{"1"}[0],
					Building:      &[]string{"สูงส่ง"}[0],
					Floor:         &[]string{"10"}[0],
					Room:          &[]string{"1001"}[0],
					Soi:           &[]string{"ต้นทาง"}[0],
					Road:          &[]string{"เรียบทาง"}[0],
					SubDistrict:   &[]string{"ท่าจั่น"}[0],
					District:      &[]string{"เมือง"}[0],
					Province:      &[]string{"กทม"}[0],
					ZipCode:       &[]string{"10002"}[0],
					Country:       &[]string{"Thailand"}[0],
					FullAddress:   "บ้านเลขที่ 11/12 หมู่ที่ 1 หมู่บ้าน/อาคาร สูงส่ง ห้อง 1001 ชั้น 10 ซอย ต้นทาง ถนน เรียบทาง แขวง/ตำบล ท่าจั่น เขต/อำเภอ เมือง จังหวัด กทม รหัสไปรษณีย์ 10002",
				},
			},
		},
	}

	for name, tc := range suite {
		t.Run(name, func(t *testing.T) {
			testTools := tests.SetupTest(t, tests.SetupOptions{Mock: tc.Mock})
			defer testTools.Teardown(t)

			testTools.C.RequireInvoke(func(db *database.DB, cs service.IClaimService) {
				input := tc.Input.(entity.CreateClaimOrderInput)
				actual, err := cs.CreateClaimOrder(context.Background(), input)
				expected := tc.Expected.(expected)
				if expected.Error != nil {
					require.Error(t, err)
					assert.EqualError(t, err, expected.Error.Error())
					return
				} else {
					require.NoError(t, err)
					assert.NotEmpty(t, actual)
				}

				actualOrder := claim_order_model.ClaimOrder{}
				db.CampaignDB.Debug().Find(&actualOrder)
				actualOrder.CreatedAt = nil
				assert.Equal(t, expected.order, actualOrder)

				if len(expected.claimItems) > 0 {
					actualClaimItem := []claim_item_model.ClaimItem{}
					db.CampaignDB.Debug().Find(&actualClaimItem)
					assert.Equal(t, expected.claimItems, actualClaimItem)
				}

				if expected.address != nil {

					actualAddress := model.DeliveryAddress{}
					db.CampaignDB.Debug().Find(&actualAddress)
					actualAddress.ID = ""
					assert.Equal(t, *expected.address, actualAddress)
				}
			})
		})
	}
}

func TestSubmitClaimOrder(t *testing.T) {
	testTools := tests.SetupSuite(t, tests.SetupOptions{})
	defer testTools.Teardown(t)

	type expected struct {
		tests.Expected
		order       claim_order_model.ClaimOrder
		claimItems  []claim_item_model.ClaimItem
		userPoint   model.UserPoint
		rewardItems []model.RewardItem
	}

	suite := tests.Testcase{
		"001.Successfully": {
			Mock: &tests.Mock{
				MockFilesPath: []string{
					"./mocks/campaigns_and_item_rewards.json",
					"./mocks/created_order.json",
					"./mocks/user_point.json",
					"./mocks/claimed_order.json",
				},
			},
			Input: entity.SubmitClaimOrderInput{
				UserId: "R0001",
				Req: &campaignv1.SubmitClaimOrderRequest{
					Id: "order01",
				},
			},
			Expected: expected{
				order: claim_order_model.ClaimOrder{
					ID:          "order01",
					UserId:      "R0001",
					CampaignId:  "campaign01",
					Quantity:    6,
					Point:       9000,
					Status:      claim_order_model.CLAIMED,
					FullAddress: "11/12 แสนสุข อาคาร:สูงส่ง ชั้น:10 ห้อง:1001 หมู่:1 ซอย:ต้นทาง ถนน:เรียบทาง ตำบล/แขวง:ท่าจั่น อำเภอ/เขต:เมือง กทม 10002 Thailand",
				},
				claimItems: []claim_item_model.ClaimItem{
					{
						ID:           "claim03",
						ClaimOrderId: "order01",
						RewardId:     "reward01",
						Quantity:     3,
						Point:        3000,
						Status:       claim_item_model.CLAIMED,
					},
					{
						ID:           "claim04",
						ClaimOrderId: "order01",
						RewardId:     "reward02",
						Quantity:     3,
						Point:        6000,
						Status:       claim_item_model.CLAIMED,
					},
				},
				userPoint: model.UserPoint{
					UserId:       "R0001",
					Cid:          "cid001",
					CampaignId:   "campaign01",
					PointBalance: &[]float64{1000}[0],
					LatestPoint:  -9000,
					TxnRef:       &[]string{"order01"}[0],
					FromRef:      &[]string{model.CLAIMED_ORDER.String()}[0],
				},
				rewardItems: []model.RewardItem{
					{
						ID:              "reward01",
						CampaignId:      "campaign01",
						Limit:           true,
						TotalQuantity:   &[]int{10}[0],
						Remaining:       &[]int{7}[0],
						RedemptionLimit: &[]int{8}[0],
						RedemptionPoint: 1000,
					},
					{
						ID:              "reward02",
						CampaignId:      "campaign01",
						Limit:           false,
						TotalQuantity:   &[]int{0}[0],
						Remaining:       &[]int{0}[0],
						RedemptionLimit: &[]int{0}[0],
						RedemptionPoint: 2000,
					},
				},
			},
		},

		"002.Insufficient Stock": {
			Mock: &tests.Mock{
				MockFilesPath: []string{
					"./mocks/campaigns_and_item_rewards.json",
					"./mocks/user_point.json",
					"./mocks/claimed_order.json",
					"./mocks/created_order_insufficient_stock.json",
				},
			},
			Input: entity.SubmitClaimOrderInput{
				UserId: "R0001",
				Req: &campaignv1.SubmitClaimOrderRequest{
					Id: "order01",
				},
			},
			Expected: expected{
				order: claim_order_model.ClaimOrder{
					ID:          "order01",
					UserId:      "R0001",
					CampaignId:  "campaign01",
					Quantity:    101,
					Point:       102000,
					Status:      claim_order_model.REJECTED,
					FullAddress: "11/12 แสนสุข อาคาร:สูงส่ง ชั้น:10 ห้อง:1001 หมู่:1 ซอย:ต้นทาง ถนน:เรียบทาง ตำบล/แขวง:ท่าจั่น อำเภอ/เขต:เมือง กทม 10002 Thailand",
					ErrorMsg: &[]string{
						error_wrapper.GrpcErrorWrapper(
							codes.FailedPrecondition,
							error_wrapper.InternalError[error_wrapper.INSUFFICIENT_STOCK].Error(),
							error_wrapper.InternalError[error_wrapper.INSUFFICIENT_STOCK],
							error_wrapper.CreateGrpcErrorDetail(error_wrapper.INSUFFICIENT_STOCK),
						).Error(),
					}[0],
				},
				claimItems: []claim_item_model.ClaimItem{
					{
						ID:           "claim03",
						ClaimOrderId: "order01",
						RewardId:     "reward02",
						Quantity:     1,
						Point:        2000,
						Status:       claim_item_model.REJECTED,
					},
					{
						ID:           "claim04",
						ClaimOrderId: "order01",
						RewardId:     "reward01",
						Quantity:     100,
						Point:        100000,
						Status:       claim_item_model.REJECTED,
						ErrorMsg: &[]string{
							error_wrapper.GrpcErrorWrapper(
								codes.FailedPrecondition,
								error_wrapper.InternalError[error_wrapper.INSUFFICIENT_STOCK].Error(),
								error_wrapper.InternalError[error_wrapper.INSUFFICIENT_STOCK],
								error_wrapper.CreateGrpcErrorDetail(error_wrapper.INSUFFICIENT_STOCK),
							).Error(),
						}[0],
					},
				},
				userPoint: model.UserPoint{
					UserId:       "R0001",
					Cid:          "cid001",
					CampaignId:   "campaign01",
					PointBalance: &[]float64{10000}[0],
					LatestPoint:  0,
					TxnRef:       &[]string{"portTxn01"}[0],
					FromRef:      &[]string{"A"}[0],
				},
				rewardItems: []model.RewardItem{
					{
						ID:              "reward01",
						CampaignId:      "campaign01",
						Limit:           true,
						TotalQuantity:   &[]int{10}[0],
						Remaining:       &[]int{10}[0],
						RedemptionLimit: &[]int{8}[0],
						RedemptionPoint: 1000,
					},
					{
						ID:              "reward02",
						CampaignId:      "campaign01",
						Limit:           false,
						TotalQuantity:   &[]int{0}[0],
						Remaining:       &[]int{0}[0],
						RedemptionLimit: &[]int{0}[0],
						RedemptionPoint: 2000,
					},
				},
			},
		},

		"003.Insufficient Balance": {
			Mock: &tests.Mock{
				MockFilesPath: []string{
					"./mocks/campaigns_and_item_rewards.json",
					"./mocks/user_point.json",
					"./mocks/claimed_order.json",
					"./mocks/created_order_insufficient_balance.json",
				},
			},
			Input: entity.SubmitClaimOrderInput{
				UserId: "R0001",
				Req: &campaignv1.SubmitClaimOrderRequest{
					Id: "order01",
				},
			},
			Expected: expected{
				order: claim_order_model.ClaimOrder{
					ID:          "order01",
					UserId:      "R0001",
					CampaignId:  "campaign01",
					Quantity:    102,
					Point:       202000,
					Status:      claim_order_model.REJECTED,
					FullAddress: "11/12 แสนสุข อาคาร:สูงส่ง ชั้น:10 ห้อง:1001 หมู่:1 ซอย:ต้นทาง ถนน:เรียบทาง ตำบล/แขวง:ท่าจั่น อำเภอ/เขต:เมือง กทม 10002 Thailand",
					ErrorMsg: &[]string{
						error_wrapper.GrpcErrorWrapper(
							codes.FailedPrecondition,
							error_wrapper.InternalError[error_wrapper.INSUFFICIENT_BALANCE].Error(),
							error_wrapper.InternalError[error_wrapper.INSUFFICIENT_BALANCE],
							error_wrapper.CreateGrpcErrorDetail(error_wrapper.INSUFFICIENT_BALANCE),
						).Error(),
					}[0],
				},
				claimItems: []claim_item_model.ClaimItem{
					{
						ID:           "claim03",
						ClaimOrderId: "order01",
						RewardId:     "reward01",
						Quantity:     2,
						Point:        2000,
						Status:       claim_item_model.REJECTED,
					},
					{
						ID:           "claim04",
						ClaimOrderId: "order01",
						RewardId:     "reward02",
						Quantity:     100,
						Point:        200000,
						Status:       claim_item_model.REJECTED,
						ErrorMsg: &[]string{
							error_wrapper.GrpcErrorWrapper(
								codes.FailedPrecondition,
								error_wrapper.InternalError[error_wrapper.INSUFFICIENT_BALANCE].Error(),
								error_wrapper.InternalError[error_wrapper.INSUFFICIENT_BALANCE],
								error_wrapper.CreateGrpcErrorDetail(error_wrapper.INSUFFICIENT_BALANCE),
							).Error(),
						}[0],
					},
				},
				userPoint: model.UserPoint{
					UserId:       "R0001",
					Cid:          "cid001",
					CampaignId:   "campaign01",
					PointBalance: &[]float64{10000}[0],
					LatestPoint:  0,
					TxnRef:       &[]string{"portTxn01"}[0],
					FromRef:      &[]string{"A"}[0],
				},
				rewardItems: []model.RewardItem{
					{
						ID:              "reward01",
						CampaignId:      "campaign01",
						Limit:           true,
						TotalQuantity:   &[]int{10}[0],
						Remaining:       &[]int{10}[0],
						RedemptionLimit: &[]int{8}[0],
						RedemptionPoint: 1000,
					},
					{
						ID:              "reward02",
						CampaignId:      "campaign01",
						Limit:           false,
						TotalQuantity:   &[]int{0}[0],
						Remaining:       &[]int{0}[0],
						RedemptionLimit: &[]int{0}[0],
						RedemptionPoint: 2000,
					},
				},
			},
		},

		"004.Limitation Exceeded": {
			Mock: &tests.Mock{
				MockFilesPath: []string{
					"./mocks/campaigns_and_item_rewards.json",
					"./mocks/user_point.json",
					"./mocks/claimed_order.json",
					"./mocks/created_order_limitation_exceeded.json",
				},
			},
			Input: entity.SubmitClaimOrderInput{
				UserId: "R0001",
				Req: &campaignv1.SubmitClaimOrderRequest{
					Id: "order01",
				},
			},
			Expected: expected{
				order: claim_order_model.ClaimOrder{
					ID:          "order01",
					UserId:      "R0001",
					CampaignId:  "campaign01",
					Quantity:    5,
					Point:       6000,
					Status:      claim_order_model.REJECTED,
					FullAddress: "11/12 แสนสุข อาคาร:สูงส่ง ชั้น:10 ห้อง:1001 หมู่:1 ซอย:ต้นทาง ถนน:เรียบทาง ตำบล/แขวง:ท่าจั่น อำเภอ/เขต:เมือง กทม 10002 Thailand",
					ErrorMsg: &[]string{
						error_wrapper.GrpcErrorWrapper(
							codes.FailedPrecondition,
							error_wrapper.InternalError[error_wrapper.LIMITATION_EXCEEDED].Error(),
							error_wrapper.InternalError[error_wrapper.LIMITATION_EXCEEDED],
							error_wrapper.CreateGrpcErrorDetail(error_wrapper.LIMITATION_EXCEEDED),
						).Error(),
					}[0],
				},
				claimItems: []claim_item_model.ClaimItem{
					{
						ID:           "claim03",
						ClaimOrderId: "order01",
						RewardId:     "reward01",
						Quantity:     4,
						Point:        4000,
						Status:       claim_item_model.REJECTED,
						ErrorMsg: &[]string{
							error_wrapper.GrpcErrorWrapper(
								codes.FailedPrecondition,
								error_wrapper.InternalError[error_wrapper.LIMITATION_EXCEEDED].Error(),
								error_wrapper.InternalError[error_wrapper.LIMITATION_EXCEEDED],
								error_wrapper.CreateGrpcErrorDetail(error_wrapper.LIMITATION_EXCEEDED),
							).Error(),
						}[0],
					},
					{
						ID:           "claim04",
						ClaimOrderId: "order01",
						RewardId:     "reward02",
						Quantity:     1,
						Point:        2000,
						Status:       claim_item_model.REJECTED,
					},
				},
				userPoint: model.UserPoint{
					UserId:       "R0001",
					Cid:          "cid001",
					CampaignId:   "campaign01",
					PointBalance: &[]float64{10000}[0],
					LatestPoint:  0,
					TxnRef:       &[]string{"portTxn01"}[0],
					FromRef:      &[]string{"A"}[0],
				},
				rewardItems: []model.RewardItem{
					{
						ID:              "reward01",
						CampaignId:      "campaign01",
						Limit:           true,
						TotalQuantity:   &[]int{10}[0],
						Remaining:       &[]int{10}[0],
						RedemptionLimit: &[]int{8}[0],
						RedemptionPoint: 1000,
					},
					{
						ID:              "reward02",
						CampaignId:      "campaign01",
						Limit:           false,
						TotalQuantity:   &[]int{0}[0],
						Remaining:       &[]int{0}[0],
						RedemptionLimit: &[]int{0}[0],
						RedemptionPoint: 2000,
					},
				},
			},
		},

		"005.Duplicate submit": {
			Mock: &tests.Mock{
				MockFilesPath: []string{
					"./mocks/campaigns_and_item_rewards.json",
					"./mocks/user_point.json",
					"./mocks/claimed_order.json",
				},
			},
			Input: entity.SubmitClaimOrderInput{
				UserId: "R0001",
				Req: &campaignv1.SubmitClaimOrderRequest{
					Id: "order00",
				},
			},
			Expected: expected{
				order: claim_order_model.ClaimOrder{
					ID:          "order00",
					UserId:      "R0001",
					CampaignId:  "campaign01",
					Quantity:    10,
					Point:       15000,
					Status:      claim_order_model.CLAIMED,
					FullAddress: "11/12 แสนสุข อาคาร:สูงส่ง ชั้น:10 ห้อง:1001 หมู่:1 ซอย:ต้นทาง ถนน:เรียบทาง ตำบล/แขวง:ท่าจั่น อำเภอ/เขต:เมือง กทม 10002 Thailand",
				},
				claimItems: []claim_item_model.ClaimItem{
					{
						ID:           "claim01",
						ClaimOrderId: "order00",
						RewardId:     "reward01",
						Quantity:     5,
						Point:        5000,
						Status:       claim_item_model.CLAIMED,
					},
					{
						ID:           "claim02",
						ClaimOrderId: "order00",
						RewardId:     "reward02",
						Quantity:     5,
						Point:        10000,
						Status:       claim_item_model.CLAIMED,
					},
				},
				userPoint: model.UserPoint{
					UserId:       "R0001",
					Cid:          "cid001",
					CampaignId:   "campaign01",
					PointBalance: &[]float64{10000}[0],
					LatestPoint:  0,
					TxnRef:       &[]string{"portTxn01"}[0],
					FromRef:      &[]string{"A"}[0],
				},
				rewardItems: []model.RewardItem{
					{
						ID:              "reward01",
						CampaignId:      "campaign01",
						Limit:           true,
						TotalQuantity:   &[]int{10}[0],
						Remaining:       &[]int{10}[0],
						RedemptionLimit: &[]int{8}[0],
						RedemptionPoint: 1000,
					},
					{
						ID:              "reward02",
						CampaignId:      "campaign01",
						Limit:           false,
						TotalQuantity:   &[]int{0}[0],
						Remaining:       &[]int{0}[0],
						RedemptionLimit: &[]int{0}[0],
						RedemptionPoint: 2000,
					},
				},
			},
		},

		"006.Submit other owner's order": {
			Mock: &tests.Mock{
				MockFilesPath: []string{
					"./mocks/campaigns_and_item_rewards.json",
					"./mocks/created_order.json",
					"./mocks/user_point.json",
					"./mocks/claimed_order.json",
				},
			},
			Input: entity.SubmitClaimOrderInput{
				UserId: "R0002",
				Req: &campaignv1.SubmitClaimOrderRequest{
					Id: "order01",
				},
			},
			Expected: expected{
				Expected: tests.Expected{
					Error: error_wrapper.GrpcErrorWrapper(
						codes.FailedPrecondition,
						error_wrapper.InternalError[error_wrapper.NOT_FOUND].Error(),
						error_wrapper.InternalError[error_wrapper.NOT_FOUND],
						error_wrapper.CreateGrpcErrorDetail(error_wrapper.NOT_FOUND),
					),
				},
			},
		},
	}

	for name, tc := range suite {
		t.Run(name, func(t *testing.T) {
			testTools := tests.SetupTest(t, tests.SetupOptions{Mock: tc.Mock})
			defer testTools.Teardown(t)

			testTools.C.RequireInvoke(func(db *database.DB, cs service.IClaimService) {
				input := tc.Input.(entity.SubmitClaimOrderInput)
				actual, err := cs.SubmitClaimOrder(context.Background(), input)
				expected := tc.Expected.(expected)
				if expected.Error != nil {
					require.Error(t, err)
					assert.EqualError(t, err, expected.Error.Error())
					return
				} else {
					require.NoError(t, err)
					assert.NotEmpty(t, actual)
				}

				actualOrder := claim_order_model.ClaimOrder{}
				db.CampaignDB.Debug().Where(`id=?`, input.Req.Id).Find(&actualOrder)
				actualOrder.CreatedAt = nil
				assert.Equal(t, expected.order, actualOrder)
				assert.Equal(t, actualOrder.ID, actual.Id)
				assert.Equal(t, actualOrder.Status.String(), actual.Status)

				// Expectation Claimed Items
				actualClaimItems := []claim_item_model.ClaimItem{}
				db.CampaignDB.Debug().
					Where(`claim_order_id=?`, input.Req.Id).
					Order("id").
					Find(&actualClaimItems)
				assert.Equal(t, expected.claimItems, actualClaimItems)
				assert.Equal(t, len(actualClaimItems), len(actual.Rewards))
				actualOrder.ClaimItems = actualClaimItems
				itemsMap := actualOrder.ClaimItemsToMap()
				for _, reward := range actual.Rewards {
					assert.Equal(t, itemsMap[reward.ClaimItemId].ID, reward.ClaimItemId)
					assert.Equal(t, itemsMap[reward.ClaimItemId].RewardId, reward.RewardId)
					assert.Equal(t, itemsMap[reward.ClaimItemId].Status.String(), reward.Status)
				}

				// Expectation User Point
				actualUserPoint := model.UserPoint{}
				db.CampaignDB.Debug().Where(`user_id=?`, input.UserId).Find(&actualUserPoint)
				actualUserPoint.CreatedAt = nil
				actualUserPoint.UpdatedAt = nil
				assert.Equal(t, expected.userPoint, actualUserPoint)

				// Expectation remaining rewards
				actualRewardItems := []model.RewardItem{}
				db.CampaignDB.Debug().Where(`campaign_id=?`, actualOrder.CampaignId).Find(&actualRewardItems)
				for i, item := range actualRewardItems {
					item.CreatedAt = nil
					item.UpdatedAt = nil
					assert.Equal(t, expected.rewardItems[i], item)
				}

			})
		})
	}
}
