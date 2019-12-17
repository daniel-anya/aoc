const fs = require('fs')

const filePath = String(__dirname + '/input.txt')
const input = fs.readFileSync(filePath).toString().split(',').map(value => parseInt(value))

const programAlarmLogic = (function(inputArray) {

inputArray[1] = 12
inputArray[2] = 2

i = 0

while (true){
if (inputArray[i] === 99){
    break
}
else if (inputArray[i] === 1) {
inputArray[inputArray[i+3]] = inputArray[inputArray[i + 1]] + inputArray[inputArray[i + 2]]

}
else if (inputArray[i] === 2) {
    inputArray[input[i+3]] = inputArray[input[i + 1]] * inputArray[inputArray[i + 2]]
    
}
i = i + 4
}
return inputArray[0]
    
})(input)

//Part One
console.log("The new initial value is %d", programAlarmLogic)