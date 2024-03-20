// Code generated by "pipeline"; DO NOT EDIT.
package razor

import (
	_ "embed"

	"fmt"
	"github.com/genshinsim/gcsim/pkg/core/action"
	"github.com/genshinsim/gcsim/pkg/core/keys"
	"github.com/genshinsim/gcsim/pkg/gcs/validation"
	"github.com/genshinsim/gcsim/pkg/model"
	"google.golang.org/protobuf/encoding/prototext"
	"slices"
)

//go:embed data_gen.textproto
var pbData []byte
var base *model.AvatarData
var paramKeysValidation = map[action.Action][]string{
	1: {"hold"},
}

func init() {
	base = &model.AvatarData{}
	err := prototext.Unmarshal(pbData, base)
	if err != nil {
		panic(err)
	}
	validation.RegisterCharParamValidationFunc(keys.Razor, ValidateParamKeys)
}

func ValidateParamKeys(a action.Action, keys []string) error {
	valid, ok := paramKeysValidation[a]
	if !ok {
		return nil
	}
	for _, v := range keys {
		if !slices.Contains(valid, v) {
			return fmt.Errorf("key %v is invalid for action %v", v, a.String())
		}
	}
	return nil
}

func (x *char) Data() *model.AvatarData {
	return base
}

var (
	attack = [][]float64{
		attack_1,
		attack_2,
		attack_3,
		attack_4,
	}
)

var (
	// attack: attack_1 = [0]
	attack_1 = []float64{
		0.9592000246047974,
		1.0246000289916992,
		1.090000033378601,
		1.1771999597549438,
		1.2425999641418457,
		1.3188999891281128,
		1.4170000553131104,
		1.5151000022888184,
		1.6131999492645264,
		1.711300015449524,
		1.809399962425232,
		1.9075000286102295,
		2.0055999755859375,
		2.1036999225616455,
		2.2018001079559326,
	}
	// attack: attack_2 = [1]
	attack_2 = []float64{
		0.8263199925422668,
		0.8826599717140198,
		0.9390000104904175,
		1.0141199827194214,
		1.0704599618911743,
		1.1361900568008423,
		1.2207000255584717,
		1.305209994316101,
		1.3897199630737305,
		1.4742300510406494,
		1.5587400197982788,
		1.6432499885559082,
		1.7277599573135376,
		1.8122700452804565,
		1.896780014038086,
	}
	// attack: attack_3 = [2]
	attack_3 = []float64{
		1.033120036125183,
		1.103559970855713,
		1.1740000247955322,
		1.2679200172424316,
		1.3383599519729614,
		1.4205399751663208,
		1.5262000560760498,
		1.6318600177764893,
		1.7375199794769287,
		1.8431799411773682,
		1.9488400220870972,
		2.054500102996826,
		2.1601600646972656,
		2.265820026397705,
		2.3714799880981445,
	}
	// attack: attack_4 = [3]
	attack_4 = []float64{
		1.3604799509048462,
		1.4532400369644165,
		1.5460000038146973,
		1.6696799993515015,
		1.7624399662017822,
		1.8706599473953247,
		2.0097999572753906,
		2.148940086364746,
		2.2880799770355225,
		2.427220106124878,
		2.5663599967956543,
		2.7054998874664307,
		2.844640016555786,
		2.9837799072265625,
		3.122920036315918,
	}
	// skill: skillHold = [1]
	skillHold = []float64{
		2.9519999027252197,
		3.1733999252319336,
		3.3947999477386475,
		3.690000057220459,
		3.911400079727173,
		4.132800102233887,
		4.427999973297119,
		4.723199844360352,
		5.018400192260742,
		5.313600063323975,
		5.608799934387207,
		5.9039998054504395,
		6.2729997634887695,
		6.642000198364258,
		7.011000156402588,
	}
	// skill: skillPress = [0]
	skillPress = []float64{
		1.9919999837875366,
		2.141400098800659,
		2.290800094604492,
		2.490000009536743,
		2.639400005340576,
		2.788800001144409,
		2.98799991607666,
		3.1872000694274902,
		3.386399984359741,
		3.585599899291992,
		3.7848000526428223,
		3.9839999675750732,
		4.232999801635742,
		4.48199987411499,
		4.730999946594238,
	}
	// burst: burstATKSpeed = [2]
	burstATKSpeed = []float64{
		0.25999999046325684,
		0.2800000011920929,
		0.30000001192092896,
		0.3199999928474426,
		0.3400000035762787,
		0.36000001430511475,
		0.3700000047683716,
		0.3799999952316284,
		0.38999998569488525,
		0.4000000059604645,
		0.4000000059604645,
		0.4000000059604645,
		0.4000000059604645,
		0.4000000059604645,
		0.4000000059604645,
	}
	// burst: burstDmg = [0]
	burstDmg = []float64{
		1.600000023841858,
		1.7200000286102295,
		1.840000033378601,
		2,
		2.119999885559082,
		2.240000009536743,
		2.4000000953674316,
		2.559999942779541,
		2.7200000286102295,
		2.880000114440918,
		3.0399999618530273,
		3.200000047683716,
		3.4000000953674316,
		3.5999999046325684,
		3.799999952316284,
	}
	// burst: wolfDmg = [1]
	wolfDmg = []float64{
		0.23999999463558197,
		0.257999986410141,
		0.2759999930858612,
		0.30000001192092896,
		0.3179999887943268,
		0.335999995470047,
		0.36000001430511475,
		0.3840000033378601,
		0.40799999237060547,
		0.4320000112056732,
		0.4560000002384186,
		0.47999998927116394,
		0.5099999904632568,
		0.5400000214576721,
		0.5699999928474426,
	}
)
