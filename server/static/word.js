var jsondata

document.getElementById("search").addEventListener("click", e => {
	searchdata()
})

function searchdata() {
	let w = document.getElementById("word").value
	// alert(w)
	fetch(`/json?word=${w}`).then(res => res.json()).then(res => {
		jsondata = res
		let ul = document.getElementById("data")
		ul.innerHTML=""
		// ul.removeChild(document.getElementsByTagName("li"))
		res.forEach(e => {
			let li = document.createElement("li")
			li.innerHTML=e
			ul.appendChild(li)
		});
	})
}

document.getElementById("filter").addEventListener("click", e=>{
	let filter = document.getElementById("filterword").value
	console.log(filter)
	let filterarray = filter.split(" ")
	console.log(filterarray)
	let result = jsondata
	console.log(result)
	let num = document.getElementById("num").value
	filterarray.forEach(element => {
		let fields = element.split(":")
		console.log(fields)
		let i = fields[0]
		let v = fields[1]
		result = result.filter(word => {
			return word[i] == v &&(!num || num && word.length == num)
		})
	});
	let ul = document.getElementById("filterdata")
	ul.innerHTML = ""
	result.forEach(e => {
		let li = document.createElement("li")
		li.innerHTML = e
		ul.appendChild(li)
	});
})