const fs = require("fs");

const filePath = String(__dirname + "/input.txt");
const input = fs
  .readFileSync(filePath)
  .toString()
  .split("\n");

class WireTemplate {
  constructor(input) {
    this.directions = input;
    this.moves = 0;
    this.position = new Map();
    this.definePath();
  }

  definePath() {
    let x = 0;
    let y = 0;
    for (let dir of this.directions) {
      if (dir[0] === "R") {
        let numOfMoves = parseInt(dir.slice(1, dir.length));
        for (let _ of Array(numOfMoves).keys()) {
          x += 1;
          this.moves += 1;
          this.position.set(`${x}` + "," + `${y}`, this.moves);
        }
      } else if (dir[0] === "L") {
        let numOfMoves = parseInt(dir.slice(1, dir.length));
        for (let _ of Array(numOfMoves).keys()) {
          x -= 1;
          this.moves += 1;
          this.position.set(`${x}` + "," + `${y}`, this.moves);
        }
      } else if (dir[0] === "U") {
        let numOfMoves = parseInt(dir.slice(1, dir.length));
        for (let _ of Array(numOfMoves).keys()) {
          y += 1;
          this.moves += 1;
          this.position.set(`${x}` + "," + `${y}`, this.moves);
        }
      } else if (dir[0] === "D") {
        let numOfMoves = parseInt(dir.slice(1, dir.length));
        for (let _ of Array(numOfMoves).keys()) {
          y -= 1;
          this.moves += 1;
          this.position.set(`${x}` + "," + `${y}`, this.moves);
        }
      }
    }
  }
}
const crossedWires = (function(dataInput) {
  const intersections = {};
  const wireOneInput = dataInput[0].split(",");
  const wireTwoInput = dataInput[1].split(",");

  const wireOnePath = new WireTemplate(wireOneInput);
  const wireTwoPath = new WireTemplate(wireTwoInput);

  function _findIntersections(wireOne, wireTwo) {
    const wireOneKeys = Array.from(wireOne.position.keys());
    const wireTwoKeys = Array.from(wireTwo.position.keys());
    intersections["key"] = wireOneKeys.filter(val => wireTwoKeys.includes(val));
  }

  function _minimumIntersectionPoint(intersections) {
    const summedIntersections = intersections.map(pos => {
      let split = pos.split(",");
      const splitArray = split.map(val => parseInt(val));
      return splitArray.reduce((acc, initVal) => {
        return Math.abs(acc) + Math.abs(initVal);
      }, 0);
    });

    return Math.min(...summedIntersections);
  }
  function _minimumIntersectionSteps(intersections, wireOne, wireTwo) {
    // const intersections = _findIntersections(wireOne, wireTwo);
    const steps = intersections.map(intersection => {
      return (
        wireOne.position.get(intersection) + wireTwo.position.get(intersection)
      );
    });

    return Math.min(...steps);
  }
  _findIntersections(wireOnePath, wireTwoPath);
  return {
    partOneAnswer: _minimumIntersectionPoint(intersections["key"]),
    partTwoAnswer: _minimumIntersectionSteps(
      intersections["key"],
      wireOnePath,
      wireTwoPath
    )
  };
})(input);

console.log("Part One answer is", crossedWires.partOneAnswer);
console.log("Part Two answer is", crossedWires.partTwoAnswer);
