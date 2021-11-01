package problem1043

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []int
	K   int
	ans int
}{

	{
		[]int{322257, 540837, 18171, 364098, 161511, 717990, 678976, 295957, 525576, 912910, 302841, 807371, 774141, 973354, 343843, 742131, 839995, 111282, 772395, 5921, 122534, 750377, 683899, 734664, 184593, 789400, 626212, 809214, 934136, 163420, 452442, 713789, 269552, 932607, 172046, 89529, 277733, 222352, 957023, 867042, 160651, 259497, 476642, 978097, 212742, 714383, 932052, 399211, 894147, 766649, 409010, 763214, 459126, 833028, 777826, 294076, 17566, 538156, 124552, 641, 606726, 429007, 507307, 360562, 133147, 833321, 496061, 966033, 801581, 695725, 323273, 888523, 260890, 835191, 305749, 410437, 505257, 235238, 535199, 66423, 63758, 654191, 288667, 760215, 837132, 41380, 256744, 474781, 302571, 486974, 93945, 417524, 733278, 338240, 165821, 938705, 722138, 676130, 887649, 130572, 585556, 126289, 508926, 355353, 812298, 138788, 751157, 20720, 832600, 562373, 529647, 819221, 57234, 345991, 129691, 240013, 224203, 436010, 969914, 320771, 512865, 735790, 299628, 724880, 359109, 295235, 352639, 514634, 527186, 352191, 958072, 281078, 286045, 983549, 222715, 419519, 234599, 439778, 344828, 901962, 266790, 572029, 551645, 775030, 433586, 852788, 654154, 709948, 707145, 940333, 434841, 548197, 535555, 302255, 99577, 468017, 258874, 244736, 779744, 582486, 66596, 582593, 206343, 571625, 543211, 948901, 259730, 720218, 391355, 4441, 512392, 256717, 264162, 689384, 108804, 743139, 456167, 622153, 34594, 642958, 682978, 492462, 960588, 544889, 59853, 618158, 826481, 46500, 848227, 441833, 60210, 971753, 993798, 106622, 265063, 323422, 150184, 927954, 7368, 233027, 803151, 715037, 741914, 576419, 862192, 47536, 876455, 190304, 672899, 263957, 308829, 925787, 99541, 311184, 654650, 247786, 236849, 372981, 976532, 975655, 893372, 744998, 33851, 914471, 790591, 958015, 559603, 597950, 283947, 386643, 168667, 867857, 105367, 528454, 73767, 928997, 534173, 723272, 416051, 878341, 753212, 446378, 375998, 65721, 866187, 33706, 449084, 397460, 850849, 111453, 747319, 501296, 115045, 712431, 34914, 744459, 665120, 548829, 680009, 889090, 804508, 1131, 216652, 592160, 619079, 250083, 451008, 387734, 221043, 34718, 503942, 807191, 890799, 403017, 558917, 187498, 753156, 161889, 326616, 748336, 753942, 648694, 225315, 388228, 418217, 916129, 435637, 152945, 353274, 51543, 819598, 320869, 329985, 708486, 555732, 291431, 566179, 195921, 708680, 723429, 62059, 617193, 713490, 839672, 427061, 787674, 658480, 758624, 426935, 87777, 146592, 143149, 691270, 228944, 21659, 912288, 405247, 274058, 156786, 742941, 346777, 474383, 355448, 568771, 666124, 677562, 765587, 858179, 528256, 961677, 611111, 818806, 897460, 208616, 659738, 64081, 207641, 889609, 892033, 133501, 266857, 621606, 794935, 396372, 567053, 52331, 582094, 104977, 569902, 496245, 390154, 550784, 55238, 641297, 23741, 860725, 861067, 283503, 924066, 404874, 392211, 897657, 46748, 490894, 304656, 751827, 221437, 730461, 421321, 695259, 158087, 259472, 14837, 744332, 850752, 149834, 924528, 43174, 374614, 567583, 410071, 636843, 783718, 844101, 951446, 703700, 201046, 511629, 145996, 54378, 241686, 285229, 2299, 992962, 948612, 215175, 710462, 977662, 184503, 939465, 439086, 18999, 536186, 659815, 203013, 965787, 255983, 456806, 840014, 261442, 770277, 955892, 940019, 884343, 686299, 676304, 716247, 11141, 236451, 124411, 478090, 680125, 175743, 178750, 461502, 633006, 694179, 47960, 21926, 320034, 524028, 732852, 603541, 799317, 334915, 276857, 253882, 262465, 828135, 394678, 384312, 922363, 222351, 746409, 20062, 157582, 131931, 252307, 626638, 78825, 767279, 505875, 153744, 975663, 64725, 826981, 427925, 926691, 748580, 289254, 3199, 561583, 109612, 470102, 130143, 583840, 224354, 775172, 609376, 235902, 351312, 271143, 575355, 324896, 787099, 855052, 748572, 999311, 926808, 139753, 866799, 537768, 788685, 199740, 950081, 835325, 162448, 518463, 781043, 752546, 985988, 237700, 801304, 846762, 254587, 448003, 91525, 308020, 237083, 12293},
		30,
		483667362,
	},

	{
		[]int{322257, 540837, 18171, 364098, 161511, 717990, 678976, 295957, 525576, 912910, 302841, 807371, 774141, 973354, 343843, 742131, 839995, 111282, 772395, 5921, 122534, 750377, 683899, 734664, 184593, 789400, 626212, 809214, 934136, 163420, 452442, 713789, 269552, 932607, 172046, 89529, 277733, 222352, 957023, 867042, 160651, 259497, 476642, 978097, 212742, 714383, 932052, 399211, 894147, 766649, 409010, 763214, 459126, 833028, 777826, 294076, 17566, 538156, 124552, 641, 606726, 429007, 507307, 360562, 133147, 833321, 496061, 966033, 801581, 695725, 323273, 888523, 260890, 835191, 305749, 410437, 505257, 235238, 535199, 66423, 63758, 654191, 288667, 760215, 837132, 41380, 256744, 474781, 302571, 486974, 93945, 417524, 733278, 338240, 165821, 938705, 722138, 676130, 887649, 130572, 585556, 126289, 508926, 355353, 812298, 138788, 751157, 20720, 832600, 562373, 529647, 819221, 57234, 345991, 129691, 240013, 224203, 436010, 969914, 320771, 512865, 735790, 299628, 724880, 359109, 295235, 352639, 514634, 527186, 352191, 958072, 281078, 286045, 983549, 222715, 419519, 234599, 439778, 344828, 901962, 266790, 572029, 551645, 775030, 433586, 852788, 654154, 709948, 707145, 940333, 434841, 548197, 535555, 302255, 99577, 468017, 258874, 244736, 779744, 582486, 66596, 582593, 206343, 571625, 543211, 948901, 259730, 720218, 391355, 4441, 512392, 256717, 264162, 689384, 108804, 743139, 456167, 622153, 34594, 642958, 682978, 492462, 960588, 544889, 59853, 618158, 826481, 46500, 848227, 441833, 60210, 971753, 993798, 106622, 265063, 323422, 150184, 927954, 7368, 233027, 803151, 715037, 741914, 576419, 862192, 47536, 876455, 190304, 672899, 263957, 308829, 925787, 99541, 311184, 654650, 247786, 236849, 372981, 976532, 975655, 893372, 744998, 33851, 914471, 790591, 958015, 559603, 597950, 283947, 386643, 168667, 867857, 105367, 528454, 73767, 928997, 534173, 723272, 416051, 878341, 753212, 446378, 375998, 65721, 866187, 33706, 449084, 397460, 850849, 111453, 747319, 501296, 115045, 712431, 34914, 744459, 665120, 548829, 680009, 889090, 804508, 1131, 216652, 592160, 619079, 250083, 451008, 387734, 221043, 34718, 503942, 807191, 890799, 403017, 558917, 187498, 753156, 161889, 326616, 748336, 753942, 648694, 225315, 388228, 418217, 916129, 435637, 152945, 353274, 51543, 819598, 320869, 329985, 708486, 555732, 291431, 566179, 195921, 708680, 723429, 62059, 617193, 713490, 839672, 427061, 787674, 658480, 758624, 426935, 87777, 146592, 143149, 691270, 228944, 21659, 912288, 405247, 274058, 156786, 742941, 346777, 474383, 355448, 568771, 666124, 677562, 765587, 858179, 528256, 961677, 611111, 818806, 897460, 208616, 659738, 64081, 207641, 889609, 892033, 133501, 266857, 621606, 794935, 396372, 567053, 52331, 582094, 104977, 569902, 496245, 390154, 550784, 55238, 641297, 23741, 860725, 861067, 283503, 924066, 404874, 392211, 897657, 46748, 490894, 304656, 751827, 221437, 730461, 421321, 695259, 158087, 259472, 14837, 744332, 850752, 149834, 924528, 43174, 374614, 567583, 410071, 636843, 783718, 844101, 951446, 703700, 201046, 511629, 145996, 54378, 241686, 285229, 2299, 992962, 948612, 215175, 710462, 977662, 184503, 939465, 439086, 18999, 536186, 659815, 203013, 965787, 255983, 456806, 840014, 261442, 770277, 955892, 940019, 884343, 686299, 676304, 716247, 11141, 236451, 124411, 478090, 680125, 175743, 178750, 461502, 633006, 694179, 47960, 21926, 320034, 524028, 732852, 603541, 799317, 334915, 276857, 253882, 262465, 828135, 394678, 384312, 922363, 222351, 746409, 20062, 157582, 131931, 252307, 626638, 78825, 767279, 505875, 153744, 975663, 64725, 826981, 427925, 926691, 748580, 289254, 3199, 561583, 109612, 470102, 130143, 583840, 224354, 775172, 609376, 235902, 351312, 271143, 575355, 324896, 787099, 855052, 748572, 999311, 926808, 139753, 866799, 537768, 788685, 199740, 950081, 835325, 162448, 518463, 781043, 752546, 985988, 237700, 801304, 846762, 254587, 448003, 91525, 308020, 237083, 12293},
		1,
		247363899,
	},

	{
		[]int{10, 9, 3, 2},
		2,
		30,
	},

	{
		[]int{1, 15, 7, 9, 2, 5, 10},
		3,
		84,
	},

	// 可以有多个 testcase
}

func Test_maxSumAfterPartitioning(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, maxSumAfterPartitioning(tc.A, tc.K), "输入:%v", tc)
	}
}

func Benchmark_maxSumAfterPartitioning(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			maxSumAfterPartitioning(tc.A, tc.K)
		}
	}
}