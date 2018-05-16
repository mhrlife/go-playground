package JSON

import (
	"encoding/json"
)



func JsonDecode(str string) map[string]interface{} {
	var json_decoded map[string]interface{}
	json.Unmarshal([]byte(str), &json_decoded)
	return json_decoded
}

func WalkObject(m map[string]interface{}, data []string) interface{}  {
	current := m
	var final interface{}
	for i,v := range data {

		if i == (len(data) -1) {
			final = current[v]
			break
		}
		current = current[v].(map[string]interface{})
	}

	return  final
}

func WalkObjStr(m map[string]interface{}, data []string) string  {
	return WalkObject(m, data).(string)
}

func WalkObjFloat64(m map[string]interface{}, data []string) float64  {
	return WalkObject(m, data).(float64)
}
func WalkObjInt32(m map[string]interface{}, data []string) int32  {
	return WalkObject(m, data).(int32)
}
func WalkObjArr(m map[string]interface{}, data []string) []interface{}  {
	return WalkObject(m, data).([]interface{})
}

