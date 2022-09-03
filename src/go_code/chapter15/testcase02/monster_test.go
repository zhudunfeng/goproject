package monster

import "testing"

func TestStore(t *testing.T) {
	monster := &Monster{
		Name:  "红孩儿~~",
		Age:   18,
		Skill: "吐火",
	}

	res := monster.Store()
	if !res {
		t.Fatalf("monster.Store 错误,希望为=%v 实际为=%v\n", true, res)
	}
	t.Logf("monster.Store 测试成功")
}

func TestReStore(t *testing.T) {
	var monster = &Monster{}
	res := monster.Restore()
	if !res {
		t.Fatalf("monster.Restore错误，希望为=%v 实际为=%v\n", true, res)
	}
	if monster.Name != "红孩儿" {
		t.Fatalf("monster.Restore错误，希望为=%v 实际为=%v\n", "红孩儿", monster.Name)
	}
	t.Logf("monster.Restore测试成功\n%v\n", *monster)
}
