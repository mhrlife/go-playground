sample code : 

```
	json := `
{
	"from":{
		"user":{
			"name":"Mohammad",
			"lastName":"Hoseini Rad"
		},
		"arr":[
			[
				"a","b"
			],
			[
				["c"],["d"]
			]
		]
	}
}
`
	decoded := JSON.JsonDecode(json)

	//searching in objects
	fmt.Println( JSON.WalkObjStr(decoded,[]string{"from","user","name"}) )


	fmt.Println( JSON.WalkObjArr(decoded,[]string{"from","arr"}) )

	//searching in array

	fmt.Println( JSON.WalkArrStr(JSON.WalkObjArr(decoded,[]string{"from","arr"}),[]int{-1,0,0}) )```
