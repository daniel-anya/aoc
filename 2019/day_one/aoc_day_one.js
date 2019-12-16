const fs = require('fs')

const filePath = String(__dirname + '/input.txt')
const input = fs.readFileSync(filePath).toString().split('\n').map(value => parseInt(value))

const calculateStaticFuelRequirement = function(mass){
    return Math.floor(mass/3) - 2
}

const calculateDynamicFuelRequirement = function(mass){
    let fuelRequirementAccumulator = 0
    let fuelRequirementRemnant = mass
    while (true){
        fuelRequirementRemnant = Math.floor(fuelRequirementRemnant/3) - 2
        if (fuelRequirementRemnant > 0){
            fuelRequirementAccumulator =  fuelRequirementAccumulator + fuelRequirementRemnant
        }else{
            break
        }
    }
    return fuelRequirementAccumulator
}

const calculateTotalFuelRequirement = function(masses, callback){
    return masses.reduce(function(accumulator, currentValue) {
        return accumulator + callback(currentValue)
      }, 0)
}

// part one
console.log(calculateTotalFuelRequirement(input, calculateStaticFuelRequirement))

// part two
console.log(calculateTotalFuelRequirement(input, calculateDynamicFuelRequirement))



