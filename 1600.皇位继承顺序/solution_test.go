package leetcode

// 1600.皇位继承顺序
// https://leetcode-cn.com/problems/throne-inheritance
type ThroneInheritance struct {
	king string
	sons map[string][]string
	dead map[string]bool
}

func Constructor(kingName string) ThroneInheritance {
	return ThroneInheritance{
		king: kingName,
		sons: map[string][]string{},
		dead: map[string]bool{},
	}
}

func (this *ThroneInheritance) Birth(parentName string, childName string) {
	this.sons[parentName] = append(this.sons[parentName], childName)
}

func (this *ThroneInheritance) Death(name string) {
	this.dead[name] = true
}

func (this *ThroneInheritance) GetInheritanceOrder() []string {
	var next func(string)
	ret := []string{}
	next = func(name string) {
		if !this.dead[name] {
			ret = append(ret, name)
		}
		for _, childName := range this.sons[name] {
			next(childName)
		}
	}
	next(this.king)
	return ret
}

/**
 * Your ThroneInheritance object will be instantiated and called as such:
 * obj := Constructor(kingName);
 * obj.Birth(parentName,childName);
 * obj.Death(name);
 * param_3 := obj.GetInheritanceOrder();
 */
// func TestSolution(t *testing.T) {
// 	testCases := []struct {

// 		want
// 		desc string
// 	}{
// 		{
//             want: ,
// 		},
// 	}
// 	for _, tC := range testCases {
// 		t.Run(tC.desc, func(t *testing.T) {
// 			get :=
// 			if !reflect.DeepEqual(tC.want,get){
// 				t.Errorf("input: %+v get: %v\n",tC,get)
// 			}
// 		})
// 	}
// }
