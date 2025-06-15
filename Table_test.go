package main

import (
	"reflect"
	"testing"
)

func TestProcChapterContent(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{
			name: "1",
			arg:  "# 第四~五章 妊娠生理与妊娠诊断\r\n\r\n妊娠期心排出量增加**30%**，心脏容量增加**10%**，血容量增加**40%~45%**\r\n\r\n妊娠|期|母体**\r\n32~3|4周**心排出|量达到最大。\r\n\r\n妊娠期内分泌系统变化：**二高一少一不变**——PRL、皮质醇升高，促性腺激素减少，游离甲状腺激素不变\r\n\r\n胎儿吞咽可以使羊水量趋于平衡\r\n\r\n母体与羊水的交换主要通过胎膜，同时其还具有缓冲、防止压迫和撞击的作用。\r\n\r\n子宫各部是**不对称**的球形\r\n\r\n妊娠晚期子宫容量约为**5000ml**\r\n\r\n妊娠晚期子宫**轻度右旋**，与乙状结肠在左有关\r\n\r\n**12周**出现不规律宫缩\r\n\r\n孕妇需要蛋白质最高的时期是**孕晚期**：孕早期基础代谢率稍下降，于孕中期逐渐增高，至孕晚期可增高**15~20%**，相应的蛋白质需要量在孕中期（4~6个月）每日增加**15g**，孕晚期（7~9个月）是需要蛋白质的最高时期，每日增加**25g**\r\n\r\n> 胎儿越大，需求越多\r\n\r\n孕妇于妊娠后期常有踝部及小腿下半部轻度水肿，但**经休息后可消退**。若休息后未能消退，应考虑妊高症、合并肾脏疾病或其他合并症。\r\n\r\n母亲孕期吸烟对母亲或胎儿的影响：**自然流产率增加**、**胎儿死亡率增加**、**低体重儿**、**胎儿畸形**，不包括新生儿死亡。\r\n\r\n### 女性不同时期雌、孕激素生成部位\r\n\r\n时期|雌激素|孕激素\r\n--|--|--\r\n排卵前|卵泡<br>膜细胞、颗粒细胞|颗粒细胞\r\n排卵后至孕10周前|黄体细胞|黄体细胞\r\n孕10周后（胎盘接管黄体）|胎盘单位（主要分泌E3）|胎盘合体滋养细胞\r\n\r\n至妊娠末期，E3值为非孕妇女的**1000倍**，E2及E1值为非孕妇女的**100倍**。\r\n\r\n胎儿肺泡表面活性物质**18~20周**开始产生，**28周**出现在羊水里，**35~36周**迅速增高至成熟水平。\r\n\r\n> 28周是早产儿的分界标志。\r\n\r\n妊娠期子宫超声改变：**5囊6芽8胎心**——5周妊娠囊，6周胎芽，8周原始心管搏动\r\n\r\n早期妊娠的确诊依据是：**B型超声检查**\r\n\r\n妊娠最早的症状：**停经**\r\n\r\n黑加征：停经**6~8周**时，双合诊检查子宫峡部极软，感觉宫体与宫颈似不相连，称黑加征，是妊娠早期**最特异**的症状。\r\n\r\n早孕反应：6周开始，12周后消失\r\n\r\n妊娠**8周**子宫增大到非孕时的2倍，**12周**到非孕时的3倍\r\n\r\n胎心音**110~160次/分**，子宫杂音和**母体心率**一致，脐带杂音和**胎心音**一致。\r\n\r\n胎姿势——胎儿在子宫内的姿势  \r\n胎产式——胎体纵轴与母体纵轴的关系  \r\n胎先露——最先进入骨盆入口的胎儿部分  \r\n胎方位——胎先露的指示点与母体骨盆的关系  \r\n骨盆轴——连接骨盆各平面中点的假想曲线\r\n\r\n我国现阶段采用的围生期规定为**从妊娠满28周至产后1周**。\r\n\r\n### 胎心音位置\r\n\r\n- 头先露时多在脐下\r\n- 臀先露时多在脐上\r\n- 肩先露时多在脐周\r\n\r\n> 胎头圆而硬，有浮球感  \r\n> 胎背宽而平坦  \r\n> 胎臀宽而软，形状不规则  \r\n> 胎儿肢体小且有不规则活动  \r\n> 24周后可区分胎头胎背\r\n\r\n### 妊娠周数的宫底位置\r\n\r\n周数|位置\r\n--|--\r\n12周末|耻骨联合上2~3横指\r\n16周末|脐耻之间\r\n20周末|脐下1横指\r\n22周末|平脐\r\n24周末|脐上1横指\r\n28周末|脐上3横指\r\n32周末|脐与剑突之间\r\n36周末|剑突下2横指\r\n40周末|脐与剑突之间或略高\r\n\r\n### 产后宫体肌纤维缩复\r\n\r\n- 产后1天——平脐\r\n- 产后1周——孕3月大小（耻骨联合）\r\n- 产后10天——盆腔内\r\n- 产后6周——非孕大小\r\n\r\n> 一日平脐周耻三，十日入盆六周还\r\n\r\n### 妊娠全过程汇总\r\n\r\n妊娠周数（末次月经起算）|生理变化\r\n--|--\r\n2周|受精卵结合于输卵管壶腹部，此时雌、孕激素来源于黄体细胞\r\n3周|受精卵着床\r\n5周|超声观察到妊娠囊\r\n6周|在此之前可使用尿妊娠试验，早孕反应出现，超声能够观察到胎芽，6~8周出现黑加征\r\n8周|羊水量5~10ml，子宫增大到非孕时的2倍，早孕反应消失，超声能够观察到心管搏动，8~10周血清hCG达到高峰\r\n10周|羊水量30ml，此时胎盘替代黄体功能，雌激素来源于胎盘单位，孕激素来源于胎盘合体滋养层细胞\r\n12周|子宫增大到非孕时的3倍，子宫峡部逐渐变软伸展扩展为宫腔的一部分，出现不规律宫缩\r\n16周|16~20周可自觉胎动\r\n18周|18~20周胎儿肺泡表面活性物质开始产生\r\n20周|羊水量400ml，可听诊胎心，触及胎体\r\n24周|可区分头背臀\r\n28周|胎儿肺泡表面活性物质出现在羊水里，此时为早产儿的分界线，我国采用的围生期规定为从妊娠满28周至产后1周\r\n32周|32~34周左侧卧位心排量达到顶峰约30%\r\n35周|35~36周羊水中的肺泡表面活性物质迅速增高至成熟水平\r\n38周|羊水量1000ml（最多），\r\n40周|羊水量800ml（足月），\r\n42周|羊水量300ml（过期），\r\n",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProcChapterContent(tt.arg); got != tt.want {
				t.Errorf("ProcChapter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_trimTunnelSpace(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{
			name: "1",
			arg:  "|a|b|c|",
			want: "a|b|c",
		},
		{
			name: "2",
			arg:  "|     a|b|c  |    ",
			want: "a|b|c",
		},
		{
			name: "3",
			arg:  "       a|b|c       ",
			want: "a|b|c",
		},
		{
			name: "4",
			arg:  "|a|b|c",
			want: "a|b|c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trimTunnelSpace(tt.arg); got != tt.want {
				t.Errorf("trimTunnelSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_right2ndLine(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want bool
	}{
		{
			name: "1",
			arg:  "|--|--|--|",
			want: true,
		},
		{
			name: "2",
			arg:  "--|--|--",
			want: true,
		},
		{
			name: "3",
			arg:  "|||||||",
			want: false,
		},
		{
			name: "4",
			arg:  "--------",
			want: false,
		},
		{
			name: "5",
			arg:  "",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsTunnelHyphen(tt.arg); got != tt.want {
				t.Errorf("right2ndLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitTableLine(t *testing.T) {
	tests := []struct {
		name string
		args string
		want []string
	}{
		{
			name: "1",
			args: "a|b|c|d",
			want: []string{"a", "b", "c", "d"},
		},
		{
			name: "2",
			args: "a | b| c |     d",
			want: []string{"a", "b", "c", "d"},
		},
		{
			name: "3",
			args: "a",
			want: []string{"a"},
		},
		{
			name: "4",
			args: "",
			want: []string{""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitTableLine(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitTableLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
