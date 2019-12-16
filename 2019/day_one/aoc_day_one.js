const fs = require('fs')

const filePath = String(__dirname + '/input.txt')
const input = fs.readFileSync(filePath).toString().split('\n').map(value => parseInt(value))

const calculateFuelRequirementFromMass = function(mass){
    return Math.floor(mass/3) - 2
}

const calculateTotalFuelRequirement = function(masses, callback){
    return masses.reduce(function(accumulator, currentValue) {
        return accumulator + callback(currentValue)
      }, 0)
}

console.log(calculateTotalFuelRequirement(input, calculateFuelRequirementFromMass))

