package impl

import (
	"github.com/iafoosball/livematches-service/models"
	"log"
	"testing"
)

func TestRotatePeople(t *testing.T) {
	positions := &models.MatchPositions{
		BlueAttack: "BlueAtt",
		RedAttack:  "RedAtt",
		RedDefense: "RedDef",
	}
	rotatePeople(positions)
	if positions.BlueAttack != "RedAtt" || positions.RedAttack != "RedDef" || positions.RedDefense != "BlueAtt" {
		log.Println(positions)
		t.Fail()
	}

}
