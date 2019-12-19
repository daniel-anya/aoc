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
  const wireOneInput = dataInput[0].split(",");

  //   const wireOneInput = "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51";
  //   const wireTwoInput = "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7";

  const wireTwoInput = dataInput[1].split(",");

  const wireOnePath = new WireTemplate(wireOneInput);
  const wireTwoPath = new WireTemplate(wireTwoInput);

  function findIntersections(wireOne, wireTwo) {
    const wireOneKeys = Array.from(wireOne.position.keys());
    const wireTwoKeys = Array.from(wireTwo.position.keys());
    return wireOneKeys.filter(val => wireTwoKeys.includes(val));
  }

  function _minimumIntersection(intersections) {
    const summedIntersections = intersections.map(pos => {
      let split = pos.split(",");
      let splitArray = split.map(val => parseInt(val));
      return splitArray.reduce((acc, initVal) => {
        return Math.abs(acc) + Math.abs(initVal);
      }, 0);
    });
    console.log(intersections);
    console.log(summedIntersections);
    return Math.min(...summedIntersections);
  }
  return {
    partOneAnswer: _minimumIntersection(
      findIntersections(wireOnePath, wireTwoPath)
    )
  };
})(input);

console.log("Part One answer is ", crossedWires.partOneAnswer);
