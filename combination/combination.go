package combination

func GetAllIndexArray(all int, pick int) [][]int {
	return getAllIndexArray(getAllIndexArrayReq{
		in:   nil,
		all:  all,
		pick: pick,
	})
}

type getAllIndexArrayReq struct {
	in   []int
	all  int
	pick int
}

func getAllIndexArray(req getAllIndexArrayReq) (out [][]int) {
	nextPick := getAllNextPick(req.all, []int{})
	if len(nextPick) == req.all-req.pick {
		return [][]int{req.in}
	}
	for _, item := range nextPick {
		out = append(out,
			getAllIndexArray(getAllIndexArrayReq{
				in:   append(append([]int{}, req.in...), item),
				all:  req.all,
				pick: req.pick,
			})...)
	}
	return
}

func getAllNextPick(all int, in []int) (out []int) {
	if len(in) == 0 {
		for i := 1; i <= all; i++ {
			out = append(out, i)
		}
		return
	}
	if len(in) == all {
		return
	}
	have := map[int]struct{}{}
	for i := 0; i < len(in); i++ {
		have[in[i]] = struct{}{}
	}
	for item := 1; item <= all; item++ {
		_, ok := have[item]
		if ok {
			continue
		}
		out = append(out, item)
	}
	return
}
