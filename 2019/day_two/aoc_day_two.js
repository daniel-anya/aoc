const fs = require("fs");

const filePath = String(__dirname + "/input.txt");
const input = fs
  .readFileSync(filePath)
  .toString()
  .split(",")
  .map(value => parseInt(value));

const programAlarmLogic = (function(inputArray) {
  function findOutput(noun, verb) {
    const clone = [...inputArray];
    clone[1] = noun;
    clone[2] = verb;

    i = 0;

    while (true) {
      if (clone[i] === 99) {
        break;
      } else if (clone[i] === 1) {
        clone[clone[i + 3]] = clone[clone[i + 1]] + clone[clone[i + 2]];
      } else if (clone[i] === 2) {
        clone[input[i + 3]] = clone[input[i + 1]] * clone[clone[i + 2]];
      }
      i = i + 4;
    }
    return clone[0];
  }

  function findInputs(finalAnswer) {
    for (let noun of [...Array(100).keys()]) {
      for (let verb of [...Array(100).keys()]) {
        if (findOutput(noun, verb) === finalAnswer) {
          return 100 * noun + verb;
        }
      }
    }

    return "No results found";
  }

  return {
    partOneAnswer: findOutput(12, 2),
    partTwoAnswer: findInputs(19690720)
  };
})(input);

//Part One Answer
console.log("The output is %d", programAlarmLogic.partOneAnswer);

//Part Two Answer
console.log("The input is", programAlarmLogic.partTwoAnswer);
