package impl

import (
	"github.com/iafoosball/livematches-service/models"
	"testing"
)

func TestRotatePeople(t *testing.T) {
	positions := &models.MatchPositions{
		BlueAttack: "BlueDef",
		RedAttack:  "RedDef",
		RedDefense: "RedAtt",
	}
	rotatePeople(positions)

}
