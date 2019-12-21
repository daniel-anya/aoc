const secureContainer = (function(input) {
  function _generateInputRange(input) {
    const stringBand = input.split("-");
    const intBand = stringBand.map(val => parseInt(val));

    const intRange = [];

    for (let i = intBand[0]; i < intBand[intBand.length - 1]; i++) {
      intRange.push(i + 1);
    }
    return intRange;
  }

  function _filterAdjacentDigits(intInput, partTwoCriteria = false) {
    const repeatedDoubleDigits = [
      "00",
      "11",
      "22",
      "33",
      "44",
      "55",
      "66",
      "77",
      "88",
      "99"
    ];
    const stringInput = intInput.toString();
    for (let val of repeatedDoubleDigits) {
      if (!partTwoCriteria && stringInput.includes(val)) {
        return true;
      } else {
        if (stringInput.includes(val) && !stringInput.includes(val + val[0])) {
          return true;
        }
      }
    }
    return false;
  }

  function _filterIncreasingDigits(intInput) {
    let state = true;
    const intInputArray = String(intInput).split("");
    for (let i = 0; i < intInputArray.length - 1; i++) {
      if (parseInt(intInputArray[i]) > parseInt(intInputArray[i + 1])) {
        state = false;
      }
    }
    return state;
  }

  function solvePuzzle(partTwoCriteria) {
    const intInputs = _generateInputRange(input);
    const adjacentDigits = intInputs.filter(val =>
      _filterAdjacentDigits(val, (partTwoCriteria = partTwoCriteria))
    );
    const increasingDigits = adjacentDigits.filter(val =>
      _filterIncreasingDigits(val)
    );
    return increasingDigits.length;
  }

  return {
    partOneAnswer: solvePuzzle(),
    partTwoAnswer: solvePuzzle((partTwoCriteria = true))
  };
})("130254-678275");

//part one answer
console.log(secureContainer.partOneAnswer);

// part two answer
console.log(secureContainer.partTwoAnswer);
