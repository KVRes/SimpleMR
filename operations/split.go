package operations

type Splitter interface {
	SplitDataIntoMPieces(data any, m int) []any
}

type Master struct {
	state      [][]MasterCell // M * R
	splitter   Splitter
	mapTracker []int
}

func (master *Master) Start(data any, n int, m, r int) {
	master.state = make([][]MasterCell, m)
	for i := 0; i < m; i++ {
		master.state[i] = make([]MasterCell, r)
	}
	splitted := master.splitter.SplitDataIntoMPieces(data, n)

}

func (master *Master) IsMapFinished(mapId int) bool {
	if len(master.state) == 0 {
		return true
	}
	col := len(master.state[mapId])

	for i := 0; i < col; i++ {
		if !master.state[mapId][i].MapOk {
			return false
		}
	}
	return true
}

type MasterCell struct {
	MapOk   bool
	MapList []any
}

func Split(data any, m int) [][]any {
	// Split data into m parts
	//l := len(data)
	//
	//if l == 0 {
	//	return [][]any{}
	//}
	// TODO: Split data into m parts
	return nil
}
