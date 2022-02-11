// const Alex ={name: "Alex", passport: {num:"11111111", e: new Error("Alex error"), d: Date()}};
// const deepClone = (obj) => {
//     if (obj ===null){
//         return null
//     }
//     const clone = {};
//     Object.keys(obj).forEach(item=>{
//         if (obj[item] instanceof Object){
//             clone[item] = deepClone(obj[item]);
//
//         }else{
//             clone[item] = obj[item];
//         }
//
//     });
//     return clone
// }
//
//
//         // if (!(typeof obj === 'object')) {
//         //     throw new TypeError("Not object");
//         //
//         // }else if(obj === null){
//         //     throw new RangeError("obj is null");
//         //
//         // } else {
//         //     return clone
//         // }
//     //}
//
//
//
// const Yurii = deepClone(Alex)
// Yurii.name = "Yurii"
// Yurii.passport.num = "12312312"
// Yurii.passport.e = new Error("Yurii err")
// Yurii.passport.d = Date().sub()
// console.log(Alex)
// console.log(Yurii)
//


const foo = async ()=>{

    const response = await fetch('http://localhost:8080/all-products', {
        method: 'GET',

    });
    await delay(1000).then(()=>{
    })
    const json  = await response.json();
    for (let prod in json){

        let p = document.createElement('p');
        p.setAttribute("id", `products${json[prod].ID}`);

        console.log(json[prod].Name);
        p.innerHTML = `Number ${json[prod].ID};     Name:   ${json[prod].Name}`;
        document.body.append(p);
        await delay(100).then(()=>{
        })
    }

}
foo()

const delay = (ms) => new Promise(resolve => {
    setTimeout(()=>{
        resolve();
    }, ms);
});