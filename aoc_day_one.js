
// const calculateFuelRequirementFromMass = function(mass){
//     return Math.floor(mass/3) - 2
// }

// const calculateTotalFuelRequirement = function(masses, callback){
//     masses.reduce(function(accumulator, currentValue, currentIndex) {
//         return accumulator + callback(currentValue)
//       }, 0)
// }

const fetch = require('node-fetch')
fetch('https://adventofcode.com/2019/day/1/input')
    .then(res => console.log(res))
    // .then(json => console.log(json))



