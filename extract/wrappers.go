package extract

// TODO: add documentation

type WrapperMap struct {
	Map map[string]string
}

func (w *WrapperMap) Get(key string) string {
	return w.Map[key]
}

type WrapperMapSlice struct {
	Map map[string][]string
}

func (ws *WrapperMapSlice) Get(key string) string {
	v := ws.Map[key]
	if len(v) == 0 {
		return ""
	}

	return v[0]
}

type WrapperMapSliceLast struct {
	Map map[string][]string
}

func (wsl *WrapperMapSliceLast) Get(key string) string {
	v := wsl.Map[key]
	if len(v) == 0 {
		return ""
	}

	return v[len(v)-1]
}

// Wrap functions

func Wrap(m map[string]string) StringGetter {
	return &WrapperMap{Map: m}
}

func WrapSlice(m map[string][]string, last bool) StringGetter {
	if last {
		return &WrapperMapSliceLast{Map: m}
	}
	return &WrapperMapSlice{Map: m}
}
