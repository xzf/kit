package combination

func GetAllIndexArray(all int, pick int) [][]int {
	return getAllIndexArray(all, pick, nil)
}

func getAllIndexArray(all int, pick int, in []int) (out [][]int) {
	nextPick := getAllNextPick(all, pick, in)
	if len(nextPick) == all-pick {
		return [][]int{in}
	}
	for _, item := range nextPick {
		out = append(out, getAllIndexArray(all, pick, append(append([]int{}, in...), item))...)
	}
	return
}

func getAllNextPick(all int, pick int, in []int) (out []int) {
	if len(in) == 0 {
		for item := 0; item < all; item++ {
			out = append(out, item)
		}
		return
	}
	if len(in) == all-pick {
		return
	}
	have := map[int]struct{}{}
	for index := 0; index < len(in); index++ {
		have[in[index]] = struct{}{}
	}
	for item := 0; item < all; item++ {
		_, ok := have[item]
		if ok {
			continue
		}
		out = append(out, item)
	}
	return
}
