package nft_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/quartz-technology/quicknode-sdks/go-sdk/client"
	"github.com/stretchr/testify/assert"
)

func TestCollectionDetails(t *testing.T) {
	c := client.New()

	type TestCase map[string]struct {
		address string
		expect  string
	}

	testCases := TestCase{
		"executes query": {
			address: "0x60e4d786628fea6478f785a6d7e704777c86a7c6",
			expect: `{
    "Contract": {
        "Address": "0x60e4d786628fea6478f785a6d7e704777c86a7c6",
        "IsVerified": true,
        "TokenStandard": "ERC721",
        "Attributes": [
            {
                "Name": "Earring",
                "Rarity": 0.57800485111214328082240854200790636241436004638671875,
                "Value": "M2 Diamond Stud",
                "ValueCount": 112
            },
            {
                "Name": "Fur",
                "Rarity": 1.7701398565309387489463688325486145913600921630859375,
                "Value": "M2 Golden Brown",
                "ValueCount": 343
            },
            {
                "Name": "Hat",
                "Rarity": 0.7741136398823347253284055113908834755420684814453125,
                "Value": "M2 Commie Hat",
                "ValueCount": 150
            },
            {
                "Name": "Mouth",
                "Rarity": 0.691541518294885637629931807168759405612945556640625,
                "Value": "M2 Rage",
                "ValueCount": 134
            },
            {
                "Name": "Name",
                "Rarity": 0.0051607575992155645117076545602685655467212200164794921875,
                "Value": "Mega DMT",
                "ValueCount": 1
            },
            {
                "Name": "Clothes",
                "Rarity": 0.7741136398823347253284055113908834755420684814453125,
                "Value": "M2 Hawaiian",
                "ValueCount": 150
            },
            {
                "Name": "Eyes",
                "Rarity": 1.821747432523094456513490513316355645656585693359375,
                "Value": "M2 Sleepy",
                "ValueCount": 353
            },
            {
                "Name": "Name",
                "Rarity": 0.0051607575992155645117076545602685655467212200164794921875,
                "Value": "Mega Death Bot",
                "ValueCount": 1
            },
            {
                "Name": "Name",
                "Rarity": 0.0051607575992155645117076545602685655467212200164794921875,
                "Value": "Mega Noise",
                "ValueCount": 1
            },
            {
                "Name": "Hat",
                "Rarity": 0.273520152758424950345528259276761673390865325927734375,
                "Value": "M2 Trippy Captain's Hat",
                "ValueCount": 53
            },
            {
                "Name": "Mouth",
                "Rarity": 5.05754244723125356841819666442461311817169189453125,
                "Value": "M2 Bored",
                "ValueCount": 980
            },
            {
                "Name": "Mouth",
                "Rarity": 2.11074985807916615243584601557813584804534912109375,
                "Value": "M1 Jovial",
                "ValueCount": 409
            },
            {
                "Name": "Eyes",
                "Rarity": 1.6772462197450586085523127621854655444622039794921875,
                "Value": "M1 Scumbag",
                "ValueCount": 325
            },
            {
                "Name": "Fur",
                "Rarity": 2.131392888476028257827010747860185801982879638671875,
                "Value": "M1 Robot",
                "ValueCount": 413
            },
            {
                "Name": "Mouth",
                "Rarity": 0.9186148526603705732895832625217735767364501953125,
                "Value": "M1 Bored Cigar",
                "ValueCount": 178
            },
            {
                "Name": "Name",
                "Rarity": 0.0051607575992155645117076545602685655467212200164794921875,
                "Value": "Mega Swamp",
                "ValueCount": 1
            },
            {
                "Name": "Name",
                "Rarity": 0.0051607575992155645117076545602685655467212200164794921875,
                "Value": "Mega Electric",
                "ValueCount": 1
            },
            {
                "Name": "Name",
                "Rarity": 0.0051607575992155645117076545602685655467212200164794921875,
                "Value": "Mega Radioactive",
                "ValueCount": 1
            },
            {
                "Name": "Name",
                "Rarity": 0.0051607575992155645117076545602685655467212200164794921875,
                "Value": "Mega Demon",
                "ValueCount": 1
            },
            {
                "Name": "Hat",
                "Rarity": 1.3727615213913402225642812481964938342571258544921875,
                "Value": "M1 Spinner Hat",
                "ValueCount": 266
            },
            {
                "Name": "Clothes",
                "Rarity": 0.443825153532538596579115619533695280551910400390625,
                "Value": "M2 Leather Punk Jacket",
                "ValueCount": 86
            },
            {
                "Name": "Clothes",
                "Rarity": 1.1663312174227176143403994501568377017974853515625,
                "Value": "M1 Leather Punk Jacket",
                "ValueCount": 226
            },
            {
                "Name": "Hat",
                "Rarity": 0.216751819167053716430615395438508130609989166259765625,
                "Value": "M2 Bandana Blue",
                "ValueCount": 42
            },
            {
                "Name": "Clothes",
                "Rarity": 0.73282757908861018147916865927982144057750701904296875,
                "Value": "M1 Blue Dress",
                "ValueCount": 142
            },
            {
                "Name": "Hat",
                "Rarity": 2.626825618000722339928643123130314052104949951171875,
                "Value": "M1 Fisherman's Hat",
                "ValueCount": 509
            },
            {
                "Name": "Name",
                "Rarity": 0.0051607575992155645117076545602685655467212200164794921875,
                "Value": "Mega Trippy",
                "ValueCount": 1
            },
            {
                "Name": "Mouth",
                "Rarity": 0.392217577540382944523145170023781247437000274658203125,
                "Value": "M1 Bored Dagger",
                "ValueCount": 76
            },
            {
                "Name": "Clothes",
                "Rarity": 1.74433606853486100618511045468039810657501220703125,
                "Value": "M1 Tanktop",
                "ValueCount": 338
            },
            {
                "Name": "Clothes",
                "Rarity": 0.2270733343654848523929246084662736393511295318603515625,
                "Value": "M2 Black Suit",
                "ValueCount": 44
            },
            {
                "Name": "Mouth",
                "Rarity": 0.16514424317489806437464494592859409749507904052734375,
                "Value": "M2 Bored Dagger",
                "ValueCount": 32
            },
            {
                "Name": "Fur",
                "Rarity": 0.7741136398823347253284055113908834755420684814453125,
                "Value": "M2 Robot",
                "ValueCount": 150
            },
            {
                "Name": "Mouth",
                "Rarity": 0.73282757908861018147916865927982144057750701904296875,
                "Value": "M1 Bored Unshaven Pipe",
                "ValueCount": 142
            },
            {
                "Name": "Hat",
                "Rarity": 2.3842700108375911582925255061127245426177978515625,
                "Value": "M1 Halo",
                "ValueCount": 462
            },
            {
                "Name": "Name",
                "Rarity": 0.0051607575992155645117076545602685655467212200164794921875,
                "Value": "Mega Robot",
                "ValueCount": 1
            },
            {
                "Name": "Name",
                "Rarity": 0.0051607575992155645117076545602685655467212200164794921875,
                "Value": "Mega Zombie",
                "ValueCount": 1
            },
            {
                "Name": "Earring",
                "Rarity": 6.0793724518759351127528134384192526340484619140625,
                "Value": "M1 Silver Stud",
                "ValueCount": 1178
            },
            {
                "Name": "Mouth",
                "Rarity": 16.88599886463332921948676812462508678436279296875,
                "Value": "M1 Bored",
                "ValueCount": 3272
            },
            {
                "Name": "Background",
                "Rarity": 3.080972286731692211247946033836342394351959228515625,
                "Value": "M2 Blue",
                "ValueCount": 597
            },
            {
                "Name": "Mouth",
                "Rarity": 0.72250606389017912878358629313879646360874176025390625,
                "Value": "M2 Phoneme Vuh",
                "ValueCount": 140
            },
            {
                "Name": "Earring",
                "Rarity": 6.39933942302730063289573081419803202152252197265625,
                "Value": "M1 Silver Hoop",
                "ValueCount": 1240
            },
            {
                "Name": "Mouth",
                "Rarity": 1.9404448573050523396688049615477211773395538330078125,
                "Value": "M1 Small Grin",
                "ValueCount": 376
            },
            {
                "Name": "Eyes",
                "Rarity": 0.423182123135676324654497193478164263069629669189453125,
                "Value": "M1 Blue Beams",
                "ValueCount": 82
            },
            {
                "Name": "Hat",
                "Rarity": 0.381896062341951780805260341367102228105068206787109375,
                "Value": "M2 Safari",
                "ValueCount": 74
            },
            {
                "Name": "Clothes",
                "Rarity": 0.38705681994116736266420275569544173777103424072265625,
                "Value": "M2 Bandolier",
                "ValueCount": 75
            },
            {
                "Name": "Hat",
                "Rarity": 0.44898591113175412292690680260420776903629302978515625,
                "Value": "M2 Irish Boho",
                "ValueCount": 87
            },
            {
                "Name": "Hat",
                "Rarity": 0.185787273571760336299263371984125114977359771728515625,
                "Value": "M2 Laurel Wreath",
                "ValueCount": 36
            },
            {
                "Name": "Clothes",
                "Rarity": 0.32512772875058060240149870878667570650577545166015625,
                "Value": "M2 Pimp Coat",
                "ValueCount": 63
            },
            {
                "Name": "Eyes",
                "Rarity": 0.335449243949011710608232306185527704656124114990234375,
                "Value": "M2 Cyborg",
                "ValueCount": 65
            },
            {
                "Name": "Mouth",
                "Rarity": 0.54704030551684990069105651855352334678173065185546875,
                "Value": "M2 Phoneme  Ooo",
                "ValueCount": 106
            },
            {
                "Name": "Hat",
                "Rarity": 0.32512772875058060240149870878667570650577545166015625,
                "Value": "M2 Prussian Helmet",
                "ValueCount": 63
            },
            {
                "Name": "Hat",
                "Rarity": 0.53155803271920321062538050682633183896541595458984375,
                "Value": "M1 Laurel Wreath",
                "ValueCount": 103
            },
            {
                "Name": "Mouth",
                "Rarity": 0.660576972699592257498579783714376389980316162109375,
                "Value": "M2 Small Grin",
                "ValueCount": 128
            },
            {
                "Name": "Eyes",
                "Rarity": 0.37157454714352067259852674396825022995471954345703125,
                "Value": "M2 Laser Eyes",
                "ValueCount": 72
            },
            {
                "Name": "Hat",
                "Rarity": 0.32512772875058060240149870878667570650577545166015625,
                "Value": "M2 Faux Hawk",
                "ValueCount": 63
            },
            {
                "Name": "Mouth",
                "Rarity": 0.319966971151365020542556294458336196839809417724609375,
                "Value": "M2 Phoneme Wah",
                "ValueCount": 62
            },
            {
                "Name": "Hat",
                "Rarity": 0.474789699127831976710467642988078296184539794921875,
                "Value": "M2 Bayc Hat Red",
                "ValueCount": 92
            },
            {
                "Name": "Mouth",
                "Rarity": 0.34061000154822729246717472051386721432209014892578125,
                "Value": "M2 Bored Cigar",
                "ValueCount": 66
            },
            {
                "Name": "Mouth",
                "Rarity": 0.392217577540382944523145170023781247437000274658203125,
                "Value": "M2 Bored Bubblegum",
                "ValueCount": 76
            },
            {
                "Name": "Mouth",
                "Rarity": 0.57800485111214328082240854200790636241436004638671875,
                "Value": "M2 Grin Multicolored",
                "ValueCount": 112
            },
            {
                "Name": "Mouth",
                "Rarity": 0.3148062135521494386836138801299966871738433837890625,
                "Value": "M2 Bored Unshaven Pipe",
                "ValueCount": 61
            },
            {
                "Name": "Hat",
                "Rarity": 0.304484698353718330476880282731144689023494720458984375,
                "Value": "M2 Party Hat 1",
                "ValueCount": 59
            },
            {
                "Name": "Mouth",
                "Rarity": 0.2838416679568560585522618566756136715412139892578125,
                "Value": "M2 Bored Party Horn",
                "ValueCount": 55
            },
            {
                "Name": "Clothes",
                "Rarity": 0.216751819167053716430615395438508130609989166259765625,
                "Value": "M2 Admirals Coat",
                "ValueCount": 42
            },
            {
                "Name": "Hat",
                "Rarity": 1.68240697734427424592240640777163207530975341796875,
                "Value": "M1 Irish Boho",
                "ValueCount": 326
            },
            {
                "Name": "Clothes",
                "Rarity": 0.319966971151365020542556294458336196839809417724609375,
                "Value": "M2 Lab Coat",
                "ValueCount": 62
            },
            {
                "Name": "Clothes",
                "Rarity": 0.34061000154822729246717472051386721432209014892578125,
                "Value": "M2 Tweed Suit",
                "ValueCount": 66
            },
            {
                "Name": "Clothes",
                "Rarity": 0.86184651906899933937467039868352003395557403564453125,
                "Value": "M2 Black T",
                "ValueCount": 167
            },
            {
                "Name": "Clothes",
                "Rarity": 0.54704030551684990069105651855352334678173065185546875,
                "Value": "M2 Sleeveless T",
                "ValueCount": 106
            },
            {
                "Name": "Earring",
                "Rarity": 3.22031274191051242183903013938106596469879150390625,
                "Value": "M1 Gold Stud",
                "ValueCount": 624
            },
            {
                "Name": "Clothes",
                "Rarity": 0.52123651752077204690749567816965281963348388671875,
                "Value": "M2 Work Vest",
                "ValueCount": 101
            },
            {
                "Name": "Hat",
                "Rarity": 0.54187954791763426332096287296735681593418121337890625,
                "Value": "M2 Vietnam Era Helmet",
                "ValueCount": 105
            },
            {
                "Name": "Hat",
                "Rarity": 1.1714919750219332517104930957430042326450347900390625,
                "Value": "M1 Stuntman Helmet",
                "ValueCount": 227
            },
            {
                "Name": "Mouth",
                "Rarity": 0.7121845486917479650657014644821174442768096923828125,
                "Value": "M2 Jovial",
                "ValueCount": 138
            },
            {
                "Name": "Clothes",
                "Rarity": 0.47995045672704750305825882605859078466892242431640625,
                "Value": "M2 Lumberjack Shirt",
                "ValueCount": 93
            },
            {
                "Name": "Clothes",
                "Rarity": 0.454146668730969704785849216932547278702259063720703125,
                "Value": "M2 Black Holes T",
                "ValueCount": 88
            },
            {
                "Name": "Eyes",
                "Rarity": 2.53909273881405805894928562338463962078094482421875,
                "Value": "M1 Sunglasses",
                "ValueCount": 492
            },
            {
                "Name": "Fur",
                "Rarity": 4.73757547607988893645369898877106606960296630859375,
                "Value": "M1 Cream",
                "ValueCount": 918
            },
            {
                "Name": "Hat",
                "Rarity": 1.8423904629199565619046552455984055995941162109375,
                "Value": "M1 Bayc Hat Black",
                "ValueCount": 357
            },
            {
                "Name": "Clothes",
                "Rarity": 0.6347731847035145147373214058461599051952362060546875,
                "Value": "M1 Pimp Coat",
                "ValueCount": 123
            },
            {
                "Name": "Clothes",
                "Rarity": 0.47995045672704750305825882605859078466892242431640625,
                "Value": "M2 Bone Necklace",
                "ValueCount": 93
            },
            {
                "Name": "Eyes",
                "Rarity": 2.074624554884656912889795421506278216838836669921875,
                "Value": "M2 Bloodshot",
                "ValueCount": 402
            },
            {
                "Name": "Clothes",
                "Rarity": 1.285028642204675719540318823419511318206787109375,
                "Value": "M1 Bandolier",
                "ValueCount": 249
            },
            {
                "Name": "Clothes",
                "Rarity": 0.2528771223615626784209098332212306559085845947265625,
                "Value": "M2 Prom Dress",
                "ValueCount": 49
            },
            {
                "Name": "Hat",
                "Rarity": 0.45930742633018528664479163126088678836822509765625,
                "Value": "M2 Sushi Chef Headband",
                "ValueCount": 89
            },
            {
                "Name": "Mouth",
                "Rarity": 0.289002425556071640411204271003953181207180023193359375,
                "Value": "M2 Bored Unshaven Cigar",
                "ValueCount": 56
            },
            {
                "Name": "Name",
                "Rarity": 0.0051607575992155645117076545602685655467212200164794921875,
                "Value": "Mega Jelly",
                "ValueCount": 1
            },
            {
                "Name": "Hat",
                "Rarity": 1.945605614904267977038898607133887708187103271484375,
                "Value": "M1 Bowler",
                "ValueCount": 377
            },
            {
                "Name": "Background",
                "Rarity": 2.79713061877483593065107925212942063808441162109375,
                "Value": "M2 Army Green",
                "ValueCount": 542
            },
            {
                "Name": "Clothes",
                "Rarity": 0.6967022758941012750000254527549259364604949951171875,
                "Value": "M2 Sailor Shirt",
                "ValueCount": 135
            },
            {
                "Name": "Eyes",
                "Rarity": 1.1663312174227176143403994501568377017974853515625,
                "Value": "M2 Coins",
                "ValueCount": 226
            },
            {
                "Name": "Mouth",
                "Rarity": 3.64865562264540432835246974718756973743438720703125,
                "Value": "M2 Bored Unshaven",
                "ValueCount": 707
            },
            {
                "Name": "Clothes",
                "Rarity": 0.41802136553646074279555477914982475340366363525390625,
                "Value": "M2 Smoking Jacket",
                "ValueCount": 81
            },
            {
                "Name": "Fur",
                "Rarity": 1.0785983382360531113164370253798551857471466064453125,
                "Value": "M2 White",
                "ValueCount": 209
            },
            {
                "Name": "Clothes",
                "Rarity": 0.97538318625174180720449612636002711951732635498046875,
                "Value": "M1 Hip Hop",
                "ValueCount": 189
            },
            {
                "Name": "Clothes",
                "Rarity": 0.6812200030964545849343494410277344286441802978515625,
                "Value": "M1 Prom Dress",
                "ValueCount": 132
            },
            {
                "Name": "Eyes",
                "Rarity": 1.548227279764669450656811022781766951084136962890625,
                "Value": "M1 Hypnotized",
                "ValueCount": 300
            },
            {
                "Name": "Hat",
                "Rarity": 0.232234091964700406496291407165699638426303863525390625,
                "Value": "M2 Ww2 Pilot Helm",
                "ValueCount": 45
            },
            {
                "Name": "Clothes",
                "Rarity": 0.412860607937245160936612364821485243737697601318359375,
                "Value": "M2 Stunt Jacket",
                "ValueCount": 80
            },
            {
                "Name": "Hat",
                "Rarity": 0.7741136398823347253284055113908834755420684814453125,
                "Value": "M2 Short Mohawk",
                "ValueCount": 150
            },
            {
                "Name": "Clothes",
                "Rarity": 0.6657377302988078948686734293005429208278656005859375,
                "Value": "M2 Biker Vest",
                "ValueCount": 129
            },
            {
                "Name": "Eyes",
                "Rarity": 1.243742581410951064668779508792795240879058837890625,
                "Value": "M2 Wide Eyed",
                "ValueCount": 241
            },
            {
                "Name": "Fur",
                "Rarity": 0.59348712390978997088808455373509787023067474365234375,
                "Value": "M2 Death Bot",
                "ValueCount": 115
            },
            {
                "Name": "Clothes",
                "Rarity": 1.641120916550549591050867093144915997982025146484375,
                "Value": "M1 Smoking Jacket",
                "ValueCount": 318
            },
            {
                "Name": "Mouth",
                "Rarity": 0.996026216648604023617963321157731115818023681640625,
                "Value": "M1 Grin Multicolored",
                "ValueCount": 193
            },
            {
                "Name": "Hat",
                "Rarity": 1.3263147029984001523672532130149193108081817626953125,
                "Value": "M1 Safari",
                "ValueCount": 257
            },
            {
                "Name": "Background",
                "Rarity": 3.050007741136398831116594010381959378719329833984375,
                "Value": "M2 Aquamarine",
                "ValueCount": 591
            },
            {
                "Name": "Clothes",
                "Rarity": 0.54704030551684990069105651855352334678173065185546875,
                "Value": "M2 Puffy Vest",
                "ValueCount": 106
            },
            {
                "Name": "Hat",
                "Rarity": 0.50059348712390983049402848337194882333278656005859375,
                "Value": "M2 Bunny Ears",
                "ValueCount": 97
            },
            {
                "Name": "Clothes",
                "Rarity": 0.93925788305723278970305045731947757303714752197265625,
                "Value": "M2 Striped Tee",
                "ValueCount": 182
            },
            {
                "Name": "Fur",
                "Rarity": 2.91582804355679403585099862539209425449371337890625,
                "Value": "M2 Brown",
                "ValueCount": 565
            },
            {
                "Name": "Clothes",
                "Rarity": 1.119884399029777544143371414975263178348541259765625,
                "Value": "M1 Sleeveless Logo T",
                "ValueCount": 217
            },
            {
                "Name": "Mouth",
                "Rarity": 0.40253909273881405272987876742263324558734893798828125,
                "Value": "M2 Grin Gold Grill",
                "ValueCount": 78
            },
            {
                "Name": "Eyes",
                "Rarity": 0.79991742787841257911196635177475400269031524658203125,
                "Value": "M2 Zombie",
                "ValueCount": 155
            },
            {
                "Name": "Eyes",
                "Rarity": 0.567683335913712117104523713351227343082427978515625,
                "Value": "M2 Scumbag",
                "ValueCount": 110
            },
            {
                "Name": "Hat",
                "Rarity": 0.691541518294885637629931807168759405612945556640625,
                "Value": "M2 Bayc Hat Black",
                "ValueCount": 134
            },
            {
                "Name": "Eyes",
                "Rarity": 3.51963668266501539250157293281517922878265380859375,
                "Value": "M1 3d",
                "ValueCount": 682
            },
            {
                "Name": "Hat",
                "Rarity": 0.95474015585487947976872646904666908085346221923828125,
                "Value": "M2 Seaman's Hat",
                "ValueCount": 185
            },
            {
                "Name": "Fur",
                "Rarity": 2.430716829230531228489553541294299066066741943359375,
                "Value": "M1 Zombie",
                "ValueCount": 471
            },
            {
                "Name": "Mouth",
                "Rarity": 1.0218300046446817663792216990259476006031036376953125,
                "Value": "M2 Bored Unshaven Cigarette",
                "ValueCount": 198
            },
            {
                "Name": "Clothes",
                "Rarity": 0.6657377302988078948686734293005429208278656005859375,
                "Value": "M2 Wool Turtleneck",
                "ValueCount": 129
            },
            {
                "Name": "Earring",
                "Rarity": 1.2179387934148733219075211309245787560939788818359375,
                "Value": "M2 Gold Stud",
                "ValueCount": 236
            },
            {
                "Name": "Hat",
                "Rarity": 0.536718790318418736973171689896844327449798583984375,
                "Value": "M2 S&m Hat",
                "ValueCount": 104
            },
            {
                "Name": "Eyes",
                "Rarity": 2.48748516282190212933755901758559048175811767578125,
                "Value": "M1 Robot",
                "ValueCount": 482
            },
            {
                "Name": "Fur",
                "Rarity": 3.741549259431284468746525817550718784332275390625,
                "Value": "M1 Pink",
                "ValueCount": 725
            },
            {
                "Name": "Mouth",
                "Rarity": 0.93925788305723278970305045731947757303714752197265625,
                "Value": "M1 Bored Pipe",
                "ValueCount": 182
            },
            {
                "Name": "Hat",
                "Rarity": 0.361253031945089564391793146569398231804370880126953125,
                "Value": "M2 Stuntman Helmet",
                "ValueCount": 70
            },
            {
                "Name": "Earring",
                "Rarity": 1.9868916756979924098658329967292957007884979248046875,
                "Value": "M2 Silver Hoop",
                "ValueCount": 385
            },
            {
                "Name": "Eyes",
                "Rarity": 1.914641069308974596907546583679504692554473876953125,
                "Value": "M1 Blindfold",
                "ValueCount": 371
            },
            {
                "Name": "Hat",
                "Rarity": 0.73282757908861018147916865927982144057750701904296875,
                "Value": "M1 Bandana Blue",
                "ValueCount": 142
            },
            {
                "Name": "Eyes",
                "Rarity": 2.47716364762347129868658157647587358951568603515625,
                "Value": "M1 Eyepatch",
                "ValueCount": 480
            },
            {
                "Name": "Mouth",
                "Rarity": 3.844764411415595883880769179086200892925262451171875,
                "Value": "M1 Dumbfounded",
                "ValueCount": 745
            },
            {
                "Name": "Mouth",
                "Rarity": 0.201269546369407026364939383711316622793674468994140625,
                "Value": "M2 Bored Unshaven Bubblegum",
                "ValueCount": 39
            },
            {
                "Name": "Clothes",
                "Rarity": 0.536718790318418736973171689896844327449798583984375,
                "Value": "M1 Kings Robe",
                "ValueCount": 104
            },
            {
                "Name": "Clothes",
                "Rarity": 1.0166692470454663510537329784710891544818878173828125,
                "Value": "M1 Bayc T Red",
                "ValueCount": 197
            },
            {
                "Name": "Clothes",
                "Rarity": 1.9662486453011303044746682644472457468509674072265625,
                "Value": "M1 Wool Turtleneck",
                "ValueCount": 381
            },
            {
                "Name": "Hat",
                "Rarity": 0.78443515508076588904629034004756249487400054931640625,
                "Value": "M1 Party Hat 1",
                "ValueCount": 152
            },
            {
                "Name": "Hat",
                "Rarity": 1.4759766733756516376985246097319759428501129150390625,
                "Value": "M1 Bunny Ears",
                "ValueCount": 286
            },
            {
                "Name": "Hat",
                "Rarity": 0.996026216648604023617963321157731115818023681640625,
                "Value": "M1 Police Motorcycle Helmet",
                "ValueCount": 193
            },
            {
                "Name": "Clothes",
                "Rarity": 0.443825153532538596579115619533695280551910400390625,
                "Value": "M2 Bone Tee",
                "ValueCount": 86
            },
            {
                "Name": "Hat",
                "Rarity": 0.76379212468390356161052068273420445621013641357421875,
                "Value": "M2 Fisherman's Hat",
                "ValueCount": 148
            },
            {
                "Name": "Eyes",
                "Rarity": 0.6192909119058678246716453941189683973789215087890625,
                "Value": "M2 Blindfold",
                "ValueCount": 120
            },
            {
                "Name": "Fur",
                "Rarity": 1.150848944625070924274723438429646193981170654296875,
                "Value": "M2 Cheetah",
                "ValueCount": 223
            },
            {
                "Name": "Mouth",
                "Rarity": 0.2270733343654848523929246084662736393511295318603515625,
                "Value": "M1 Bored Unshaven Dagger",
                "ValueCount": 44
            },
            {
                "Name": "Clothes",
                "Rarity": 0.2115910615678381623272485967390821315348148345947265625,
                "Value": "M2 Kings Robe",
                "ValueCount": 41
            },
            {
                "Name": "Mouth",
                "Rarity": 1.8423904629199565619046552455984055995941162109375,
                "Value": "M1 Phoneme  Ooo",
                "ValueCount": 357
            },
            {
                "Name": "Clothes",
                "Rarity": 0.361253031945089564391793146569398231804370880126953125,
                "Value": "M2 Bayc T Red",
                "ValueCount": 70
            },
            {
                "Name": "Fur",
                "Rarity": 0.7121845486917479650657014644821174442768096923828125,
                "Value": "M2 Noise",
                "ValueCount": 138
            },
            {
                "Name": "Clothes",
                "Rarity": 0.50059348712390983049402848337194882333278656005859375,
                "Value": "M1 Admirals Coat",
                "ValueCount": 97
            },
            {
                "Name": "Clothes",
                "Rarity": 1.50178046137172938045978298760019242763519287109375,
                "Value": "M1 Leather Jacket",
                "ValueCount": 291
            },
            {
                "Name": "Fur",
                "Rarity": 1.238581823811735649343290788237936794757843017578125,
                "Value": "M1 Noise",
                "ValueCount": 240
            },
            {
                "Name": "Eyes",
                "Rarity": 3.932497290602260608949336528894491493701934814453125,
                "Value": "M2 Bored",
                "ValueCount": 762
            },
            {
                "Name": "Mouth",
                "Rarity": 0.428342880734891906513439607806503772735595703125,
                "Value": "M2 Discomfort",
                "ValueCount": 83
            },
            {
                "Name": "Eyes",
                "Rarity": 0.54187954791763426332096287296735681593418121337890625,
                "Value": "M2 X Eyes",
                "ValueCount": 105
            },
            {
                "Name": "Hat",
                "Rarity": 2.637147133199153614668830414302647113800048828125,
                "Value": "M1 Cowboy Hat",
                "ValueCount": 511
            },
            {
                "Name": "Clothes",
                "Rarity": 1.0527945502399751465105737224803306162357330322265625,
                "Value": "M1 Lab Coat",
                "ValueCount": 204
            },
            {
                "Name": "Clothes",
                "Rarity": 1.6307994013521185383552847270038910210132598876953125,
                "Value": "M1 Lumberjack Shirt",
                "ValueCount": 316
            },
            {
                "Name": "Clothes",
                "Rarity": 1.5740310677607471934180694006499834358692169189453125,
                "Value": "M1 Bayc T Black",
                "ValueCount": 305
            },
            {
                "Name": "Hat",
                "Rarity": 0.9599009134540950061165176521171815693378448486328125,
                "Value": "M1 Bayc Hat Red",
                "ValueCount": 186
            },
            {
                "Name": "Mouth",
                "Rarity": 1.5843525829591784681582566918223164975643157958984375,
                "Value": "M1 Tongue Out",
                "ValueCount": 307
            },
            {
                "Name": "Mouth",
                "Rarity": 0.7121845486917479650657014644821174442768096923828125,
                "Value": "M1 Grin Gold Grill",
                "ValueCount": 138
            },
            {
                "Name": "Hat",
                "Rarity": 0.51607575992155652055970449509914033114910125732421875,
                "Value": "M2 Spinner Hat",
                "ValueCount": 100
            },
            {
                "Name": "Fur",
                "Rarity": 3.0087216803426741762450546957552433013916015625,
                "Value": "M2 Dark Brown",
                "ValueCount": 583
            },
            {
                "Name": "Fur",
                "Rarity": 0.263198637559993786627643430620082654058933258056640625,
                "Value": "M2 Solid Gold",
                "ValueCount": 51
            },
            {
                "Name": "Clothes",
                "Rarity": 0.6141301543066521873015517485328018665313720703125,
                "Value": "M2 Vietnam Jacket",
                "ValueCount": 119
            },
            {
                "Name": "Eyes",
                "Rarity": 0.43350363833410743286123079087701626121997833251953125,
                "Value": "M2 Holographic",
                "ValueCount": 84
            },
            {
                "Name": "Eyes",
                "Rarity": 0.2993239407545027486179378684028051793575286865234375,
                "Value": "M2 Blue Beams",
                "ValueCount": 58
            },
            {
                "Name": "Clothes",
                "Rarity": 0.505754244723125356841819666442461311817169189453125,
                "Value": "M2 Prison Jumpsuit",
                "ValueCount": 98
            },
            {
                "Name": "Hat",
                "Rarity": 0.32512772875058060240149870878667570650577545166015625,
                "Value": "M2 Girl's Hair Short",
                "ValueCount": 63
            },
            {
                "Name": "Eyes",
                "Rarity": 0.57800485111214328082240854200790636241436004638671875,
                "Value": "M1 Laser Eyes",
                "ValueCount": 112
            },
            {
                "Name": "Hat",
                "Rarity": 0.81539970067605926917764236350194551050662994384765625,
                "Value": "M1 Party Hat 2",
                "ValueCount": 158
            },
            {
                "Name": "Clothes",
                "Rarity": 2.105589100479950293021147444960661232471466064453125,
                "Value": "M1 Sailor Shirt",
                "ValueCount": 408
            },
            {
                "Name": "Hat",
                "Rarity": 1.8991587965113279068418705719523131847381591796875,
                "Value": "M1 Horns",
                "ValueCount": 368
            },
            {
                "Name": "Mouth",
                "Rarity": 0.53155803271920321062538050682633183896541595458984375,
                "Value": "M1 Bored Unshaven Party Horn",
                "ValueCount": 103
            },
            {
                "Name": "Fur",
                "Rarity": 0.397378335139598470870936353094293735921382904052734375,
                "Value": "M1 Solid Gold",
                "ValueCount": 77
            },
            {
                "Name": "Mouth",
                "Rarity": 3.1222583475254168661194853484630584716796875,
                "Value": "M1 Bored Unshaven Cigarette",
                "ValueCount": 605
            },
            {
                "Name": "Clothes",
                "Rarity": 0.43350363833410743286123079087701626121997833251953125,
                "Value": "M2 Caveman Pelt",
                "ValueCount": 84
            },
            {
                "Name": "Hat",
                "Rarity": 0.660576972699592257498579783714376389980316162109375,
                "Value": "M2 Horns",
                "ValueCount": 128
            },
            {
                "Name": "Clothes",
                "Rarity": 2.97259637714816538078821395174600183963775634765625,
                "Value": "M1 Striped Tee",
                "ValueCount": 576
            },
            {
                "Name": "Mouth",
                "Rarity": 0.7276668214893946551313774762093089520931243896484375,
                "Value": "M1 Bored Unshaven Cigar",
                "ValueCount": 141
            },
            {
                "Name": "Clothes",
                "Rarity": 1.6514424317489806437464494592859409749507904052734375,
                "Value": "M1 Vietnam Jacket",
                "ValueCount": 320
            },
            {
                "Name": "Mouth",
                "Rarity": 0.11869742478195799417761691074701957404613494873046875,
                "Value": "M2 Bored Unshaven Pizza",
                "ValueCount": 23
            },
            {
                "Name": "Hat",
                "Rarity": 0.8824895494658615557881375934812240302562713623046875,
                "Value": "M1 Ww2 Pilot Helm",
                "ValueCount": 171
            },
            {
                "Name": "Mouth",
                "Rarity": 0.2064303039686225804683061824107426218688488006591796875,
                "Value": "M2 Bored Kazoo",
                "ValueCount": 40
            },
            {
                "Name": "Mouth",
                "Rarity": 0.495432729524694248635086069043609313666820526123046875,
                "Value": "M2 Phoneme Oh",
                "ValueCount": 96
            },
            {
                "Name": "Eyes",
                "Rarity": 0.95474015585487947976872646904666908085346221923828125,
                "Value": "M2 Heart",
                "ValueCount": 185
            },
            {
                "Name": "Hat",
                "Rarity": 0.660576972699592257498579783714376389980316162109375,
                "Value": "M2 Army Hat",
                "ValueCount": 128
            },
            {
                "Name": "Clothes",
                "Rarity": 1.7391753109356453688150168090942315757274627685546875,
                "Value": "M1 Tuxedo Tee",
                "ValueCount": 337
            },
            {
                "Name": "Fur",
                "Rarity": 0.676059245497238947564255795441567897796630859375,
                "Value": "M1 Trippy",
                "ValueCount": 131
            },
            {
                "Name": "Eyes",
                "Rarity": 2.96743561954894996546272523119114339351654052734375,
                "Value": "M1 Heart",
                "ValueCount": 575
            },
            {
                "Name": "Hat",
                "Rarity": 0.629612427104298877367227760259993374347686767578125,
                "Value": "M1 King's Crown",
                "ValueCount": 122
            },
            {
                "Name": "Mouth",
                "Rarity": 1.088919853434484164012019391520880162715911865234375,
                "Value": "M2 Dumbfounded",
                "ValueCount": 211
            },
            {
                "Name": "Hat",
                "Rarity": 2.2139650100634771234808795270510017871856689453125,
                "Value": "M1 Army Hat",
                "ValueCount": 429
            },
            {
                "Name": "Hat",
                "Rarity": 0.438664395933323014720173205205355770885944366455078125,
                "Value": "M2 Baby's Bonnet",
                "ValueCount": 85
            },
            {
                "Name": "Hat",
                "Rarity": 0.7895959126799814153940815231180749833583831787109375,
                "Value": "M1 Girl's Hair Pink",
                "ValueCount": 153
            },
            {
                "Name": "Clothes",
                "Rarity": 0.4902719719254786667761436547152698040008544921875,
                "Value": "M2 Toga",
                "ValueCount": 95
            },
            {
                "Name": "Mouth",
                "Rarity": 1.6359601589513339536807734475587494671344757080078125,
                "Value": "M1 Discomfort",
                "ValueCount": 317
            },
            {
                "Name": "Clothes",
                "Rarity": 1.6617639469474119184866367504582740366458892822265625,
                "Value": "M1 Toga",
                "ValueCount": 322
            },
            {
                "Name": "Clothes",
                "Rarity": 2.131392888476028257827010747860185801982879638671875,
                "Value": "M1 Hawaiian",
                "ValueCount": 413
            },
            {
                "Name": "Eyes",
                "Rarity": 3.10161531712855431663911076611839234828948974609375,
                "Value": "M1 Angry",
                "ValueCount": 601
            },
            {
                "Name": "Fur",
                "Rarity": 1.4346906125819269828269852951052598655223846435546875,
                "Value": "M2 Cream",
                "ValueCount": 278
            },
            {
                "Name": "Fur",
                "Rarity": 1.1095628838313464914477890488342382013797760009765625,
                "Value": "M2 Red",
                "ValueCount": 215
            },
            {
                "Name": "Hat",
                "Rarity": 1.3108324302007534623015772012877278029918670654296875,
                "Value": "M2 Beanie",
                "ValueCount": 254
            },
            {
                "Name": "Clothes",
                "Rarity": 1.672085462145842971182219116599299013614654541015625,
                "Value": "M1 Guayabera",
                "ValueCount": 324
            },
            {
                "Name": "Hat",
                "Rarity": 0.2993239407545027486179378684028051793575286865234375,
                "Value": "M2 Party Hat 2",
                "ValueCount": 58
            },
            {
                "Name": "Clothes",
                "Rarity": 0.32512772875058060240149870878667570650577545166015625,
                "Value": "M1 Black Suit",
                "ValueCount": 63
            },
            {
                "Name": "Clothes",
                "Rarity": 2.559735769210920164340450355666689574718475341796875,
                "Value": "M1 Black T",
                "ValueCount": 496
            },
            {
                "Name": "Clothes",
                "Rarity": 0.79991742787841257911196635177475400269031524658203125,
                "Value": "M2 Navy Striped Tee",
                "ValueCount": 155
            },
            {
                "Name": "Clothes",
                "Rarity": 1.2489033390101667020388731543789617717266082763671875,
                "Value": "M1 Caveman Pelt",
                "ValueCount": 242
            },
            {
                "Name": "Hat",
                "Rarity": 0.9134540950611549359194896169356070458889007568359375,
                "Value": "M2 Fez",
                "ValueCount": 177
            },
            {
                "Name": "Mouth",
                "Rarity": 0.423182123135676324654497193478164263069629669189453125,
                "Value": "M2 Tongue Out",
                "ValueCount": 82
            },
            {
                "Name": "Fur",
                "Rarity": 0.8050781854776281054597575348452664911746978759765625,
                "Value": "M2 Dmt",
                "ValueCount": 156
            },
            {
                "Name": "Eyes",
                "Rarity": 0.8670072766682148657224615817540325224399566650390625,
                "Value": "M1 Cyborg",
                "ValueCount": 168
            },
            {
                "Name": "Mouth",
                "Rarity": 0.185787273571760336299263371984125114977359771728515625,
                "Value": "M1 Bored Unshaven Pizza",
                "ValueCount": 36
            },
            {
                "Name": "Clothes",
                "Rarity": 0.485111214326263084917201240386930294334888458251953125,
                "Value": "M2 Leather Jacket",
                "ValueCount": 94
            },
            {
                "Name": "Eyes",
                "Rarity": 1.0837590958352685266419257459347136318683624267578125,
                "Value": "M2 Angry",
                "ValueCount": 210
            },
            {
                "Name": "Clothes",
                "Rarity": 0.1806265159725447821958965732846991159021854400634765625,
                "Value": "M2 Blue Dress",
                "ValueCount": 35
            },
            {
                "Name": "Eyes",
                "Rarity": 1.5585487949631005033523933889227919280529022216796875,
                "Value": "M2 Closed",
                "ValueCount": 302
            },
            {
                "Name": "Fur",
                "Rarity": 1.393404551788202550000050905509851872920989990234375,
                "Value": "M2 Tan",
                "ValueCount": 270
            },
            {
                "Name": "Hat",
                "Rarity": 0.92377561025958609963737444559228606522083282470703125,
                "Value": "M2 Cowboy Hat",
                "ValueCount": 179
            },
            {
                "Name": "Fur",
                "Rarity": 3.52479744026423080782706165337003767490386962890625,
                "Value": "M1 Blue",
                "ValueCount": 683
            },
            {
                "Name": "Hat",
                "Rarity": 0.56252257831449659075673253028071485459804534912109375,
                "Value": "M1 Trippy Captain's Hat",
                "ValueCount": 109
            },
            {
                "Name": "Hat",
                "Rarity": 1.4140475821850648774358205628232099115848541259765625,
                "Value": "M1 Sushi Chef Headband",
                "ValueCount": 274
            },
            {
                "Name": "Fur",
                "Rarity": 1.4604944005780049476328485980047844350337982177734375,
                "Value": "M1 Death Bot",
                "ValueCount": 283
            },
            {
                "Name": "Earring",
                "Rarity": 0.54187954791763426332096287296735681593418121337890625,
                "Value": "M2 Cross",
                "ValueCount": 105
            },
            {
                "Name": "Mouth",
                "Rarity": 0.89281106466429271950602242213790304958820343017578125,
                "Value": "M1 Bored Bubblegum",
                "ValueCount": 173
            },
            {
                "Name": "Background",
                "Rarity": 3.323527893894823836973273500916548073291778564453125,
                "Value": "M2 Yellow",
                "ValueCount": 644
            },
            {
                "Name": "Mouth",
                "Rarity": 0.1961087887701914722615725850118906237185001373291015625,
                "Value": "M2 Bored Pizza",
                "ValueCount": 38
            },
            {
                "Name": "Eyes",
                "Rarity": 0.90829333746193940957169843386509455740451812744140625,
                "Value": "M2 Sunglasses",
                "ValueCount": 176
            },
            {
                "Name": "Hat",
                "Rarity": 0.278680910357640476693319442347274161875247955322265625,
                "Value": "M2 Girl's Hair Pink",
                "ValueCount": 54
            },
            {
                "Name": "Clothes",
                "Rarity": 0.2838416679568560585522618566756136715412139892578125,
                "Value": "M2 Service",
                "ValueCount": 55
            },
            {
                "Name": "Eyes",
                "Rarity": 1.285028642204675719540318823419511318206787109375,
                "Value": "M2 Sad",
                "ValueCount": 249
            },
            {
                "Name": "Clothes",
                "Rarity": 1.0785983382360531113164370253798551857471466064453125,
                "Value": "M1 Service",
                "ValueCount": 209
            },
            {
                "Name": "Mouth",
                "Rarity": 0.37157454714352067259852674396825022995471954345703125,
                "Value": "M1 Bored Pizza",
                "ValueCount": 72
            },
            {
                "Name": "Fur",
                "Rarity": 1.1869742478195799417761691074701957404613494873046875,
                "Value": "M2 Pink",
                "ValueCount": 230
            },
            {
                "Name": "Background",
                "Rarity": 3.019043195541105450985241986927576363086700439453125,
                "Value": "M2 Gray",
                "ValueCount": 585
            },
            {
                "Name": "Eyes",
                "Rarity": 1.1560097022242865616448170840158127248287200927734375,
                "Value": "M2 3d",
                "ValueCount": 224
            },
            {
                "Name": "Clothes",
                "Rarity": 1.780461371729369801641951198689639568328857421875,
                "Value": "M1 Bone Tee",
                "ValueCount": 345
            },
            {
                "Name": "Fur",
                "Rarity": 1.7082107653403519886836647856398485600948333740234375,
                "Value": "M1 Dmt",
                "ValueCount": 331
            },
            {
                "Name": "Clothes",
                "Rarity": 0.366413789544305090739584329639910720288753509521484375,
                "Value": "M2 Hip Hop",
                "ValueCount": 71
            },
            {
                "Name": "Eyes",
                "Rarity": 1.19213500541879557914626275305636227130889892578125,
                "Value": "M1 Holographic",
                "ValueCount": 231
            },
            {
                "Name": "Background",
                "Rarity": 3.03452546833875214105091799865476787090301513671875,
                "Value": "M2 Purple",
                "ValueCount": 588
            },
            {
                "Name": "Earring",
                "Rarity": 1.2179387934148733219075211309245787560939788818359375,
                "Value": "M2 Gold Hoop",
                "ValueCount": 236
            },
            {
                "Name": "Fur",
                "Rarity": 0.443825153532538596579115619533695280551910400390625,
                "Value": "M2 Trippy",
                "ValueCount": 86
            },
            {
                "Name": "Clothes",
                "Rarity": 0.52639727511998757325528686124016530811786651611328125,
                "Value": "M2 Tanktop",
                "ValueCount": 102
            },
            {
                "Name": "Mouth",
                "Rarity": 0.485111214326263084917201240386930294334888458251953125,
                "Value": "M2 Phoneme L",
                "ValueCount": 94
            },
            {
                "Name": "Clothes",
                "Rarity": 0.2993239407545027486179378684028051793575286865234375,
                "Value": "M2 Sleeveless Logo T",
                "ValueCount": 58
            },
            {
                "Name": "Fur",
                "Rarity": 0.72250606389017912878358629313879646360874176025390625,
                "Value": "M2 Zombie",
                "ValueCount": 140
            },
            {
                "Name": "Mouth",
                "Rarity": 0.361253031945089564391793146569398231804370880126953125,
                "Value": "M2 Grin Diamond Grill",
                "ValueCount": 70
            },
            {
                "Name": "Clothes",
                "Rarity": 0.56252257831449659075673253028071485459804534912109375,
                "Value": "M2 Tuxedo Tee",
                "ValueCount": 109
            },
            {
                "Name": "Clothes",
                "Rarity": 1.56370955256231614072248703450895845890045166015625,
                "Value": "M1 Puffy Vest",
                "ValueCount": 303
            },
            {
                "Name": "Clothes",
                "Rarity": 0.8360427310729214855911095582996495068073272705078125,
                "Value": "M1 Space Suit",
                "ValueCount": 162
            },
            {
                "Name": "Hat",
                "Rarity": 1.7133715229395676260537584312260150909423828125,
                "Value": "M1 Vietnam Era Helmet",
                "ValueCount": 332
            },
            {
                "Name": "Clothes",
                "Rarity": 1.13536667182742423420904742670245468616485595703125,
                "Value": "M1 Tie Dye",
                "ValueCount": 220
            },
            {
                "Name": "Mouth",
                "Rarity": 0.12385818238117356215877151726090232841670513153076171875,
                "Value": "M2 Bored Unshaven Dagger",
                "ValueCount": 24
            },
            {
                "Name": "Hat",
                "Rarity": 0.366413789544305090739584329639910720288753509521484375,
                "Value": "M2 King's Crown",
                "ValueCount": 71
            },
            {
                "Name": "Clothes",
                "Rarity": 0.32512772875058060240149870878667570650577545166015625,
                "Value": "M2 Cowboy Shirt",
                "ValueCount": 63
            },
            {
                "Name": "Hat",
                "Rarity": 0.95474015585487947976872646904666908085346221923828125,
                "Value": "M1 Prussian Helmet",
                "ValueCount": 185
            },
            {
                "Name": "Clothes",
                "Rarity": 0.9444186406564483160508416403899900615215301513671875,
                "Value": "M1 Cowboy Shirt",
                "ValueCount": 183
            },
            {
                "Name": "Fur",
                "Rarity": 1.2024565206172266318418451191973872482776641845703125,
                "Value": "M2 Blue",
                "ValueCount": 233
            },
            {
                "Name": "Fur",
                "Rarity": 9.129380193012334387958617298863828182220458984375,
                "Value": "M1 Black",
                "ValueCount": 1769
            },
            {
                "Name": "Eyes",
                "Rarity": 1.9662486453011303044746682644472457468509674072265625,
                "Value": "M1 X Eyes",
                "ValueCount": 381
            },
            {
                "Name": "Eyes",
                "Rarity": 0.59348712390978997088808455373509787023067474365234375,
                "Value": "M2 Hypnotized",
                "ValueCount": 115
            },
            {
                "Name": "Mouth",
                "Rarity": 0.154822727976466956167911348529742099344730377197265625,
                "Value": "M2 Bored Unshaven Party Horn",
                "ValueCount": 30
            },
            {
                "Name": "Mouth",
                "Rarity": 0.65541621510037673115078860064386390149593353271484375,
                "Value": "M1 Bored Party Horn",
                "ValueCount": 127
            },
            {
                "Name": "Background",
                "Rarity": 9.5267585281519320261622851830907166004180908203125,
                "Value": "M1 Purple",
                "ValueCount": 1846
            },
            {
                "Name": "Clothes",
                "Rarity": 0.55220106311606542703884770162403583526611328125,
                "Value": "M2 Guayabera",
                "ValueCount": 107
            },
            {
                "Name": "Clothes",
                "Rarity": 0.38705681994116736266420275569544173777103424072265625,
                "Value": "M2 Tie Dye",
                "ValueCount": 75
            },
            {
                "Name": "Hat",
                "Rarity": 0.2373948495639159883552338214940391480922698974609375,
                "Value": "M2 Police Motorcycle Helmet",
                "ValueCount": 46
            },
            {
                "Name": "Hat",
                "Rarity": 0.7431490942870413451970534879365004599094390869140625,
                "Value": "M2 Bayc Flipped Brim",
                "ValueCount": 144
            },
            {
                "Name": "Hat",
                "Rarity": 0.9599009134540950061165176521171815693378448486328125,
                "Value": "M1 Faux Hawk",
                "ValueCount": 186
            },
            {
                "Name": "Mouth",
                "Rarity": 0.1909480311709758904026301706835511140525341033935546875,
                "Value": "M2 Bored Unshaven Kazoo",
                "ValueCount": 37
            },
            {
                "Name": "Background",
                "Rarity": 9.129380193012334387958617298863828182220458984375,
                "Value": "M1 Army Green",
                "ValueCount": 1769
            },
            {
                "Name": "Clothes",
                "Rarity": 0.2683593951592093684865858449484221637248992919921875,
                "Value": "M2 Rainbow Suspenders",
                "ValueCount": 52
            },
            {
                "Name": "Mouth",
                "Rarity": 1.243742581410951064668779508792795240879058837890625,
                "Value": "M1 Phoneme Wah",
                "ValueCount": 241
            },
            {
                "Name": "Background",
                "Rarity": 3.019043195541105450985241986927576363086700439453125,
                "Value": "M2 Orange",
                "ValueCount": 585
            },
            {
                "Name": "Clothes",
                "Rarity": 0.35609227434587398253285073224105872213840484619140625,
                "Value": "M2 Space Suit",
                "ValueCount": 69
            },
            {
                "Name": "Eyes",
                "Rarity": 0.841203488672137122961203203885816037654876708984375,
                "Value": "M2 Eyepatch",
                "ValueCount": 163
            },
            {
                "Name": "Hat",
                "Rarity": 0.7431490942870413451970534879365004599094390869140625,
                "Value": "M2 Halo",
                "ValueCount": 144
            },
            {
                "Name": "Mouth",
                "Rarity": 0.381896062341951780805260341367102228105068206787109375,
                "Value": "M2 Bored Pipe",
                "ValueCount": 74
            },
            {
                "Name": "Hat",
                "Rarity": 2.7661660731795425505197272286750376224517822265625,
                "Value": "M1 Fez",
                "ValueCount": 536
            },
            {
                "Name": "Clothes",
                "Rarity": 1.1714919750219332517104930957430042326450347900390625,
                "Value": "M1 Tweed Suit",
                "ValueCount": 227
            },
            {
                "Name": "Eyes",
                "Rarity": 6.53867987820612039939760506968013942241668701171875,
                "Value": "M1 Bloodshot",
                "ValueCount": 1267
            },
            {
                "Name": "Hat",
                "Rarity": 3.173865923517572351642002104199491441249847412109375,
                "Value": "M1 Seaman's Hat",
                "ValueCount": 615
            },
            {
                "Name": "Earring",
                "Rarity": 1.6359601589513339536807734475587494671344757080078125,
                "Value": "M1 Diamond Stud",
                "ValueCount": 317
            },
            {
                "Name": "Clothes",
                "Rarity": 0.64509469990194556743290377198718488216400146484375,
                "Value": "M2 Bayc T Black",
                "ValueCount": 125
            },
            {
                "Name": "Hat",
                "Rarity": 0.6502554575011612048029974175733514130115509033203125,
                "Value": "M2 Bowler",
                "ValueCount": 126
            },
            {
                "Name": "Eyes",
                "Rarity": 12.85028642204675719540318823419511318206787109375,
                "Value": "M1 Bored",
                "ValueCount": 2490
            },
            {
                "Name": "Mouth",
                "Rarity": 0.51607575992155652055970449509914033114910125732421875,
                "Value": "M1 Bored Unshaven Bubblegum",
                "ValueCount": 100
            },
            {
                "Name": "Fur",
                "Rarity": 2.616504102802291509277665682020597159862518310546875,
                "Value": "M2 Black",
                "ValueCount": 507
            },
            {
                "Name": "Mouth",
                "Rarity": 1.6049956133560405735494214241043664515018463134765625,
                "Value": "M2 Bored Cigarette",
                "ValueCount": 311
            },
            {
                "Name": "Hat",
                "Rarity": 2.373948495639159883552338214940391480922698974609375,
                "Value": "M1 Short Mohawk",
                "ValueCount": 460
            },
            {
                "Name": "Mouth",
                "Rarity": 1.7133715229395676260537584312260150909423828125,
                "Value": "M1 Phoneme Oh",
                "ValueCount": 332
            },
            {
                "Name": "Background",
                "Rarity": 8.938432161841358691845016437582671642303466796875,
                "Value": "M1 Gray",
                "ValueCount": 1732
            },
            {
                "Name": "Mouth",
                "Rarity": 2.472002890024255439271883005858398973941802978515625,
                "Value": "M1 Phoneme Vuh",
                "ValueCount": 479
            },
            {
                "Name": "Background",
                "Rarity": 9.53708004335036463317010202445089817047119140625,
                "Value": "M1 Aquamarine",
                "ValueCount": 1848
            },
            {
                "Name": "Fur",
                "Rarity": 3.6176910770501109482211177237331867218017578125,
                "Value": "M1 Red",
                "ValueCount": 701
            },
            {
                "Name": "Clothes",
                "Rarity": 1.919801826908190012233035304234363138675689697265625,
                "Value": "M1 Biker Vest",
                "ValueCount": 372
            },
            {
                "Name": "Clothes",
                "Rarity": 2.368787738039944468226849494385533034801483154296875,
                "Value": "M1 Navy Striped Tee",
                "ValueCount": 459
            },
            {
                "Name": "Fur",
                "Rarity": 10.0892811064664300602089497260749340057373046875,
                "Value": "M1 Dark Brown",
                "ValueCount": 1955
            },
            {
                "Name": "Hat",
                "Rarity": 4.2782680497497036498089073575101792812347412109375,
                "Value": "M1 Beanie",
                "ValueCount": 829
            },
            {
                "Name": "Background",
                "Rarity": 9.1809877690044903175703439046628773212432861328125,
                "Value": "M1 Blue",
                "ValueCount": 1779
            },
            {
                "Name": "Clothes",
                "Rarity": 1.945605614904267977038898607133887708187103271484375,
                "Value": "M1 Sleeveless T",
                "ValueCount": 377
            },
            {
                "Name": "Earring",
                "Rarity": 1.1456881870258552869046297928434796631336212158203125,
                "Value": "M1 Cross",
                "ValueCount": 222
            },
            {
                "Name": "Eyes",
                "Rarity": 3.566083501057955462698600967996753752231597900390625,
                "Value": "M1 Coins",
                "ValueCount": 691
            },
            {
                "Name": "Mouth",
                "Rarity": 1.837229705320741146579166525043547153472900390625,
                "Value": "M1 Phoneme L",
                "ValueCount": 356
            },
            {
                "Name": "Background",
                "Rarity": 9.1913092842029211482213213457725942134857177734375,
                "Value": "M1 New Punk Blue",
                "ValueCount": 1781
            },
            {
                "Name": "Eyes",
                "Rarity": 2.42555607163131536907485497067682445049285888671875,
                "Value": "M1 Zombie",
                "ValueCount": 470
            },
            {
                "Name": "Fur",
                "Rarity": 3.71574547143520650394066251465119421482086181640625,
                "Value": "M1 Gray",
                "ValueCount": 720
            },
            {
                "Name": "Hat",
                "Rarity": 1.7133715229395676260537584312260150909423828125,
                "Value": "M1 S&m Hat",
                "ValueCount": 332
            },
            {
                "Name": "Mouth",
                "Rarity": 5.2175259328069358844004455022513866424560546875,
                "Value": "M1 Bored Cigarette",
                "ValueCount": 1011
            },
            {
                "Name": "Fur",
                "Rarity": 3.086133044330907626573434754391200840473175048828125,
                "Value": "M1 Cheetah",
                "ValueCount": 598
            },
            {
                "Name": "Fur",
                "Rarity": 1.176652732621148889080586741329170763492584228515625,
                "Value": "M2 Gray",
                "ValueCount": 228
            },
            {
                "Name": "Eyes",
                "Rarity": 5.25365123600144467985728624626062810420989990234375,
                "Value": "M1 Closed",
                "ValueCount": 1018
            },
            {
                "Name": "Clothes",
                "Rarity": 1.764979098931723111576275186962448060512542724609375,
                "Value": "M1 Prison Jumpsuit",
                "ValueCount": 342
            },
            {
                "Name": "Clothes",
                "Rarity": 1.56370955256231614072248703450895845890045166015625,
                "Value": "M1 Black Holes T",
                "ValueCount": 303
            },
            {
                "Name": "Background",
                "Rarity": 3.03452546833875214105091799865476787090301513671875,
                "Value": "M2 New Punk Blue",
                "ValueCount": 588
            },
            {
                "Name": "Earring",
                "Rarity": 2.131392888476028257827010747860185801982879638671875,
                "Value": "M2 Silver Stud",
                "ValueCount": 413
            },
            {
                "Name": "Eyes",
                "Rarity": 0.8824895494658615557881375934812240302562713623046875,
                "Value": "M2 Robot",
                "ValueCount": 171
            },
            {
                "Name": "Hat",
                "Rarity": 0.68638076069567011128214062409824691712856292724609375,
                "Value": "M2 Sea Captain's Hat",
                "ValueCount": 133
            },
            {
                "Name": "Mouth",
                "Rarity": 0.63993394230273004108511258891667239367961883544921875,
                "Value": "M1 Grin Diamond Grill",
                "ValueCount": 124
            },
            {
                "Name": "Clothes",
                "Rarity": 1.5224234917685917078955526449135504662990570068359375,
                "Value": "M1 Bone Necklace",
                "ValueCount": 295
            },
            {
                "Name": "Background",
                "Rarity": 9.80543943850957333552287309430539608001708984375,
                "Value": "M1 Yellow",
                "ValueCount": 1900
            },
            {
                "Name": "Eyes",
                "Rarity": 5.6974763895339837205256117158569395542144775390625,
                "Value": "M1 Sleepy",
                "ValueCount": 1104
            },
            {
                "Name": "Hat",
                "Rarity": 2.317180162047788538615122888586483895778656005859375,
                "Value": "M1 Commie Hat",
                "ValueCount": 449
            },
            {
                "Name": "Eyes",
                "Rarity": 3.117097589926201006704786777845583856105804443359375,
                "Value": "M1 Crazy",
                "ValueCount": 604
            },
            {
                "Name": "Fur",
                "Rarity": 9.965422924085256539683541632257401943206787109375,
                "Value": "M1 Brown",
                "ValueCount": 1931
            },
            {
                "Name": "Clothes",
                "Rarity": 0.99086545904938849727017213808721862733364105224609375,
                "Value": "M1 Rainbow Suspenders",
                "ValueCount": 192
            },
            {
                "Name": "Hat",
                "Rarity": 1.176652732621148889080586741329170763492584228515625,
                "Value": "M1 Baby's Bonnet",
                "ValueCount": 228
            },
            {
                "Name": "Clothes",
                "Rarity": 1.3417969757960468424329292247421108186244964599609375,
                "Value": "M1 Stunt Jacket",
                "ValueCount": 260
            },
            {
                "Name": "Eyes",
                "Rarity": 4.00990865459049405927771658753044903278350830078125,
                "Value": "M1 Sad",
                "ValueCount": 777
            },
            {
                "Name": "Mouth",
                "Rarity": 1.9559271301026990297344809732749126851558685302734375,
                "Value": "M1 Rage",
                "ValueCount": 379
            },
            {
                "Name": "Background",
                "Rarity": 9.4390256489652681892721375334076583385467529296875,
                "Value": "M1 Orange",
                "ValueCount": 1829
            },
            {
                "Name": "Earring",
                "Rarity": 3.390617742684626012561466268380172550678253173828125,
                "Value": "M1 Gold Hoop",
                "ValueCount": 657
            },
            {
                "Name": "Fur",
                "Rarity": 5.6871548743355528898746342747472226619720458984375,
                "Value": "M1 Golden Brown",
                "ValueCount": 1102
            },
            {
                "Name": "Hat",
                "Rarity": 1.8475512205191721992747488911845721304416656494140625,
                "Value": "M1 Bayc Flipped Brim",
                "ValueCount": 358
            },
            {
                "Name": "Mouth",
                "Rarity": 0.55220106311606542703884770162403583526611328125,
                "Value": "M1 Bored Kazoo",
                "ValueCount": 107
            },
            {
                "Name": "Eyes",
                "Rarity": 0.88765030706507719315823123906739056110382080078125,
                "Value": "M2 Crazy",
                "ValueCount": 172
            },
            {
                "Name": "Mouth",
                "Rarity": 1.74949682613407642151059917523525655269622802734375,
                "Value": "M2 Grin",
                "ValueCount": 339
            },
            {
                "Name": "Mouth",
                "Rarity": 11.5239717190483563769021202460862696170806884765625,
                "Value": "M1 Bored Unshaven",
                "ValueCount": 2233
            },
            {
                "Name": "Fur",
                "Rarity": 4.47437683851989476124799693934619426727294921875,
                "Value": "M1 Tan",
                "ValueCount": 867
            },
            {
                "Name": "Hat",
                "Rarity": 2.250090313257986363026930121122859418392181396484375,
                "Value": "M1 Sea Captain's Hat",
                "ValueCount": 436
            },
            {
                "Name": "Mouth",
                "Rarity": 0.485111214326263084917201240386930294334888458251953125,
                "Value": "M1 Bored Unshaven Kazoo",
                "ValueCount": 94
            },
            {
                "Name": "Clothes",
                "Rarity": 1.367600763792124585194187602610327303409576416015625,
                "Value": "M1 Work Vest",
                "ValueCount": 265
            },
            {
                "Name": "Eyes",
                "Rarity": 4.03055168498735572057967146974988281726837158203125,
                "Value": "M1 Wide Eyed",
                "ValueCount": 781
            },
            {
                "Name": "Fur",
                "Rarity": 2.9364710739536565853313732077367603778839111328125,
                "Value": "M1 White",
                "ValueCount": 569
            },
            {
                "Name": "Mouth",
                "Rarity": 5.341384115188109404925853596068918704986572265625,
                "Value": "M1 Grin",
                "ValueCount": 1035
            },
            {
                "Name": "Hat",
                "Rarity": 1.1250451566289931815134650605614297091960906982421875,
                "Value": "M1 Girl's Hair Short",
                "ValueCount": 218
            },
            {
                "Name": "Earring",
                "Rarity": 70.454662744490889281223644502460956573486328125,
                "Value": "",
                "ValueCount": 13652
            },
            {
                "Name": "Fur",
                "Rarity": 0.89281106466429271950602242213790304958820343017578125,
                "Value": "",
                "ValueCount": 173
            },
            {
                "Name": "Hat",
                "Rarity": 22.392527222996335467541939578950405120849609375,
                "Value": "",
                "ValueCount": 4339
            },
            {
                "Name": "Mouth",
                "Rarity": 0.89281106466429271950602242213790304958820343017578125,
                "Value": "",
                "ValueCount": 173
            },
            {
                "Name": "Name",
                "Rarity": 99.9432316664086357604901422746479511260986328125,
                "Value": "",
                "ValueCount": 19366
            },
            {
                "Name": "Clothes",
                "Rarity": 18.800639933942303372305104858241975307464599609375,
                "Value": "",
                "ValueCount": 3643
            },
            {
                "Name": "Eyes",
                "Rarity": 0.89281106466429271950602242213790304958820343017578125,
                "Value": "",
                "ValueCount": 173
            },
            {
                "Name": "Background",
                "Rarity": 0.89281106466429271950602242213790304958820343017578125,
                "Value": "",
                "ValueCount": 173
            }
        ],
        "CirculatingSupply": 19377,
        "Name": "Mutant Ape Yacht Club",
        "Stats": {
            "Average": 12.615470588235293547541004954837262630462646484375,
            "Ceiling": 17.5,
            "Floor": 11.5129999999999999005240169935859739780426025390625,
            "TotalSales": 17,
            "Volume": 214.462999999999993860910763032734394073486328125
        },
        "Symbol": "MAYC"
    }
}`,
		},
	}

	for name, tt := range testCases {
		t.Run(name, func(t *testing.T) {
			res, err := c.GetCollectionDetails(context.TODO(), tt.address)

			assert.Nil(t, err)
			data, err := json.Marshal(res)
			assert.Nil(t, err)
			assert.JSONEq(t, tt.expect, string(data))
		})
	}
}
