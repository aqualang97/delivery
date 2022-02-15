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


const main = async ()=>{

    const response = await fetch('http://localhost:8080/all-products', {
        method: 'GET',

    });
    await delay(1000)
    const json  = await response.json();
    for (let prod in json){

        let p = document.createElement('p');
        let price = document.createElement('p');

        let img = document.createElement('img');
        p.setAttribute("id", `products${json[prod].ID}`);
        price.setAttribute("id", `products${json[prod].ID}`);
        img.src
        console.log(json[prod].Name);
        p.innerHTML = `Number ${json[prod].ID};     Name:   ${json[prod].name}`;
        img.src = json[prod].image;

        img.style.width = "20%"
        document.body.append(p);
        document.body.append(img);
        price.innerHTML = `Price: ${json[prod].price}`
        document.body.append(price);

        await delay(100)
    }

}
main()

const delay = (ms) => new Promise(resolve => {
    setTimeout(()=>{
        resolve();
    }, ms);
});