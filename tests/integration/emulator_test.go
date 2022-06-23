// +build integration

package integration_test

/*
import (
	"context"
	"os"
	"testing"

	"github.com/Feride3d/payment-service-emulator/internal/pb"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
)

type EmulatorSuite struct {
	suite.Suite
	ctx           context.Context
	emulatorConn   *grpc.ClientConn // устанавливаем соединение для дальнейшего переиспользования
	emulatorClient pb.EmulatorClient
}

func (s *EmulatorSuite) SetupSuite() {
	emulatorHost := os.Getenv("EMULATOR_SERVER_HOST")
	if emulatorHost == "" { // если запускаем в локальной сети, а не в CI-CD
		emulatorHost = "127.0.0.1:9001"
	}

	emulatorConn, err := grpc.Dial(emulatorHost, grpc.WithInsecure())
	s.Require().NoError(err) // требуется, чтобы не было ошибки при подключении

	s.ctx = context.Background()
	s.emulatorClient = pb.NewEmulatorClient(emulatorConn)
	s.emulatorClient.CreateBanner
}

// test CreateBanner method
func (s *EmulatorSuite) TestCreateBanner() {
	req := &pb.CreateBannerRequest{
//		BannerId: "banner_id",
	}
	_, err := s.emulatorClient.CreateBanner(s.ctx, req)
	s.Require().NoError(err)

	s.Require().Equal(pb.BannerId, resp.GetBannerId() // проверить появилось ли что-то в бд
}

// test AddBanner method
func (s *EmulatorSuite) TestAddBanner() {

	bannerID := resp.GetBannerId()
	slotID := resp.GetSlotId()

	sendReq := &pb.AddBannerRequest{
		BannerId: bannerID,
		SlotId: slotID,
	}
	resp, err := s.emulatorClient.AddBanner(s.ctx, req)
	s.Require().NoError(err)

	s.Require().Equal([]string{""}, resp.GetBannerId()) // мб проверить через слот в бд
}



func TestEmulatorSuite(t *testing.T) {
	suite.Run(t, new(EmulatorSuite))
}
*/
