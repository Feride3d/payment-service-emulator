package cmd

import (
	"fmt"
	"net"
	"net/http"

	"github.com/caarlos0/env/v6"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/Feride3d/payment-service-emulator/internal/storage"
	"github.com/kulti/tFeride3d/payment-service-emulator/internal/router"
)

type serverCmdFlags struct {
	Port uint16 `env:"PORT" envDefault:"0"`
}

const (
	emulatorCredentialPath = "/etc/tl/emulator/credentials.json"
	emulatorIDsPath        = "/etc/tl/emulator/emulators.json"
)

func newServerCmd(dbFlags dbFlags) *cobra.Command {
	var serverCmdFlags serverCmdFlags
	if err := env.Parse(&serverCmdFlags); err != nil {
		zap.S().Fatalw("failed to parse server cmd flags", zap.Error(err))
	}

	serverCmd := &cobra.Command{
		Use:   "server",
		Short: "Starts payment service emulator server",
		Run: func(cmd *cobra.Command, args []string) {
			listener, err := net.Listen("tcp", fmt.Sprintf(":%d", serverCmdFlags.Port))
			if err != nil {
				zap.S().Fatalw("failed to listen", zap.Error(err), "port", serverCmdFlags.Port)
			}
			zap.S().Infow("listen at", "addr", listener.Addr().String())

			dbStore, err := storage.New(dbFlags.URL())
			if err != nil {
				zap.S().Fatalw("failed to connect to db", zap.Error(err))
			}

			paymentStore := paymentstore.New(dbStore)
			paymentTmpl := paymenttmpl.New(dbStore, newPaymentService())
			router := router.New(paymentStore, paymentTmpl)

			err = http.Serve(listener, router.RootHandler())
			if err != nil {
				zap.S().Fatalw("failed to graceful server shutdown", zap.Error(err))
			}
		},
	}

	return serverCmd
}

func newPaymentService() paymenttmpl.PayService {
	ps, err := payservice.New(payservice.Options{
		CredentialPath:  emulatorCredentialPath,
		EmulatorIDsPath: emulatorIDsPath,
	})
	if err != nil {
		zap.L().Warn("failed to create payment service", zap.Error(err))
		return nil
	}

	zap.L().Info("payment service created")
	return ps
}
