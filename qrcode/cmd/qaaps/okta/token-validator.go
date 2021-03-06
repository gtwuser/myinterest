package main

import (
	"fmt"
	jwtverifier "github.com/okta/okta-jwt-verifier-golang"
	log "github.com/sirupsen/logrus"
)

func main() {
	toValidate := map[string]string{}
	toValidate["aud"] = "http://localhost:8000"
	//toValidate["cid"] = "0oarhbhjvB2iZjpOX4x6"

	jwtVerifierSetup := jwtverifier.JwtVerifier{
		Issuer:           "https://dev-285665.okta.com/oauth2/ausks0dmxpZ0cXyqm4x6",
		ClaimsToValidate: toValidate,
	}

	verifier := jwtVerifierSetup.New()

	token, err := verifier.VerifyAccessToken("eyJraWQiOiJFbTR6ZThPZWdIY3M3ZEN6NzM4anVhVFY3ZS12UmNTdjNrSUpZeDRIdnJFIiwiYWxnIjoiUlMyNTYifQ.eyJ2ZXIiOjEsImp0aSI6IkFULmVoQWpqNHh6TDltXy1rWmpFdC1NWGZsYnloSGp6N2tmNDJ2UHlkMzllX0kiLCJpc3MiOiJodHRwczovL2Rldi0yODU2NjUub2t0YS5jb20vb2F1dGgyL2F1c2tzMGRteHBaMGNYeXFtNHg2IiwiYXVkIjoiaHR0cDovL2xvY2FsaG9zdDo4MDAwIiwiaWF0IjoxNTk4MjA2MTk3LCJleHAiOjE1OTgyMDk3OTcsImNpZCI6IjBvYXJoYmhqdkIyaVpqcE9YNHg2Iiwic2NwIjpbInJlYWQiXSwic3ViIjoiMG9hcmhiaGp2QjJpWmpwT1g0eDYifQ.YUV6W7cKIle2UeKCt13MZ-Sk036lBMUkDLn4Tehm99cvd6ViUTG7CkWwGma28rQcAXGGTxPZLYUdOld1xTYVoq1vHEAGFBPlE8Kouwv_qv_rDG31UWkTHkkzME2vXoPckmq_BlDNTfF9Kl57dhL0c6JdKUx1norpdxo3TfDwXxIOKYPO6DzmFObfISqmc6pEHjFvB0KLJnRh9WIAcEneoPAhr9ywezmGIe9tJtlxfZMYBfX0Ii4KmZQRfZL8t-1vzgg7JgGFDEPaTgcHLoIz591E4bDPuvXpMVUF_X2t9vnBklNyFVjjp5AveiT6YdtRL4FQQzaLZZAEtf0jvgzrGw")
	//token, err := verifier.VerifyAccessToken("eyJraWQiOiJFbTR6ZThPZWdIY3M3ZEN6NzM4anVhVFY3ZS12UmNTdjNrSUpZeDRIdnJFIiwiYWxnIjoiUlMyNTYifQ.eyJ2ZXIiOjEsImp0aSI6IkFULmVoQWpqNHh6TDltXy1rWmpFdC1NWGZsYnloSGp6N2tmNDJ2UHlkMzllX0kiLCJpc3MiOiJodHRwczovL2Rldi0yODU2NjUub2t0YS5jb20vb2F1dGgyL2F1c2tzMGRteHBaMGNYeXFtNHg2IiwiYXVkIjoiaHR0cDovL2xvY2FsaG9zdDo4MDgwIiwiaWF0IjoxNTk4MjA2MTk3LCJleHAiOjE1OTgyMDk3OTcsImNpZCI6IjBvYXJoYmhqdkIyaVpqcE9YNHg2Iiwic2NwIjpbInJlYWQiXSwic3ViIjoiMG9hcmhiaGp2QjJpWmpwT1g0eDYifQ.YUV6W7cKIle2UeKCt13MZ-Sk036lBMUkDLn4Tehm99cvd6ViUTG7CkWwGma28rQcAXGGTxPZLYUdOld1xTYVoq1vHEAGFBPlE8Kouwv_qv_rDG31UWkTHkkzME2vXoPckmq_BlDNTfF9Kl57dhL0c6JdKUx1norpdxo3TfDwXxIOKYPO6DzmFObfISqmc6pEHjFvB0KLJnRh9WIAcEneoPAhr9ywezmGIe9tJtlxfZMYBfX0Ii4KmZQRfZL8t-1vzgg7JgGFDEPaTgcHLoIz591E4bDPuvXpMVUF_X2t9vnBklNyFVjjp5AveiT6YdtRL4FQQzaLZZAEtf0jvgzrGw")
	if err != nil {
		log.Fatal("Error 😡😡", err)
	}
	for k, v := range token.Claims {
		fmt.Println(k, v)
	}
}
