package cli

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"jwtcli/pkg/handlers"
	"os"
	"time"

	"github.com/spf13/cobra"
)

type CreateConfig struct {
	iss     string
	sub     string
	aud     []string
	exp     string
	method  string
	keyfile string
}

var createConfig = CreateConfig{}

func createCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create",
		Run: func(cmd *cobra.Command, args []string) {
			sm, err := handlers.GetSigningMethod(createConfig.method)
			if err != nil {
				fmt.Printf("failed: %s", err.Error())
				os.Exit(1)
			}

			duration, err := time.ParseDuration(createConfig.exp)
			if err != nil {
				fmt.Printf("failed parsing exp: %s", err.Error())
				os.Exit(1)
			}

			token := jwt.New(sm)
			token.Claims = &jwt.RegisteredClaims{
				Issuer:    createConfig.iss,
				Subject:   createConfig.sub,
				Audience:  createConfig.aud,
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
				NotBefore: jwt.NewNumericDate(time.Now()),
			}

			var rawKey []byte
			if createConfig.keyfile != "" {
				rawKey, err = os.ReadFile(createConfig.keyfile)
				if err != nil {
					fmt.Printf("failed reading keyfile: %s", err.Error())
					os.Exit(1)
				}
			} else {
				rawKey = []byte(os.Getenv("JWT_KEY"))
				if len(rawKey) == 0 {
					fmt.Printf("failed loading key from JWT_KEY env")
					os.Exit(1)
				}
			}
			key, err := handlers.ParseKey(rawKey, createConfig.method)
			if err != nil {
				fmt.Printf("failed parsing key: %s", err.Error())
				os.Exit(1)
			}

			signed, err := token.SignedString(key)
			if err != nil {
				fmt.Printf("failed signing token: %s", err.Error())
				os.Exit(1)
			}

			fmt.Printf("%s\n", signed)
		},
	}
	cmd.PersistentFlags().StringVarP(&createConfig.iss, "iss", "i", "issuer_uri", "jwt issuer claim value")
	cmd.PersistentFlags().StringVarP(&createConfig.sub, "sub", "s", "subject", "jwt subject claim value")
	cmd.PersistentFlags().StringSliceVarP(&createConfig.aud, "aud", "a", []string{"audience"}, "jwt audience claim value")
	cmd.PersistentFlags().StringVarP(&createConfig.exp, "exp", "e", "60m", "jwt expires claim value, specify like 60s, 2h etc.")
	//cmd.PersistentFlags().StringVarP(&createConfig.custom, "custom", "c", "", "jwt custom claim value, specify as claim=value, can be used multiple times")
	cmd.PersistentFlags().StringVarP(&createConfig.keyfile, "keyfile", "k", "", "private key file for signing")
	cmd.PersistentFlags().StringVarP(&createConfig.method, "method", "m", "", "token signing method (rs256, rs512, hs256 etc.)")
	// cmd.PersistentFlags().StringVarP(createConfig.exp, "output", "o", "", "output format")
	return cmd
}
