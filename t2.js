var str="abataxios.Get("https://jsonplaceholder.typicode.com/users").then((res)=>{
	for(let i=0;i<res.data.length();i++){
			emp=res.data[i]
			if(emp.address.zipcode.contains("7")){
				op.push(emp)
			}
		}	
	})
	"
let isNotPlallindrome=false
for(let i=0;i<str.length;i++){
	let n=str.length
	if(str[i]!=str[n-1-i]){
		isNotPlallindrome=true
	}
}
if(isNotPlallindrome){
	console.log("Not Plalindrome, Input Str:", str)
}else{
	console.log("Plalindrome, Input Str:", str)
}
// if(str==str.reverse()){
// 	console.log("Pallindrome")
// }