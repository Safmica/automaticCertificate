package main

import (
	"automaticCertificate/entity"
	"fmt"
)

func main() {
	certificate := "certificate.png"
	participants, err := entity.ReadSpreadsheet("participants.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, participant := range participants {
		name := participant.Name
		//email := participant.Email

		entity.GenerateCertificate(name, certificate)

		//entity.SendEmail(email, "Your Certificate", "Please find your certificate attached.", certificatePath)
	}
}
