// Code generated by "pipeline"; DO NOT EDIT.
package yelan

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
	1: {"marked"},
	3: {"travel"},
	7: {"hold", "travel", "weakspot"},
}

func init() {
	base = &model.AvatarData{}
	err := prototext.Unmarshal(pbData, base)
	if err != nil {
		panic(err)
	}
	validation.RegisterCharParamValidationFunc(keys.Yelan, ValidateParamKeys)
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
	attack = [][][]float64{
		{attack_1},
		{attack_2},
		{attack_3},
		attack_4,
	}
)

var (
	// attack: aim = [4]
	aim = []float64{
		0.43860000371932983,
		0.47429999709129333,
		0.5099999904632568,
		0.5609999895095825,
		0.5967000126838684,
		0.637499988079071,
		0.6935999989509583,
		0.7497000098228455,
		0.8058000206947327,
		0.8669999837875366,
		0.9282000064849854,
		0.9894000291824341,
		1.0506000518798828,
		1.111799955368042,
		1.1729999780654907,
	}
	// attack: attack_1 = [0]
	attack_1 = []float64{
		0.4067800045013428,
		0.4398899972438812,
		0.4729999899864197,
		0.5202999711036682,
		0.553409993648529,
		0.5912500023841858,
		0.643280029296875,
		0.6953099966049194,
		0.7473400235176086,
		0.804099977016449,
		0.8608599901199341,
		0.9176200032234192,
		0.9743800163269043,
		1.0311399698257446,
		1.0879000425338745,
	}
	// attack: attack_2 = [1]
	attack_2 = []float64{
		0.3904399871826172,
		0.42221999168395996,
		0.45399999618530273,
		0.49939998984336853,
		0.5311800241470337,
		0.5674999952316284,
		0.6174399852752686,
		0.6673799753189087,
		0.7173200249671936,
		0.7717999815940857,
		0.8262799978256226,
		0.8807600140571594,
		0.9352399706840515,
		0.9897199869155884,
		1.0441999435424805,
	}
	// attack: attack_3 = [2]
	attack_3 = []float64{
		0.515999972820282,
		0.5580000281333923,
		0.6000000238418579,
		0.6600000262260437,
		0.7020000219345093,
		0.75,
		0.8159999847412109,
		0.8820000290870667,
		0.9480000138282776,
		1.0199999809265137,
		1.0920000076293945,
		1.1640000343322754,
		1.2359999418258667,
		1.3079999685287476,
		1.3799999952316284,
	}
	// attack: attack_4 = [3 3]
	attack_4 = [][]float64{
		{
			0.325080007314682,
			0.3515399992465973,
			0.3779999911785126,
			0.415800005197525,
			0.4422599971294403,
			0.4724999964237213,
			0.5140799880027771,
			0.5556600093841553,
			0.5972399711608887,
			0.6425999999046326,
			0.6879600286483765,
			0.7333199977874756,
			0.7786800265312195,
			0.8240399956703186,
			0.8694000244140625,
		},
		{
			0.325080007314682,
			0.3515399992465973,
			0.3779999911785126,
			0.415800005197525,
			0.4422599971294403,
			0.4724999964237213,
			0.5140799880027771,
			0.5556600093841553,
			0.5972399711608887,
			0.6425999999046326,
			0.6879600286483765,
			0.7333199977874756,
			0.7786800265312195,
			0.8240399956703186,
			0.8694000244140625,
		},
	}
	// attack: barb = [6]
	barb = []float64{
		0.11575999855995178,
		0.12444200366735458,
		0.1331239938735962,
		0.14470000565052032,
		0.15338200330734253,
		0.16206400096416473,
		0.17363999783992767,
		0.1852159947156906,
		0.19679200649261475,
		0.20836800336837769,
		0.21994400024414062,
		0.23151999711990356,
		0.24598999321460724,
		0.2604599893093109,
		0.2749300003051758,
	}
	// attack: fullaim = [5]
	fullaim = []float64{
		1.2400000095367432,
		1.3329999446868896,
		1.4259999990463257,
		1.5499999523162842,
		1.6430000066757202,
		1.7359999418258667,
		1.8600000143051147,
		1.9839999675750732,
		2.1080000400543213,
		2.2320001125335693,
		2.3559999465942383,
		2.4800000190734863,
		2.634999990463257,
		2.7899999618530273,
		2.944999933242798,
	}
	// skill: skill = [0]
	skill = []float64{
		0.2261359989643097,
		0.2430959939956665,
		0.2600559890270233,
		0.2826699912548065,
		0.29962998628616333,
		0.31659001111984253,
		0.33920401334762573,
		0.36181798577308655,
		0.38443100452423096,
		0.40704500675201416,
		0.4296579957008362,
		0.4522719979286194,
		0.4805389940738678,
		0.5088059902191162,
		0.537073016166687,
	}
	// burst: burst = [0]
	burst = []float64{
		0.07308000326156616,
		0.07856100052595139,
		0.08404199779033661,
		0.0913499966263771,
		0.09683100134134293,
		0.10231199860572815,
		0.10961999744176865,
		0.11692799627780914,
		0.12423600256443024,
		0.13154399394989014,
		0.13885200023651123,
		0.14616000652313232,
		0.1552949994802475,
		0.16443000733852386,
		0.17356500029563904,
	}
	// burst: burstDice = [1]
	burstDice = []float64{
		0.04871999844908714,
		0.05237400159239769,
		0.05602800101041794,
		0.0608999989926815,
		0.06455399841070175,
		0.0682080015540123,
		0.07308000326156616,
		0.07795199751853943,
		0.08282399922609329,
		0.08769600093364716,
		0.09256800264120102,
		0.09743999689817429,
		0.10352999716997147,
		0.10961999744176865,
		0.11570999771356583,
	}
)
