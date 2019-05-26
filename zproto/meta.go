package zproto

type KeyVer struct {
	Id   int //in msg
	Ver  int
	Rm   int //0 raw, 1 msg, if array Type+1
	Type int //details
}

type Item struct {
	Type       int    //enum, msg, item
	Key        string //Var
	Repeated   bool
	Val        string //Value or Type(raw,msg)
	Ver        KeyVer //@
	Comment    string //#
	CommentPre string //
}

type DTypeInfo struct {
	Id   int
	Size int
}

var TargetRawType map[string][]string
var RawTypeIndex = []string{"bool", "i8", "i16", "i32", "i64",
	"u8", "u16", "u32", "u64", "string", "bytes"}
var dTypePool map[string]DTypeInfo
var rawTypeMax = 0
var dTypeIndex = 0

func TargetRawInit() {
	TargetRawType = make(map[string][]string)
	TargetRawType["go"] = []string{"bool", "int8", "int16", "int32", "int64",
		"uint8", "uint16", "uint32", "uint64", "string", "bytes"}
	TargetRawType["cs"] = []string{"Boolean", "SByte", "Int16", "Int32", "Int64",
		"Byte", "UInt16", "UInt32", "UInt64", "String", "Byte[]"}
}

func DTypeInit() {
	dTypePool = make(map[string]DTypeInfo)
	for id, val := range RawTypeIndex {
		size := -1
		if id == 0 {
			size = 1
		} else if id >= 1 && id <= 4 {
			size = 2 ^ (id - 1)
		} else if id >= 5 && id <= 8 {
			size = 2 ^ (id - 5)
		}
		dTypePool[val] = DTypeInfo{id, size}
	}
	rawTypeMax = len(RawTypeIndex)
	dTypeIndex = rawTypeMax

	TargetRawInit()
}

func TypeSize(t string) int {
	tr, ok := dTypePool[t]
	if ok == false {
		return 0
	}
	return tr.Size
}

func IsRawType(t string) bool {
	tr, ok := dTypePool[t]
	if ok == false {
		return false
	}
	if tr.Size < rawTypeMax {
		return true
	}
	return false
}

func AddNewDType(t string) {
	dTypeIndex++
	dTypePool[t] = DTypeInfo{dTypeIndex, -1}
}
